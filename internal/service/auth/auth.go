package auth

import (
	"errors"
	"sakshyahere/tuko/internal/repository/auth"
	"sakshyahere/tuko/internal/util"
)

type AuthService interface {
	LoginService(email string, password string) (string, error)
	SignupService(email string, firstName string, lastName string, password string) (string, error)
	EmailExistsService(email string) error
}

type authService struct {
	repo auth.AuthRepository
}

func NewAuthService(authRepo auth.AuthRepository) AuthService {
	return &authService{
		repo: authRepo,
	}
}

func (s *authService) LoginService(email string, password string) (string, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}
	credible := util.CheckPasswordHash(password, user.Password)
	if !credible {
		return "", errors.New("invalid password")
	}
	token, err := util.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *authService) SignupService(email string, firstName string, lastName string, password string) (string, error) {
	hash, err := util.HashPassword(password)
	if err != nil {
		return "", err
	}
	user, err := s.repo.CreateUser(email, hash, firstName, lastName)
	if err != nil {
		return "", err
	}
	token, err := util.GenerateJWT(user.ID)

	return token, nil
}

func (s *authService) EmailExistsService(email string) error {
	return s.repo.EmailExists(email)
}
