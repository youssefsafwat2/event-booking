package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/youssefsafwat2/event-booking/utils"
)

func Authenticate(context *gin.Context) {
	token := context.GetHeader("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized, please login",
			"status":  "error",
		})
		return
	}

	userID, err := utils.ValidateToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized, log in and try again",
			"status":  "error",
		})
		return
	}
	context.Set("userID", userID)
	context.Next()
}
