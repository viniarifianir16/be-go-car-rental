package controllers

import (
	"be-go-car-rental/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type bookingTypeInput struct {
	BookingType string `json:"booking_type" binding:"required"`
	Description string `json:"description" binding:"required"`
}

// GetAllBookingType godoc
// @Summary Get All Booking Type.
// @Description Get a list of Booking Type.
// @Tags Booking Type
// @Produce json
// @Success 200 {object} []models.BookingType
// @Router /bookingtype [get]
func GetAllBookingType(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var bookingType []models.BookingType
	if err := db.Find(&bookingType).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": bookingType})
}

// GetBookingTypeById godoc
// @Summary Get Booking Type by ID.
// @Description Get a Booking Type by ID.
// @Tags Booking Type
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.BookingType
// @Router /bookingtype/{id} [get]
func GetBookingTypeByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var bookingType models.BookingType

	if err := db.Where("id = ?", c.Param("id")).First(&bookingType).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bookingType})
}

// CreateBookingType godoc
// @Summary Create New Booking Type.
// @Description Create a new Booking Type.
// @Tags Booking Type
// @Param Body body bookingTypeInput true "The body to create a new Booking Type"
// @Produce json
// @Success 200 {object} models.BookingType
// @Router /bookingtype [post]
func CreateBookingType(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var bookingType models.BookingType

	var input bookingTypeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bookingType = models.BookingType{
		BookingType: input.BookingType,
		Description: input.Description,
	}

	if err := db.Create(&bookingType).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": bookingType})
}

// UpdateBookingType godoc
// @Summary Update Booking Type.
// @Description Update Booking Type by id.
// @Tags Booking Type
// @Param id path string true "BookingType ID"
// @Param Body body bookingTypeInput true "The body to update an Booking Type"
// @Produce json
// @Success 200 {object} models.BookingType
// @Router /bookingtype/{id} [patch]
func UpdateBookingType(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var bookingType models.BookingType
	var input bookingTypeInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.First(&bookingType, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "BookingType not found"})
		return
	}

	bookingType.BookingType = input.BookingType
	bookingType.Description = input.Description

	if err := db.Save(&bookingType).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bookingType})
}

// DeleteBookingType godoc
// @Summary Delete one Booking Type.
// @Description Delete a Booking Type by id.
// @Tags Booking Type
// @Param id path string true "BookingType ID"
// @Produce json
// @Success 200 {object} map[string]boolean
// @Router /bookingtype/{id} [delete]
func DeleteBookingType(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	if err := db.Delete(&models.BookingType{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking Type deleted succesfully"})
}
