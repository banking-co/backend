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

func GetAll(keys []string) (*[][]byte, error) {
	if len(keys) == 0 {
		return nil, nil
	}

	pipe := RedisDB.Pipeline()
	ctx := context.Background()

	cmders := make([]*redis.StringCmd, len(keys))
	for i, key := range keys {
		cmders[i] = pipe.Get(ctx, key)
	}

	_, err := pipe.Exec(ctx)
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	}

	results := make([][]byte, len(keys))
	allNil := true

	for i, cmd := range cmders {
		result, err := cmd.Bytes()
		if err != nil && err != redis.Nil {
			results[i] = nil
		} else {
			results[i] = result
			if result != nil {
				allNil = false
			}
		}
	}

	if allNil {
		return nil, nil
	}

	return &results, nil
}

func Set(key string, val []byte, exp time.Duration) error {
	if len(key) <= 0 || len(val) <= 0 {
		return nil
	}

	return RedisDB.Set(context.Background(), key, val, exp).Err()
}

func SetAll(pairs map[string][]byte, exp time.Duration) error {
	if len(pairs) == 0 {
		return nil
	}

	pipe := RedisDB.Pipeline()
	ctx := context.Background()

	for key, val := range pairs {
		if len(key) > 0 && len(val) > 0 {
			pipe.Set(ctx, key, val, exp)
		}
	}

	_, err := pipe.Exec(ctx)
	return err
}
