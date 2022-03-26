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
	r.GET("/errors", controllers.GetErrorsOfCsv)
	r.POST("/addError", controllers.AddErrorList)
	// r.GET("/deleteErro",controllers.DeleteErrorFromTAab)
	r.Run(":8080")

}
