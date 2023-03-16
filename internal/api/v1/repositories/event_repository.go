package repositories

import (
	"errors"
	"fmt"
	"log"
	models "rudderstack/internal/api/v1/models"

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
	events := make([]*models.Event, 0)
	var total int64

	// Get the total number of tracking plans
	if err := r.db.Model(&models.Event{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Limit(limit).Offset(offset).Order("created_at DESC").Find(&events).Error; err != nil {
		return nil, total, err
	}
	return events, total, nil
}

// GetEventByID retrieves a single event by ID from the repository.
func (r *EventRepository) GetEventByID(id int64) (*models.Event, error) {
	event := new(models.Event)
	result := r.db.First(&event, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return event, fmt.Errorf("No event rule found for name %d", id)
	}
	return event, result.Error
}

// GetEventByName retrieves a single event by name from the repository.
func (r *EventRepository) GetEventByName(name string) (*models.Event, error) {
	event := new(models.Event)
	result := r.db.Table(event.TableName()).Where("name = ?", name).First(&event)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return event, fmt.Errorf("No event rule found for name '%s'", name)
	}
	return event, result.Error
}

// GetEventByName retrieves a single event by name from the repository.
func (r *EventRepository) GetAllEventsByEventRule(name string, limit, offset int) ([]*models.Event, int64, error) {
	events := make([]*models.Event, 0)
	var total int64

	// Get the total number of tracking plans
	if err := r.db.Model(&models.Event{}).Where("name = ?", name).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Model(&models.Event{}).Where(
		"name = ?", name).Limit(limit).Offset(offset).Order(
		"created_at DESC").Find(&events).Error; err != nil {
		return nil, total, err
	}
	return events, total, nil
}

// GetEventRuleByName retrieves a single Event rule by event name from the repository.
func (r *EventRepository) GetEventRulesByName(eventName string) ([]*models.EventRule, error) {
	eventRules := make([]*models.EventRule, 0)

	result := r.db.Model(models.EventRule{}).Where("name = ?", eventName).Find(&eventRules)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return eventRules, fmt.Errorf("No event rule found for name '%s'", eventName)
	}
	return eventRules, result.Error
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
