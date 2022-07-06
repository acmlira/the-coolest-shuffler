package repository

import (
	"context"

	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

type Cache struct {
	Redis   *redis.Client
	Context context.Context
}

func NewCache(url string, database int, password string) *Cache{
	return &Cache{
		Redis: redis.NewClient(&redis.Options{
			Addr:     url,
			Password: password,
			DB:       database,
		}), 
		Context: context.Background(),
	}
}

func (c *Cache) Set(key string, value string) {
	err := c.Redis.Set(c.Context, key, value, 0).Err()
	if err != nil {
		log.Error(err)
	}
}

func (c *Cache) Get(key string) string {
	value, err := c.Redis.Get(c.Context, key).Result()
	if err == redis.Nil {
		log.Info(key + " does not exist, cannot get")
	} else if err != nil {
		log.Error(err)
	} else {
		log.Debug("Got " + value + " for " + key)
	}
	return value
}

func (c *Cache) Del(key string) {
	err := c.Redis.Del(c.Context, key).Err()
	if err == redis.Nil {
		log.Info(key + " does not exist, cannot delete")
	} else if err != nil {
		log.Error(err)
	} else {
		log.Debug(key + " deleted")
	}
}
