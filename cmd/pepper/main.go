package main

import (
	"fmt"

	"github.com/ukrainskykirill/pepper/pkg/api"
	"github.com/ukrainskykirill/pepper/pkg/config"
	"github.com/ukrainskykirill/pepper/pkg/database"
	"github.com/ukrainskykirill/pepper/pkg/repositories"
	"github.com/ukrainskykirill/pepper/pkg/services"
	"github.com/ukrainskykirill/pepper/pkg/types"
)

func main() {
	config := config.NewConfig()
	validator := types.NewValidator()
	db := database.NewPostgresPool(config)
	queries := database.New(db)
	repositories := repositories.NewRepositories(queries)
	services := services.NewServices(repositories, validator)
	handlers := api.NewHandler(services)
	router := api.InitServer(handlers)
	err := router.Run()
	if err != nil {
		fmt.Println(err)
	}
}
