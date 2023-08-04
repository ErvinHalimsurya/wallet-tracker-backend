package repositories

import (
	ur "wallet-tracker-backend/internal/repositories/user"

	"gorm.io/gorm"
)

type Repositories struct {
	UserRepository ur.UserRepository
}

type RepositoriesConfig struct {
	UserRepository ur.UserRepository
}

func InitRepositories() *Repositories {
	userRepository := ur.NewUserRepository(
		&ur.UserRepositoryConfig{
			DB: &gorm.DB{},
		},
	)

	r := &Repositories{
		UserRepository: userRepository,
	}

	return r
}
