package config

import (
	"be-go-car-rental/models"
	"be-go-car-rental/seeders"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase() *gorm.DB {
	// if error run "go clean -cache"
	dbProvider := GetEnvOrDefault("DB_PROVIDER", "mysql")

	var db *gorm.DB

	if dbProvider == "postgres" {
		username := os.Getenv("DB_USERNAME")
		password := os.Getenv("DB_PASSWORD")
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		database := os.Getenv("DB_NAME")

		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require stmtcache.mode=describe",
			host, username, password, database, port)
		dbGorm, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			PrepareStmt: false,
			Logger:      logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			panic(err.Error())
		}

		db = dbGorm

	} else {
		username := GetEnvOrDefault("DB_USERNAME", "root")
		password := GetEnvOrDefault("DB_PASSWORD", "")
		host := GetEnvOrDefault("DB_HOST", "127.0.0.1")
		port := GetEnvOrDefault("DB_PORT", "3306")
		database := GetEnvOrDefault("DB_NAME", "db_car_rental")

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			username, password, host, port, database)
		dbGorm, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			PrepareStmt: false,
			Logger:      logger.Default.LogMode(logger.Silent),
		})
		log.Printf("Connecting to database with: host=%s port=%s user=%s dbname=%s", host, port, username, database)

		if err != nil {
			panic(err.Error())
		}

		db = dbGorm
	}

	environment := os.Getenv("ENVIRONMENT")
	if environment == "local" {
		err := db.AutoMigrate(
			&models.Customer{},
			&models.Cars{},
			&models.Booking{},
			&models.Membership{},
			&models.BookingType{},
			&models.Driver{},
			&models.DriverIncentive{},
		)
		if err != nil {
			log.Fatal("Failed to migrate database:", err)
		}

		log.Println("Running seeders in local environment...")

		var customerCount int64
		db.Model(&models.Customer{}).Count(&customerCount)
		if customerCount == 0 {
			seeders.SeedCustomer(db)
		}

		var carCount int64
		db.Model(&models.Cars{}).Count(&carCount)
		if carCount == 0 {
			seeders.SeedCars(db)
		}

		var membershipCount int64
		db.Model(&models.Membership{}).Count(&membershipCount)
		if membershipCount == 0 {
			seeders.SeedMembership(db)
		}

		var bookingTypeCount int64
		db.Model(&models.BookingType{}).Count(&bookingTypeCount)
		if bookingTypeCount == 0 {
			seeders.SeedBookingType(db)
		}

		var driver int64
		db.Model(&models.Driver{}).Count(&driver)
		if driver == 0 {
			seeders.SeedDriver(db)
		}
	} else {
		log.Println("Skipping seeders in production environment.")
	}

	return db
}

func GetEnvOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
