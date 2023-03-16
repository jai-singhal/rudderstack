package controllers

import (
	"errors"
	"fmt"
	"net/http"
	models "rudderstack/internal/api/v1/models"
	repositories "rudderstack/internal/api/v1/repositories"
	utils "rudderstack/internal/api/v1/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type EventController struct {
	repo *repositories.EventRepository
}

func NewEventController(repo *repositories.EventRepository) *EventController {
	return &EventController{repo}
}

func (c *EventController) GetAllEventsHandler(ctx *gin.Context) {
	limitStr := ctx.Query("limit")
	offsetStr := ctx.Query("offset")
	eventRule := ctx.Query("rule")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 5 // default limit
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0 // default offset
	}

	var events []*models.Event
	var total int64
	var err1 error
	if len(eventRule) == 0 {
		events, total, err1 = c.repo.GetAllEvents(limit, offset)
	} else {
		events, total, err1 = c.repo.GetAllEventsByEventRule(eventRule, limit, offset)
	}
	if err1 != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"items": events,
		"pagination": gin.H{
			"count": len(events),
			"total": total,
		},
	})
}

func (c *EventController) GetEventHandler(ctx *gin.Context) {
	eventID, err := strconv.ParseInt(ctx.Param("eventID"), 10, 64)

	if err != nil {
		log.Printf("failed to parse event id")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "failed to parse event id"})
		return
	}

	event, err := c.repo.GetEventByID(eventID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatus(http.StatusNotFound)
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func (c *EventController) CreateEventHandler(ctx *gin.Context) {
	var eventRequestBody models.EventRequestBody

	if err := ctx.ShouldBindJSON(&eventRequestBody); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	eventRules, err := c.repo.GetEventRulesByName(eventRequestBody.Name)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(eventRules) == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": fmt.Sprintf("No Event rules found for '%s'", eventRequestBody.Name)})
		return
	}

	event := &models.Event{
		Name:       eventRequestBody.Name,
		Properties: eventRequestBody.Properties,
	}

	for _, eventRule := range eventRules {
		if err := utils.ValidateJSON(eventRule.Rules, eventRequestBody.Properties); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"error": fmt.Sprintf("Error while vaildating event rule %s. %s", eventRule.Name, err.Error())})
			return
		}

		eventTracking := &models.EventTracking{
			EventRuleID: eventRule.ID,
		}

		if err := c.repo.CreateEvent(event, eventTracking); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError,
				gin.H{"error": fmt.Sprintf("Error Creating Event for event rule: %s. %s", eventRule.Name, err.Error())})
			return
		}
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "data": event})
}

func (c *EventController) UpdateEventHandler(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "failed to parse event id"})
		return
	}

	event, err := c.repo.GetEventByID(eventId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var eventRequestBody models.EventRequestBody
	if err := ctx.ShouldBindJSON(&eventRequestBody); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	eventRules, err := c.repo.GetEventRulesByName(eventRequestBody.Name)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, eventRule := range eventRules {
		if err := utils.ValidateJSON(eventRule.Rules, eventRequestBody.Properties); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"error": fmt.Sprintf("Error while vaildating event rule %s. %s", eventRule.Name, err.Error())})
			return
		}

		updatedEvent := &models.Event{
			ID:         eventId,
			Name:       eventRequestBody.Name,
			Properties: eventRequestBody.Properties,
			CreatedAt:  event.CreatedAt,
			UpdatedAt:  time.Now(),
		}

		if err := c.repo.UpdateEvent(updatedEvent); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}
