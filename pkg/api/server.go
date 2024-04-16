package api

import (
	"github.com/gin-gonic/gin"
    "github.com/ukrainskykirill/pepper/pkg/services"
)

type Handler struct {
	Users *UsersHandler
}

func NewHandler(services services.IServices) *Handler {
	return &Handler{
		Users: NewUsersHandler(services.GetUsersService()),
	}
}

func InitServer(handler *Handler) *gin.Engine {
    mainRouter := gin.Default()
    userRouter := mainRouter.Group("/user")
    {
        userRouter.POST("/create", handler.Users.createUser)
        userRouter.DELETE("/:id", handler.Users.deleteUser)
        userRouter.GET("/:id", handler.Users.getUser)
        userRouter.PATCH("/:id", handler.Users.updateUser)
        userRouter.GET("/ping", handler.Users.ping)
    }
    return mainRouter
}