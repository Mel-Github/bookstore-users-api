package users

import (
	"fmt"

	"github.com/mel-github/bookstore-users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

// Save - Save the user record to persistent layer
func (user *User) Save() *errors.RestErr {
	current := usersDB[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exist", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exist", user.ID))
	}

	// user does not exist. Creating
	usersDB[user.ID] = user
	return nil
}

// Get - Get the user record from persistent layer
func (user *User) Get() *errors.RestErr {
	result := usersDB[user.ID]
	if result == nil {
		return errors.NewNotFoundtError(fmt.Sprintf("user %d not found", user.ID))
	}

	user.ID = result.ID
	user.Email = result.Email
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.DateCreated = result.DateCreated

	return nil
}
