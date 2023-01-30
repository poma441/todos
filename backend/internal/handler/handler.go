package handler

import (
	"todos/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := new(gin.Engine)

	// Привязка маршрутов к обработчикам

	return router
}
