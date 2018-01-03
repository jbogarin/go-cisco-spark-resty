package ciscosparkresty

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"
	"gopkg.in/resty.v1"
)

// TeamMembershipsService is the service to communicate with the TeamMemberships API endpoint
type TeamMembershipsService service

// CreateTeamMembership Add someone to a team by Person ID or email address; optionally making them a moderator.
/* Add someone to a team by Person ID or email address; optionally making them a moderator.
@param teamMemberhipCreateRequest
@return TeamMembership
*/
func (s *TeamMembershipsService) CreateTeamMembership(teamMemberhipCreateRequest *TeamMembershipCreateRequest) (*TeamMembership, *resty.Response, error) {

	path := "/team/memberships/"

	response, err := RestyClient.R().
		SetBody(teamMemberhipCreateRequest).
		SetResult(&TeamMembership{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*TeamMembership)
	return result, response, err

}

// DeleteTeamMembership Deletes a team membership, by ID.
/* Deletes a team membership, by ID.
Specify the team membership ID in the membershipId URI parameter.

 @param membershipID Membership ID.
 @return
*/
func (s *TeamMembershipsService) DeleteTeamMembership(membershipID string) (*resty.Response, error) {

	path := "/team/memberships/{membershipId}"
	path = strings.Replace(path, "{"+"membershipId"+"}", fmt.Sprintf("%v", membershipID), -1)

	response, err := RestyClient.R().
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// GetTeamMembership Shows details for a team membership, by ID.
/* Shows details for a team membership, by ID.
Specify the team membership ID in the membershipId URI parameter.

 @param membershipId Membership ID.
 @return TeamMembership
*/
func (s *TeamMembershipsService) GetTeamMembership(membershipID string) (*TeamMembership, *resty.Response, error) {

	path := "/team/memberships/{membershipId}"
	path = strings.Replace(path, "{"+"membershipId"+"}", fmt.Sprintf("%v", membershipID), -1)

	response, err := RestyClient.R().
		SetResult(&TeamMembership{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*TeamMembership)
	return result, response, err

}

// ListTeamMemberhipsQueryParams are the query params for the ListTeamMemberhips API Call
type ListTeamMemberhipsQueryParams struct {
	TeamID string `url:"teamId,omitempty"` // Team ID.
	Max    int32  `url:"max,omitempty"`    // Limit the maximum number of items in the response.
}

// ListTeamMemberhips Lists all team memberships for a given team, specified by the teamId query parameter.
/* Lists all team memberships for a given team, specified by the teamId query parameter.
Use query parameters to filter the response.

 @param teamID Team ID.
 @param "max" (int32) Limit the maximum number of items in the response.
 @return TeamMemberships
*/
func (s *TeamMembershipsService) ListTeamMemberhips(queryParams *ListTeamMemberhipsQueryParams) (*TeamMemberships, *resty.Response, error) {

	path := "/team/memberships/"

	queryParamsString, _ := query.Values(queryParams)

	response, err := RestyClient.R().
		SetQueryString(queryParamsString.Encode()).
		SetResult(&TeamMemberships{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*TeamMemberships)
	return result, response, err

}

// UpdateTeamMembership Updates a team membership, by ID.
/* Updates a team membership, by ID.
Specify the team membership ID in the membershipId URI parameter.

 @param membershipId Membership ID.
 @param teamMembershipUpdateRequest
 @return TeamMembership
*/
func (s *TeamMembershipsService) UpdateTeamMembership(membershipID string, teamMembershipUpdateRequest *TeamMembershipUpdateRequest) (*TeamMembership, *resty.Response, error) {

	path := "/team/memberships/{membershipId}"
	path = strings.Replace(path, "{"+"membershipId"+"}", fmt.Sprintf("%v", membershipID), -1)

	response, err := RestyClient.R().
		SetBody(teamMembershipUpdateRequest).
		SetResult(&TeamMembership{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*TeamMembership)
	return result, response, err

}
