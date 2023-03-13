package api

import (
	"rudderstack/api/routers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func IntializeRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// Register event routes
	eventRoutes := router.Group("events")
	routers.RegisterEventRoutes(eventRoutes, db)

	// Register tracking plan routes
	trackingPlanRoutes := router.Group("tracking-plans")
	routers.RegisterTrackingPlanRoutes(trackingPlanRoutes, db)

	return router
}
