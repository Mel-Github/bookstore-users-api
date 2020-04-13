package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mel-github/bookstore-users-api/logger"
)

var (
	router = gin.Default()
)

// StartApplication method
func StartApplication() {
	mapUrls()

	logger.Info("about to start application")
	router.Run(":8080")
}
