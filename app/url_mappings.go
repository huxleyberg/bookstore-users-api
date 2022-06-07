package app

import (
	"github.com/huxleyberg/bookstore-users-api/controllers/ping"
	"github.com/huxleyberg/bookstore-users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	router.GET("/users/search", users.SearchUser)
	router.POST("/users", users.CreateUser)
}
