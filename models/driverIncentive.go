package models

import "time"

type DriverIncentive struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncremet"`
	BookingID uint      `json:"booking_id" binding:"required"`
	Incentive uint      `json:"incentive" binding:"required"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Booking   Booking   `json:"booking" gorm:"foreignKey:BookingID"`
}
