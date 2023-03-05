package main

import (
	"fmt"
	"log"
	"todos/config"
	"todos/internal/handler"
	"todos/internal/repository"
	"todos/internal/server"
	"todos/internal/service"
	"todos/pkg/postgres"
	"todos/pkg/redis"
)

func main() {
	cfg, err := config.InitConfig("../config")
	if err != nil {
		log.Fatal("Ошибка инициализации: ", err)
	}

	// Создание соединения с Postgres
	db, err := postgres.NewConnectDB(cfg)
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}

	// Создание соединения с кэшем Redis
	redisConn, err := redis.NewConnectRedis(&cfg)
	if err != nil {
		log.Fatal("Ошибка подключения к кэшу Redis:", err)
	}

	// Внедрение зависимостей
	repo := repository.NewRepository(db, redisConn)
	service := service.NewService(repo)
	handler := handler.NewHandler(service, &cfg)

	fmt.Printf("Попытка запуска сервера на: %s:%s\n", cfg.Server.Host, cfg.Server.Port)

	if err := server.Run(cfg, handler.InitRoutes()); err != nil {
		log.Fatal("Ошибка запуска сервера: ", err)
	}
}
