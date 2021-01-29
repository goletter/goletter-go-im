package redis

import (
	"fmt"
	"github.com/go-redis/redis"
)

var (
	client *redis.Client
)

func ExampleNewClient() {
	client = redis.NewClient(&redis.Options{
		Addr:         "127.0.0.1:6379",
		Password:     "",
		DB:           0,
		PoolSize:     30,
		MinIdleConns: 30,
	})

	pong, err := client.Ping().Result()
	fmt.Println("初始化redis:", pong, err)
	// Output: PONG <nil>
}

func GetClient() (c *redis.Client) {

	return client
}
