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

// NotificationService contains methods and other services that help with
// interacting with the courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNotificationService] method instead.
type NotificationService struct {
	Options []option.RequestOption
	Draft   NotificationDraftService
	Checks  NotificationCheckService
}

// NewNotificationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNotificationService(opts ...option.RequestOption) (r NotificationService) {
	r = NotificationService{}
	r.Options = opts
	r.Draft = NewNotificationDraftService(opts...)
	r.Checks = NewNotificationCheckService(opts...)
	return
}

func (r *NotificationService) List(ctx context.Context, query NotificationListParams, opts ...option.RequestOption) (res *NotificationListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "notifications"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

func (r *NotificationService) GetContent(ctx context.Context, id string, opts ...option.RequestOption) (res *shared.NotificationGetContent, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("notifications/%s/content", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type NotificationListResponse struct {
	Paging  shared.Paging                    `json:"paging,required"`
	Results []NotificationListResponseResult `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Paging      respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationListResponse) RawJSON() string { return r.JSON.raw }
func (r *NotificationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationListResponseResult struct {
	ID        string                             `json:"id,required"`
	CreatedAt int64                              `json:"created_at,required"`
	Note      string                             `json:"note,required"`
	Routing   shared.MessageRouting              `json:"routing,required"`
	TopicID   string                             `json:"topic_id,required"`
	UpdatedAt int64                              `json:"updated_at,required"`
	Tags      NotificationListResponseResultTags `json:"tags,nullable"`
	Title     string                             `json:"title,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Note        respjson.Field
		Routing     respjson.Field
		TopicID     respjson.Field
		UpdatedAt   respjson.Field
		Tags        respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationListResponseResult) RawJSON() string { return r.JSON.raw }
func (r *NotificationListResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationListResponseResultTags struct {
	Data []NotificationListResponseResultTagsData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationListResponseResultTags) RawJSON() string { return r.JSON.raw }
func (r *NotificationListResponseResultTags) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationListResponseResultTagsData struct {
	ID   string `json:"id,required"`
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationListResponseResultTagsData) RawJSON() string { return r.JSON.raw }
func (r *NotificationListResponseResultTagsData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationListParams struct {
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Retrieve the notes from the Notification template settings.
	Notes param.Opt[bool] `query:"notes,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NotificationListParams]'s query parameters as `url.Values`.
func (r NotificationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
