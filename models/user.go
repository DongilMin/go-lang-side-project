package models

import (
    "time"
    "gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"not null"`
	Email     string         `json:"email" gorm:"unique;not null"`
	Phone     string         `json:"phone"`
	Address   string         `json:"address"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	Orders    []Order        `json:"orders,omitempty" gorm:"foreignKey:UserID"`
	CartItems []CartItem     `json:"cart_items,omitempty" gorm:"foreignKey:UserID"`
}