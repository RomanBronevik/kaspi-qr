package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io"
	"log"
)

func (h *Handler) tradePoints(c *gin.Context) {

	//AuthorizationToken := c.Request.Header["AuthorizationToken"]

	organizationBIN := c.Param("organizationBIN")

	req, err := kaspiTradePoints(organizationBIN)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Something went wrong",
			"error":   err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"output": req,
	})
}

func (h *Handler) deviceRegistration(c *gin.Context) {
	//var bodyRequest tradePointRegistration

	body := c.Request.Body
	kaspiDeviceRegistration(body)

	//x, err := ioutil.ReadAll(body)
	//
	//if err != nil {
	//	c.JSON(400, gin.H{
	//		"message": "problems with JSON body - 1",
	//	})
	//}
	//
	//errJson := json.Unmarshal(x, &bodyRequest)
	//
	//if errJson != nil {
	//	c.JSON(400, gin.H{
	//		"message": "problems with JSON body - 2",
	//	})
	//}
}

func (h *Handler) deleteOrOffDevice(c *gin.Context) {

}

func kaspiDeviceRegistration(requestBody io.ReadCloser) (RegistrationOutputSt, error) {
	var bodyRequest RegistrationOutputSt

	client, err := getHttpClientTls()

	if err != nil {
		log.Fatal(err.Error())
	}

	resp, err := client.Get(viper.GetString("kaspiURL") + "device/register/")
	if err != nil {
		panic(err)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
		return RegistrationOutputSt{}, err
	}

	errJson := json.Unmarshal(bytes, &bodyRequest)
	if errJson != nil {
		log.Fatal(err.Error())
		return RegistrationOutputSt{}, err
	}

	return bodyRequest, nil
}

func kaspiTradePoints(organizationBIN string) (tradePointSt, error) {

	var bodyRequest tradePointSt

	client, err := getHttpClientTls()

	if err != nil {
		log.Fatal(err.Error())
	}

	resp, err := client.Get(viper.GetString("kaspiURL") + "partner/tradepoints/" + organizationBIN)
	if err != nil {
		panic(err)
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
