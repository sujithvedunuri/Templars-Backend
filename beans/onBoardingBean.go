package beans

import "gorm.io/gorm"

type OnBoardingEmployees struct {
	OBID         int            `json:"obid" gorm:"primary_key"`
	Email        string         `json:"email"`
	FirstName    string         `json:"first_name"`
	LastName     string         `json:"last_name"`
	Location     string         `json:"location"`
	Department   string         `json:"department"`
	Role         string         `json:"role"`
	ManagerName  string         `json:"manager_name"`
	HRName       string         `json:"hr_name"`
	AssigneeID   int            `json:"assignee_id"`
	AssigneeName string         `json:"assignee_name"`
	IssuerID     int            `json:"issuer_id"`
	IssuerName   string         `json:"issuer_name"`
	PhoneNumber  int            `json:"phone_number"`
	Region       string         `json:"region"`
	CreatedAt    int64          `gorm:"autoCreateTime"`
	UpdatedAt    int64          `gorm:"autoUpdateTime:milli"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Status       string         `json:"status"`
}

type EmployeeInfo struct {
	EID       int    `json:"eid"`
	FirstName string `json:"first_name"`
}

type OBID struct {
	ObID int `json:"obid"`
}
