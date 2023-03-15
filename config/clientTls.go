package config

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"net/http"
	"os"

	"golang.org/x/crypto/pkcs12"
)

func GetHttpClientTls() (*http.Client, error) {
	pfxFile := os.Getenv("CERTIFICATE_PATH")
	pfxData, err := os.ReadFile(pfxFile)

	if err != nil {
		return nil, err
	}

	blocks, err := pkcs12.ToPEM(pfxData, os.Getenv("CERTIFICATE_PASSWORD"))
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
