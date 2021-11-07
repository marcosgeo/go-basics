package services

import (
	"github.com/marcosgeo/go-basics/mvc/domain"
	"github.com/marcosgeo/go-basics/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	// return data from domain DAO
	return domain.GetUser(userId)
}
