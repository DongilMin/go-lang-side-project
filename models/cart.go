package models

import "time"

type CartItem struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`
	User      User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
	ProductID uint      `json:"product_id"`
	Product   Product   `json:"product,omitempty" gorm:"foreignKey:ProductID"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
