package routes

import (
	"github.com/CkCreative/rest/book/controllers"
	"github.com/gin-gonic/gin"
)

// Router for books
func Router(r *gin.Engine) {
	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
}
