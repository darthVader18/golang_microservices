package domain

import (
	"fmt"
	"golang_microservices/mvc/utils"
	"net/http"
)

var(
	users = map[int64]*User{
		123 : {Id: 123, FirstName: "Chaitanya Singh", LastName: "Bisht", Email: "csb.net.in@gmail.com"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message: fmt.Sprintf("User %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code: "not_found",
	}
}

