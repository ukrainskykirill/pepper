package api

import (
	"github.com/gin-gonic/gin"
    "github.com/ukrainskykirill/pepper/pkg/services"
)

type Handler struct {
	Users *UsersHandler
}

func NewHandler(services *services.Services) *Handler {
	return &Handler{
		Users: NewUsersHandler(services.Users),
	}
}

func InitServer(handler *Handler) {
    mainRouter := gin.Default()
    userRouter := mainRouter.Group("/user")
    {
        userRouter.POST("/create", handler.Users.createUser)
        userRouter.DELETE("/:id", handler.Users.deleteUser)
        userRouter.GET("/:id", handler.Users.getUser)
        userRouter.PUT("/:id", handler.Users.updateUser)
    }
    mainRouter.Run()
}