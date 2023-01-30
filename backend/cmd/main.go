package main

import (
	"todos/config"
	"todos/internal/server"
)

func main() {
	if err := server.Run(config.InitConfig()); err != nil {
		return
	}
}
