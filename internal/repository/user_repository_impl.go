package repository

type userRepo struct {
}

func (repo *userRepo) GetUser() string {
	return "Hello World"
}
