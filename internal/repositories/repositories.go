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

// func New(hc *RepositoriesConfig) *Repositories {
// 	return &Repositories{
// 		UserRepository: hc.UserRepository,
// 	}
// }

func InitRepositories() *Repositories {
	userRepository := ur.NewUserRepository(
		&ur.UserRepositoryConfig{
			DB: &gorm.DB{},
		},
	)

	// r := New(&RepositoriesConfig{
	// 	UserRepository: userRepository,
	// })

	r := &Repositories{
		UserRepository: userRepository,
	}

	return r
}
