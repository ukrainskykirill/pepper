package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingPongRoute(t *testing.T) {
	prepareTestDatabase()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/ping", nil)
	testRouter.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}