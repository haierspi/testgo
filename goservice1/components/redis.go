package components

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func redisCreateClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "password",
		DB:       0,
	})

	// 通过 cient.Ping() 来检查是否成功连接到了 redis 服务器
	pong, err := client.Ping(context.TODO()).Result()
	fmt.Println(pong, err)

	return client
}
