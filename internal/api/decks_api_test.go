package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"the-coolest-shuffler/internal/api/mocks"
	"the-coolest-shuffler/internal/model"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeck(t *testing.T) {
	id := uuid.New()
	service := &mocks.Shuffler{}
	service.On("Create", mock.Anything).Return(&model.Deck{})
	service.On("Show", mock.Anything).Return(&model.Deck{})

	t.Run(`CreateNewDeck`, func(t *testing.T) {
		tested := NewDeckAPI(service)
		e := echo.New()
		request := httptest.NewRequest(http.MethodGet, "/the-coolest-shuffler/v1", nil)
		recorder := httptest.NewRecorder()
		context := e.NewContext(request, recorder)
		context.SetPath("/deck/new")
		if assert.NoError(t, tested.Create(context)) {
			assert.Equal(t, http.StatusOK, recorder.Code)
		}
	})

	t.Run(`OpenDeck`, func(t *testing.T) {
		tested := DeckAPI{service}
		e := echo.New()
		request := httptest.NewRequest(http.MethodGet, "/the-coolest-shuffler/v1", nil)
		recorder := httptest.NewRecorder()
		context := e.NewContext(request, recorder)
		context.SetPath("/deck/%s")
		context.SetParamNames("id")
		context.SetParamValues(id.String())

		if assert.NoError(t, tested.Show(context)) {
			assert.Equal(t, http.StatusOK, recorder.Code)
		}
	})
}
