package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/events", getEvents)
	server.Run() // listen and seserverve on 0.0.0.0:8080
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}
