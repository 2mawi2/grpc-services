package main

import (
	"context"
	gen "github.com/marius/grpc-services/gen"
	"github.com/marius/grpc-services/login"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
)

var lis *bufconn.Listener

func init() {
	const bufSize = 1024 * 1024
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	server := login.Server{
		LoginService: NewService(),
	}
	gen.RegisterLoginServiceServer(s, &server)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func Test_LoginServer_responds_with_message_when_grpc_login_successful(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := gen.NewLoginServiceClient(conn)

	loginResponse, err := client.Login(ctx, &gen.User{
		Name:     "Test",
		LastName: "Test",
		Email:    "marius.wichtner@email.com",
		Password: "secretPassword",
	})

	assert.NotNil(t, loginResponse)
	assert.NotEmpty(t, loginResponse.Message)
	assert.Nil(t, err)
}

func Test_LoginServer_responds_with_error_when_grpc_login_unsuccessful(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := gen.NewLoginServiceClient(conn)

	loginResponse, err := client.Login(ctx, &gen.User{
		Name:     "Test",
		LastName: "Test",
		Email:    "marius.wichtner@email.com",
		Password: "wrongPassord",
	})

	assert.NotNil(t, err)
	assert.Nil(t, loginResponse)
}
