package repositories

import (
	"errors"
	"fmt"
	models "rudderstack/internal/api/v1/models"

	"gorm.io/gorm"
)

type TrackingPlanRepository struct {
	db *gorm.DB
}

func NewTrackingPlanRepository(db *gorm.DB) *TrackingPlanRepository {
	return &TrackingPlanRepository{db}
}

func (r *TrackingPlanRepository) GetAllTrackingPlans(limit, offset int) ([]*models.TrackingPlan, int64, error) {
	trackingPlans := make([]*models.TrackingPlan, 0)
	var total int64

	// Get the total number of tracking plans
	if err := r.db.Model(&models.TrackingPlan{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Retrieve the tracking plans based on the limit and offset
	if err := r.db.Model(&models.TrackingPlan{}).Limit(limit).Offset(offset).Order("created_at DESC").Find(&trackingPlans).Error; err != nil {
		return nil, 0, err
	}
	return trackingPlans, total, nil
}

func (r *TrackingPlanRepository) GetEventRulesByTrackingPlanId(tracking_plan_id int64) ([]*models.EventRule, error) {
	eventRules := make([]*models.EventRule, 0)
	result := r.db.Model(&models.EventRule{}).Where("tracking_plan_id = ?", tracking_plan_id).Find(&eventRules)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return eventRules, fmt.Errorf("No Event rules found for tracking_plan_id '%d'", tracking_plan_id)
	}
	return eventRules, result.Error
}

func (r *TrackingPlanRepository) CheckEventRuleExists(name string) bool {
	eventRule := new(models.EventRule)

	result := r.db.Model(&models.EventRule{}).Where("name = ?", name).First(&eventRule)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

func (r *TrackingPlanRepository) GetTrackingPlanByID(id int64) (*models.TrackingPlan, error) {
	trackingPlan := new(models.TrackingPlan)
	result := r.db.Table(trackingPlan.TableName()).Where("id = ?", id).First(&trackingPlan)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return trackingPlan, fmt.Errorf("No Tracking plan found for id '%d'", id)
	}
	return trackingPlan, result.Error
}

func (r *TrackingPlanRepository) CreateTrackingPlan(trackingPlan *models.TrackingPlan, eventRules []*models.EventRule) error {
	// begin the transaction
	tx := r.db.Begin()

	if err := tx.Table(trackingPlan.TableName()).Create(trackingPlan).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, eventRule := range eventRules {
		eventRule.TrackingPlanID = trackingPlan.ID

		if err := tx.Table(eventRule.TableName()).Create(eventRule).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// Commit the transaction if all operations succeed
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (r *TrackingPlanRepository) CheckTrackingPlanExists(trackingPlan *models.TrackingPlan) error {
	// Check if tracking plan exists
	var count int64
	err := r.db.Table(trackingPlan.TableName()).Where("display_name = ?", trackingPlan.DisplayName).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("tracking plan with name `%s` already exists!", trackingPlan.DisplayName)
	}
	return nil
}

func (r *TrackingPlanRepository) UpdateTrackingPlan(trackingPlan *models.TrackingPlan, eventRules []*models.EventRule) error {
	// begin the transaction
	tx := r.db.Begin()

	if err := tx.Table(trackingPlan.TableName()).Save(trackingPlan).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("tracking_plan_id = ?", trackingPlan.ID).Delete(&models.EventRule{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, eventRule := range eventRules {
		if err := tx.Table(eventRule.TableName()).Create(eventRule).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// Commit the transaction if all operations succeed
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}
