package entity

import (
	"time"
)

// User represents a user.
type User struct {
	ID   string
	Name string
	AppleID string
	Email string
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetID returns the user ID.
func (u User) GetID() string {
	return u.ID
}

// GetName returns the user name.
func (u User) GetName() string {
	return u.Name
}
