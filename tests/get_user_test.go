package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"fmt"


	"github.com/ukrainskykirill/pepper/pkg/api"
	"github.com/ukrainskykirill/pepper/pkg/repositories"
	"github.com/ukrainskykirill/pepper/pkg/services"
	"github.com/ukrainskykirill/pepper/pkg/database"
	"github.com/ukrainskykirill/pepper/pkg/config"
	"github.com/ukrainskykirill/pepper/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	prepareTestDatabase()
	config := config.NewConfig()
	validator := types.NewValidator()
	db := database.NewPostgresPool(config)
	queries := database.New(db)
	repositories := repositories.NewRepositories(queries)
	services := services.NewServices(repositories, validator)
	handlers := api.NewHandler(services)
	router := api.InitServer(handlers)
	w := httptest.NewRecorder()
	url := fmt.Sprintf("/user/%s", User1Id)
	
	req, _ := http.NewRequest("GET", url, nil)
	router.ServeHTTP(w, req)
	fmt.Println(w.Body)
	assert.Equal(t, 202, w.Code)
}