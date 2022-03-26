package daos

import (
	"fmt"

	"templars.com/backend/beans"
	"templars.com/backend/resources/database"
)

func FetchErrorMail() beans.CsvErrorBean {
	var csvErrorScheet beans.CsvErrorBean
	database.Db.Find(&csvErrorScheet)
	fmt.Println(csvErrorScheet)
	return csvErrorScheet

}
