//go:build all
// +build all

package service

import (
	"context"
	"io/ioutil"
	"testing"

	"the-coolest-shuffler/configs"
	"the-coolest-shuffler/internal/model"
	"the-coolest-shuffler/internal/repository"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestDeck(t *testing.T) {
	type scenario struct {
		shuffle  bool
		amount   int
		codes    []string
		values   []string
		suits    []string
		expected *model.Deck
		name     string
	}

	it := scenario{
		false,
		1,
		[]string{"AC"},
		[]string{"ACE"},
		[]string{"CLUBS"},
		model.NewDeck(uuid.New(), []model.Card{{Value: "ACE", Suit: "CLUBS", Code: "AC"}}, 1, false, 1),
		"integration",
	}

	log.SetOutput(ioutil.Discard)

	configs.Init()
	repository.Database(configs.GetPostgresDSN())
	repository.Cache(
		configs.GetRedisHost(),
		configs.GetRedisPort(),
		configs.GetRedisDatabase(),
		configs.GetRedisPassword())

	t.Run(it.name, func(t *testing.T) {
		var deckService = NewShuffler()
		var newDeck = deckService.CreateNewDeck(
			context.Background(),
			it.shuffle,
			it.amount,
			it.codes,
			it.values,
			it.suits)

		assert.Equal(t, newDeck.Cards, it.expected.Cards)

		var openedDeck = deckService.OpenDeck(context.Background(), newDeck.Id)

		assert.Equal(t, openedDeck, newDeck)
	})
}
