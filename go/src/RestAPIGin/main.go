package main

import (
	"example.com/restApiGin/db"
	routes "example.com/restApiGin/route"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")

}
