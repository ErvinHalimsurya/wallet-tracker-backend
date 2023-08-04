package user

import (
	userService "wallet-tracker-backend/internal/services/user"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	GetUser(c echo.Context) error
}

type userHandler struct {
	userService userService.UserService
}

type UserHandlerConfig struct {
	UserService userService.UserService
}

func NewUserHandler(uhc *UserHandlerConfig) UserHandler {
	return &userHandler{
		userService: uhc.UserService,
	}
}
