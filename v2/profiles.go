package courier

import (
	"context"
	"encoding/json"
)

// GetProfile calls the GET /profiles/:id endpoint of the Courier API
func (c *Client) GetProfile(ctx context.Context, id string) (map[string]json.RawMessage, error) {
	return c.API.SendRequestWithMaps(ctx, "GET", "/profiles/"+id, nil)
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
