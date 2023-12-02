package routes

import (
	"github.com/elsierra/go-echo-zik/handlers"
	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	e.GET("/", handlers.HomeHandler)
}
