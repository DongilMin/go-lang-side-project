package models

import (
	"gorm.io/gorm"
	"time"
)

type Category struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	Products  []Product      `json:"products,omitempty" gorm:"foreignKey:CategoryID"`
}
