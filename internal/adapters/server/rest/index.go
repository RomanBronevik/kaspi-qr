package rest

import (
	"context"
	"kaspi-qr/internal/domain/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rendau/dop/adapters/logger"
	dopHttps "github.com/rendau/dop/adapters/server/https"
	swagFiles "github.com/swaggo/files"
	ginSwag "github.com/swaggo/gin-swagger"
)

type St struct {
	lg  logger.Lite
	ucs *usecases.St
}

func GetHandler(lg logger.Lite, ucs *usecases.St, withCors bool) http.Handler {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	// middlewares

	r.Use(dopHttps.MwRecovery(lg, nil))
	if withCors {
		r.Use(dopHttps.MwCors())
	}

	// handlers

	// doc
	r.GET("/doc/*any", ginSwag.WrapHandler(swagFiles.Handler, func(c *ginSwag.Config) {
		c.DefaultModelsExpandDepth = 0
		c.DocExpansion = "none"
	}))

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
	r.GET("/payment/:id", s.hPaymentGet)
	r.PUT("/payment/:id", s.hPaymentUpdate)
	r.DELETE("/payment/:id", s.hPaymentDelete)
	r.GET("/payment/:id/qr_picture", s.hPaymentGetQrPicture)

	// emu
	r.POST("/payment/:id/scan", s.hPaymentEmuPaymentScan)
	r.POST("/payment/:id/scan_error", s.hPaymentEmuPaymentScanError)
	r.POST("/payment/:id/confirm", s.hPaymentEmuPaymentConfirm)
	r.POST("/payment/:id/confirm_error", s.hPaymentEmuPaymentConfirmError)

	// trade_point
	r.GET("/trade_point", s.hTradePointList)

	// src
	r.GET("/src", s.hSrcList)
	r.POST("/src", s.hSrcCreate)
	r.GET("/src/:id", s.hSrcGet)
	r.PUT("/src/:id", s.hSrcUpdate)
	r.DELETE("/src/:id", s.hSrcDelete)

	return r
}

func (o *St) getRequestContext(c *gin.Context) context.Context {
	return context.Background()
}
