package main

import (
	"log"
	"todos/services/auth-svc/config"
	"todos/services/auth-svc/models"
	"todos/services/auth-svc/pkg/postgres"
)

func main() {
	cfg, err := config.InitConfig("../config")
	if err != nil {
		log.Fatal("Ошибка файла конфигурации: ", err)
	}

	db, err := postgres.NewConnectDB(cfg)
	if err != nil {
		log.Fatal("Ошибка подключения к postgres: ", err)
	}

	db.AutoMigrate(&models.User{})
}
