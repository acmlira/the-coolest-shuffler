package api

import (
	"context"
	"net/http"

	"the-coolest-shuffler/internal/model"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type DeckManager interface {
	CreateNewDeck(context context.Context, shuffle bool, amount int, codes []string, values []string, suits []string) model.Deck
	OpenDeck(context context.Context, id uuid.UUID) model.Deck
	DrawCard(context context.Context, id uuid.UUID, count int) model.Draw
}

type Deck struct {
	Manager DeckManager
}

func (d Deck) Register(server *echo.Echo) {
	v1 := server.Group("the-coolest-shuffler/v1")
	v1.GET("/deck/new", d.CreateNewDeck)
	v1.GET("/deck/:id", d.OpenDeck)
	v1.GET("/deck/:id/draw", d.DrawCard)
}

func (d Deck) CreateNewDeck(c echo.Context) error {
	shuffle := OptionalBool(c.QueryParam("shuffle"), "?shuffle=", false)
	amount := OptionalInt(c.QueryParam("amount"), "?amount=", 1)
	codes := OptionalStringList(c.QueryParam("codes"), "?codes=")
	values := OptionalStringList(c.QueryParam("values"), "?values=")
	suits := OptionalStringList(c.QueryParam("suits"), "?suits=")
	return c.JSON(http.StatusOK, d.Manager.CreateNewDeck(c.Request().Context(), shuffle, amount, codes, values, suits))
}

func (d Deck) OpenDeck(c echo.Context) error {
	id, err := RequiredUUID(c.Param("id"), ":id")
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, d.Manager.OpenDeck(c.Request().Context(), id))
}

func (d Deck) DrawCard(c echo.Context) error {
	id, err := RequiredUUID(c.Param("id"), ":id")
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	count := OptionalInt(c.QueryParam("count"), "?count=", 1)
	return c.JSON(http.StatusOK, d.Manager.DrawCard(c.Request().Context(), id, count))
}
