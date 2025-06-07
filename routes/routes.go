package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/youssefsafwat2/event-booking/handlers"
	"github.com/youssefsafwat2/event-booking/middlewares"
)

func RegisterEventRoutes(server *gin.Engine) {
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	server.GET("/events", handlers.GetEvents)
	server.GET("/events/:id", handlers.GetEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/events", handlers.CreateEvent)
	authenticated.PUT("/events/:id", handlers.UpdateEvent)
	authenticated.DELETE("/events/:id", handlers.DeleteEvent)
	authenticated.POST("/events/:id/register", handlers.RegisterForEvent)
	authenticated.DELETE("/events/:id/register", handlers.CancelRegistration)

	server.POST("/signup", handlers.Signup)
	server.GET("/users", handlers.GetUsers)
	server.POST("/login", handlers.Login)
}
