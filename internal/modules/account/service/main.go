package service

import (
	"wallet-tracker-backend/internal/modules/user"

	"gorm.io/gorm"
)

type AccountService interface {
	GetUser() string
}

func NewUserService() {
	var db *gorm.DB
	user.InitUserModule(db)

}
