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

// TenantService contains methods and other services that help with interacting
// with the Courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTenantService] method instead.
type TenantService struct {
	Options     []option.RequestOption
	Preferences TenantPreferenceService
	Templates   TenantTemplateService
}

// NewTenantService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTenantService(opts ...option.RequestOption) (r TenantService) {
	r = TenantService{}
	r.Options = opts
	r.Preferences = NewTenantPreferenceService(opts...)
	r.Templates = NewTenantTemplateService(opts...)
	return
}

// Get a Tenant
func (r *TenantService) Get(ctx context.Context, tenantID string, opts ...option.RequestOption) (res *Tenant, err error) {
	opts = slices.Concat(r.Options, opts)
	if tenantID == "" {
		err = errors.New("missing required tenant_id parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s", tenantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Create or Replace a Tenant
func (r *TenantService) Update(ctx context.Context, tenantID string, body TenantUpdateParams, opts ...option.RequestOption) (res *Tenant, err error) {
	opts = slices.Concat(r.Options, opts)
	if tenantID == "" {
		err = errors.New("missing required tenant_id parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s", tenantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Get a List of Tenants
func (r *TenantService) List(ctx context.Context, query TenantListParams, opts ...option.RequestOption) (res *TenantListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tenants"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a Tenant
func (r *TenantService) Delete(ctx context.Context, tenantID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if tenantID == "" {
		err = errors.New("missing required tenant_id parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s", tenantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get Users in Tenant
func (r *TenantService) ListUsers(ctx context.Context, tenantID string, query TenantListUsersParams, opts ...option.RequestOption) (res *TenantListUsersResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if tenantID == "" {
		err = errors.New("missing required tenant_id parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s/users", tenantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type BaseTemplateTenantAssociation struct {
	// The template's id
	ID string `json:"id,required"`
	// The timestamp at which the template was created
	CreatedAt string `json:"created_at,required"`
	// The timestamp at which the template was published
	PublishedAt string `json:"published_at,required"`
	// The timestamp at which the template was last updated
	UpdatedAt string `json:"updated_at,required"`
	// The version of the template
	Version string `json:"version,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		PublishedAt respjson.Field
		UpdatedAt   respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaseTemplateTenantAssociation) RawJSON() string { return r.JSON.raw }
func (r *BaseTemplateTenantAssociation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DefaultPreferences struct {
	Items []DefaultPreferencesItem `json:"items,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DefaultPreferences) RawJSON() string { return r.JSON.raw }
func (r *DefaultPreferences) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DefaultPreferences to a DefaultPreferencesParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DefaultPreferencesParam.Overrides()
func (r DefaultPreferences) ToParam() DefaultPreferencesParam {
	return param.Override[DefaultPreferencesParam](json.RawMessage(r.RawJSON()))
}

type DefaultPreferencesItem struct {
	// Topic ID
	ID string `json:"id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	SubscriptionTopicNew
}

// Returns the unmodified JSON received from the API
func (r DefaultPreferencesItem) RawJSON() string { return r.JSON.raw }
func (r *DefaultPreferencesItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DefaultPreferencesParam struct {
	Items []DefaultPreferencesItemParam `json:"items,omitzero"`
	paramObj
}

func (r DefaultPreferencesParam) MarshalJSON() (data []byte, err error) {
	type shadow DefaultPreferencesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DefaultPreferencesParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DefaultPreferencesItemParam struct {
	// Topic ID
	ID string `json:"id,required"`
	SubscriptionTopicNewParam
}

func (r DefaultPreferencesItemParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*DefaultPreferencesItemParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// Request body for publishing a tenant template version
type PostTenantTemplatePublishRequestParam struct {
	// The version of the template to publish (e.g., "v1", "v2", "latest"). If not
	// provided, defaults to "latest".
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

func (r PostTenantTemplatePublishRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow PostTenantTemplatePublishRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PostTenantTemplatePublishRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response from publishing a tenant template
type PostTenantTemplatePublishResponse struct {
	// The template ID
	ID string `json:"id,required"`
	// The timestamp when the template was published
	PublishedAt string `json:"published_at,required"`
	// The published version of the template
	Version string `json:"version,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		PublishedAt respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostTenantTemplatePublishResponse) RawJSON() string { return r.JSON.raw }
func (r *PostTenantTemplatePublishResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body for creating or updating a tenant notification template
//
// The property Template is required.
type PutTenantTemplateRequestParam struct {
	// Template configuration for creating or updating a tenant notification template
	Template TenantTemplateInputParam `json:"template,omitzero,required"`
	// Whether to publish the template immediately after saving. When true, the
	// template becomes the active/published version. When false (default), the
	// template is saved as a draft.
	Published param.Opt[bool] `json:"published,omitzero"`
	paramObj
}

func (r PutTenantTemplateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow PutTenantTemplateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PutTenantTemplateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response from creating or updating a tenant notification template
type PutTenantTemplateResponse struct {
	// The template ID
	ID string `json:"id,required"`
	// The version of the saved template
	Version string `json:"version,required"`
	// The timestamp when the template was published. Only present if the template was
	// published as part of this request.
	PublishedAt string `json:"published_at,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Version     respjson.Field
		PublishedAt respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PutTenantTemplateResponse) RawJSON() string { return r.JSON.raw }
func (r *PutTenantTemplateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionTopicNew struct {
	// Any of "OPTED_OUT", "OPTED_IN", "REQUIRED".
	Status SubscriptionTopicNewStatus `json:"status,required"`
	// The default channels to send to this tenant when has_custom_routing is enabled
	CustomRouting []shared.ChannelClassification `json:"custom_routing,nullable"`
	// Override channel routing with custom preferences. This will override any
	// template preferences that are set, but a user can still customize their
	// preferences
	HasCustomRouting bool `json:"has_custom_routing,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status           respjson.Field
		CustomRouting    respjson.Field
		HasCustomRouting respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionTopicNew) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionTopicNew) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SubscriptionTopicNew to a SubscriptionTopicNewParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SubscriptionTopicNewParam.Overrides()
func (r SubscriptionTopicNew) ToParam() SubscriptionTopicNewParam {
	return param.Override[SubscriptionTopicNewParam](json.RawMessage(r.RawJSON()))
}

type SubscriptionTopicNewStatus string

const (
	SubscriptionTopicNewStatusOptedOut SubscriptionTopicNewStatus = "OPTED_OUT"
	SubscriptionTopicNewStatusOptedIn  SubscriptionTopicNewStatus = "OPTED_IN"
	SubscriptionTopicNewStatusRequired SubscriptionTopicNewStatus = "REQUIRED"
)

// The property Status is required.
type SubscriptionTopicNewParam struct {
	// Any of "OPTED_OUT", "OPTED_IN", "REQUIRED".
	Status SubscriptionTopicNewStatus `json:"status,omitzero,required"`
	// Override channel routing with custom preferences. This will override any
	// template preferences that are set, but a user can still customize their
	// preferences
	HasCustomRouting param.Opt[bool] `json:"has_custom_routing,omitzero"`
	// The default channels to send to this tenant when has_custom_routing is enabled
	CustomRouting []shared.ChannelClassification `json:"custom_routing,omitzero"`
	paramObj
}

func (r SubscriptionTopicNewParam) MarshalJSON() (data []byte, err error) {
	type shadow SubscriptionTopicNewParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubscriptionTopicNewParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Tenant struct {
	// Id of the tenant.
	ID string `json:"id,required"`
	// Name of the tenant.
	Name string `json:"name,required"`
	// Brand to be used for the account when one is not specified by the send call.
	BrandID string `json:"brand_id,nullable"`
	// Defines the preferences used for the account when the user hasn't specified
	// their own.
	DefaultPreferences DefaultPreferences `json:"default_preferences,nullable"`
	// Tenant's parent id (if any).
	ParentTenantID string `json:"parent_tenant_id,nullable"`
	// Arbitrary properties accessible to a template.
	Properties map[string]any `json:"properties,nullable"`
	// A user profile object merged with user profile on send.
	UserProfile map[string]any `json:"user_profile,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Name               respjson.Field
		BrandID            respjson.Field
		DefaultPreferences respjson.Field
		ParentTenantID     respjson.Field
		Properties         respjson.Field
		UserProfile        respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Tenant) RawJSON() string { return r.JSON.raw }
func (r *Tenant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantAssociation struct {
	// Tenant ID for the association between tenant and user
	TenantID string `json:"tenant_id,required"`
	// Additional metadata to be applied to a user profile when used in a tenant
	// context
	Profile map[string]any `json:"profile,nullable"`
	// Any of "user".
	Type TenantAssociationType `json:"type,nullable"`
	// User ID for the association between tenant and user
	UserID string `json:"user_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TenantID    respjson.Field
		Profile     respjson.Field
		Type        respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenantAssociation) RawJSON() string { return r.JSON.raw }
func (r *TenantAssociation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TenantAssociation to a TenantAssociationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TenantAssociationParam.Overrides()
func (r TenantAssociation) ToParam() TenantAssociationParam {
	return param.Override[TenantAssociationParam](json.RawMessage(r.RawJSON()))
}

type TenantAssociationType string

const (
	TenantAssociationTypeUser TenantAssociationType = "user"
)

// The property TenantID is required.
type TenantAssociationParam struct {
	// Tenant ID for the association between tenant and user
	TenantID string `json:"tenant_id,required"`
	// User ID for the association between tenant and user
	UserID param.Opt[string] `json:"user_id,omitzero"`
	// Additional metadata to be applied to a user profile when used in a tenant
	// context
	Profile map[string]any `json:"profile,omitzero"`
	// Any of "user".
	Type TenantAssociationType `json:"type,omitzero"`
	paramObj
}

func (r TenantAssociationParam) MarshalJSON() (data []byte, err error) {
	type shadow TenantAssociationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TenantAssociationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Template configuration for creating or updating a tenant notification template
//
// The property Content is required.
type TenantTemplateInputParam struct {
	// Template content configuration including blocks, elements, and message structure
	Content shared.ElementalContentParam `json:"content,omitzero,required"`
	// Channel-specific delivery configuration (email, SMS, push, etc.)
	Channels map[string]TenantTemplateInputChannelParam `json:"channels,omitzero"`
	// Provider-specific delivery configuration for routing to specific email/SMS
	// providers
	Providers map[string]TenantTemplateInputProviderParam `json:"providers,omitzero"`
	// Message routing configuration for multi-channel delivery strategies
	Routing shared.MessageRoutingParam `json:"routing,omitzero"`
	paramObj
}

func (r TenantTemplateInputParam) MarshalJSON() (data []byte, err error) {
	type shadow TenantTemplateInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TenantTemplateInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantTemplateInputChannelParam struct {
	// Brand id used for rendering.
	BrandID param.Opt[string] `json:"brand_id,omitzero"`
	// JS conditional with access to data/profile.
	If       param.Opt[string]                       `json:"if,omitzero"`
	Metadata TenantTemplateInputChannelMetadataParam `json:"metadata,omitzero"`
	// Channel specific overrides.
	Override map[string]any `json:"override,omitzero"`
	// Providers enabled for this channel.
	Providers []string `json:"providers,omitzero"`
	// Defaults to `single`.
	//
	// Any of "all", "single".
	RoutingMethod string                                  `json:"routing_method,omitzero"`
	Timeouts      TenantTemplateInputChannelTimeoutsParam `json:"timeouts,omitzero"`
	paramObj
}

func (r TenantTemplateInputChannelParam) MarshalJSON() (data []byte, err error) {
	type shadow TenantTemplateInputChannelParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TenantTemplateInputChannelParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[TenantTemplateInputChannelParam](
		"routing_method", "all", "single",
	)
}

type TenantTemplateInputChannelMetadataParam struct {
	Utm shared.UtmParam `json:"utm,omitzero"`
	paramObj
}

func (r TenantTemplateInputChannelMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow TenantTemplateInputChannelMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TenantTemplateInputChannelMetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantTemplateInputChannelTimeoutsParam struct {
	Channel  param.Opt[int64] `json:"channel,omitzero"`
	Provider param.Opt[int64] `json:"provider,omitzero"`
	paramObj
}

func (r TenantTemplateInputChannelTimeoutsParam) MarshalJSON() (data []byte, err error) {
	type shadow TenantTemplateInputChannelTimeoutsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TenantTemplateInputChannelTimeoutsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantTemplateInputProviderParam struct {
	// JS conditional with access to data/profile.
	If       param.Opt[string]                        `json:"if,omitzero"`
	Timeouts param.Opt[int64]                         `json:"timeouts,omitzero"`
	Metadata TenantTemplateInputProviderMetadataParam `json:"metadata,omitzero"`
	// Provider-specific overrides.
	Override map[string]any `json:"override,omitzero"`
	paramObj
}

func (r TenantTemplateInputProviderParam) MarshalJSON() (data []byte, err error) {
	type shadow TenantTemplateInputProviderParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TenantTemplateInputProviderParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantTemplateInputProviderMetadataParam struct {
	Utm shared.UtmParam `json:"utm,omitzero"`
	paramObj
}

func (r TenantTemplateInputProviderMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow TenantTemplateInputProviderMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TenantTemplateInputProviderMetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantListResponse struct {
	// Set to true when there are more pages that can be retrieved.
	HasMore bool `json:"has_more,required"`
	// An array of Tenants
	Items []Tenant `json:"items,required"`
	// Always set to "list". Represents the type of this object.
	//
	// Any of "list".
	Type TenantListResponseType `json:"type,required"`
	// A url that may be used to generate these results.
	URL string `json:"url,required"`
	// A pointer to the next page of results. Defined only when has_more is set to
	// true.
	Cursor string `json:"cursor,nullable"`
	// A url that may be used to generate fetch the next set of results. Defined only
	// when has_more is set to true
	NextURL string `json:"next_url,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore     respjson.Field
		Items       respjson.Field
		Type        respjson.Field
		URL         respjson.Field
		Cursor      respjson.Field
		NextURL     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenantListResponse) RawJSON() string { return r.JSON.raw }
func (r *TenantListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Always set to "list". Represents the type of this object.
type TenantListResponseType string

const (
	TenantListResponseTypeList TenantListResponseType = "list"
)

type TenantListUsersResponse struct {
	// Set to true when there are more pages that can be retrieved.
	HasMore bool `json:"has_more,required"`
	// Always set to `list`. Represents the type of this object.
	//
	// Any of "list".
	Type TenantListUsersResponseType `json:"type,required"`
	// A url that may be used to generate these results.
	URL string `json:"url,required"`
	// A pointer to the next page of results. Defined only when `has_more` is set to
	// true
	Cursor string              `json:"cursor,nullable"`
	Items  []TenantAssociation `json:"items,nullable"`
	// A url that may be used to generate fetch the next set of results. Defined only
	// when `has_more` is set to true
	NextURL string `json:"next_url,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore     respjson.Field
		Type        respjson.Field
		URL         respjson.Field
		Cursor      respjson.Field
		Items       respjson.Field
		NextURL     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenantListUsersResponse) RawJSON() string { return r.JSON.raw }
func (r *TenantListUsersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Always set to `list`. Represents the type of this object.
type TenantListUsersResponseType string

const (
	TenantListUsersResponseTypeList TenantListUsersResponseType = "list"
)

type TenantUpdateParams struct {
	// Name of the tenant.
	Name string `json:"name,required"`
	// Brand to be used for the account when one is not specified by the send call.
	BrandID param.Opt[string] `json:"brand_id,omitzero"`
	// Tenant's parent id (if any).
	ParentTenantID param.Opt[string] `json:"parent_tenant_id,omitzero"`
	// Arbitrary properties accessible to a template.
	Properties map[string]any `json:"properties,omitzero"`
	// A user profile object merged with user profile on send.
	UserProfile map[string]any `json:"user_profile,omitzero"`
	// Defines the preferences used for the tenant when the user hasn't specified their
	// own.
	DefaultPreferences DefaultPreferencesParam `json:"default_preferences,omitzero"`
	paramObj
}

func (r TenantUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TenantUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TenantUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantListParams struct {
	// Continue the pagination with the next cursor
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// The number of tenants to return (defaults to 20, maximum value of 100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter the list of tenants by parent_id
	ParentTenantID param.Opt[string] `query:"parent_tenant_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TenantListParams]'s query parameters as `url.Values`.
func (r TenantListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TenantListUsersParams struct {
	// Continue the pagination with the next cursor
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// The number of accounts to return (defaults to 20, maximum value of 100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TenantListUsersParams]'s query parameters as `url.Values`.
func (r TenantListUsersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
