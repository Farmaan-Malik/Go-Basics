package main

import (
	"example/rest-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "My Get request"})
	})
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8080")
}

func getEvents(ctx *gin.Context) {
	events := models.GetAllEvents()
	ctx.JSON(200, gin.H{"status": "success", "data": events})
}

func createEvent(ctx *gin.Context) {
	var event models.Events
	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(400, gin.H{"status": "failure", "error": err})
		return
	}
	err = event.Save()
	if err != nil {
		ctx.JSON(400, gin.H{"status": "failure", "error": err})
		return
	}
	ctx.JSON(201, gin.H{"status": "success", "data": event})
}
