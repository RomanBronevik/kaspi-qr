package rest

import (
	"github.com/gin-gonic/gin"
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

	city, err := h.usc.FindOneCityByCityCode(c, input.Code)
	if err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, "City not found")
		return
	}

	cityDoesntExist := h.usc.IsEmptyCity(city)

	if cityDoesntExist {
		errs.NewErrorResponse(c, http.StatusBadRequest, "City not found")
		return
	}

	device, err := h.usc.FindOneDevice(c, city.OrganizationBin)

	if err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, "Device not exist")
		return
	}

	qrToken := entities.KaspiPaymentInput{
		OrganizationBin: city.OrganizationBin,
		DeviceToken:     device.Token,
		Amount:          input.Amount,
		ExternalId:      input.OrderNumber,
	}

	output, err := h.kaspi.CreateQrToken(qrToken)

	if err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	output.Message = h.usc.SetMessageByStatusCode(output.StatusCode)

	c.JSON(http.StatusOK, output)

	if output.StatusCode == 0 {
		err = h.usc.QrCreateOrderRecords(c, qrToken, output)

		if err != nil {
			errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	}
}

func (h *Handler) paymentLink(c *gin.Context) {
	var input entities.PaymentLinkRequestInput

	if err := c.BindJSON(&input); err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	city, err := h.usc.FindOneCityByCityCode(c, input.Code)
	if err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	device, err := h.usc.FindOneDevice(c, city.OrganizationBin)

	if err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	paymentLink := entities.KaspiPaymentInput{
		OrganizationBin: city.OrganizationBin,
		DeviceToken:     device.Token,
		Amount:          input.Amount,
		ExternalId:      input.OrderNumber,
	}

	output, err := h.kaspi.CreatePaymentLink(paymentLink)

	if err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	output.Message = h.usc.SetMessageByStatusCode(output.StatusCode)

	c.JSON(http.StatusOK, output)

	if output.StatusCode == 0 {

		err = h.usc.LinkCreateOrderRecords(c, paymentLink, output)
		if err != nil {
			errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	}
}

func (h *Handler) operationStatus(c *gin.Context) {
	QrPaymentId := c.Param("QrPaymentId")

	req, err := h.kaspi.OperationStatus(QrPaymentId)

	if err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, req)
}

func (h *Handler) checkOrdersForPayment(c *gin.Context) {
	err := h.usc.CheckPaymentStatus(c)
	if err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, entities.StatusSt{"OK"})
}
