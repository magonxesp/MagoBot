package helpers

import (
	"context"
	"github.com/go-redis/redis/v9"
	"os"
)

var redisContext = context.Background()
var redisClient = redis.NewClient(&redis.Options{
	Addr:     os.Getenv("MAGOBOT_REDIS_HOST"),
	Password: os.Getenv("MAGOBOT_REDIS_PASSWORD"),
	DB:       0,
})

func GetRedisClient() *redis.Client {
	return redisClient
}

func GetRedisContext() *context.Context {
	return &redisContext
}
