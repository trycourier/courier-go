package courier

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// AuditEventActor is the type for audit event actor entity
type AuditEventActor struct {
	Email string `json:"email"`
	Id    string `json:"id"`
}

// AuditEventTarget is the type for audit event target entity
type AuditEventTarget struct {
	Email string `json:"email"`
	Id    string `json:"id"`
}

// AuditEvent is the type for audit event entity
type AuditEvent struct {
	AuditEventId string `json:"auditEventId"`
	Source       string `json:"source"`
	Timestamp    string `json:"timestamp"`
	Type         string `json:"type"`
	Actor        *AuditEventActor
	Target       *AuditEventTarget
}

// ListAuditEventsResponse is type for the GET /audit-events response
type ListAuditEventsResponse struct {
	Results []AuditEvent `json:"results"`
	Paging  *PagingResponse
}

// GetAuditEvent calls the GET /audit-events/:auditEventId endpoint of the Courier API
func (c *Client) GetAuditEvent(ctx context.Context, auditEventID string) (*AuditEvent, error) {
	if auditEventID == "" {
		return nil, errors.New("audit Event ID is required")
	}
	url := fmt.Sprintf(c.API.BaseURL+"/audit-events/%s", auditEventID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	bytes, err := c.API.ExecuteRequest(req)
	if err != nil {
		return nil, err
	}

	var data AuditEvent
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// ListAuditEvents calls the GET /audit-events endpoint of the Courier API
func (c *Client) ListAuditEvents(ctx context.Context, cursor string) (*ListAuditEventsResponse, error) {
	url := fmt.Sprintf(c.API.BaseURL + "/audit-events")

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	if cursor != "" {
		q.Add("cursor", cursor)
	}
	req.URL.RawQuery = q.Encode()

	bytes, err := c.API.ExecuteRequest(req)
	if err != nil {
		return nil, err
	}

	var data ListAuditEventsResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
