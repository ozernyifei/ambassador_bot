// Package interfaces provides definitions for repository interfaces
// that abstract data access operations for the Ambassador Bot.
package interfaces

import (
	"context"
	"time"

	"github.com/ozernyifei/ambassador_bot/internal/domain"
)

// UserRepository defines the methods required to access user data.
// TelegramID is used as the primary key.
type UserRepository interface {
	// Save persists a new user in the repository.
	Save(ctx context.Context, user *domain.User) error

	// GetByTelegramID retrieves a user by their TelegramID.
	GetByTelegramID(ctx context.Context, telegramID int64) (*domain.User, error)
}

// EventRepository defines the methods required to access event data.
type EventRepository interface {
	// Save persists a new event in the repository.
	Save(ctx context.Context, event *domain.Event) error

	// GetUpcoming retrieves events scheduled for the future.
	GetUpcoming(ctx context.Context) ([]*domain.Event, error)
}

// NotificationRepository defines the methods required to access notification data.
type NotificationRepository interface {
	// Save persists a new notification in the repository.
	Save(ctx context.Context, notification *domain.Notification) error

	// GetScheduled retrieves notifications scheduled to be sent before the specified time.
	GetScheduled(ctx context.Context, before time.Time) ([]*domain.Notification, error)
}
