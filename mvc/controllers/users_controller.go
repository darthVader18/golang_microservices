package controllers

import (
	"github.com/gin-gonic/gin"
	"golang_microservices/mvc/services"
	"golang_microservices/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id") ,10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}

		utils.RespondError(c, apiErr)
		//Just return the bad request to the client.
		return
	}

	user, apiErr := services.UsersService.GetUser(userId)
	if apiErr != nil {
		utils.RespondError(c, apiErr)
		//Handle the err and return to the client
		return
	}

	//return user to client
	utils.Respond(c, http.StatusOK, user)

}
