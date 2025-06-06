package main

import (
	"github.com/gin-gonic/gin"
	"github.com/youssefsafwat2/event-booking/db"
	"github.com/youssefsafwat2/event-booking/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterEventRoutes(server)
	server.Run(":8080")

}
