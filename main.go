package main

import (
	"bukumanga-api/controller"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "5000"
	}
	return ":" + port
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", controller.Hello())
	e.GET("/entries", controller.GetEntries())

	// Start server
	e.Logger.Fatal(e.Start(port()))
}
