package ciscosparkresty

import (
	"time"
)

// Person is the Person definition
type Person struct {
	ID           string    `json:"id,omitempty"`           // Person ID.
	Emails       []string  `json:"emails,omitempty"`       // Person email array.
	DisplayName  string    `json:"displayName,omitempty"`  // Person display name.
	NickName     string    `json:"nickName,omitempty"`     // Person nickname.
	FirstName    string    `json:"firstName,omitempty"`    // Person first name.
	LastName     string    `json:"lastName,omitempty"`     // Person last name.
	Avatar       string    `json:"avatar,omitempty"`       // Person avatar URL.
	OrgID        string    `json:"orgId,omitempty"`        // Person organization ID.
	Roles        []string  `json:"roles,omitempty"`        // Person roles.
	Licenses     []string  `json:"licenses,omitempty"`     // Person licenses.
	Created      time.Time `json:"created,omitempty"`      // Person creation date/time.
	TimeZone     string    `json:"timeZone,omitempty"`     // Person time zone.
	LastActivity time.Time `json:"lastActivity,omitempty"` // Person last active date/time.
	Status       string    `json:"status,omitempty"`       // Person presence status (active or inactive).
	PersonType   string    `json:"type,omitempty"`         // Person type (person or bot).
}
