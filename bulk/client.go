// This file was auto-generated by Fern from our API Definition.

package bulk

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	errors "errors"
	fmt "fmt"
	v3 "github.com/trycourier/courier-go/v3"
	core "github.com/trycourier/courier-go/v3/core"
	option "github.com/trycourier/courier-go/v3/option"
	io "io"
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
		caller:  core.NewCaller(options.HTTPClient),
		header:  options.ToHeader(),
	}
}

func (c *Client) CreateJob(
	ctx context.Context,
	request *v3.BulkCreateJobParams,
	opts ...option.IdempotentRequestOption,
) (*v3.BulkCreateJobResponse, error) {
	options := core.NewIdempotentRequestOptions(opts...)

	baseURL := "https://api.courier.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/" + "bulk"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(v3.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *v3.BulkCreateJobResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodPost,
			Headers:      headers,
			Client:       options.HTTPClient,
			Request:      request,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Ingest user data into a Bulk Job
func (c *Client) IngestUsers(
	ctx context.Context,
	// A unique identifier representing the bulk job
	jobId string,
	request *v3.BulkIngestUsersParams,
	opts ...option.IdempotentRequestOption,
) error {
	options := core.NewIdempotentRequestOptions(opts...)

	baseURL := "https://api.courier.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"bulk/%v", jobId)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:     endpointURL,
			Method:  http.MethodPost,
			Headers: headers,
			Client:  options.HTTPClient,
			Request: request,
		},
	); err != nil {
		return err
	}
	return nil
}

// Run a bulk job
func (c *Client) RunJob(
	ctx context.Context,
	// A unique identifier representing the bulk job
	jobId string,
	opts ...option.IdempotentRequestOption,
) error {
	options := core.NewIdempotentRequestOptions(opts...)

	baseURL := "https://api.courier.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"bulk/%v/run", jobId)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(v3.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodPost,
			Headers:      headers,
			Client:       options.HTTPClient,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return err
	}
	return nil
}

// Get a bulk job
func (c *Client) GetJob(
	ctx context.Context,
	// A unique identifier representing the bulk job
	jobId string,
	opts ...option.RequestOption,
) (*v3.BulkGetJobResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.courier.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"bulk/%v", jobId)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(v3.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *v3.BulkGetJobResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodGet,
			Headers:      headers,
			Client:       options.HTTPClient,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Get Bulk Job Users
func (c *Client) GetUsers(
	ctx context.Context,
	// A unique identifier representing the bulk job
	jobId string,
	opts ...option.RequestOption,
) (*v3.BulkGetJobUsersResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.courier.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"bulk/%v/users", jobId)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(v3.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *v3.BulkGetJobUsersResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodGet,
			Headers:      headers,
			Client:       options.HTTPClient,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
