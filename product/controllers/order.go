package controllers

import (
	"fmt"
	"net/http"

	"github.com/CkCreative/rest/product/models"
	"github.com/CkCreative/rest/utils"
	"github.com/gin-gonic/gin"
)

// FindOrders : GET /orders
// Get all orders
func FindOrders(c *gin.Context) {
	var orders []models.CustomerOrder
	utils.DB.Find(&orders)

	c.JSON(http.StatusOK, gin.H{"data": orders})
}

// CreateOrder POST /orders
// Create new order
func CreateOrder(c *gin.Context) {
	var order models.CustomerOrder
	if err := c.ShouldBindJSON(&order); err != nil {
		fmt.Println(err)

		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not decode data"})
		return
	}
	// Create order
	utils.DB.Create(&order)

	c.JSON(http.StatusOK, gin.H{"data": order})
}

// FindOrder GET /orders/:id
// Find a order
func FindOrder(c *gin.Context) { // Get model if exist
	var order models.CustomerOrder

	if err := utils.DB.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}

//UpdateOrder PATCH /books/:id
// Update a book
func UpdateOrder(c *gin.Context) {
	// Get model if exist
	var order models.CustomerOrder
	if err := utils.DB.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.CustomerOrder
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not decode data"})
		return
	}

	utils.DB.Model(&order).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": order})
}

// DeleteOrder DELETE /orders/:id
// Delete a order
func DeleteOrder(c *gin.Context) {
	// Get model if exist
	var order models.CustomerOrder
	if err := utils.DB.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	utils.DB.Delete(&order, c.Param("id"))

	c.JSON(http.StatusOK, gin.H{"data": true})
}
