package services

import (
	"context"
	"fmt"

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
		fmt.Println("validator error")
		return InvalidInputData{
			fmt.Errorf("invalid input data %w", err),
		}
	}
	isExists, err := service.repo.IsExistsByLogin(ctx, input.Login)
	if err != nil {
		return err
	}
	if isExists {
		return AlreadyExistsByLogin{
			fmt.Errorf("already exists by login %s", input.Login),
		}
	}
	if err := service.repo.CreateUser(ctx, input); err != nil {
		return err
	}
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