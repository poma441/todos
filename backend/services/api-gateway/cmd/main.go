package main

import (
	"fmt"
	"log"
	"todos/services/api-gateway/config"
	"todos/services/api-gateway/grpc-svc-clients/auth"
	server "todos/services/api-gateway/internal/server"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.InitConfig("../config")
	if err != nil {
		log.Fatal("Ошибка инициализации api-gateway: ", err)
	}

	router := gin.New()
	//authSvc := *auth.RegisterRoutes(router, &cfg)
	auth.RegisterRoutes(router, &cfg)

	fmt.Printf("Попытка запуска сервера на: %s:%s\n", cfg.ApiGatewayServer.Host, cfg.ApiGatewayServer.Port)
	if err := server.Run(cfg, router); err != nil {
		log.Fatal("Ошибка запуска сервера api-gateway: ", err)
	}
}
