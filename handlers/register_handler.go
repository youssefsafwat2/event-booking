package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/youssefsafwat2/event-booking/models"
)

func RegisterForEvent(context *gin.Context) {
	userID := context.GetInt64("userID")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{
			"message": "Invalid event ID",
			"status":  "error",
		})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not retrieve event, Try again later",
			"status":  "error",
		})
		return
	}
	if event == nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Event not found",
			"status":  "error",
		})
		return
	}
	err = event.RegisterUserForEvent(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not register for event, Try again later",
			"status":  "error",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Registered for event successfully",
		"status":  "success",
		"event":   event,
	})
}

func CancelRegistration(context *gin.Context) {
	userID := context.GetInt64("userID")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{
			"message": "Invalid event ID",
			"status":  "error",
		})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not retrieve event, Try again later",
			"status":  "error",
		})
		return
	}
	if event == nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Event not found",
			"status":  "error",
		})
		return
	}
	err = event.CancelRegistration(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not canceling the registration, Try again later",
			"status":  "error",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Registration canceled successfully",
		"status":  "success",
	})
}
