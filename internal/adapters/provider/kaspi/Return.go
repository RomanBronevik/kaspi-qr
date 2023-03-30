package kaspi

import (
	bytes2 "bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"kaspi-qr/internal/domain/entities"
)

func (s *St) KaspiOperationDetails(input entities.OperationGetSt) (entities.OperationDetails, error) {
	var bodyRequest entities.OperationDetails

	client, err := GetHttpClientTls(s.certPath, s.certPassword)

	if err != nil {
		return entities.OperationDetails{}, err
	}

	req, err := http.NewRequest("GET", s.kaspiUrl+"payment/details?QrPaymentId="+fmt.Sprint(input.QrPaymentId)+"&DeviceToken="+input.DeviceToken, nil)
	if err != nil {
		return entities.OperationDetails{}, err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return entities.OperationDetails{}, err
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return entities.OperationDetails{}, err
	}

	errJson := json.Unmarshal(bytes, &bodyRequest)
	if errJson != nil {
		return entities.OperationDetails{}, err
	}

	return bodyRequest, nil
}

func (s *St) KaspiReturnWithoutClient(input entities.ReturnRequestInput) (entities.ReturnSt, error) {
	var bodyRequest entities.ReturnSt

	client, err := GetHttpClientTls(s.certPath, s.certPassword)

	if err != nil {
		log.Fatal(err.Error())
	}

	requestBody, err := json.Marshal(input)

	if err != nil {
		return entities.ReturnSt{}, err
	}

	req, err := http.NewRequest("POST", s.kaspiUrl+"payment/return", bytes2.NewBuffer(requestBody))
	if err != nil {
		return entities.ReturnSt{}, err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return entities.ReturnSt{}, err
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return entities.ReturnSt{}, err
	}

	errJson := json.Unmarshal(bytes, &bodyRequest)
	if errJson != nil {
		return entities.ReturnSt{}, err
	}

	return bodyRequest, nil
}
