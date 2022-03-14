package store

import (
	"context"
	"strconv"

	"github.com/gin-app-frame/internal/pkg/global"
	"github.com/go-redis/redis/v8"
)

func InitRedis() *redis.Client {
	cfg := global.App.Config
	zLogger = global.App.Logger

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host + ":" + strconv.Itoa(cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		zLogger.Panic("Redis connect failed: %s", err.Error())
		return nil
	}

	return client
}
