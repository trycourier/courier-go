package courier

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type SingleFilter struct {
	Operator string `json:"operator"`
	Path     string `json:"path"`
	Value    string `json:"value"`
}

type NestedFilter struct {
	Operator string         `json:"operator"`
	Rules    []SingleFilter `json:"rules"`
}

type Audience struct {
	Description string      `json:"description,omitempty"`
	Filter      interface{} `json:"filter"`
	Name        string      `json:"name,omitempty"`
}

type AudienceResponseBody struct {
	Id          string      `json:"id"`
	Description string      `json:"description"`
	CreatedAt   string      `json:"created_at"`
	Filter      interface{} `json:"filter"`
	Name        string      `json:"name"`
	UpdatedAt   string      `json:"updated_at"`
}

type AudienceResponse struct {
	Audience AudienceResponseBody `json:"audience"`
}

type AudienceMember struct {
	AudienceId      string `json:"audience_id"`
	AddedAt         string `json:"added_at"`
	AudienceVersion int    `json:"audience_version"`
	MemberId        string `json:"member_id"`
	Reason          string `json:"reason"`
}

type GetAudienceMembersResponse struct {
	Items  []*AudienceMember `json:"items"`
	Paging *PagingResponse   `json:"paging"`
}

type GetAudiencesResponse struct {
	Items  []*AudienceResponseBody `json:"items"`
	Paging *PagingResponse
}

func (c *Client) PutAudience(ctx context.Context, audienceId string, audience Audience) (*AudienceResponse, error) {
	if audienceId == "" {
		return nil, errors.New("Audience ID is required")
	}

	_, singleFilterOk := audience.Filter.(SingleFilter)
	_, nestedFilterOk := audience.Filter.(NestedFilter)

	if !singleFilterOk && !nestedFilterOk {
		return nil, errors.New("Audience filter must be of type SingleFilter or NestFilter")
	}

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(audience)

	jsonBody, err := json.Marshal(audience)

	if err != nil {
		return nil, err
	}

	bytes, err := c.API.SendRequestWithBytes(ctx, http.MethodPut, fmt.Sprintf("/audiences/%s", audienceId), jsonBody)

	if err != nil {
		return nil, err
	}

	var data AudienceResponse

	err = json.Unmarshal(bytes, &data)

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (c *Client) GetAudience(ctx context.Context, audienceId string) (*AudienceResponseBody, error) {
	if audienceId == "" {
		return nil, errors.New("Audience ID is required")
	}

	url := fmt.Sprintf(c.API.BaseURL+"/audiences/%s", audienceId)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

	if err != nil {
		return nil, err
	}

	bytes, err := c.API.ExecuteRequest(req)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("AudienceId %s, not found", audienceId))
	}

	var data AudienceResponseBody

	err = json.Unmarshal(bytes, &data)

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (c *Client) GetAudienceMembers(ctx context.Context, audienceId string, cursor string) (*GetAudienceMembersResponse, error) {
	if audienceId == "" {
		return nil, errors.New("Audience ID is required")
	}

	url := fmt.Sprintf(c.API.BaseURL+"/audiences/%s/members", audienceId)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	if cursor != "" {
		q.Add("cursor", cursor)
	}

	bytes, err := c.API.ExecuteRequest(req)

	if err != nil {
		return nil, err
	}

	var data GetAudienceMembersResponse

	err = json.Unmarshal(bytes, &data)

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (c *Client) GetAudiences(ctx context.Context, cursor string) (*GetAudiencesResponse, error) {

	url := fmt.Sprintf(c.API.BaseURL + "/audiences")

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	if cursor != "" {
		q.Add("cursor", cursor)
	}

	bytes, err := c.API.ExecuteRequest(req)

	if err != nil {
		return nil, err
	}

	var data GetAudiencesResponse

	err = json.Unmarshal(bytes, &data)

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (c *Client) DeleteAudience(ctx context.Context, audienceId string) error {
	if audienceId == "" {
		return errors.New("Audience ID is required")
	}

	url := fmt.Sprintf(c.API.BaseURL+"/audiences/%s", audienceId)

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)

	if err != nil {
		return err
	}

	_, err = c.API.ExecuteRequest(req)

	if err != nil {
		return errors.New(fmt.Sprintf("AudienceId %s, not found", audienceId))
	}

	return nil
}
