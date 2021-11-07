package services

import (
	"github.com/marcosge/go-basics/mvc/domain"
)

func GetUser(userId int64) (*domain.User, error) {
	return domain.GetUser(userId)
}
