package domain

import (
	"fmt"
	"log"
	"net/http"

	"github.com/marcosgeo/go-basics/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Antonio", LastName: "Marcos", Email: "marcosgeo@yahoo.com"},
	}

	UserDao usersDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type usersDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct{}

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	log.Println("We are accessing the database")
	user := users[userId]
	if user == nil {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("User %v does not exists", userId),
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	}
	return user, nil
}
