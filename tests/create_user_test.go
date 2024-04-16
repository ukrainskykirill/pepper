package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"
	"fmt"
	"encoding/json"


	"github.com/ukrainskykirill/pepper/pkg/api"
	"github.com/ukrainskykirill/pepper/pkg/repositories"
	"github.com/ukrainskykirill/pepper/pkg/services"
	"github.com/ukrainskykirill/pepper/pkg/database"
	"github.com/ukrainskykirill/pepper/pkg/config"
	"github.com/ukrainskykirill/pepper/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	config := config.NewConfig()
	validator := types.NewValidator()
	db := database.NewPostgresPool(config)
	queries := database.New(db)
	repositories := repositories.NewRepositories(queries)
	services := services.NewServices(repositories, validator)
	handlers := api.NewHandler(services)
	router := api.InitServer(handlers)
	user := types.UserInput{
		Login: GenString(13),
		Name: nil,
		Phone: "79158965663",

	}
	marshalled, err := json.Marshal(user)
	if err != nil {
		fmt.Println("impossible to marshall user")
	}
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/user/create", bytes.NewReader(marshalled))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}