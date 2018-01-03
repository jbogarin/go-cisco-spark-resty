package ciscosparkresty

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"
	"gopkg.in/resty.v1"
)

// RolesService is the service to communicate with the Roles API endpoint
type RolesService service

// GetRole Shows details for a role, by ID.
/* Shows details for a role, by ID.
Specify the role ID in the roleId parameter in the URI.

 @param roleID Role ID.
 @return Role
*/
func (s *RolesService) GetRole(roleID string) (*Role, *resty.Response, error) {

	path := "/roles/{roleId}"
	path = strings.Replace(path, "{"+"roleId"+"}", fmt.Sprintf("%v", roleID), -1)

	response, err := RestyClient.R().
		SetResult(&Role{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Role)
	return result, response, err

}

// GetRolesQueryParams are the query params for the GetRoles API Call
type GetRolesQueryParams struct {
	Max int32 `url:"max,omitempty"` // Limit the maximum number of items in the response.
}

// GetRoles List all roles.
/* List all roles.
@param "max" (int32) Limit the maximum number of items in the response.
@return Roles
*/
func (s *RolesService) GetRoles(queryParams *GetRolesQueryParams) (*Roles, *resty.Response, error) {

	path := "/roles/"

	queryParamsString, _ := query.Values(queryParams)

	response, err := RestyClient.R().
		SetQueryString(queryParamsString.Encode()).
		SetResult(&Roles{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Roles)
	return result, response, err

}
