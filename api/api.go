package api

import (
	"log"
	"net/http"

	fmbot "github.com/amaghzaz-y/fm-bot/internal"
	"github.com/labstack/echo/v4"
)

type API struct {
	Bot *fmbot.Bot
}

func New() *API {
	bot := fmbot.New()
	bot.Init()
	return &API{
		Bot: bot,
	}
}

// GET /chat/:query
func (a *API) replyHandler(c echo.Context) error {
	query := c.Param("query")
	if len(query) < 2 {
		return c.String(http.StatusBadRequest, "bad query !!")
	}
	log.Println(query)
	reply, err := a.Bot.Reply(query)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.String(200, reply.Reply)
}

func (a *API) Start() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "FARMING MARS BOT API v0.1 (alpha)")
	})
	e.GET("/chat/:query", a.replyHandler)
	e.Logger.Fatal(e.Start("127.0.0.1:1323"))
}
