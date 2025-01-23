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
// @Description Get a list of Booking.
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

// GetBookingById godoc
// @Summary Get Booking by ID.
// @Description Get a Booking by ID.
// @Tags Booking
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Booking
// @Router /booking/{id} [get]
func GetBookingbyID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var booking models.Booking

	if err := db.Preload("Customer").Preload("Cars").Where("id = ?", c.Param("id")).First(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, booking)
}

// GetAllBookingWithDiscount godoc
// @Summary Get All Booking With Discount.
// @Description Get a list of Booking with Discount.
// @Tags Booking
// @Produce json
// @Success 200 {object} []models.Booking
// @Router /booking/discount [get]
func GetAllBookingWithDiscount(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var booking []models.Booking

	if err := db.Preload("Customer.Membership").Preload("Cars").Where("id = ?", c.Param("id")).First(&booking).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	for i := range booking {
		daysOfRent := int(booking[i].EndRent.Truncate(24*time.Hour).Sub(booking[i].StartRent.Truncate(24*time.Hour)).Hours()/24) + 1 // Jumlah hari sewa
		dailyCarRent := int(booking[i].Cars.DailyRent)
		membershipDiscount := int(booking[i].Customer.Membership.Discount) / 100

		// Hitung diskon (jika ada membership)
		var discount int
		if booking[i].Customer.MembershipID != 0 {
			// Discount = ( Days_of_Rent * Daily_car_Rent ) * Membership_discount
			discount = int(daysOfRent*dailyCarRent) * membershipDiscount
		}

		// Total biaya tanpa diskon
		//  days * daily_rent
		totalCost := daysOfRent * dailyCarRent
		booking[i].TotalCost = uint(totalCost)

		// Total biaya dengan diskon
		// total cost * membership disc
		totalCostDiscount := totalCost * membershipDiscount
		booking[i].TotalCost = uint(totalCostDiscount)

		// Return data booking dengan total biaya dan diskon
		c.JSON(http.StatusOK, gin.H{
			"id":                  booking[i].ID,
			"finished":            booking[i].Finished,
			"start_rent":          booking[i].StartRent,
			"end_rent":            booking[i].EndRent,
			"days_of_rent":        daysOfRent,
			"daily_rent":          dailyCarRent,
			"discount":            discount,
			"total_cost":          totalCost,
			"total_cost_discount": totalCostDiscount,
			"customer_name":       booking[i].Customer.Name,
			"membership_name":     booking[i].Customer.Membership.MembershipName,
			"membershipDiscount":  booking[i].Customer.Membership.Discount,
			"car":                 booking[i].Cars.Name,
			"car_daily_rent":      booking[i].Cars.DailyRent,
		})
	}
}

// GetBookingByIdWithDiscount godoc
// @Summary Get Booking by ID with Discount.
// @Description Get a Booking by ID with Discount.
// @Tags Booking
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Booking
// @Router /booking/{id}/discount [get]
func GetBookingbyIDWithDiscount(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var booking models.Booking

	if err := db.Preload("Customer.Membership").Preload("Cars").Where("id = ?", c.Param("id")).First(&booking).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	daysOfRent := int(booking.EndRent.Truncate(24*time.Hour).Sub(booking.StartRent.Truncate(24*time.Hour)).Hours()/24) + 1 // Jumlah hari sewa
	dailyCarRent := int(booking.Cars.DailyRent)
	membershipDiscount := int(booking.Customer.Membership.Discount) / 100

	// Hitung diskon (jika ada membership)
	var discount int
	if booking.Customer.MembershipID != 0 {
		// Discount = ( Days_of_Rent * Daily_car_Rent ) * Membership_discount
		discount = int(daysOfRent*dailyCarRent) * membershipDiscount
	}

	// Total biaya tanpa diskon
	//  days * daily_rent
	totalCost := daysOfRent * dailyCarRent
	booking.TotalCost = uint(totalCost)

	// Total biaya dengan diskon
	// total cost * membership disc
	totalCostDiscount := totalCost * membershipDiscount
	booking.TotalCost = uint(totalCostDiscount)

	// Return data booking dengan total biaya dan diskon
	c.JSON(http.StatusOK, gin.H{
		"id":                  booking.ID,
		"finished":            booking.Finished,
		"start_rent":          booking.StartRent,
		"end_rent":            booking.EndRent,
		"days_of_rent":        daysOfRent,
		"daily_rent":          dailyCarRent,
		"discount":            discount,
		"total_cost":          totalCost,
		"total_cost_discount": totalCostDiscount,
		"customer_name":       booking.Customer.Name,
		"membership_name":     booking.Customer.Membership.MembershipName,
		"membershipDiscount":  booking.Customer.Membership.Discount,
		"car":                 booking.Cars.Name,
		"car_daily_rent":      booking.Cars.DailyRent,
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
