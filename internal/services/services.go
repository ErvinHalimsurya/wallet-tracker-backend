package services

import (
	"wallet-tracker-backend/internal/repositories"
	us "wallet-tracker-backend/internal/services/user"
)

type Services struct {
	UserService us.UserService
}

// type ServicesConfig struct {
// 	UserService us.UserService
// }

// func New(hc *ServicesConfig) *Services {
// 	return &Services{
// 		UserService: hc.UserService,
// 	}
// }

func InitServices(r *repositories.Repositories) *Services {

	userService := us.NewUserService(
		&us.UserServiceConfig{
			UserRepository: r.UserRepository,
		},
	)

	// s := New(&ServicesConfig{
	// 	UserService: userService,
	// })

	s := &Services{
		UserService: userService,
	}

	return s
}
