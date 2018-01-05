package ciscosparkapi

import (
	"time"
)

// Membership is the Membership definition
type Membership struct {
	ID                string    `json:"id,omitempty"`                // Membership ID.
	RoomID            string    `json:"roomId,omitempty"`            // Room ID.
	PersonID          string    `json:"personId,omitempty"`          // Person ID.
	PersonEmail       string    `json:"personEmail,omitempty"`       // Person email.
	PersonDisplayName string    `json:"personDisplayName,omitempty"` // Person display name.
	IsModerator       bool      `json:"isModerator,omitempty"`       // Membership is moderator.
	IsMonitor         bool      `json:"isMonitor,omitempty"`         // Membership is monitor.
	Created           time.Time `json:"created,omitempty"`           // Membership creation date/time.
}
