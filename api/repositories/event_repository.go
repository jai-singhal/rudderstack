package repositories

import (
	"errors"
	"fmt"
	"log"
	models "rudderstack/api/models"

	"gorm.io/gorm"
)

// EventRepository is a concrete implementation of EventRepository using GORM.
type EventRepository struct {
	db *gorm.DB
}

// NewEventRepository CreateEvents a new instance of the EventRepository.
func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{db: db}
}

// GetAllEvents returns a slice of all events with pagination.
func (r *EventRepository) GetAllEvents(limit, offset int) ([]*models.Event, int64, error) {
	var events []*models.Event
	var total int64

	// Get the total number of tracking plans
	if err := r.db.Model(&models.TrackingPlan{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Limit(limit).Offset(offset).Order("created_at DESC").Find(&events).Error; err != nil {
		return nil, total, err
	}
	return events, total, nil
}

// GetEventByID retrieves a single event by ID from the repository.
func (r *EventRepository) GetEventByID(id int64) (*models.Event, error) {
	var event *models.Event
	result := r.db.First(&event, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return event, fmt.Errorf("No event rule found for name %d", id)
	}
	return event, result.Error
}

// GetEventByName retrieves a single event by name from the repository.
func (r *EventRepository) GetEventByName(name string) (*models.Event, error) {
	var event *models.Event
	result := r.db.Table(event.TableName()).Where("name = ?", name).First(&event)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return event, fmt.Errorf("No event rule found for name '%s'", name)
	}
	return event, result.Error
}

func (r *EventRepository) GetEventRuleByName(eventName string) (*models.EventRule, error) {
	var eventRule *models.EventRule
	result := r.db.Table(eventRule.TableName()).Where("name = ?", eventName).First(&eventRule)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return eventRule, fmt.Errorf("No event rule found for name '%s'", eventName)
	}
	return eventRule, result.Error
}

// CreateEvent adds a new event to the repository.
func (r *EventRepository) CreateEvent(event *models.Event, eventTracking *models.EventTracking) error {
	log.Printf("Model before insertion: %+v\n", event)
	// begin the transaction
	tx := r.db.Begin()

	if err := tx.Table(event.TableName()).Create(event).Error; err != nil {
		tx.Rollback()
		return err
	}

	eventTracking.EventID = event.ID
	if err := tx.Table(eventTracking.TableName()).Create(eventTracking).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction if all operations succeed
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

// UpdateEvent modifies an existing event in the repository.
func (r *EventRepository) UpdateEvent(event *models.Event) error {
	result := r.db.Save(event)
	return result.Error
}
