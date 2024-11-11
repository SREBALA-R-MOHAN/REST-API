package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent) //id is dynamic events/5,events/1
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent) //PUT commonly used for update requests
	server.DELETE("/events/:id", deleteEvent)
}
