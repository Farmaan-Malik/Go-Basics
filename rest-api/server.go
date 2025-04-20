package main

import (
	"example/rest-api/db"
	"example/rest-api/models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "My Get request"})
	})
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)
	server.POST("/events", createEvent)
	server.Run(":8080")
}

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(500, gin.H{"status": "failure", "error": fmt.Sprint(err)})
		return
	}
	ctx.JSON(200, gin.H{"status": "success", "data": events})
}

func createEvent(ctx *gin.Context) {
	var event models.Events
	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(400, gin.H{"status": "failure", "error": fmt.Sprint(err)})
		return
	}
	err = event.Save()
	if err != nil {
		ctx.JSON(400, gin.H{"status": "failure", "error": fmt.Sprint(err)})
		return
	}
	ctx.JSON(201, gin.H{"status": "success", "data": event})
}

func getEventById(ctx *gin.Context) {

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"status": "failure", "error": "Could not parse event id"})
		return
	}
	data, err := models.GetEventById(id)
	if err != nil {
		ctx.JSON(500, gin.H{"status": "failure", "error": fmt.Sprint(err)})
		return
	}
	ctx.JSON(200, gin.H{"status": "success", "data": data})
}
