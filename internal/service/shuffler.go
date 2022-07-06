package service

import (
	"context"
	"encoding/json"
	"fmt"
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
	var m = make(map[string][]string, 10)
	m["code"] = codes
	m["value"] = values
	m["suit"] = suits

	var cardsRaw = s.Database.Select("cards", m, model.Card{})

	cards := []model.Card{}
	for _, v := range cardsRaw {
		card := model.Card{
			Code:  v["code"].(string),
			Suit:  v["suit"].(string),
			Value: v["value"].(string),
		}
		cards = append(cards, card)
	}

	fmt.Printf("cards: %v\n", cards)

	var deck = model.NewDeck(uuid.New(), cards, len(cards), shuffle, amount)
	var deckAmounted = applyAmount(deck)
	var deckShuffled = applyShuffle(deckAmounted)
	s.putOnCache(deckShuffled)
	return deckShuffled
}

func (s *Shuffler) OpenDeck(ctx context.Context, id uuid.UUID) *model.Deck {
	return s.getFromCache(id)
}

func (s *Shuffler) DrawCard(ctx context.Context, id uuid.UUID, count int) *model.Draw {
	deck := s.getFromCache(id)
	deck, draw := doDraw(deck, count)
	s.updateOnCache(id, deck)
	return draw
}

func (s *Shuffler) putOnCache(deck *model.Deck) {
	encoded, err := json.Marshal(deck)
	if err == nil {
		s.Cache.Set(deck.Id.String(), string(encoded))
	} else {
		log.Warn("Cannot insert key: value into cache")
	}
}

func (s *Shuffler) getFromCache(id uuid.UUID) *model.Deck {
	deck := &model.Deck{}
	deckJson := s.Cache.Get(id.String())
	json.Unmarshal([]byte(deckJson), deck)
	return deck
}

func (s *Shuffler) updateOnCache(id uuid.UUID, deck *model.Deck) {
	s.Cache.Del(id.String())
	s.putOnCache(deck)
}

func doDraw(deck *model.Deck, count int) (*model.Deck, *model.Draw) {
	var cards = deck.Cards
	var drawCards = []model.Card{}

	for i := 0; i < count; i++ {
		drawCard := cards[0]
		drawCards = append(drawCards, drawCard)
		cards = append(cards[:0], cards[1:]...)
	}
	return model.NewDeck(deck.Id, cards, len(cards), deck.Shuffle, deck.Amount), model.NewDraw(drawCards)
}

func applyShuffle(deck *model.Deck) *model.Deck {
	if deck.Shuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(deck.Cards), func(i, j int) {
			deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
		})
	}
	return model.NewDeck(deck.Id, deck.Cards, deck.Remaining, deck.Shuffle, deck.Amount)
}

func applyAmount(deck *model.Deck) *model.Deck {
	var amount = deck.Amount
	var cards = deck.Cards
	if amount == 0 {
		return deck
	}
	if amount > 1 {
		for i := 1; i < amount; i++ {
			cards = append(cards, cards...)
		}
	}
	return model.NewDeck(deck.Id, cards, len(cards), deck.Shuffle, deck.Amount)
}
