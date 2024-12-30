package cache

import (
	"context"
	"time"

	"github.com/kadekchresna/pastely/config"
	"github.com/redis/go-redis/v9"
)

type Cache interface {
	Set(context context.Context, key string, expiredInSecond int, value interface{}) error
	Get(context context.Context, key string, value interface{}) error
}

type CacheClient struct {
	Redis *redis.Client
}

func InitCache(config config.Config) Cache {

	return &CacheClient{
		Redis: redis.NewClient(&redis.Options{
			Addr:     config.RedisHost,
			Password: config.RedisPassword,
			Username: config.RedisUsername,
			DB:       config.RedisDB,
		}),
	}
}

func (c *CacheClient) Set(context context.Context, key string, expiredInSecond int, value interface{}) error {
	if err := c.Redis.SetNX(context, key, value, time.Duration(expiredInSecond)*time.Second).Err(); err != nil {
		return err
	}

	return nil
}

func (c *CacheClient) Get(context context.Context, key string, value interface{}) error {

	if err := c.Redis.Get(context, key).Scan(value); err != nil {
		return err
	}

	return nil
}
