package redis

import (
	"fmt"
	"todos/services/auth-svc/config"

	"github.com/go-redis/redis"
)

func NewConnectRedis(cfg *config.Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.Db,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return client, err
	}

	return client, nil
}
