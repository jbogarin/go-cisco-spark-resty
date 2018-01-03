package ciscosparkresty

// RoomCreateRequest is the Room Create Request Parameters
type RoomCreateRequest struct {
	Title  string `json:"title,omitempty"`  // A user-friendly name for the room.
	TeamID string `json:"teamId,omitempty"` // The ID for the team with which this room is associated.
}
