package user

import (
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUser() string
}

type userRepository struct {
	db *gorm.DB
}

type UserRepositoryConfig struct {
	DB *gorm.DB
}

func NewUserRepository(urc *UserRepositoryConfig) UserRepository {
	return &userRepository{
		db: urc.DB,
	}
}
