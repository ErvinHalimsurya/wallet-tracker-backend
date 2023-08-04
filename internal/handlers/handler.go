package handlers

import (
	uh "wallet-tracker-backend/internal/handlers/user"
	r "wallet-tracker-backend/internal/repositories"
	s "wallet-tracker-backend/internal/services"
)

type Handlers struct {
	UserHandler uh.UserHandler
}

func InitHandlers() *Handlers {
	repositories := r.InitRepositories()
	services := s.InitServices(repositories)

	userHandler := uh.NewUserHandler(
		&uh.UserHandlerConfig{
			UserService: services.UserService,
		},
	)

	h := &Handlers{
		UserHandler: userHandler,
	}

	return h
}
