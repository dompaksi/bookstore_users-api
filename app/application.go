package app

import (
	"github.com/dompaksi/bookstore_utils-go/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	PingMapUrls()
	UsersMapUrls()
	logger.Info("about to start the application...")
	router.Run(":8000")
}
