package api

import (
	"net/http"

	"the-coolest-shuffler/internal/model"

	"github.com/labstack/echo/v4"
)

type ShufflerService interface {
	Create(deck *model.Request) *model.Deck
	Show(deck *model.Request) *model.Deck
}

type DeckAPI struct {
	ShufflerService ShufflerService
}

func NewDeckAPI(shufflerService ShufflerService) *DeckAPI {
	return &DeckAPI{
		ShufflerService: shufflerService,
	}
}

func (d DeckAPI) Register(server *echo.Echo) {
	v1 := server.Group("the-coolest-shuffler/v1")
	v1.POST("/deck", d.Create)
	v1.GET("/deck/new", d.New)
	v1.GET("/deck/:id", d.Show)
}

func (d DeckAPI) Create(c echo.Context) error {
	deck := new(model.Request)
	if err := c.Bind(deck); err != nil {
		return badRequest(c, err)
	}

	return ok(c, d.ShufflerService.Create(deck))
}

func (d DeckAPI) New(c echo.Context) error {
	deck := new(model.Request)
	if err := c.Bind(deck); err != nil {
		return badRequest(c, err)
	}

	return ok(c, d.ShufflerService.Create(deck))
}

func (d DeckAPI) Show(c echo.Context) error {
	deck := new(model.Request)
	if err := c.Bind(deck); err != nil {
		return badRequest(c, err)
	}

	return ok(c, d.ShufflerService.Show(deck))
}

func ok(context echo.Context, object interface{}) error {
	return context.JSON(http.StatusOK, object)
}

func badRequest(context echo.Context, object error) error {
	return context.JSON(http.StatusBadRequest, object)
}
