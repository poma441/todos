package handler

import (
	"todos/config"
	"todos/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
	config   *config.Config
}

func NewHandler(services *service.Service, cfg *config.Config) *Handler {
	return &Handler{
		services: services,
		config:   cfg,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/signup" /*h.ValidateUser,*/, h.SignUp)
		auth.POST("/signin" /*h.ValidateUser,*/, h.SignIn)
		auth.POST("/logout", h.Logout)
		auth.POST("/refresh", h.Refresh)
	}

	todos := router.Group("/todos", h.UserIdentify)
	{
		todos.GET("/:userid", h.GetToDoItemsList)
		todos.POST("/:userid", h.AddToDoItem)
		todos.PUT("/:id", h.UpdateToDoItem)
		todos.DELETE("/:id", h.DeleteToDoItem)
	}

	return router
}
