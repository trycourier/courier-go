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

// TenantTemplateService contains methods and other services that help with
// interacting with the Courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTenantTemplateService] method instead.
type TenantTemplateService struct {
	Options []option.RequestOption
}

// NewTenantTemplateService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTenantTemplateService(opts ...option.RequestOption) (r TenantTemplateService) {
	r = TenantTemplateService{}
	r.Options = opts
	return
}

// Get a Template in Tenant
func (r *TenantTemplateService) Get(ctx context.Context, templateID string, query TenantTemplateGetParams, opts ...option.RequestOption) (res *BaseTemplateTenantAssociation, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.TenantID == "" {
		err = errors.New("missing required tenant_id parameter")
		return
	}
	if templateID == "" {
		err = errors.New("missing required template_id parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s/templates/%s", query.TenantID, templateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Templates in Tenant
func (r *TenantTemplateService) List(ctx context.Context, tenantID string, query TenantTemplateListParams, opts ...option.RequestOption) (res *TenantTemplateListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if tenantID == "" {
		err = errors.New("missing required tenant_id parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s/templates", tenantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type TenantTemplateListResponse struct {
	// Set to true when there are more pages that can be retrieved.
	HasMore bool `json:"has_more,required"`
	// Always set to `list`. Represents the type of this object.
	//
	// Any of "list".
	Type TenantTemplateListResponseType `json:"type,required"`
	// A url that may be used to generate these results.
	URL string `json:"url,required"`
	// A pointer to the next page of results. Defined only when `has_more` is set to
	// true
	Cursor string                           `json:"cursor,nullable"`
	Items  []TenantTemplateListResponseItem `json:"items,nullable"`
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
func (r TenantTemplateListResponse) RawJSON() string { return r.JSON.raw }
func (r *TenantTemplateListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Always set to `list`. Represents the type of this object.
type TenantTemplateListResponseType string

const (
	TenantTemplateListResponseTypeList TenantTemplateListResponseType = "list"
)

type TenantTemplateListResponseItem struct {
	// The template's data containing it's routing configs
	Data TenantTemplateListResponseItemData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BaseTemplateTenantAssociation
}

// Returns the unmodified JSON received from the API
func (r TenantTemplateListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *TenantTemplateListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The template's data containing it's routing configs
type TenantTemplateListResponseItemData struct {
	Routing shared.MessageRouting `json:"routing,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Routing     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenantTemplateListResponseItemData) RawJSON() string { return r.JSON.raw }
func (r *TenantTemplateListResponseItemData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantTemplateGetParams struct {
	TenantID string `path:"tenant_id,required" json:"-"`
	paramObj
}

type TenantTemplateListParams struct {
	// Continue the pagination with the next cursor
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// The number of templates to return (defaults to 20, maximum value of 100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TenantTemplateListParams]'s query parameters as
// `url.Values`.
func (r TenantTemplateListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
