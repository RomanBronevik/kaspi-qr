package handlers

import (
	"kaspi-qr/internal/adapters/logger"
	"kaspi-qr/internal/domain/usecases"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	lg  logger.Lite
	usc *usecases.St
}

func NewHandler(lg logger.Lite, usc *usecases.St) *Handler {
	return &Handler{
		lg:  lg,
		usc: usc,
	}
}

func (h *Handler) InitRoutes(withCors bool) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	// middlewares

	router.Use(MwRecovery(h.lg, nil))
	if withCors {
		router.Use(MwCors())
	}

	// handlers

	api := router.Group("/return") // operation details and return payment without client
	{
		api.GET("/details", h.details)
		api.POST("/selfreturn", h.selfReturn)
	}

	device := router.Group("/device") // all movement with device and tradepoints
	{
		tradePoints := device.Group("/tradepoints")
		{
			tradePoints.GET("/:organizationBIN", h.tradePoints)
		}
		device.POST("/registration", h.deviceRegistration)
		device.POST("/delete", h.deleteOrOffDevice)
	}

	payment := router.Group("/payment") // qr token generation and payment link
	{
		payment.POST("/QR", h.QR)
		payment.POST("/link", h.paymentLink)

		status := payment.Group("/status")
		{
			status.GET("/:QrPaymentId", h.operationStatus)
			status.POST("/checkOrdersForPayment", h.checkOrdersForPayment)
		}
	}

	return router
}
