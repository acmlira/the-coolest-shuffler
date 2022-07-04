package service

import (
	"context"
	"math/rand"
	"encoding/json"
	"time"

	"the-coolest-shuffler/internal/dao"
	"the-coolest-shuffler/internal/model"
	"the-coolest-shuffler/internal/repository"

	uuid "github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type Shuffler struct {}

func NewShuffler() *Shuffler {
	return &Shuffler{}
}

func (s *Shuffler) CreateNewDeck(ctx context.Context, shuffle bool, amount int, codes []string, values []string, suits []string) model.Deck {
	var cards = dao.NewCard().Select(codes, values, suits)
	var deck = *model.NewDeck(cards, shuffle, amount)
	deck = applyShuffle(applyAmount(deck))
	putOnCache(deck)
	return deck
}

func (s *Shuffler) OpenDeck(ctx context.Context, id uuid.UUID) model.Deck {
	return getFromCache(id)
}

func (s *Shuffler) DrawCard(ctx context.Context, id uuid.UUID, count int) model.Draw {
	deck := getFromCache(id)
	draw := draw(deck, count)
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

func draw(deck model.Deck, count int) model.Draw {
	var cards []model.Card
	for i := 0; i < count; i++ {
		card := deck.Cards[0]
		cards = append(cards, card)
		deck.Cards = append(deck.Cards[:0], deck.Cards[1:]...)
		deck.Remaining = len(deck.Cards)
	}
	return model.Draw{Cards: cards}
}

func applyShuffle(deck model.Deck) model.Deck {
	if deck.Shuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(deck.Cards), func(i, j int) { 
			deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i] 
		})
	}
	return deck
}

func applyAmount(deck model.Deck) model.Deck {
	if deck.Amount == 0 {
		return deck
	}
	if deck.Amount > 1 {
		for i := 1; i < deck.Amount; i++ {
			deck.Cards = append(deck.Cards, deck.Cards...)
		}
	}
	return deck
}