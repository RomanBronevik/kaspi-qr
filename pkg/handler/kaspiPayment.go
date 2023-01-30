package handler

import (
	"encoding/json"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
)

func kaspiQR(requestBody io.ReadCloser) (QRToken, error) {
	var bodyRequest QRToken

	client, err := getHttpClientTls()

	if err != nil {
		log.Fatal(err.Error())
	}

	req, err := http.NewRequest("POST", viper.GetString("kaspiURL")+"qr/create", requestBody)
	if err != nil {
		log.Fatal(err.Error())
		return QRToken{}, err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
		return QRToken{}, err
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
		return QRToken{}, err
	}

	errJson := json.Unmarshal(bytes, &bodyRequest)
	if errJson != nil {
		log.Fatal(err.Error())
		return QRToken{}, err
	}

	return bodyRequest, nil
}

func kaspiPaymentLink(requestBody io.ReadCloser) (PaymentLink, error) {
	var bodyRequest PaymentLink

	client, err := getHttpClientTls()

	if err != nil {
		log.Fatal(err.Error())
	}

	req, err := http.NewRequest("POST", viper.GetString("kaspiURL")+"qr/create-link", requestBody)
	if err != nil {
		log.Fatal(err.Error())
		return PaymentLink{}, err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
		return PaymentLink{}, err
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
		return PaymentLink{}, err
	}

	errJson := json.Unmarshal(bytes, &bodyRequest)
	if errJson != nil {
		log.Fatal(err.Error())
		return PaymentLink{}, err
	}

	return bodyRequest, nil
}
