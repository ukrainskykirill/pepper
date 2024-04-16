package tests

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
	_ "github.com/lib/pq"
)

var (
	db *sql.DB
	fixtures *testfixtures.Loader
	User1Id string = "4d9db4cd-78d4-4fea-8d66-bfc146f58608"
)

func TestMain(m *testing.M) {
	var err error
	db, err = sql.Open("postgres", "user=postgres password=postgres dbname=pepper_test sslmode=disable host=localhost port=32432")
	if err != nil {
			fmt.Println("error sql open")
	}
	fixtures, err = testfixtures.New(
			testfixtures.Database(db),
			testfixtures.Dialect("postgres"),
			testfixtures.Files("testdata/users.yaml"),
	)
	if err != nil {
		fmt.Println("error sql conntext")
	}
	exitVal := m.Run()
	fmt.Println("Do stuff AFTER the tests!")
	os.Exit(exitVal)
}

func prepareTestDatabase() {
	if err := fixtures.Load(); err != nil {
		fmt.Println("+___________________")
		fmt.Println(err)
		fmt.Println("failed prepareTestDatabase")
	}
	fmt.Println("prepareTestDatabase")
}