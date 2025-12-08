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

// NotificationService contains methods and other services that help with
// interacting with the Courier API.
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

func (r *NotificationService) GetContent(ctx context.Context, id string, opts ...option.RequestOption) (res *NotificationGetContent, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("notifications/%s/content", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type BaseCheck struct {
	ID string `json:"id,required"`
	// Any of "RESOLVED", "FAILED", "PENDING".
	Status BaseCheckStatus `json:"status,required"`
	// Any of "custom".
	Type BaseCheckType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaseCheck) RawJSON() string { return r.JSON.raw }
func (r *BaseCheck) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BaseCheck to a BaseCheckParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BaseCheckParam.Overrides()
func (r BaseCheck) ToParam() BaseCheckParam {
	return param.Override[BaseCheckParam](json.RawMessage(r.RawJSON()))
}

type BaseCheckStatus string

const (
	BaseCheckStatusResolved BaseCheckStatus = "RESOLVED"
	BaseCheckStatusFailed   BaseCheckStatus = "FAILED"
	BaseCheckStatusPending  BaseCheckStatus = "PENDING"
)

type BaseCheckType string

const (
	BaseCheckTypeCustom BaseCheckType = "custom"
)

// The properties ID, Status, Type are required.
type BaseCheckParam struct {
	ID string `json:"id,required"`
	// Any of "RESOLVED", "FAILED", "PENDING".
	Status BaseCheckStatus `json:"status,omitzero,required"`
	// Any of "custom".
	Type BaseCheckType `json:"type,omitzero,required"`
	paramObj
}

func (r BaseCheckParam) MarshalJSON() (data []byte, err error) {
	type shadow BaseCheckParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaseCheckParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Check struct {
	Updated int64 `json:"updated,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Updated     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BaseCheck
}

// Returns the unmodified JSON received from the API
func (r Check) RawJSON() string { return r.JSON.raw }
func (r *Check) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationGetContent struct {
	Blocks   []NotificationGetContentBlock   `json:"blocks,nullable"`
	Channels []NotificationGetContentChannel `json:"channels,nullable"`
	Checksum string                          `json:"checksum,nullable"`
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
func (r NotificationGetContent) RawJSON() string { return r.JSON.raw }
func (r *NotificationGetContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationGetContentBlock struct {
	ID string `json:"id,required"`
	// Any of "action", "divider", "image", "jsonnet", "list", "markdown", "quote",
	// "template", "text".
	Type     string                                            `json:"type,required"`
	Alias    string                                            `json:"alias,nullable"`
	Checksum string                                            `json:"checksum,nullable"`
	Content  NotificationGetContentBlockContentUnion           `json:"content,nullable"`
	Context  string                                            `json:"context,nullable"`
	Locales  map[string]NotificationGetContentBlockLocaleUnion `json:"locales,nullable"`
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
func (r NotificationGetContentBlock) RawJSON() string { return r.JSON.raw }
func (r *NotificationGetContentBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// NotificationGetContentBlockContentUnion contains all possible properties and
// values from [string],
// [NotificationGetContentBlockContentNotificationContentHierarchy].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type NotificationGetContentBlockContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is from variant
	// [NotificationGetContentBlockContentNotificationContentHierarchy].
	Children string `json:"children"`
	// This field is from variant
	// [NotificationGetContentBlockContentNotificationContentHierarchy].
	Parent string `json:"parent"`
	JSON   struct {
		OfString respjson.Field
		Children respjson.Field
		Parent   respjson.Field
		raw      string
	} `json:"-"`
}

func (u NotificationGetContentBlockContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NotificationGetContentBlockContentUnion) AsNotificationContentHierarchy() (v NotificationGetContentBlockContentNotificationContentHierarchy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u NotificationGetContentBlockContentUnion) RawJSON() string { return u.JSON.raw }

func (r *NotificationGetContentBlockContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationGetContentBlockContentNotificationContentHierarchy struct {
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
func (r NotificationGetContentBlockContentNotificationContentHierarchy) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationGetContentBlockContentNotificationContentHierarchy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// NotificationGetContentBlockLocaleUnion contains all possible properties and
// values from [string],
// [NotificationGetContentBlockLocaleNotificationContentHierarchy].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type NotificationGetContentBlockLocaleUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is from variant
	// [NotificationGetContentBlockLocaleNotificationContentHierarchy].
	Children string `json:"children"`
	// This field is from variant
	// [NotificationGetContentBlockLocaleNotificationContentHierarchy].
	Parent string `json:"parent"`
	JSON   struct {
		OfString respjson.Field
		Children respjson.Field
		Parent   respjson.Field
		raw      string
	} `json:"-"`
}

func (u NotificationGetContentBlockLocaleUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NotificationGetContentBlockLocaleUnion) AsNotificationContentHierarchy() (v NotificationGetContentBlockLocaleNotificationContentHierarchy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u NotificationGetContentBlockLocaleUnion) RawJSON() string { return u.JSON.raw }

func (r *NotificationGetContentBlockLocaleUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationGetContentBlockLocaleNotificationContentHierarchy struct {
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
func (r NotificationGetContentBlockLocaleNotificationContentHierarchy) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationGetContentBlockLocaleNotificationContentHierarchy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationGetContentChannel struct {
	ID       string                                         `json:"id,required"`
	Checksum string                                         `json:"checksum,nullable"`
	Content  NotificationGetContentChannelContent           `json:"content,nullable"`
	Locales  map[string]NotificationGetContentChannelLocale `json:"locales,nullable"`
	Type     string                                         `json:"type,nullable"`
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
func (r NotificationGetContentChannel) RawJSON() string { return r.JSON.raw }
func (r *NotificationGetContentChannel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationGetContentChannelContent struct {
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
func (r NotificationGetContentChannelContent) RawJSON() string { return r.JSON.raw }
func (r *NotificationGetContentChannelContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationGetContentChannelLocale struct {
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
func (r NotificationGetContentChannelLocale) RawJSON() string { return r.JSON.raw }
func (r *NotificationGetContentChannelLocale) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	ID        string `json:"id,required"`
	CreatedAt int64  `json:"created_at,required"`
	// Array of event IDs associated with this notification
	EventIDs  []string                           `json:"event_ids,required"`
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
		EventIDs    respjson.Field
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
