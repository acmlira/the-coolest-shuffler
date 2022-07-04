package service

import (
	"context"
	"encoding/json"

	"the-coolest-shuffler/internal/dao"
	"the-coolest-shuffler/internal/model"
	"the-coolest-shuffler/internal/repository"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type Deck struct {
}

func NewDeckService() *Deck {
	return &Deck{}
}

func (d *Deck) CreateNewDeck(ctx context.Context, shuffle bool, amount int, codes []string, values []string, suits []string) model.Deck {
	var cards = dao.NewCard().Select(codes, values, suits)
	var deck = model.NewDeck(cards, shuffle, amount)
	putOnCache(deck)
	return deck
}

func (d *Deck) OpenDeck(ctx context.Context, id uuid.UUID) model.Deck {
	return getFromCache(id)
}

func (d *Deck) DrawCard(ctx context.Context, id uuid.UUID, count int) model.Draw {
	deck := getFromCache(id)
	draw := deck.Draw(count)
	updateOnCache(id, deck)
	return draw
}

func putOnCache(deck model.Deck) {
	encoded, err := json.Marshal(deck)
	if err == nil {
		repository.Set(deck.Id.String(), string(encoded))
	} else {
		log.Warn("Cannot insert key: value into cache")
	}
}

func getFromCache(id uuid.UUID) model.Deck {
	deck := model.Deck{}
	deckJson := repository.Get(id.String())
	json.Unmarshal([]byte(deckJson), &deck)
	return deck
}

func updateOnCache(id uuid.UUID, deck model.Deck) {
	repository.Del(id.String())
	putOnCache(deck)
}
