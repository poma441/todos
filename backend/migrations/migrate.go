package main

import (
	"log"
	"todos/config"
	"todos/internal/entity"
	"todos/pkg/postgres"
)

func main() {
	cfg, err := config.InitConfig("../config")
	if err != nil {
		log.Fatal("Ошибка файла конфигурации:", err)
	}

	db, err := postgres.NewConnectDB(cfg)
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	db.AutoMigrate(&entity.ToDoItem{})
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Student{})
}
