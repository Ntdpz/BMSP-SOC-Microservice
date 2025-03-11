package models

type FilterAlarm struct {
	IsOpen       *bool  `query:"is_open"`
	CustomerName string `query:"customer_name"`
}
