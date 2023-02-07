package handler

import (
	"github.com/gin-gonic/gin"
	"kaspi-qr/internal/provider/kaspi"
	"kaspi-qr/pkg/errors"
	"net/http"
)

func (h *Handler) QR(c *gin.Context) {
	body := c.Request.Body

	output, err := provider.KaspiQR(body)

	if err != nil {
		errors.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, output)
}

func (h *Handler) paymentLink(c *gin.Context) {
	body := c.Request.Body

	output, err := provider.KaspiPaymentLink(body)

	if err != nil {
		errors.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, output)
}

func (h *Handler) operationStatus(c *gin.Context) {
	QrPaymentId := c.Param("QrPaymentId")

	req, err := provider.OperationStatus(QrPaymentId)

	if err != nil {
		errors.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, req)
}
