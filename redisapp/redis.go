package redisapp

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client

func StartRedis() {
	Rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// try ping
	ctx := context.Background()
	pong, err := Rdb.Ping(ctx).Result()
	fmt.Println("----- redis ping test ------")
	fmt.Println(pong, err)
	fmt.Println("---------- end -------------")
}
