package service

import (
	"context"
	"encoding/json"
	"math/rand"
	"time"

	"the-coolest-shuffler/internal/dao"
	"the-coolest-shuffler/internal/model"
	"the-coolest-shuffler/internal/repository"

	uuid "github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type Shuffler struct{}

func NewShuffler() *Shuffler {
	return &Shuffler{}
}

func (s *Shuffler) CreateNewDeck(ctx context.Context, shuffle bool, amount int, codes []string, values []string, suits []string) *model.Deck {
	var cards = dao.NewCard().Select(codes, values, suits)
	var deck = model.NewDeck(uuid.New(), cards, len(cards), shuffle, amount)
	var deckAmounted = applyAmount(deck)
	var deckShuffled = applyShuffle(deckAmounted)
	putOnCache(deckShuffled)
	return deckShuffled
}

func (s *Shuffler) OpenDeck(ctx context.Context, id uuid.UUID) *model.Deck {
	return getFromCache(id)
}

func (s *Shuffler) DrawCard(ctx context.Context, id uuid.UUID, count int) *model.Draw {
	deck := getFromCache(id)
	deck, draw := doDraw(deck, count)
	updateOnCache(id, deck)
	return draw
}

func putOnCache(deck *model.Deck) {
	encoded, err := json.Marshal(deck)
	if err == nil {
		repository.Set(deck.Id.String(), string(encoded))
	} else {
		log.Warn("Cannot insert key: value into cache")
	}
}

func getFromCache(id uuid.UUID) *model.Deck {
	deck := &model.Deck{}
	deckJson := repository.Get(id.String())
	json.Unmarshal([]byte(deckJson), deck)
	return deck
}

func updateOnCache(id uuid.UUID, deck *model.Deck) {
	repository.Del(id.String())
	putOnCache(deck)
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
