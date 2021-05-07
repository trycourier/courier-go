package courier

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type BrandColors struct {
	Primary   string `json:"primary,omitempty"`
	Secondary string `json:"secondary,omitempty"`
	Tertiary  string `json:"tertiary,omitempty"`
}

type BrandEmail struct {
	Header interface{} `json:"header,omitempty"`
	Footer interface{} `json:"footer,omitempty"`
}

type BrandSettings struct {
	Colors BrandColors `json:"colors,omitempty"`
	Email  BrandEmail  `json:"email,omitempty"`
}

type BrandSnippetItem struct {
	Format string `json:"format,omitempty"`
	Name   string `json:"name,omitempty"`
	Value  string `json:"value,omitempty"`
}

type BrandSnippets struct {
	Items []*BrandSnippetItem `json:"items,omitempty"`
}

type BrandResponse struct {
	ID        string
	Name      string
	Created   int64
	Published int64
	Updated   int64
	Settings  BrandSettings
	Snippets  BrandSnippets
	Version   string
}

type BrandsResponse struct {
	Paging  *PagingResponse
	Results []*BrandResponse
}

type PostBrandBody struct {
	ID       string        `json:"id,omitempty"`
	Name     string        `json:"name"`
	Settings BrandSettings `json:"settings"`
	Snippets BrandSnippets `json:"snippets,omitempty"`
}

type PutBrandBody struct {
	Name     string        `json:"name"`
	Settings BrandSettings `json:"settings"`
	Snippets BrandSnippets `json:"snippets,omitempty"`
}

// GetBrands calls the GET /brands endpoint of the Courier API
func (c *Client) GetBrands(ctx context.Context, cursor string) (*BrandsResponse, error) {
	url := c.API.BaseURL + "/brands"

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

	var data BrandsResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// GetBrand calls the GET /brands/{brand_id} endpoint of the Courier API
func (c *Client) GetBrand(ctx context.Context, brandID string) (*BrandResponse, error) {
	if brandID == "" {
		return nil, errors.New("Brand ID is required")
	}

	url := fmt.Sprintf(c.API.BaseURL+"/brands/%s", brandID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	bytes, err := c.API.ExecuteRequest(req)
	if err != nil {
		return nil, err
	}

	var data BrandResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// PostBrand calls the POST /brands endpoint of the Courier API
func (c *Client) PostBrand(ctx context.Context, body PostBrandBody) (*BrandResponse, error) {
	payload, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	url := c.API.BaseURL + "/brands"

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}

	bytes, err := c.API.ExecuteRequest(req)
	if err != nil {
		return nil, err
	}

	var data BrandResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// PutBrand calls the PUT /brands/{brand_id} endpoint of the Courier API
func (c *Client) PutBrand(ctx context.Context, brandID string, body PutBrandBody) (*BrandResponse, error) {
	if brandID == "" {
		return nil, errors.New("Brand ID is required")
	}

	payload, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(c.API.BaseURL+"/brands/%s", brandID)

	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}

	bytes, err := c.API.ExecuteRequest(req)
	if err != nil {
		return nil, err
	}

	var data BrandResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// DeleteBrand calls the DELETE /brands/{brand_id} endpoint of the Courier API
func (c *Client) DeleteBrand(ctx context.Context, brandID string) error {
	if brandID == "" {
		return errors.New("Brand ID is required")
	}

	url := fmt.Sprintf(c.API.BaseURL+"/brands/%s", brandID)

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
