package models

import (
	"time"
)

// Review represent review model
type Review struct {
	ID             int64     `json:"id"`
	UserID         string    `json:"user_id"`
	Title          string    `json:"title" validate:"required"`
	Content        string    `json:"content" validate:"required"`
	Date           time.Time `json:"date"`
	Classification string    `json:"classification"` // public or private
	Amount         float64   `json:"amount"`
	Insurance      string    `json:"insurance"`
	OverallRate    int32     `json:"overall_rate"`
	PlaceID        int       `json:"place_id"`
	PlaceRate      int32     `json:"place_rate"`
	DoctorID       int       `json:"doctor_id"`
	DoctorRate     int32     `json:"doctor_rate"`
	MidwifeID      int       `json:"midwife"`
	MidwifeRate    int32     `json:"midwife_rate"`
	DoulaID        int       `json:"doula_id"`
	DoulaRate      int32     `json:"doula_rate"`
	Team           string    `json:"team"`
	TeamRate       int32     `json:"team_rate"`
	UpdatedAt      time.Time `json:"updated_at"`
	CreatedAt      time.Time `json:"created_at"`
	Image          string    `json:"image"`
}

type Doctor struct {
	ID          int       `json:"id"`
	Name        int       `json:"name"`
	CRM         int       `json:"crm"`
	AverageRate int32     `json:"average_rate"`
	CreatedAt   time.Time `json:"created_at"`
}

type Place struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Address        string    `json:"address"`
	City           string    `json:"city"`
	State          string    `json:"state"`
	Country        string    `json:"country"`
	Classification string    `json:"classification"` // public or private
	AverageRate    int32     `json:"average_rate"`
	CreatedAt      time.Time `json:"created_at"`
}

type Midwife struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	AverageRate int32     `json:"average_rate"`
	CreatedAt   time.Time `json:"created_at"`
}

type Doula struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	AverageRate int32     `json:"average_rate"`
	CreatedAt   time.Time `json:"created_at"`
}
