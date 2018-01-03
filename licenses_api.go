package ciscosparkresty

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"
	"gopkg.in/resty.v1"
)

// LicensesService is the service to communicate with the Licenses API endpoint
type LicensesService service

// GetLicense Shows details for a license, by ID.
/* Shows details for a license, by ID.
Specify the license ID in the licenseID parameter in the URI.

 @param licenseId License ID.
 @return License
*/
func (s *LicensesService) GetLicense(licenseID string) (*License, *resty.Response, error) {

	path := "/licenses/{licenseId}"
	path = strings.Replace(path, "{"+"licenseId"+"}", fmt.Sprintf("%v", licenseID), -1)

	response, err := RestyClient.R().
		SetResult(&License{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*License)
	return result, response, err

}

// ListLicensesQueryParams are the query params for the ListLicenses API Call
type ListLicensesQueryParams struct {
	Max int32 `url:"max,omitempty"` // Limit the maximum number of items in the response.
}

// ListLicenses List all licenses for a given organization.
/* List all licenses for a given organization.
If no orgId is specified, the default is the organization of the authenticated user.

 @param "max" (int32) Limit the maximum number of items in the response.
 @return Licenses
*/
func (s *LicensesService) ListLicenses(queryParams *ListLicensesQueryParams) (*Licenses, *resty.Response, error) {

	path := "/licenses/"

	queryParamsString, _ := query.Values(queryParams)

	response, err := RestyClient.R().
		SetQueryString(queryParamsString.Encode()).
		SetResult(&Licenses{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Licenses)
	return result, response, err

}
