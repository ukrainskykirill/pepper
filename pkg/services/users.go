package services

import (
	"context"
	"fmt"

	vd "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
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
		return InvalidInputData{
			fmt.Errorf("invalid input data %w", err),
		}
	}
	isExists, err := service.repo.IsExistsByLogin(ctx, input.Login)
	if err != nil {
		return err
	} else if isExists {
		return AlreadyExistsByLogin{
			fmt.Errorf("already exists by login %s", input.Login),
		}
	}
	if err := service.repo.CreateUser(ctx, input); err != nil {
		return err
	}
	return nil
}

func (service *UsersService) DeleteUser(id uuid.UUID) error {
	ctx := context.Background()
	isExists, err := service.repo.IsExistsById(ctx, id)
	if err != nil {
		return err
	} else if !isExists {
		return UserDoesntExists{
			fmt.Errorf("user exists by id %s", id),
		}
	}
	if err := service.repo.DeleteUser(ctx, id); err != nil {
		return err
	}
	return nil
}
func (service *UsersService) GetUser() {
}
func (service *UsersService) UpdateUser() {
}