package auth

import (
	"errors"
	"gorm.io/gorm"
	"sakshyahere/tuko/internal/model"
)

type authRepository struct {
	db *gorm.DB
}

func (r *authRepository) CreateUser(email string, password string, firstName string, lastName string) (*model.User, error) {
	user := &model.User{
		Email:     email,
		Password:  password,
		FirstName: firstName,
		LastName:  lastName,
	}

	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *authRepository) EmailExists(email string) error {
	var user model.User
	err := r.db.First(&user, "email = ?", email).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	if err != nil {
		return err // could be gorm.ErrRecordNotFound or a real DB error
	}
	return errors.New("email already exists") // email is already taken
}

func (r *authRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err // could be gorm.ErrRecordNotFound or a real DB error
	}
	return &user, nil // email is already taken
}
