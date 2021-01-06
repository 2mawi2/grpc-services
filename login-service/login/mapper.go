package login

import login "github.com/marius/grpc-services/gen"

func mapProtoUserToUser(protoUser *login.User) User {
	return User{
		Name:     protoUser.Name,
		LastName: protoUser.LastName,
		Email:    protoUser.Email,
		Password: protoUser.Password,
	}
}

func mapUserToProtoUser(user *User) login.User {
	return login.User{
		Name:     user.Name,
		LastName: user.LastName,
		Email:    user.Email,
		Password: user.Password,
	}
}

func mapUsersToProtoUsers(users []*User) []*login.User {
	var protoUsers []*login.User
	for _, user := range users {
		protoUser := mapUserToProtoUser(user)
		protoUsers = append(protoUsers, &protoUser)
	}
	return protoUsers
}

func mapProtoUserRoleToUserRole(enum *login.UserRole) UserRole {
	var userRole UserRole
	switch *enum {
	case login.UserRole_ADMIN:
		userRole = ADMIN
	case login.UserRole_DEVELOPER:
		userRole = DEVELOPER
	}
	return userRole
}
