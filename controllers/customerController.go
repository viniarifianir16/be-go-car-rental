package controllers

import (
	"be-go-car-rental/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type customerInput struct {
	Name        string `json:"name" binding:"required"`
	NIK         string `json:"nik" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

// GetAllCustomer godoc
// @Summary Get All Customer.
// @Description Get a list of customer.
// @Tags Customer
// @Produce json
// @Success 200 {object} []models.Customer
// @Router /customer [get]
func GetAllCustomer(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var customer []models.Customer
	if err := db.Find(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customer)
}

// CreateCustomer godoc
// @Summary Create New Customer.
// @Description Create a new customer.
// @Tags Customer
// @Param Body body customerInput true "The body to create a new Customer"
// @Produce json
// @Success 200 {object} models.Customer
// @Router /customer [post]
func CreateCustomer(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var customer models.Customer

	var input customerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer = models.Customer{
		Name:        input.Name,
		NIK:         input.NIK,
		PhoneNumber: input.PhoneNumber,
	}

	if err := db.Create(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, customer)
}

// UpdateCustomer godoc
// @Summary Update Customer.
// @Description Update customer by id.
// @Tags Customer
// @Param id path string true "Customer ID"
// @Param Body body customerInput true "The body to update an Customer"
// @Produce json
// @Success 200 {object} models.Customer
// @Router /customer/{id} [patch]
func UpdateCustomer(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var customer models.Customer
	var input customerInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	customer.Name = input.Name
	customer.NIK = input.NIK
	customer.PhoneNumber = input.PhoneNumber

	if err := db.Save(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customer)
}

// DeleteCustomer godoc
// @Summary Delete one customer.
// @Description Delete a customer by id.
// @Tags Customer
// @Param id path string true "Customer ID"
// @Produce json
// @Success 200 {object} map[string]boolean
// @Router /customer/{id} [delete]
func DeleteCustomer(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	if err := db.Delete(&models.Customer{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted succesfully"})
}
