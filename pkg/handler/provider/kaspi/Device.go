package provider

import (
	bytes2 "bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"kaspi-qr/configs"
	"kaspi-qr/pkg/handler/entities"
	"log"
	"net/http"
)

func KaspiTradePoints(organizationBIN string) (entities.TradePointSt, error) {

	var bodyRequest entities.TradePointSt

	client, err := configs.GetHttpClientTls()

	if err != nil {
		log.Fatal(err.Error())
		return entities.TradePointSt{}, err
	}

	req, err := http.NewRequest("GET", viper.GetString("kaspiURL")+"partner/tradepoints/"+organizationBIN, nil)
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

func KaspiDeviceRegistration(input entities.DeviceInputReg) (entities.DeviceOutputReg, error) {
	var bodyRequest entities.DeviceOutputReg

	client, err := configs.GetHttpClientTls()

	if err != nil {
		log.Fatal(err.Error())
	}

	requestBody, err := json.Marshal(input)

	if err != nil {
		log.Fatal(err.Error())
		return entities.DeviceOutputReg{}, err
	}

	req, err := http.NewRequest("POST", viper.GetString("kaspiURL")+"device/register/", bytes2.NewBuffer(requestBody))
	if err != nil {
		log.Fatal(err.Error())
		return entities.DeviceOutputReg{}, err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
		return entities.DeviceOutputReg{}, err
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
		return entities.DeviceOutputReg{}, err
	}

	errJson := json.Unmarshal(bytes, &bodyRequest)
	if errJson != nil {
		fmt.Println(errJson.Error())
		log.Fatal(err.Error())
		return entities.DeviceOutputReg{}, err
	}

	return bodyRequest, nil
}

func KaspiDeviceDelete(requestBody io.ReadCloser) (entities.DeviceOutputDel, error) {
	var bodyRequest entities.DeviceOutputDel

	client, err := configs.GetHttpClientTls()

	if err != nil {
		log.Fatal(err.Error())
	}

	req, err := http.NewRequest("POST", viper.GetString("kaspiURL")+"device/delete", requestBody)
	if err != nil {
		log.Fatal(err.Error())
		return entities.DeviceOutputDel{}, err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
		return entities.DeviceOutputDel{}, err
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
		return entities.DeviceOutputDel{}, err
	}

	errJson := json.Unmarshal(bytes, &bodyRequest)
	if errJson != nil {
		fmt.Println(errJson.Error())
		log.Fatal(err.Error())
		return entities.DeviceOutputDel{}, err
	}

	return bodyRequest, nil
}
