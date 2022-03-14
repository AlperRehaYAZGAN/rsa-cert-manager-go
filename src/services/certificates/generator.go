package certificates

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"

	"golang.org/x/crypto/ssh"
)

// GeneratePrivateKey creates a RSA Private Key of specified byte size
func GeneratePrivateKey(bitSize int) (*rsa.PrivateKey, error) {
	// Private Key generation
	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return nil, err
	}

	// Validate Private Key
	err = privateKey.Validate()
	if err != nil {
		return nil, err
	}

	log.Println("Private Key generated")
	return privateKey, nil
}

// GeneratePublicKey take a rsa.PublicKey and return bytes suitable for writing to .pub file
// returns in the format "ssh-rsa ..."
func GenerateSSHPublicKeyPem(privatekey *rsa.PublicKey) ([]byte, error) {
	publicRsaKey, err := ssh.NewPublicKey(privatekey)
	if err != nil {
		return nil, err
	}

	pubKeyBytes := ssh.MarshalAuthorizedKey(publicRsaKey)

	return pubKeyBytes, nil
}

func GenerateSSLPublicKeyPem(privatekey *rsa.PublicKey) ([]byte, error) {
	publicRsaKey, err := x509.MarshalPKIXPublicKey(privatekey)
	if err != nil {
		return nil, err
	}

	pubPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicRsaKey,
	})

	return pubPEM, nil
}

// EncodePrivateKeyToPEM encodes Private Key from RSA to PEM format
func EncodePrivateKeyToPEM(privateKey *rsa.PrivateKey) []byte {
	// Get ASN.1 DER format
	privDER := x509.MarshalPKCS1PrivateKey(privateKey)

	// pem.Block
	privBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privDER,
	}

	// Private key in PEM format
	privatePEM := pem.EncodeToMemory(&privBlock)

	return privatePEM
}
