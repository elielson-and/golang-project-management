// cmd/api/main.go
package main

import (
	_ "github.com/elielson-and/golang-project-management/docs" // Import necessário para o Swag
	"github.com/elielson-and/golang-project-management/internal/routes"
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
func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
