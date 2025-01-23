package controllers

import (
	"be-go-car-rental/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type driverIncentiveInput struct {
	BookingID uint `json:"booking_id" binding:"required"`
	Incentive uint `json:"incentive" binding:"required"`
}

// GetAllDriverIncentive godoc
// @Summary Get All Driver Incentive.
// @Description Get a list of Driver Incentive.
// @Tags Driver Incentive
// @Produce json
// @Success 200 {object} []models.DriverIncentive
// @Router /driverincentive [get]
func GetAllDriverIncentive(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var driverIncentive []models.DriverIncentive

	if err := db.Find(&driverIncentive).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": driverIncentive})
}

// GetDriverById godoc
// @Summary Get Driver by ID.
// @Description Get a Driver by ID.
// @Tags Driver Incentive
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Driver
// @Router /driverincentive/{id} [get]
func GetDriverIncentiveByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var driverIncentive models.DriverIncentive

	if err := db.Where("id = ?", c.Param("id")).First(&driverIncentive).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": driverIncentive})
}

// CreateDriverIncentive godoc
// @Summary Create New Driver Incentive.
// @Description Create a new Driver Incentive.
// @Tags Driver Incentive
// @Param Body body driverIncentiveInput true "The body to create a new Driver Incentive"
// @Produce json
// @Success 200 {object} models.DriverIncentive
// @Router /driverincentive [post]
func CreateDriverIncentive(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var driverIncentive models.DriverIncentive
	var booking models.Booking

	var input driverIncentiveInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// cek id
	if err := db.Where("id = ?", input.BookingID).First(&booking).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "booking_id not found!"})
		return
	}

	daysOfRent := int(booking.EndRent.Truncate(24*time.Hour).Sub(booking.StartRent.Truncate(24*time.Hour)).Hours()/24) + 1 // Jumlah hari sewa
	dailyCarRent := int(booking.Cars.DailyRent)

	// Incentive per Booking = ( Days_of_Rent * Daily_car_Rent ) * 5%
	IncentivePerBooking := (daysOfRent * dailyCarRent) * 5 / 100

	driverIncentive = models.DriverIncentive{
		BookingID: input.BookingID,
		Incentive: uint(IncentivePerBooking),
	}

	if err := db.Create(&driverIncentive).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": driverIncentive})
}

// UpdateDriverIncentive godoc
// @Summary Update Driver Incentive.
// @Description Update Driver Incentive by id.
// @Tags Driver Incentive
// @Param id path string true "DriverIncentive ID"
// @Param Body body driverIncentiveInput true "The body to update an Driver Incentive"
// @Produce json
// @Success 200 {object} models.DriverIncentive
// @Router /driverincentive/{id} [patch]
func UpdateDriverIncentive(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var driverIncentive models.DriverIncentive
	var booking models.Booking
	var input driverIncentiveInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.First(&driverIncentive, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Driver Incentive not found"})
		return
	}

	daysOfRent := int(booking.EndRent.Truncate(24*time.Hour).Sub(booking.StartRent.Truncate(24*time.Hour)).Hours()/24) + 1 // Jumlah hari sewa
	dailyCarRent := int(booking.Cars.DailyRent)

	// Incentive per Booking = ( Days_of_Rent * Daily_car_Rent ) * 5%
	IncentivePerBooking := (daysOfRent * dailyCarRent) * 5 / 100

	driverIncentive.BookingID = input.BookingID
	driverIncentive.Incentive = uint(IncentivePerBooking)

	if err := db.Save(&driverIncentive).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": driverIncentive})
}

// DeleteDriverIncentive godoc
// @Summary Delete one Driver Incentive.
// @Description Delete a Driver Incentive by id.
// @Tags Driver Incentive
// @Param id path string true "DriverIncentive ID"
// @Produce json
// @Success 200 {object} map[string]boolean
// @Router /driverincentive/{id} [delete]
func DeleteDriverIncentive(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	if err := db.Delete(&models.DriverIncentive{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Driver Incentive deleted succesfully"})
}
