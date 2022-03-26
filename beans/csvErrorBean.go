package beans

import "gorm.io/gorm"

type CsvErrorBean struct {
	Email       string         `json:"email"`
	FirstName   string         `json:"first_name"`
	LastName    string         `json:"last_name"`
	Location    string         `json:"location"`
	Department  string         `json:"department"`
	Role        string         `json:"role"`
	ManagerName string         `json:"manager_name"`
	HRName      string         `json:"hr_name"`
	PhoneNumber int            `json:"phone_number"`
	Issuer      string         `json:"issuer"`
	Region      string         `json:"region"`
	CreatedAt   int64          `gorm:"autoCreateTime"`
	UpdatedAt   int64          `gorm:"autoUpdateTime:milli"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
