package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
)

type tradePoint struct {
	BIN int
}

func (h *Handler) tradepoints(c *gin.Context) {

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

}

func (h *Handler) deleteOrOffDevice(c *gin.Context) {

}

//type tradePoint struct {
//	TradePointId int
//	TradePointName string
//}

func kaspiTradePoints(organizationBIN string) ([]byte, error) {
	client, err := getHttpClientTsl()

	if err != nil {
		log.Fatal(err.Error())
	}

	resp, err := client.Get("https://mtokentest.kaspi.kz:8545/r3/v01/partner/tradepoints/" + organizationBIN)
	if err != nil {
		panic(err)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	fmt.Println(string(bytes))

	//jsonErr := json.Unmarshal(body, &people1)
	//if jsonErr != nil {
	//	log.Fatal(jsonErr)
	//}

	//fmt.Println(people1.Number)

	return bytes, nil

	//testbdy := new(bytes.Buffer)
	//testbdy.ReadFrom(resp.Body)
	//fmt.Println(testbdy)
}
