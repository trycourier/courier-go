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

// UserTenantService contains methods and other services that help with interacting
// with the courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserTenantService] method instead.
type UserTenantService struct {
	Options []option.RequestOption
}

// NewUserTenantService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserTenantService(opts ...option.RequestOption) (r UserTenantService) {
	r = UserTenantService{}
	r.Options = opts
	return
}

// Returns a paginated list of user tenant associations.
func (r *UserTenantService) List(ctx context.Context, userID string, query UserTenantListParams, opts ...option.RequestOption) (res *UserTenantListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("users/%s/tenants", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// This endpoint is used to add a user to multiple tenants in one call. A custom
// profile can also be supplied for each tenant. This profile will be merged with
// the user's main profile when sending to the user with that tenant.
func (r *UserTenantService) AddMultiple(ctx context.Context, userID string, body UserTenantAddMultipleParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("users/%s/tenants", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// This endpoint is used to add a single tenant.
//
// A custom profile can also be supplied with the tenant. This profile will be
// merged with the user's main profile when sending to the user with that tenant.
func (r *UserTenantService) AddSingle(ctx context.Context, tenantID string, params UserTenantAddSingleParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if params.UserID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	if tenantID == "" {
		err = errors.New("missing required tenant_id parameter")
		return
	}
	path := fmt.Sprintf("users/%s/tenants/%s", params.UserID, tenantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, nil, opts...)
	return
}

// Removes a user from any tenants they may have been associated with.
func (r *UserTenantService) RemoveAll(ctx context.Context, userID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("users/%s/tenants", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Removes a user from the supplied tenant.
func (r *UserTenantService) RemoveSingle(ctx context.Context, tenantID string, body UserTenantRemoveSingleParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.UserID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	if tenantID == "" {
		err = errors.New("missing required tenant_id parameter")
		return
	}
	path := fmt.Sprintf("users/%s/tenants/%s", body.UserID, tenantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
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

type UserTenantListResponse struct {
	// Set to true when there are more pages that can be retrieved.
	HasMore bool `json:"has_more,required"`
	// Always set to `list`. Represents the type of this object.
	//
	// Any of "list".
	Type UserTenantListResponseType `json:"type,required"`
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
func (r UserTenantListResponse) RawJSON() string { return r.JSON.raw }
func (r *UserTenantListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Always set to `list`. Represents the type of this object.
type UserTenantListResponseType string

const (
	UserTenantListResponseTypeList UserTenantListResponseType = "list"
)

type UserTenantListParams struct {
	// Continue the pagination with the next cursor
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// The number of accounts to return (defaults to 20, maximum value of 100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UserTenantListParams]'s query parameters as `url.Values`.
func (r UserTenantListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserTenantAddMultipleParams struct {
	Tenants []TenantAssociationParam `json:"tenants,omitzero,required"`
	paramObj
}

func (r UserTenantAddMultipleParams) MarshalJSON() (data []byte, err error) {
	type shadow UserTenantAddMultipleParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserTenantAddMultipleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserTenantAddSingleParams struct {
	UserID  string         `path:"user_id,required" json:"-"`
	Profile map[string]any `json:"profile,omitzero"`
	paramObj
}

func (r UserTenantAddSingleParams) MarshalJSON() (data []byte, err error) {
	type shadow UserTenantAddSingleParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserTenantAddSingleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserTenantRemoveSingleParams struct {
	UserID string `path:"user_id,required" json:"-"`
	paramObj
}
