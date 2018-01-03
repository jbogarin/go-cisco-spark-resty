package ciscosparkresty

import (
	"time"
)

// Message is the Message definition
type Message struct {
	ID              string    `json:"id,omitempty"`              // Message ID.
	RoomID          string    `json:"roomId,omitempty"`          // Room ID.
	RoomType        string    `json:"roomType,omitempty"`        // Room type (group or direct).
	ToPersonID      string    `json:"toPersonId,omitempty"`      // Person ID (for type=direct).
	ToPersonEmail   string    `json:"toPersonEmail,omitempty"`   // Person email (for type=direct).
	Text            string    `json:"text,omitempty"`            // Message in plain text format.
	Markdown        string    `json:"markdown,omitempty"`        // Message in markdown format.
	Files           []string  `json:"files,omitempty"`           // File URL array.
	PersonID        string    `json:"personId,omitempty"`        // Person ID.
	PersonEmail     string    `json:"personEmail,omitempty"`     // Person Email.
	Created         time.Time `json:"created,omitempty"`         // Message creation date/time.
	MentionedPeople []string  `json:"mentionedPeople,omitempty"` // Person ID array.
	MentionedGroups []string  `json:"mentionedGroups,omitempty"` // Groups array.
}
