package controllers

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/CkCreative/rest/user/models"
	"github.com/CkCreative/rest/utils"
	"github.com/gin-gonic/gin"
)

// CreateUser function
func CreateUser(c *gin.Context) {
	// CreateUserInput Validate input
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not decode input"})
		return
	}

	//user := models.User{Name: input.Name, Email: input.Name, Password: input.Password, Gender: input.Gender}

	pass, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password Encryption  failed"})
		return
	}
	input.Password = string(pass)

	if err := utils.DB.Create(&input); err.Error != nil {
		fmt.Println(err.Error)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Entry could not be created"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Entry created successfully"})

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
	utils.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// UpdatePassword function
func UpdatePassword(c *gin.Context) {
	type U struct {
		Password string `json:"password"`
		Email    string `json:"email"`
		Code     string `json:"code"`
	}
	var u models.User
	user := &U{}
	if err := c.ShouldBindJSON(user); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not decode input"})
		return
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password Encryption  failed"})
		return
	}
	user.Password = string(pass)

	utils.DB.Where("email = ?", user.Email).First(&u)

	if user.Code == u.Code {
		if s := utils.DB.Model(&u).Update("password", user.Password); s.Error != nil {
			fmt.Println(s.Error)
			c.JSON(http.StatusBadGateway, gin.H{"data": "Update failed"})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Invalid data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Update successful"})
}

// UpdateUser function
func UpdateUser(c *gin.Context) {
	// Get model if exist
	var user models.User
	if err := utils.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not decode data"})
		return
	}

	utils.DB.Model(&user).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": "Update was successful"})
}
