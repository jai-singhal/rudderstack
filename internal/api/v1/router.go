package api

import (
	routers "rudderstack/internal/api/v1/routers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func IntializeRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// Register event routes
	eventRoutes := router.Group("/api/v1/events")
	routers.RegisterEventRoutes(eventRoutes, db)

	// Register tracking plan routes
	trackingPlanRoutes := router.Group("/api/v1/tracking-plans")
	routers.RegisterTrackingPlanRoutes(trackingPlanRoutes, db)

	return router
}
