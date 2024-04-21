package tests

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-testfixtures/testfixtures/v3"
	_ "github.com/lib/pq"
	"github.com/ukrainskykirill/pepper/pkg/api"
	"github.com/ukrainskykirill/pepper/pkg/config"
	"github.com/ukrainskykirill/pepper/pkg/database"
	"github.com/ukrainskykirill/pepper/pkg/repositories"
	"github.com/ukrainskykirill/pepper/pkg/services"
	"github.com/ukrainskykirill/pepper/pkg/types"
)

var (
	db         *sql.DB
	fixtures   *testfixtures.Loader
	User1Login string = "johndoe"
	User1Id    string = "4d9db4cd-78d4-4fea-8d66-bfc146f58608"
	testRouter *gin.Engine
)

func TestMain(m *testing.M) {
	var err error
	db, err = sql.Open("postgres", "user=test_user password=test_password dbname=test_user host=postgres port=5432")
	if err != nil {
		fmt.Println("error sql open")
	}
	fixtures, err = testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("postgres"),
		testfixtures.Files("testdata/users.yaml"),
	)
	if err != nil {
		fmt.Println(err)
		fmt.Println("error sql conntext")
	}
	config := config.NewConfig()
	validator := types.NewValidator()
	db := database.NewPostgresPool(config)
	queries := database.New(db)
	repositories := repositories.NewRepositories(queries)
	services := services.NewServices(repositories, validator)
	handlers := api.NewHandler(services)
	testRouter = api.InitServer(handlers)
	exitVal := m.Run()
	fmt.Println("Do stuff AFTER the tests!")
	os.Exit(exitVal)
}

func prepareTestDatabase() {
	if err := fixtures.Load(); err != nil {
		fmt.Println("failed prepareTestDatabase")
	}
	fmt.Println("prepareTestDatabase")
}
