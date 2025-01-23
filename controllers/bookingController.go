package controllers

import (
	"be-go-car-rental/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type bookingInput struct {
	CustomerID    uint      `json:"customer_id" binding:"required"`
	CarsID        uint      `json:"cars_id" binding:"required"`
	BookingTypeID uint      `json:"booking_type_id" binding:"required"`
	DriverID      uint      `json:"driver_id"`
	StartRent     time.Time `json:"start_rent" binding:"required"`
	EndRent       time.Time `json:"end_rent" binding:"required"`
	Finished      bool      `json:"finished,omitempty"`
}

// GetAllBooking godoc
// @Summary Get All Booking.
// @Description Get a list of Booking.
// @Tags Booking
// @Produce json
// @Success 200 {object} []models.Booking
// @Router /booking [get]
func GetAllBooking(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var booking []models.Booking

	if err := db.Find(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, booking)
}

// GetBookingById godoc
// @Summary Get Booking by ID.
// @Description Get a Booking by ID.
// @Tags Booking
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Booking
// @Router /booking/{id} [get]
func GetBookingByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var booking models.Booking

	if err := db.Where("id = ?", c.Param("id")).First(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, booking)
}

// GetBookingByIdWithDetail godoc
// @Summary Get Booking by ID with Detail.
// @Description Get a Booking by ID with Detail.
// @Tags Booking
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Booking
// @Router /booking/{id}/detail [get]
func GetBookingbyIDWithDetail(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var booking models.Booking

	if err := db.Preload("Customer").Preload("Customer.Membership").Preload("Cars").Preload("BookingType").Preload("Driver").Where("id = ?", c.Param("id")).First(&booking).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	// Jumlah hari sewa
	daysOfRent := int(booking.EndRent.Truncate(24*time.Hour).Sub(booking.StartRent.Truncate(24*time.Hour)).Hours()/24) + 1

	totalCostDiscount := booking.TotalCost * uint(booking.Customer.Membership.Discount) / 100

	// Return data booking dengan total biaya dan diskon
	c.JSON(http.StatusOK, gin.H{
		"id":                  booking.ID,
		"start_rent":          booking.StartRent,
		"end_rent":            booking.EndRent,
		"days_of_rent":        daysOfRent,
		"daily_rent":          booking.Cars.DailyRent,
		"discount":            booking.Discount,
		"total_cost":          booking.TotalCost,
		"total_cost_discount": totalCostDiscount, // total harga setelah discount
		"customer_name":       booking.Customer.Name,
		"membership_name":     booking.Customer.Membership.MembershipName,
		"membership_discount": booking.Customer.Membership.Discount,
		"car":                 booking.Cars.Name,
		"car_daily_rent":      booking.Cars.DailyRent,
		"booking_type":        booking.BookingType.Description,
		"driver_name":         booking.Driver.Name,
		"total_driver_cost":   booking.TotalDriverCost,
	})
}

// CreateBooking godoc
// @Summary Create New Booking.
// @Description Create a new Booking.
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
	var bookingType models.BookingType
	var driver models.Driver

	var input bookingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cek id
	if err := db.Where("id = ?", input.CustomerID).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "customer_id not found!"})
		return
	}

	if err := db.Where("id = ?", input.CarsID).First(&cars).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cars_id not found!"})
		return
	}

	if err := db.Where("id = ?", input.BookingTypeID).First(&bookingType).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "booking_type_id not found!"})
		return
	}

	if err := db.Where("id = ?", input.DriverID).First(&driver).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "driver_id not found!"})
		return
	}

	// Jumlah hari sewa
	daysOfRent := int(booking.EndRent.Truncate(24*time.Hour).Sub(booking.StartRent.Truncate(24*time.Hour)).Hours()/24) + 1

	dailyCarRent := int(booking.Cars.DailyRent)
	membershipDiscount := int(booking.Customer.Membership.Discount) / 100

	// Hitung diskon (jika ada membership)
	var discount int
	if booking.Customer.MembershipID != 0 {
		// Discount = ( Days_of_Rent * Daily_car_Rent ) * Membership_discount
		discount = int(daysOfRent*dailyCarRent) * membershipDiscount
	} else {
		discount = 1
	}

	// Total biaya tanpa diskon days * daily_rent
	totalCost := daysOfRent * dailyCarRent

	// Total biaya dengan diskon total cost * membership disc
	totalCostDiscount := totalCost * membershipDiscount

	// Days_of_Rent * harga daily cost driver
	driverCost := booking.Driver.DailyCost
	totalDriverCost := daysOfRent * int(driverCost)

	booking = models.Booking{
		CustomerID:      input.CustomerID,
		CarsID:          input.CarsID,
		BookingtypeID:   input.BookingTypeID,
		DriverID:        input.DriverID,
		StartRent:       input.StartRent,
		EndRent:         input.EndRent,
		TotalCost:       uint(totalCostDiscount),
		Finished:        input.Finished,
		Discount:        uint(discount),
		TotalDriverCost: uint(totalDriverCost),
	}

	if err := db.Create(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, booking)
}

// UpdateBooking godoc
// @Summary Update Booking.
// @Description Update Booking by id.
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

	// Jumlah hari sewa
	daysOfRent := int(booking.EndRent.Truncate(24*time.Hour).Sub(booking.StartRent.Truncate(24*time.Hour)).Hours()/24) + 1

	dailyCarRent := int(booking.Cars.DailyRent)
	membershipDiscount := int(booking.Customer.Membership.Discount) / 100

	// Hitung diskon (jika ada membership)
	var discount int
	if booking.Customer.MembershipID != 0 {
		// Discount = ( Days_of_Rent * Daily_car_Rent ) * Membership_discount
		discount = int(daysOfRent*dailyCarRent) * membershipDiscount
	} else {
		discount = 1
	}

	// Total biaya tanpa diskon days * daily_rent
	totalCost := daysOfRent * dailyCarRent

	// Total biaya dengan diskon total cost * membership disc
	totalCostDiscount := totalCost * membershipDiscount

	// Days_of_Rent * harga daily cost driver
	driverCost := booking.Driver.DailyCost
	totalDriverCost := daysOfRent * int(driverCost)

	booking.CustomerID = input.CustomerID
	booking.CarsID = input.CarsID
	booking.BookingtypeID = input.BookingTypeID
	booking.DriverID = input.DriverID
	booking.StartRent = input.StartRent
	booking.EndRent = input.EndRent
	booking.TotalCost = uint(totalCostDiscount)
	booking.Finished = input.Finished
	booking.Discount = uint(discount)
	booking.TotalDriverCost = uint(totalDriverCost)

	if err := db.Save(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, booking)
}

// DeleteBooking godoc
// @Summary Delete one Booking.
// @Description Delete a Booking by id.
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
