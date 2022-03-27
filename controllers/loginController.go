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

	fmt.Println("data ", data)
  
	user := daos.GetLoginCredentials(data)
	if (user.UserName == data.UserName) && (user.Password == data.Password){
		c.IndentedJSON(http.StatusOK, user)
		fmt.Println(user)
  }
}
