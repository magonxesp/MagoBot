package utils

import (
	"context"
	"github.com/go-redis/redis/v9"
	"os"
)

var RedisContext = context.Background()

func CreateRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("MAGOBOT_REDIS_HOST"),
		Password: os.Getenv("MAGOBOT_REDIS_PASSWORD"),
		DB:       0,
	})
}
