package entity

import (
	"time"
)

type TransactionPayment struct {
	ID            string    `json:"id"`
	TripId        string    `json:"trip_id"`
	TripMemberId  string    `json:"trip_member_id"`
	TransactionId string    `json:"transaction_id"`
	Nominal       int64     `json:"nominal"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
