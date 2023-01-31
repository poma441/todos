package main

import (
	"fmt"
	"log"
	"todos/config"
	"todos/internal/handler"
	"todos/internal/repository"
	"todos/internal/server"
	"todos/internal/service"
)

func main() {
	cfg, err := config.InitConfig("../config")
	if err != nil {
		log.Fatal("Ошибка инициализации ", err)
	}

	// Внедрение зависимостей
	repo := repository.NewRepository()
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	fmt.Printf("Попытка запуска сервера на: %s:%s\n", cfg.Server.Host, cfg.Server.Port)

	if err := server.Run(cfg, handler.InitRoutes()); err != nil {
		log.Fatal("Ошибка запуска сервера", err)
	}
}
