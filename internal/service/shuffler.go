package service

import (
	"encoding/json"

	"the-coolest-shuffler/internal/filter"
	"the-coolest-shuffler/internal/handler"
	"the-coolest-shuffler/internal/model"

	uuid "github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type Cache interface {
	Set(key string, value string)
	Get(key string) string
	Del(key string)
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
	cards := s.Database.Select("cards", []model.Card{}, cardFilter.Filter).([]model.Card)
	deck := model.NewDeck(uuid.New(), cards, shuffle, amount)
	handler.NewDeckHandler().Handle(deck)
	s.set(deck)
	return deck
}

func (s *Shuffler) OpenDeck(id uuid.UUID) *model.Deck {
	return s.get(id)
}

func (s *Shuffler) DrawCard(id uuid.UUID, count int) *model.Draw {
	deck := s.get(id)
	draw := []model.Card{}
	for i := 0; i < count; i++ {
		card := deck.Cards[0]
		draw = append(draw, card)
		deck.Cards = append(deck.Cards[:0], deck.Cards[1:]...)
	}
	deck.Remaining = len(deck.Cards)
	s.set(deck)
	return model.NewDraw(draw)
}

func (s *Shuffler) set(deck *model.Deck) {
	encoded, err := json.Marshal(deck)
	if err == nil {
		s.Cache.Set(deck.Id.String(), string(encoded))
	} else {
		log.Warn("Cannot insert key: value into cache")
	}
}

func (s *Shuffler) get(id uuid.UUID) *model.Deck {
	deck := &model.Deck{}
	json.Unmarshal([]byte(s.Cache.Get(id.String())), deck)
	return deck
}