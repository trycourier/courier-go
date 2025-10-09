// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/trycourier/courier-go/v3/internal/apijson"
	"github.com/trycourier/courier-go/v3/internal/apiquery"
	"github.com/trycourier/courier-go/v3/internal/requestconfig"
	"github.com/trycourier/courier-go/v3/option"
	"github.com/trycourier/courier-go/v3/packages/param"
	"github.com/trycourier/courier-go/v3/packages/respjson"
	"github.com/trycourier/courier-go/v3/shared"
)

// BulkService contains methods and other services that help with interacting with
// the courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBulkService] method instead.
type BulkService struct {
	Options []option.RequestOption
}

// NewBulkService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBulkService(opts ...option.RequestOption) (r BulkService) {
	r = BulkService{}
	r.Options = opts
	return
}

// Ingest user data into a Bulk Job
func (r *BulkService) AddUsers(ctx context.Context, jobID string, body BulkAddUsersParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return
	}
	path := fmt.Sprintf("bulk/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Create a bulk job
func (r *BulkService) NewJob(ctx context.Context, body BulkNewJobParams, opts ...option.RequestOption) (res *BulkNewJobResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "bulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get Bulk Job Users
func (r *BulkService) ListUsers(ctx context.Context, jobID string, query BulkListUsersParams, opts ...option.RequestOption) (res *BulkListUsersResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return
	}
	path := fmt.Sprintf("bulk/%s/users", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get a bulk job
func (r *BulkService) GetJob(ctx context.Context, jobID string, opts ...option.RequestOption) (res *BulkGetJobResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return
	}
	path := fmt.Sprintf("bulk/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Run a bulk job
func (r *BulkService) RunJob(ctx context.Context, jobID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return
	}
	path := fmt.Sprintf("bulk/%s/run", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

type BulkNewJobResponse struct {
	JobID string `json:"jobId,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		JobID       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BulkNewJobResponse) RawJSON() string { return r.JSON.raw }
func (r *BulkNewJobResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BulkListUsersResponse struct {
	Items  []BulkListUsersResponseItem `json:"items,required"`
	Paging shared.Paging               `json:"paging,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BulkListUsersResponse) RawJSON() string { return r.JSON.raw }
func (r *BulkListUsersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BulkListUsersResponseItem struct {
	// Any of "PENDING", "ENQUEUED", "ERROR".
	Status    string `json:"status,required"`
	MessageID string `json:"messageId,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		MessageID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.InboundBulkMessageUser
}

// Returns the unmodified JSON received from the API
func (r BulkListUsersResponseItem) RawJSON() string { return r.JSON.raw }
func (r *BulkListUsersResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BulkGetJobResponse struct {
	Job BulkGetJobResponseJob `json:"job,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Job         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BulkGetJobResponse) RawJSON() string { return r.JSON.raw }
func (r *BulkGetJobResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BulkGetJobResponseJob struct {
	Definition shared.InboundBulkMessageUnion `json:"definition,required"`
	Enqueued   int64                          `json:"enqueued,required"`
	Failures   int64                          `json:"failures,required"`
	Received   int64                          `json:"received,required"`
	// Any of "CREATED", "PROCESSING", "COMPLETED", "ERROR".
	Status string `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Definition  respjson.Field
		Enqueued    respjson.Field
		Failures    respjson.Field
		Received    respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BulkGetJobResponseJob) RawJSON() string { return r.JSON.raw }
func (r *BulkGetJobResponseJob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BulkAddUsersParams struct {
	Users []shared.InboundBulkMessageUserParam `json:"users,omitzero,required"`
	paramObj
}

func (r BulkAddUsersParams) MarshalJSON() (data []byte, err error) {
	type shadow BulkAddUsersParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BulkAddUsersParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BulkNewJobParams struct {
	Message shared.InboundBulkMessageUnionParam `json:"message,omitzero,required"`
	paramObj
}

func (r BulkNewJobParams) MarshalJSON() (data []byte, err error) {
	type shadow BulkNewJobParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BulkNewJobParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BulkListUsersParams struct {
	// A unique identifier that allows for fetching the next set of users added to the
	// bulk job
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BulkListUsersParams]'s query parameters as `url.Values`.
func (r BulkListUsersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
