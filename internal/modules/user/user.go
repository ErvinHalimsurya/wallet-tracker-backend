package user

import (
	"wallet-tracker-backend/internal/modules/account"
	h "wallet-tracker-backend/internal/modules/user/handler"
	r "wallet-tracker-backend/internal/modules/user/repository"
	s "wallet-tracker-backend/internal/modules/user/service"

	"gorm.io/gorm"
)

type UserModule struct {
	UserHandler    h.UserHandler
	UserSerivce    s.UserService
	UserRepository r.UserRepository
}

type UserModuleConfig struct {
	userHandler    h.UserHandler
	userService    s.UserService
	userRepository r.UserRepository
}

func NewUserModule(umc *UserModuleConfig) *UserModule {
	return &UserModule{
		UserHandler:    umc.userHandler,
		UserSerivce:    umc.userService,
		UserRepository: umc.userRepository,
	}
}

func InitUserModule(db *gorm.DB) *UserModule {

	account.InitAccountModule()

	userRepository := r.NewUserRepository(&r.UserRepositoryConfig{
		DB: db,
	})

	userService := s.NewUserService(
		&s.UserServiceConfig{
			UserRepository: userRepository,
		},
	)

	userHandler := h.NewuserHandler(
		&h.UserHandlerConfig{
			UserService: userService,
		},
	)

	userModule := NewUserModule(&UserModuleConfig{
		userHandler:    userHandler,
		userService:    userService,
		userRepository: userRepository,
	})

	return userModule

}
