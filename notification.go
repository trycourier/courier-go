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
	shimjson "github.com/trycourier/courier-go/v4/internal/encoding/json"
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

// Create a notification template. Requires all fields in the notification object.
// Templates are created in draft state by default.
func (r *NotificationService) New(ctx context.Context, body NotificationNewParams, opts ...option.RequestOption) (res *NotificationTemplateMutationResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "notifications"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a notification template by ID. Returns the published version by
// default. Pass version=draft to retrieve an unpublished template.
func (r *NotificationService) Get(ctx context.Context, id string, query NotificationGetParams, opts ...option.RequestOption) (res *NotificationTemplateGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("notifications/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List notification templates in your workspace.
func (r *NotificationService) List(ctx context.Context, query NotificationListParams, opts ...option.RequestOption) (res *NotificationListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "notifications"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Archive a notification template.
func (r *NotificationService) Archive(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("notifications/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// List versions of a notification template.
func (r *NotificationService) ListVersions(ctx context.Context, id string, query NotificationListVersionsParams, opts ...option.RequestOption) (res *NotificationTemplateVersionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("notifications/%s/versions", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Publish a notification template. Publishes the current draft by default. Pass a
// version in the request body to publish a specific historical version.
func (r *NotificationService) Publish(ctx context.Context, id string, body NotificationPublishParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("notifications/%s/publish", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Replace a notification template. All fields are required.
func (r *NotificationService) Replace(ctx context.Context, id string, body NotificationReplaceParams, opts ...option.RequestOption) (res *NotificationTemplateMutationResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("notifications/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

func (r *NotificationService) GetContent(ctx context.Context, id string, opts ...option.RequestOption) (res *NotificationGetContent, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("notifications/%s/content", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type BaseCheck struct {
	ID string `json:"id" api:"required"`
	// Any of "RESOLVED", "FAILED", "PENDING".
	Status BaseCheckStatus `json:"status" api:"required"`
	// Any of "custom".
	Type BaseCheckType `json:"type" api:"required"`
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
	ID string `json:"id" api:"required"`
	// Any of "RESOLVED", "FAILED", "PENDING".
	Status BaseCheckStatus `json:"status,omitzero" api:"required"`
	// Any of "custom".
	Type BaseCheckType `json:"type,omitzero" api:"required"`
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
	Updated int64 `json:"updated" api:"required"`
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
	Blocks   []NotificationGetContentBlock   `json:"blocks" api:"nullable"`
	Channels []NotificationGetContentChannel `json:"channels" api:"nullable"`
	Checksum string                          `json:"checksum" api:"nullable"`
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
	ID string `json:"id" api:"required"`
	// Any of "action", "divider", "image", "jsonnet", "list", "markdown", "quote",
	// "template", "text".
	Type     string                                            `json:"type" api:"required"`
	Alias    string                                            `json:"alias" api:"nullable"`
	Checksum string                                            `json:"checksum" api:"nullable"`
	Content  NotificationGetContentBlockContentUnion           `json:"content" api:"nullable"`
	Context  string                                            `json:"context" api:"nullable"`
	Locales  map[string]NotificationGetContentBlockLocaleUnion `json:"locales" api:"nullable"`
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
	Children string `json:"children" api:"nullable"`
	Parent   string `json:"parent" api:"nullable"`
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
	Children string `json:"children" api:"nullable"`
	Parent   string `json:"parent" api:"nullable"`
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
	ID       string                                         `json:"id" api:"required"`
	Checksum string                                         `json:"checksum" api:"nullable"`
	Content  NotificationGetContentChannelContent           `json:"content" api:"nullable"`
	Locales  map[string]NotificationGetContentChannelLocale `json:"locales" api:"nullable"`
	Type     string                                         `json:"type" api:"nullable"`
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
	Subject string `json:"subject" api:"nullable"`
	Title   string `json:"title" api:"nullable"`
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
	Subject string `json:"subject" api:"nullable"`
	Title   string `json:"title" api:"nullable"`
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

// Request body for creating a notification template.
//
// The property Notification is required.
type NotificationTemplateCreateRequestParam struct {
	// Full document shape used in POST and PUT request bodies, and returned inside the
	// GET response envelope.
	Notification NotificationTemplatePayloadParam `json:"notification,omitzero" api:"required"`
	// Template state after creation. Case-insensitive input, normalized to uppercase
	// in the response. Defaults to "DRAFT".
	//
	// Any of "DRAFT", "PUBLISHED".
	State NotificationTemplateCreateRequestState `json:"state,omitzero"`
	paramObj
}

func (r NotificationTemplateCreateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow NotificationTemplateCreateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NotificationTemplateCreateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Template state after creation. Case-insensitive input, normalized to uppercase
// in the response. Defaults to "DRAFT".
type NotificationTemplateCreateRequestState string

const (
	NotificationTemplateCreateRequestStateDraft     NotificationTemplateCreateRequestState = "DRAFT"
	NotificationTemplateCreateRequestStatePublished NotificationTemplateCreateRequestState = "PUBLISHED"
)

// Envelope response for GET /notifications/{id}. The notification object mirrors
// the POST/PUT input shape. Nullable fields return null when unset.
type NotificationTemplateGetResponse struct {
	// Epoch milliseconds when the template was created.
	Created int64 `json:"created" api:"required"`
	// User ID of the creator.
	Creator string `json:"creator" api:"required"`
	// Full document shape used in POST and PUT request bodies, and returned inside the
	// GET response envelope.
	Notification NotificationTemplateGetResponseNotification `json:"notification" api:"required"`
	// The template state. Always uppercase.
	//
	// Any of "DRAFT", "PUBLISHED".
	State NotificationTemplateGetResponseState `json:"state" api:"required"`
	// Epoch milliseconds of last update.
	Updated int64 `json:"updated"`
	// User ID of the last updater.
	Updater string `json:"updater"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Created      respjson.Field
		Creator      respjson.Field
		Notification respjson.Field
		State        respjson.Field
		Updated      respjson.Field
		Updater      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationTemplateGetResponse) RawJSON() string { return r.JSON.raw }
func (r *NotificationTemplateGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Full document shape used in POST and PUT request bodies, and returned inside the
// GET response envelope.
type NotificationTemplateGetResponseNotification struct {
	// The template ID.
	ID string `json:"id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	NotificationTemplatePayload
}

// Returns the unmodified JSON received from the API
func (r NotificationTemplateGetResponseNotification) RawJSON() string { return r.JSON.raw }
func (r *NotificationTemplateGetResponseNotification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The template state. Always uppercase.
type NotificationTemplateGetResponseState string

const (
	NotificationTemplateGetResponseStateDraft     NotificationTemplateGetResponseState = "DRAFT"
	NotificationTemplateGetResponseStatePublished NotificationTemplateGetResponseState = "PUBLISHED"
)

// Response returned by POST and PUT operations.
type NotificationTemplateMutationResponse struct {
	Notification NotificationTemplateMutationResponseNotification `json:"notification" api:"required"`
	// The template state after the operation. Always uppercase.
	//
	// Any of "DRAFT", "PUBLISHED".
	State NotificationTemplateMutationResponseState `json:"state" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Notification respjson.Field
		State        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationTemplateMutationResponse) RawJSON() string { return r.JSON.raw }
func (r *NotificationTemplateMutationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationTemplateMutationResponseNotification struct {
	// The ID of the created or updated template.
	ID string `json:"id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationTemplateMutationResponseNotification) RawJSON() string { return r.JSON.raw }
func (r *NotificationTemplateMutationResponseNotification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The template state after the operation. Always uppercase.
type NotificationTemplateMutationResponseState string

const (
	NotificationTemplateMutationResponseStateDraft     NotificationTemplateMutationResponseState = "DRAFT"
	NotificationTemplateMutationResponseStatePublished NotificationTemplateMutationResponseState = "PUBLISHED"
)

// Full document shape used in POST and PUT request bodies, and returned inside the
// GET response envelope.
type NotificationTemplatePayload struct {
	// Brand reference, or null for no brand.
	Brand NotificationTemplatePayloadBrand `json:"brand" api:"required"`
	// Elemental content definition.
	Content shared.ElementalContent `json:"content" api:"required"`
	// Display name for the template.
	Name string `json:"name" api:"required"`
	// Routing strategy reference, or null for none.
	Routing NotificationTemplatePayloadRouting `json:"routing" api:"required"`
	// Subscription topic reference, or null for none.
	Subscription NotificationTemplatePayloadSubscription `json:"subscription" api:"required"`
	// Tags for categorization. Send empty array for none.
	Tags []string `json:"tags" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Brand        respjson.Field
		Content      respjson.Field
		Name         respjson.Field
		Routing      respjson.Field
		Subscription respjson.Field
		Tags         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationTemplatePayload) RawJSON() string { return r.JSON.raw }
func (r *NotificationTemplatePayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this NotificationTemplatePayload to a
// NotificationTemplatePayloadParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// NotificationTemplatePayloadParam.Overrides()
func (r NotificationTemplatePayload) ToParam() NotificationTemplatePayloadParam {
	return param.Override[NotificationTemplatePayloadParam](json.RawMessage(r.RawJSON()))
}

// Brand reference, or null for no brand.
type NotificationTemplatePayloadBrand struct {
	ID string `json:"id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationTemplatePayloadBrand) RawJSON() string { return r.JSON.raw }
func (r *NotificationTemplatePayloadBrand) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Routing strategy reference, or null for none.
type NotificationTemplatePayloadRouting struct {
	StrategyID string `json:"strategy_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StrategyID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationTemplatePayloadRouting) RawJSON() string { return r.JSON.raw }
func (r *NotificationTemplatePayloadRouting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Subscription topic reference, or null for none.
type NotificationTemplatePayloadSubscription struct {
	TopicID string `json:"topic_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TopicID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationTemplatePayloadSubscription) RawJSON() string { return r.JSON.raw }
func (r *NotificationTemplatePayloadSubscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Full document shape used in POST and PUT request bodies, and returned inside the
// GET response envelope.
//
// The properties Brand, Content, Name, Routing, Subscription, Tags are required.
type NotificationTemplatePayloadParam struct {
	// Brand reference, or null for no brand.
	Brand NotificationTemplatePayloadBrandParam `json:"brand,omitzero" api:"required"`
	// Routing strategy reference, or null for none.
	Routing NotificationTemplatePayloadRoutingParam `json:"routing,omitzero" api:"required"`
	// Subscription topic reference, or null for none.
	Subscription NotificationTemplatePayloadSubscriptionParam `json:"subscription,omitzero" api:"required"`
	// Elemental content definition.
	Content shared.ElementalContentParam `json:"content,omitzero" api:"required"`
	// Display name for the template.
	Name string `json:"name" api:"required"`
	// Tags for categorization. Send empty array for none.
	Tags []string `json:"tags,omitzero" api:"required"`
	paramObj
}

func (r NotificationTemplatePayloadParam) MarshalJSON() (data []byte, err error) {
	type shadow NotificationTemplatePayloadParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NotificationTemplatePayloadParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Brand reference, or null for no brand.
//
// The property ID is required.
type NotificationTemplatePayloadBrandParam struct {
	ID string `json:"id" api:"required"`
	paramObj
}

func (r NotificationTemplatePayloadBrandParam) MarshalJSON() (data []byte, err error) {
	type shadow NotificationTemplatePayloadBrandParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NotificationTemplatePayloadBrandParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Routing strategy reference, or null for none.
//
// The property StrategyID is required.
type NotificationTemplatePayloadRoutingParam struct {
	StrategyID string `json:"strategy_id" api:"required"`
	paramObj
}

func (r NotificationTemplatePayloadRoutingParam) MarshalJSON() (data []byte, err error) {
	type shadow NotificationTemplatePayloadRoutingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NotificationTemplatePayloadRoutingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Subscription topic reference, or null for none.
//
// The property TopicID is required.
type NotificationTemplatePayloadSubscriptionParam struct {
	TopicID string `json:"topic_id" api:"required"`
	paramObj
}

func (r NotificationTemplatePayloadSubscriptionParam) MarshalJSON() (data []byte, err error) {
	type shadow NotificationTemplatePayloadSubscriptionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NotificationTemplatePayloadSubscriptionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional request body for publishing a notification template. Omit or send an
// empty object to publish the current draft.
type NotificationTemplatePublishRequestParam struct {
	// Historical version to publish (e.g. "v001"). Omit to publish the current draft.
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

func (r NotificationTemplatePublishRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow NotificationTemplatePublishRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NotificationTemplatePublishRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V2 (CDS) template summary returned in list responses.
type NotificationTemplateSummary struct {
	ID string `json:"id" api:"required"`
	// Epoch milliseconds when the template was created.
	Created int64 `json:"created" api:"required"`
	// User ID of the creator.
	Creator string `json:"creator" api:"required"`
	Name    string `json:"name" api:"required"`
	// Any of "DRAFT", "PUBLISHED".
	State NotificationTemplateSummaryState `json:"state" api:"required"`
	Tags  []string                         `json:"tags" api:"required"`
	// Epoch milliseconds of last update.
	Updated int64 `json:"updated"`
	// User ID of the last updater.
	Updater string `json:"updater"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Created     respjson.Field
		Creator     respjson.Field
		Name        respjson.Field
		State       respjson.Field
		Tags        respjson.Field
		Updated     respjson.Field
		Updater     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationTemplateSummary) RawJSON() string { return r.JSON.raw }
func (r *NotificationTemplateSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationTemplateSummaryState string

const (
	NotificationTemplateSummaryStateDraft     NotificationTemplateSummaryState = "DRAFT"
	NotificationTemplateSummaryStatePublished NotificationTemplateSummaryState = "PUBLISHED"
)

// Request body for replacing a notification template. Same shape as create. All
// fields required (PUT = full replacement).
//
// The property Notification is required.
type NotificationTemplateUpdateRequestParam struct {
	// Full document shape used in POST and PUT request bodies, and returned inside the
	// GET response envelope.
	Notification NotificationTemplatePayloadParam `json:"notification,omitzero" api:"required"`
	// Template state after update. Case-insensitive input, normalized to uppercase in
	// the response. Defaults to "DRAFT".
	//
	// Any of "DRAFT", "PUBLISHED".
	State NotificationTemplateUpdateRequestState `json:"state,omitzero"`
	paramObj
}

func (r NotificationTemplateUpdateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow NotificationTemplateUpdateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NotificationTemplateUpdateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Template state after update. Case-insensitive input, normalized to uppercase in
// the response. Defaults to "DRAFT".
type NotificationTemplateUpdateRequestState string

const (
	NotificationTemplateUpdateRequestStateDraft     NotificationTemplateUpdateRequestState = "DRAFT"
	NotificationTemplateUpdateRequestStatePublished NotificationTemplateUpdateRequestState = "PUBLISHED"
)

type NotificationTemplateVersionListResponse struct {
	Paging   shared.Paging `json:"paging" api:"required"`
	Versions []VersionNode `json:"versions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Paging      respjson.Field
		Versions    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationTemplateVersionListResponse) RawJSON() string { return r.JSON.raw }
func (r *NotificationTemplateVersionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A version entry for a notification template.
type VersionNode struct {
	// Epoch milliseconds when this version was created.
	Created int64 `json:"created" api:"required"`
	// User ID of the version creator.
	Creator string `json:"creator" api:"required"`
	// Version identifier. One of "draft", "published:vNNN" (current published
	// version), or "vNNN" (historical version).
	Version string `json:"version" api:"required"`
	// Whether the draft has unpublished changes. Only present on the draft version.
	HasChanges bool `json:"has_changes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Created     respjson.Field
		Creator     respjson.Field
		Version     respjson.Field
		HasChanges  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionNode) RawJSON() string { return r.JSON.raw }
func (r *VersionNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationListResponse struct {
	Paging shared.Paging `json:"paging" api:"required"`
	// Notification templates in this workspace.
	Results []NotificationListResponseResultUnion `json:"results" api:"required"`
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

// NotificationListResponseResultUnion contains all possible properties and values
// from [NotificationListResponseResultNotification],
// [NotificationTemplateSummary].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type NotificationListResponseResultUnion struct {
	ID string `json:"id"`
	// This field is from variant [NotificationListResponseResultNotification].
	CreatedAt int64 `json:"created_at"`
	// This field is from variant [NotificationListResponseResultNotification].
	EventIDs []string `json:"event_ids"`
	// This field is from variant [NotificationListResponseResultNotification].
	Note string `json:"note"`
	// This field is from variant [NotificationListResponseResultNotification].
	Routing shared.MessageRouting `json:"routing"`
	// This field is from variant [NotificationListResponseResultNotification].
	TopicID string `json:"topic_id"`
	// This field is from variant [NotificationListResponseResultNotification].
	UpdatedAt int64 `json:"updated_at"`
	// This field is a union of [NotificationListResponseResultNotificationTags],
	// [[]string]
	Tags NotificationListResponseResultUnionTags `json:"tags"`
	// This field is from variant [NotificationListResponseResultNotification].
	Title string `json:"title"`
	// This field is from variant [NotificationTemplateSummary].
	Created int64 `json:"created"`
	// This field is from variant [NotificationTemplateSummary].
	Creator string `json:"creator"`
	// This field is from variant [NotificationTemplateSummary].
	Name string `json:"name"`
	// This field is from variant [NotificationTemplateSummary].
	State NotificationTemplateSummaryState `json:"state"`
	// This field is from variant [NotificationTemplateSummary].
	Updated int64 `json:"updated"`
	// This field is from variant [NotificationTemplateSummary].
	Updater string `json:"updater"`
	JSON    struct {
		ID        respjson.Field
		CreatedAt respjson.Field
		EventIDs  respjson.Field
		Note      respjson.Field
		Routing   respjson.Field
		TopicID   respjson.Field
		UpdatedAt respjson.Field
		Tags      respjson.Field
		Title     respjson.Field
		Created   respjson.Field
		Creator   respjson.Field
		Name      respjson.Field
		State     respjson.Field
		Updated   respjson.Field
		Updater   respjson.Field
		raw       string
	} `json:"-"`
}

func (u NotificationListResponseResultUnion) AsNotification() (v NotificationListResponseResultNotification) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NotificationListResponseResultUnion) AsNotificationTemplateSummary() (v NotificationTemplateSummary) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u NotificationListResponseResultUnion) RawJSON() string { return u.JSON.raw }

func (r *NotificationListResponseResultUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// NotificationListResponseResultUnionTags is an implicit subunion of
// [NotificationListResponseResultUnion]. NotificationListResponseResultUnionTags
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [NotificationListResponseResultUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfStringArray]
type NotificationListResponseResultUnionTags struct {
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field is from variant [NotificationListResponseResultNotificationTags].
	Data []NotificationListResponseResultNotificationTagsData `json:"data"`
	JSON struct {
		OfStringArray respjson.Field
		Data          respjson.Field
		raw           string
	} `json:"-"`
}

func (r *NotificationListResponseResultUnionTags) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationListResponseResultNotification struct {
	ID        string `json:"id" api:"required"`
	CreatedAt int64  `json:"created_at" api:"required"`
	// Array of event IDs associated with this notification
	EventIDs  []string                                       `json:"event_ids" api:"required"`
	Note      string                                         `json:"note" api:"required"`
	Routing   shared.MessageRouting                          `json:"routing" api:"required"`
	TopicID   string                                         `json:"topic_id" api:"required"`
	UpdatedAt int64                                          `json:"updated_at" api:"required"`
	Tags      NotificationListResponseResultNotificationTags `json:"tags" api:"nullable"`
	Title     string                                         `json:"title" api:"nullable"`
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
func (r NotificationListResponseResultNotification) RawJSON() string { return r.JSON.raw }
func (r *NotificationListResponseResultNotification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationListResponseResultNotificationTags struct {
	Data []NotificationListResponseResultNotificationTagsData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationListResponseResultNotificationTags) RawJSON() string { return r.JSON.raw }
func (r *NotificationListResponseResultNotificationTags) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationListResponseResultNotificationTagsData struct {
	ID   string `json:"id" api:"required"`
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationListResponseResultNotificationTagsData) RawJSON() string { return r.JSON.raw }
func (r *NotificationListResponseResultNotificationTagsData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationNewParams struct {
	// Request body for creating a notification template.
	NotificationTemplateCreateRequest NotificationTemplateCreateRequestParam
	paramObj
}

func (r NotificationNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.NotificationTemplateCreateRequest)
}
func (r *NotificationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationGetParams struct {
	// Version to retrieve. One of "draft", "published", or a version string like
	// "v001". Defaults to "published".
	Version param.Opt[string] `query:"version,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NotificationGetParams]'s query parameters as `url.Values`.
func (r NotificationGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NotificationListParams struct {
	// Opaque pagination cursor from a previous response. Omit for the first page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Include template notes in the response. Only applies to legacy templates.
	Notes param.Opt[bool] `query:"notes,omitzero" json:"-"`
	// Filter to templates linked to this event map ID.
	EventID param.Opt[string] `query:"event_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NotificationListParams]'s query parameters as `url.Values`.
func (r NotificationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NotificationListVersionsParams struct {
	// Opaque pagination cursor from a previous response. Omit for the first page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of versions to return per page. Default 10, max 10.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NotificationListVersionsParams]'s query parameters as
// `url.Values`.
func (r NotificationListVersionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NotificationPublishParams struct {
	// Optional request body for publishing a notification template. Omit or send an
	// empty object to publish the current draft.
	NotificationTemplatePublishRequest NotificationTemplatePublishRequestParam
	paramObj
}

func (r NotificationPublishParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.NotificationTemplatePublishRequest)
}
func (r *NotificationPublishParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationReplaceParams struct {
	// Request body for replacing a notification template. Same shape as create. All
	// fields required (PUT = full replacement).
	NotificationTemplateUpdateRequest NotificationTemplateUpdateRequestParam
	paramObj
}

func (r NotificationReplaceParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.NotificationTemplateUpdateRequest)
}
func (r *NotificationReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
