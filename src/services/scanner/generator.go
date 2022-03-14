package scanner

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
)

/**
*	@param address string - The address to scan like "www.google.com:443"
*	@return string,error - The certificate in PEM format or an error
 */
func GetCertificatesPEM(address string) (string, error) {
	conn, err := tls.Dial("tcp", address, &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		return "", err
	}
	defer conn.Close()
	var b bytes.Buffer
	for _, cert := range conn.ConnectionState().PeerCertificates {
		err := pem.Encode(&b, &pem.Block{
			Type:  "CERTIFICATE",
			Bytes: cert.Raw,
		})
		if err != nil {
			return "", err
		}
	}
	return b.String(), nil
}

/**
*	@param address string - The address to scan like "www.google.com:443"
*	@return x509.Certificate,error - The certificate as a x509.Cert or an error
 */
func GetCertificates(address string) ([]*x509.Certificate, error) {
	conn, err := tls.Dial("tcp", address, &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	// get the certificate from the connection state
	certs := conn.ConnectionState().PeerCertificates
	return certs, nil
}
