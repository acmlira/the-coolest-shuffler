package repository

import (
	"context"
	"encoding/json"
	"the-coolest-shuffler/internal/model"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type DecksRepository struct {
	Redis   *redis.Client
	Context context.Context
}

func NewDecksRepository(url string, database int, password string) *DecksRepository{
	return &DecksRepository{
		Redis: redis.NewClient(&redis.Options{
			Addr:     url,
			Password: password,
			DB:       database,
		}), 
		Context: context.Background(),
	}
}

func (c *DecksRepository) Set(deck *model.Deck) *model.Deck {
	value, err := json.Marshal(deck)
	if err != nil {
		log.Warn("Cannot insert key: value into cache")
	}
	err = c.Redis.Set(c.Context, deck.Id.String(), value, 0).Err()
	if err != nil {
		log.Error(err)
	}
	return deck
}

func (c *DecksRepository) Get(key uuid.UUID) *model.Deck {
	k := key.String()
	value, err := c.Redis.Get(c.Context, k).Result()
	if err == redis.Nil {
		log.Info(k + " does not exist, cannot get")
	} else if err != nil {
		log.Error(err)
	} else {
		log.Debug("Got " + value + " for " + k)
	}
	deck := &model.Deck{}
	json.Unmarshal([]byte(value), deck)
	return deck
}

func (c *DecksRepository) Del(key uuid.UUID) {
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
