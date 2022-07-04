//go:build all
// +build all

package dao

import (
	"testing"
	"the-coolest-shuffler/configs"
	"the-coolest-shuffler/internal/model"
	"the-coolest-shuffler/internal/repository"

	"github.com/stretchr/testify/assert"
)

func TestSelect(t *testing.T) {
	type scenario struct {
		codes    []string
		values   []string
		suits    []string
		expected interface{}
		name     string
	}

	scenarios := []scenario{
		{[]string{"AC"}, []string{"ACE"}, []string{"CLUBS"}, []model.Card{{Value: "ACE", Suit: "CLUBS", Code: "AC"}}, "full filled ok"},
	}

	for _, it := range scenarios {
		t.Run(it.name, func(t *testing.T) {
			configs.Init()
			repository.Database(configs.GetPostgresDSN())
			value := NewCard().Select(it.codes, it.values, it.suits)
			assert.Equal(t, it.expected, value)
		})
	}
}
