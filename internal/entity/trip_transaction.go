package entity

import (
	"time"
)

type Transaction struct {
	ID          string    `json:"id"`
	TripId      string    `json:"trip_id"`
	UserPaidId  string    `json:"user_paid_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
