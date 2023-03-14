package models

import (
	"encoding/json"
	"time"
)

type TrackingPlan struct {
	ID          int64     `json:"id"`
	DisplayName string    `json:"display_name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (TrackingPlan) TableName() string {
	return "tracking_plans"
}

type EventRule struct {
	ID             int64           `json:"id"`
	TrackingPlanID int64           `json:"tracking_plan_id"`
	Name           string          `json:"name"`
	Description    string          `json:"description"`
	Rules          json.RawMessage `json:"rules"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (EventRule) TableName() string {
	return "event_rules"
}

type TrackingPlanRequestBody struct {
	TrackingPlan struct {
		DisplayName string `json:"display_name" binding:"required"`
		Description string `json:"description"`
		Rules       struct {
			Events []struct {
				Name        string          `json:"name" binding:"required"`
				Description string          `json:"description"`
				Rules       json.RawMessage `json:"rules"`
			} `json:"events" binding:"required"`
		} `json:"rules" binding:"required"`
	} `json:"tracking_plan" binding:"required"`
}
