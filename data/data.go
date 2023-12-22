package data

import (
	"context"
	"time"

	"github.com/go-redis/redis/v9"
)

func SetRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:8080",
		DB:   0,
	})

	return rdb
}

func AddDataToRedis(key string, value string) {
	client := SetRedis()
	client.Set(context.Background(), key, value, 24 * time.Hour)

}
