package services

import (
	"net/http"

	"github.com/marcosgeo/go-basics/mvc/domain"
	"github.com/marcosgeo/go-basics/mvc/utils"
)

type itemsService struct{}

var (
	ItemsService itemsService
)

func (s *itemsService) GetItem(itemId string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "not implemented",
		StatusCode: http.StatusInternalServerError,
	}
}
