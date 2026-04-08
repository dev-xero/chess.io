package models

import (
	"time"

	"gorm.io/gorm"
)

// Country record model.
type Country struct {
	Id        uint           `gorm:"primaryKey" json:"country_id"`
	Name      string         `json:"country_name"`
	Code      string         `gorm:"uniqueIndex" json:"country_code"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Country) TableName() string {
	return "countries"
}
