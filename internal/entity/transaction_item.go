package entity

import (
	"time"
)

type TransactionItem struct {
	ID        		string  `json:"id"`
	TripId      	string  `json:"trip_id"`
	TransactionId   string  `json:"transaction_id"`
	Title     		string  `json:"title"`
	Description     string  `json:"description"`
	Price			int64		`json:"price"`
	CreatedAt 		time.Time `json:"created_at"`
	UpdatedAt 		time.Time `json:"updated_at"`
}
