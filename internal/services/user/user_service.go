package user

import userRepository "wallet-tracker-backend/internal/repositories/user"

type UserService interface {
	GetUser() string
}

type userService struct {
	userRepository userRepository.UserRepository
}

type UserServiceConfig struct {
	UserRepository userRepository.UserRepository
}

func NewUserService(usc *UserServiceConfig) UserService {
	return &userService{
		userRepository: usc.UserRepository,
	}
}
