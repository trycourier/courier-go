// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/trycourier/courier-go/v4/internal/apiquery"
	"github.com/trycourier/courier-go/v4/internal/requestconfig"
	"github.com/trycourier/courier-go/v4/option"
	"github.com/trycourier/courier-go/v4/packages/param"
)

// DigestScheduleService contains methods and other services that help with
// interacting with the Courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDigestScheduleService] method instead.
type DigestScheduleService struct {
	Options []option.RequestOption
}

// NewDigestScheduleService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDigestScheduleService(opts ...option.RequestOption) (r DigestScheduleService) {
	r = DigestScheduleService{}
	r.Options = opts
	return
}

// List the digest instances for a schedule. Each instance represents the events
// accumulated for a single user against the schedule, and can be used to monitor
// digest accumulation before the digest is released.
func (r *DigestScheduleService) ListInstances(ctx context.Context, scheduleID string, query DigestScheduleListInstancesParams, opts ...option.RequestOption) (res *DigestInstanceListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if scheduleID == "" {
		err = errors.New("missing required schedule_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("digests/schedules/%s/instances", scheduleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Send a digest now instead of waiting for its scheduled time, so your users get
// what they have collected so far right away.
func (r *DigestScheduleService) Release(ctx context.Context, scheduleID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if scheduleID == "" {
		err = errors.New("missing required schedule_id parameter")
		return err
	}
	path := fmt.Sprintf("digests/schedules/%s/trigger", scheduleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

type DigestScheduleListInstancesParams struct {
	// A cursor token from a previous response, used to fetch the next page of results.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// The maximum number of digest instances to return. Defaults to 20, with a maximum
	// of 100.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DigestScheduleListInstancesParams]'s query parameters as
// `url.Values`.
func (r DigestScheduleListInstancesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
