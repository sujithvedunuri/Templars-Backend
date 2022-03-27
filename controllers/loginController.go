package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"templars.com/backend/beans"
	"templars.com/backend/daos"
)

func LoginController(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	var data beans.LoginBean
	if err := c.ShouldBindJSON(&data); err != nil {
		log.Fatal("error while binding to json")
	}

	user := daos.GetLoginCredentials(data)
	if (user.UserName == data.UserName) && (user.Password == data.Password) {
		c.IndentedJSON(http.StatusOK, user)
		// fmt.Println(user)
	} else {
		fmt.Println("Invalid login credentials")
		c.IndentedJSON(401, "Invalid Login Check Credentials")
	}

}
