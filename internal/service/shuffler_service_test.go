package service

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"the-coolest-shuffler/internal/model"
	"the-coolest-shuffler/internal/service/mocks"
)

type Scenario struct {
	request  interface{}
	expected *model.Deck
	name     string
}

var deckId = uuid.New()

var response = &model.Deck{
	Id:        deckId,
	Remaining: 1,
	Shuffle:   false,
	Amount:    1,
	Cards: []model.Card{
		{
			Value: "ACE",
			Suit:  "CLUBS",
			Code:  "AC",
		},
	},
}

var createRequest = &model.CreateRequest{
	Shuffle: false,
	Amount:  1,
	Codes:   []string{"AC"},
	Values:  []string{"ACE"},
	Suits:   []string{"CLUBS"},
}

var showRequest = &model.ShowRequest{
	DeckId: deckId,
}

func TestCreateADeck(t *testing.T) {
	scenario := Scenario{
		createRequest,
		response,
		"shouldCreateADeck",
	}

	t.Run(scenario.name, func(t *testing.T) {
		var cardsRepository = mocks.NewCardsRepository(t)
		var decksRepository = mocks.NewDecksRepository(t)

		cardsRepository.On("Get", mock.Anything, mock.Anything, mock.Anything).Return(scenario.expected.Cards)
		decksRepository.On("Set", mock.Anything).Return(scenario.expected)

		var shufflerService = NewShufflerService(cardsRepository, decksRepository)
		var deck = shufflerService.Create(scenario.request.(*model.CreateRequest))

		assert.Equal(t, deck.Remaining, scenario.expected.Remaining)
		assert.Equal(t, deck.Shuffle, scenario.expected.Shuffle)
		assert.Equal(t, deck.Amount, scenario.expected.Amount)
		assert.Equal(t, deck.Cards, scenario.expected.Cards)
	})
}

func TestShowADeck(t *testing.T) {
	scenario := Scenario{
		showRequest,
		response,
		"shouldShowADeck",
	}

	t.Run(scenario.name, func(t *testing.T) {
		var cardsRepository = mocks.NewCardsRepository(t)
		var decksRepository = mocks.NewDecksRepository(t)

		decksRepository.On("Get", mock.Anything).Return(scenario.expected)

		var shufflerService = NewShufflerService(cardsRepository, decksRepository)
		var deck = shufflerService.Show(scenario.request.(*model.ShowRequest))

		assert.Equal(t, deck.Remaining, scenario.expected.Remaining)
		assert.Equal(t, deck.Shuffle, scenario.expected.Shuffle)
		assert.Equal(t, deck.Amount, scenario.expected.Amount)
		assert.Equal(t, deck.Cards, scenario.expected.Cards)
	})
}
