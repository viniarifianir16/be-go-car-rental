package models

import "time"

type Driver struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name" binding:"required"`
	NIK         uint      `json:"nik" gorm:"unique" binding:"required"`
	PhoneNumber string    `json:"phone_number" binding:"required"`
	DailyCost   uint      `json:"daily_cost"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
