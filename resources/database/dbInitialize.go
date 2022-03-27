package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"templars.com/backend/beans"
)

var Db *gorm.DB

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//"root:password@tcp(127.0.0.1:3306)/TabRemDatabase?charset=utf8mb4&parseTime=True&loc=Local"

	dbformat := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	db, error := gorm.Open(mysql.Open(dbformat), &gorm.Config{})
	Db = db
	if error != nil {
		log.Fatal("unable to connect to the database")
	}
	Db.AutoMigrate(&beans.Employee{})
	Db.AutoMigrate(&beans.OnBoardingEmployees{})

}

func DbTemp() {

	employees := []beans.Employee{
		{
			EID:         1,
			Email:       "sujith@gmail.com",
			FirstName:   "sujith",
			LastName:    "vedunuri",
			Location:    "hyd",
			UserName:    "sujith",
			Password:    "sujith",
			Department:  "IT",
			Role:        "engineer",
			ManagerName: "mani",
			HRName:      "pracheth",
			PhoneNumber: 1234567891,
			Region:      "hyd",
		},
		{
			EID:         2,
			Email:       "rahul@gmail.com",
			FirstName:   "rahul",
			LastName:    "tumukunta",
			Location:    "hyd",
			UserName:    "rahul",
			Password:    "rahul",
			Department:  "IT",
			Role:        "engineer",
			ManagerName: "mani",
			HRName:      "pracheth",
			PhoneNumber: 1234567892,
			Region:      "hyd",
		},
		{
			EID:         3,
			Email:       "rohith@gmail.com",
			FirstName:   "rohith",
			LastName:    "vanteru",
			Location:    "hyd",
			UserName:    "rohith",
			Password:    "rohith",
			Department:  "IT",
			Role:        "engineer",
			ManagerName: "mani",
			HRName:      "pracheth",
			PhoneNumber: 1234567893,
			Region:      "hyd",
		},
		{
			EID:         4,
			Email:       "kakashi@gmail.com",
			FirstName:   "kakashi",
			LastName:    "vanteru",
			Location:    "hyd",
			UserName:    "kakashi",
			Password:    "kakashi",
			Department:  "IT",
			Role:        "engineer",
			ManagerName: "mani",
			HRName:      "pracheth",
			PhoneNumber: 1234567894,
			Region:      "hyd",
		},
	}

	if err := Db.Create(&employees).Error; err != nil {
		fmt.Println("error creating new row to Employee Table ")
		log.Fatalln(err)
	}

	OnBoarders := []beans.OnBoardingEmployees{
		{
			OBID:         1,
			Email:        "test1@gmail.com",
			FirstName:    "test",
			LastName:     "1",
			Location:     "sgp",
			Department:   "Finance",
			Role:         "junior financer",
			ManagerName:  "rahul",
			HRName:       "rohith",
			AssigneeID:   "2",
			AssigneeName: "rahul",
			IssuerID:     "4",
			IssuerName:   "kakashi",
			PhoneNumber:  1234567895,
			Region:       "sgp",
			Status:       "pending",
		},
		{
			OBID:         2,
			Email:        "test2@gmail.com",
			FirstName:    "test",
			LastName:     "2",
			Location:     "thai",
			Department:   "IT",
			Role:         "junior engineer",
			ManagerName:  "rahul",
			HRName:       "sujith",
			AssigneeID:   "1",
			AssigneeName: "sujith",
			IssuerID:     "4",
			IssuerName:   "kakashi",
			PhoneNumber:  1234567896,
			Region:       "thai",
			Status:       "pending",
		},
		{
			OBID:         3,
			Email:        "test3@gmail.com",
			FirstName:    "test",
			LastName:     "3",
			Location:     "viet",
			Department:   "accounting",
			Role:         "junior accounter",
			ManagerName:  "rahul",
			HRName:       "rohith",
			AssigneeID:   "3",
			AssigneeName: "rohith",
			IssuerID:     "4",
			IssuerName:   "kakashi",
			PhoneNumber:  1234567897,
			Region:       "sgp",
			Status:       "pending",
		},
		{
			OBID:         4,
			Email:        "test4@gmail.com",
			FirstName:    "test",
			LastName:     "4",
			Location:     "jpn",
			Department:   "Human Resources",
			Role:         "junior HR",
			ManagerName:  "rahul",
			HRName:       "sujith",
			AssigneeID:   "2",
			AssigneeName: "rahul",
			IssuerID:     "4",
			IssuerName:   "kakashi",
			PhoneNumber:  1234567897,
			Region:       "jpn",
			Status:       "pending",
		},
	}

	if err := Db.Create(&OnBoarders).Error; err != nil {
		fmt.Println("error creating new row to OnBoarding Table ")
		log.Fatalln(err)
	}

}
