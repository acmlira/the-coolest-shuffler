package api

import (
	"net/http"

	"the-coolest-shuffler/internal/model"

	"github.com/labstack/echo/v4"

	swaggo "github.com/swaggo/echo-swagger"
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
	v1.GET("/deck/:deckId", d.Show)

	server.GET("/swagger/*", swaggo.WrapHandler)
}

// Deck API
// @Summary 	Create new Deck
// @Description Create new Deck based in predefined cards
// @Tags 		Deck
// @Accept		json
// @Produce 	json
// @Param		Request body model.Request false "Deck properties"
// @Success		200	{object} model.Deck
// @Failure		400	{object} string
// @Failure		500 {object} string
// @Router      /deck [post]
func (d DeckAPI) Create(c echo.Context) error {
	request := new(model.Request)
	if err := c.Bind(request); err != nil {
		return badRequest(c, err)
	}

	return ok(c, d.ShufflerService.Create(request))
}

// Deck API
// @Summary 	Create new Deck
// @Description Create new Deck based in predefined cards
// @Tags 		Deck
// @Accept		json
// @Produce 	json
// @Param		shuffle query bool     false "shuffle cards"
// @Param		amount  query int      false "amount of card sets (before filters)"
// @Param		codes   query []string false "code filter"
// @Param		values  query []string false "value filter"
// @Param		suits   query []string false "suit filter"
// @Success		200	{object} model.Deck
// @Failure		400	{object} string
// @Failure		500 {object} string
// @Router      /deck/new [get]
func (d DeckAPI) New(c echo.Context) error {
	request := new(model.Request)
	if err := c.Bind(request); err != nil {
		return badRequest(c, err)
	}

	return ok(c, d.ShufflerService.Create(request))
}

// Deck API
// @Summary 	Show a deck
// @Description Show a created deck
// @Tags 		Deck
// @Accept		json
// @Produce 	json
// @Param		deckId path string false "code filter"
// @Success		200	{object} model.Deck
// @Failure		400	{object} string
// @Failure		500 {object} string
// @Router      /deck/{deckId} [get]
func (d DeckAPI) Show(c echo.Context) error {
	request := new(model.Request)
	if err := c.Bind(request); err != nil {
		return badRequest(c, err)
	}

	return ok(c, d.ShufflerService.Show(request))
}

func ok(context echo.Context, object interface{}) error {
	return context.JSON(http.StatusOK, object)
}

func badRequest(context echo.Context, object error) error {
	return context.JSON(http.StatusBadRequest, object)
}
