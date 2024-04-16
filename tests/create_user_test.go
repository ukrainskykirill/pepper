package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"github.com/ukrainskykirill/pepper/pkg/types"
)

func TestUserCreatingFailed1(t *testing.T) {
	t.Run(
		"Test case: Faild creation user due to invalid phone number (not russian)",
		func(t *testing.T) {
			userIn := types.UserInput{
				Name:  nil,
				Login: gofakeit.Username(),
				Phone: "+447800767690",
			}
			marshalled, err := json.Marshal(userIn)
			if err != nil {
				fmt.Println("impossible to marshall user")
			}
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(
				"POST",
				"/user/create",
				bytes.NewReader(marshalled),
			)
			testRouter.ServeHTTP(w, req)
			assert.Equal(t, 422, w.Code)
		})
}

func TestUserCreatingFailed2(t *testing.T) {
	t.Run(
		"Test case: Faild creation user due to login already exist",
		func(t *testing.T) {
			prepareTestDatabase()
			userIn := types.UserInput{
				Name:  nil,
				Login: User1Login,
				Phone: "79158865662",
			}
			marshalled, err := json.Marshal(userIn)
			if err != nil {
				fmt.Println("impossible to marshall user")
			}
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(
				"POST",
				"/user/create",
				bytes.NewReader(marshalled),
			)
			testRouter.ServeHTTP(w, req)
			assert.Equal(t, 409, w.Code)
		})
}

func TestSuccessfulUserCreating(t *testing.T) {
	t.Run(
		"Test case: Successful creation user creation",
		func(t *testing.T) {
			var userIn types.UserInput
			err := gofakeit.Struct(&userIn)
			if err != nil {
				fmt.Println("error with fakeit struct")
			}
			marshalled, err := json.Marshal(userIn)
			if err != nil {
				fmt.Println("impossible to marshall user")
			}
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(
				"POST",
				"/user/create",
				bytes.NewReader(marshalled),
			)
			testRouter.ServeHTTP(w, req)
			assert.Equal(t, 200, w.Code)
		})
}
