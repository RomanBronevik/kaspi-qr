package http

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"kaspi-qr/internal/adapters/logger"
	"kaspi-qr/internal/adapters/notifier"
	"net/http"
	"time"
)

type St struct {
	lg logger.Lite
}

func New(lg logger.Lite) *St {
	return &St{
		lg: lg,
	}
}

func (o *St) NotifyOrderStatusChange(uri string, obj *notifier.OrderStatusChangeReqSt) error {
	o.lg.Infow("NotifyOrderStatusChange", "uri", uri, "obj", obj)

	if uri == "" {
		return nil
	}

	httpClient := http.Client{
		Timeout:   5 * time.Second,
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}

	reqBody, err := json.Marshal(obj)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(reqBody))
	if err != nil {
		return fmt.Errorf("http.NewRequest: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("client.Do: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("resp.StatusCode: %d", resp.StatusCode)
	}

	return nil
}
