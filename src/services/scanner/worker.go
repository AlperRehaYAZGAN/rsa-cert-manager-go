package scanner

import (
	"crypto/tls"
	"crypto/x509"
)

func ScanUrl(url string) ([]*x509.Certificate, error) {
	conn, err := tls.Dial("tcp", url, &tls.Config{
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
