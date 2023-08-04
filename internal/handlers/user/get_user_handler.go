package user_handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (uh *userHandler) GetUser(c echo.Context) error {

	return c.String(http.StatusOK, uh.userService.GetUser())
}
