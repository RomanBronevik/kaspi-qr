package kaspi

import (
	bytes2 "bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"kaspi-qr/configs"
	"kaspi-qr/internal/domain/entities"
	"log"
	"net/http"
)

func (s *St) KaspiOperationDetails(requestBody io.ReadCloser) (entities.OperationDetails, error) {
	var bodyRequest entities.OperationDetails

	var inputBody entities.OperationGetSt

	body := requestBody
	x, err := io.ReadAll(body)

	if err != nil {
		log.Fatal(err.Error())
		return entities.OperationDetails{}, err
	}

	errJsonUnmarshall := json.Unmarshal(x, &inputBody)
	if errJsonUnmarshall != nil {
		log.Fatal(err.Error())
		return entities.OperationDetails{}, err
	}

	client, err := configs.GetHttpClientTls()

	if err != nil {
		log.Fatal(err.Error())
	}

	req, err := http.NewRequest("GET", viper.GetString("kaspiURL")+"payment/details?QrPaymentId="+fmt.Sprint(inputBody.QrPaymentId)+"&DeviceToken="+inputBody.DeviceToken, nil)
	if err != nil {
		log.Fatal(err.Error())
		return entities.OperationDetails{}, err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
		return entities.OperationDetails{}, err
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
		return entities.OperationDetails{}, err
	}

	errJson := json.Unmarshal(bytes, &bodyRequest)
	if errJson != nil {
		log.Fatal(err.Error())
		return entities.OperationDetails{}, err
	}

	return bodyRequest, nil
}

func (s *St) KaspiReturnWithoutClient(input entities.ReturnRequestInput) (entities.ReturnSt, error) {
	var bodyRequest entities.ReturnSt

	client, err := configs.GetHttpClientTls()

	if err != nil {
		log.Fatal(err.Error())
	}

	requestBody, err := json.Marshal(input)

	if err != nil {
		return entities.ReturnSt{}, err
	}

	req, err := http.NewRequest("POST", viper.GetString("kaspiURL")+"payment/return", bytes2.NewBuffer(requestBody))
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
