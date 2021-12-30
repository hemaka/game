package db

import (
	"log"

	"github.com/go-redis/redis/v7"
)

var redisAddr = "redis://root:@127.0.0.1:6379/"

func RedisRun() *redis.Client {

	log.Println("Contact to redis server: ", redisAddr)

	opt, err := redis.ParseURL(redisAddr)
	if err != nil {
		panic(err)
	}

	return redis.NewClient(opt)
}
