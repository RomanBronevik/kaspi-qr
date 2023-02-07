package handler

import (
	"github.com/gin-gonic/gin"
	"kaspi-qr/internal/entities"
	"kaspi-qr/internal/provider/kaspi"
	"kaspi-qr/pkg/errors"
	"net/http"
)

func (h *Handler) tradePoints(c *gin.Context) {

	organizationBIN := c.Param("organizationBIN")

	req, err := provider.KaspiTradePoints(organizationBIN)

	if err != nil {
		errors.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, req)
}

func (h *Handler) deviceRegistration(c *gin.Context) {

	var input entities.DeviceInputReg

	if err := c.BindJSON(&input); err != nil {
		errors.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	output, err := provider.KaspiDeviceRegistration(input)

	if err != nil {
		errors.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, output)
}

func (h *Handler) deleteOrOffDevice(c *gin.Context) {
	body := c.Request.Body

	output, err := provider.KaspiDeviceDelete(body)

	if err != nil {
		errors.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, output)
}
