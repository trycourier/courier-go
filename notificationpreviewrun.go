// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/trycourier/courier-go/v4/internal/apijson"
	"github.com/trycourier/courier-go/v4/internal/apiquery"
	shimjson "github.com/trycourier/courier-go/v4/internal/encoding/json"
	"github.com/trycourier/courier-go/v4/internal/requestconfig"
	"github.com/trycourier/courier-go/v4/option"
	"github.com/trycourier/courier-go/v4/packages/param"
	"github.com/trycourier/courier-go/v4/packages/respjson"
	"github.com/trycourier/courier-go/v4/shared"
)

// Render a template's email content on real email clients and read back the
// screenshots, so you can check how it looks before you send it.
//
// NotificationPreviewRunService contains methods and other services that help with
// interacting with the Courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNotificationPreviewRunService] method instead.
type NotificationPreviewRunService struct {
	Options []option.RequestOption
}

// NewNotificationPreviewRunService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNotificationPreviewRunService(opts ...option.RequestOption) (r NotificationPreviewRunService) {
	r = NotificationPreviewRunService{}
	r.Options = opts
	return
}

// Render this template's email content on each of the requested devices.
//
// Returns as soon as the run exists and its render is queued — the screenshots are
// produced asynchronously. Poll
// `GET /notifications/{id}/previews/runs/{previewRunId}` until every result
// reaches a terminal status.
//
// Name the devices either with `device_set_id`, for a saved set, or with
// `device_ids`, for a one-off list. Exactly one of the two is required. Inline
// `device_ids` must be ids listed by `GET /previews/devices`; any other id is a
// 422, refused before the run exists or is billed.
//
// A template that does not exist is a 404. One that exists but cannot be previewed
// — not a Design Studio template, no email channel, or no such `template_version`
// — is a 422, also refused before the run exists or is billed.
//
// Preview runs are a metered add-on. A workspace without it, or with its billing
// suspended, receives a 402.
func (r *NotificationPreviewRunService) New(ctx context.Context, id string, params NotificationPreviewRunNewParams, opts ...option.RequestOption) (res *PreviewRun, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.XIdempotencyExpiration) {
		opts = append(opts, option.WithHeader("x-idempotency-expiration", fmt.Sprintf("%v", params.XIdempotencyExpiration.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("notifications/%s/previews/runs", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieve one of this template's preview runs together with its per-device
// results.
//
// A run is only readable under the template it previewed: under any other template
// it is a 404, the same as a run that does not exist.
//
// `thumbnail_url` and `screenshot_url` are short-lived signed URLs, re-signed on
// every read. Fetch them now rather than storing them. Both are null until
// Courier's own copy of the image exists, which is what `status: COMPLETED` on a
// result means.
func (r *NotificationPreviewRunService) Get(ctx context.Context, previewRunID string, query NotificationPreviewRunGetParams, opts ...option.RequestOption) (res *PreviewRunDetail, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if previewRunID == "" {
		err = errors.New("missing required previewRunId parameter")
		return nil, err
	}
	path := fmt.Sprintf("notifications/%s/previews/runs/%s", query.ID, previewRunID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List this template's preview runs, newest first. Cursor-paginated.
//
// A template that does not exist is a 404, the same as every other
// `/notifications/{id}` route.
func (r *NotificationPreviewRunService) List(ctx context.Context, id string, query NotificationPreviewRunListParams, opts ...option.RequestOption) (res *PreviewRunListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("notifications/%s/previews/runs", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Request body for creating a preview run of the template in the path. Provide
// exactly one of `device_set_id` or `device_ids`. The template is the path's
// `{id}`; a `template_id` here is an unknown key and a 400.
type CreatePreviewRunRequestParam struct {
	// A saved device set naming the devices to render on. Mutually exclusive with
	// `device_ids`.
	DeviceSetID param.Opt[string] `json:"device_set_id,omitzero"`
	// Render the template's content for this locale, e.g. "fr-FR".
	Locale param.Opt[string] `json:"locale,omitzero"`
	// Which version of the template to render. Omit for the latest saved draft, which
	// always exists and is what the editor shows. `published` renders the live
	// version; a zero-padded `v002` renders that specific publish. Versions are
	// 1-based, so `v000` is not a version, and the unpadded `v2` is rejected — that
	// spelling belongs to journeys' AutomationVersionId, a different scheme in which
	// `v0` means published.
	TemplateVersion param.Opt[string] `json:"template_version,omitzero"`
	// Template variables to render with, the same shape as the `data` object on a
	// send.
	Data map[string]any `json:"data,omitzero"`
	// The devices to render on, by `PreviewDevice.id`, for a one-off run. Mutually
	// exclusive with `device_set_id`.
	DeviceIDs []string `json:"device_ids,omitzero"`
	paramObj
}

func (r CreatePreviewRunRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreatePreviewRunRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreatePreviewRunRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One device's result within a preview run.
type PreviewResult struct {
	// The device this result is for, by `PreviewDevice.id`.
	DeviceID string `json:"device_id" api:"required"`
	// Short-lived signed URL for the full-sized image. Null until the screenshot
	// exists. Re-signed on every read, so fetch it rather than storing it.
	ScreenshotURL string `json:"screenshot_url" api:"required"`
	// One device's outcome. `COMPLETED` means the screenshot exists and its URLs are
	// populated. `UNSUPPORTED`, `TIMED_OUT` and `FAILED` are all terminal, and none
	// stands in for another — `UNSUPPORTED` means the device was retired at the
	// vendor, `TIMED_OUT` means it did not report in time.
	//
	// Any of "PENDING", "PROCESSING", "COMPLETED", "UNSUPPORTED", "TIMED_OUT",
	// "FAILED".
	Status PreviewResultStatus `json:"status" api:"required"`
	// Short-lived signed URL for the grid-sized image. Null until the screenshot
	// exists. Re-signed on every read, so fetch it rather than storing it.
	ThumbnailURL string `json:"thumbnail_url" api:"required"`
	// Why one device's render failed, when its `status` is `FAILED` and the cause has
	// a public name. `DELIVERY_FAILED` means the rendering service could not deliver
	// the message to its own capture mailbox — infrastructure, not anything wrong with
	// the template.
	//
	// Any of "DELIVERY_FAILED".
	FailureReason PreviewResultFailureReason `json:"failure_reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeviceID      respjson.Field
		ScreenshotURL respjson.Field
		Status        respjson.Field
		ThumbnailURL  respjson.Field
		FailureReason respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PreviewResult) RawJSON() string { return r.JSON.raw }
func (r *PreviewResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Why one device's render failed, when its `status` is `FAILED` and the cause has
// a public name. `DELIVERY_FAILED` means the rendering service could not deliver
// the message to its own capture mailbox — infrastructure, not anything wrong with
// the template.
type PreviewResultFailureReason string

const (
	PreviewResultFailureReasonDeliveryFailed PreviewResultFailureReason = "DELIVERY_FAILED"
)

// One device's outcome. `COMPLETED` means the screenshot exists and its URLs are
// populated. `UNSUPPORTED`, `TIMED_OUT` and `FAILED` are all terminal, and none
// stands in for another — `UNSUPPORTED` means the device was retired at the
// vendor, `TIMED_OUT` means it did not report in time.
type PreviewResultStatus string

const (
	PreviewResultStatusPending     PreviewResultStatus = "PENDING"
	PreviewResultStatusProcessing  PreviewResultStatus = "PROCESSING"
	PreviewResultStatusCompleted   PreviewResultStatus = "COMPLETED"
	PreviewResultStatusUnsupported PreviewResultStatus = "UNSUPPORTED"
	PreviewResultStatusTimedOut    PreviewResultStatus = "TIMED_OUT"
	PreviewResultStatusFailed      PreviewResultStatus = "FAILED"
)

// One render of a template across a set of devices. Billable.
type PreviewRun struct {
	// Unique identifier for the preview run.
	ID string `json:"id" api:"required"`
	// ISO-8601 timestamp of when the run was created.
	CreatedAt string `json:"created_at" api:"required"`
	// The devices this run was submitted for, snapshotted when the run was created.
	DeviceIDs []string `json:"device_ids" api:"required"`
	// Where the run itself has got to. `PENDING` and `RENDERED` mean Courier is still
	// preparing the email, `SUBMITTED` means it is with the rendering service, and
	// `COMPLETED` means every device has reported. `FAILED` is the run as a whole
	// failing — an individual device failing never fails the run.
	//
	// Any of "PENDING", "RENDERED", "SUBMITTED", "COMPLETED", "FAILED".
	Status PreviewRunStatus `json:"status" api:"required"`
	// The template that was rendered.
	TemplateID string `json:"template_id" api:"required"`
	// Why the run failed, when `status` is `FAILED`. `NO_EMAIL_CHANNEL` and
	// `TEMPLATE_NOT_SUPPORTED` mean there was nothing to render;
	// `ALL_DEVICES_UNSUPPORTED` means every requested device has been retired and the
	// request can be fixed by choosing others.
	//
	// Any of "TEMPLATE_NOT_SUPPORTED", "NO_EMAIL_CHANNEL", "RENDER_FAILED",
	// "ALL_DEVICES_UNSUPPORTED", "VENDOR_ERROR".
	FailureReason PreviewRunFailureReason `json:"failure_reason"`
	// The version of the template that was rendered — `draft`, or a zero-padded
	// published version such as `v002`. Absent until the render settles.
	TemplateVersion string `json:"template_version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		DeviceIDs       respjson.Field
		Status          respjson.Field
		TemplateID      respjson.Field
		FailureReason   respjson.Field
		TemplateVersion respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PreviewRun) RawJSON() string { return r.JSON.raw }
func (r *PreviewRun) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A preview run together with its per-device results.
type PreviewRunDetail struct {
	// Unique identifier for the preview run.
	ID string `json:"id" api:"required"`
	// ISO-8601 timestamp of when the run was created.
	CreatedAt string `json:"created_at" api:"required"`
	// The devices this run was submitted for, snapshotted when the run was created.
	DeviceIDs []string `json:"device_ids" api:"required"`
	// One entry per device in `device_ids`.
	Results []PreviewResult `json:"results" api:"required"`
	// Where the run itself has got to. `PENDING` and `RENDERED` mean Courier is still
	// preparing the email, `SUBMITTED` means it is with the rendering service, and
	// `COMPLETED` means every device has reported. `FAILED` is the run as a whole
	// failing — an individual device failing never fails the run.
	//
	// Any of "PENDING", "RENDERED", "SUBMITTED", "COMPLETED", "FAILED".
	Status PreviewRunStatus `json:"status" api:"required"`
	// The template that was rendered.
	TemplateID string `json:"template_id" api:"required"`
	// Why the run failed, when `status` is `FAILED`. `NO_EMAIL_CHANNEL` and
	// `TEMPLATE_NOT_SUPPORTED` mean there was nothing to render;
	// `ALL_DEVICES_UNSUPPORTED` means every requested device has been retired and the
	// request can be fixed by choosing others.
	//
	// Any of "TEMPLATE_NOT_SUPPORTED", "NO_EMAIL_CHANNEL", "RENDER_FAILED",
	// "ALL_DEVICES_UNSUPPORTED", "VENDOR_ERROR".
	FailureReason PreviewRunFailureReason `json:"failure_reason"`
	// The version of the template that was rendered — `draft`, or a zero-padded
	// published version such as `v002`. Absent until the render settles.
	TemplateVersion string `json:"template_version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		DeviceIDs       respjson.Field
		Results         respjson.Field
		Status          respjson.Field
		TemplateID      respjson.Field
		FailureReason   respjson.Field
		TemplateVersion respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PreviewRunDetail) RawJSON() string { return r.JSON.raw }
func (r *PreviewRunDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Why the run failed, when `status` is `FAILED`. `NO_EMAIL_CHANNEL` and
// `TEMPLATE_NOT_SUPPORTED` mean there was nothing to render;
// `ALL_DEVICES_UNSUPPORTED` means every requested device has been retired and the
// request can be fixed by choosing others.
type PreviewRunFailureReason string

const (
	PreviewRunFailureReasonTemplateNotSupported  PreviewRunFailureReason = "TEMPLATE_NOT_SUPPORTED"
	PreviewRunFailureReasonNoEmailChannel        PreviewRunFailureReason = "NO_EMAIL_CHANNEL"
	PreviewRunFailureReasonRenderFailed          PreviewRunFailureReason = "RENDER_FAILED"
	PreviewRunFailureReasonAllDevicesUnsupported PreviewRunFailureReason = "ALL_DEVICES_UNSUPPORTED"
	PreviewRunFailureReasonVendorError           PreviewRunFailureReason = "VENDOR_ERROR"
)

// Paginated list of preview runs, newest first.
type PreviewRunListResponse struct {
	Paging  shared.Paging `json:"paging" api:"required"`
	Results []PreviewRun  `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Paging      respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PreviewRunListResponse) RawJSON() string { return r.JSON.raw }
func (r *PreviewRunListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Where the run itself has got to. `PENDING` and `RENDERED` mean Courier is still
// preparing the email, `SUBMITTED` means it is with the rendering service, and
// `COMPLETED` means every device has reported. `FAILED` is the run as a whole
// failing — an individual device failing never fails the run.
type PreviewRunStatus string

const (
	PreviewRunStatusPending   PreviewRunStatus = "PENDING"
	PreviewRunStatusRendered  PreviewRunStatus = "RENDERED"
	PreviewRunStatusSubmitted PreviewRunStatus = "SUBMITTED"
	PreviewRunStatusCompleted PreviewRunStatus = "COMPLETED"
	PreviewRunStatusFailed    PreviewRunStatus = "FAILED"
)

type NotificationPreviewRunNewParams struct {
	// Request body for creating a preview run of the template in the path. Provide
	// exactly one of `device_set_id` or `device_ids`. The template is the path's
	// `{id}`; a `template_id` here is an unknown key and a 400.
	CreatePreviewRunRequest CreatePreviewRunRequestParam
	IdempotencyKey          param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	XIdempotencyExpiration  param.Opt[string] `header:"x-idempotency-expiration,omitzero" json:"-"`
	paramObj
}

func (r NotificationPreviewRunNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreatePreviewRunRequest)
}
func (r *NotificationPreviewRunNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationPreviewRunGetParams struct {
	ID string `path:"id" api:"required" json:"-"`
	paramObj
}

type NotificationPreviewRunListParams struct {
	// Opaque pagination cursor from a previous response. Omit for the first page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of results per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NotificationPreviewRunListParams]'s query parameters as
// `url.Values`.
func (r NotificationPreviewRunListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
