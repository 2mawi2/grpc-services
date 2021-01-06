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

func (s Server) Login(ctx context.Context, request *login.User) (*login.LoginResponse, error) {
	user := mapProtoUserToUser(request)
	isValidUser := s.LoginService.IsValidUser(user)
	if isValidUser {
		log.Printf("Received valid user: %s", request.Name)
		return &login.LoginResponse{Message: "Login success"}, nil
	} else {
		log.Printf("Received invalid user: %s", request.Name)
		return nil, status.Error(codes.Unauthenticated, "user was not found")
	}
}

func (s Server) GetUserByRole(ctx context.Context, request *login.UserRoleRequest) (*login.UserRoleResponse, error) {
	userRole := mapProtoUserRoleToUserRole(request.UserRole.Enum())
	usersByRole := s.LoginService.GetUsersByRole(userRole)
	protoUsersByRole := mapUsersToProtoUsers(usersByRole)
	if len(usersByRole) > 0 {
		return &login.UserRoleResponse{Users: protoUsersByRole}, nil
	} else {
		return nil, status.Error(codes.NotFound, "No users found with given role")
	}
}
