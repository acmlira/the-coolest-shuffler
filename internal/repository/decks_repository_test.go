package repository

import (
	"context"
	"testing"
	"the-coolest-shuffler/internal/model"
	"time"

	redisMock "github.com/go-redis/redismock/v8"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type Suite struct {
	key      uuid.UUID
	value    interface{}
	expected *model.Deck
	name     string
}

func TesDeckRepository(t *testing.T) {
	scenario := Suite{
		key: uuid.MustParse("fae9ebce-d23c-484d-95a9-f92a67ef44bb"),
		value: "{\"id\":\"fae9ebce-d23c-484d-95a9-f92a67ef44bb\",\"remaining\":1,\"shuffled\":false,\"cards\":[{\"value\":\"ACE\", \"suit\":\"CLUBS\", \"code\":\"AC\"}],\"amount\":1}",
		expected: &model.Deck{
			Id:        uuid.MustParse("fae9ebce-d23c-484d-95a9-f92a67ef44bb"),
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
		},
		name: "shouldGetADeck",
	}

	t.Run(scenario.name, func(t *testing.T) {
		redis, mock := redisMock.NewClientMock()
		mock.ExpectGet(scenario.key.String()).SetVal(scenario.value.(string))
		decksRepository := DecksRepository{
			Redis: redis, 
			Context: context.Background(),
		}
		deck := decksRepository.Get(scenario.key)
		assert.Equal(t, scenario.expected, deck)
	})

	t.Run(scenario.name, func(t *testing.T) {
		redis, mock := redisMock.NewClientMock()
		mock.ExpectSet(scenario.key.String(), scenario.value, time.Millisecond).SetErr(nil)
		decksRepository := DecksRepository{
			Redis: redis, 
			Context: context.Background(),
		}
		deck := decksRepository.Set(scenario.expected)
		assert.Equal(t, scenario.expected, deck)
	})
}