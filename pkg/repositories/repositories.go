package repositories

import "github.com/ukrainskykirill/pepper/pkg/database"


type Repositories struct {
	Users *UsersRepository
}

func NewRepositories(queries *database.Queries) *Repositories {
	return &Repositories{
		Users: NewUserRepository(queries),
	}
}