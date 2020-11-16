package controllers

import (
	"fmt"
	"net/http"

	"github.com/CkCreative/rest/product/models"
	"github.com/CkCreative/rest/utils"
	"github.com/gin-gonic/gin"
)

// FindProducts : GET /product
// Get all product
func FindProducts(c *gin.Context) {
	var p []models.Product
	utils.DB.Find(&p)

	c.JSON(http.StatusOK, gin.H{"data": p})
}

// CreateProduct POST /product
// Create new product
func CreateProduct(c *gin.Context) {
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)

		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not decode data"})
		return
	}
	// Create product
	utils.DB.Create(&input)

	c.JSON(http.StatusOK, gin.H{"data": input})
}

// FindProduct GET /product/:id
// Find a product
func FindProduct(c *gin.Context) { // Get model if exist
	var product models.Product

	if err := utils.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

//UpdateProduct PATCH /product/:id
// Update a product
func UpdateProduct(c *gin.Context) {
	// Get model if exist
	var product models.Product
	if err := utils.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not decode data"})
		return
	}

	utils.DB.Model(&product).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// DeleteProduct DELETE /product/:id
// Delete a product
func DeleteProduct(c *gin.Context) {
	// Get model if exist
	var product models.Product
	if err := utils.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	utils.DB.Delete(&product, c.Param("id"))

	c.JSON(http.StatusOK, gin.H{"data": true})
}
