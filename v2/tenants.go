package courier

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Account is the type for account entity
type Tenant struct {
	Id                 string      `json:"id"`
	Name               string      `json:"name"`
	ParentTenantId     string      `json:"parent_tenant_id,omitempty"`
	DefaultPreferences interface{} `json:"default_preferences,omitempty"`
	Properties         interface{} `json:"properties,omitempty"`
	UserProfile        interface{} `json:"user_profile,omitempty"`
	BrandId            string      `json:"brand_id,omitempty"`
}

// ListTenantsResponse is type for the GET /tenants response
type ListTenantsResponse struct {
	Items   []Tenant `json:"items"`
	HasMore bool     `json:"has_more"`
	Url     string   `json:"url"`
	NextUrl string   `json:"next_url,omitempty"`
}

// GetTenant calls the GET /tenants/:tenantId endpoint of the Courier API
func (c *Client) GetTenant(ctx context.Context, tenantId string) (*Tenant, error) {
	if tenantId == "" {
		return nil, errors.New("tenant ID is required")
	}
	url := fmt.Sprintf(c.API.BaseURL+"/tenants/%s", tenantId)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	bytes, err := c.API.ExecuteRequest(req)
	if err != nil {
		return nil, err
	}

	var data Tenant
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// GetAccounts calls the GET /tenants endpoint of the Courier API
func (c *Client) GetTenants(ctx context.Context, limit string, cursor string) (*ListTenantsResponse, error) {
	url := fmt.Sprintf(c.API.BaseURL + "/tenants")

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	if limit != "" {
		q.Add("limit", limit)
	}
	if cursor != "" {
		q.Add("cursor", cursor)
	}
	req.URL.RawQuery = q.Encode()

	bytes, err := c.API.ExecuteRequest(req)
	if err != nil {
		return nil, err
	}

	var data ListTenantsResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// Calls the PUT /tenants/:tenantId endpoint of the Courier API
func (c *Client) PutTenant(ctx context.Context, tenantId string, body interface{}) (*Tenant, error) {
	if tenantId == "" {
		return nil, errors.New("Tenant ID is required")
	}

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(body)

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	bytes, err := c.API.SendRequestWithBytes(ctx, http.MethodPut, fmt.Sprintf("/tenants/%s", tenantId), jsonBody)

	if err != nil {
		return nil, err
	}

	var data Tenant

	err = json.Unmarshal(bytes, &data)

	if err != nil {
		return nil, err
	}

	return &data, nil
}

// DeleteAccount calls the DELETE /accounts/:account_id endpoint of the Courier API
func (c *Client) DeleteTenant(ctx context.Context, tenantId string) error {
	if tenantId == "" {
		return errors.New("Tenant ID is required")
	}

	url := fmt.Sprintf(c.API.BaseURL+"/tenants/%s", tenantId)

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
