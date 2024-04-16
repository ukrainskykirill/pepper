package services

import (
	vd "github.com/go-playground/validator/v10"
	"github.com/ukrainskykirill/pepper/pkg/repositories"
)

type IServices interface {
    GetUsersService() IUsersService
}

type Services struct {
	Users IUsersService
	Validator *vd.Validate
}

func NewServices(repo *repositories.Repositories, validator *vd.Validate) IServices {
	return &Services{
		Users: NewUserService(repo.Users, validator),
	}
}

func (s *Services) GetUsersService() IUsersService {
    return s.Users
}

