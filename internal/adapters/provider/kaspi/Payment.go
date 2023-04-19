package kaspi

import (
	bytes2 "bytes"
	"encoding/json"
	"io"
	"kaspi-qr/internal/adapters/provider"
	"log"
	"net/http"
)

func (s *St) CreateQrToken(input provider.PaymentCreateReqSt) (*provider.PaymentSt, error) {
	var bodyRequest *provider.PaymentCreateRepSt

	client, err := GetHttpClientTls(s.certPath, s.certPassword)

	if err != nil {
		log.Fatal(err.Error())
	}

	requestBody, err := json.Marshal(input)

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	req, err := http.NewRequest("POST", s.kaspiUrl+"qr/create", bytes2.NewBuffer(requestBody))
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	errJson := json.Unmarshal(bytes, &bodyRequest)
	if errJson != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return &bodyRequest.Data, nil
}

func (s *St) CreatePaymentLink(input provider.PaymentLinkCreateReqSt) (*provider.PaymentLinkSt, error) {
	var bodyRequest provider.PaymentLinkCreateRepSt

	client, err := GetHttpClientTls(s.certPath, s.certPassword)

	if err != nil {
		log.Fatal(err.Error())
	}

	requestBody, err := json.Marshal(input)

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	req, err := http.NewRequest("POST", s.kaspiUrl+"qr/create-link", bytes2.NewBuffer(requestBody))
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	errJson := json.Unmarshal(bytes, &bodyRequest)
	if errJson != nil {
		return nil, err
	}

	return bodyRequest.Data, nil
}

func (s *St) OperationStatus(QrPaymentId string) (string, error) {

	var bodyRequest provider.OperationStatusRepSt

	client, err := GetHttpClientTls(s.certPath, s.certPassword)

	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}

	req, err := http.NewRequest("GET", s.kaspiUrl+"payment/status/"+QrPaymentId, nil)
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}

	errJson := json.Unmarshal(bytes, &bodyRequest)
	if errJson != nil {
		log.Fatal(err.Error())
		return "", err
	}

	return bodyRequest.Data.Status, nil
}
