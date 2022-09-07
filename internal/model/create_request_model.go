package model

type CreateRequest struct {
	Shuffle bool     `json:"shuffle" param:"shuffle" query:"shuffle"`
	Amount  int      `json:"amount" param:"amount" query:"amount"`
	Codes   []string `json:"codes" param:"codes" query:"codes"`
	Values  []string `json:"values" param:"values" query:"values"`
	Suits   []string `json:"suits" param:"suits" query:"suits"`
}
