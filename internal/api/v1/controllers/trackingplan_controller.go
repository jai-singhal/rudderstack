package controllers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	models "rudderstack/internal/api/v1/models"
	repositories "rudderstack/internal/api/v1/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Controller handles HTTP requests for the tracking plan resource.
type TrackingPlanController struct {
	repo *repositories.TrackingPlanRepository
}

// NewTrackingPlanController creates a new instance of the tracking plan controller.
func NewTrackingPlanController(repo *repositories.TrackingPlanRepository) *TrackingPlanController {
	return &TrackingPlanController{
		repo: repo,
	}
}

// GetTrackingPlanHandler returns an HTTP handler for retrieving a tracking plan by ID.
func (c *TrackingPlanController) GetTrackingPlanHandler(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		log.Printf("failed to parse event id")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "failed to parse event id"})
		return
	}

	trackingPlan, err := c.repo.GetTrackingPlanByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatus(http.StatusNotFound)
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	eventRules, err := c.repo.GetEventRulesByTrackingPlanId(id)
	trackingPlanReturn := gin.H{
		"id":    trackingPlan.ID,
		"name":  trackingPlan.DisplayName,
		"rules": eventRules,
	}
	ctx.JSON(http.StatusOK, trackingPlanReturn)
}

func (c *TrackingPlanController) GetAllTrackingPlansHandler(ctx *gin.Context) {
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 5 // default limit
	}
	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		offset = 0 // default offset
	}

	trackingPlans, total, err := c.repo.GetAllTrackingPlans(limit, offset)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"items": trackingPlans,
		"pagination": gin.H{
			"count": len(trackingPlans),
			"total": total,
		},
	})
}

// CreateTrackingPlanHandler returns an HTTP handler for creating a new tracking plan.
func (c *TrackingPlanController) CreateTrackingPlanHandler(ctx *gin.Context) {
	var requestBody models.TrackingPlanRequestBody

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	trackingPlan := &models.TrackingPlan{
		DisplayName: requestBody.TrackingPlan.DisplayName,
		Description: requestBody.TrackingPlan.Description,
	}

	eventRules := make([]*models.EventRule, len(requestBody.TrackingPlan.Rules.Events))
	for i, event := range requestBody.TrackingPlan.Rules.Events {
		eventRules[i] = &models.EventRule{
			TrackingPlanID: trackingPlan.ID,
			Name:           event.Name,
			Description:    event.Description,
			Rules:          event.Rules,
		}
	}

	if err := c.repo.CheckTrackingPlanExists(trackingPlan); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, eventRule := range eventRules {
		if exists := c.repo.CheckEventRuleExists(eventRule.Name); exists == true {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"error": fmt.Sprintf("Event rule '%s' already exists.", eventRule.Name)})
			return
		}
	}

	if err := c.repo.CreateTrackingPlan(trackingPlan, eventRules); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Tracking plan created successfully", "data": trackingPlan})
}

// UpdateTrackingPlanHandler returns an HTTP handler for updating an existing tracking plan.
func (c *TrackingPlanController) UpdateTrackingPlanHandler(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var requestBody models.TrackingPlanRequestBody

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	trackingPlan, err := c.repo.GetTrackingPlanByID(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	trackingPlan.DisplayName = requestBody.TrackingPlan.DisplayName
	trackingPlan.Description = requestBody.TrackingPlan.Description

	eventRules := make([]*models.EventRule, len(requestBody.TrackingPlan.Rules.Events))
	for i, event := range requestBody.TrackingPlan.Rules.Events {
		eventRules[i] = &models.EventRule{
			TrackingPlanID: trackingPlan.ID,
			Name:           event.Name,
			Description:    event.Description,
			Rules:          event.Rules,
		}
	}

	if err := c.repo.UpdateTrackingPlan(trackingPlan, eventRules); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Tracking plan updated successfully", "data": trackingPlan})
}
