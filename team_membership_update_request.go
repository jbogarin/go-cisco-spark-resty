package ciscosparkresty

// TeamMembershipUpdateRequest is the Team Membership Update Request object
type TeamMembershipUpdateRequest struct {
	IsModerator bool `json:"isModerator,omitempty"` // Team Membership is a moderator.
}
