package model

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Deck struct {
	Id        uuid.UUID `json:"deck_id"`
	Remaining int       `json:"remaining"`
	Shuffle   bool      `json:"shuffled"`
	Cards     []Card    `json:"cards"`
}

func NewDeck(cards []Card, shuffle bool, amount int) Deck {
	if shuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
	}

	if amount == 0 {
		cards = []Card{}
	}

	if amount > 1 {
		for i := 1; i < amount; i++ {
			cards = append(cards, cards...)
		}
	}

	return Deck{
		Id:        uuid.New(),
		Remaining: len(cards),
		Shuffle:   shuffle,
		Cards:     cards,
	}
}

func (deck *Deck) Draw(count int) Draw {
	var cards []Card
	for i := 0; i < count; i++ {
		card := deck.Cards[0]
		cards = append(cards, card)
		deck.Cards = append(deck.Cards[:0], deck.Cards[1:]...)
		deck.Remaining = len(deck.Cards)
	}
	return Draw{Cards: cards}
}
