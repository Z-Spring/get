package myredis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"log"
)

var Args = []string{
	"gin",
	"iris",
	"gjson",
	"elastic",
	"gorm",
	"goquery",
	"redis",
	"beego",
	"mysql",
	"oracle",
	"viper",
	"cobra",
	"jwt",
	"zap",
	"simpletable",
	"postgresql",
}
var Repos = []string{
	"github.com/gin-gonic/gin",
	"github.com/tidwall/gjson",
	"github.com/PuerkitoBio/goquery",
	"github.com/beego/beego/v2@latest",
	"github.com/beego/beego",
	"github.com/go-redis/redis/v9",
	"github.com/go-redis/redis/v8",
	"github.com/go-redis/redis/v7",
	"github.com/elastic/go-elasticsearch/v8",
	"github.com/elastic/go-elasticsearch/v7",
	"github.com/golang-jwt/jwt/v4",
}

func NewRedis() *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr: "101.42.224.110:6379",
		DB:   0,
	})
	// 连接远程Redis服务器
	/*opt, _ := redis.ParseURL("rediss://:562a2ad2986f407097761e0d94ef507f@us1-vocal-bulldog-37309.upstash.io:37309")
	rdb := redis.NewClient(opt)*/
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

func AddPkgToRedis2() {
	client := NewRedis()
	for _, v := range Args {

		n := client.LPush(context.Background(), "pkgNames", v)
		_, err := n.Result()
		if err != nil {
			log.Panicln(err)
		}
	}
	for _, v := range Repos {
		n := client.LPush(context.Background(), "pkgRepos", v)
		_, err := n.Result()
		if err != nil {
			log.Panicln(err)
		}
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
