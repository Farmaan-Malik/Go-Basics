package routes

import (
	"example/rest-api/models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

func updateEvent(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		fmt.Println("here", err)
		ctx.JSON(400, gin.H{"status": "failure", "error": "Could not parse event id"})
		return
	}
	_, err = models.GetEventById(id)
	if err != nil {
		fmt.Println("here 2", err)

		ctx.JSON(500, gin.H{"status": "failure", "error": fmt.Sprint(err)})
		return
	}
	var updatedEvent models.Events
	err = ctx.ShouldBindJSON(&updatedEvent)
	if err != nil {
		fmt.Println("here 3", err)

		ctx.JSON(400, gin.H{"status": "failure", "error": fmt.Sprint(err)})
		return
	}
	updatedEvent.ID = id
	err = updatedEvent.UpdateEvent()
	if err != nil {
		ctx.JSON(500, gin.H{"status": "failure", "error": fmt.Sprint(err)})
		return
	}
	ctx.JSON(200, gin.H{"status": "success", "data": "Event updated successfully!"})

}


