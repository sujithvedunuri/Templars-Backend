package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"templars.com/backend/beans"
	"templars.com/backend/daos"
)

func EmployeeOnBoarding(c *gin.Context) {

	var emplInfo beans.EmployeeInfo

	if err := c.ShouldBindJSON(&emplInfo); err != nil {
		fmt.Println("error binding json to OnBoarding bean")
	}

	onBoardingInfo, err := daos.RetreiveOnBoardingEmployees(emplInfo)
	if err != nil {
		c.IndentedJSON(http.StatusOK, onBoardingInfo)
	}

}

func AddNewEmployee(c *gin.Context) {

	var ApprovalList []beans.OnBoardingEmployees

	if err := c.ShouldBindJSON(&ApprovalList); err != nil {
		fmt.Println("error binding json to onboarding ApprovalList bean")
	}

	err := daos.InsertNewEmployees(ApprovalList)
	if err != nil {
		c.IndentedJSON(http.StatusOK, "Added employees succesfully")
	}
}

func UpdateEmployeeStatus(c *gin.Context) {

	var Id beans.OBID
	if err := c.ShouldBindJSON(&Id); err != nil {
		fmt.Println("error binding json to onboarding OBID bean")
	}
	var ID = Id.ObID
	err := daos.UpdateStatus(ID)
	if err != nil {
		c.IndentedJSON(http.StatusOK, "Updated status succesfully")
	}

}
