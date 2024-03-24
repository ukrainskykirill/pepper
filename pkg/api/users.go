package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ukrainskykirill/pepper/pkg/services"
)

type UsersHandler struct {
	service *services.UsersService
}

func NewUsersHandler(service *services.UsersService) *UsersHandler{
	return &UsersHandler{
		service: service,
	}
}

func (handler *UsersHandler) createUser(ctx *gin.Context) {
	handler.service.CreateUser()
	return
}
func (handler *UsersHandler) deleteUser(ctx *gin.Context) {
	handler.service.CreateUser()
	return
}
func (handler *UsersHandler) getUser(ctx *gin.Context) {
	handler.service.CreateUser()
	return
}
func (handler *UsersHandler) updateUser(ctx *gin.Context) {
	handler.service.CreateUser()
	return
}