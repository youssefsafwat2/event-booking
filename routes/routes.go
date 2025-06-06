package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/youssefsafwat2/event-booking/handlers"
)

func RegisterEventRoutes(server *gin.Engine) {
	server.GET("/events", handlers.GetEvents)
	server.POST("/events", handlers.CreateEvent)
	server.GET("/events/:id", handlers.GetEvent)
	server.PUT("/events/:id", handlers.UpdateEvent)
	server.DELETE("/events/:id", handlers.DeleteEvent)
	server.POST("/signup", handlers.Signup)
	server.GET("/users", handlers.GetUsers)
	server.POST("/login", handlers.Login)
}
