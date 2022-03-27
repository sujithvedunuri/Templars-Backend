package daos

import (
	"fmt"
	"log"

	"templars.com/backend/beans"
	"templars.com/backend/resources/database"
)

func RetreiveOnBoardingEmployees(emplInfo beans.EmployeeInfo) ([]beans.OnBoardingEmployees, error) {

	var onboardingInfo []beans.OnBoardingEmployees

	err := database.Db.Where("assignee_id = ? ", emplInfo.EID).Where("assignee_name = ? ", emplInfo.FirstName).Find(&onboardingInfo)
	if err.Error != nil {
		log.Fatal("error retreiving data from OnBoarding table")
		fmt.Println(err.Error)
	}
	return onboardingInfo, err.Error

}

func InsertNewEmployees(ApprovalList []beans.OnBoardingEmployees) error {

	err := database.Db.Create(&ApprovalList)
	if err.Error != nil {
		fmt.Println("error creating new rows to OnBoarding Table ")
		log.Fatalln(err.Error)
	}

	return err.Error
}

func UpdateStatus(ID int) error {

	err := database.Db.Model(&beans.OnBoardingEmployees{}).Where("ob_id = ?", ID).Update("status", "INPROGRESS")
	if err.Error != nil {
		fmt.Println("error updating status to inprogress in OnBoarding Table ")
		log.Fatalln(err.Error)
	}

	return err.Error
}
