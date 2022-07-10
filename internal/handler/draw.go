package handler

import "the-coolest-shuffler/internal/model"

func Draw(deck *model.Deck, draw *model.Draw) *model.Draw {
	handle(deck, draw)
	return draw
}

func handle(deck *model.Deck, draw *model.Draw) {
	for i := 0; i < draw.Count; i++ {
		card := deck.Cards[0]
		draw.Cards = append(draw.Cards, card)
		deck.Cards = append(deck.Cards[:0], deck.Cards[1:]...)
	}
	deck.Remaining = len(deck.Cards)
}

