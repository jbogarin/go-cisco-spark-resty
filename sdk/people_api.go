package ciscosparkapi

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"
	"gopkg.in/resty.v1"
)

// PeopleService is the service to communicate with the People API endpoint
type PeopleService service

// CreatePerson Create a new user account for a given organization.
/* Create a new user account for a given organization. Only an admin can create a new user account.
Currently, users may have only one email address associated with their account. The emails parameter is an array, which accepts multiple values to allow for future expansion, but currently only one email address will be used for the new user.

 @param personRequest
 @return Person
*/
func (s *PeopleService) CreatePerson(personRequest *PersonRequest) (*Person, *resty.Response, error) {

	path := "/people/"

	response, err := RestyClient.R().
		SetBody(personRequest).
		SetResult(&Person{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Person)
	return result, response, err

}

// Delete Remove a person from the system. Only an admin can remove a person.
/* Remove a person from the system. Only an admin can remove a person.
Specify the person ID in the personID parameter in the URI.

 @param personID Person ID.
 @return
*/
func (s *PeopleService) DeletePerson(personID string) (*resty.Response, error) {

	path := "/people/{personId}"
	path = strings.Replace(path, "{"+"personId"+"}", fmt.Sprintf("%v", personID), -1)

	response, err := RestyClient.R().
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// GetMe Show the profile for the authenticated user.
/* Show the profile for the authenticated user. This is the same as GET /people/:id using the Person ID associated with your Auth token.
@return Person
*/
func (s *PeopleService) GetMe() (*Person, *resty.Response, error) {

	path := "/people/me"

	response, err := RestyClient.R().
		SetResult(&Person{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Person)
	return result, response, err

}

// GetPerson Shows details for a person, by ID.
/* Shows details for a person, by ID. Certain fields, such as status or lastActivity, will only be displayed for people within your organization or an organzation you manage.
Specify the person ID in the personID parameter in the URI.

 @param personID Person ID.
 @return Person
*/
func (s *PeopleService) GetPerson(personID string) (*Person, *resty.Response, error) {

	path := "/people/{personId}"
	path = strings.Replace(path, "{"+"personId"+"}", fmt.Sprintf("%v", personID), -1)

	response, err := RestyClient.R().
		SetResult(&Person{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Person)
	return result, response, err

}

// ListPeopleQueryParams are the query params for the ListPeople API Call
type ListPeopleQueryParams struct {
	ID          string `url:"id,omitempty"`          // List people by ID. Accepts up to 85 person IDs separated by commas.
	Email       string `url:"email,omitempty"`       // List people with this email address. For non-admin requests, either this or displayName are required.
	DisplayName string `url:"displayName,omitempty"` // List people whose name starts with this string. For non-admin requests, either this or email are required.
	Max         int32  `url:"max,omitempty"`         // Limit the maximum number of items in the response.
	OrgID       string `url:"orgId,omitempty"`       // List people in this organization. Only admin users of another organization (such as partners) may use this parameter.
}

// ListPeople List people in your organization.
/* List people in your organization. For most users, either the email or displayName parameter is required.
Admin users can omit these fields and list all users in their organization.

 @param "id" (string) List people by ID. Accepts up to 85 person IDs separated by commas.
 @param "email" (string) List people with this email address. For non-admin requests, either this or displayName are required.
 @param "displayName" (string) List people whose name starts with this string. For non-admin requests, either this or email are required.
 @param "max" (int32) Limit the maximum number of items in the response.
 @param "orgId" (string) List people in this organization. Only admin users of another organization (such as partners) may use this parameter.
 @return People
*/
func (s *PeopleService) ListPeople(queryParams *ListPeopleQueryParams) (*People, *resty.Response, error) {

	path := "/people/"

	queryParamsString, _ := query.Values(queryParams)

	response, err := RestyClient.R().
		SetQueryString(queryParamsString.Encode()).
		SetResult(&People{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*People)
	return result, response, err

}

// Update Update details for a person, by ID.
/* Update details for a person, by ID.
Specify the person ID in the personID parameter in the URI. Only an admin can update a person details. Email addresses for a person cannot be changed via the Cisco Spark API.
Include all details for the person. This action expects all user details to be present in the request. A common approach is to first GET the person&#39;s details, make changes, then PUT both the changed and unchanged values.

 @param personID Person ID.
 @param personRequest
 @return Person
*/
func (s *PeopleService) Update(personID string, personRequest *PersonRequest) (*Person, *resty.Response, error) {

	path := "/people/{personId}"
	path = strings.Replace(path, "{"+"personId"+"}", fmt.Sprintf("%v", personID), -1)

	response, err := RestyClient.R().
		SetBody(personRequest).
		SetResult(&Person{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Person)
	return result, response, err

}
