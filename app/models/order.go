package models

import (
	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	OrderID   uint    `json:"order_id"`
	MenuItemID uint   `json:"menu_item_id"`
	MenuItem  MenuItem `json:"menu_item" gorm:"foreignKey:MenuItemID"`
	Quantity  int     `json:"quantity"`
	Subtotal  float64 `json:"subtotal"`
}

type Order struct {
	gorm.Model
	UserID  uint        `json:"user_id"`
	User    User        `json:"user" gorm:"foreignKey:UserID"`
	TableID int         `json:"table_id"`
	Items   []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
	Total   float64     `json:"total"`
	Status  string      `json:"status"`
}
