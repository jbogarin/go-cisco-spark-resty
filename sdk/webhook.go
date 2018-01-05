package ciscosparkapi

import (
	"time"
)

// Webhook is the Webhook definition
type Webhook struct {
	ID        string    `json:"id,omitempty"`        // Webhook ID.
	Name      string    `json:"name,omitempty"`      // Webhook name.
	TargetURL string    `json:"targetUrl,omitempty"` // Webhook target URL.
	Resource  string    `json:"resource,omitempty"`  // Webhook resource.
	Event     string    `json:"event,omitempty"`     // Webhook event.
	OrgID     string    `json:"orgId,omitempty"`     // Webhook organization ID.
	CreatedBy string    `json:"createdBy,omitempty"` // Webhook created by Person ID.
	AppID     string    `json:"appId,omitempty"`     // Webhook application ID.
	OwnedBy   string    `json:"ownedBy,omitempty"`   // Webhook owner Person ID.
	Filter    string    `json:"filter,omitempty"`    // Webhook filter.
	Status    string    `json:"status,omitempty"`    // Webhook status.
	Secret    string    `json:"secret,omitempty"`    // Webhook secret.
	Created   time.Time `json:"created,omitempty"`   // Webhook creation date/time.
}
