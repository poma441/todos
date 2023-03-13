package auth

import (
	"fmt"
	"todos/services/api-gateway/config"
	"todos/services/api-gateway/grpc-svc-clients/auth/pb"
	"todos/services/api-gateway/grpc-svc-clients/auth/routes"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	cc, err := grpc.Dial(c.AuthSvcServer.Host+":"+c.AuthSvcServer.Port, grpc.WithInsecure())
	if err != nil {
		fmt.Println("ошибка подключения к сервису auth: ", err)
	}

	return pb.NewAuthServiceClient(cc)
}

func (svc *ServiceClient) SignUp(ctx *gin.Context) {
	routes.SignUp(ctx, svc.Client)
}
