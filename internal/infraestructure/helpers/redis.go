package helpers

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"os"
)

var redisContext = context.Background()
var redisClient *redis.Client

func ConnectRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("MAGOBOT_REDIS_HOST"), os.Getenv("MAGOBOT_REDIS_PORT")),
		Password: os.Getenv("MAGOBOT_REDIS_PASSWORD"),
		DB:       0,
	})
}

func DisconnectRedis() {
	err := redisClient.Close()

	if err != nil {
		panic(err)
	}
}

func GetRedisClient() *redis.Client {
	return redisClient
}

func GetRedisContext() *context.Context {
	return &redisContext
}
