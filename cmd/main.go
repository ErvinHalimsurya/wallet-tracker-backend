package main

import (
	"wallet-tracker-backend/config"
	"wallet-tracker-backend/internal/handlers"
	"wallet-tracker-backend/internal/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	config.Init()

	e := echo.New()
	h := handlers.InitHandlers()

	routes.InitiateRoutes(e, h)

	e.Logger.Fatal(e.Start(":1323"))
}
