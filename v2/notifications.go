package courier

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type NotificationResponse struct {
	ID    string
	Title string
}

type NotificationsResponse struct {
	Paging  *PagingResponse
	Results []*NotificationResponse
}

type NotificationContentResponse struct {
	Blocks   []interface{} `json:"blocks"`
	Channels []interface{} `json:"channels,omitempty"`
	Checksum string        `json:"checksum"`
}

type NotificationVariationsRequestBody struct {
	Blocks   []interface{} `json:"blocks,omitempty"`
	Channels []interface{} `json:"channels,omitempty"`
}

type SubmissionCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Type    string `json:"type"`
	Updated int    `json:"updated,omitempty"`
}

type NotificationSubmissionChecksResponse struct {
	Checks []*SubmissionCheck
}

type NotificationSubmissionChecksRequest struct {
	Checks []SubmissionCheck `json:"checks"`
}

// GetNotifications calls the GET /notifications endpoint of the Courier API
func (c *Client) GetNotifications(ctx context.Context, cursor string) (*NotificationsResponse, error) {
	url := c.API.BaseURL + "/notifications"

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

	var data NotificationsResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// GetNotificationContent calls the GET /notifications/{notification_id}/content endpoint of the Courier API
func (c *Client) GetNotificationContent(ctx context.Context, notificationID string) (*NotificationContentResponse, error) {
	if notificationID == "" {
		return nil, errors.New("Notification ID is required")
	}

	url := fmt.Sprintf(c.API.BaseURL+"/notifications/%s/content", notificationID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	bytes, err := c.API.ExecuteRequest(req)
	if err != nil {
		return nil, err
	}

	var data NotificationContentResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// GetNotificationDraftContent calls the GET /notifications/{notification_id}/draft/content endpoint of the Courier API
func (c *Client) GetNotificationDraftContent(ctx context.Context, notificationID string) (*NotificationContentResponse, error) {
	if notificationID == "" {
		return nil, errors.New("Notification ID is required")
	}

	url := fmt.Sprintf(c.API.BaseURL+"/notifications/%s/draft/content", notificationID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	bytes, err := c.API.ExecuteRequest(req)
	if err != nil {
		return nil, err
	}

	var data NotificationContentResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// PostNotificationVariations calls the POST /notifications/{notification_id}/variations endpoint of the Courier API
func (c *Client) PostNotificationVariations(ctx context.Context, notificationID string, body interface{}) error {
	bodyMap, err := toJSONMap(body)
	if err != nil {
		return err
	}

	_, err = c.API.SendRequestWithMaps(ctx, "POST", "/notifications/"+notificationID+"/variations", bodyMap)
	if err != nil {
		return err
	}

	return nil
}

// PostNotificationDraftVariations calls the POST /notifications/{notification_id}/draft/variations endpoint of the Courier API
func (c *Client) PostNotificationDraftVariations(ctx context.Context, notificationID string, body interface{}) error {
	bodyMap, err := toJSONMap(body)
	if err != nil {
		return err
	}

	_, err = c.API.SendRequestWithMaps(ctx, "POST", "/notifications/"+notificationID+"/draft/variations", bodyMap)
	if err != nil {
		return err
	}

	return nil
}

// GetNotificationSubmissionChecks calls the GET /notifications/{notification_id}/{submission_id}/checks endpoint of the Courier API
func (c *Client) GetNotificationSubmissionChecks(ctx context.Context, notificationID string, submissionID string) (*NotificationSubmissionChecksResponse, error) {
	if notificationID == "" {
		return nil, errors.New("Notification ID is required")
	}

	if submissionID == "" {
		return nil, errors.New("Submission ID is required")
	}

	url := fmt.Sprintf(c.API.BaseURL+"/notifications/%s/%s/checks", notificationID, submissionID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	bytes, err := c.API.ExecuteRequest(req)
	if err != nil {
		return nil, err
	}

	var data NotificationSubmissionChecksResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// PutNotificationSubmissionChecks calls the PUT /notifications/{notification_id}/{submission_id}/checks endpoint of the Courier API
func (c *Client) PutNotificationSubmissionChecks(ctx context.Context, notificationID string, submissionID string, body interface{}) error {
	bodyMap, err := toJSONMap(body)
	if err != nil {
		return err
	}

	_, err = c.API.SendRequestWithMaps(ctx, "PUT", "/notifications/"+notificationID+"/"+submissionID+"/checks", bodyMap)
	if err != nil {
		return err
	}

	return nil
}

// CancelNotificationSubmission calls the DELETE /notifications/{notification_id}/{submission_id}/checks endpoint of the Courier API
func (c *Client) CancelNotificationSubmission(ctx context.Context, notificationID string, submissionID string) error {
	if notificationID == "" {
		return errors.New("Notification ID is required")
	}

	if submissionID == "" {
		return errors.New("Submission ID is required")
	}

	url := fmt.Sprintf(c.API.BaseURL+"/notifications/%s/%s/checks", notificationID, submissionID)

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return err
	}

	_, err = c.API.ExecuteRequest(req)
	if err != nil {
		return err
	}

	return nil
}
