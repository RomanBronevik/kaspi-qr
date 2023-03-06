package rest

import (
	"github.com/gin-gonic/gin"
	"kaspi-qr/internal/domain/entities"
	"kaspi-qr/internal/domain/errs"
	"net/http"
)

func (h *Handler) tradePoints(c *gin.Context) {

	organizationBIN := c.Param("organizationBIN")

	req, err := h.kaspi.GetAllTradePoints(organizationBIN)

	if err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	req.Message = h.usc.SetMessageByStatusCode(req.StatusCode)

	c.JSON(http.StatusOK, req)
}

func (h *Handler) deviceRegistration(c *gin.Context) {

	var input entities.DeviceInputReg

	if err := c.BindJSON(&input); err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	output, err := h.kaspi.DeviceRegistration(input)

	if err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	output.Message = h.usc.SetMessageByStatusCode(output.StatusCode)

	c.JSON(http.StatusOK, output)

	if output.StatusCode == 0 {
		err = h.usc.CreateDeviceRecord(c, input, output)
		if err != nil {
			errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	}
}

func (h *Handler) deleteOrOffDevice(c *gin.Context) {
	var input entities.DeviceInputDel

	if err := c.BindJSON(&input); err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	output, err := h.kaspi.DeviceDelete(input)

	if err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	output.Message = h.usc.SetMessageByStatusCode(output.StatusCode)

	c.JSON(http.StatusOK, output)

	if output.StatusCode == 0 {
		err = h.usc.DeleteDevice(c, input.OrganizationBin, input.DeviceToken)
		if err != nil {
			errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	}
}
