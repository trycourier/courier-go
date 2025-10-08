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

// BrandService contains methods and other services that help with interacting with
// the courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBrandService] method instead.
type BrandService struct {
	Options []option.RequestOption
}

// NewBrandService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBrandService(opts ...option.RequestOption) (r BrandService) {
	r = BrandService{}
	r.Options = opts
	return
}

// Create a new brand
func (r *BrandService) New(ctx context.Context, body BrandNewParams, opts ...option.RequestOption) (res *shared.Brand, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "brands"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch a specific brand by brand ID.
func (r *BrandService) Get(ctx context.Context, brandID string, opts ...option.RequestOption) (res *shared.Brand, err error) {
	opts = slices.Concat(r.Options, opts)
	if brandID == "" {
		err = errors.New("missing required brand_id parameter")
		return
	}
	path := fmt.Sprintf("brands/%s", brandID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Replace an existing brand with the supplied values.
func (r *BrandService) Update(ctx context.Context, brandID string, body BrandUpdateParams, opts ...option.RequestOption) (res *shared.Brand, err error) {
	opts = slices.Concat(r.Options, opts)
	if brandID == "" {
		err = errors.New("missing required brand_id parameter")
		return
	}
	path := fmt.Sprintf("brands/%s", brandID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Get the list of brands.
func (r *BrandService) List(ctx context.Context, query BrandListParams, opts ...option.RequestOption) (res *BrandListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "brands"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a brand by brand ID.
func (r *BrandService) Delete(ctx context.Context, brandID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if brandID == "" {
		err = errors.New("missing required brand_id parameter")
		return
	}
	path := fmt.Sprintf("brands/%s", brandID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type BrandListResponse struct {
	Paging  shared.Paging  `json:"paging,required"`
	Results []shared.Brand `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Paging      respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandListResponse) RawJSON() string { return r.JSON.raw }
func (r *BrandListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandNewParams struct {
	Name     string                    `json:"name,required"`
	ID       param.Opt[string]         `json:"id,omitzero"`
	Settings shared.BrandSettingsParam `json:"settings,omitzero"`
	Snippets shared.BrandSnippetsParam `json:"snippets,omitzero"`
	paramObj
}

func (r BrandNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BrandNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandUpdateParams struct {
	// The name of the brand.
	Name     string                    `json:"name,required"`
	Settings shared.BrandSettingsParam `json:"settings,omitzero"`
	Snippets shared.BrandSnippetsParam `json:"snippets,omitzero"`
	paramObj
}

func (r BrandUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow BrandUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandListParams struct {
	// A unique identifier that allows for fetching the next set of brands.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BrandListParams]'s query parameters as `url.Values`.
func (r BrandListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
