package api

import (
	"net/http"

	"the-coolest-shuffler/internal/filter"
	"the-coolest-shuffler/internal/model"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Shuffler interface {
	CreateNewDeck(shuffle bool, amount int, cardFilter *filter.CardFilter) *model.Deck
	OpenDeck(id uuid.UUID) *model.Deck
	DrawCard(id uuid.UUID, count int) *model.Draw
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
	v1.GET("/deck/new", d.CreateNewDeck)
	v1.GET("/deck/:id", d.OpenDeck)
	v1.GET("/deck/:id/draw", d.DrawCard)
}

func (d Decks) CreateNewDeck(c echo.Context) error {
	shuffle := optionalBool(c.QueryParam("shuffle"), "?shuffle=", false)
	amount := optionalInt(c.QueryParam("amount"), "?amount=", 1)
	cardFilter := filter.NewCardFilter(
		optionalStringList(c.QueryParam("codes"), "?codes="),
		optionalStringList(c.QueryParam("values"), "?values="),
		optionalStringList(c.QueryParam("suits"), "?suits="),
	)
	return ok(c, d.Shuffler.CreateNewDeck(shuffle, amount, cardFilter))
}

func (d Decks) OpenDeck(c echo.Context) error {
	id, e := requiredUUID(c.Param("id"), ":id")
	if e != nil {
		return notFound(c, e)
	}
	return ok(c, d.Shuffler.OpenDeck(id))
}

func (d Decks) DrawCard(c echo.Context) error {
	id, e := requiredUUID(c.Param("id"), ":id")
	if e != nil {
		return notFound(c, e)
	}
	count := optionalInt(c.QueryParam("count"), "?count=", 1)
	return ok(c, d.Shuffler.DrawCard(id, count))
}

func ok(context echo.Context, object interface{}) error {
	return context.JSON(http.StatusOK, object)
}

func notFound(context echo.Context, object error) error {
	return context.JSON(http.StatusNotFound, object)
}
