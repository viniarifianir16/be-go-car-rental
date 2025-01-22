package seeders

import (
	"be-go-car-rental/models"
	"log"

	"gorm.io/gorm"
)

func SeedCustomer(db *gorm.DB) {
	customers := []models.Customer{
		{Name: "Wawan Hermawan", NIK: "3372093912739", PhoneNumber: "081237123682"},
		{Name: "Philip Walker", NIK: "3372093912785", PhoneNumber: "081237123683"},
		{Name: "Hugo Fleming", NIK: "3372093912800", PhoneNumber: "081237123684"},
		{Name: "Maximillian Mendez", NIK: "3372093912848", PhoneNumber: "081237123685"},
		{Name: "Felix Dixon", NIK: "3372093912851", PhoneNumber: "081237123686"},
		{Name: "Nicholas Riddle", NIK: "3372093912929", PhoneNumber: "081237123687"},
		{Name: "Stephen Wheeler", NIK: "3372093912976", PhoneNumber: "081237123688"},
		{Name: "Roy Brennan", NIK: "3372093913022", PhoneNumber: "081237123689"},
		{Name: "Eliza Le", NIK: "3372093913106", PhoneNumber: "081237123690"},
		{Name: "Jesse Taylor", NIK: "3372093913126", PhoneNumber: "081237123691"},
		{Name: "Damien Kaufman", NIK: "3372093913202", PhoneNumber: "081237123692"},
		{Name: "Ayesha Richardson", NIK: "3372093913257", PhoneNumber: "081237123693"},
		{Name: "Margaret Stokes", NIK: "3372093913262", PhoneNumber: "081237123694"},
		{Name: "Sara Livingston", NIK: "3372093913268", PhoneNumber: "081237123695"},
		{Name: "Callie Townsend", NIK: "3372093913281", PhoneNumber: "081237123696"},
		{Name: "Lilly Fischer", NIK: "3372093913325", PhoneNumber: "081237123697"},
		{Name: "Theresa Barton", NIK: "3372093913335", PhoneNumber: "081237123698"},
		{Name: "Mia Curtis", NIK: "3372093913343", PhoneNumber: "081237123699"},
		{Name: "Flora Barlow", NIK: "3372093913400", PhoneNumber: "081237123700"},
		{Name: "Vanessa Patton", NIK: "3372093913434", PhoneNumber: "081237123701"},
	}

	for _, customer := range customers {
		var existingCustomer models.Customer
		if err := db.Where("nik = ?", customer.NIK).First(&existingCustomer).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&customer).Error; err != nil {
					log.Printf("Error adding customer: %v", err)
				}
			} else {
				log.Printf("Error checking customer: %v", err)
			}
		}
	}
}
