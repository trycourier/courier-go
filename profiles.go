package courier

import (
	"bytes"
	"context"
	"encoding/json"
)

// GetProfile calls the GET /profiles/:id endpoint of the Courier API
func (c *Client) GetProfile(ctx context.Context, id string) (map[string]json.RawMessage, error) {
	response, err := c.http.SendRequestWithMaps(ctx, "GET", "/profiles/"+id, nil)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// MergeProfile calls the POST /profiles/:id endpoint of the Courier API
func (c *Client) MergeProfile(ctx context.Context, id string, profile []byte) error {
	_, err := c.http.SendRequestWithBytes(ctx, "POST", "/profiles/"+id, bytes.NewBuffer(profile))
	return err
}

// UpdateProfile calls the PUT /profiles/:id endpoint of the Courier API
func (c *Client) UpdateProfile(ctx context.Context, id string, profile []byte) error {
	_, err := c.http.SendRequestWithBytes(ctx, "PUT", "/profiles/"+id, bytes.NewBuffer(profile))
	return err
}
