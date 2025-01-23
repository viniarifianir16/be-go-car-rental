package seeders

import (
	"be-go-car-rental/models"

	"gorm.io/gorm"
)

func SeedDriverIncentive(db *gorm.DB) {
	driverIncentive := []models.DriverIncentive{
		{BookingID: 6, Incentive: 40000},
		{BookingID: 7, Incentive: 75000},
		{BookingID: 8, Incentive: 25000},
	}

	for _, driverIncentive := range driverIncentive {
		db.Create(&driverIncentive)
	}
}
