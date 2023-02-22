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
	router := gin.New()

	router.POST("/signup", h.Signup)
	router.POST("/signin", h.Signin)
	router.POST("/validate", h.UserIdentify)
	router.POST("/valid", h.ValidateUser)

	todos := router.Group("/todos")
	{
		todos.GET("/:userid", h.GetToDoItemsList)
		todos.POST("/:userid", h.AddToDoItem)
		todos.PUT("/:id", h.UpdateToDoItem)
		todos.DELETE("/:id", h.DeleteToDoItem)
	}

	return router
}
