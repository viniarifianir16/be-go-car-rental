package models

import "time"

type Customer struct {
	ID           uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	MembershipID uint       `json:"membership_id"`
	Name         string     `json:"name" binding:"required"`
	NIK          uint       `json:"nik" gorm:"unique" binding:"required"`
	PhoneNumber  string     `json:"phone_number" binding:"required"`
	CreatedAt    time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	Membership   Membership `json:"membership" gorm:"foreignKey:MembershipID"`
}
