// This shows how to generate a SSH RSA Private/Public key pair and save it locally
package certificates

type SSHCertificate struct {
	PrivateKey []byte
	PublicKey  []byte
}

func GenerateSSHKeyPair() (*SSHCertificate, error) {
	bitSize := 4096

	privateKey, err := GeneratePrivateKey(bitSize)
	if err != nil {
		return nil, err
	}

	publicKeyBytes, err := GenerateSSHPublicKeyPem(&privateKey.PublicKey)
	if err != nil {
		return nil, err
	}

	privateKeyBytes := EncodePrivateKeyToPEM(privateKey)

	return &SSHCertificate{
		PrivateKey: privateKeyBytes,
		PublicKey:  publicKeyBytes,
	}, nil

}
