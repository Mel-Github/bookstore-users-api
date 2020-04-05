package app

import (
	"github.com/mel-github/bookstore-users-api/controllers/ping"
	"github.com/mel-github/bookstore-users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("users/:user_id", users.GetUser)
	// router.GET("users/search", controllers.SearchUser)
	router.POST("users", users.CreateUser)
}
