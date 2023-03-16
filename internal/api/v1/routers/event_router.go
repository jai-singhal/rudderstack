package routers

import (
	controllers "rudderstack/internal/api/v1/controllers"
	repositories "rudderstack/internal/api/v1/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterEventRoutes(routerGroup *gin.RouterGroup, db *gorm.DB) {
	eventRepo := repositories.NewEventRepository(db)
	eventController := controllers.NewEventController(eventRepo)

	eventRoutes := routerGroup
	{
		eventRoutes.GET("", eventController.GetAllEventsHandler)
		eventRoutes.POST("", eventController.CreateEventHandler)
		eventRoutes.GET("/:id", eventController.GetEventHandler)

		eventRoutes.PUT("/:id", eventController.UpdateEventHandler)

	}
}
