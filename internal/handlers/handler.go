package handlers

import (
	uh "wallet-tracker-backend/internal/handlers/user"
	"wallet-tracker-backend/internal/repositories"
	"wallet-tracker-backend/internal/services"
)

type Handlers struct {
	UserHandler uh.UserHandler
}

// type HandlersConfig struct {
// 	UserHandler uh.UserHandler
// }

// func New(hc *HandlersConfig) *Handlers {
// 	return &Handlers{
// 		UserHandler: hc.UserHandler,
// 	}
// }

func InitHandlers() *Handlers {
	repositories := repositories.InitRepositories()
	services := services.InitServices(repositories)

	userHandler := uh.NewUserHandler(
		&uh.UserHandlerConfig{
			UserService: services.UserService,
		},
	)

	// h := New(&HandlersConfig{
	// 	UserHandler: userHandler,
	// })

	h := &Handlers{
		UserHandler: userHandler,
	}

	return h
}
