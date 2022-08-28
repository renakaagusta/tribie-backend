package entity

import (
	"time"
)

type TripMember struct {
	ID        	string  `json:"id"`
	TripId      string  `json:"trip_id"`
	UserId      string  `json:"user_id"`
	Name 		string 	`json:"name"`
	CreatedAt 	time.Time `json:"created_at"`
	UpdatedAt 	time.Time `json:"updated_at"`
}
