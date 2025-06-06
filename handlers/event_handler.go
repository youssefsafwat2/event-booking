package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/youssefsafwat2/event-booking/models"
)

func GetEvents(context *gin.Context) {
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

func CreateEvent(context *gin.Context) {
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

func GetEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event ID",
			"status":  "error",
		})
		return
	}
	event, err := models.GetEventByID(id)

	if event == nil && err == nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Event not found",
			"status":  "error",
		})
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not retrieve event, Try again later",
			"status":  "error",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"event":   event,
		"message": "Event retrieved successfully",
		"status":  "success",
	})

}

func UpdateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event ID",
			"status":  "error",
		})
		return
	}
	var event *models.Event
	event, err = models.GetEventByID(id)
	if event == nil && err == nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Event not found",
			"status":  "error",
		})
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not retrieve event, Try again later",
			"status":  "error",
		})
		return
	}
	var updatedEvent models.Event

	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
			"status":  "error",
		})
		return
	}

	updatedEvent.ID = id

	err = updatedEvent.UpdateEvent()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not update event, Try again later",
			"status":  "error",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully",
		"status":  "success",
	})

}

func DeleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event ID",
			"status":  "error",
		})
		return
	}
	var event *models.Event
	event, err = models.GetEventByID(id)
	if event == nil && err == nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Event not found",
			"status":  "error",
		})
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not retrieve event, Try again later",
			"status":  "error",
		})
		return
	}
	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not delete event, Try again later",
			"status":  "error",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event deleted successfully",
		"status":  "success",
	})

}
