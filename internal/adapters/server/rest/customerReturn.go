package rest

import (
	"github.com/gin-gonic/gin"
	"kaspi-qr/internal/domain/entities"
	"kaspi-qr/internal/domain/errs"
	"net/http"
)

func (h *Handler) details(c *gin.Context) {
	body := c.Request.Body

	output, err := h.kaspi.KaspiOperationDetails(body)

	if err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, output)
}

func (h *Handler) selfReturn(c *gin.Context) {
	var input entities.ReturnRequestInput

	if err := c.BindJSON(&input); err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	output, err := h.kaspi.KaspiReturnWithoutClient(input)

	if err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, output)

	if output.StatusCode == 0 {
		err = h.usc.ReturnOrder(c, input.QrPaymentId)
		if err != nil {
			errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	}
}
