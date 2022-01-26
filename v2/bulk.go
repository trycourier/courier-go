package courier

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Job can be used to cast GET /bulk/jobId response to
type Job struct {
	Definition interface{} `json:"definition"`
	Enqueued   int         `json:"enqueued"`
	Failures   int         `json:"failures"`
	Received   int         `json:"received"`
	Status     string      `json:"status"`
}

// CreateJobBody can be used as request body to POST /bulk
type CreateJobBody struct {
	Message interface{} `json:"message"`
}

// IngestJobBody can be used as request body to POST /bulk/:jobId
type IngestJobBody struct {
	Users []interface{} `json:"users"`
}

// GetBulkJobResponse is type for the GET /bulk/:jobId response
type GetBulkJobResponse struct {
	Job interface{} `json:"job"`
}

// GetBulkJobUsersResponse is type for the GET /bulk/:jobId/users response
type GetBulkJobUsersResponse struct {
	Items  []interface{} `json:"items"`
	Paging *PagingResponse
}

// IngestBulkJobResponse can be used to deserialize POST /bulk/:jobId response
type IngestBulkJobResponse struct {
	Errors []interface{} `json:"errors"`
	Total  int           `json:"total"`
}

// CreateJob calls the POST /bulk endpoint of the Courier API
func (c *Client) CreateJob(ctx context.Context, body interface{}) (string, error) {
	bodyMap, err := toJSONMap(body)
	if err != nil {
		return "", err
	}

	response, err := c.API.SendRequestWithMaps(ctx, "POST", "/bulk", bodyMap)
	if err != nil {
		return "", err
	}

	var jobID string
	err = json.Unmarshal(response["jobId"], &jobID)
	if err != nil {
		return "", err
	}
	return jobID, nil
}

// IngestJob calls the POST /bulk/:jobId endpoint of the Courier API
func (c *Client) IngestJob(ctx context.Context, jobID string, body interface{}) (*IngestBulkJobResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	bytes, err := c.API.SendRequestWithBytes(ctx, "POST", "/bulk/"+jobID, jsonBody)
	if err != nil {
		return nil, err
	}

	var data IngestBulkJobResponse

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// RunJob calls the POST /bulk/:jobId/run endpoint of the Courier API
func (c *Client) RunJob(ctx context.Context, jobID string) error {
	_, err := c.API.SendRequestWithBytes(ctx, "POST", "/bulk/"+jobID+"/run", nil)
	return err
}

// GetJob calls the GET /bulk/:jobId endpoint of the Courier API
func (c *Client) GetJob(ctx context.Context, jobID string) (*GetBulkJobResponse, error) {
	if jobID == "" {
		return nil, errors.New("Job ID is required")
	}
	url := fmt.Sprintf(c.API.BaseURL+"/bulk/%s", jobID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	bytes, err := c.API.ExecuteRequest(req)
	if err != nil {
		return nil, err
	}

	var data GetBulkJobResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// GetJobUsers calls the GET /bulk/:jobId/users endpoint of the Courier API
func (c *Client) GetJobUsers(ctx context.Context, jobID string, cursor string) (*GetBulkJobUsersResponse, error) {
	if jobID == "" {
		return nil, errors.New("Job ID is required")
	}

	url := fmt.Sprintf(c.API.BaseURL+"/bulk/%s/users", jobID)

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

	var data GetBulkJobUsersResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
