package main

//go:generate bash ./gen-protos.sh
//go:generate bash ./gen-open-api-definitions.sh

import (
	"context"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	gen "github.com/marius/grpc-services/gen"
	"github.com/marius/grpc-services/login"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9000", "gRPC server endpoint")
)

func SetupGRPCGatewayEndpoint() {
	mux := runtime.NewServeMux()
	ctx := context.Background()
	dialOptions := []grpc.DialOption{grpc.WithInsecure()}
	if err := gen.RegisterLoginServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, dialOptions); err != nil {
		log.Fatalf("failed to register service handler from endpoint %s", *grpcServerEndpoint)
	}
	grpcGatewayPort := "8081"
	if err := http.ListenAndServe(fmt.Sprintf(":%s", grpcGatewayPort), mux); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func SetupGRPCServer(loginService login.Service) {
	grpcPort := "9000"
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen to grpcPort %s", grpcPort)
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
