package service

import (
	"the-coolest-shuffler/internal/handler"
	"the-coolest-shuffler/internal/model"
	"the-coolest-shuffler/internal/request"

	uuid "github.com/google/uuid"
)

type CardsRepository interface {
	Set(deck *model.Deck) *model.Deck
	Get(key uuid.UUID) *model.Deck
}

type DecksRepository interface {
	Get(codes []string, values []string, suits []string) []model.Card
}

type Shuffler struct{
	CardsRepository CardsRepository
	DecksRepository DecksRepository
}

func NewShuffler(cardsRepository CardsRepository, decksRepository DecksRepository) *Shuffler {
	return &Shuffler{
		CardsRepository: cardsRepository,
		DecksRepository: decksRepository,
	}
}

func (s *Shuffler) Create(request *request.Deck) *model.Deck {
	deck := model.NewDeck(
		uuid.New(), 
		s.DecksRepository.Get(request.Codes, request.Values, request.Suits), 
		request.Shuffle, 
		request.Amount)

	deck = handler.Deck(deck)

	return s.CardsRepository.Set(deck)
}

func (s *Shuffler) Show(request *request.Deck) *model.Deck  {
	return s.CardsRepository.Get(request.Id)
}
