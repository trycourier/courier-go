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

// Create a one-off send to a list or audience, author its content, then send it
// immediately or schedule it for later.
//
// BroadcastService contains methods and other services that help with interacting
// with the Courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBroadcastService] method instead.
type BroadcastService struct {
	Options []option.RequestOption
}

// NewBroadcastService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBroadcastService(opts ...option.RequestOption) (r BroadcastService) {
	r = BroadcastService{}
	r.Options = opts
	return
}

// Create a broadcast. Provisions a private notification template for the broadcast
// and returns the new broadcast in the draft state. Exactly one channel is
// required.
func (r *BroadcastService) New(ctx context.Context, body BroadcastNewParams, opts ...option.RequestOption) (res *Broadcast, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "broadcasts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a broadcast by ID. Archived broadcasts return 404.
func (r *BroadcastService) Get(ctx context.Context, broadcastID string, opts ...option.RequestOption) (res *Broadcast, err error) {
	opts = slices.Concat(r.Options, opts)
	if broadcastID == "" {
		err = errors.New("missing required broadcastId parameter")
		return nil, err
	}
	path := fmt.Sprintf("broadcasts/%s", broadcastID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a broadcast's name. Content is edited via the broadcast's notification
// template, not this endpoint.
func (r *BroadcastService) Update(ctx context.Context, broadcastID string, body BroadcastUpdateParams, opts ...option.RequestOption) (res *Broadcast, err error) {
	opts = slices.Concat(r.Options, opts)
	if broadcastID == "" {
		err = errors.New("missing required broadcastId parameter")
		return nil, err
	}
	path := fmt.Sprintf("broadcasts/%s", broadcastID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// List broadcasts in your workspace. Cursor-paginated; returns broadcasts
// newest-first.
func (r *BroadcastService) List(ctx context.Context, query BroadcastListParams, opts ...option.RequestOption) (res *BroadcastListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "broadcasts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Archive a broadcast. This is a soft delete — the archived broadcast is returned
// and no longer appears in list results.
func (r *BroadcastService) Archive(ctx context.Context, broadcastID string, opts ...option.RequestOption) (res *Broadcast, err error) {
	opts = slices.Concat(r.Options, opts)
	if broadcastID == "" {
		err = errors.New("missing required broadcastId parameter")
		return nil, err
	}
	path := fmt.Sprintf("broadcasts/%s", broadcastID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Cancel a broadcast's pending schedule, returning it to the draft state. Only
// valid for a scheduled broadcast.
func (r *BroadcastService) Cancel(ctx context.Context, broadcastID string, opts ...option.RequestOption) (res *Broadcast, err error) {
	opts = slices.Concat(r.Options, opts)
	if broadcastID == "" {
		err = errors.New("missing required broadcastId parameter")
		return nil, err
	}
	path := fmt.Sprintf("broadcasts/%s/cancel", broadcastID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Duplicate a broadcast (and its template) into a new draft named "{source name}
// (copy)".
func (r *BroadcastService) Duplicate(ctx context.Context, broadcastID string, opts ...option.RequestOption) (res *Broadcast, err error) {
	opts = slices.Concat(r.Options, opts)
	if broadcastID == "" {
		err = errors.New("missing required broadcastId parameter")
		return nil, err
	}
	path := fmt.Sprintf("broadcasts/%s/duplicate", broadcastID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Author the broadcast's content by replacing the draft elemental content of its
// private notification template. The draft is published automatically when the
// broadcast is sent or scheduled.
func (r *BroadcastService) PutContent(ctx context.Context, broadcastID string, body BroadcastPutContentParams, opts ...option.RequestOption) (res *NotificationContentMutationResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if broadcastID == "" {
		err = errors.New("missing required broadcastId parameter")
		return nil, err
	}
	path := fmt.Sprintf("broadcasts/%s/content", broadcastID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Retrieve the broadcast's content — the elemental content of its private
// notification template. Defaults to the working draft, since broadcast content is
// authored as a draft until the broadcast is sent.
func (r *BroadcastService) GetContent(ctx context.Context, broadcastID string, query BroadcastGetContentParams, opts ...option.RequestOption) (res *NotificationContentGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if broadcastID == "" {
		err = errors.New("missing required broadcastId parameter")
		return nil, err
	}
	path := fmt.Sprintf("broadcasts/%s/content", broadcastID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Schedule a broadcast for a future send to a list or audience. Publishes the
// broadcast template first. Not allowed once the broadcast is sending or sent. For
// an immediate send use POST /broadcasts/{broadcastId}/send.
func (r *BroadcastService) Schedule(ctx context.Context, broadcastID string, body BroadcastScheduleParams, opts ...option.RequestOption) (res *Broadcast, err error) {
	opts = slices.Concat(r.Options, opts)
	if broadcastID == "" {
		err = errors.New("missing required broadcastId parameter")
		return nil, err
	}
	path := fmt.Sprintf("broadcasts/%s/schedule", broadcastID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Send a broadcast immediately to a list or audience. Publishes the broadcast
// template first. Not allowed once the broadcast is sending or sent.
func (r *BroadcastService) Send(ctx context.Context, broadcastID string, body BroadcastSendParams, opts ...option.RequestOption) (res *Broadcast, err error) {
	opts = slices.Concat(r.Options, opts)
	if broadcastID == "" {
		err = errors.New("missing required broadcastId parameter")
		return nil, err
	}
	path := fmt.Sprintf("broadcasts/%s/send", broadcastID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// A broadcast — a single-channel message delivered to a known set of recipients (a
// list or audience).
type Broadcast struct {
	// The broadcast ID (bst\_ prefix).
	ID string `json:"id" api:"required"`
	// The broadcast's delivery channel.
	//
	// Any of "email", "sms", "push", "inbox", "slack", "msteams".
	Channel BroadcastChannel `json:"channel" api:"required"`
	// ISO 8601 timestamp when the broadcast was created.
	CreatedAt string `json:"created_at" api:"required"`
	// Actor that created the broadcast.
	CreatedBy string `json:"created_by" api:"required"`
	// Human-readable name.
	Name string `json:"name" api:"required"`
	// Lifecycle status of the broadcast.
	//
	// Any of "draft", "scheduled", "sending", "sent".
	Status BroadcastStatus `json:"status" api:"required"`
	// ISO 8601 timestamp of the last update.
	UpdatedAt string `json:"updated_at" api:"required"`
	// Actor that last updated the broadcast.
	UpdatedBy string `json:"updated_by" api:"required"`
	// ISO 8601 timestamp when the broadcast was archived, if archived.
	ArchivedAt string `json:"archived_at" api:"nullable"`
	// Actor that archived the broadcast, if archived.
	ArchivedBy string `json:"archived_by" api:"nullable"`
	// The delivery schedule and recipient targeting for a broadcast.
	Schedule BroadcastSchedule `json:"schedule" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Channel     respjson.Field
		CreatedAt   respjson.Field
		CreatedBy   respjson.Field
		Name        respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		UpdatedBy   respjson.Field
		ArchivedAt  respjson.Field
		ArchivedBy  respjson.Field
		Schedule    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Broadcast) RawJSON() string { return r.JSON.raw }
func (r *Broadcast) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The broadcast's delivery channel.
type BroadcastChannel string

const (
	BroadcastChannelEmail   BroadcastChannel = "email"
	BroadcastChannelSMS     BroadcastChannel = "sms"
	BroadcastChannelPush    BroadcastChannel = "push"
	BroadcastChannelInbox   BroadcastChannel = "inbox"
	BroadcastChannelSlack   BroadcastChannel = "slack"
	BroadcastChannelMsteams BroadcastChannel = "msteams"
)

// Lifecycle status of the broadcast.
type BroadcastStatus string

const (
	BroadcastStatusDraft     BroadcastStatus = "draft"
	BroadcastStatusScheduled BroadcastStatus = "scheduled"
	BroadcastStatusSending   BroadcastStatus = "sending"
	BroadcastStatusSent      BroadcastStatus = "sent"
)

// Paginated list of broadcasts.
type BroadcastListResponse struct {
	Paging  shared.Paging `json:"paging" api:"required"`
	Results []Broadcast   `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Paging      respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BroadcastListResponse) RawJSON() string { return r.JSON.raw }
func (r *BroadcastListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The delivery schedule and recipient targeting for a broadcast.
type BroadcastSchedule struct {
	// ID of the target list or audience.
	RecipientID string `json:"recipient_id" api:"required"`
	// Whether the broadcast targets a list or an audience.
	//
	// Any of "list", "audience".
	RecipientType BroadcastScheduleRecipientType `json:"recipient_type" api:"required"`
	// Wall-clock timestamp of the scheduled send, no timezone offset (e.g.
	// "2026-07-21T20:00:00").
	ScheduledTo string `json:"scheduled_to" api:"nullable"`
	// IANA timezone for the scheduled send (e.g. America/New_York).
	Timezone string `json:"timezone" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RecipientID   respjson.Field
		RecipientType respjson.Field
		ScheduledTo   respjson.Field
		Timezone      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BroadcastSchedule) RawJSON() string { return r.JSON.raw }
func (r *BroadcastSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the broadcast targets a list or an audience.
type BroadcastScheduleRecipientType string

const (
	BroadcastScheduleRecipientTypeList     BroadcastScheduleRecipientType = "list"
	BroadcastScheduleRecipientTypeAudience BroadcastScheduleRecipientType = "audience"
)

// Request body for creating a broadcast.
//
// The properties Channel, Name are required.
type CreateBroadcastRequestParam struct {
	// The single delivery channel for this broadcast.
	//
	// Any of "email", "sms", "push", "inbox", "slack", "msteams".
	Channel CreateBroadcastRequestChannel `json:"channel,omitzero" api:"required"`
	// Human-readable name.
	Name string `json:"name" api:"required"`
	paramObj
}

func (r CreateBroadcastRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateBroadcastRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateBroadcastRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The single delivery channel for this broadcast.
type CreateBroadcastRequestChannel string

const (
	CreateBroadcastRequestChannelEmail   CreateBroadcastRequestChannel = "email"
	CreateBroadcastRequestChannelSMS     CreateBroadcastRequestChannel = "sms"
	CreateBroadcastRequestChannelPush    CreateBroadcastRequestChannel = "push"
	CreateBroadcastRequestChannelInbox   CreateBroadcastRequestChannel = "inbox"
	CreateBroadcastRequestChannelSlack   CreateBroadcastRequestChannel = "slack"
	CreateBroadcastRequestChannelMsteams CreateBroadcastRequestChannel = "msteams"
)

// Request body for scheduling a broadcast for a future send.
//
// The properties RecipientID, RecipientType, ScheduledTo are required.
type ScheduleBroadcastRequestParam struct {
	// ID of the target list or audience.
	RecipientID string `json:"recipient_id" api:"required"`
	// Whether the broadcast targets a list or an audience.
	//
	// Any of "list", "audience".
	RecipientType ScheduleBroadcastRequestRecipientType `json:"recipient_type,omitzero" api:"required"`
	// Wall-clock timestamp of the future send, no timezone offset (e.g.
	// "2026-07-21T20:00:00"). The zone is given by `timezone`.
	ScheduledTo string `json:"scheduled_to" api:"required"`
	// IANA timezone for the scheduled send (e.g. America/New_York).
	Timezone param.Opt[string] `json:"timezone,omitzero"`
	paramObj
}

func (r ScheduleBroadcastRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ScheduleBroadcastRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScheduleBroadcastRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the broadcast targets a list or an audience.
type ScheduleBroadcastRequestRecipientType string

const (
	ScheduleBroadcastRequestRecipientTypeList     ScheduleBroadcastRequestRecipientType = "list"
	ScheduleBroadcastRequestRecipientTypeAudience ScheduleBroadcastRequestRecipientType = "audience"
)

// Request body for sending a broadcast immediately.
//
// The properties RecipientID, RecipientType are required.
type SendBroadcastRequestParam struct {
	// ID of the target list or audience.
	RecipientID string `json:"recipient_id" api:"required"`
	// Whether the broadcast targets a list or an audience.
	//
	// Any of "list", "audience".
	RecipientType SendBroadcastRequestRecipientType `json:"recipient_type,omitzero" api:"required"`
	paramObj
}

func (r SendBroadcastRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow SendBroadcastRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendBroadcastRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the broadcast targets a list or an audience.
type SendBroadcastRequestRecipientType string

const (
	SendBroadcastRequestRecipientTypeList     SendBroadcastRequestRecipientType = "list"
	SendBroadcastRequestRecipientTypeAudience SendBroadcastRequestRecipientType = "audience"
)

// Request body for updating a broadcast. Only the name is mutable.
//
// The property Name is required.
type UpdateBroadcastRequestParam struct {
	// New human-readable name.
	Name string `json:"name" api:"required"`
	paramObj
}

func (r UpdateBroadcastRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateBroadcastRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateBroadcastRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BroadcastNewParams struct {
	// Request body for creating a broadcast.
	CreateBroadcastRequest CreateBroadcastRequestParam
	paramObj
}

func (r BroadcastNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateBroadcastRequest)
}
func (r *BroadcastNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BroadcastUpdateParams struct {
	// Request body for updating a broadcast. Only the name is mutable.
	UpdateBroadcastRequest UpdateBroadcastRequestParam
	paramObj
}

func (r BroadcastUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateBroadcastRequest)
}
func (r *BroadcastUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BroadcastListParams struct {
	// Opaque pagination cursor from a previous response. Omit for the first page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of results per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BroadcastListParams]'s query parameters as `url.Values`.
func (r BroadcastListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BroadcastPutContentParams struct {
	// Request body for replacing the elemental content of a notification template.
	NotificationContentPutRequest NotificationContentPutRequestParam
	paramObj
}

func (r BroadcastPutContentParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.NotificationContentPutRequest)
}
func (r *BroadcastPutContentParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BroadcastGetContentParams struct {
	// Accepts `draft`, `published`, or a version string (e.g. `v001`). Defaults to
	// `draft`.
	Version param.Opt[string] `query:"version,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BroadcastGetContentParams]'s query parameters as
// `url.Values`.
func (r BroadcastGetContentParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BroadcastScheduleParams struct {
	// Request body for scheduling a broadcast for a future send.
	ScheduleBroadcastRequest ScheduleBroadcastRequestParam
	paramObj
}

func (r BroadcastScheduleParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ScheduleBroadcastRequest)
}
func (r *BroadcastScheduleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BroadcastSendParams struct {
	// Request body for sending a broadcast immediately.
	SendBroadcastRequest SendBroadcastRequestParam
	paramObj
}

func (r BroadcastSendParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SendBroadcastRequest)
}
func (r *BroadcastSendParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
