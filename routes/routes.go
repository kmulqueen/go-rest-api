package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kmulqueen/go-rest-api/controllers"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", controllers.GetEvents)
	server.GET("/events/:id", controllers.GetEvent)
	server.POST("/events", controllers.CreateEvent)
	server.PUT("/events/:id", controllers.UpdateEvent)
}
