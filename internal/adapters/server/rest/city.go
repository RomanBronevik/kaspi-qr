package rest

import (
	"github.com/gin-gonic/gin"
	"kaspi-qr/internal/adapters/provider/mechtaWeb"
	"kaspi-qr/internal/domain/errs"
	"net/http"
)

func (h *Handler) UpdateCities(c *gin.Context) {

	cities := mechtaWeb.GetCitiesFromSite(c)

	err := h.usc.UpdateCities(c, cities)

	if err != nil {
		errs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}
