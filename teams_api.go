package ciscosparkresty

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"
	"gopkg.in/resty.v1"
)

// TeamsService is the service to communicate with the Teams API endpoint
type TeamsService service

// CreateTeam Creates a team.
/* Creates a team. The authenticated user is automatically added as a member of the team.
See the Team Memberships API to learn how to add more people to the team.

 @param teamCreateRequest
 @return Team
*/
func (s *TeamsService) CreateTeam(teamCreateRequest *TeamCreateRequest) (*Team, *resty.Response, error) {

	path := "/teams/"

	response, err := RestyClient.R().
		SetBody(teamCreateRequest).
		SetResult(&Team{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Team)
	return result, response, err

}

// DeleteTeam Deletes a team, by ID.
/* Deletes a team, by ID.
Specify the team ID in the teamId parameter in the URI.

 @param teamID Team ID.
 @return
*/
func (s *TeamsService) DeleteTeam(teamID string) (*resty.Response, error) {

	path := "/teams/{teamId}"
	path = strings.Replace(path, "{"+"teamId"+"}", fmt.Sprintf("%v", teamID), -1)

	response, err := RestyClient.R().
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// GetTeam Shows details for a team, by ID.
/* Shows details for a team, by ID.
Specify the team ID in the teamId parameter in the URI.

 @param teamId Team ID.
 @return Team
*/
func (s *TeamsService) GetTeam(teamID string) (*Team, *resty.Response, error) {

	path := "/teams/{teamId}"
	path = strings.Replace(path, "{"+"teamId"+"}", fmt.Sprintf("%v", teamId), -1)

	response, err := RestyClient.R().
		SetResult(&Team{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Team)
	return result, response, err

}

// ListTeamsQueryParams are the query params for the ListTeams API Call
type ListTeamsQueryParams struct {
	Max int32 `url:"max,omitempty"` // Limit the maximum number of items in the response.
}

// ListTeams Lists teams to which the authenticated user belongs.
/* Lists teams to which the authenticated user belongs.
@param "max" (int32) Limit the maximum number of items in the response.
@return Teams
*/
func (s *TeamsService) ListTeams(queryParams *ListTeamsQueryParams) (*Teams, *resty.Response, error) {

	path := "/teams/"

	queryParamsString, _ := query.Values(queryParams)

	response, err := RestyClient.R().
		SetQueryString(queryParamsString.Encode()).
		SetResult(&Teams{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Teams)
	return result, response, err

}

// UpdateTeam Updates details for a team, by ID.
/* Updates details for a team, by ID.
Specify the team ID in the teamId parameter in the URI.

 @param teamId Team ID.
 @param teamUpdateRequest
 @return Team
*/
func (s *TeamsService) UpdateTeam(teamId string, teamUpdateRequest *TeamUpdateRequest) (*Team, *resty.Response, error) {

	path := "/teams/{teamId}"
	path = strings.Replace(path, "{"+"teamId"+"}", fmt.Sprintf("%v", teamId), -1)

	response, err := RestyClient.R().
		SetBody(teamUpdateRequest).
		SetResult(&Team{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Team)
	return result, response, err

}
