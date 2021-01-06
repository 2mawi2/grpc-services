package main

import (
	"github.com/marius/grpc-services/login"
	"sync"
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
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func() {
		SetupGRPCServer(loginService)
		wg.Done()
	}()
	go func() {
		SetupGRPCGatewayEndpoint()
		wg.Done()
	}()
	wg.Wait()
}
