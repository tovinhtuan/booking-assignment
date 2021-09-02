package models

import (
	"time"

	"github.com/google/uuid"
)

type Flight struct {
	Id             uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey`
	Name           string     `gorm:"type:varchar(256)"`
	From           string     `gorm:"type:varchar(256)"`
	To             string     `gorm:"type:varchar(256)"`
	Date           time.Time  `gorm:"type:timestamp"`
	Status         string     `gorm:"type:varchar(256)"`
	Avaliable_slot int64      `gorm:"type:bigint"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
