package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ukrainskykirill/pepper/pkg/services"
	"github.com/ukrainskykirill/pepper/pkg/database"
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
	var userIn database.CreateUserParams
	if err := ctx.ShouldBindJSON(&userIn); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := handler.service.CreateUser(ctx.Request.Context(), &userIn)
	if errors.As(err, &services.InvalidInputData{}) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	} else if errors.As(err, &services.AlreadyExistsByLogin{}) {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
func (handler *UsersHandler) deleteUser(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	parsedId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	err = handler.service.DeleteUser(ctx.Request.Context(), parsedId)
	if errors.As(err, &services.UserDoesntExists{}) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
func (handler *UsersHandler) getUser(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	parsedId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	user, err := handler.service.GetUser(ctx.Request.Context(), parsedId)
	if errors.As(err, &services.UserDoesntExists{}) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else { ctx.JSON(http.StatusAccepted, user) }

}
func (handler *UsersHandler) updateUser(ctx *gin.Context) {
	handler.service.UpdateUser()
}