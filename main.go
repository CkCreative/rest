package main

import (
	"github.com/CkCreative/rest/controllers"
	"github.com/CkCreative/rest/middleware"
	"github.com/CkCreative/rest/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.GET("/users", controllers.FindUsers)
	r.POST("/users", controllers.CreateUser)
	r.POST("/login", controllers.Login)
	r.GET("/users/:id", controllers.FindBook)
	r.PATCH("/users/:id", controllers.UpdateBook)
	r.DELETE("/users/:id", controllers.DeleteBook)

	r.Use(middleware.JwtVerify())
	r.GET("/auth", controllers.FindUsers)

	r.Run()

}
