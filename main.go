package main

import (
	"github.com/gin-gonic/gin"
	"templars.com/backend/controllers"
	"templars.com/backend/resources/database"
)

func main() {
	r := gin.Default()
	database.Init()
	r.GET("/login", controllers.LoginController)
	r.Run(":8080")

}
