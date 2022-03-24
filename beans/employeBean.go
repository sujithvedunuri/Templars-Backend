package beans

import (
	"gorm.io/gorm"
)

type Employee struct {
	EID         int            `json:"eid"`
	Email       string         `json:"email"`
	FirstName   string         `json:"first_name"`
	LastName    string         `json:"last_name"`
	Location    string         `json:"location"`
	UserName    string         `json:"username"`
	Password    string         `json:"-"`
	Department  string         `json:"department"`
	Role        string         `json:"role"`
	ManagerName string         `json:"manager_name"`
	HRName      string         `json:"hr_name"`
	PhoneNumber int            `json:"phone_number"`
	Region      string         `json:"region"`
	CreatedAt   int64          `gorm:"autoCreateTime"`
	UpdatedAt   int64          `gorm:"autoUpdateTime:milli"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type LoginBean struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
