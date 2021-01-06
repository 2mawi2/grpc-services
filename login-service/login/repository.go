package login

type IRepository interface {
	GetAllUsers() []User
	GetUsersByRole(role UserRole) []*User
}
type Repository struct{}

var (
	users = []User{
		{
			Name:     "Marius",
			LastName: "Wichtner",
			Email:    "marius.wichtner@email.com",
			Password: "$2a$10$bibXdmTJ9o0Ybizr8byhZe0yc5huOtldVKgDa41gdrn7l4.nXPmwS",
			Role:     ADMIN,
		},
		{
			Name:     "Developer",
			LastName: "DeveloperLastName",
			Email:    "developer.lastname@email.com",
			Password: "$2a$10$bibXdmTJ9o0Ybizr8byhZe0yc5huOtldVKgDa41gdrn7l4.nXPwwS",
			Role:     DEVELOPER,
		},
	}
)

func (r *Repository) GetAllUsers() []User {
	return users
}

func (r *Repository) GetUsersByRole(role UserRole) []*User {
	var usersByRole []*User
	for _, user := range users {
		if user.Role == role {
			usersByRole = append(usersByRole, &user)
		}
	}
	return usersByRole
}
