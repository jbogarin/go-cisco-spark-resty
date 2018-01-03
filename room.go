package ciscosparkresty

import (
	"time"
)

// Room is the Room definition
type Room struct {
	ID           string    `json:"id,omitempty"`           // Room ID.
	Title        string    `json:"title,omitempty"`        // Room title.
	RoomType     string    `json:"type,omitempty"`         // Room type (group or direct).
	IsLocked     bool      `json:"isLocked,omitempty"`     // Room is moderated.
	TeamID       string    `json:"teamId,omitempty"`       // Room Team ID.
	CreatorID    string    `json:"creatorId,omitempty"`    // Room creator Person ID.
	LastActivity time.Time `json:"lastActivity,omitempty"` // Room last activity date/time.
	Created      time.Time `json:"created,omitempty"`      // Room creation date/time.
}
