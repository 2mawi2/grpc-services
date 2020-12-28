package login

import (
	context "context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type Server struct {
	LoginService Service
}

func (s Server) Login(ctx context.Context, request *UserRequest) (*LoginResponse, error) {
	user := mapUserRequestToUser(request)
	isValidUser := s.LoginService.IsValidUser(user)
	if isValidUser {
		log.Printf("Received valid user: %s", request.Name)
		return &LoginResponse{Message: "Login success"}, nil
	} else {
		log.Printf("Received invalid user: %s", request.Name)
		return nil, status.Error(codes.Unauthenticated, "user was not found")
	}
}
