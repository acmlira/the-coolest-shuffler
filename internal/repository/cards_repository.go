package repository

import (
	"strconv"
	"the-coolest-shuffler/internal/model"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CardsRepository struct {
	database *gorm.DB
}

func NewCardsRepository(dsn string) *CardsRepository {
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error(err)
		return &CardsRepository{}
	}
	return &CardsRepository{database: database}
}

func (d *CardsRepository) Get(codes []string, values []string, suits []string) []model.Card {
	var query *gorm.DB = d.database

	if len(codes) > 0 {
		query = query.Where("code IN ?", codes)
	}

	if len(values) > 0 {
		query = query.Where("value IN ?", values)
	}

	if len(suits) > 0 {
		query = query.Where("suit IN ?", suits)
	}

	var cards = []model.Card{}
	var result = query.Table("cards").Find(&cards)
	if result.Error != nil {
		log.Error(result.Error)
	} else {
		log.Debug("Retrived " + strconv.Itoa(int(result.RowsAffected)) + " objects")
	}
	return cards
}
