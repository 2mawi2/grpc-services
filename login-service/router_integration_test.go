package main

import (
	"bytes"
	"encoding/json"
	"github.com/marius/grpc-services/login"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)


func Test_LoginRouter_responds_with_200_when_request_valid(t *testing.T) {
	service := NewService()
	ts := httptest.NewServer(SetupRouter(service))
	defer ts.Close()

	jsonStr, _ := json.Marshal(login.User{
		Email:    "marius.wichtner@email.com",
		Password: "secretPassword",
	})
	resp, err := http.Post(ts.URL+"/login", "application/json", bytes.NewBuffer(jsonStr))

	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, 200)
}

func Test_LoginRouter_responds_with_401_when_user_unauthorized(t *testing.T) {
	service := NewService()
	ts := httptest.NewServer(SetupRouter(service))
	defer ts.Close()

	jsonStr, _ := json.Marshal(login.User{
		Email:    "marius.wichtner@email.com",
		Password: "wrong password",
	})
	resp, err := http.Post(ts.URL+"/login", "application/json", bytes.NewBuffer(jsonStr))

	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, 401)
}

func Test_LoginRouter_responds_with_400_when_user_body_is_empty(t *testing.T) {
	service := NewService()
	ts := httptest.NewServer(SetupRouter(service))
	defer ts.Close()

	resp, err := http.Post(ts.URL+"/login", "application/json", nil)

	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, 400)
}
