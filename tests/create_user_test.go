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

func TestUserCreatingFailed(t *testing.T) {
	var userIn types.UserInput
	err := gofakeit.Struct(&userIn)
	if err != nil {
		fmt.Println("error with fakeit struct")
	}
    testCases := []struct {
        testName   string
        userInput  types.UserInput
        expectedCode int
    }{
        {
            testName: "Test case: Faild creation user due to invalid phone number (not russian)",
            userInput: types.UserInput{
                Name:  nil,
                Login: gofakeit.Username(),
                Phone: "+447800767690",
            },
            expectedCode: 422,
        },
		{
            testName: "Test case: Successful creation user creation",
            userInput: userIn,
            expectedCode: 200,
        },
        
    }
    for _, tc := range testCases {
        t.Run(tc.testName, func(t *testing.T) {
            marshalled, err := json.Marshal(tc.userInput)
            if err != nil {
                t.Fatal("impossible to marshall user")
            }

            w := httptest.NewRecorder()
            req, _ := http.NewRequest(
                "POST",
                "/user/create",
                bytes.NewReader(marshalled),
            )

            testRouter.ServeHTTP(w, req)
            assert.Equal(t, tc.expectedCode, w.Code)
        })
    }
}