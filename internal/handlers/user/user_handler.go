package user_handler

import (
	user_service "wallet-tracker-backend/internal/services/user"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	GetUser(c echo.Context) error
}

type userHandler struct {
	userService user_service.UserService
}

type UserHandlerConfig struct {
	UserService user_service.UserService
}

func NewUserHandler(uhc *UserHandlerConfig) UserHandler {
	return &userHandler{
		userService: uhc.UserService,
	}
}
