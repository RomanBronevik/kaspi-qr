package rest

import (
	"github.com/gin-gonic/gin"
	"kaspi-qr/internal/adapters/provider/kaspi"
	"kaspi-qr/internal/domain/entities"
	"kaspi-qr/internal/domain/errs"
	"net/http"
)

func (h *Handler) QR(c *gin.Context) {

	var input entities.QrTokenRequestInput

	if err := c.BindJSON(&input); err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	city, err := h.usc.FindOneCity(c, input.City)
	if err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	device, err := h.usc.FindOneDevice(c, city.OrganizationBin)

	if err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	qrToken := entities.QrTokenInput{
		OrganizationBin: city.OrganizationBin,
		DeviceToken:     device.Token,
		Amount:          input.Amount,
		ExternalId:      input.OrderNumber,
	}

	output, err := kaspi.CreateQrToken(qrToken)

	if err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, output)
}

func (h *Handler) paymentLink(c *gin.Context) {
	body := c.Request.Body

	output, err := kaspi.CreatePaymentLink(body)

	if err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, output)
}

func (h *Handler) operationStatus(c *gin.Context) {
	QrPaymentId := c.Param("QrPaymentId")

	req, err := kaspi.OperationStatus(QrPaymentId)

	if err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, req)
}
