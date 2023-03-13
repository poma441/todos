package auth

import (
	"fmt"
	"todos/services/api-gateway/config"
	"todos/services/api-gateway/grpc-svc-clients/auth/pb"

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
