package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mel-github/bookstore-users-api/domain/users"
	"github.com/mel-github/bookstore-users-api/services"
	"github.com/mel-github/bookstore-users-api/utils/errors"
)

// CreateUser - creates a new user
func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// GetUser - gets an existing user based on user_id
func GetUser(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("Invalid user id. Should be number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

// // SearchUser - return the list of existing users.
// func SearchUser(c *gin.Context) {
// 	c.String(http.StatusNotImplemented, "implementing search users")
// }
