package login

type IRepository interface {
	getAllUsers() []User
}

type Repository struct{}

func (r *Repository) getAllUsers() []User {
	return []User{
		{
			Name:     "Marius",
			LastName: "Wichtner",
			Email:    "marius.wichtner@email.com",
			Password: "$2a$10$bibXdmTJ9o0Ybizr8byhZe0yc5huOtldVKgDa41gdrn7l4.nXPmwS",
		},
	}
}
