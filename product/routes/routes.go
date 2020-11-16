package routes

import (
	"github.com/CkCreative/rest/product/controllers"
	"github.com/gin-gonic/gin"
)

// Router for product
func Router(r *gin.Engine) {
	r.GET("/products", controllers.FindProducts)
	r.POST("/products", controllers.CreateProduct)
	r.GET("/products/:id", controllers.FindProduct)
	r.PATCH("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)

	r.GET("/orders", controllers.FindOrders)
	r.POST("/orders", controllers.CreateOrder)
	r.GET("/orders/:id", controllers.FindOrder)
	r.PATCH("/orders/:id", controllers.UpdateOrder)
	r.DELETE("/orders/:id", controllers.DeleteOrder)

	r.GET("/addresses", controllers.FindAddresses)
	r.POST("/addresses", controllers.CreateAddress)
	r.GET("/addresses/:id", controllers.FindAddress)
	r.PATCH("/addresses/:id", controllers.UpdateAddress)
	r.DELETE("/addresses/:id", controllers.DeleteAddress)

	r.GET("/customers", controllers.FindCustomers)
	r.POST("/customers", controllers.CreateCustomer)
	r.GET("/customers/:id", controllers.FindCustomer)
	r.PATCH("/customers/:id", controllers.UpdateCustomer)
	r.DELETE("/customers/:id", controllers.DeleteCustomer)

	r.GET("/inventory", controllers.FindInventories)
	r.POST("/inventory", controllers.CreateInventory)
	r.GET("/inventory/:id", controllers.FindInventory)
	r.PATCH("/inventory/:id", controllers.UpdateInventory)
	r.DELETE("/inventory/:id", controllers.DeleteInventory)
}
