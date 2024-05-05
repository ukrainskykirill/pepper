package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	prepareTestDatabase()
	testCases := []struct {
        testName   string
        userId  string
        expectedCode int
    }{
        {
            testName: "Test case: Faild get user due to invalid uuid",
            userId: "User1Id",
            expectedCode: 422,
        },
		{
            testName: "Test case: Successful get user",
            userId: user1Id,
            expectedCode: 200,
        },
        
    }
	for _, tc := range testCases {
        t.Run(tc.testName, func(t *testing.T) {
            w := httptest.NewRecorder()
			url := fmt.Sprintf("/user/%s", tc.userId)
            req, _ := http.NewRequest(
                "GET",
                url,
                nil,
            )
            testRouter.ServeHTTP(w, req)
            assert.Equal(t, tc.expectedCode, w.Code)
        })
    }
}