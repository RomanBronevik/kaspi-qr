package kaspi

import (
	bytes2 "bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"kaspi-qr/internal/domain/entities"
)

func (s *St) CreateQrToken(input entities.KaspiPaymentInput) (entities.QrTokenOutput, error) {
	var bodyRequest entities.QrTokenOutput

	client, err := GetHttpClientTls(s.certPath, s.certPassword)

	if err != nil {
		log.Fatal(err.Error())
	}

	requestBody, err := json.Marshal(input)

	if err != nil {
		log.Fatal(err.Error())
		return entities.QrTokenOutput{}, err
	}

	req, err := http.NewRequest("POST", s.kaspiUrl+"qr/create", bytes2.NewBuffer(requestBody))
	if err != nil {
		log.Fatal(err.Error())
		return entities.QrTokenOutput{}, err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
		return entities.QrTokenOutput{}, err
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
		return entities.QrTokenOutput{}, err
	}
	errJson := json.Unmarshal(bytes, &bodyRequest)
	if errJson != nil {
		log.Fatal(err.Error())
		return entities.QrTokenOutput{}, err
	}

	return bodyRequest, nil
}

func (s *St) CreatePaymentLink(input entities.KaspiPaymentInput) (entities.PaymentLinkRequestOutput, error) {
	var bodyRequest entities.PaymentLinkRequestOutput

	client, err := GetHttpClientTls(s.certPath, s.certPassword)

	if err != nil {
		log.Fatal(err.Error())
	}

	requestBody, err := json.Marshal(input)

	if err != nil {
		log.Fatal(err.Error())
		return entities.PaymentLinkRequestOutput{}, err
	}

	req, err := http.NewRequest("POST", s.kaspiUrl+"qr/create-link", bytes2.NewBuffer(requestBody))
	if err != nil {
		log.Fatal(err.Error())
		return entities.PaymentLinkRequestOutput{}, err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return entities.PaymentLinkRequestOutput{}, err
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return entities.PaymentLinkRequestOutput{}, err
	}

	errJson := json.Unmarshal(bytes, &bodyRequest)
	if errJson != nil {
		return entities.PaymentLinkRequestOutput{}, err
	}

	return bodyRequest, nil
}

func (s *St) OperationStatus(QrPaymentId string) (entities.OperationStatus, error) {

	var bodyRequest entities.OperationStatus

	client, err := GetHttpClientTls(s.certPath, s.certPassword)

	if err != nil {
		log.Fatal(err.Error())
		return entities.OperationStatus{}, err
	}

	req, err := http.NewRequest("GET", s.kaspiUrl+"payment/status/"+QrPaymentId, nil)
	if err != nil {
		log.Fatal(err.Error())
		return entities.OperationStatus{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
		return entities.OperationStatus{}, err
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
		return entities.OperationStatus{}, err
	}

	errJson := json.Unmarshal(bytes, &bodyRequest)
	if errJson != nil {
		log.Fatal(err.Error())
		return entities.OperationStatus{}, err
	}

	return bodyRequest, nil
}
