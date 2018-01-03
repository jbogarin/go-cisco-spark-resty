package ciscosparkresty

// MembershipCreateRequest is the Create Membership Request Parameters
type MembershipCreateRequest struct {
	RoomID      string `json:"roomId,omitempty"`      // Room ID.
	PersonID    string `json:"personId,omitempty"`    // Person ID.
	PersonEmail string `json:"personEmail,omitempty"` // Person email.
	IsModerator bool   `json:"isModerator,omitempty"` // Membership is a moderator.
}
