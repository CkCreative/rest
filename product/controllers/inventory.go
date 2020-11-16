package controllers

import (
	"fmt"
	"net/http"

	"github.com/CkCreative/rest/product/models"
	"github.com/CkCreative/rest/utils"
	"github.com/gin-gonic/gin"
)

// FindInventories : GET /inventory
// Get all inventory
func FindInventories(c *gin.Context) {
	var p []models.Inventory
	utils.DB.Find(&p)

	c.JSON(http.StatusOK, gin.H{"data": p})
}

// CreateInventory POST /inventory
// Create new inventory
func CreateInventory(c *gin.Context) {
	var input models.Inventory
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)

		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not decode data"})
		return
	}
	// Create inventory
	utils.DB.Create(&input)

	c.JSON(http.StatusOK, gin.H{"data": input})
}

// FindInventory GET /inventory/:id
// Find a inventory
func FindInventory(c *gin.Context) { // Get model if exist
	var inventory models.Inventory

	if err := utils.DB.Where("id = ?", c.Param("id")).First(&inventory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": inventory})
}

//UpdateInventory PATCH /inventory/:id
// Update a inventory
func UpdateInventory(c *gin.Context) {
	// Get model if exist
	var inventory models.Inventory
	if err := utils.DB.Where("id = ?", c.Param("id")).First(&inventory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.Inventory
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not decode data"})
		return
	}

	utils.DB.Model(&inventory).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": inventory})
}

// DeleteInventory DELETE /inventory/:id
// Delete a inventory
func DeleteInventory(c *gin.Context) {
	// Get model if exist
	var inventory models.Inventory
	if err := utils.DB.Where("id = ?", c.Param("id")).First(&inventory).Error; err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	utils.DB.Delete(&inventory, c.Param("id"))

	c.JSON(http.StatusOK, gin.H{"data": true})
}
