package services

import (
	"golang_microservices/mvc/domain"
	"golang_microservices/mvc/utils"
	"net/http"
)

type itemsService struct {}

var (
	ItemsService itemsService
)

func (s *itemsService) GetItem(itemId string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message: "Implement me",
		StatusCode: http.StatusInternalServerError,
	}
}
