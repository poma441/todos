package config

import (
	"github.com/spf13/viper"
)

type (

	// Config
	Config struct {
		Server `yaml:"server"`
		DB     `yaml:"db"`
	}

	// Server
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}

	// DB
	DB struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Name     string `yaml:"name"`
		Password string `yaml:"password"`
		SslMode  string `yaml:"sslmode"`
	}
)

func InitConfig(path string) (config Config, err error) {
	// Инициализация конфига
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
