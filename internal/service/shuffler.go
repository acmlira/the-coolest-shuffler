package service

import (
	"context"
	"encoding/json"
	"math/rand"
	"time"

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
	Select(table string, filter map[string][]string, model interface{}) []map[string]interface{}
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

func (s *Shuffler) CreateNewDeck(ctx context.Context, shuffle bool, amount int, codes []string, values []string, suits []string) *model.Deck {
	filter := map[string][]string{"code": codes, "value": values, "suit": suits}
	raw := s.Database.Select("cards", filter, model.Card{})
	cards := []model.Card{}

	for _, v := range raw {
		card := model.NewCard(v["code"].(string), v["value"].(string), v["suit"].(string))
		cards = append(cards, *card)
	}

	deck := model.NewDeck(uuid.New(), cards, len(cards), shuffle, amount)
	s.handleAmount(deck)
	s.handleShuffle(deck)
	s.set(deck)
	return deck
}

func (s *Shuffler) OpenDeck(ctx context.Context, id uuid.UUID) *model.Deck {
	return s.get(id)
}

func (s *Shuffler) DrawCard(ctx context.Context, id uuid.UUID, count int) *model.Draw {
	deck := s.get(id)
	draw := []model.Card{}
	for i := 0; i < count; i++ {
		card := deck.Cards[0]
		draw = append(draw, card)
		deck.Cards = append(deck.Cards[:0], deck.Cards[1:]...)
	}
	deck.Remaining = len(deck.Cards)
	s.del(id, deck)
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
	deckJson := s.Cache.Get(id.String())
	json.Unmarshal([]byte(deckJson), deck)
	return deck
}

func (s *Shuffler) del(id uuid.UUID, deck *model.Deck) {
	s.Cache.Del(id.String())
}

func (s *Shuffler) handleShuffle(deck *model.Deck) {
	if deck.Shuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(deck.Cards), func(i, j int) {
			deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
		})
	}
}

func (s *Shuffler) handleAmount(deck *model.Deck) {
	if deck.Amount == 0 {
		deck.Cards = []model.Card{}
		deck.Remaining = 0
	}
	if deck.Amount > 1 {
		cards := deck.Cards
		for i := 1; i < deck.Amount; i++ {
			deck.Cards = append(deck.Cards, cards...)
		}
		deck.Remaining = len(deck.Cards)
	}
}
