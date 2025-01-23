package seeders

import (
	"be-go-car-rental/models"
	"log"

	"gorm.io/gorm"
)

func SeedDriver(db *gorm.DB) {
	driver := []models.Driver{
		{Name: "Stanley Baxter", NIK: 3220132938273, PhoneNumber: "081992048712", DailyCost: 150000},
		{Name: "Halsey Quinn", NIK: 3220132938293, PhoneNumber: "081992048713", DailyCost: 135000},
		{Name: "Kingsley Alvarez", NIK: 3220132938313, PhoneNumber: "081992048714", DailyCost: 150000},
		{Name: "Cecilia Flowers", NIK: 3220132938330, PhoneNumber: "081992048715", DailyCost: 155000},
		{Name: "Clarissa Brown", NIK: 3220132938351, PhoneNumber: "081992048716", DailyCost: 145000},
		{Name: "Zeph Larson", NIK: 3220132938372, PhoneNumber: "081992048717", DailyCost: 130000},
		{Name: "Zach Reynolds", NIK: 3220132938375, PhoneNumber: "081992048718", DailyCost: 140000},
		{Name: "Zach Reynolds", NIK: 3220132938375, PhoneNumber: "081992048718", DailyCost: 0},
	}

	for _, driver := range driver {
		var existingDriver models.Driver
		if err := db.Where("nik = ?", driver.NIK).First(&existingDriver).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&driver).Error; err != nil {
					log.Printf("Error adding driver: %v", err)
				}
			} else {
				log.Printf("Error checking driver: %v", err)
			}
		}
	}
}
