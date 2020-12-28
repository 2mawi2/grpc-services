package login

import (
	context "context"
	login "github.com/marius/grpc-services/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type Server struct {
	LoginService Service
	login.UnimplementedLoginServiceServer
}

func (s Server) Login(ctx context.Context, request *login.UserRequest) (*login.LoginResponse, error) {
	user := mapUserRequestToUser(request)
	isValidUser := s.LoginService.IsValidUser(user)
	if isValidUser {
		log.Printf("Received valid user: %s", request.Name)
		return &login.LoginResponse{Message: "Login success"}, nil
	} else {
		log.Printf("Received invalid user: %s", request.Name)
		return nil, status.Error(codes.Unauthenticated, "user was not found")
	}
}
