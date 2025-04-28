package routes

import (
	"example/rest-api/models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func createRegistration(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, exists := ctx.Params.Get("id")
	if !exists {
		ctx.JSON(404, gin.H{"status": "failure", "message": "could not find event"})
		return
	}
	id, err := strconv.ParseInt(eventId, 10, 64)
	if err != nil {
		ctx.JSON(404, gin.H{"status": "failure", "message": "could not parse event id"})
		return
	}
	event, err := models.GetEventById(id)
	if err != nil {
		ctx.JSON(404, gin.H{"status": "failure", "message": "could not find event"})
		return
	}
	err = event.MakeRegistration(userId)
	if err != nil {
		ctx.JSON(404, gin.H{"status": "failure", "message": fmt.Sprint(err)})
		return
	}
	ctx.JSON(201, gin.H{"status": "success", "message": "event created successfully"})
}

func deleteRegistration(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(404, gin.H{"status": "failure", "message": "could not find event"})
		return
	}
	var event models.Events
	event.ID = eventId
	err = event.DeleteRegistration(userId)
	if err != nil {
		ctx.JSON(404, gin.H{"status": "failure", "message": "error deleting registration"})
		return
	}
	ctx.JSON(200, gin.H{"status": "success", "message": "registration deleted successfully"})
}

func getAllRegistrations(ctx *gin.Context) {
	data, err := models.GetRegistrations()
	if err != nil {
		ctx.JSON(404, gin.H{"status": "failure", "message": fmt.Sprint(err)})
		return
	}
	ctx.JSON(200, gin.H{"status": "success", "data": data})
}
