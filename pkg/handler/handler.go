package handler

import (
	"github.com/gin-gonic/gin"
	"kaspi-qr/pkg/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/return") // для регистрации и авторизации
	{
		api.GET("/details", h.details)
		api.POST("/selfreturn", h.selfReturn)
	}

	device := router.Group("/device") // для работы с endpoint со списками и их задачами
	{
		tradePoints := device.Group("/tradepoints")
		{
			tradePoints.GET("/:organizationBIN", h.tradePoints)
		}
		device.POST("/registration", h.deviceRegistration)
		device.POST("/delete", h.deleteOrOffDevice)
	}

	payment := router.Group("/payment") // для работы с endpoint со списками и их задачами
	{
		payment.GET("/QR", h.QR)
		payment.POST("/payment-link", h.paymentLink)
		payment.POST("/status", h.operationStatus)
	}

	return router
}
