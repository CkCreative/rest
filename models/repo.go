package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" //connector
)

// DB variable
var DB *gorm.DB

//ConnectDatabase initial
func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Book{})
	database.AutoMigrate(&User{})

	DB = database
}
