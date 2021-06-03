package app

import (
	"github.com/dompaksi/bookstore_users-api/controllers"
	"github.com/dompaksi/bookstore_users-api/controllers/users"
)

func UsersMapUrls() {
	router.POST("/users", users.Create)
	router.GET("/users/:userId", users.Get)
	router.PUT("/users/:userId", users.Update)
	router.PATCH("/users/:userId", users.Update)
	router.DELETE("/users/:userId", users.Delete)
	router.GET("/internal/users/search", users.Search)
	router.POST("/users/login", users.Login)
}

func PingMapUrls() {
	router.GET("/ping", controllers.Ping)
}
