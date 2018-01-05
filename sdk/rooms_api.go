package ciscosparkapi

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"
	"gopkg.in/resty.v1"
)

// RoomsService is the service to communicate with the Rooms API endpoint
type RoomsService service

// CreateRoom Creates a room. The authenticated user is automatically added as a member of the room.
/* Creates a room. The authenticated user is automatically added as a member of the room. See the Memberships API to learn how to add more people to the room.
To create a 1-to-1 room, use the Create Messages endpoint to send a message directly to another person by using the toPersonID or toPersonEmail parameters.

 @param roomCreateRequest
 @return Room
*/
func (s *RoomsService) CreateRoom(roomCreateRequest *RoomCreateRequest) (*Room, *resty.Response, error) {

	path := "/rooms/"

	response, err := RestyClient.R().
		SetBody(roomCreateRequest).
		SetResult(&Room{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Room)
	return result, response, err

}

// DeleteRoom Deletes a room, by ID. Deleted rooms cannot be recovered.
/* Deletes a room, by ID. Deleted rooms cannot be recovered.
Deleting a room that is part of a team will archive the room instead.
Specify the room ID in the roomID parameter in the URI

 @param roomID Room ID.
 @return
*/
func (s *RoomsService) DeleteRoom(roomID string) (*resty.Response, error) {

	path := "/rooms/{roomId}"
	path = strings.Replace(path, "{"+"roomId"+"}", fmt.Sprintf("%v", roomID), -1)

	response, err := RestyClient.R().
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// GetRoom Shows details for a room, by ID.
/* Shows details for a room, by ID.
The title of the room for 1-to-1 rooms will be the display name of the other person.
Specify the room ID in the roomID parameter in the URI.

 @param roomID Room ID.
 @return Room
*/
func (s *RoomsService) GetRoom(roomID string) (*Room, *resty.Response, error) {

	path := "/rooms/{roomId}"
	path = strings.Replace(path, "{"+"roomId"+"}", fmt.Sprintf("%v", roomID), -1)

	response, err := RestyClient.R().
		SetResult(&Room{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Room)
	return result, response, err

}

// ListRoomsQueryParams are the query params for the ListRooms API Call
type ListRoomsQueryParams struct {
	TeamID string `url:"teamId,omitempty"` // Limit the rooms to those associated with a team, by ID.
	Type_  string `url:"type_,omitempty"`  // direct returns all 1-to-1 rooms. group returns all group rooms.
	SortBy string `url:"sortBy,omitempty"` // Sort results by room ID (id), most recent activity (lastactivity), or most recently created (created).
	Max    int32  `url:"max,omitempty"`    // Limit the maximum number of items in the response.
}

// ListRooms List rooms.
/* List rooms.
The title of the room for 1-to-1 rooms will be the display name of the other person.
By default, lists rooms to which the authenticated user belongs.
Long result sets will be split into pages.

 @param "teamId" (string) Limit the rooms to those associated with a team, by ID.
 @param "type_" (string) direct returns all 1-to-1 rooms. group returns all group rooms.
 @param "sortBy" (string) Sort results by room ID (id), most recent activity (lastactivity), or most recently created (created).
 @param "max" (int32) Limit the maximum number of items in the response.
 @return Rooms
*/
func (s *RoomsService) ListRooms(queryParams *ListRoomsQueryParams) (*Rooms, *resty.Response, error) {

	path := "/rooms/"

	queryParamsString, _ := query.Values(queryParams)

	response, err := RestyClient.R().
		SetQueryString(queryParamsString.Encode()).
		SetResult(&Rooms{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Rooms)
	return result, response, err

}

// UpdateRoom Updates details for a room, by ID.
/* Updates details for a room, by ID.
Specify the room ID in the roomID parameter in the URI.

 @param roomID Room ID.
 @param roomUpdateRequest
 @return Room
*/
func (s *RoomsService) UpdateRoom(roomID string, roomUpdateRequest *RoomUpdateRequest) (*Room, *resty.Response, error) {

	path := "/rooms/{roomId}"
	path = strings.Replace(path, "{"+"roomId"+"}", fmt.Sprintf("%v", roomID), -1)

	response, err := RestyClient.R().
		SetBody(roomUpdateRequest).
		SetResult(&Room{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Room)
	return result, response, err

}
