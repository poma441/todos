package postgres

import (
	"fmt"
	"log"
	"todos/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnectDB(config config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.DB.Host, config.DB.Username, config.DB.Password, config.DB.Name, config.DB.Port, config.DB.SslMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Не удалось подключиться к БД:", err)
	}

	return db, nil
}
