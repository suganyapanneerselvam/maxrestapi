package main

import (
	"practice/restapi/api-tsst/db"
	"practice/restapi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default() //localhost
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
