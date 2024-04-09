package repositories

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/ukrainskykirill/pepper/pkg/database"
)

type UsersRepository struct {
	queries *database.Queries
}

func NewUserRepository(queries *database.Queries) *UsersRepository{
	return &UsersRepository{
		queries: queries,
	}
}

func (repo *UsersRepository) CreateUser(ctx context.Context, input *database.CreateUserParams) error {
	err := repo.queries.CreateUser(ctx, database.CreateUserParams{
		Name: input.Name,
		Login: input.Login,
		Phone: input.Phone,
	})
	if err != nil {
		return fmt.Errorf("error with creating user: %w", err)
	}
	return nil
}
func (repo *UsersRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	err := repo.queries.DeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("error with deliting user: %w", err)
	}
	return nil
}
func (repo *UsersRepository) GetUser(ctx context.Context, id uuid.UUID) (database.GetUserRow, error) {
	user, err := repo.queries.GetUser(ctx, id)
	if err != nil {
		return user, fmt.Errorf("error with deliting user: %w", err)
	}
	return user, err

}
func (repo *UsersRepository) UpdateUser() {
}
func (repo *UsersRepository) IsExistsByLogin(ctx context.Context, login string) (bool, error) {
	isExists, err := repo.queries.IsExistsByLogin(ctx, login)
	if err != nil {
		return isExists, fmt.Errorf("failed to do db request IsExistsByLogin: %w", err)
	}
	return isExists, err
}
func (repo *UsersRepository) IsExistsById(ctx context.Context, id uuid.UUID) (bool, error) {
	isExists, err := repo.queries.IsExistsById(ctx, id)
	if err != nil {
		return isExists, fmt.Errorf("failed to do db request IsExistsById: %w", err)
	}
	return isExists, err
}