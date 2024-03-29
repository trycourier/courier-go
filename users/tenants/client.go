// This file was auto-generated by Fern from our API Definition.

package tenants

import (
	context "context"
	fmt "fmt"
	core "github.com/trycourier/courier-go/v3/core"
	option "github.com/trycourier/courier-go/v3/option"
	users "github.com/trycourier/courier-go/v3/users"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	return &Client{
		baseURL: options.BaseURL,
		caller: core.NewCaller(
			&core.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header: options.ToHeader(),
	}
}

// This endpoint is used to add a user to
// multiple tenants in one call.
// A custom profile can also be supplied for each tenant.
// This profile will be merged with the user's main
// profile when sending to the user with that tenant.
func (c *Client) AddMultple(
	ctx context.Context,
	// The user's ID. This can be any uniquely identifiable string.
	userId string,
	request *users.AddUserToMultipleTenantsParams,
	opts ...option.RequestOption,
) error {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.courier.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"users/%v/tenants", userId)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:         endpointURL,
			Method:      http.MethodPut,
			MaxAttempts: options.MaxAttempts,
			Headers:     headers,
			Client:      options.HTTPClient,
			Request:     request,
		},
	); err != nil {
		return err
	}
	return nil
}

// This endpoint is used to add a single tenant.
//
// A custom profile can also be supplied with the tenant.
// This profile will be merged with the user's main profile
// when sending to the user with that tenant.
func (c *Client) Add(
	ctx context.Context,
	// Id of the user to be added to the supplied tenant.
	userId string,
	// Id of the tenant the user should be added to.
	tenantId string,
	request *users.AddUserToSingleTenantsParams,
	opts ...option.RequestOption,
) error {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.courier.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"users/%v/tenants/%v", userId, tenantId)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:         endpointURL,
			Method:      http.MethodPut,
			MaxAttempts: options.MaxAttempts,
			Headers:     headers,
			Client:      options.HTTPClient,
			Request:     request,
		},
	); err != nil {
		return err
	}
	return nil
}

// Removes a user from any tenants they may have been associated with.
func (c *Client) RemoveAll(
	ctx context.Context,
	// Id of the user to be removed from the supplied tenant.
	userId string,
	opts ...option.RequestOption,
) error {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.courier.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"users/%v/tenants", userId)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:         endpointURL,
			Method:      http.MethodDelete,
			MaxAttempts: options.MaxAttempts,
			Headers:     headers,
			Client:      options.HTTPClient,
		},
	); err != nil {
		return err
	}
	return nil
}

// Removes a user from the supplied tenant.
func (c *Client) Remove(
	ctx context.Context,
	// Id of the user to be removed from the supplied tenant.
	userId string,
	// Id of the tenant the user should be removed from.
	tenantId string,
	opts ...option.RequestOption,
) error {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.courier.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"users/%v/tenants/%v", userId, tenantId)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:         endpointURL,
			Method:      http.MethodDelete,
			MaxAttempts: options.MaxAttempts,
			Headers:     headers,
			Client:      options.HTTPClient,
		},
	); err != nil {
		return err
	}
	return nil
}

// Returns a paginated list of user tenant associations.
func (c *Client) List(
	ctx context.Context,
	// Id of the user to retrieve all associated tenants for.
	userId string,
	request *users.ListTenantsForUserParams,
	opts ...option.RequestOption,
) (*users.ListTenantsForUserResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.courier.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"users/%v/tenants", userId)

	queryParams, err := core.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *users.ListTenantsForUserResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:         endpointURL,
			Method:      http.MethodGet,
			MaxAttempts: options.MaxAttempts,
			Headers:     headers,
			Client:      options.HTTPClient,
			Response:    &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
