package user_service

import user_repository "wallet-tracker-backend/internal/repositories/user"

type UserService interface {
	GetUser() string
}

type userService struct {
	userRepository user_repository.UserRepository
}

type UserServiceConfig struct {
	UserRepository user_repository.UserRepository
}

func NewUserService(usc *UserServiceConfig) UserService {
	return &userService{
		userRepository: usc.UserRepository,
	}
}
