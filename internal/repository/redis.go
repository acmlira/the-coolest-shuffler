package repository

import (
	"context"

	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

var ctx = context.Background()
var r *redis.Client

func Cache(host string, port string, database int, password string) *redis.Client {
	r = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       database,
	})

	return r
}

func GetCache() *redis.Client {
	return r
}

func Set(key string, value string) {
	err := r.Set(ctx, key, value, 0).Err()
	if err != nil {
		log.Error(err)
	}
}

func Get(key string) string {
	value, err := r.Get(ctx, key).Result()
	if err == redis.Nil {
		log.Info(key + " does not exist, cannot get")
	} else if err != nil {
		log.Error(err)
	} else {
		log.Debug("Got " + value + " for " + key)
	}
	return value
}

func Del(key string) {
	err := r.Del(ctx, key).Err()
	if err == redis.Nil {
		log.Info(key + " does not exist, cannot delete")
	} else if err != nil {
		log.Error(err)
	} else {
		log.Info(key + " deleted")
	}
}
