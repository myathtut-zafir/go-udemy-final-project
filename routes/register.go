package routes

import (
	"net/http"
	"strconv"

	"example.com/event-api/models"
	"github.com/gin-gonic/gin"
)

func registerEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event"})
		return
	}
	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Registered successfully"})
}
func cancelRegisterEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event"})
		return
	}
	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Registered successfully"})
}
