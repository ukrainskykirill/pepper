package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	prepareTestDatabase()
	w := httptest.NewRecorder()
	url := fmt.Sprintf("/user/%s", User1Id)
	req, _ := http.NewRequest("GET", url, nil)
	testRouter.ServeHTTP(w, req)
	assert.Equal(t, 202, w.Code)
}