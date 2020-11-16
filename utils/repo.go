package utils

import (
	b "github.com/CkCreative/rest/book/models"
	p "github.com/CkCreative/rest/product/models"
	u "github.com/CkCreative/rest/user/models"
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

	database.AutoMigrate(&b.Book{})
	database.AutoMigrate(&u.User{})
	database.AutoMigrate(&p.Product{})
	database.AutoMigrate(&p.Customer{})
	database.AutoMigrate(&p.Address{})
	database.AutoMigrate(&p.Inventory{})
	database.AutoMigrate(&p.CustomerOrder{})
	database.AutoMigrate(&p.CustomerAddress{})
	database.AutoMigrate(&p.CustomerOrderDelivery{})

	DB = database
}
