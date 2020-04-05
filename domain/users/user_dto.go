package users

import (
	"strings"

	"github.com/mel-github/bookstore-users-api/utils/errors"
)

// User - user key fields
type User struct {
	ID          int64  `json:"ID"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

// Validate - Check user email is valid
func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}
