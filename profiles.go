package courier

import (
	"bytes"
	"context"
	"net/http"
)

// Profile represents the return of the Courier profiles API
type Profile struct {
	Profile interface{} `json:"profile"`
}

// GetProfile calls the GET /profiles/:id endpoint of the Courier API
func (c *Client) GetProfile(ctx context.Context, id string) (*Profile, error) {
	var response Profile
	err := c.http.SendRequestWithJSON(ctx, "GET", "/profiles/"+id, nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// MergeProfile calls the POST /profiles/:id endpoint of the Courier API
func (c *Client) MergeProfile(ctx context.Context, id string, profile []byte) error {
	req, err := http.NewRequest("POST", c.http.BaseURL+"/profiles/"+id, bytes.NewBuffer(profile))
	if err != nil {
		return err
	}
	_, err = c.http.ExecuteRequest(req)
	return err
}

// UpdateProfile calls the PUT /profiles/:id endpoint of the Courier API
func (c *Client) UpdateProfile(ctx context.Context, id string, profile []byte) error {
	req, err := http.NewRequest("PUT", c.http.BaseURL+"/profiles/"+id, bytes.NewBuffer(profile))
	if err != nil {
		return err
	}
	_, err = c.http.ExecuteRequest(req)
	return err
}
