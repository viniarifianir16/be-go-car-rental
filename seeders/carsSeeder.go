package seeders

import (
	"be-go-car-rental/models"

	"gorm.io/gorm"
)

func SeedCars(db *gorm.DB) {
	cars := []models.Cars{
		{Name: "Toyota Camry", Stock: 2, DailyRent: 500000},
		{Name: "Toyota Avalon", Stock: 2, DailyRent: 500000},
		{Name: "Toyota Yaris", Stock: 2, DailyRent: 400000},
		{Name: "Toyota Agya", Stock: 2, DailyRent: 400000},
		{Name: "Toyota Fortuner", Stock: 1, DailyRent: 700000},
		{Name: "Toyota Rush", Stock: 1, DailyRent: 600000},
		{Name: "Toyota Hiace", Stock: 1, DailyRent: 1000000},
		{Name: "Honda Brio", Stock: 3, DailyRent: 500000},
		{Name: "Honda Civic", Stock: 1, DailyRent: 500000},
		{Name: "Honda Jazz", Stock: 1, DailyRent: 500000},
		{Name: "Honda Mobilio", Stock: 2, DailyRent: 700000},
		{Name: "Honda Amaze", Stock: 1, DailyRent: 700000},
		{Name: "Honda Breeze", Stock: 2, DailyRent: 700000},
		{Name: "Mitsubishi Pajero Sport", Stock: 5, DailyRent: 800000},
		{Name: "Mitsubishi Mirage", Stock: 3, DailyRent: 600000},
	}

	for _, cars := range cars {
		db.Create(&cars)
	}
}
