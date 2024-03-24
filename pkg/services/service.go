package services

import "github.com/ukrainskykirill/pepper/pkg/repositories"

type Services struct {
	Users *UsersService
}

func NewServices(repo *repositories.Repositories) *Services {
	return &Services{
		Users: NewUserService(repo.Users),
	}
}
