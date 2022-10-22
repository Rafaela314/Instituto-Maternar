package models

import (
	"time"
)

//Review represent review model
type Review struct {
	ID         int       `json:"id"`
	Username   string    `json:"username"`
	Title      string    `json:"title" validate:"required"`
	Content    string    `json:"content" validate:"required"`
	Date       string    `json:"date"`
	City       int       `json:"city"`
	State      int       `json:"state"`
	Place      string    `json:"place"`
	PlaceRate  int       `json:"placerate"`
	Doctor     string    `json:"doctor"`
	DoctorRate int       `json:"doctorrate"`
	Team       string    `json:"team"`
	TeamRate   int       `json:"teamrate"`
	UpdatedAt  time.Time `json:"updated_at"`
	CreatedAt  time.Time `json:"created_at"`
}
