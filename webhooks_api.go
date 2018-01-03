package ciscosparkresty

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"
	"gopkg.in/resty.v1"
)

// WebhooksService is the service to communicate with the Webhooks API endpoint
type WebhooksService service

// CreateWebhook Creates a webhook.
/* Creates a webhook.
@param webhookCreateRequest
@return Webhook
*/
func (s *WebhooksService) CreateWebhook(webhookCreateRequest *WebhookCreateRequest) (*Webhook, *resty.Response, error) {

	path := "/webhooks/"

	response, err := RestyClient.R().
		SetBody(webhookCreateRequest).
		SetResult(&Webhook{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Webhook)
	return result, response, err

}

// DeleteWebhook Deletes a webhook, by ID.
/* Deletes a webhook, by ID.
Specify the webhook ID in the webhookId parameter in the URI.

 @param webhookID Webhook ID.
 @return
*/
func (s *WebhooksService) DeleteWebhook(webhookID string) (*resty.Response, error) {

	path := "/webhooks/{webhookId}"
	path = strings.Replace(path, "{"+"webhookId"+"}", fmt.Sprintf("%v", webhookID), -1)

	response, err := RestyClient.R().
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// GetWebhook Shows details for a webhook, by ID.
/* Shows details for a webhook, by ID.
Specify the webhook ID in the webhookId parameter in the URI.

 @param webhookID Webhook ID.
 @return Webhook
*/
func (s *WebhooksService) GetWebhook(webhookID string) (*Webhook, *resty.Response, error) {

	path := "/webhooks/{webhookId}"
	path = strings.Replace(path, "{"+"webhookId"+"}", fmt.Sprintf("%v", webhookID), -1)

	response, err := RestyClient.R().
		SetResult(&Webhook{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Webhook)
	return result, response, err

}

// ListWebhooksQueryParams are the query params for the ListWebhooks API Call
type ListWebhooksQueryParams struct {
	Max int32 `url:"max,omitempty"` // Limit the maximum number of items in the response.
}

// ListWebhooks Lists all of your webhooks.
/* Lists all of your webhooks.
@param "max" (int32) Limit the maximum number of items in the response.
@return Webhooks
*/
func (s *WebhooksService) ListWebhooks(queryParams *ListWebhooksQueryParams) (*Webhooks, *resty.Response, error) {

	path := "/webhooks/"

	queryParamsString, _ := query.Values(queryParams)

	response, err := RestyClient.R().
		SetQueryString(queryParamsString.Encode()).
		SetResult(&Webhooks{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Webhooks)
	return result, response, err

}

// UpdateWebhook Updates a webhook, by ID.
/* Updates a webhook, by ID.
Specify the webhook ID in the webhookId parameter in the URI.

 @param webhookID Webhook ID.
 @param updateWebhookRequest
 @return Webhook
*/
func (s *WebhooksService) UpdateWebhook(webhookID string, updateWebhookRequest *UpdateWebhookRequest) (*Webhook, *resty.Response, error) {

	path := "/webhooks/{webhookId}"
	path = strings.Replace(path, "{"+"webhookId"+"}", fmt.Sprintf("%v", webhookID), -1)

	response, err := RestyClient.R().
		SetBody(updateWebhookRequest).
		SetResult(&Webhook{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Webhook)
	return result, response, err

}
