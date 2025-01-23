package controllers

import (
	"be-go-car-rental/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type membershipInput struct {
	MembershipName string `json:"membership_name" binding:"required"`
	Discount       int    `json:"discount" binding:"required"`
}

// GetAllMembership godoc
// @Summary Get All Membership.
// @Description Get a list of Membership.
// @Tags Membership
// @Produce json
// @Success 200 {object} []models.Membership
// @Router /membership [get]
func GetAllMembership(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var membership []models.Membership
	if err := db.Find(&membership).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, membership)
}

// GetMembershipById godoc
// @Summary Get Membership by ID.
// @Description Get a Membership by ID.
// @Tags Membership
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Membership
// @Router /membership/{id} [get]
func GetMembershipByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var membership models.Membership

	if err := db.Where("id = ?", c.Param("id")).First(&membership).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, membership)
}

// CreateMembership godoc
// @Summary Create New Membership.
// @Description Create a new Membership.
// @Tags Membership
// @Param Body body membershipInput true "The body to create a new Membership"
// @Produce json
// @Success 200 {object} models.Membership
// @Router /membership [post]
func CreateMembership(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var membership models.Membership

	var input membershipInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	membership = models.Membership{
		MembershipName: input.MembershipName,
		Discount:       input.Discount,
	}

	if err := db.Create(&membership).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, membership)
}

// UpdateMembership godoc
// @Summary Update Membership.
// @Description Update Membership by id.
// @Tags Membership
// @Param id path string true "Membership ID"
// @Param Body body membershipInput true "The body to update an Membership"
// @Produce json
// @Success 200 {object} models.Membership
// @Router /membership/{id} [patch]
func UpdateMembership(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var membership models.Membership
	var input membershipInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.First(&membership, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Membership not found"})
		return
	}

	membership.MembershipName = input.MembershipName
	membership.Discount = input.Discount

	if err := db.Save(&membership).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, membership)
}

// DeleteMembership godoc
// @Summary Delete one Membership.
// @Description Delete a Membership by id.
// @Tags Membership
// @Param id path string true "Membership ID"
// @Produce json
// @Success 200 {object} map[string]boolean
// @Router /membership/{id} [delete]
func DeleteMembership(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	if err := db.Delete(&models.Membership{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Membership deleted succesfully"})
}
