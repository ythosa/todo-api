package repository

import "github.com/go-redis/redis/v8"

type RedisConfig struct {
	Address  string
	Password string
	DB       int
}

func NewRedisCache(cfg RedisConfig) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.Address,
		Password: cfg.Password,
		DB: cfg.DB,
	})

	return rdb
}
