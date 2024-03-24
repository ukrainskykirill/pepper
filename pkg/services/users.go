package services

import (
	"github.com/ukrainskykirill/pepper/pkg/repositories"
)

type UsersService struct {
	repo *repositories.UsersRepository
}

func NewUserService(repo *repositories.UsersRepository) *UsersService{
	return &UsersService{
		repo: repo,
	}
}

func (service *UsersService) CreateUser() {
	return 
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