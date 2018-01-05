package ciscosparkapi

// WebhookCreateRequest is the Webhook Create Request Parameters
type WebhookCreateRequest struct {
	Name      string `json:"name,omitempty"`      // Webhook name.
	TargetURL string `json:"targetUrl,omitempty"` // Webhook target URL.
	Resource  string `json:"resource,omitempty"`  // Webhook resource.
	Event     string `json:"event,omitempty"`     // Webhook event.
	Filter    string `json:"filter,omitempty"`    // Webhook filter.
	Secret    string `json:"secret,omitempty"`    // Webhook secret.
}
