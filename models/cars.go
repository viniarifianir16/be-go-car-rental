package models

import "time"

type Cars struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" binding:"required"`
	Stock     uint      `json:"stock" binding:"required"`
	DailyRent uint      `json:"daily_rent" binding:"required"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
