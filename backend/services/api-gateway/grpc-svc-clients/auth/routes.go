package auth

import (
	"todos/services/api-gateway/config"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	auth := r.Group("/auth")
	{
		auth.POST("/signup", svc.SignUp)
	}

	return svc
}
