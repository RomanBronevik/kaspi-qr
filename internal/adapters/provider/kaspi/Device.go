package kaspi

import (
	bytes2 "bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"kaspi-qr/internal/domain/entities"
)

func (s *St) GetAllTradePoints(organizationBIN string) (entities.TradePointSt, error) {

	var bodyRequest entities.TradePointSt

	client, err := GetHttpClientTls(s.certPath, s.certPassword)

	if err != nil {
		log.Fatal(err.Error())
		return entities.TradePointSt{}, err
	}

	req, err := http.NewRequest("GET", s.kaspiUrl+"partner/tradepoints/"+organizationBIN, nil)
	if err != nil {
		log.Fatal(err.Error())
		return entities.TradePointSt{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
		return entities.TradePointSt{}, err
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
		return entities.TradePointSt{}, err
	}

	errJson := json.Unmarshal(bytes, &bodyRequest)
	if errJson != nil {
		log.Fatal(err.Error())
		return entities.TradePointSt{}, err
	}

	return bodyRequest, nil
}

func (s *St) DeviceRegistration(input entities.DeviceInputReg) (entities.DeviceOutputReg, error) {
	var bodyRequest entities.DeviceOutputReg

	client, err := GetHttpClientTls(s.certPath, s.certPassword)

	if err != nil {
		return entities.DeviceOutputReg{}, err
	}

	requestBody, err := json.Marshal(input)

	if err != nil {
		return entities.DeviceOutputReg{}, err
	}

	req, err := http.NewRequest("POST", s.kaspiUrl+"device/register/", bytes2.NewBuffer(requestBody))
	if err != nil {
		return entities.DeviceOutputReg{}, err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return entities.DeviceOutputReg{}, err
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return entities.DeviceOutputReg{}, err
	}

	errJson := json.Unmarshal(bytes, &bodyRequest)
	if errJson != nil {
		return entities.DeviceOutputReg{}, err
	}

	return bodyRequest, nil
}

func (s *St) DeviceDelete(input entities.DeviceInputDel) (entities.DeviceOutputDel, error) {
	var bodyRequest entities.DeviceOutputDel

	client, err := GetHttpClientTls(s.certPath, s.certPassword)

	if err != nil {
		return entities.DeviceOutputDel{}, err
	}

	requestBody, err := json.Marshal(input)

	if err != nil {
		return entities.DeviceOutputDel{}, err
	}

	req, err := http.NewRequest("POST", s.kaspiUrl+"device/delete", bytes2.NewBuffer(requestBody))
	if err != nil {
		return entities.DeviceOutputDel{}, err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return entities.DeviceOutputDel{}, err
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return entities.DeviceOutputDel{}, err
	}

	err = json.Unmarshal(bytes, &bodyRequest)
	if err != nil {
		return entities.DeviceOutputDel{}, err
	}

	return bodyRequest, nil
}
