package login

func mapUserRequestToUser(userRequest *UserRequest) User {
	return User{
		Name:     userRequest.Name,
		LastName: userRequest.LastName,
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}
}
