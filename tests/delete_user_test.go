package tests

// import (
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestDeleteUser(t *testing.T) {
// 	prepareTestDatabase()
// 	w := httptest.NewRecorder()
// 	url := fmt.Sprintf("/user/%s", User1Id)
// 	req, _ := http.NewRequest("DELETE", url, nil)
// 	testRouter.ServeHTTP(w, req)
// 	fmt.Println(w.Body)
// 	assert.Equal(t, 200, w.Code)
// 	w = httptest.NewRecorder()
// 	req, _ = http.NewRequest("GET", url, nil)
// 	testRouter.ServeHTTP(w, req)
// 	fmt.Println(w.Body)
// 	assert.Equal(t, 404, w.Code)
// }