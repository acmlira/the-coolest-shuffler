package model

import "github.com/google/uuid"

type Deck struct {
	Id        uuid.UUID `json:"deck_id"`
	Remaining int       `json:"remaining"`
	Shuffle   bool      `json:"shuffled"`
	Amount    int       `json:"amount"`
	Cards     []Card    `json:"cards"`
}

func NewDeck(cards []Card, shuffle bool, amount int) *Deck {
	return &Deck{
		Id:        uuid.New(),
		Remaining: len(cards),
		Shuffle:   shuffle,
		Amount:    amount,
		Cards:     cards,
	}
}
