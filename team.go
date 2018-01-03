package ciscosparkresty

import (
	"time"
)

// Team is the Team definition
type Team struct {
	ID        string    `json:"id,omitempty"`        // Team ID.
	Name      string    `json:"name,omitempty"`      // Team Name.
	CreatorID string    `json:"creatorId,omitempty"` // Team creator ID.
	Created   time.Time `json:"created,omitempty"`   // Team creation date/time.
}
