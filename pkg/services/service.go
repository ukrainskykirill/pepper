package services

import (
	vd "github.com/go-playground/validator/v10"
	"github.com/ukrainskykirill/pepper/pkg/repositories"
)

type Services struct {
	Users *UsersService
	Validator *vd.Validate
}

func NewServices(repo *repositories.Repositories, validator *vd.Validate) *Services {
	return &Services{
		Users: NewUserService(repo.Users, validator),
	}
}
