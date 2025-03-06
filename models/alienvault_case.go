package models

type AlienvaultCase struct {
	ID        int    `json:"id" db:"id"`
	AlarmID   string `json:"alarm_id" db:"alarm_id"`
	Title     string `json:"title" db:"title"`
	Priority  string `json:"priority" db:"priority"`
	Timestamp string `json:"timestamp" db:"timestamp"`
	EventName string `json:"event_name" db:"event_name"`
}
