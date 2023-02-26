package config

import (
	"time"

	"github.com/spf13/viper"
)

type (

	// Config
	Config struct {
		Server `yaml:"server"`
		DB     `yaml:"db"`
		Redis  `yaml:"redis"`
		Token  `yaml:"token"`
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

	// Кэш redis
	Redis struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Password string `yaml:"password"`
		Db       int    `yaml:"db"`
	}

	Token struct {
		Keys    `yaml:"keys"`
		Access  `yaml:"access"`
		Refresh `yaml:"refresh"`
	}

	Keys struct {
		PrivateKey string `yaml:"privatekey"`
		PublicKey  string `yaml:"publickey"`
	}

	Access struct {
		TTL time.Duration `yaml:"ttl"`
	}

	Refresh struct {
		TTL               time.Duration `yaml:"ttl"`
		SecondsTTL        int           `yaml:"secondsTtl"`
		RefreshCookiePath string        `yaml:"refreshCookiePath"`
		LogoutCookiePath  string        `yaml:"logoutCookiePath"`
	}
)

func InitConfig(path string) (config Config, err error) {
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
