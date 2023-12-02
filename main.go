package main

import (
	"github.com/elsierra/go-echo-zik/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.Init(e)
	e.Start(":8080")
}
