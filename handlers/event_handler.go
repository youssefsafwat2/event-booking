package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/youssefsafwat2/event-booking/models"
	"github.com/youssefsafwat2/event-booking/utils"
)

func authenticate(context *gin.Context) (*models.User, bool) {
	token := context.GetHeader("Authorization")
	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized, please login",
			"status":  "error",
		})
		return nil, false
	}

	rawUser, exists := context.Get("user")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized, log in and try again",
			"status":  "error",
		})
		return nil, false
	}

	user := rawUser.(models.User)
	if _, err := utils.ValidateToken(token, user.Email, user.ID); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized, log in and try again",
			"status":  "error",
		})
		return nil, false
	}

	return &user, true
}

func GetEvents(context *gin.Context) {
	events, err := models.GetEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not retrieve events, try again later",
			"status":  "error",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"events":  events,
		"message": "Events retrieved successfully",
		"status":  "success",
	})
}

func CreateEvent(context *gin.Context) {
	user, ok := authenticate(context)
	if !ok {
		return
	}

	var event models.Event
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
			"status":  "error",
		})
		return
	}

	if err := event.Save(user.ID); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create event, try again later",
			"status":  "error",
		})
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
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not retrieve event, try again later",
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

	context.JSON(http.StatusOK, gin.H{
		"event":   event,
		"message": "Event retrieved successfully",
		"status":  "success",
	})
}

func UpdateEvent(context *gin.Context) {
	user, ok := authenticate(context)
	if !ok {
		return
	}

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event ID",
			"status":  "error",
		})
		return
	}

	event, err := models.GetEventByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not retrieve event, try again later",
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
	if event.UserID != user.ID {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized access",
			"status":  "error",
		})
		return
	}

	var updatedEvent models.Event
	if err := context.ShouldBindJSON(&updatedEvent); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
			"status":  "error",
		})
		return
	}
	updatedEvent.ID = id

	if err := updatedEvent.UpdateEvent(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not update event, try again later",
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
	user, ok := authenticate(context)
	if !ok {
		return
	}

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event ID",
			"status":  "error",
		})
		return
	}

	event, err := models.GetEventByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not retrieve event, try again later",
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
	if event.UserID != user.ID {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized access",
			"status":  "error",
		})
		return
	}

	if err := event.Delete(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not delete event, try again later",
			"status":  "error",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event deleted successfully",
		"status":  "success",
	})
}
