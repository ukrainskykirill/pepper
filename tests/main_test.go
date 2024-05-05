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
	user1Id    string = "4d9db4cd-78d4-4fea-8d66-bfc146f58608"
	testRouter *gin.Engine
)

func TestMain(m *testing.M) {
	var err error
	config := config.NewConfig()
	validator := types.NewValidator()
	db_con := database.NewPostgresPool(config)
	queries := database.New(db_con)
	repositories := repositories.NewRepositories(queries)
	services := services.NewServices(repositories, validator)
	handlers := api.NewHandler(services)
	testRouter = api.InitServer(handlers)
	db , err = sql.Open(
		"postgres", 
		fmt.Sprintf(
			"user=%s password=%s dbname=%s sslmode=disable host=%s port=%s",
			config.PostgresUser, 
			config.PostgresPassword,
			config.PostgresDB,
			config.PostgresHost,
			config.PostgresPort,
		),
	)
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
	exitVal := m.Run()
	os.Exit(exitVal)
}

func prepareTestDatabase() {
	if err := fixtures.Load(); err != nil {
		fmt.Println("failed prepareTestDatabase")
	}
	fmt.Println("prepareTestDatabase")
}
