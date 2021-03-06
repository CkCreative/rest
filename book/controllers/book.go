package controllers

import (
	"fmt"
	"net/http"

	"github.com/CkCreative/rest/book/models"
	"github.com/CkCreative/rest/utils"
	"github.com/gin-gonic/gin"
)

// CreateBookInput type
type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

// UpdateBookInput type
type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

// FindBooks : GET /books
// Get all books
func FindBooks(c *gin.Context) {
	var books []models.Book
	utils.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// CreateBook POST /books
// Create new book
func CreateBook(c *gin.Context) {
	// CreateBookInput Validate input
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)

		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not decode data"})
		return
	}
	// Create book
	book := models.Book{Title: input.Title, Author: input.Author}
	utils.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// FindBook GET /books/:id
// Find a book
func FindBook(c *gin.Context) { // Get model if exist
	var book models.Book

	if err := utils.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

//UpdateBook PATCH /books/:id
// Update a book
func UpdateBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := utils.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not decode data"})
		return
	}

	utils.DB.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DeleteBook DELETE /books/:id
// Delete a book
func DeleteBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := utils.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	utils.DB.Delete(&book, c.Param("id"))

	c.JSON(http.StatusOK, gin.H{"data": true})
}
