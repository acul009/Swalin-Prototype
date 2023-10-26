package pki

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"rahnit-rmm/util"
)

type PrivateKey ecdsa.PrivateKey

func (key *PrivateKey) MarshalJSON() ([]byte, error) {
	return nil, fmt.Errorf("cannot marshal private key")
}

func (key *PrivateKey) UnmarshalJSON(data []byte) error {
	return fmt.Errorf("cannot unmarshal private key")
}

func (key *PrivateKey) BinaryEncode(password []byte) ([]byte, error) {
	if password == nil {
		return nil, fmt.Errorf("password cannot be nil")
	}

	if len(password) == 0 {
		return nil, fmt.Errorf("password cannot be empty")
	}

	ecdsaKey := ecdsa.PrivateKey(*key)

	keyBytes, err := x509.MarshalECPrivateKey(&ecdsaKey)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal private key: %w", err)
	}

	encryptedBytes, err := util.EncryptDataWithPassword(password, keyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt private key: %w", err)
	}

	return encryptedBytes, nil

}

func PrivateKeyFromBinary(keyPEM []byte, password []byte) (*PrivateKey, error) {
	keyBytes, err := util.DecryptDataWithPassword(password, keyPEM)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt private key: %w", err)
	}

	// Parse the CA private key
	key, err := x509.ParseECPrivateKey(keyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse CA private key: %w", err)
	}

	return ImportPrivateKey(*key), nil
}

func (key *PrivateKey) PemEncode(password []byte) ([]byte, error) {
	encryptedBytes, err := key.BinaryEncode(password)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal private key: %w", err)
	}

	return pem.EncodeToMemory(
		&pem.Block{Type: "EC PRIVATE KEY",
			Bytes:   encryptedBytes,
			Headers: map[string]string{"Proc-Type": "4,ENCRYPTED", "DEK-Info": "AES-CFB"},
		},
	), nil
}

func PrivateKeyFromPem(keyPEM []byte, password []byte) (*PrivateKey, error) {
	// Decode the PEM-encoded CA private key
	block, _ := pem.Decode(keyPEM)
	if block == nil {
		return nil, fmt.Errorf("failed to decode CA private key PEM")
	}

	return PrivateKeyFromBinary(block.Bytes, password)
}

func (key *PrivateKey) GetPublicKey() *PublicKey {
	pub := PublicKey(key.PublicKey)
	return &pub
}

func ImportPrivateKey(key ecdsa.PrivateKey) *PrivateKey {
	keyTyped := PrivateKey(key)
	return &keyTyped
}

func (key *PrivateKey) ToEcdsa() *ecdsa.PrivateKey {
	keyTyped := ecdsa.PrivateKey(*key)
	return &keyTyped
}