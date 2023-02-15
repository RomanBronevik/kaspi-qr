package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kaspi-qr/internal/adapters/provider/kaspi"
	"kaspi-qr/internal/domain/entities"
	"kaspi-qr/internal/domain/errs"
	"net/http"
)

func (h *Handler) tradePoints(c *gin.Context) {

	organizationBIN := c.Param("organizationBIN")

	req, err := kaspi.KaspiTradePoints(organizationBIN)

	if err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, req)
}

func (h *Handler) deviceRegistration(c *gin.Context) {

	var input entities.DeviceInputReg

	if err := c.BindJSON(&input); err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	output, err := kaspi.KaspiDeviceRegistration(input)

	if err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	output.Message = h.usc.SetMessageByStatusCode(output.StatusCode)

	c.JSON(http.StatusOK, output)

	dtoSt := entities.CreateDeviceDTO{
		Token:           output.Data.DeviceToken,
		DeviceId:        input.DeviceId,
		OrganizationBin: input.OrganizationBin,
		TradePointId:    input.TradePointId,
	}

	err = h.usc.CreateDevice(c, &dtoSt)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (h *Handler) deleteOrOffDevice(c *gin.Context) {
	var input entities.DeviceInputDel

	if err := c.BindJSON(&input); err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	output, err := kaspi.KaspiDeviceDelete(input)

	if err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	output.Message = h.usc.SetMessageByStatusCode(output.StatusCode)

	c.JSON(http.StatusOK, output)

	err = h.usc.DeleteDevice(c, input.OrganizationBin, input.DeviceToken)
	if err != nil {
		fmt.Println(err.Error())
	}
}
