package myredis

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v9"
	"log"
	"time"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
// Discarded
func NewRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "101.42.224.110:6379",
		Password: "ubuntu",
		DB:       0,
	})
	return rdb
}

// AddNameToRedis add name and pkg to redis-string
func AddNameToRedis(pkgName string, pkgRepo string) {
	client := NewRedis()

	client.Set(context.Background(), pkgName, pkgRepo, 100*time.Hour)
}

// AddNameToRedis2 add name to redis-list
func AddNameToRedis2(pkgName string) {
	client := NewRedis()

	n := client.LPush(context.Background(), "pkgNames", pkgName)
	_, err := n.Result()
	if err != nil {
		log.Panicln(err)
	}
}

// GetPkg get name from key[pkgName],it uses redis-string
func GetPkg(pkgName string) (string, error) {
	client := NewRedis()
	result := client.Get(context.Background(), pkgName)
	pkgRepo, err := result.Result()
	if err != nil {
		error := errors.New("can't find this pkgName")
		return "", error
	}
	return pkgRepo, nil
}

// GetNamesFromRedis get all names from key[pkgNames],it uses redis-list
func GetNamesFromRedis() []string {
	client := NewRedis()
	result := client.LRange(context.Background(), "pkgNames", 0, -1)
	vals := result.Val()
	return vals
}

func GetReposFromRedis() []string {
	client := NewRedis()
	result := client.LRange(context.Background(), "pkgRepos", 0, -1)
	vals := result.Val()
	return vals
}
