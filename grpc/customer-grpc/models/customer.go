package models

import (
	// "database/sql"
	"time"

	"github.com/google/uuid"
	// "gorm.io/gorm"
)

type Customer struct {
	Id          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey`
	Name        string    `gorm:"type:varchar(256)"`
	Address     string    `gorm:"type:varchar(256)"`
	LicenseID   string    `gorm:"type:varchar(256)"`
	PhoneNumber string    `gorm:"type:varchar(12)"`
	Email       string    `gorm:"type:varchar(256)"`
	Password    string    `gorm:"type:varchar(256)"`
	Active      bool      `gorm:"default:true"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
