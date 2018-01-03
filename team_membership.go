package ciscosparkresty

import (
	"time"
)

// TeamMembership is the Team Membership definition
type TeamMembership struct {
	ID                string    `json:"id,omitempty"`                // Team Membership ID.
	TeamID            string    `json:"teamId,omitempty"`            // Team ID.
	PersonID          string    `json:"personId,omitempty"`          // Person ID.
	PersonEmail       string    `json:"personEmail,omitempty"`       // Person email.
	PersonDisplayName string    `json:"personDisplayName,omitempty"` // Person display name.
	IsModerator       bool      `json:"isModerator,omitempty"`       // Team Membership is moderator.
	Created           time.Time `json:"created,omitempty"`           // Team Membership creation date/time.
}
