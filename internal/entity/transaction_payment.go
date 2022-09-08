package entity

import (
	"time"
)

type TransactionPayment struct {
	ID            string    `json:"id"`
	TripId        string    `json:"trip_id"`
	TripMemberId  string    `json:"trip_member_id"`
	TransactionId string    `json:"transaction_id"`
	UserFromId    string    `json:"user_from_id"`
	UserToId      string    `json:"user_to_id"`
	Nominal       int64     `json:"nominal"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
