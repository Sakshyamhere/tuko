package repository

type UserRepository interface {
	GetUser() string
}

func NewUserRepository() UserRepository {
	return &userRepo{}
}
