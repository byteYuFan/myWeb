package redisMiddleware

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

var Client *redis.Client
var Ctx context.Context

func init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "docker:6379",
		Password: "",
		DB:       0,
	})
	Ctx = context.Background()
}

func Insert() {
	pong, err := Client.Ping(context.Background()).Result()
	log.Println(pong, err)
}
func StoreCode(key string, field string, value interface{}, expiration time.Duration) error {
	cmd := Client.HSet(Ctx, key, field, value)
	if err := cmd.Err(); err != nil {
		return err
	}
	// set key expiration
	if err := Client.Expire(Ctx, key, expiration).Err(); err != nil {
		return err
	}

	return nil
}
func GetEmailCode(key, filed string) (string, error) {
	value, err := Client.HGet(Ctx, key, filed).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}
