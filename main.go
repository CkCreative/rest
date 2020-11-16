package main

import (
	bookRoutes "github.com/CkCreative/rest/book/routes"
	userRoutes "github.com/CkCreative/rest/user/routes"
	"github.com/CkCreative/rest/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	utils.ConnectDatabase()

	bookRoutes.Router(router)
	userRoutes.Router(router)

	router.Run()

}
