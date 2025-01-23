package models

import "time"

type Membership struct {
	ID             uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	MembershipName string    `json:"membership_name" binding:"required"`
	Discount       int       `json:"discount" binding:"required"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
