// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	model "the-coolest-shuffler/internal/model"

	mock "github.com/stretchr/testify/mock"
)

// ShufflerService is an autogenerated mock type for the ShufflerService type
type ShufflerService struct {
	mock.Mock
}

// Create provides a mock function with given fields: request
func (_m *ShufflerService) Create(request *model.CreateRequest) *model.Deck {
	ret := _m.Called(request)

	var r0 *model.Deck
	if rf, ok := ret.Get(0).(func(*model.CreateRequest) *model.Deck); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Deck)
		}
	}

	return r0
}

// Show provides a mock function with given fields: request
func (_m *ShufflerService) Show(request *model.ShowRequest) *model.Deck {
	ret := _m.Called(request)

	var r0 *model.Deck
	if rf, ok := ret.Get(0).(func(*model.ShowRequest) *model.Deck); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Deck)
		}
	}

	return r0
}

type mockConstructorTestingTNewShufflerService interface {
	mock.TestingT
	Cleanup(func())
}

// NewShufflerService creates a new instance of ShufflerService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewShufflerService(t mockConstructorTestingTNewShufflerService) *ShufflerService {
	mock := &ShufflerService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}