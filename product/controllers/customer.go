package controllers

import (
	"fmt"
	"net/http"

	"github.com/CkCreative/rest/product/models"
	"github.com/CkCreative/rest/utils"
	"github.com/gin-gonic/gin"
)

// FindCustomers : GET /customers
// Get all customer
func FindCustomers(c *gin.Context) {
	var customer []models.Customer
	utils.DB.Find(&customer)

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

// CreateCustomer POST /customers
// Create new customer
func CreateCustomer(c *gin.Context) {
	var input models.Customer
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)

		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not decode data"})
		return
	}
	// Create customer
	utils.DB.Create(&input)

	c.JSON(http.StatusOK, gin.H{"data": input})
}

// FindCustomer GET /customers/:id
// Find a customer
func FindCustomer(c *gin.Context) { // Get model if exist
	var customer models.Customer

	if err := utils.DB.Where("id = ?", c.Param("id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

//UpdateCustomer PATCH /customers/:id
// Update a customer
func UpdateCustomer(c *gin.Context) {
	// Get model if exist
	var customer models.Customer
	if err := utils.DB.Where("id = ?", c.Param("id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.Customer
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not decode data"})
		return
	}

	utils.DB.Model(&customer).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

// DeleteCustomer DELETE /customers/:id
// Delete a customer
func DeleteCustomer(c *gin.Context) {
	// Get model if exist
	var customer models.Customer
	if err := utils.DB.Where("id = ?", c.Param("id")).First(&customer).Error; err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	utils.DB.Delete(&customer, c.Param("id"))

	c.JSON(http.StatusOK, gin.H{"data": true})
}
