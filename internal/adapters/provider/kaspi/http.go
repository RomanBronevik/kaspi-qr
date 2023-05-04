package kaspi

import (
	"crypto/tls"
	"encoding/pem"
	"fmt"
	"net/http"
	"os"

	"github.com/rendau/dop/dopErrs"

	"github.com/rendau/dop/adapters/client/httpc"
	"github.com/rendau/dop/adapters/client/httpc/httpclient"
	"golang.org/x/crypto/pkcs12"
)

func createCert(certPath, certKey string) (tls.Certificate, error) {
	pfxData, err := os.ReadFile(certPath)
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("os.ReadFile: %w", err)
	}

	blocks, err := pkcs12.ToPEM(pfxData, certKey)
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("pkcs12.ToPEM: %w", err)
	}

	var pemData []byte
	for _, b := range blocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}

	cert, err := tls.X509KeyPair(pemData, pemData)
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("tls.X509KeyPair: %w", err)
	}

	//caCertPool := x509.NewCertPool()
	//caCertPool.AppendCertsFromPEM(pemData)

	return cert, nil
}

func (s *St) sendRequest(method, path string, reqObj, repObj any) (*httpc.RespSt, error) {
	httpClient := httpclient.New(s.lg, &httpc.OptionsSt{
		Client: &http.Client{
			Timeout: RequestTimeout,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					Certificates:       []tls.Certificate{s.cert},
					InsecureSkipVerify: true,
				}},
		},
		LogPrefix: "Kaspi: ",
	})

	resp, err := httpClient.Send(&httpc.OptionsSt{
		Uri:    s.uri + path,
		Method: method,
		//LogFlags: httpc.LogResponse,

		ReqObj: reqObj,
		RepObj: repObj,
	})
	if err != nil {
		return nil, dopErrs.ServiceNA
	}

	return resp, nil
}
