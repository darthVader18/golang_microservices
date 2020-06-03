package services

import (
	"golang_microservices/mvc/domain"
	"golang_microservices/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
