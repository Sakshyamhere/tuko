package service

import "sakshyahere/tuko/internal/repository"

type UserService interface {
	GetUser() string
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) GetUser() string {
	return s.userRepo.GetUser()
}
