package routes

import (
	"github.com/CkCreative/rest/user/controllers"
	"github.com/CkCreative/rest/user/middleware"
	"github.com/gin-gonic/gin"
)

// Router for books
func Router(r *gin.Engine) {
	r.GET("/users", controllers.FindUsers)
	r.POST("/users", controllers.CreateUser)
	r.POST("/login", controllers.Login)
	r.PATCH("/users/:id", controllers.UpdateUser)
	r.PATCH("/users/password/:id", controllers.UpdatePassword)

	r.Use(middleware.JwtVerify())
	r.GET("/auth", controllers.FindUsers)
}
