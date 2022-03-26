package daos

import (
	"fmt"

	"templars.com/backend/beans"
	"templars.com/backend/resources/database"
)

func FetchErrorCsv() []beans.CsvErrorBean {

	var csvErrorScheet []beans.CsvErrorBean
	database.Db.Where("issuer=?", "nayeem").Find(&csvErrorScheet)
	fmt.Println(csvErrorScheet)
	return csvErrorScheet
}

func AddErrorCsvTable(csvError beans.CsvErrorBean) {
	database.Db.Create(&csvError)
}
func RemoveErrorFromTable(mail string) (beans.CsvErrorBean, error) {
	var ErrorData beans.CsvErrorBean
	fmt.Println(mail)
	result := database.Db.Where("email = ?", mail).Find(&ErrorData).Delete(&ErrorData)
	if result.Error != nil {
		return ErrorData, result.Error
	} else {
		// result := database.Db.Delete(&ErrorData, mail)
		return ErrorData, result.Error

	}
}
