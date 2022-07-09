package model

type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

func NewCard(code string, value string, suit string) *Card {
	return &Card{
		Value: value,
		Code:  code,
		Suit:  suit,
	}
}
