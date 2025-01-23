package models

import "time"

type BookingType struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	BookingType string    `json:"booking_type" binding:"required"`
	Description string    `json:"description" binding:"required"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
