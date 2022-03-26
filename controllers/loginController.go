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
	if err := c.BindJSON(&data); err != nil {
		log.Fatal("error")
	}
	fmt.Println(data)
	user := daos.GetLoginCredentials(data)
    if (user.UserName == data.UserName) && (user.Password == data.Password) {
        c.IndentedJSON(http.StatusOK, user)
        fmt.Println(user)
    }

}
