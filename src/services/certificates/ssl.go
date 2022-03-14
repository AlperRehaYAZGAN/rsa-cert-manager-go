// This shows how to generate a SSH RSA Private/Public key pair and save it locally
package certificates

type SSLCertificate struct {
	PrivateKey []byte
	PublicKey  []byte
}

func GenerateSSLKeyPair() (*SSLCertificate, error) {
	bitSize := 2048

	privateKey, err := GeneratePrivateKey(bitSize)
	if err != nil {
		return nil, err
	}

	publicKeyBytes, err := GenerateSSLPublicKeyPem(&privateKey.PublicKey)
	if err != nil {
		return nil, err
	}

	privateKeyBytes := EncodePrivateKeyToPEM(privateKey)

	return &SSLCertificate{
		PrivateKey: privateKeyBytes,
		PublicKey:  publicKeyBytes,
	}, nil

}
