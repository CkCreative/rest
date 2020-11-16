package models

import "github.com/jinzhu/gorm"

// User model
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"type:varchar(100);unique_index"`
	Gender   string `json:"gender"`
	Password string `json:"password"`
	Code     string `json:"code"`
}
