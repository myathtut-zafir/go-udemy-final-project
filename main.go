package main

import (
	"example.com/event-api/db"
	"example.com/event-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run() // listen and seserverve on 0.0.0.0:8080
}
