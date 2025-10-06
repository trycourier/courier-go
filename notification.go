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

	"github.com/trycourier/courier-go/internal/apijson"
	"github.com/trycourier/courier-go/internal/apiquery"
	"github.com/trycourier/courier-go/internal/requestconfig"
	"github.com/trycourier/courier-go/option"
	"github.com/trycourier/courier-go/packages/param"
	"github.com/trycourier/courier-go/packages/respjson"
)

// NotificationService contains methods and other services that help with
// interacting with the courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNotificationService] method instead.
type NotificationService struct {
	Options []option.RequestOption
	Checks  NotificationCheckService
	Draft   NotificationDraftService
}

// NewNotificationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNotificationService(opts ...option.RequestOption) (r NotificationService) {
	r = NotificationService{}
	r.Options = opts
	r.Checks = NewNotificationCheckService(opts...)
	r.Draft = NewNotificationDraftService(opts...)
	return
}

func (r *NotificationService) List(ctx context.Context, query NotificationListParams, opts ...option.RequestOption) (res *NotificationListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "notifications"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

func (r *NotificationService) GetContent(ctx context.Context, id string, opts ...option.RequestOption) (res *NotificationContent, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("notifications/%s/content", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type NotificationContent struct {
	Blocks   []NotificationContentBlock   `json:"blocks,nullable"`
	Channels []NotificationContentChannel `json:"channels,nullable"`
	Checksum string                       `json:"checksum,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Blocks      respjson.Field
		Channels    respjson.Field
		Checksum    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationContent) RawJSON() string { return r.JSON.raw }
func (r *NotificationContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationContentBlock struct {
	ID string `json:"id,required"`
	// Any of "action", "divider", "image", "jsonnet", "list", "markdown", "quote",
	// "template", "text".
	Type     string                                         `json:"type,required"`
	Alias    string                                         `json:"alias,nullable"`
	Checksum string                                         `json:"checksum,nullable"`
	Content  NotificationContentBlockContentUnion           `json:"content,nullable"`
	Context  string                                         `json:"context,nullable"`
	Locales  map[string]NotificationContentBlockLocaleUnion `json:"locales,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		Alias       respjson.Field
		Checksum    respjson.Field
		Content     respjson.Field
		Context     respjson.Field
		Locales     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationContentBlock) RawJSON() string { return r.JSON.raw }
func (r *NotificationContentBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// NotificationContentBlockContentUnion contains all possible properties and values
// from [string], [NotificationContentBlockContentNotificationContentHierarchy].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type NotificationContentBlockContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is from variant
	// [NotificationContentBlockContentNotificationContentHierarchy].
	Children string `json:"children"`
	// This field is from variant
	// [NotificationContentBlockContentNotificationContentHierarchy].
	Parent string `json:"parent"`
	JSON   struct {
		OfString respjson.Field
		Children respjson.Field
		Parent   respjson.Field
		raw      string
	} `json:"-"`
}

func (u NotificationContentBlockContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NotificationContentBlockContentUnion) AsNotificationContentHierarchy() (v NotificationContentBlockContentNotificationContentHierarchy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u NotificationContentBlockContentUnion) RawJSON() string { return u.JSON.raw }

func (r *NotificationContentBlockContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationContentBlockContentNotificationContentHierarchy struct {
	Children string `json:"children,nullable"`
	Parent   string `json:"parent,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Children    respjson.Field
		Parent      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationContentBlockContentNotificationContentHierarchy) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationContentBlockContentNotificationContentHierarchy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// NotificationContentBlockLocaleUnion contains all possible properties and values
// from [string], [NotificationContentBlockLocaleNotificationContentHierarchy].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type NotificationContentBlockLocaleUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is from variant
	// [NotificationContentBlockLocaleNotificationContentHierarchy].
	Children string `json:"children"`
	// This field is from variant
	// [NotificationContentBlockLocaleNotificationContentHierarchy].
	Parent string `json:"parent"`
	JSON   struct {
		OfString respjson.Field
		Children respjson.Field
		Parent   respjson.Field
		raw      string
	} `json:"-"`
}

func (u NotificationContentBlockLocaleUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NotificationContentBlockLocaleUnion) AsNotificationContentHierarchy() (v NotificationContentBlockLocaleNotificationContentHierarchy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u NotificationContentBlockLocaleUnion) RawJSON() string { return u.JSON.raw }

func (r *NotificationContentBlockLocaleUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationContentBlockLocaleNotificationContentHierarchy struct {
	Children string `json:"children,nullable"`
	Parent   string `json:"parent,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Children    respjson.Field
		Parent      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationContentBlockLocaleNotificationContentHierarchy) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationContentBlockLocaleNotificationContentHierarchy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationContentChannel struct {
	ID       string                                      `json:"id,required"`
	Checksum string                                      `json:"checksum,nullable"`
	Content  NotificationContentChannelContent           `json:"content,nullable"`
	Locales  map[string]NotificationContentChannelLocale `json:"locales,nullable"`
	Type     string                                      `json:"type,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Checksum    respjson.Field
		Content     respjson.Field
		Locales     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationContentChannel) RawJSON() string { return r.JSON.raw }
func (r *NotificationContentChannel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationContentChannelContent struct {
	Subject string `json:"subject,nullable"`
	Title   string `json:"title,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Subject     respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationContentChannelContent) RawJSON() string { return r.JSON.raw }
func (r *NotificationContentChannelContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationContentChannelLocale struct {
	Subject string `json:"subject,nullable"`
	Title   string `json:"title,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Subject     respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationContentChannelLocale) RawJSON() string { return r.JSON.raw }
func (r *NotificationContentChannelLocale) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationListResponse struct {
	Paging  Paging                           `json:"paging,required"`
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
	Routing   MessageRouting                     `json:"routing,required"`
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
