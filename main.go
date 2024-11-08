package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	address := "localhost:16379"
	redis := NewRedis(address)

	key := "example_key"
	value := "example_value"

	if err := Set(redis, ctx, key, value, 0); err != nil {
		fmt.Printf("failed setting value: %v\n", err)
		return
	}

	result, err := Get(redis, ctx, key)
	if err != nil {
		fmt.Printf("Error getting value: %v\n", err)
		return
	}

	fmt.Printf("Result: %v\n", result)
}

func NewRedis(addr string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
		PoolSize: 1000,
	})
}

func Set(client *redis.Client, ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	if err := client.Set(ctx, key, value, expiration).Err(); err != nil {
		return err
	}
	return nil
}

func Get(client *redis.Client, ctx context.Context, key string) (interface{}, error) {
	value, err := client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return value, nil
}
