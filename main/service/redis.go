package service

import (
	"golang-boilerplate/main/config"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.AppConfig.Redis.URL,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return nil
}

func CloseRedis() {
	if RedisClient != nil {
		err := RedisClient.Close()
		if err != nil {
			return
		}
	}
}
