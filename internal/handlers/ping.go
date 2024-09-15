package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// PingHandler responde com "pong"
// @Summary Ping the server
// @Description This is a simple ping-pong endpoint to test server responsiveness
// @Tags ping
// @Success 200 {string} string "pong"
// @Router /ping [get]
func PingHandler(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

// PingHandler2 responde com "pong2"
// @Summary Ping the server2
// @Description This is a simple ping-pong endpoint to test server responsiveness
// @Tags ping
// @Success 200 {string} string "pong2"
// @Router /ping2 [get]
func PingHandler2(c echo.Context) error {
	return c.String(http.StatusOK, "pong2")
}
