package service

import (
	"math/rand"
	"the-coolest-shuffler/internal/model"
	"time"

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

func (s *ShufflerService) Create(request *model.CreateRequest) *model.Deck {
	deck := model.NewDeck(
		uuid.New(),
		s.CardsRepository.Get(request.Codes, request.Values, request.Suits),
		request.Shuffle,
		request.Amount)

	doShuffle(deck)
	applyAmount(deck)

	return s.DecksRepository.Set(deck)
}

func (s *ShufflerService) Show(request *model.ShowRequest) *model.Deck {
	return s.DecksRepository.Get(request.DeckId)
}

func doShuffle(deck *model.Deck) {
	if deck.Shuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(deck.Cards), func(i int, j int) {
			deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
		})
	}
}

func applyAmount(deck *model.Deck) {
	if deck.Amount == 0 {
		deck.Cards = []model.Card{}
		deck.Remaining = 0
	}
	if deck.Amount > 1 {
		cards := deck.Cards
		for i := 1; i < deck.Amount; i++ {
			deck.Cards = append(deck.Cards, cards...)
		}
		deck.Remaining = len(deck.Cards)
	}
}
