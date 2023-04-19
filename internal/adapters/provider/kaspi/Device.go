package kaspi

import (
	bytes2 "bytes"
	"encoding/json"
	"io"
	"kaspi-qr/internal/adapters/provider"
	"kaspi-qr/internal/domain/errs"
	"log"
	"net/http"
)

func (s *St) GetAllTradePoints(orgBin string) ([]*provider.TradePointSt, error) {
	var bodyRequest provider.TradePointListRepSt

	client, err := GetHttpClientTls(s.certPath, s.certPassword)

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	req, err := http.NewRequest("GET", s.kaspiUrl+"partner/tradepoints/"+orgBin, nil)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

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

	return bodyRequest.Data, nil
}

func (s *St) DeviceRegistration(input provider.DeviceCreateReqSt) (string, error) {
	var bodyRequest provider.DeviceCreateRepSt

	client, err := GetHttpClientTls(s.certPath, s.certPassword)

	if err != nil {
		return "", err
	}

	requestBody, err := json.Marshal(input)

	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", s.kaspiUrl+"device/register/", bytes2.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	errJson := json.Unmarshal(bytes, &bodyRequest)
	if errJson != nil {
		return "", err
	}

	return bodyRequest.Data.DeviceToken, nil
}

func (s *St) DeviceDelete(input provider.DeviceRemoveReqSt) error {
	var bodyRequest provider.BaseRepSt

	client, err := GetHttpClientTls(s.certPath, s.certPassword)

	if err != nil {
		return err
	}

	requestBody, err := json.Marshal(input)

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", s.kaspiUrl+"device/delete", bytes2.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &bodyRequest)
	if err != nil {
		return err
	}

	if bodyRequest.StatusCode != SuccessStatus {
		return errs.ServiceNA
	}

	return nil
}
