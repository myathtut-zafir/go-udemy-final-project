package main

import (
	"net/http"

	"example.com/event-api/db"
	"example.com/event-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run() // listen and seserverve on 0.0.0.0:8080
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, events)
}
func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(& )
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// event.ID = 1
	// event.UserID = 1
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Created", "event": event})
}
