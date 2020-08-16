package courier

import (
	"context"
	"encoding/json"
)

// GetProfile calls the GET /profiles/:id endpoint of the Courier API
func (c *Client) GetProfile(ctx context.Context, id string) (map[string]interface{}, error) {
	type Response struct {
		Profile map[string]interface{}
	}

	bytes, err := c.API.SendRequestWithBytes(ctx, "GET", "/profiles/"+id, nil)

	var response Response
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return nil, err
	}

	return response.Profile, nil
}

// MergeProfileBytes calls the POST /profiles/:id endpoint of the Courier API
func (c *Client) MergeProfileBytes(ctx context.Context, id string, profile []byte) error {
	_, err := c.API.SendRequestWithBytes(ctx, "POST", "/profiles/"+id, profile)
	return err
}

// UpdateProfileBytes calls the PUT /profiles/:id endpoint of the Courier API
func (c *Client) UpdateProfileBytes(ctx context.Context, id string, profile []byte) error {
	_, err := c.API.SendRequestWithBytes(ctx, "PUT", "/profiles/"+id, profile)
	return err
}
