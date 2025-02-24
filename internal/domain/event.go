// Package domain contains the core business models for the Ambassador Bot system.
package domain

import "time"

// Event represents an event within the Ambassador Bot system, such as a meeting,
// training session, lecture, online session, volunteer project, or discussion group.
// It contains essential information about the event including its title, description,
// and scheduling details.
type Event struct {
	// ID is the unique identifier for the event.
	// This may be automatically assigned by the database.
	ID int64

	// Title is the name or title of the event.
	Title string

	// Description provides detailed information about the event.
	Description string

	// StartTime specifies when the event begins.
	StartTime time.Time

	// EndTime specifies when the event ends.
	EndTime time.Time

	// CreatedAt is the timestamp when the event was created.
	CreatedAt time.Time
}

// NewEvent creates a new Event instance with the specified title, description, start time, and end time.
// The CreatedAt field is set to the current time.
// Parameters:
//   - title: the title or name of the event
//   - description: a detailed description of the event
//   - startTime: the starting time of the event
//   - endTime: the ending time of the event
//
// Returns:
//   - a pointer to the newly created Event instance
func NewEvent(title, description string, startTime, endTime time.Time) *Event {
	return &Event{
		Title:       title,
		Description: description,
		StartTime:   startTime,
		EndTime:     endTime,
		CreatedAt:   time.Now(),
	}
}
