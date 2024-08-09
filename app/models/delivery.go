package models

import "time"

type Delivery struct {
	ID           int       `json:"id"`
	OrderID      int       `json:"order_id"`
	DriverID     int       `json:"driver_id"`
	Status       string    `json:"status"`
	Address      string    `json:"address"`
	EstimatedETA time.Time `json:"estimated_eta"`
	ActualETA    time.Time `json:"actual_eta"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
