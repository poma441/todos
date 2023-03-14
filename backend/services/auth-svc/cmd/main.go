package main

import (
	"fmt"
	"log"
	"net"
	"todos/services/auth-svc/config"
	"todos/services/auth-svc/internal/pb"
	"todos/services/auth-svc/internal/repository"
	"todos/services/auth-svc/internal/service"
	"todos/services/auth-svc/pkg/postgres"
	"todos/services/auth-svc/pkg/redis"

	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.InitConfig("../config")
	if err != nil {
		log.Fatal("не удалось инициализировать сервис auth: ", err)
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
	service := service.NewService(repo, &cfg)

	// Запуск TCP сервера
	lis, err := net.Listen("tcp", cfg.Server.Host+":"+cfg.Server.Port)
	if err != nil {
		log.Fatal("Не удалось запустить tcp сервер для сервиса auth: ", err)
	}

	// Создание grpc сервера
	grpcServer := grpc.NewServer()

	// Связывание созданного grpc сервера и слоя бизнес логики
	pb.RegisterAuthServiceServer(grpcServer, service.Authorization)

	fmt.Println("Сервер аутентификации запущен")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Не удалось запустить grpc сервер для сервиса auth: ", err)
	}
}
