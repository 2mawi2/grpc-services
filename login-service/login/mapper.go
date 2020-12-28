package login

import login "github.com/marius/grpc-services/gen"

func mapUserRequestToUser(userRequest *login.UserRequest) User {
	return User{
		Name:     userRequest.Name,
		LastName: userRequest.LastName,
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}
}
