package repository

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var redisCtx = context.TODO()

const redisTTL = time.Hour * 24 * 31

type RedisConfig struct {
	Address  string
	Password string
	DB       int
}

func NewRedisCache(cfg RedisConfig) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	return rdb
}
