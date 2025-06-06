package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/youssefsafwat2/event-booking/models"
	"github.com/youssefsafwat2/event-booking/utils"
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
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
			"status":  "error",
		})
		return
	}

	existingUser, err := models.GetUserByEmail(user.Email)
	if err != nil || existingUser == nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid email or password",
			"status":  "error",
		})
		return
	}

	if !models.CheckPasswordHash(user.Password, existingUser.Password) {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid email or password",
			"status":  "error",
		})
		return
	}
	user.ID = existingUser.ID
	token, err := utils.GenerateToken(existingUser.Email, existingUser.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not generate token, Try again later",
			"status":  "error",
		})
		return
	}
	// existingUser.Token = token
	existingUser.Password = "" // Clear password before sending response
	context.Set("user", existingUser)

	context.JSON(http.StatusOK, gin.H{
		"user":    existingUser,
		"token":   token,
		"message": "Login successful",
		"status":  "success",
	})

}
