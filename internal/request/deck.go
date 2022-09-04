package request

import "github.com/google/uuid"

type Deck struct {
	Id      uuid.UUID `json:"id" param:"id" query:"id"`
	Shuffle bool      `json:"shuffle" param:"shuffle" query:"shuffle"`
	Amount  int       `json:"amount" param:"amount" query:"amount"`
	Codes   []string  `json:"codes" param:"codes" query:"codes"`
	Values  []string  `json:"values" param:"values" query:"values"`
	Suits   []string  `json:"suits" param:"suits" query:"suits"`
	Count   int       `json:"count" param:"count" query:"count"`
}
