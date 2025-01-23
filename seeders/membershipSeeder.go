package seeders

import (
	"be-go-car-rental/models"

	"gorm.io/gorm"
)

func SeedMembership(db *gorm.DB) {
	membership := []models.Membership{
		{MembershipName: "Bronze", Discount: 4},
		{MembershipName: "Silver", Discount: 7},
		{MembershipName: "Gold", Discount: 15},
	}

	for _, membership := range membership {
		db.Create(&membership)
	}
}
