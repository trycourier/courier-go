package courier

// WebhookResponse represents the payload provided by an outbound webhook
type WebhookResponse struct {
	Data MessageResponse `json:"data"`
	Type string          `json:"type"`
}
