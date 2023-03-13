package models

import (
	"encoding/json"
	"time"
)

type Event struct {
	ID         int64           `json:"id"`
	Name       string          `json:"name"`
	Properties json.RawMessage `json:"properties"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at"`
}

func (Event) TableName() string {
	return "events"
}

type EventTracking struct {
	ID          int64     `json:"id"`
	EventRuleID int64     `json:"event_rule_id"`
	EventID     int64     `json:"event_id"`
	CreatedAt   time.Time `json:"created_at"`
}

func (EventTracking) TableName() string {
	return "event_tracking"
}

type EventRequestBody struct {
	Name       string          `json:"name"`
	Properties json.RawMessage `json:"properties"`
}
