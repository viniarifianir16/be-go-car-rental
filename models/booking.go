package models

import "time"

type Booking struct {
	ID            uint        `json:"id" gorm:"primaryKey;autoIncremet"`
	CustomerID    uint        `json:"customer_id" binding:"required"`
	CarsID        uint        `json:"cars_id" binding:"required"`
	BookingtypeID uint        `json:"booking_type_id,omitempty" binding:"required"`
	DriverID      uint        `json:"driver_id,omitempty" binding:"required"`
	StartRent     time.Time   `json:"start_rent" gorm:"type:TIMESTAMP" binding:"required"`
	EndRent       time.Time   `json:"end_rent" gorm:"type:TIMESTAMP" binding:"required"`
	TotalCost     uint        `json:"total_cost" binding:"required"`
	Finished      bool        `json:"finished" gorm:"type:BOOLEAN"`
	CreatedAt     time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
	Customer      Customer    `json:"customer" gorm:"foreignKey:CustomerID"`
	Cars          Cars        `json:"cars" gorm:"foreignKey:CarsID"`
	BookingType   BookingType `json:"booking_type" gorm:"foreignKey:BookingtypeID"`
}
