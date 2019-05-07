package routers_test

import (
	"bytes"
	"github.com/haithanh079/go-leaderboard/routers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLeaderboardRouters(t *testing.T) {
	router := routers.Router{}
	router.Init(true)
	w := httptest.NewRecorder()
	/*
	Ping to Get Leaderboard
	 */
	req, err := http.NewRequest("GET", "/leaderboard/get", nil)
	router.ServeHTTP(w, req)
	if err != nil {
		t.Log(err)
	}else {
		t.Log("Test Get leaderboard success!")
	}
	assert.Equal(t, err, nil, "GET LEADERBOARD OK")
	/*
	Ping to Add user -> Leaderboard
	 */
	var jsonBodyRequest = []byte(`{"username":"haithanh","score":"1",}`)
	req, err = http.NewRequest("POST", "/leaderboard/add", bytes.NewBuffer(jsonBodyRequest))
	router.ServeHTTP(w, req)
	if err != nil {
		t.Log(err)
	}else {
		t.Log("Test Add user success!")
	}
	assert.Equal(t, err, nil, "ADD USER -> LEADERBOARD OK")
	/*
	Add an empty request
	 */
	jsonBodyRequest = []byte(`{"username":"","score":"",}`)
	req, err = http.NewRequest("POST", "/leaderboard/add", nil)
	router.ServeHTTP(w, req)
	if err != nil{
		t.Log(err)
	}else {
		t.Log("Test Add user with empty body success!")
	}
	assert.Equal(t, err, nil, "ADD EMPTY REQUEST BODY OK")
}

