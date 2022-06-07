package services

import (
	"github.com/huxleyberg/bookstore-users-api/domain/users"
)

func CreateUser(user *users.User) (*users.User, error) {
	return user, nil
}
