package config

import (
	"github.com/spf13/viper"
)

type (

	// Config
	Config struct {
		ApiGatewayServer `yaml:"apiGatewayServer"`
		AuthSvcServer    `yaml:"authSvcServer"`
		TodosSvcServer   `yaml:"todosSvcServer"`
	}

	// Api-Gateway Server
	ApiGatewayServer struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}

	// Auth-Svc Server
	AuthSvcServer struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}

	// Todos-Svc Server
	TodosSvcServer struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
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
