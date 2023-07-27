package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitiateRoutes(e *echo.Echo) {

	e.GET("/", hello)
	e.GET("/ecample/:id", example)

}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello")
}

// should be handler
func example(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
