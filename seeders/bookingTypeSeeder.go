package seeders

import (
	"be-go-car-rental/models"

	"gorm.io/gorm"
)

func SeedBookingType(db *gorm.DB) {
	bookingType := []models.BookingType{
		{BookingType: "Car Only", Description: "Rent Car only"},
		{BookingType: "Car & Driver", Description: "Rent Car and a Driver"},
	}

	for _, bookingType := range bookingType {
		db.Create(&bookingType)
	}
}
