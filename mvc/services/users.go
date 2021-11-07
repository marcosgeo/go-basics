package services

import (
	"github.com/marcosgeo/go-basics/mvc/domain"
	"github.com/marcosgeo/go-basics/mvc/utils"
)

type usersService struct{}

var (
	UsersService usersService
)

func (u *usersService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	// return data from domain DAO
	user, err := domain.UserDao.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
