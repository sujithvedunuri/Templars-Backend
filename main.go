package main

import (
	"github.com/gin-gonic/gin"
	"templars.com/backend/controllers"
	"templars.com/backend/resources/database"
)

func main() {
	r := gin.Default()
	database.Init()
	r.POST("/login", controllers.LoginController)
	r.GET("/errors", controllers.GetErrorsOfCsv)
	r.POST("/adderror", controllers.AddErrorList)
	r.GET("/delete", controllers.DeleteErrorFromTable)
	r.Run(":8080")

}
