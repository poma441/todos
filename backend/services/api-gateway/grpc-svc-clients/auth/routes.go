package auth

import (
	"todos/services/api-gateway/config"
	"todos/services/api-gateway/grpc-svc-clients/auth/routes"

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

func (svc *ServiceClient) SignUp(ctx *gin.Context) {
	routes.SignUp(ctx, svc.Client)
}
