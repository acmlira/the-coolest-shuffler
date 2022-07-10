package repository

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
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

func (c *Cache) Set(key uuid.UUID, target interface{}) interface{} {
	value, err := json.Marshal(target)
	if err != nil {
		log.Warn("Cannot insert key: value into cache")
	}
	err = c.Redis.Set(c.Context, key.String(), value, 0).Err()
	if err != nil {
		log.Error(err)
	}
	return target
}

func (c *Cache) Get(key uuid.UUID, target interface{}) interface{} {
	k := key.String()
	value, err := c.Redis.Get(c.Context, k).Result()
	if err == redis.Nil {
		log.Info(k + " does not exist, cannot get")
	} else if err != nil {
		log.Error(err)
	} else {
		log.Debug("Got " + value + " for " + k)
	}
	json.Unmarshal([]byte(value), target)
	return target
}

func (c *Cache) Del(key uuid.UUID) {
	k := key.String()
	err := c.Redis.Del(c.Context, k).Err()
	if err == redis.Nil {
		log.Info(k + " does not exist, cannot delete")
	} else if err != nil {
		log.Error(err)
	} else {
		log.Debug(k + " deleted")
	}
}
