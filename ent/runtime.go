// Code generated by ent, DO NOT EDIT.

package ent

import (
	"rahnit-rmm/ent/device"
	"rahnit-rmm/ent/revocation"
	"rahnit-rmm/ent/schema"
	"rahnit-rmm/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	deviceFields := schema.Device{}.Fields()
	_ = deviceFields
	// deviceDescPublicKey is the schema descriptor for public_key field.
	deviceDescPublicKey := deviceFields[0].Descriptor()
	// device.PublicKeyValidator is a validator for the "public_key" field. It is called by the builders before save.
	device.PublicKeyValidator = deviceDescPublicKey.Validators[0].(func(string) error)
	// deviceDescCertificate is the schema descriptor for certificate field.
	deviceDescCertificate := deviceFields[1].Descriptor()
	// device.CertificateValidator is a validator for the "certificate" field. It is called by the builders before save.
	device.CertificateValidator = deviceDescCertificate.Validators[0].(func(string) error)
	revocationFields := schema.Revocation{}.Fields()
	_ = revocationFields
	// revocationDescHash is the schema descriptor for hash field.
	revocationDescHash := revocationFields[1].Descriptor()
	// revocation.HashValidator is a validator for the "hash" field. It is called by the builders before save.
	revocation.HashValidator = revocationDescHash.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[0].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescPasswordDoubleHashed is the schema descriptor for password_double_hashed field.
	userDescPasswordDoubleHashed := userFields[3].Descriptor()
	// user.PasswordDoubleHashedValidator is a validator for the "password_double_hashed" field. It is called by the builders before save.
	user.PasswordDoubleHashedValidator = userDescPasswordDoubleHashed.Validators[0].(func([]byte) error)
	// userDescCertificate is the schema descriptor for certificate field.
	userDescCertificate := userFields[4].Descriptor()
	// user.CertificateValidator is a validator for the "certificate" field. It is called by the builders before save.
	user.CertificateValidator = userDescCertificate.Validators[0].(func(string) error)
	// userDescPublicKey is the schema descriptor for public_key field.
	userDescPublicKey := userFields[5].Descriptor()
	// user.PublicKeyValidator is a validator for the "public_key" field. It is called by the builders before save.
	user.PublicKeyValidator = userDescPublicKey.Validators[0].(func(string) error)
	// userDescEncryptedPrivateKey is the schema descriptor for encrypted_private_key field.
	userDescEncryptedPrivateKey := userFields[6].Descriptor()
	// user.EncryptedPrivateKeyValidator is a validator for the "encrypted_private_key" field. It is called by the builders before save.
	user.EncryptedPrivateKeyValidator = userDescEncryptedPrivateKey.Validators[0].(func([]byte) error)
}
