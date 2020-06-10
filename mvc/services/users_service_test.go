package services

import (
	"github.com/stretchr/testify/assert"
	"golang_microservices/mvc/utils"
	"golang_microservices/mvc/domain"
	"net/http"
	"testing"
)

var (
	userDaoMock usersDaoMock

	getUserFunction func(userId int64) (*domain.User, *utils.ApplicationError)
)

func init() {
	domain.UserDao = &usersDaoMock{}
}

type usersDaoMock struct {}

func (m *usersDaoMock) GetUser(userId int64)(*domain.User, *utils.ApplicationError) {
	return getUserFunction(userId)
}

func TestGetUserNotFoundInDatabase(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message: "User 0 does not exists",
		}
	}
	user, err := UsersService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.Equal(t, "User 0 does not exists", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
			Id: 123,
		}, nil
	}

	user, err := UsersService.GetUser(123)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.Equal(t, 123, user.Id)
}
