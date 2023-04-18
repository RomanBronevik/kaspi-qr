package rest

import (
	"kaspi-qr/internal/adapters/logger"
	"kaspi-qr/internal/adapters/server/rest/handlers"
	"kaspi-qr/internal/domain/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type St struct {
	lg  logger.Lite
	ucs *usecases.St
}

func GetHandler(lg logger.Lite, ucs *usecases.St, withCors bool) http.Handler {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	// middlewares

	r.Use(handlers.MwRecovery(lg, nil))
	if withCors {
		r.Use(handlers.MwCors())
	}

	// handlers

	s := &St{lg: lg, ucs: ucs}

	// healthcheck
	r.GET("/healthcheck", func(c *gin.Context) { c.Status(http.StatusOK) })

	// organisation
	r.GET("/organisation", s.hOrganisationList)
	r.POST("/organisation", s.hOrganisationCreate)
	r.GET("/organisation/:id", s.hOrganisationGet)
	r.PUT("/organisation/:id", s.hOrganisationUpdate)
	r.DELETE("/organisation/:id", s.hOrganisationDelete)

	return r
}
