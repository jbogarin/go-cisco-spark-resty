package ciscosparkresty

// MembershipUpdateRequest is the Update Membership Request object
type MembershipUpdateRequest struct {
	IsModerator bool `json:"isModerator,omitempty"` // Membership is a moderator.
}
