package models

type FilterAlarm struct {
	EventStatus  string `query:"eventstatus"`
	CustomerName string `query:"customer_name"`
}
