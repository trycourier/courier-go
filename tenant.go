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

	"github.com/stainless-sdks/courier-go/internal/apijson"
	"github.com/stainless-sdks/courier-go/internal/apiquery"
	"github.com/stainless-sdks/courier-go/internal/requestconfig"
	"github.com/stainless-sdks/courier-go/option"
	"github.com/stainless-sdks/courier-go/packages/param"
	"github.com/stainless-sdks/courier-go/packages/respjson"
)

// TenantService contains methods and other services that help with interacting
// with the courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTenantService] method instead.
type TenantService struct {
	Options            []option.RequestOption
	DefaultPreferences TenantDefaultPreferenceService
}

// NewTenantService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTenantService(opts ...option.RequestOption) (r TenantService) {
	r = TenantService{}
	r.Options = opts
	r.DefaultPreferences = NewTenantDefaultPreferenceService(opts...)
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
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
	type shadow DefaultPreferencesItemParam
	return param.MarshalObject(r, (*shadow)(&r))
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
