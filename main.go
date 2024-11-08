package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	host string
	port string
}

func main() {
	ctx := context.Background()

	redisCfg := RedisConfig{
		host: "localhost",
		port: "16379",
	}
	redis, err := NewRedis(ctx, redisCfg)
	if err != nil {
		fmt.Printf("Error connecting redis: %v\n", err)
		return
	}

	key := "example_key"
	value := "example_value"

	if err := redis.Set(ctx, key, value, 0); err != nil {
		fmt.Printf("Error setting value: %v\n", err)
		return
	}

	result, err := redis.Get(ctx, key)
	if err != nil {
		fmt.Printf("Error getting value: %v\n", err)
		return
	}

	fmt.Printf("Result: %v\n", result)
}

type Redis struct {
	client *redis.Client
}

func NewRedis(ctx context.Context, cfg RedisConfig) (*Redis, error) {
	addr := fmt.Sprintf("%s:%s", cfg.host, cfg.port)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
		PoolSize: 1000,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &Redis{client: client}, nil
}

func (r *Redis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	if err := r.client.Set(ctx, key, value, expiration).Err(); err != nil {
		return err
	}
	return nil
}

func (r *Redis) Get(ctx context.Context, key string) (interface{}, error) {
	value, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return value, nil
}
