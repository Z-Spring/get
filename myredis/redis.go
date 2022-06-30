package myredis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"log"
)

func NewRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "101.42.224.110:6379",
		DB:   0,
	})
	return rdb
}

func AddPkgToRedis(pkgName string, pkgRepo string) {
	client := NewRedis()

	n := client.LPush(context.Background(), "pkgNames", pkgName)
	_, err := n.Result()
	if err != nil {
		log.Panicln(err)
	}

	r := client.LPush(context.Background(), "pkgRepos", pkgRepo)
	_, err2 := r.Result()
	if err != nil {
		log.Panicln(err2)
	}
	fmt.Println("done")
}

func GetPkgNameFromRedis() []string {
	client := NewRedis()
	result := client.LRange(context.Background(), "pkgNames", 0, -1)
	vals := result.Val()
	return vals
}

func GetPkgRepoFromRedis() []string {
	client := NewRedis()
	result := client.LRange(context.Background(), "pkgRepos", 0, -1)
	vals := result.Val()
	return vals
}
