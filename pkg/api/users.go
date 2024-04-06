package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ukrainskykirill/pepper/pkg/services"
	"github.com/ukrainskykirill/pepper/pkg/types"
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
	var userIn types.UserInput
	if err := ctx.ShouldBindJSON(&userIn); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err := handler.service.CreateUser(&userIn); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	}
}
func (handler *UsersHandler) deleteUser(ctx *gin.Context) {
	handler.service.DeleteUser()
	return
}
func (handler *UsersHandler) getUser(ctx *gin.Context) {
	handler.service.GetUser()
	return
}
func (handler *UsersHandler) updateUser(ctx *gin.Context) {
	handler.service.UpdateUser()
	return
}