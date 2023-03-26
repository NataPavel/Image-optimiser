package handler

import (
	"github.com/gin-gonic/gin"
	"imageOptimisation/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// initializing endpoints
	api := router.Group("/api")
	{
		api.POST("/upload-image", h.uploadImage)
		api.GET("/get-image/:id", h.getImageById)
	}
	return router
}
