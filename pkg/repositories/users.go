package repositories

import (
	"context"
	"fmt"

	"github.com/ukrainskykirill/pepper/pkg/database"
	"github.com/ukrainskykirill/pepper/pkg/types"
)

type UsersRepository struct {
	queries *database.Queries
}

func NewUserRepository(queries *database.Queries) *UsersRepository{
	return &UsersRepository{
		queries: queries,
	}
}

func (repo *UsersRepository) CreateUser(ctx context.Context, input *types.UserInput) error {
	err := repo.queries.CreateUser(ctx, database.CreateUserParams{
		Name: input.Name,
		Login: input.Login,
		Phone: input.Phone,
	})
	if err != nil {
		return fmt.Errorf("Error with creating user %w", err)
	}
	return nil
}
func (repo *UsersRepository) DeleteUser() {
	return
}
func (repo *UsersRepository) GetUser() {
	return 
}
func (repo *UsersRepository) UpdateUser() {
	return 
}
func (repo *UsersRepository) IsExistsByLogin(ctx context.Context, login string) (bool, error) {
	isExists, err := repo.queries.IsExistsByLogin(ctx, login)
	if err != nil {
		return isExists, fmt.Errorf("failed to do db request IsExistsById")
	}
	return isExists, err
}