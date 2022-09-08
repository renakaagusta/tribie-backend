package entity

import (
	"time"
)

type Transaction struct {
	ID            string    `json:"id"`
	TripId        string    `json:"trip_id"`
	UserPaidId    string    `json:"user_paid_id"`
	GrandTotal    int       `json:"grand_total"`
	Method        string    `json:"method"`
	SubTotal      int       `json:"sub_total"`
	ServiceCharge int       `json:"service_charge"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
