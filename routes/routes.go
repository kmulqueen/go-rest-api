package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kmulqueen/go-rest-api/controllers"
	"github.com/kmulqueen/go-rest-api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	//* v1 API
	{
		v1 := server.Group("/api/v1")

		//* Event Routes
		{
			//* Public Event Routes
			eventRoutes := v1.Group("/events")
			eventRoutes.GET("/", controllers.GetEvents)
			eventRoutes.GET("/:id", controllers.GetEvent)

			//* Protected Event Routes
			protectedEventRoutes := eventRoutes.Group("/")
			protectedEventRoutes.Use(middlewares.Authenticate)
			protectedEventRoutes.POST("/", controllers.CreateEvent)
			protectedEventRoutes.PUT("/:id", controllers.UpdateEvent)
			protectedEventRoutes.DELETE("/:id", controllers.DeleteEvent)
		}

		//* User Routes
		{
			userRoutes := v1.Group("/users")
			userRoutes.POST("/signup", controllers.Signup)
			userRoutes.POST("/login", controllers.Login)
		}
	}

}
