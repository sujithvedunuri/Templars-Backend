package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"templars.com/backend/beans"
	"templars.com/backend/daos"
)

func GetErrorsOfCsv(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	csvErroList := daos.FetchErrorCsv()
	c.IndentedJSON(http.StatusOK, csvErroList)
}

func AddErrorList(c *gin.Context) {
	c.Writer.Header().Set("Acces-Control-Allow-Origin", "*")
	var ErrorsCsv beans.CsvErrorBean
	if err := c.ShouldBindJSON(&ErrorsCsv); err == nil {
		daos.AddErrorCsvTable(ErrorsCsv)
	} else {
		log.Fatal("error adding datapoint to database")
	}
	c.IndentedJSON(http.StatusOK, ErrorsCsv)

}
func DeleteErrorFromTable(c *gin.Context) {
	// id := c.Param("id")
	// daos.RemoveErrorFromTable(strconv.(id))
}
