package repositories

import (
	"context"

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

func (repo *UsersRepository) CreateUser(ctx context.Context, input *types.UserInput) {
	err := repo.queries.CreateUser(ctx, database.CreateUserParams{
		Name: input.Name,
		Login: input.Login,
		Phone: input.Phone,
	})
	if err != nil {
		println(err)
	}
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