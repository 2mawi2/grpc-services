package login

import (
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Repository IRepository
}

func (s *Service) isValidUser(loginUser User) bool {
	users := s.Repository.getAllUsers()
	for _, user := range users {
		isValidEmail := loginUser.Email == user.Email
		isValidPassword := s.isValidPassword(user, loginUser)
		if isValidEmail && isValidPassword {
			return true
		}
	}
	return false
}

func (s *Service) isValidPassword(user User, loginUser User) bool {
	hashedPassword := user.Password
	clearTextPassword := loginUser.Password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(clearTextPassword))
	return err == nil
}
