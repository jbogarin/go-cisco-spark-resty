package ciscosparkresty

// UpdateWebhookRequest is Update Webhook Request Parameters
type UpdateWebhookRequest struct {
	Name      string `json:"name,omitempty"`      // Webhook name.
	TargetURL string `json:"targetUrl,omitempty"` // Webhook target URL.
}
