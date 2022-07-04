package dao

import (
	"strconv"
	"sync"
	"the-coolest-shuffler/internal/model"
	"the-coolest-shuffler/internal/repository"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var instance *Card
var once sync.Once

type Card struct {
	database *gorm.DB
}

func NewCard() *Card {
	once.Do(func() {
		instance = &Card{database: repository.GetDatabase()}
	})
	return instance
}

func (card *Card) Select(codes []string, values []string, suits []string) []model.Card {
	var query *gorm.DB = card.database
	if len(codes) > 0 {
		query = query.Where("code IN ?", codes)
	}

	if len(values) > 0 {
		query = query.Where("value IN ?", values)
	}

	if len(suits) > 0 {
		query = query.Where("suit IN ?", suits)
	}

	var cards []model.Card
	var result = query.Find(&cards)
	resultHandler(result)
	return cards
}

func resultHandler(result *gorm.DB) {
	if result.Error != nil {
		log.Error(result.Error)
	} else {
		log.Debug("Retrived " + strconv.Itoa(int(result.RowsAffected)) + " objects")
	}
}
