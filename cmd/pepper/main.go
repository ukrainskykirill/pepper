package main

import (
	"github.com/ukrainskykirill/pepper/pkg/api"
	"github.com/ukrainskykirill/pepper/pkg/repositories"
	"github.com/ukrainskykirill/pepper/pkg/services"
	"github.com/ukrainskykirill/pepper/pkg/database"
	"github.com/ukrainskykirill/pepper/pkg/config"
)

func main() {
	config := config.NewConfig()
	db := database.NewPostgresPool(config)
	queries := database.New(db)
	repositories := repositories.NewRepositories(queries)
	services := services.NewServices(repositories)
	handlers := api.NewHandler(services)
	api.InitServer(handlers)
}
