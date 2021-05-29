package app

import (
	"github.com/dompaksi/bookstore_users-api/controllers"
	"github.com/dompaksi/bookstore_users-api/controllers/users"
)

func UsersMapUrls() {
	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	//router.GET("/users/search", controllers.FindUser)
}

func PingMapUrls() {
	router.GET("/ping", controllers.Ping)
}
