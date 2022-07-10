package model

type Draw struct {
	Cards []Card `json:"cards"`
	Count int    `json:"-"`
}

func NewDraw(cards []Card, count int) *Draw {
	return &Draw{
		Cards: cards,
		Count: count,
	}
}