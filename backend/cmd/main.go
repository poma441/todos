package main

import (
	"fmt"
	"log"
	"todos/config"
)

func main() {
	cfg, err := config.InitConfig("../config")
	if err != nil {
		log.Fatal("Ошибка инициализации")
	}
	fmt.Println(cfg.DB.Host)
	//server.Run(cfg)
}
