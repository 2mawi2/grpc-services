package main

import (
	"fmt"
	gen "github.com/marius/grpc-services/gen"
	"github.com/marius/grpc-services/login"
	"google.golang.org/grpc"
	"log"
	"net"
)

//go:generate  protoc -I . --go_out ./gen/ --go_opt paths=source_relative --go-grpc_out ./gen/ --go-grpc_opt paths=source_relative login.proto
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
	gen.RegisterLoginServiceServer(grpcServer, &loginServiceServer)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
