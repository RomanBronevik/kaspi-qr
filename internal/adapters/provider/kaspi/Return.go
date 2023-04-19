package kaspi

import (
	bytes2 "bytes"
	"encoding/json"
	"fmt"
	"io"
	"kaspi-qr/internal/adapters/provider"
	"log"
	"net/http"
)

func (s *St) KaspiOperationDetails(input provider.OperationGetReqSt) (*provider.OperationDetailsSt, error) {
	var bodyRequest provider.OperationGetRepSt

	client, err := GetHttpClientTls(s.certPath, s.certPassword)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", s.kaspiUrl+"payment/details?QrPaymentId="+fmt.Sprint(input.QrPaymentId)+"&DeviceToken="+input.DeviceToken, nil)
	if err != nil {
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

	return &bodyRequest.Data, nil
}

func (s *St) KaspiReturnWithoutClient(input provider.ReturnReqSt) (int64, error) {
	var bodyRequest provider.ReturnRepSt

	client, err := GetHttpClientTls(s.certPath, s.certPassword)

	if err != nil {
		log.Fatal(err.Error())
	}

	requestBody, err := json.Marshal(input)

	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest("POST", s.kaspiUrl+"payment/return", bytes2.NewBuffer(requestBody))
	if err != nil {
		return 0, err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	errJson := json.Unmarshal(bytes, &bodyRequest)
	if errJson != nil {
		return 0, err
	}

	return bodyRequest.ReturnOperationDataSt.ReturnOperationId, nil
}
