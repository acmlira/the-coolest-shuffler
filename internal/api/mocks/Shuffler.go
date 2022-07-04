// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"
	model "the-coolest-shuffler/internal/model"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// Shuffler is an autogenerated mock type for the Shuffler type
type Shuffler struct {
	mock.Mock
}

// CreateNewDeck provides a mock function with given fields: _a0, shuffle, amount, codes, values, suits
func (_m *Shuffler) CreateNewDeck(_a0 context.Context, shuffle bool, amount int, codes []string, values []string, suits []string) *model.Deck {
	ret := _m.Called(_a0, shuffle, amount, codes, values, suits)

	var r0 *model.Deck
	if rf, ok := ret.Get(0).(func(context.Context, bool, int, []string, []string, []string) *model.Deck); ok {
		r0 = rf(_a0, shuffle, amount, codes, values, suits)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Deck)
		}
	}

	return r0
}

// DrawCard provides a mock function with given fields: _a0, id, count
func (_m *Shuffler) DrawCard(_a0 context.Context, id uuid.UUID, count int) *model.Draw {
	ret := _m.Called(_a0, id, count)

	var r0 *model.Draw
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, int) *model.Draw); ok {
		r0 = rf(_a0, id, count)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Draw)
		}
	}

	return r0
}

// OpenDeck provides a mock function with given fields: _a0, id
func (_m *Shuffler) OpenDeck(_a0 context.Context, id uuid.UUID) *model.Deck {
	ret := _m.Called(_a0, id)

	var r0 *model.Deck
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *model.Deck); ok {
		r0 = rf(_a0, id)
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
