package handler

import (
	"encoding/json"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
)

func kaspiOperationDetails(requestBody io.ReadCloser) (OperationDetails, error) {
	var bodyRequest OperationDetails

	var inputBody OperationGetSt

	body := requestBody
	x, err := io.ReadAll(body)

	if err != nil {
		log.Fatal(err.Error())
		return OperationDetails{}, err
	}

	errJsonUnmarshall := json.Unmarshal(x, &inputBody)
	if errJsonUnmarshall != nil {
		log.Fatal(err.Error())
		return OperationDetails{}, err
	}

	client, err := getHttpClientTls()

	if err != nil {
		log.Fatal(err.Error())
	}

	req, err := http.NewRequest("GET", viper.GetString("kaspiURL")+"payment/details?QrPaymentId="+string(inputBody.QrPaymentId)+"&DeviceToken="+inputBody.DeviceToken, nil)
	if err != nil {
		log.Fatal(err.Error())
		return OperationDetails{}, err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
		return OperationDetails{}, err
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
		return OperationDetails{}, err
	}

	errJson := json.Unmarshal(bytes, &bodyRequest)
	if errJson != nil {
		log.Fatal(err.Error())
		return OperationDetails{}, err
	}

	return bodyRequest, nil
}

func kaspiReturnWithoutClient(requestBody io.ReadCloser) (ReturnSt, error) {
	var bodyRequest ReturnSt

	client, err := getHttpClientTls()

	if err != nil {
		log.Fatal(err.Error())
	}

	req, err := http.NewRequest("POST", viper.GetString("kaspiURL")+"payment/return", requestBody)
	if err != nil {
		log.Fatal(err.Error())
		return ReturnSt{}, err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
		return ReturnSt{}, err
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
		return ReturnSt{}, err
	}

	errJson := json.Unmarshal(bytes, &bodyRequest)
	if errJson != nil {
		log.Fatal(err.Error())
		return ReturnSt{}, err
	}

	return bodyRequest, nil
}
