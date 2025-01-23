package controllers

import (
	"be-go-car-rental/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type driverInput struct {
	Name        string `json:"name" binding:"required"`
	NIK         uint   `json:"nik" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	DailyCost   uint   `json:"daily_cost" binding:"required"`
}

// GetAllDriver godoc
// @Summary Get All Driver.
// @Description Get a list of Driver.
// @Tags Driver
// @Produce json
// @Success 200 {object} []models.Driver
// @Router /driver [get]
func GetAllDriver(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var driver []models.Driver
	if err := db.Find(&driver).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": driver})
}

// GetDriverById godoc
// @Summary Get Driver by ID.
// @Description Get a Driver by ID.
// @Tags Driver
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Driver
// @Router /driver/{id} [get]
func GetDriverByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var driver models.Driver

	if err := db.Where("id = ?", c.Param("id")).First(&driver).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": driver})
}

// CreateDriver godoc
// @Summary Create New Driver.
// @Description Create a new Driver.
// @Tags Driver
// @Param Body body driverInput true "The body to create a new Driver"
// @Produce json
// @Success 200 {object} models.Driver
// @Router /driver [post]
func CreateDriver(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var driver models.Driver

	var input driverInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	driver = models.Driver{
		Name:        input.Name,
		NIK:         input.NIK,
		PhoneNumber: input.PhoneNumber,
		DailyCost:   input.DailyCost,
	}

	if err := db.Create(&driver).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": driver})
}

// UpdateDriver godoc
// @Summary Update Driver.
// @Description Update Driver by id.
// @Tags Driver
// @Param id path string true "Driver ID"
// @Param Body body driverInput true "The body to update an Driver"
// @Produce json
// @Success 200 {object} models.Driver
// @Router /driver/{id} [patch]
func UpdateDriver(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var driver models.Driver
	var input driverInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.First(&driver, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Driver not found"})
		return
	}

	driver.Name = input.Name
	driver.NIK = input.NIK
	driver.PhoneNumber = input.PhoneNumber
	driver.DailyCost = input.DailyCost

	if err := db.Save(&driver).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": driver})
}

// DeleteDriver godoc
// @Summary Delete one driver.
// @Description Delete a Driver by id.
// @Tags Driver
// @Param id path string true "Driver ID"
// @Produce json
// @Success 200 {object} map[string]boolean
// @Router /driver/{id} [delete]
func DeleteDriver(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	if err := db.Delete(&models.Driver{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Driver deleted succesfully"})
}
