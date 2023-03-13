package routers

import (
	controllers "rudderstack/api/controllers"
	repositories "rudderstack/api/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterTrackingPlanRoutes(routerGroup *gin.RouterGroup, db *gorm.DB) {
	trackingPlanRepo := repositories.NewTrackingPlanRepository(db)
	trackingPlanController := controllers.NewTrackingPlanController(trackingPlanRepo)

	trackingPlanRoutes := routerGroup
	{
		trackingPlanRoutes.GET("/", trackingPlanController.GetAllTrackingPlansHandler)
		trackingPlanRoutes.POST("/", trackingPlanController.CreateTrackingPlanHandler)
		trackingPlanRoutes.GET("/:id", trackingPlanController.GetTrackingPlanHandler)
		trackingPlanRoutes.PUT("/:id", trackingPlanController.UpdateTrackingPlanHandler)
	}
}
