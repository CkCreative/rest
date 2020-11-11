package controllers

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/CkCreative/rest/models"
	"github.com/gin-gonic/gin"
)

// CreateUser function
func CreateUser(c *gin.Context) {
	// CreateUserInput Validate input
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//user := models.User{Name: input.Name, Email: input.Name, Password: input.Password, Gender: input.Gender}

	pass, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.Password = string(pass)

	if err := models.DB.Create(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Entry could not be created"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})

}

// FindUsers : GET /users
// Get all users
func FindUsers(c *gin.Context) {
	//subset of fields
	type User struct {
		Name  string
		Email string
	}
	var users []User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}
