package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)
type Booking struct {
	Id             uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey`
	CustomerID     uuid.UUID  `gorm:"type:uuid;not null"`
	FlightID       uuid.UUID  `gorm:"type:uuid;not null"`
	Code		   int64     `gorm:"type:bigint`
	Status         string	  `gorm:"type:varchar(20)`
	BookedDate     time.Time  `gorm:"type:timestamp"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}