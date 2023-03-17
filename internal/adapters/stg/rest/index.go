package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"kaspi-qr/internal/adapters/logger"
)

const requestTimout = 10 * time.Second

type St struct {
	lg     logger.Lite
	apiUrl string

	httpClient *http.Client
}

func New(lg logger.Lite, apiUrl string) *St {
	return &St{
		lg:     lg,
		apiUrl: strings.TrimRight(apiUrl, "/") + "/",

		httpClient: &http.Client{
			Timeout: requestTimout,
		},
	}
}

func (o *St) sendRequest(method, path string, reqObj any, reqPars url.Values, repObj any) ([]byte, error) {
	var reqStream io.Reader = nil

	if reqObj != nil {
		raw, err := json.Marshal(reqObj)
		if err != nil {
			return nil, err
		}

		reqStream = bytes.NewReader(raw)
	}

	req, err := http.NewRequest(method, o.apiUrl+path, reqStream)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest: %w", err)
	}

	if reqObj != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Accept", "application/json")

	if reqPars != nil {
		req.URL.RawQuery = reqPars.Encode()
	}

	o.lg.Infow("Send request", "method", method, "path", path, "reqPars", reqPars)

	resp, err := o.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("httpClient.Do: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("resp.StatusCode: %d", resp.StatusCode)
	}

	repData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll body: %w", err)
	}

	if len(repData) > 0 && repObj != nil {
		err = json.Unmarshal(repData, repObj)
		if err != nil {
			return nil, fmt.Errorf("json.Unmarshal: %w", err)
		}
	}

	return repData, nil
}
