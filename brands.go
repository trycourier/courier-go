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
	items []*BrandSnippetItem `json:"items,omitempty"`
}

type BrandResponse struct {
	Id        string
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

type CreateBrandRequest struct {
	Id       string        `json:"id,omitempty"`
	Name     string        `json:"name"`
	Settings BrandSettings `json:"settings"`
	Snippets BrandSnippets `json:"snippets,omitempty"`
}

type UpdateBrandRequest struct {
	Name     string        `json:"name"`
	Settings BrandSettings `json:"settings"`
	Snippets BrandSnippets `json:"snippets,omitempty"`
}

func (c *Client) GetBrands(ctx context.Context, cursor string) (*BrandsResponse, error) {
	url := c.BaseUrl + "/brands"

	if cursor != "" {
		url = url + "?cursor=" + cursor
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	bytes, err := c.doRequest(req)
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

func (c *Client) GetBrand(ctx context.Context, brandID string) (*BrandResponse, error) {
	if brandID == "" {
		return nil, errors.New("brandID is required")
	}

	url := fmt.Sprintf(c.BaseUrl+"/brands/%s", brandID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	bytes, err := c.doRequest(req)
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

func (c *Client) PostBrand(ctx context.Context, brandID string, brandName string, brandSettings BrandSettings, brandSnippets BrandSnippets, idempotencyKey string) (*BrandResponse, error) {
	if brandName == "" {
		return nil, errors.New("Brand Name is required")
	}

	payload := CreateBrandRequest{
		Id:       brandID,
		Name:     brandName,
		Settings: brandSettings,
		Snippets: brandSnippets,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	url := c.BaseUrl + "/brands"

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	// Idempotent Request check
	if idempotencyKey != "" {
		req.Header.Set("Idempotency-Key", idempotencyKey)
	}

	bytes, err := c.doRequest(req)
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

func (c *Client) PutBrand(ctx context.Context, brandID string, brandName string, brandSettings BrandSettings, brandSnippets BrandSnippets) (*BrandResponse, error) {
	if brandID == "" {
		return nil, errors.New("Brand ID is required")
	}

	if brandName == "" {
		return nil, errors.New("Brand Name is required")
	}

	payload := UpdateBrandRequest{
		Name:     brandName,
		Settings: brandSettings,
		Snippets: brandSnippets,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(c.BaseUrl+"/brands/%s", brandID)

	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	bytes, err := c.doRequest(req)
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

func (c *Client) DeleteBrand(ctx context.Context, brandID string) error {
	if brandID == "" {
		return errors.New("brandID is required")
	}

	url := fmt.Sprintf(c.BaseUrl+"/brands/%s", brandID)

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return err
	}

	bytes, err := c.doRequest(req)
	if err != nil {
		return err
	}

	var data string
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return err
	}

	return nil
}
