package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func StartRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// try ping
	ctx := context.Background()
	pong, err := rdb.Ping(ctx).Result()
	fmt.Println("----- redis ping test ------")
	fmt.Println(pong, err)
	fmt.Println("---------- end -------------")
}
