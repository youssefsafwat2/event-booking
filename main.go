package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/youssefsafwat2/event-booking/db"
	"github.com/youssefsafwat2/event-booking/models"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8080")

}

func getEvents(context *gin.Context) {
	events, err := models.GetEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not retrieve events, Try again later",
			"status":  "error"})
		return
	}
	context.JSON(200, gin.H{
		"events":  events,
		"message": "Events retrieved successfully",
		"status":  "success",
	})
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(400, gin.H{
			"message": "Invalid input",
			"status":  "error",
		})
		return
	}
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create event, Try again later",
			"status":  "error"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"event":   event,
		"message": "Event created successfully",
		"status":  "success",
	})
}
