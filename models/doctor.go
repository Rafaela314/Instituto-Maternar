package models

import (
	"time"
)

type Doctor struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	CRM         string    `json:"crm"`
	AverageRate int32     `json:"average_rate"`
	CreatedAt   time.Time `json:"created_at"`
}
