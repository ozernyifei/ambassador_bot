// Package domain contains the core business models for the Ambassador Bot system.
package domain

import "time"

// Notification represents a message or alert that is scheduled to be sent to users.
// Notifications can serve various purposes, such as reminders for events or sending
// motivational quotes. Each notification includes its type, content, and scheduling information.
type Notification struct {
	// ID is the unique identifier of the notification.
	// This may be assigned by the database when the notification is persisted.
	ID int64

	// Type specifies the category of the notification, for example "reminder" or "quote".
	Type string

	// Message is the content to be delivered to the user.
	Message string

	// ScheduledAt indicates the time when the notification should be sent.
	ScheduledAt time.Time

	// CreatedAt is the timestamp when the notification was created.
	CreatedAt time.Time
}

// NewNotification creates a new Notification instance with the given type, message, and scheduled time.
// It sets the CreatedAt field to the current time.
// Parameters:
//   - notificationType: the category/type of the notification (e.g., "reminder", "quote")
//   - message: the text content of the notification
//   - scheduledAt: the time at which the notification should be delivered
//
// Returns:
//   - a pointer to the created Notification instance
func NewNotification(notificationType, message string, scheduledAt time.Time) *Notification {
	return &Notification{
		Type:        notificationType,
		Message:     message,
		ScheduledAt: scheduledAt,
		CreatedAt:   time.Now(),
	}
}
