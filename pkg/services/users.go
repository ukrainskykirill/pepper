package services

import (
	"context"

	vd "github.com/go-playground/validator/v10"
	"github.com/ukrainskykirill/pepper/pkg/repositories"
	"github.com/ukrainskykirill/pepper/pkg/types"
)

type UsersService struct {
	repo *repositories.UsersRepository
	validator *vd.Validate
}

func NewUserService(repo *repositories.UsersRepository, validator *vd.Validate) *UsersService{
	return &UsersService{
		repo: repo,
		validator: validator,
	}
}

func (service *UsersService) CreateUser(input *types.UserInput) error {
	ctx := context.Background()
	if err := service.validator.Struct(input); err != nil {
		return err
	}
	service.repo.CreateUser(ctx, input)
	return nil
}

func (service *UsersService) DeleteUser() {
	return 
}
func (service *UsersService) GetUser() {
	return 
}
func (service *UsersService) UpdateUser() {
	return 
}