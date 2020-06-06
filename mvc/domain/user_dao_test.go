package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
	//Initialization:

	//Execution
	user, err := GetUser(0)

	//Validation
	assert.Nil(t, user, "We were not expecting a user with ID 0")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "User 0 does not exists", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Chaitanya Singh", user.FirstName)
	assert.EqualValues(t, "Bisht", user.LastName)
	assert.EqualValues(t, "csb.net.in@gmail.com", user.Email)

}
