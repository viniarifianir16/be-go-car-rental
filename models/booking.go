package models

import "time"

type Booking struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncremet"`
	CustomerID uint      `json:"customer_id"`
	CarsID     uint      `json:"cars_id"`
	StartRent  time.Time `json:"start_rent"`
	EndRent    time.Time `json:"end_rent"`
	TotalCost  uint      `json:"total_cost"`
	Finished   bool      `json:"finished"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
