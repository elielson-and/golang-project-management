// internal/routes/routes.go
package routes

import (
	"github.com/elielson-and/golang-project-management/internal/handlers"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/ping", handlers.PingHandler)
	e.GET("/ping2", handlers.PingHandler2)
}
