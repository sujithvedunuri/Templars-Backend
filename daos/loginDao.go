package daos

import (
	"fmt"

	"templars.com/backend/beans"
	"templars.com/backend/resources/database"
)

func GetLoginCredentials(data beans.LoginBean) beans.Employee {
	var user beans.Employee
	err := database.Db.Where("user_name = ? ", data.UserName).First(&user)
	if err.Error != nil {
		fmt.Println("error retreiving from OnBoarding table")
	}
	return user
}
