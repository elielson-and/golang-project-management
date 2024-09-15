package main

import (
	"net/http"

	_ "github.com/elielson-and/golang-project-management/docs" // Import necessário para o Swag
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server for Swagger.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

// PingHandler responde com "pong"
// @Summary Ping the server
// @Description This is a simple ping-pong endpoint to test server responsiveness
// @Tags ping
// @Success 200 {string} string "pong"
// @Router /ping [get]
func PingHandler(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

// PingHandler2 responde com "pong"
// @Summary Ping the server
// @Description This is a simple ping-pong endpoint to test server responsiveness
// @Tags ping
// @Success 200 {string} string "pong"
// @Router /ping [get]
func PingHandler2(c echo.Context) error {
	return c.String(http.StatusOK, "pong2")
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/ping", PingHandler)
	e.GET("/ping2", PingHandler2)

	e.Logger.Fatal(e.Start(":8080"))
}
