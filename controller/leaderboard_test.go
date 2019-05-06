package controller_test

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/haithanh079/go-leaderboard/routers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddUserToLeaderboard(t *testing.T) {
	body := gin.H{
		"username": "haithanh079",
		"score": "10",
	}
	router := routers.Router{}
	router.Init()
	w := performRequest(&router, "POST", "/leaderboard/add")
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	value, exists := response["data"]
	fmt.Println(value)
	assert.Nil(t, err)

	assert.True(t, exists)

	assert.Equal(t, body, value)
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}