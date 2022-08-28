package entity

import (
	"time"
)

type TransactionExpenses struct {
	ID            string    `json:"id"`
	TripId        string    `json:"trip_id"`
	TripMemberId  string    `json:"trip_member_id"`
	TransactionId string    `json:"transaction_id"`
	ItemId        string    `json:"item_id"`
	Quantity      int64     `json:"quantity"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
