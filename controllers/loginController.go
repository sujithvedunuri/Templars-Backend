package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"templars.com/backend/beans"
	"templars.com/backend/resources/database"
)

func LoginController(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var data beans.LoginBean
	if err := c.BindJSON(&data); err != nil {
		log.Fatal("error")
	}
	var user beans.Employee

	database.Db.Where("user_name = ? ", data.UserName).First(&user)
	if (user.UserName == data.UserName) && (user.Password == data.Password) {
		c.IndentedJSON(http.StatusOK, user)
		fmt.Println(user)
	}

}
