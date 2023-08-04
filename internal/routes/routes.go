package routes

import (
	"net/http"
	"wallet-tracker-backend/internal/handlers"

	"github.com/labstack/echo/v4"
)

func InitiateRoutes(e *echo.Echo, h *handlers.Handlers) {

	e.GET("/", hello)
	e.GET("/ecample/:id", example)
	e.GET("/user", h.UserHandler.GetUser)
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello")
}

// should be handler
func example(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
