package model

type Draw struct {
	Cards []Card `json:"cards"`
}

func NewDraw(cards []Card) *Draw {
	return &Draw{Cards: cards}
}