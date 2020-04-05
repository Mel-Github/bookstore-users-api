package services

import (
	"github.com/mel-github/bookstore-users-api/domain/users"
	"github.com/mel-github/bookstore-users-api/utils/errors"
)

// CreateUser - Service to create user
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUser - Service to get existing user based on user id
func GetUser(userID int64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}
