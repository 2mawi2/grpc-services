package main

import (
	"fmt"
	"github.com/marius/grpc-services/login"
	"google.golang.org/grpc"
	"log"
	"net"
)

//go:generate protoc --go_out=plugins=grpc:login login.proto
func SetupGRPCServer(loginService login.Service) {

	port := "9000"
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen to port %s", port)
	}

	grpcServer := grpc.NewServer()
	loginServiceServer := login.Server{
		LoginService: loginService,
	}
	login.RegisterLoginServiceServer(grpcServer, &loginServiceServer)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
