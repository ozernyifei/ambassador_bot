// Package interfaces provides interface definitions for external API interactions
// in the Ambassador Bot system.
package interfaces

import "context"

// TelegramAPI defines an interface for interacting with the Telegram Bot API.
// This abstraction enables the application to send, edit, or delete messages
// without being tightly coupled to a specific HTTP client or implementation.
type TelegramAPI interface {
	// SendMessage sends a message to a specific chat.
	// Parameters:
	//   - ctx: context for cancellation and timeout control.
	//   - chatID: the unique identifier of the target chat.
	//   - text: the text content of the message to be sent.
	// Returns an error if the message could not be sent.
	SendMessage(ctx context.Context, chatID int64, text string) error

	// EditMessage edits an existing message in a chat.
	// Parameters:
	//   - ctx: context for cancellation and timeout control.
	//   - chatID: the unique identifier of the target chat.
	//   - messageID: the identifier of the message to be edited.
	//   - text: the new text for the message.
	// Returns an error if the message could not be edited.
	EditMessage(ctx context.Context, chatID int64, messageID int, text string) error

	// DeleteMessage deletes a message from a chat.
	// Parameters:
	//   - ctx: context for cancellation and timeout control.
	//   - chatID: the unique identifier of the target chat.
	//   - messageID: the identifier of the message to delete.
	// Returns an error if the message could not be deleted.
	DeleteMessage(ctx context.Context, chatID int64, messageID int) error
}
