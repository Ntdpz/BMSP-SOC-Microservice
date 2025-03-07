package models

import "time"

type Alarm struct {
	ID            int        `json:"id" db:"id"`
	AlarmID       string     `json:"alarm_id" db:"alarm_id"`
	Title         string     `json:"title" db:"title"`
	Priority      string     `json:"priority" db:"priority"`
	Timestamp     string     `json:"timestamp" db:"timestamp"`
	EventName     string     `json:"event_name" db:"event_name"`
	ETDATimestamp *time.Time `json:"etda_timestamp,omitempty" db:"etda_timestamp"`
	CustomerName  string     `json:"customer_name" db:"customer_name"`
	Platform      string     `json:"platform" db:"platform"`
	URL           string     `json:"url" db:"url"`
}
