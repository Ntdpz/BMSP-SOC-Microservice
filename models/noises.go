package models

type Noises struct {
	ID     int    `json:"id" db:"id"`
	Noises string `json:"noises" db:"noises"`
}
