package h_gin

import (
	"context"
	"net/http"

	"kaspi-qr/internal/domain/entities"
	"kaspi-qr/internal/domain/errs"

	"github.com/gin-gonic/gin"
)

func (h *Handler) tradePoints(c *gin.Context) {

	organizationBIN := c.Param("organizationBIN")

	req, err := h.usc.GetAllTradePoints(organizationBIN)

	if err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, errs.BadStatusCode, err.Error())
		return
	}

	req.Message = h.usc.SetMessageByStatusCode(req.StatusCode)

	c.JSON(http.StatusOK, req)
}

func (h *Handler) deviceRegistration(c *gin.Context) {

	var input entities.DeviceInputReg

	if err := c.BindJSON(&input); err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, errs.BadJson, err.Error())
		return
	}

	output, err := h.usc.DeviceRegistration(input)

	if err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, errs.BadStatusCode, err.Error())
		return
	}

	output.Message = h.usc.SetMessageByStatusCode(output.StatusCode)

	c.JSON(http.StatusOK, output)

	if output.StatusCode == 0 {
		err = h.usc.CreateDeviceRecord(context.Background(), input, output)
		if err != nil {
			errs.NewErrorResponse(c, http.StatusBadRequest, errs.NotImplemented, err.Error())
			return
		}
	}
}

func (h *Handler) deleteOrOffDevice(c *gin.Context) {
	var input entities.DeviceInputDel

	if err := c.BindJSON(&input); err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, errs.BadJson, err.Error())
		return
	}

	output, err := h.usc.DeleteOrOffDevice(input)

	if err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, errs.BadStatusCode, err.Error())
		return
	}

	output.Message = h.usc.SetMessageByStatusCode(output.StatusCode)

	c.JSON(http.StatusOK, output)

	if output.StatusCode == 0 {
		err = h.usc.DeleteDevice(c, input.OrganizationBin, input.DeviceToken)
		if err != nil {
			errs.NewErrorResponse(c, http.StatusBadRequest, errs.NotImplemented, err.Error())
			return
		}
	}
}
