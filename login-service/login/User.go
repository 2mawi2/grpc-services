package login

type User struct {
	Name     string   `json:"name"`
	LastName string   `json:"lastName"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Role     UserRole `json:"type"`
}
type UserRole int32

const (
	ADMIN     UserRole = 0
	DEVELOPER          = 1
)
