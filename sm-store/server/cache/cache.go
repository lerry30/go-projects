package cache

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

var ErrCacheMiss = errors.New("cache miss")

// SET
func Set(rdb *redis.Client, key string, value any, ttl time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return rdb.Set(ctx, key, data, ttl).Err()
}

// GET
func Get[T any](rdb *redis.Client, key string) (T, error) {
	var result T

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	data, err := rdb.Get(ctx, key).Bytes()
	if errors.Is(err, redis.Nil) {
		return result, ErrCacheMiss // key doesn't exist
	}
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

// DELETE (e.g. when data is updated)
func Delete(rdb *redis.Client, key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return rdb.Del(ctx, key).Err()
}
