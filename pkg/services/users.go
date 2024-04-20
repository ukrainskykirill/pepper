package services

import (
	"context"
	"fmt"

	vd "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/ukrainskykirill/pepper/pkg/types"
	"github.com/ukrainskykirill/pepper/pkg/database"
	"github.com/ukrainskykirill/pepper/pkg/repositories"
)

type IUsersService interface {
	CreateUser(ctx context.Context, input *types.UserInput) error
	GetUser(ctx context.Context, id uuid.UUID) (database.GetUserRow, error)
	UpdateUser(ctx context.Context, input *types.UserInputUpd) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type UsersService struct {
	repo *repositories.UsersRepository
	validator *vd.Validate
}

func NewUserService(repo *repositories.UsersRepository, validator *vd.Validate) IUsersService{
	return &UsersService{
		repo: repo,
		validator: validator,
	}
}

func (service *UsersService) CreateUser(ctx context.Context, input *types.UserInput) error {
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
	if err := service.repo.CreateUsers(ctx, input); err != nil {
		return err
	}
	return nil
}

func (service *UsersService) DeleteUser(ctx context.Context, id uuid.UUID) error {
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
func (service *UsersService) GetUser(ctx context.Context, id uuid.UUID) (database.GetUserRow, error){
	isExists, err := service.repo.IsExistsById(ctx, id)
	if err != nil {
		return database.GetUserRow{}, err
	} else if !isExists {
		return database.GetUserRow{}, UserDoesntExists{
			fmt.Errorf("user exists by id %s", id),
		}
	}
	user, err := service.repo.GetUser(ctx, id)
	return user, err
}

func (service *UsersService) UpdateUser(ctx context.Context, input *types.UserInputUpd) error {
	isExists, err := service.repo.IsExistsById(ctx, input.ID)
	if err != nil {
		return err
	} else if !isExists {
		return UserDoesntExists{
			fmt.Errorf("user exists by id %s", input.ID),
		}
	}
	err = service.repo.UpdateUser(ctx, input)
	return err
}