package ciscosparkapi

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"
	"gopkg.in/resty.v1"
)

// OrganizationsService is the service to communicate with the Organizations API endpoint
type OrganizationsService service

// GetOrganization Shows details for an organization, by ID.
/* Shows details for an organization, by ID.
Specify the org ID in the orgID parameter in the URI.

 @param orgID Organization ID.
 @return Organization
*/
func (s *OrganizationsService) GetOrganization(orgID string) (*Organization, *resty.Response, error) {

	path := "/organizations/{orgId}"
	path = strings.Replace(path, "{"+"orgId"+"}", fmt.Sprintf("%v", orgID), -1)

	response, err := RestyClient.R().
		SetResult(&Organization{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Organization)
	return result, response, err

}

// ListOrganizationsQueryParams are the query params for the ListOrganizations API Call
type ListOrganizationsQueryParams struct {
	Max int32 `url:"max,omitempty"` // Limit the maximum number of items in the response.
}

// ListOrganizations List all organizations visible by your account. The results will not be paginated.
/* List all organizations visible by your account. The results will not be paginated.
@param "max" (int32) Limit the maximum number of items in the response.
@return Organizations
*/
func (s *OrganizationsService) ListOrganizations(queryParams *ListOrganizationsQueryParams) (*Organizations, *resty.Response, error) {

	path := "/organizations/"

	queryParamsString, _ := query.Values(queryParams)

	response, err := RestyClient.R().
		SetQueryString(queryParamsString.Encode()).
		SetResult(&Organizations{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Organizations)
	return result, response, err

}
