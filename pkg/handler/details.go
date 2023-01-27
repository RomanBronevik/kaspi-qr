package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) tradePoints(c *gin.Context) {

	organizationBIN := c.Param("organizationBIN")

	req, err := kaspiTradePoints(organizationBIN)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Something went wrong",
			"error":   err.Error(),
		})
	}

	c.JSON(200, req)
}

func (h *Handler) deviceRegistration(c *gin.Context) {
	body := c.Request.Body

	output, err := kaspiDeviceRegistration(body)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(200, output)
}

func (h *Handler) deleteOrOffDevice(c *gin.Context) {
	body := c.Request.Body

	output, err := kaspiDeviceRegistration(body)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(200, output)
}
