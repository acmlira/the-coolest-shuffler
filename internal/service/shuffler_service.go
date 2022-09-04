package service

import (
	"the-coolest-shuffler/internal/handler"
	"the-coolest-shuffler/internal/model"

	uuid "github.com/google/uuid"
)

type DecksRepository interface {
	Set(deck *model.Deck) *model.Deck
	Get(key uuid.UUID) *model.Deck
}

type CardsRepository interface {
	Get(codes []string, values []string, suits []string) []model.Card
}

type ShufflerService struct {
	CardsRepository CardsRepository
	DecksRepository DecksRepository
}

func NewShufflerService(cardsRepository CardsRepository, decksRepository DecksRepository) *ShufflerService {
	return &ShufflerService{
		CardsRepository: cardsRepository,
		DecksRepository: decksRepository,
	}
}

func (s *ShufflerService) Create(request *model.Request) *model.Deck {
	deck := model.NewDeck(
		uuid.New(),
		s.CardsRepository.Get(request.Codes, request.Values, request.Suits),
		request.Shuffle,
		request.Amount)

	deck = handler.Deck(deck)

	return s.DecksRepository.Set(deck)
}

func (s *ShufflerService) Show(request *model.Request) *model.Deck {
	return s.DecksRepository.Get(request.Id)
}
