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
func (r *BrandService) New(ctx context.Context, body BrandNewParams, opts ...option.RequestOption) (res *Brand, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "brands"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch a specific brand by brand ID.
func (r *BrandService) Get(ctx context.Context, brandID string, opts ...option.RequestOption) (res *Brand, err error) {
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
func (r *BrandService) Update(ctx context.Context, brandID string, body BrandUpdateParams, opts ...option.RequestOption) (res *Brand, err error) {
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

type Brand struct {
	// The date/time of when the brand was created. Represented in milliseconds since
	// Unix epoch.
	Created int64 `json:"created,required"`
	// Brand name
	Name string `json:"name,required"`
	// The date/time of when the brand was published. Represented in milliseconds since
	// Unix epoch.
	Published int64         `json:"published,required"`
	Settings  BrandSettings `json:"settings,required"`
	// The date/time of when the brand was updated. Represented in milliseconds since
	// Unix epoch.
	Updated int64 `json:"updated,required"`
	// The version identifier for the brand
	Version string `json:"version,required"`
	// Brand Identifier
	ID       string        `json:"id,nullable"`
	Snippets BrandSnippets `json:"snippets,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Created     respjson.Field
		Name        respjson.Field
		Published   respjson.Field
		Settings    respjson.Field
		Updated     respjson.Field
		Version     respjson.Field
		ID          respjson.Field
		Snippets    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Brand) RawJSON() string { return r.JSON.raw }
func (r *Brand) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandSettings struct {
	Colors BrandSettingsColors `json:"colors,nullable"`
	Email  BrandSettingsEmail  `json:"email,nullable"`
	Inapp  any                 `json:"inapp"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Colors      respjson.Field
		Email       respjson.Field
		Inapp       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandSettings) RawJSON() string { return r.JSON.raw }
func (r *BrandSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BrandSettings to a BrandSettingsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BrandSettingsParam.Overrides()
func (r BrandSettings) ToParam() BrandSettingsParam {
	return param.Override[BrandSettingsParam](json.RawMessage(r.RawJSON()))
}

type BrandSettingsColors struct {
	Primary   string `json:"primary,nullable"`
	Secondary string `json:"secondary,nullable"`
	Tertiary  string `json:"tertiary,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Primary     respjson.Field
		Secondary   respjson.Field
		Tertiary    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandSettingsColors) RawJSON() string { return r.JSON.raw }
func (r *BrandSettingsColors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandSettingsEmail struct {
	Footer any `json:"footer,required"`
	Header any `json:"header,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Footer      respjson.Field
		Header      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandSettingsEmail) RawJSON() string { return r.JSON.raw }
func (r *BrandSettingsEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandSettingsParam struct {
	Colors BrandSettingsColorsParam `json:"colors,omitzero"`
	Email  BrandSettingsEmailParam  `json:"email,omitzero"`
	Inapp  any                      `json:"inapp,omitzero"`
	paramObj
}

func (r BrandSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandSettingsColorsParam struct {
	Primary   param.Opt[string] `json:"primary,omitzero"`
	Secondary param.Opt[string] `json:"secondary,omitzero"`
	Tertiary  param.Opt[string] `json:"tertiary,omitzero"`
	paramObj
}

func (r BrandSettingsColorsParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandSettingsColorsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandSettingsColorsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Footer, Header are required.
type BrandSettingsEmailParam struct {
	Footer any `json:"footer,omitzero,required"`
	Header any `json:"header,omitzero,required"`
	paramObj
}

func (r BrandSettingsEmailParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandSettingsEmailParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandSettingsEmailParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandSnippets struct {
	Items []BrandSnippetsItem `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandSnippets) RawJSON() string { return r.JSON.raw }
func (r *BrandSnippets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BrandSnippets to a BrandSnippetsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BrandSnippetsParam.Overrides()
func (r BrandSnippets) ToParam() BrandSnippetsParam {
	return param.Override[BrandSnippetsParam](json.RawMessage(r.RawJSON()))
}

type BrandSnippetsItem struct {
	// Any of "handlebars".
	Format string `json:"format,required"`
	Name   string `json:"name,required"`
	Value  string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Format      respjson.Field
		Name        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandSnippetsItem) RawJSON() string { return r.JSON.raw }
func (r *BrandSnippetsItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Items is required.
type BrandSnippetsParam struct {
	Items []BrandSnippetsItemParam `json:"items,omitzero,required"`
	paramObj
}

func (r BrandSnippetsParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandSnippetsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandSnippetsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Format, Name, Value are required.
type BrandSnippetsItemParam struct {
	// Any of "handlebars".
	Format string `json:"format,omitzero,required"`
	Name   string `json:"name,required"`
	Value  string `json:"value,required"`
	paramObj
}

func (r BrandSnippetsItemParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandSnippetsItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandSnippetsItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BrandSnippetsItemParam](
		"format", "handlebars",
	)
}

type BrandListResponse struct {
	Paging  Paging  `json:"paging,required"`
	Results []Brand `json:"results,required"`
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
	// The name of the brand.
	Name     string             `json:"name,required"`
	Settings BrandSettingsParam `json:"settings,omitzero,required"`
	ID       param.Opt[string]  `json:"id,omitzero"`
	Snippets BrandSnippetsParam `json:"snippets,omitzero"`
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
	Name     string             `json:"name,required"`
	Settings BrandSettingsParam `json:"settings,omitzero"`
	Snippets BrandSnippetsParam `json:"snippets,omitzero"`
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
