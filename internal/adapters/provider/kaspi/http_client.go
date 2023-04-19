package kaspi

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"kaspi-qr/internal/adapters/logger"
	"net/http"
	"os"
	"strings"

	"golang.org/x/crypto/pkcs12"
)

type httpClientSt struct {
	lg     logger.Lite
	uri    string
	client *http.Client
}

func newHttpClient(lg logger.Lite, uri string, certPath, certKey string) (*httpClientSt, error) {
	pfxData, err := os.ReadFile(certPath)
	if err != nil {
		return nil, fmt.Errorf("os.ReadFile: %w", err)
	}

	blocks, err := pkcs12.ToPEM(pfxData, certKey)
	if err != nil {
		return nil, fmt.Errorf("pkcs12.ToPEM: %w", err)
	}

	var pemData []byte
	for _, b := range blocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}

	cert, err := tls.X509KeyPair(pemData, pemData)
	if err != nil {
		return nil, fmt.Errorf("tls.X509KeyPair: %w", err)
	}

	//caCertPool := x509.NewCertPool()
	//caCertPool.AppendCertsFromPEM(pemData)

	return &httpClientSt{
		lg:  lg,
		uri: strings.TrimRight(uri, "/") + "/",
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					Certificates:       []tls.Certificate{cert},
					InsecureSkipVerify: true,
				},
			},
		},
	}, nil
}

func (c *httpClientSt) sendRequest(method, path string, reqObj, repObj any) (*httpRespSt, error) {
	result := &httpRespSt{
		lg:        c.lg,
		reqMethod: method,
		reqPath:   path,
	}

	var reqStream io.Reader

	if reqObj != nil {
		requestBody, err := json.Marshal(reqObj)
		if err != nil {
			return result, fmt.Errorf("json.Marshal: %w", err)
		}
		result.reqBody = string(requestBody)
		reqStream = bytes.NewBuffer(requestBody)
	}

	req, err := http.NewRequest(method, c.uri+path, reqStream)
	if err != nil {
		return result, fmt.Errorf("http.NewRequest: %w", err)
	}

	if reqObj != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return result, fmt.Errorf("client.Do: %w", err)
	}
	defer resp.Body.Close()

	result.repStatusCode = resp.StatusCode

	repBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, fmt.Errorf("io.ReadAll: %w", err)
	}
	result.repBody = string(repBody)

	if repObj != nil {
		err = json.Unmarshal(repBody, &repObj)
		if err != nil {
			return result, fmt.Errorf("json.Unmarshal: %w", err)
		}
	}

	return result, nil
}

// httpRespSt

type httpRespSt struct {
	lg logger.Lite

	// req
	reqMethod string
	reqPath   string
	reqBody   string

	// rep
	repStatusCode int
	repBody       string
}

func (r *httpRespSt) LogError(msg string, err error) {
	r.lg.Errorw(
		msg, err,
		"reqMethod", r.reqMethod,
		"reqPath", r.reqPath,
		"reqBody", r.reqBody,
		"repStatusCode", r.repStatusCode,
		"repBody", r.repBody,
	)
}
