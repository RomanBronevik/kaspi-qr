package rest

import (
	"github.com/gin-gonic/gin"
	"kaspi-qr/internal/adapters/provider/kaspi"
	"kaspi-qr/internal/domain/errs"
	"net/http"
)

func (h *Handler) QR(c *gin.Context) {
	body := c.Request.Body

	//var input entities.DeviceInputReg
	//
	//if err := c.BindJSON(&input); err != nil {
	//	errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	//	return
	//}

	output, err := kaspi.KaspiQR(body)

	if err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, output)
}

func (h *Handler) paymentLink(c *gin.Context) {
	body := c.Request.Body

	output, err := kaspi.KaspiPaymentLink(body)

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
