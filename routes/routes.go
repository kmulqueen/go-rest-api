package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kmulqueen/go-rest-api/controllers"
)

func RegisterRoutes(server *gin.Engine) {
	//* v1 API
	{
		v1 := server.Group("/api/v1")

		//* Event Routes
		{
			eventRoutes := v1.Group("/events")
			eventRoutes.GET("/", controllers.GetEvents)
			eventRoutes.GET("/:id", controllers.GetEvent)
			eventRoutes.POST("/", controllers.CreateEvent)
			eventRoutes.PUT("/:id", controllers.UpdateEvent)
			eventRoutes.DELETE("/:id", controllers.DeleteEvent)
		}
	}

}
