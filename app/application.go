package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	PingMapUrls()
	UsersMapUrls()
	router.Run(":8000")
}
