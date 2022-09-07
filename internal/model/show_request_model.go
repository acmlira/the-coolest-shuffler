package model

import (
	"github.com/google/uuid"
)

type ShowRequest struct {
	DeckId uuid.UUID `json:"deckId" param:"deckId" query:"deckId"`
}
