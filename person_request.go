package ciscosparkresty

// PersonRequest is the Create Person Request Parameters
type PersonRequest struct {
	Emails      []string `json:"emails,omitempty"`      // Email addresses of the person
	DisplayName string   `json:"displayName,omitempty"` // Full name of the person
	FirstName   string   `json:"firstName,omitempty"`   // First name of the person
	LastName    string   `json:"lastName,omitempty"`    // Last name of the person
	Avatar      string   `json:"avatar,omitempty"`      // URL to the person's avatar in PNG format
	OrgID       string   `json:"orgId,omitempty"`       // ID of the organization to which this person belongs
	Roles       []string `json:"roles,omitempty"`       // Roles of the person
	Licenses    []string `json:"licenses,omitempty"`    // Licenses allocated to the person
}
