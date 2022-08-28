package entity

import (
	"time"
)

type Trip struct {
	ID        	string    `json:"id"`
	Title      	string    `json:"title"`
	Description string    `json:"description"`
	Place	 	string    `json:"place"`
	CreatedAt 	time.Time `json:"created_at"`
	UpdatedAt 	time.Time `json:"updated_at"`
}
