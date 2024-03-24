package repositories

import "github.com/ukrainskykirill/pepper/pkg/database"

type UsersRepository struct {
	queries *database.Queries
}

func NewUserRepository(queries *database.Queries) *UsersRepository{
	return &UsersRepository{
		queries: queries,
	}
}

func (repo *UsersRepository) CreateUser() {
	return 
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