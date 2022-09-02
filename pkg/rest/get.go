package rest

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Get(c echo.Context) error {
	nick := strings.TrimSpace(c.QueryParam("nick"))
	if nick == "" {
		return c.JSONPretty(http.StatusBadRequest, echo.Map{
			"error":         "wrong request",
			"how to use it": "make a `GET /?nick=TikTokNickName`",
		}, "  ")
	}

	return nil
}
