// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/trycourier/courier-go/v4/internal/apijson"
	"github.com/trycourier/courier-go/v4/internal/apiquery"
	"github.com/trycourier/courier-go/v4/internal/requestconfig"
	"github.com/trycourier/courier-go/v4/option"
	"github.com/trycourier/courier-go/v4/packages/param"
	"github.com/trycourier/courier-go/v4/packages/respjson"
	"github.com/trycourier/courier-go/v4/shared"
)

// BulkService contains methods and other services that help with interacting with
// the Courier API.
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

// Ingest user data into a Bulk Job.
//
// **Important**: For email-based bulk jobs, each user must include `profile.email`
// for provider routing to work correctly. The `to.email` field is not sufficient
// for email provider routing.
func (r *BulkService) AddUsers(ctx context.Context, jobID string, body BulkAddUsersParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return
	}
	path := fmt.Sprintf("bulk/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Creates a new bulk job for sending messages to multiple recipients.
//
// **Required**: `message.event` (event ID or notification ID)
//
// **Optional (V2 format)**: `message.template` (notification ID) or
// `message.content` (Elemental content) can be provided to override the
// notification associated with the event.
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
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return
	}
	path := fmt.Sprintf("bulk/%s/run", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Bulk message definition. Supports two formats:
//
//   - V1 format: Requires `event` field (event ID or notification ID)
//   - V2 format: Optionally use `template` (notification ID) or `content` (Elemental
//     content) in addition to `event`
type InboundBulkMessage struct {
	// Event ID or Notification ID (required). Can be either a Notification ID (e.g.,
	// "FRH3QXM9E34W4RKP7MRC8NZ1T8V8") or a custom Event ID (e.g., "welcome-email")
	// mapped to a notification.
	Event string `json:"event,required"`
	Brand string `json:"brand,nullable"`
	// Elemental content (optional, for V2 format). When provided, this will be used
	// instead of the notification associated with the `event` field.
	Content  InboundBulkMessageContentUnion `json:"content,nullable"`
	Data     map[string]any                 `json:"data,nullable"`
	Locale   map[string]map[string]any      `json:"locale,nullable"`
	Override map[string]any                 `json:"override,nullable"`
	// Notification ID or template ID (optional, for V2 format). When provided, this
	// will be used instead of the notification associated with the `event` field.
	Template string `json:"template,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Event       respjson.Field
		Brand       respjson.Field
		Content     respjson.Field
		Data        respjson.Field
		Locale      respjson.Field
		Override    respjson.Field
		Template    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboundBulkMessage) RawJSON() string { return r.JSON.raw }
func (r *InboundBulkMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this InboundBulkMessage to a InboundBulkMessageParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// InboundBulkMessageParam.Overrides()
func (r InboundBulkMessage) ToParam() InboundBulkMessageParam {
	return param.Override[InboundBulkMessageParam](json.RawMessage(r.RawJSON()))
}

// InboundBulkMessageContentUnion contains all possible properties and values from
// [shared.ElementalContentSugar], [shared.ElementalContent].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type InboundBulkMessageContentUnion struct {
	// This field is from variant [shared.ElementalContentSugar].
	Body string `json:"body"`
	// This field is from variant [shared.ElementalContentSugar].
	Title string `json:"title"`
	// This field is from variant [shared.ElementalContent].
	Elements []shared.ElementalNodeUnion `json:"elements"`
	// This field is from variant [shared.ElementalContent].
	Version string `json:"version"`
	JSON    struct {
		Body     respjson.Field
		Title    respjson.Field
		Elements respjson.Field
		Version  respjson.Field
		raw      string
	} `json:"-"`
}

func (u InboundBulkMessageContentUnion) AsElementalContentSugar() (v shared.ElementalContentSugar) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InboundBulkMessageContentUnion) AsElementalContent() (v shared.ElementalContent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InboundBulkMessageContentUnion) RawJSON() string { return u.JSON.raw }

func (r *InboundBulkMessageContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bulk message definition. Supports two formats:
//
//   - V1 format: Requires `event` field (event ID or notification ID)
//   - V2 format: Optionally use `template` (notification ID) or `content` (Elemental
//     content) in addition to `event`
//
// The property Event is required.
type InboundBulkMessageParam struct {
	// Event ID or Notification ID (required). Can be either a Notification ID (e.g.,
	// "FRH3QXM9E34W4RKP7MRC8NZ1T8V8") or a custom Event ID (e.g., "welcome-email")
	// mapped to a notification.
	Event string            `json:"event,required"`
	Brand param.Opt[string] `json:"brand,omitzero"`
	// Notification ID or template ID (optional, for V2 format). When provided, this
	// will be used instead of the notification associated with the `event` field.
	Template param.Opt[string] `json:"template,omitzero"`
	// Elemental content (optional, for V2 format). When provided, this will be used
	// instead of the notification associated with the `event` field.
	Content  InboundBulkMessageContentUnionParam `json:"content,omitzero"`
	Data     map[string]any                      `json:"data,omitzero"`
	Locale   map[string]map[string]any           `json:"locale,omitzero"`
	Override map[string]any                      `json:"override,omitzero"`
	paramObj
}

func (r InboundBulkMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow InboundBulkMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InboundBulkMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InboundBulkMessageContentUnionParam struct {
	OfElementalContentSugar *shared.ElementalContentSugarParam `json:",omitzero,inline"`
	OfElementalContent      *shared.ElementalContentParam      `json:",omitzero,inline"`
	paramUnion
}

func (u InboundBulkMessageContentUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfElementalContentSugar, u.OfElementalContent)
}
func (u *InboundBulkMessageContentUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *InboundBulkMessageContentUnionParam) asAny() any {
	if !param.IsOmitted(u.OfElementalContentSugar) {
		return u.OfElementalContentSugar
	} else if !param.IsOmitted(u.OfElementalContent) {
		return u.OfElementalContent
	}
	return nil
}

type InboundBulkMessageUser struct {
	// User-specific data that will be merged with message.data
	Data        any                         `json:"data"`
	Preferences shared.RecipientPreferences `json:"preferences,nullable"`
	// User profile information. For email-based bulk jobs, `profile.email` is required
	// for provider routing to determine if the message can be delivered. The email
	// address should be provided here rather than in `to.email`.
	Profile map[string]any `json:"profile,nullable"`
	// User ID (legacy field, use profile or to.user_id instead)
	Recipient string `json:"recipient,nullable"`
	// Optional recipient information. Note: For email provider routing, use
	// `profile.email` instead of `to.email`. The `to` field is primarily used for
	// recipient identification and data merging.
	To shared.UserRecipient `json:"to,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Preferences respjson.Field
		Profile     respjson.Field
		Recipient   respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboundBulkMessageUser) RawJSON() string { return r.JSON.raw }
func (r *InboundBulkMessageUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this InboundBulkMessageUser to a InboundBulkMessageUserParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// InboundBulkMessageUserParam.Overrides()
func (r InboundBulkMessageUser) ToParam() InboundBulkMessageUserParam {
	return param.Override[InboundBulkMessageUserParam](json.RawMessage(r.RawJSON()))
}

type InboundBulkMessageUserParam struct {
	// User ID (legacy field, use profile or to.user_id instead)
	Recipient param.Opt[string] `json:"recipient,omitzero"`
	// User profile information. For email-based bulk jobs, `profile.email` is required
	// for provider routing to determine if the message can be delivered. The email
	// address should be provided here rather than in `to.email`.
	Profile map[string]any `json:"profile,omitzero"`
	// User-specific data that will be merged with message.data
	Data        any                              `json:"data,omitzero"`
	Preferences shared.RecipientPreferencesParam `json:"preferences,omitzero"`
	// Optional recipient information. Note: For email provider routing, use
	// `profile.email` instead of `to.email`. The `to` field is primarily used for
	// recipient identification and data merging.
	To shared.UserRecipientParam `json:"to,omitzero"`
	paramObj
}

func (r InboundBulkMessageUserParam) MarshalJSON() (data []byte, err error) {
	type shadow InboundBulkMessageUserParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InboundBulkMessageUserParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	InboundBulkMessageUser
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
	// Bulk message definition. Supports two formats:
	//
	//   - V1 format: Requires `event` field (event ID or notification ID)
	//   - V2 format: Optionally use `template` (notification ID) or `content` (Elemental
	//     content) in addition to `event`
	Definition InboundBulkMessage `json:"definition,required"`
	Enqueued   int64              `json:"enqueued,required"`
	Failures   int64              `json:"failures,required"`
	Received   int64              `json:"received,required"`
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
	Users []InboundBulkMessageUserParam `json:"users,omitzero,required"`
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
	// Bulk message definition. Supports two formats:
	//
	//   - V1 format: Requires `event` field (event ID or notification ID)
	//   - V2 format: Optionally use `template` (notification ID) or `content` (Elemental
	//     content) in addition to `event`
	Message InboundBulkMessageParam `json:"message,omitzero,required"`
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
