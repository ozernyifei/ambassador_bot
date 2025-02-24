// Package domain contains the core business models of the Ambassador Bot system.
package domain

import "time"

// User represents a user in the Ambassador Bot system.
// TelegramID is used as the primary identifier.
type User struct {
	// TelegramID is the unique Telegram user identifier.
	TelegramID int64
	// Name is the user's display name.
	Name string
	// CreatedAt is the timestamp when the user was registered.
	CreatedAt time.Time
}

// NewUser creates a new User instance with the given TelegramID and name.
// It sets the CreatedAt field to the current time.
func NewUser(telegramID int64, name string) *User {
	return &User{
		TelegramID: telegramID,
		Name:       name,
		CreatedAt:  time.Now(),
	}
}
