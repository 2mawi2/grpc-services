package login

import (
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

//go:generate mockery --name=IRepository --structname=MockIRepository --output=../domain/mocks --inpackage
var (
	repository = new(MockIRepository)
)

func generateTestHash(testPassword string) []byte {
	testHash, _ := bcrypt.GenerateFromPassword([]byte(testPassword), bcrypt.DefaultCost)
	return testHash
}

func getTestExistingUsers() []User {
	hashedTestPassword := generateTestHash("secretPassword")
	existingUsers := []User{
		{
			Name:     "Marius",
			LastName: "Wichtner",
			Email:    "marius.wichtner@email.com",
			Password: string(hashedTestPassword),
		},
	}
	return existingUsers
}

func CreateService() *Service {
	textExistingUsers := getTestExistingUsers()
	repository.On("getAllUsers", mock.Anything).Return(textExistingUsers)
	service := Service{Repository: repository}
	return &service
}

func Test_isValidUser_should_return_true_for_authenticate_user(t *testing.T) {
	// arrange
	authenticatedUser := User{Name: "M", LastName: "W", Email: "marius.wichtner@email.com", Password: "secretPassword"}
	// act
	isAuthenticated := CreateService().IsValidUser(authenticatedUser)
	// assert
	assert.Equal(t, isAuthenticated, true)
}

func Test_isValidUser_should_return_false_for_invalid_email(t *testing.T) {
	// arrange
	invalidEmailUser := User{Name: "M", LastName: "WWichtner", Email: "INVALID", Password: "secretPassword"}
	// act
	isAuthenticated := CreateService().IsValidUser(invalidEmailUser)
	// assert
	assert.Equal(t, isAuthenticated, false)
}

func Test_isValidUser_should_return_false_for_invalid_password(t *testing.T) {
	// arrange
	invalidPasswordUser := User{Name: "M", LastName: "W", Email: "marius.wichtner@email.com", Password: "INVALID"}
	// act
	isAuthenticated := CreateService().IsValidUser(invalidPasswordUser)
	// assert
	assert.Equal(t, isAuthenticated, false)
}
