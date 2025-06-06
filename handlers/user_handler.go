package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/youssefsafwat2/event-booking/models"
)

func GetUsers(context *gin.Context) {
	users, err := models.GetUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not retrieve users, Try again later",
			"status":  "error"})
		return
	}
	context.JSON(200, gin.H{
		"users":   users,
		"message": "Users retrieved successfully",
		"status":  "success",
	})
}

func Signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
			"status":  "error",
		})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create user, Try again later",
			"status":  "error",
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"user":    user,
		"message": "User created successfully",
		"status":  "success",
	})

}

func Login(context *gin.Context) {

}
