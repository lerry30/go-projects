package redis

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisConnection struct {
	Client *redis.Client
}

func (rdb *RedisConnection) Close() {
	rdb.Client.Close()
}

func NewRedisConnection(addr, password string) *RedisConnection {
	opts := &redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0, // default DB, Redis supports 0-15

		// --- Connection Pool ---
		PoolSize:        10, // max connections in pool
		MinIdleConns:    3,  // keep at least 3 idle connections ready
		MaxIdleConns:    5,
		ConnMaxLifetime: 30 * time.Minute, // recycle connections after 30 min
		ConnMaxIdleTime: 5 * time.Minute,  // close idle connections after 5 min

		// --- Timeouts ---
		DialTimeout:  5 * time.Second, // timeout for establishing connection
		ReadTimeout:  3 * time.Second, // timeout for reads
		WriteTimeout: 3 * time.Second, // timeout for writes
		PoolTimeout:  4 * time.Second, // wait time for pool when all conns are busy

		// --- TLS (required by most managed Redis in production) ---
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	}

	client := redis.NewClient(opts)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatalf("failed to connect to Redis: %v", err)
	}

	fmt.Println("successfully connected to Redis")

	return &RedisConnection{
		Client: client,
	}
}

/*
func (r *RedisConnection) Set(key string, value any, ttl time.Duration) error {

}
*/