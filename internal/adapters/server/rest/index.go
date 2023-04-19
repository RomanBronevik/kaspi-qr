package rest

import (
	"context"
	"kaspi-qr/internal/adapters/logger"
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

	r.Use(MwRecovery(lg, nil))
	if withCors {
		r.Use(MwCors())
	}

	// handlers

	s := &St{lg: lg, ucs: ucs}

	// healthcheck
	r.GET("/healthcheck", func(c *gin.Context) { c.Status(http.StatusOK) })

	// city
	r.GET("/city", s.hCityList)
	r.POST("/city", s.hCityCreate)
	r.GET("/city/:id", s.hCityGet)
	r.PUT("/city/:id", s.hCityUpdate)
	r.DELETE("/city/:id", s.hCityDelete)

	// device
	r.GET("/device", s.hDeviceList)
	r.POST("/device", s.hDeviceCreate)
	r.GET("/device/:id", s.hDeviceGet)
	r.PUT("/device/:id", s.hDeviceUpdate)
	r.DELETE("/device/:id", s.hDeviceDelete)

	// ord
	r.GET("/ord", s.hOrdList)
	r.POST("/ord", s.hOrdCreate)
	r.GET("/ord/:id", s.hOrdGet)
	r.PUT("/ord/:id", s.hOrdUpdate)
	r.DELETE("/ord/:id", s.hOrdDelete)

	// payment
	r.GET("/payment", s.hPaymentList)
	r.POST("/payment", s.hPaymentCreate)
	r.GET("/payment/:id", s.hPaymentGet)
	r.PUT("/payment/:id", s.hPaymentUpdate)
	r.DELETE("/payment/:id", s.hPaymentDelete)

	return r
}

func (o *St) getRequestContext(c *gin.Context) context.Context {
	return context.Background()
}
