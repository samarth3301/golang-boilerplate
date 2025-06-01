package service

import (
	"github.com/redis/go-redis/v9"
	"os"
)

var RedisClient *redis.Client

func InitRedis() error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URL"),
	})
	
	return nil
}

func CloseRedis() {
	if RedisClient != nil {
		RedisClient.Close()
	}
} 