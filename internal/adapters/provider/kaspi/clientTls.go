package kaspi

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"net/http"
	"os"

	"golang.org/x/crypto/pkcs12"
)

func GetHttpClientTls(CertPath string, CertPassword string) (*http.Client, error) {
	pfxFile := CertPath
	pfxData, err := os.ReadFile(pfxFile)

	if err != nil {
		return nil, err
	}

	blocks, err := pkcs12.ToPEM(pfxData, CertPassword)
	if err != nil {
		return nil, err
	}
	var pemData []byte
	for _, b := range blocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}
	cert, err := tls.X509KeyPair(pemData, pemData)

	config := &tls.Config{Certificates: []tls.Certificate{cert}}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(pemData)

	config = &tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: config,
		},
	}

	return client, nil
}
