package routes

import (
	"example/rest-api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	authenticated := server.Group("/")
	authenticated.Use(middleware.Authentication)
	authenticated.POST("/events", createEvent)
	authenticated.GET("/events/:id", getEventById)
	authenticated.GET("/events", getEvents)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.POST("/events/:id/register", createRegistration)
	authenticated.DELETE("/events/:id/register", deleteRegistration)
	authenticated.GET("/events/register", getAllRegistrations)

	server.POST("/signup", userSignup)
	server.POST("/login", userLogin)
}
