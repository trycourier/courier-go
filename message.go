// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/trycourier/courier-go/internal/apijson"
	"github.com/trycourier/courier-go/internal/apiquery"
	"github.com/trycourier/courier-go/internal/requestconfig"
	"github.com/trycourier/courier-go/option"
	"github.com/trycourier/courier-go/packages/param"
	"github.com/trycourier/courier-go/packages/respjson"
	"github.com/trycourier/courier-go/shared"
)

// MessageService contains methods and other services that help with interacting
// with the courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessageService] method instead.
type MessageService struct {
	Options []option.RequestOption
}

// NewMessageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMessageService(opts ...option.RequestOption) (r MessageService) {
	r = MessageService{}
	r.Options = opts
	return
}

// Fetch the status of a message you've previously sent.
func (r *MessageService) Get(ctx context.Context, messageID string, opts ...option.RequestOption) (res *MessageGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return
	}
	path := fmt.Sprintf("messages/%s", messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetch the statuses of messages you've previously sent.
func (r *MessageService) List(ctx context.Context, query MessageListParams, opts ...option.RequestOption) (res *MessageListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "messages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Cancel a message that is currently in the process of being delivered. A
// well-formatted API call to the cancel message API will return either `200`
// status code for a successful cancellation or `409` status code for an
// unsuccessful cancellation. Both cases will include the actual message record in
// the response body (see details below).
func (r *MessageService) Cancel(ctx context.Context, messageID string, opts ...option.RequestOption) (res *shared.MessageDetails, err error) {
	opts = slices.Concat(r.Options, opts)
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return
	}
	path := fmt.Sprintf("messages/%s/cancel", messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Get message content
func (r *MessageService) Content(ctx context.Context, messageID string, opts ...option.RequestOption) (res *MessageContentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return
	}
	path := fmt.Sprintf("messages/%s/output", messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetch the array of events of a message you've previously sent.
func (r *MessageService) History(ctx context.Context, messageID string, query MessageHistoryParams, opts ...option.RequestOption) (res *MessageHistoryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return
	}
	path := fmt.Sprintf("messages/%s/history", messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type MessageGetResponse struct {
	Providers []map[string]any `json:"providers,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Providers   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.MessageDetails
}

// Returns the unmodified JSON received from the API
func (r MessageGetResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageListResponse struct {
	// Paging information for the result set.
	Paging shared.Paging `json:"paging,required"`
	// An array of messages with their details.
	Results []shared.MessageDetails `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Paging      respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageListResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageContentResponse struct {
	// An array of render output of a previously sent message.
	Results []MessageContentResponseResult `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageContentResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageContentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageContentResponseResult struct {
	// The channel used for rendering the message.
	Channel string `json:"channel,required"`
	// The ID of channel used for rendering the message.
	ChannelID string `json:"channel_id,required"`
	// Content details of the rendered message.
	Content MessageContentResponseResultContent `json:"content,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Channel     respjson.Field
		ChannelID   respjson.Field
		Content     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageContentResponseResult) RawJSON() string { return r.JSON.raw }
func (r *MessageContentResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Content details of the rendered message.
type MessageContentResponseResultContent struct {
	// The blocks of the rendered message.
	Blocks []MessageContentResponseResultContentBlock `json:"blocks,required"`
	// The body of the rendered message.
	Body string `json:"body,required"`
	// The html content of the rendered message.
	HTML string `json:"html,required"`
	// The subject of the rendered message.
	Subject string `json:"subject,required"`
	// The text of the rendered message.
	Text string `json:"text,required"`
	// The title of the rendered message.
	Title string `json:"title,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Blocks      respjson.Field
		Body        respjson.Field
		HTML        respjson.Field
		Subject     respjson.Field
		Text        respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageContentResponseResultContent) RawJSON() string { return r.JSON.raw }
func (r *MessageContentResponseResultContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageContentResponseResultContentBlock struct {
	// The block text of the rendered message block.
	Text string `json:"text,required"`
	// The block type of the rendered message block.
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageContentResponseResultContentBlock) RawJSON() string { return r.JSON.raw }
func (r *MessageContentResponseResultContentBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageHistoryResponse struct {
	Results []map[string]any `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageHistoryResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageHistoryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageListParams struct {
	// A boolean value that indicates whether archived messages should be included in
	// the response.
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	// A unique identifier that allows for fetching the next set of messages.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// The enqueued datetime of a message to filter out messages received before.
	EnqueuedAfter param.Opt[string] `query:"enqueued_after,omitzero" json:"-"`
	// A unique identifier representing the event that was used to send the event.
	Event param.Opt[string] `query:"event,omitzero" json:"-"`
	// A unique identifier representing the list the message was sent to.
	List param.Opt[string] `query:"list,omitzero" json:"-"`
	// A unique identifier representing the message_id returned from either /send or
	// /send/list.
	MessageID param.Opt[string] `query:"messageId,omitzero" json:"-"`
	// A unique identifier representing the notification that was used to send the
	// event.
	Notification param.Opt[string] `query:"notification,omitzero" json:"-"`
	// A unique identifier representing the recipient associated with the requested
	// profile.
	Recipient param.Opt[string] `query:"recipient,omitzero" json:"-"`
	// A comma delimited list of 'tags'. Messages will be returned if they match any of
	// the tags passed in.
	Tags param.Opt[string] `query:"tags,omitzero" json:"-"`
	// Messages sent with the context of a Tenant
	TenantID param.Opt[string] `query:"tenant_id,omitzero" json:"-"`
	// The unique identifier used to trace the requests
	TraceID param.Opt[string] `query:"traceId,omitzero" json:"-"`
	// The key assocated to the provider you want to filter on. E.g., sendgrid, inbox,
	// twilio, slack, msteams, etc. Allows multiple values to be set in query
	// parameters.
	Provider []string `query:"provider,omitzero" json:"-"`
	// An indicator of the current status of the message. Allows multiple values to be
	// set in query parameters.
	Status []string `query:"status,omitzero" json:"-"`
	// A tag placed in the metadata.tags during a notification send. Allows multiple
	// values to be set in query parameters.
	Tag []string `query:"tag,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessageListParams]'s query parameters as `url.Values`.
func (r MessageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessageHistoryParams struct {
	// A supported Message History type that will filter the events returned.
	Type param.Opt[string] `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessageHistoryParams]'s query parameters as `url.Values`.
func (r MessageHistoryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
