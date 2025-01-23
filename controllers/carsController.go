package controllers

import (
	"be-go-car-rental/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type carsInput struct {
	Name      string `json:"name" binding:"required"`
	Stock     uint   `json:"stock" binding:"required"`
	DailyRent uint   `json:"daily_rent" binding:"required"`
}

// GetAllCars godoc
// @Summary Get All Cars.
// @Description Get a list of Cars.
// @Tags Cars
// @Produce json
// @Success 200 {object} []models.Cars
// @Router /cars [get]
func GetAllCars(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var cars []models.Cars

	if err := db.Find(&cars).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cars})
}

// GetCarsById godoc
// @Summary Get Cars by ID.
// @Description Get a Cars by ID.
// @Tags Cars
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Cars
// @Router /cars/{id} [get]
func GetCarsByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var cars models.Cars

	if err := db.Where("id = ?", c.Param("id")).First(&cars).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cars})
}

// CreateCars godoc
// @Summary Create New Cars.
// @Description Create a new Cars.
// @Tags Cars
// @Param Body body carsInput true "The body to create a new Cars"
// @Produce json
// @Success 200 {object} models.Cars
// @Router /cars [post]
func CreateCars(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var cars models.Cars

	var input carsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cars = models.Cars{
		Name:      input.Name,
		DailyRent: input.DailyRent,
	}

	if err := db.Create(&cars).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": cars})
}

// UpdateCars godoc
// @Summary Update Cars.
// @Description Update Cars by id.
// @Tags Cars
// @Param id path string true "Cars ID"
// @Param Body body carsInput true "The body to update an Cars"
// @Produce json
// @Success 200 {object} models.Cars
// @Router /cars/{id} [patch]
func UpdateCars(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var cars models.Cars
	var input carsInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.First(&cars, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cars not found"})
		return
	}

	cars.Name = input.Name
	cars.Stock = input.Stock
	cars.DailyRent = input.DailyRent

	if err := db.Save(&cars).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cars})
}

// DeleteCars godoc
// @Summary Delete one cars.
// @Description Delete a Cars by id.
// @Tags Cars
// @Param id path string true "Cars ID"
// @Produce json
// @Success 200 {object} map[string]boolean
// @Router /cars/{id} [delete]
func DeleteCars(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	if err := db.Delete(&models.Cars{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cars deleted succesfully"})
}
