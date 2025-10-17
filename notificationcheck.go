// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/trycourier/courier-go/v3/internal/apijson"
	"github.com/trycourier/courier-go/v3/internal/requestconfig"
	"github.com/trycourier/courier-go/v3/option"
	"github.com/trycourier/courier-go/v3/packages/param"
	"github.com/trycourier/courier-go/v3/packages/respjson"
)

// NotificationCheckService contains methods and other services that help with
// interacting with the courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNotificationCheckService] method instead.
type NotificationCheckService struct {
	Options []option.RequestOption
}

// NewNotificationCheckService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNotificationCheckService(opts ...option.RequestOption) (r NotificationCheckService) {
	r = NotificationCheckService{}
	r.Options = opts
	return
}

func (r *NotificationCheckService) Update(ctx context.Context, submissionID string, params NotificationCheckUpdateParams, opts ...option.RequestOption) (res *NotificationCheckUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if submissionID == "" {
		err = errors.New("missing required submissionId parameter")
		return
	}
	path := fmt.Sprintf("notifications/%s/%s/checks", params.ID, submissionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

func (r *NotificationCheckService) List(ctx context.Context, submissionID string, query NotificationCheckListParams, opts ...option.RequestOption) (res *NotificationCheckListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ID == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if submissionID == "" {
		err = errors.New("missing required submissionId parameter")
		return
	}
	path := fmt.Sprintf("notifications/%s/%s/checks", query.ID, submissionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *NotificationCheckService) Delete(ctx context.Context, submissionID string, body NotificationCheckDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.ID == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if submissionID == "" {
		err = errors.New("missing required submissionId parameter")
		return
	}
	path := fmt.Sprintf("notifications/%s/%s/checks", body.ID, submissionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type NotificationCheckUpdateResponse struct {
	Checks []Check `json:"checks,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Checks      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationCheckUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *NotificationCheckUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationCheckListResponse struct {
	Checks []Check `json:"checks,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Checks      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationCheckListResponse) RawJSON() string { return r.JSON.raw }
func (r *NotificationCheckListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationCheckUpdateParams struct {
	ID     string           `path:"id,required" json:"-"`
	Checks []BaseCheckParam `json:"checks,omitzero,required"`
	paramObj
}

func (r NotificationCheckUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow NotificationCheckUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NotificationCheckUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationCheckListParams struct {
	ID string `path:"id,required" json:"-"`
	paramObj
}

type NotificationCheckDeleteParams struct {
	ID string `path:"id,required" json:"-"`
	paramObj
}
