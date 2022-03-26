package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"templars.com/backend/daos"
)

func GetErrorsOfCsv(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	csvErroList := daos.FetchErrorMail()
	c.IndentedJSON(http.StatusOK, csvErroList)

}
