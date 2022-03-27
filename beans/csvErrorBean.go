package beans

import "gorm.io/gorm"

type CsvErrorBean struct {
	ID           int            `gorm:"primaryKey"`
	Email        string         `json:"email"`
	FirstName    string         `json:"first_name"`
	LastName     string         `json:"last_name"`
	Location     string         `json:"location"`
	Department   string         `json:"department"`
	Role         string         `json:"role"`
	ManagerName  string         `json:"manager_name"`
	HRName       string         `json:"hr_name"`
	PhoneNumber  int            `json:"phone_number"`
	AssigneeID   int            `json:"assignee_id"`
	AssigneeName string         `json:"assignee_name"`
	IssuerID     int            `json:"issuer_id"`
	IssuerName   string         `json:"issuer_name"`
	Region       string         `json:"region"`
	CreatedAt    int64          `gorm:"autoCreateTime"`
	UpdatedAt    int64          `gorm:"autoUpdateTime:milli"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
