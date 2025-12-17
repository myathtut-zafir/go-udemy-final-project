package routes

import (
	"example.com/event-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventByID)
	auth := server.Group("/")
	auth.Use(middlewares.Authenticate)
	auth.POST("/events", createEvent)
	auth.PUT("/events/:id", updateEvent)
	auth.DELETE("/events/:id", deleteEvent)
	auth.POST("/events/:id/register", registerEvent)
	auth.DELETE("/events/:id/register", cancelRegisterEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)
}
