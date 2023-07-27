package main

import (
	"wallet-tracker-backend/config"
	"wallet-tracker-backend/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	config.Init()

	e := echo.New()

	routes.InitiateRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
