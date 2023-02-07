package handler

import (
	"github.com/gin-gonic/gin"
	"kaspi-qr/internal/provider/kaspi"
	"kaspi-qr/pkg/errors"
	"net/http"
)

func (h *Handler) details(c *gin.Context) {
	body := c.Request.Body

	output, err := provider.KaspiOperationDetails(body)

	if err != nil {
		errors.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, output)
}

func (h *Handler) selfReturn(c *gin.Context) {
	body := c.Request.Body

	output, err := provider.KaspiReturnWithoutClient(body)

	if err != nil {
		errors.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, output)
}
