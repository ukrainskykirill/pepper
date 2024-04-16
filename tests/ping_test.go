package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ukrainskykirill/pepper/pkg/api"
	"github.com/ukrainskykirill/pepper/pkg/repositories"
	"github.com/ukrainskykirill/pepper/pkg/services"
	"github.com/ukrainskykirill/pepper/pkg/database"
	"github.com/ukrainskykirill/pepper/pkg/config"
	"github.com/ukrainskykirill/pepper/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestPingPongRoute(t *testing.T) {
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
	req, _ := http.NewRequest("GET", "/user/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}