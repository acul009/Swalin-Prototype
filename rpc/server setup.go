package rpc

import (
	"context"
	"fmt"
	"io"
	"log"
	"rahnit-rmm/config"
	"rahnit-rmm/pki"

	"github.com/quic-go/quic-go"
)

func WaitForServerSetup(listenAddr string) error {
	ok, err := pki.CurrentAvailable()
	if err != nil {
		return fmt.Errorf("failed to check if current cert exists: %w", err)
	}

	if ok {
		// Server already initialized
		return nil
	}
	tlsConf, err := GetTlsServerConfig([]TlsConnectionProto{ProtoServerInit})
	if err != nil {
		return fmt.Errorf("error getting server tls config: %w", err)
	}

	quicConf := &quic.Config{}
	listener, err := quic.ListenAddr(listenAddr, tlsConf, quicConf)
	if err != nil {
		return fmt.Errorf("error creating QUIC server: %w", err)
	}

	initNonceStorage = NewNonceStorage()

	for {
		conn, err := listener.Accept(context.Background())
		if err != nil {
			err := fmt.Errorf("error accepting QUIC connection: %w", err)
			log.Println(err)
			continue
		}

		log.Printf("Connection accepted")

		err = acceptServerInitialization(conn)
		if err != nil {
			err := fmt.Errorf("error initializing server: %w", err)
			log.Println(err)
		} else {
			// no error, initialization was successful
			err = listener.Close()
			if err != nil {
				return fmt.Errorf("error closing listener: %w", err)
			}
			return nil
		}
	}
}

var initNonceStorage *nonceStorage = nil

type serverInitRequest struct {
	ServerPubKey string
}

type serverInitResponse struct {
	RootCa     string
	ServerCert string
}

func acceptServerInitialization(quicConn quic.Connection) error {
	conn := newRpcConnection(quicConn, nil, RpcRoleInit, initNonceStorage, nil, ProtoServerInit)

	log.Printf("Opening init QUIC stream...")

	session, err := conn.AcceptSession(context.Background())
	if err != nil {
		return fmt.Errorf("error opening QUIC stream: %w", err)
	}

	err = session.mutateState(RpcSessionCreated, RpcSessionOpen)
	if err != nil {
		return fmt.Errorf("error mutating session state: %w", err)
	}

	log.Printf("Session opened, reading public key...")

	pubRootMarshalled := ""
	sender, err := pki.ReadAndUnmarshalAndVerify(session, &pubRootMarshalled)
	if err != nil {
		return fmt.Errorf("error reading public key: %w", err)
	}

	pubRoot, err := pki.DecodePubFromString(pubRootMarshalled)
	if err != nil {
		return fmt.Errorf("error decoding public key: %w", err)
	}

	if !sender.Equal(pubRoot) {
		return fmt.Errorf("root public key does not match sender")
	}

	session.partner = pubRoot

	log.Printf("preparing init request...")

	pubMe, err := pki.GetCurrentPublicKey()
	if err != nil {
		return fmt.Errorf("error getting current public key: %w", err)
	}

	pubMeMarshalled, err := pki.EncodePubToString(pubMe)
	if err != nil {
		return fmt.Errorf("error marshalling public key: %w", err)
	}

	initRequest := &serverInitRequest{
		ServerPubKey: pubMeMarshalled,
	}

	log.Printf("Sending init request...")

	err = WriteMessage[*serverInitRequest](session, initRequest)
	if err != nil {
		return fmt.Errorf("error writing message: %w", err)
	}

	log.Printf("Awaiting init response...")

	response := &serverInitResponse{}
	sender, err = readMessageFromUnknown[*serverInitResponse](session, response)
	if err != nil {
		return fmt.Errorf("error reading message: %w", err)
	}

	log.Printf("Init response received")

	if !pubRoot.Equal(sender) {
		return fmt.Errorf("root public key does not match sender")
	}

	rootCert, err := pki.DecodeCertificate([]byte(response.RootCa))
	if err != nil {
		return fmt.Errorf("error decoding root certificate: %w", err)
	}

	if !pubRoot.Equal(rootCert.PublicKey) {
		return fmt.Errorf("root public key does not match certificate")
	}

	serverCert, err := pki.DecodeCertificate([]byte(response.ServerCert))
	if err != nil {
		return fmt.Errorf("error decoding server certificate: %w", err)
	}

	err = pki.SaveCurrentCert(serverCert)
	if err != nil {
		return fmt.Errorf("error saving server certificate: %w", err)
	}

	err = pki.SaveRootCert(rootCert)
	if err != nil {
		return fmt.Errorf("error saving root certificate: %w", err)
	}

	session.Close()
	conn.Close(200, "done")

	return nil
}

func SetupServer(addr string, rootPassword []byte, nameForServer string) error {
	err := pki.UnlockAsRoot(rootPassword)
	if err != nil {
		return fmt.Errorf("error unlocking root cert: %w", err)
	}

	tlsConf := GetTlsClientConfig(ProtoServerInit)

	quicConf := &quic.Config{}

	quicConn, err := quic.DialAddr(context.Background(), addr, tlsConf, quicConf)
	if err != nil {
		qErr, ok := err.(*quic.TransportError)
		if ok && uint8(qErr.ErrorCode) == 120 {
			return fmt.Errorf("server not in init mode: %w", err)
		}
		return fmt.Errorf("error creating QUIC connection: %w", err)
	}

	initNonceStorage = NewNonceStorage()

	conn := newRpcConnection(quicConn, nil, RpcRoleInit, initNonceStorage, nil, ProtoServerInit)

	log.Printf("Connection opened to %s\n", addr)

	session, err := conn.OpenSession(context.Background())
	if err != nil {
		return fmt.Errorf("error opening QUIC stream: %w", err)
	}

	err = session.mutateState(RpcSessionCreated, RpcSessionOpen)
	if err != nil {
		return fmt.Errorf("error mutating session state: %w", err)
	}

	log.Printf("Session opened, sending public key")

	pubKey, err := pki.GetCurrentPublicKey()
	if err != nil {
		return fmt.Errorf("error getting current public key: %w", err)
	}

	key, err := pki.GetCurrentKey()
	if err != nil {
		return fmt.Errorf("error getting current key: %w", err)
	}

	pubMarshalled, err := pki.EncodePubToString(pubKey)
	if err != nil {
		return fmt.Errorf("error marshalling public key: %w", err)
	}

	payload, err := pki.MarshalAndSign(pubMarshalled, key, pubKey)
	if err != nil {
		return fmt.Errorf("error marshalling message: %w", err)
	}

	n, err := session.Write(payload)
	if err != nil {
		return fmt.Errorf("error writing message: %w", err)
	}
	if n != len(payload) {
		return fmt.Errorf("error writing message: %w", io.ErrShortWrite)
	}

	log.Printf("reading initialization request^...")

	req := &serverInitRequest{}

	sender, err := readMessageFromUnknown[*serverInitRequest](session, req)
	if err != nil {
		return fmt.Errorf("error reading message: %w", err)
	}

	serverPubKey, err := pki.DecodePubFromString(req.ServerPubKey)
	if err != nil {
		return fmt.Errorf("error decoding server public key: %w", err)
	}

	if !sender.Equal(serverPubKey) {
		return fmt.Errorf("server public key does not match sender")
	}

	session.partner = sender

	log.Printf("Received request with pubkey: %s\n", req.ServerPubKey)

	serverHostCert, err := pki.CreateServerCertWithCurrent(nameForServer, serverPubKey)
	if err != nil {
		return fmt.Errorf("error creating server certificate: %w", err)
	}

	log.Printf("Created server certificate:\n%s\n\n", string(pki.EncodeCertificate(serverHostCert)))

	rootCert, err := pki.GetRootCert()
	if err != nil {
		return fmt.Errorf("error getting root certificate: %w", err)
	}

	response := &serverInitResponse{
		RootCa:     string(pki.EncodeCertificate(rootCert)),
		ServerCert: string(pki.EncodeCertificate(serverHostCert)),
	}

	err = WriteMessage[*serverInitResponse](session, response)
	if err != nil {
		return fmt.Errorf("error writing message: %w", err)
	}

	session.Close()
	conn.Close(200, "done")

	config.V().Set("upstream.address", addr)
	err = config.Save()
	if err != nil {
		return fmt.Errorf("error saving config: %w", err)
	}

	err = pki.SaveUpstreamCert(serverHostCert)
	if err != nil {
		return fmt.Errorf("error saving upstream certificate: %w", err)
	}

	return nil
}
