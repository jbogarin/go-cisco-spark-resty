package ciscosparkapi

// WebhookUpdateRequest is the Update Webhook Request Parameters
type WebhookUpdateRequest struct {
	Name      string `json:"name,omitempty"`      // Webhook name.
	TargetURL string `json:"targetUrl,omitempty"` // Webhook target URL.
}
