package api

import (
	"net/http"

	"the-coolest-shuffler/internal/model"
	"the-coolest-shuffler/internal/request"

	"github.com/labstack/echo/v4"
)

type Shuffler interface {
	Create(deck *request.Deck) *model.Deck
	Show(deck *request.Deck) *model.Deck
}

type Decks struct {
	Shuffler Shuffler
}

func NewDecks(shuffler Shuffler) *Decks {
	return &Decks{
		Shuffler: shuffler,
	}
}

func (d Decks) Register(server *echo.Echo) {
	v1 := server.Group("the-coolest-shuffler/v1")
	v1.POST("/deck", d.Create)
	v1.GET("/deck/new", d.New)
	v1.GET("/deck/:id", d.Show)
}

func (d Decks) Create(c echo.Context) error {
	deck := new(request.Deck)
	if err := c.Bind(deck); err != nil {
		return badRequest(c, err)
	}

	return ok(c, d.Shuffler.Create(deck))
}

func (d Decks) New(c echo.Context) error {
	deck := new(request.Deck)
	if err := c.Bind(deck); err != nil {
		return badRequest(c, err)
	}

	return ok(c, d.Shuffler.Create(deck))
}

func (d Decks) Show(c echo.Context) error {
	deck := new(request.Deck)
	if err := c.Bind(deck); err != nil {
		return badRequest(c, err)
	}

	return ok(c, d.Shuffler.Show(deck))
}

func ok(context echo.Context, object interface{}) error {
	return context.JSON(http.StatusOK, object)
}

func badRequest(context echo.Context, object error) error {
	return context.JSON(http.StatusBadRequest, object)
}
