// Package interfaces provides interface definitions for handling bot commands
// and interactions in the Ambassador Bot system.
package interfaces

import (
	"context"

	"github.com/ozernyifei/ambassador_bot/internal/domain"
)

// BotController defines the interface for handling Telegram bot commands and user interactions.
// This interface abstracts the application logic for processing commands such as /start, /register,
// fetching events, and sending notifications.
type BotController interface {
	// HandleStart processes the /start command.
	// It typically welcomes the user and performs any necessary initial setup.
	// Parameters:
	//   - ctx: a context for cancellation and timeout control.
	//   - telegramID: the unique Telegram identifier of the user.
	//   - userName: the name of the user.
	// Returns an error if the operation fails.
	HandleStart(ctx context.Context, telegramID int64, userName string) error

	// HandleRegister registers a new user using their TelegramID and name.
	// Parameters:
	//   - ctx: a context for cancellation and timeout control.
	//   - telegramID: the unique Telegram identifier of the user.
	//   - userName: the name of the user.
	// Returns an error if registration is unsuccessful.
	HandleRegister(ctx context.Context, telegramID int64, userName string) error

	// HandleEvents retrieves a list of upcoming events for the user.
	// Parameters:
	//   - ctx: a context for cancellation and timeout control.
	//   - telegramID: the unique Telegram identifier of the user.
	// Returns:
	//   - a slice of pointers to Event objects representing upcoming events.
	//   - an error if event retrieval fails.
	HandleEvents(ctx context.Context, telegramID int64) ([]*domain.Event, error)

	// HandleNotification processes notification requests for the user.
	// This may involve sending motivational quotes or event reminders.
	// Parameters:
	//   - ctx: a context for cancellation and timeout control.
	//   - telegramID: the unique Telegram identifier of the user.
	// Returns an error if the notification cannot be sent.
	HandleNotification(ctx context.Context, telegramID int64) error
}
