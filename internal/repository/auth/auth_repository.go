package auth

import (
	"gorm.io/gorm"
	"sakshyahere/tuko/internal/model"
)

type AuthRepository interface {
	CreateUser(email string, password string, firstName string, lastName string) (*model.User, error)
	EmailExists(email string) error
	GetUserByEmail(email string) (*model.User, error)
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}
