package daos

import (
	"fmt"

	"templars.com/backend/beans"
	"templars.com/backend/resources/database"
)

func FetchErrorCsv() []beans.CsvErrorBean {
	var csvErrorScheet []beans.CsvErrorBean
	database.Db.Find(&csvErrorScheet)
	fmt.Println(csvErrorScheet)
	return csvErrorScheet
}

func AddErrorCsvTable(csvError beans.CsvErrorBean) {
	database.Db.Create(&csvError)
}
func RemoveErrorFromTable(id int) (beans.CsvErrorBean, error) {
	var ErrorData beans.CsvErrorBean
	fmt.Println(id)
	result := database.Db.Where("id = ?", id).First(&ErrorData)
	if result.Error != nil {
		return ErrorData, result.Error
	} else {
		result := database.Db.Delete(&ErrorData, id)
		return ErrorData, result.Error

	}
}
