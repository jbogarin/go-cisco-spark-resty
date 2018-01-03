package ciscosparkresty

// TeamMembershipCreateRequest is the Team Membership Create Request Parameters
type TeamMembershipCreateRequest struct {
	TeamID      string `json:"teamId,omitempty"`      // Team ID.
	PersonID    string `json:"personId,omitempty"`    // Person ID.
	PersonEmail string `json:"personEmail,omitempty"` // Person Email.
	IsModerator bool   `json:"isModerator,omitempty"` // Team Membership is a moderator.
}
