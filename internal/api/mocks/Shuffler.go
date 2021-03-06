// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	filter "the-coolest-shuffler/internal/filter"

	mock "github.com/stretchr/testify/mock"

	model "the-coolest-shuffler/internal/model"

	uuid "github.com/google/uuid"
)

// Shuffler is an autogenerated mock type for the Shuffler type
type Shuffler struct {
	mock.Mock
}

// CreateNewDeck provides a mock function with given fields: shuffle, amount, cardFilter
func (_m *Shuffler) CreateNewDeck(shuffle bool, amount int, cardFilter *filter.CardFilter) *model.Deck {
	ret := _m.Called(shuffle, amount, cardFilter)

	var r0 *model.Deck
	if rf, ok := ret.Get(0).(func(bool, int, *filter.CardFilter) *model.Deck); ok {
		r0 = rf(shuffle, amount, cardFilter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Deck)
		}
	}

	return r0
}

// DrawCard provides a mock function with given fields: id, count
func (_m *Shuffler) DrawCard(id uuid.UUID, count int) *model.Draw {
	ret := _m.Called(id, count)

	var r0 *model.Draw
	if rf, ok := ret.Get(0).(func(uuid.UUID, int) *model.Draw); ok {
		r0 = rf(id, count)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Draw)
		}
	}

	return r0
}

// OpenDeck provides a mock function with given fields: id
func (_m *Shuffler) OpenDeck(id uuid.UUID) *model.Deck {
	ret := _m.Called(id)

	var r0 *model.Deck
	if rf, ok := ret.Get(0).(func(uuid.UUID) *model.Deck); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Deck)
		}
	}

	return r0
}

type mockConstructorTestingTNewShuffler interface {
	mock.TestingT
	Cleanup(func())
}

// NewShuffler creates a new instance of Shuffler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewShuffler(t mockConstructorTestingTNewShuffler) *Shuffler {
	mock := &Shuffler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
