package controllers

import (
	"fmt"
	"net/http"

	"github.com/CkCreative/rest/product/models"
	"github.com/CkCreative/rest/utils"
	"github.com/gin-gonic/gin"
)

// FindAddresses : GET /addresses
// Get all address
func FindAddresses(c *gin.Context) {
	var address []models.Address
	utils.DB.Find(&address)

	c.JSON(http.StatusOK, gin.H{"data": address})
}

// CreateAddress POST /addresses
// Create new address
func CreateAddress(c *gin.Context) {
	var input models.Address
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)

		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not decode data"})
		return
	}
	// Create address
	utils.DB.Create(&input)

	c.JSON(http.StatusOK, gin.H{"data": input})
}

// FindAddress GET /addresses/:id
// Find a address
func FindAddress(c *gin.Context) { // Get model if exist
	var address models.Address

	if err := utils.DB.Where("id = ?", c.Param("id")).First(&address).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": address})
}

//UpdateAddress PATCH /addresses/:id
// Update a address
func UpdateAddress(c *gin.Context) {
	// Get model if exist
	var address models.Address
	if err := utils.DB.Where("id = ?", c.Param("id")).First(&address).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.Address
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not decode data"})
		return
	}

	utils.DB.Model(&address).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": address})
}

// DeleteAddress DELETE /addresses/:id
// Delete a address
func DeleteAddress(c *gin.Context) {
	// Get model if exist
	var address models.Address
	if err := utils.DB.Where("id = ?", c.Param("id")).First(&address).Error; err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	utils.DB.Delete(&address, c.Param("id"))

	c.JSON(http.StatusOK, gin.H{"data": true})
}
