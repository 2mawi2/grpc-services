package main

import (
	"github.com/marius/grpc-services/login"
)

func main() {
	loginService := NewService()
	setupProtocols(loginService)
}

func NewService() login.Service {
	repository := login.Repository{}
	service := login.Service{Repository: &repository}
	return service
}

func setupProtocols(loginService login.Service) {
	go func() {
		SetupRouter(loginService).Run()
	}()
	SetupGRPCServer(loginService)
}
