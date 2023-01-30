package handler

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
)

func kaspiTradePoints(organizationBIN string) (tradePointSt, error) {

	var bodyRequest tradePointSt

	client, err := getHttpClientTls()

	if err != nil {
		log.Fatal(err.Error())
		return tradePointSt{}, err
	}

	req, err := http.NewRequest("GET", viper.GetString("kaspiURL")+"partner/tradepoints/"+organizationBIN, nil)
	if err != nil {
		log.Fatal(err.Error())
		return tradePointSt{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
		return tradePointSt{}, err
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
		return tradePointSt{}, err
	}

	errJson := json.Unmarshal(bytes, &bodyRequest)
	if errJson != nil {
		log.Fatal(err.Error())
		return tradePointSt{}, err
	}

	return bodyRequest, nil
}

func kaspiDeviceRegistration(requestBody io.ReadCloser) (RegistrationOutputSt, error) {
	var bodyRequest RegistrationOutputSt

	client, err := getHttpClientTls()

	if err != nil {
		log.Fatal(err.Error())
	}

	req, err := http.NewRequest("POST", viper.GetString("kaspiURL")+"device/register/", requestBody)
	if err != nil {
		log.Fatal(err.Error())
		return RegistrationOutputSt{}, err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
		return RegistrationOutputSt{}, err
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
		return RegistrationOutputSt{}, err
	}

	errJson := json.Unmarshal(bytes, &bodyRequest)
	if errJson != nil {
		fmt.Println(errJson.Error())
		log.Fatal(err.Error())
		return RegistrationOutputSt{}, err
	}

	return bodyRequest, nil
}

func kaspiDeviceDelete(requestBody io.ReadCloser) (DeleteOutputSt, error) {
	var bodyRequest DeleteOutputSt

	client, err := getHttpClientTls()

	if err != nil {
		log.Fatal(err.Error())
	}

	req, err := http.NewRequest("POST", viper.GetString("kaspiURL")+"device/delete", requestBody)
	if err != nil {
		log.Fatal(err.Error())
		return DeleteOutputSt{}, err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
		return DeleteOutputSt{}, err
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
		return DeleteOutputSt{}, err
	}

	errJson := json.Unmarshal(bytes, &bodyRequest)
	if errJson != nil {
		fmt.Println(errJson.Error())
		log.Fatal(err.Error())
		return DeleteOutputSt{}, err
	}

	return bodyRequest, nil
}
