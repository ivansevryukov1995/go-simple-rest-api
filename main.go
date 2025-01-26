// main.go
package main

import (
	"github.com/ivansevryukov1995/go-simple-rest-api/api"
	"github.com/ivansevryukov1995/go-simple-rest-api/database"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Connect()
	db.AutoMigrate(&api.User{})

	r := gin.Default()
	userHandler := api.UserHandler{DB: db}

	r.POST("/users", userHandler.CreateUser)
	r.GET("/user/:id", userHandler.GetUser)
	r.PATCH("/user/:id", userHandler.EditUser)

	r.Run(":8080")
}
