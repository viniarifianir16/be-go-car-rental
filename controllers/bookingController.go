package controllers

import (
	"be-go-car-rental/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type bookingInput struct {
	CustomerID uint      `json:"customer_id" binding:"required"`
	CarsID     uint      `json:"cars_id" binding:"required"`
	StartRent  time.Time `json:"start_rent" binding:"required"`
	EndRent    time.Time `json:"end_rent" binding:"required"`
	TotalCost  uint      `json:"total_cost" binding:"required"`
	Finished   bool      `json:"finished,omitempty"`
}

// GetAllBooking godoc
// @Summary Get All Booking.
// @Description Get a list of booking.
// @Tags Booking
// @Produce json
// @Success 200 {object} []models.Booking
// @Router /booking [get]
func GetAllBooking(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var booking []models.Booking

	if err := db.Preload("Customer").Preload("Cars").Find(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, booking)
}

// CreateBooking godoc
// @Summary Create New Booking.
// @Description Create a new booking.
// @Tags Booking
// @Param Body body bookingInput true "The body to create a new Booking"
// @Produce json
// @Success 200 {object} models.Booking
// @Router /booking [post]
func CreateBooking(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var booking models.Booking
	var customer models.Customer
	var cars models.Cars

	var input bookingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// cek id
	if err := db.Where("id = ?", input.CustomerID).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "customer_id not found!"})
		return
	}

	if err := db.Where("id = ?", input.CarsID).First(&cars).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cars_id not found!"})
		return
	}

	booking = models.Booking{
		CustomerID: input.CustomerID,
		CarsID:     input.CarsID,
		StartRent:  input.StartRent,
		EndRent:    input.EndRent,
		TotalCost:  input.TotalCost,
		Finished:   input.Finished,
	}

	if err := db.Create(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, booking)
}

// UpdateBooking godoc
// @Summary Update Booking.
// @Description Update booking by id.
// @Tags Booking
// @Param id path string true "Booking ID"
// @Param Body body bookingInput true "The body to update an Booking"
// @Produce json
// @Success 200 {object} models.Booking
// @Router /booking/{id} [patch]
func UpdateBooking(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var booking models.Booking
	var input bookingInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.First(&booking, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	booking.CustomerID = input.CustomerID
	booking.CarsID = input.CarsID
	booking.StartRent = input.StartRent
	booking.EndRent = input.EndRent
	booking.TotalCost = input.TotalCost
	booking.Finished = input.Finished

	if err := db.Save(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, booking)
}

// DeleteBooking godoc
// @Summary Delete one booking.
// @Description Delete a booking by id.
// @Tags Booking
// @Param id path string true "Booking ID"
// @Produce json
// @Success 200 {object} map[string]boolean
// @Router /booking/{id} [delete]
func DeleteBooking(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	if err := db.Delete(&models.Booking{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking deleted succesfully"})
}
