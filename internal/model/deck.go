package model

import "github.com/google/uuid"

type Deck struct {
	Id        uuid.UUID `json:"deck_id"`
	Remaining int       `json:"remaining"`
	Shuffle   bool      `json:"shuffled"`
	Amount    int       `json:"amount"`
	Cards     []Card    `json:"cards"`
}

func NewDeck(id uuid.UUID, cards []Card, remaining int, shuffle bool, amount int) *Deck {
	return &Deck{
		Id:        id,
		Remaining: remaining,
		Shuffle:   shuffle,
		Amount:    amount,
		Cards:     cards,
	}
}
