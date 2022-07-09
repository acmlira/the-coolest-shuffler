package handler

import (
	"math/rand"
	"time"
	
	"the-coolest-shuffler/internal/model"
)

type DeckHandler struct {}

func NewDeckHandler() *DeckHandler {
	return &DeckHandler{}
}

func (dh *DeckHandler) Handle(deck *model.Deck) {
	dh.amount(deck)
	dh.shuffle(deck)
}

func (dh *DeckHandler) shuffle(deck *model.Deck) {
	if deck.Shuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(deck.Cards), func(i, j int) {
			deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
		})
	}
}

func (dh *DeckHandler) amount(deck *model.Deck) {
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
