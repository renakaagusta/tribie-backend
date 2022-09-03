package entity

import (
	"time"
)

// User represents a user.
type UserDefault struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	AppleId   string    `json:"apple_id"`
	DeviceId  string    `json:"device_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetID returns the user ID.
func (u UserDefault) GetID() string {
	return u.ID
}

// GetName returns the user name.
func (u UserDefault) GetEmail() string {
	return u.Email
}
