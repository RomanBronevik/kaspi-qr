package kaspi

import (
	bytes2 "bytes"
	"encoding/json"
	"github.com/spf13/viper"
	"io"
	"kaspi-qr/configs"
	"kaspi-qr/internal/domain/entities"
	"log"
	"net/http"
)

func CreateQrToken(input entities.QrTokenInput) (entities.QrTokenOutput, error) {
	var bodyRequest entities.QrTokenOutput

	client, err := configs.GetHttpClientTls()

	if err != nil {
		log.Fatal(err.Error())
	}

	requestBody, err := json.Marshal(input)

	if err != nil {
		log.Fatal(err.Error())
		return entities.QrTokenOutput{}, err
	}

	req, err := http.NewRequest("POST", viper.GetString("kaspiURL")+"qr/create", bytes2.NewBuffer(requestBody))
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

func CreatePaymentLink(requestBody io.ReadCloser) (entities.PaymentLink, error) {
	var bodyRequest entities.PaymentLink

	client, err := configs.GetHttpClientTls()

	if err != nil {
		log.Fatal(err.Error())
	}

	req, err := http.NewRequest("POST", viper.GetString("kaspiURL")+"qr/create-link", requestBody)
	if err != nil {
		log.Fatal(err.Error())
		return entities.PaymentLink{}, err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
		return entities.PaymentLink{}, err
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
		return entities.PaymentLink{}, err
	}

	errJson := json.Unmarshal(bytes, &bodyRequest)
	if errJson != nil {
		log.Fatal(err.Error())
		return entities.PaymentLink{}, err
	}

	return bodyRequest, nil
}

func OperationStatus(QrPaymentId string) (entities.OperationStatus, error) {

	var bodyRequest entities.OperationStatus

	client, err := configs.GetHttpClientTls()

	if err != nil {
		log.Fatal(err.Error())
		return entities.OperationStatus{}, err
	}

	req, err := http.NewRequest("GET", viper.GetString("kaspiURL")+"payment/status/"+QrPaymentId, nil)
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
