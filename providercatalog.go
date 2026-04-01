// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
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

// ProviderCatalogService contains methods and other services that help with
// interacting with the Courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProviderCatalogService] method instead.
type ProviderCatalogService struct {
	Options []option.RequestOption
}

// NewProviderCatalogService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProviderCatalogService(opts ...option.RequestOption) (r ProviderCatalogService) {
	r = ProviderCatalogService{}
	r.Options = opts
	return
}

// Returns the catalog of available provider types with their display names,
// descriptions, and configuration schema fields (snake_case, with `type` and
// `required`). Providers with no configurable schema return only `provider`,
// `name`, and `description`.
func (r *ProviderCatalogService) List(ctx context.Context, query ProviderCatalogListParams, opts ...option.RequestOption) (res *ProviderCatalogListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "providers/catalog"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Paginated list of available provider types with their configuration schemas.
type ProviderCatalogListResponse struct {
	Paging  shared.Paging           `json:"paging" api:"required"`
	Results []ProvidersCatalogEntry `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Paging      respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProviderCatalogListResponse) RawJSON() string { return r.JSON.raw }
func (r *ProviderCatalogListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProviderCatalogListParams struct {
	// Exact match (case-insensitive) against the provider channel taxonomy (e.g.
	// `email`, `sms`, `push`).
	Channel param.Opt[string] `query:"channel,omitzero" json:"-"`
	// Comma-separated provider keys to filter by (e.g. `sendgrid,twilio`).
	Keys param.Opt[string] `query:"keys,omitzero" json:"-"`
	// Case-insensitive substring match against the provider display name.
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ProviderCatalogListParams]'s query parameters as
// `url.Values`.
func (r ProviderCatalogListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
