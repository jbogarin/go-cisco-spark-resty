package ciscosparkresty

import (
	"time"
)

// Organization is the Organization definition
type Organization struct {
	ID          string    `json:"id,omitempty"`          // Organization ID.
	DisplayName string    `json:"displayName,omitempty"` // Organization Display Name.
	Created     time.Time `json:"created,omitempty"`     // Organization creation date/time.
}
