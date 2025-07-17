package models

import (
    "time"
    "gorm.io/gorm"
)

type Order struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	UserID      uint           `json:"user_id"`
	User        User           `json:"user,omitempty" gorm:"foreignKey:UserID"`
	TotalAmount float64        `json:"total_amount"`
	Status      string         `json:"status" gorm:"default:'pending'"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	OrderItems  []OrderItem    `json:"order_items,omitempty" gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Product   Product `json:"product,omitempty" gorm:"foreignKey:ProductID"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}