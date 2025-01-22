package models

import "time"

type Customer struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name" binding:"required"`
	NIK         string    `json:"nik" gorm:"unique" binding:"required"`
	PhoneNumber string    `json:"phone_number" binding:"required"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
