package models

import (
	"time"

	userModels "github.com/CkCreative/rest/user/models"
	"github.com/jinzhu/gorm"
)

// Product type
type Product struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Size        string `json:"size"`
	Color       string `json:"color"`
}

// Inventory type
type Inventory struct {
	gorm.Model
	Product  Product `json:"product"`
	Quantity string  `json:"quantity"`
}

// Customer type
type Customer struct {
	gorm.Model
	User userModels.User `json:"user"`
}

// Address type
type Address struct {
	gorm.Model
	Street    string `json:"street"`
	City      string `json:"city"`
	ZipPostal string `json:"zip_postal"`
	Country   string `json:"country"`
}

// CustomerAddress type
type CustomerAddress struct {
	gorm.Model
	Address  Address  `json:"address"`
	Customer Customer `json:"customer"`
}

// CustomerOrder type
type CustomerOrder struct {
	gorm.Model
	Customer Customer  `json:"customer"`
	Product  Product   `json:"product"`
	DatePaid time.Time `json:"date_paid"`
}

// CustomerOrderDelivery type
type CustomerOrderDelivery struct {
	gorm.Model
	CustomerOrder CustomerOrder `json:"customer_order"`
	Status        string        `json:"status"`
}
