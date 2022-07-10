package service

import (
	"the-coolest-shuffler/internal/filter"
	"the-coolest-shuffler/internal/handler"
	"the-coolest-shuffler/internal/model"

	uuid "github.com/google/uuid"
)

type Cache interface {
	Set(key uuid.UUID, target interface{}) interface{}
	Get(key uuid.UUID, target interface{}) interface{}
}

type Database interface {
	Select(table string, target interface{}, filter map[string][]string) interface{}
}

type Shuffler struct{
	Cache    Cache
	Database Database
}

func NewShuffler(cache Cache, database Database) *Shuffler {
	return &Shuffler{
		Cache:    cache,
		Database: database,
	}
}

func (s *Shuffler) CreateNewDeck(shuffle bool, amount int, cardFilter *filter.CardFilter) *model.Deck {
	id := uuid.New()
	cards := s.Database.Select("cards", []model.Card{}, cardFilter.Filter).([]model.Card)
	deck := handler.Deck(model.NewDeck(id, cards, shuffle, amount))
	s.Cache.Set(id, deck)
	return deck
}

func (s *Shuffler) OpenDeck(id uuid.UUID) *model.Deck {
	deck := s.Cache.Get(id, &model.Deck{}).(*model.Deck)
	return deck
}

func (s *Shuffler) DrawCard(id uuid.UUID, count int) *model.Draw {
	deck := s.Cache.Get(id, &model.Deck{}).(*model.Deck)
	draw := handler.Draw(deck, model.NewDraw([]model.Card{}, count))
	s.Cache.Set(id, deck)
	return draw
}