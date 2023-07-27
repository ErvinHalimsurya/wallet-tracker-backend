package service

import r "wallet-tracker-backend/internal/modules/user/repository"

type UserService interface {
	GetUser() string
}

type userService struct {
	userRepository r.UserRepository
}

type UserServiceConfig struct {
	UserRepository r.UserRepository
}

func NewUserService(usc *UserServiceConfig) UserService {
	return &userService{
		userRepository: usc.UserRepository,
	}
}
