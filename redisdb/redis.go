package redisdb

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

var RedisDB *redis.Client

func Init() {
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func Get(key string) ([]byte, error) {
	if len(key) <= 0 {
		return nil, nil
	}

	val, err := RedisDB.Get(context.Background(), key).Bytes()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	}

	return val, err
}

func Set(key string, val []byte, exp time.Duration) error {
	if len(key) <= 0 || len(val) <= 0 {
		return nil
	}

	return RedisDB.Set(context.Background(), key, val, exp).Err()
}
