package controller

import (
	"be-go-car-rental/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type customersInput struct {
	Name        string `json:"name" binding:"required"`
	NIK         string `json:"nik" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

// GetAllCustomers godoc
// @Summary Get All Customers.
// @Description Get a list of customers.
// @Tags Customers
// @Produce json
// @Success 200 {object} []models.Customers
// @Router /customers [get]
func GetAllCustomers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var customers []models.Customers
	if err := db.Find(&customers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customers)
}

// CreateCustomers godoc
// @Summary Create New Customers.
// @Description Create a new customers.
// @Tags Customers
// @Param Body body customersInput true "The body to create a new Customers"
// @Produce json
// @Success 200 {object} models.Customers
// @Router /customers [post]
func CreateCustomers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var customers models.Customers

	var input customersInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customers = models.Customers{
		Name:        input.Name,
		NIK:         input.NIK,
		PhoneNumber: input.PhoneNumber,
	}

	if err := db.Create(&customers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, customers)
}

// UpdateCustomers godoc
// @Summary Update Customers.
// @Description Update customers by id.
// @Tags Customers
// @Param id path string true "Customers ID"
// @Param Body body customersInput true "The body to update an Customers"
// @Produce json
// @Success 200 {object} models.Customers
// @Router /customers/{id} [patch]
func UpdateCustomers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var customers models.Customers
	var input customersInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.First(&customers, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customers not found"})
		return
	}

	customers.Name = input.Name
	customers.NIK = input.NIK
	customers.PhoneNumber = input.PhoneNumber

	if err := db.Save(&customers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customers)
}

// DeleteCustomers godoc
// @Summary Delete one customers.
// @Description Delete a customers by id.
// @Tags Customers
// @Param id path string true "Customers ID"
// @Produce json
// @Success 200 {object} map[string]boolean
// @Router /customers/{id} [delete]
func DeleteCustomers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	if err := db.Delete(&models.Customers{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customers deleted succesfully"})
}
