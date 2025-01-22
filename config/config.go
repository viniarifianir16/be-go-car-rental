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
	dbProvider := GetEnvOrDefault("DB_PROVIDER", "mysql")

	var db *gorm.DB

	if dbProvider == "postgres" {
		username := os.Getenv("DB_USERNAME")
		password := os.Getenv("DB_PASSWORD")
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		database := os.Getenv("DB_NAME")

		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
			host, username, password, database, port)
		dbGorm, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			PrepareStmt: false,
			Logger:      logger.Default.LogMode(logger.Info),
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
			Logger:      logger.Default.LogMode(logger.Info),
		})
		log.Printf("Connecting to database with: host=%s port=%s user=%s dbname=%s", host, port, username, database)

		if err != nil {
			panic(err.Error())
		}

		db = dbGorm
	}

	// go clean -cache
	if db != nil {
		return db
	}

	err := db.AutoMigrate(
		&models.Customer{},
		&models.Cars{},
		&models.Booking{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	var customerCount int64
	var carCount int64

	db.Model(&models.Customer{}).Count(&customerCount)
	db.Model(&models.Cars{}).Count(&carCount)

	if customerCount == 0 {
		seeders.SeedCustomer(db)
	} else {
		log.Println("Data already exists in the customer table. Seeder will not run.")
	}

	if carCount == 0 {
		seeders.SeedCars(db)
	} else {
		log.Println("Data already exists in the cars table. Seeder will not run.")
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
