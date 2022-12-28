package models

import (
	"time"
)

// Review represent review model
type Review struct {
	ID          int64     `json:"id"`
	Username    string    `json:"username"`
	Title       string    `json:"title" validate:"required"`
	Content     string    `json:"content" validate:"required"`
	Date        time.Time `json:"date"`
	City        string    `json:"city"`
	State       string    `json:"state"`
	Country     string    `json:"country"`
	Place       string    `json:"place"`
	PlaceRate   int32     `json:"place_rate"`
	Doctor      string    `json:"doctor"`
	DoctorRate  int32     `json:"doctor_rate"`
	CRM         string    `json:"crm"`
	Midwife     string    `json:"midwife"`
	MidwifeRate int32     `json:"midwife_rate"`
	Doula       string    `json:"doula"`
	DoulaRate   int32     `json:"doula_rate"`
	Team        string    `json:"team"`
	TeamRate    int32     `json:"team_rate"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
	Image       string    `json:"image"`
}
