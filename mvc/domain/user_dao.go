package domain

import (
	"fmt"
)

var (
	users = map[int64]*User{
		123: {Id: 1, FirstName: "Antonio", LastName: "Marcos", Email: "marcosgeo@yahoo.com"},
	}
)

func GetUser(userId int64) (*User, error) {
	user := users[userId]
	if user == nil {
		return nil, fmt.Errorf("User %v was not found", userId)
	}
	return user, nil
}
