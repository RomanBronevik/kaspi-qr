package rest

import (
	"github.com/gin-gonic/gin"
	"kaspi-qr/internal/adapters/provider/kaspi"
	"kaspi-qr/internal/domain/errs"
	"net/http"
)

func (h *Handler) details(c *gin.Context) {
	body := c.Request.Body

	output, err := kaspi.KaspiOperationDetails(body)

	if err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, output)
}

func (h *Handler) selfReturn(c *gin.Context) {
	body := c.Request.Body

	output, err := kaspi.KaspiReturnWithoutClient(body)

	if err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, output)
}
