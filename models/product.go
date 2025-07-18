package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null;size:255"`
	Description string         `json:"description" gorm:"type:text"`
	Price       float64        `json:"price" gorm:"not null;check:price >= 0"`
	Stock       int            `json:"stock" gorm:"default:0;check:stock >= 0"`
	CategoryID  uint           `json:"category_id" gorm:"not null"`
	Category    Category       `json:"category,omitempty" gorm:"foreignKey:CategoryID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}