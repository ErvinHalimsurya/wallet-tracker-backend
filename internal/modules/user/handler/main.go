package handler

import s "wallet-tracker-backend/internal/modules/user/service"

type UserHandler interface {
	GetUserHandler()
}

type userHandler struct {
	userService s.UserService
}

type UserHandlerConfig struct {
	UserService s.UserService
}

func NewuserHandler(uhc *UserHandlerConfig) UserHandler {
	return &userHandler{
		userService: uhc.UserService,
	}
}
