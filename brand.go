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
	ID        string        `json:"id,required"`
	Created   int64         `json:"created,required"`
	Name      string        `json:"name,required"`
	Updated   int64         `json:"updated,required"`
	Published int64         `json:"published,nullable"`
	Settings  BrandSettings `json:"settings,nullable"`
	Snippets  BrandSnippets `json:"snippets,nullable"`
	Version   string        `json:"version,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Created     respjson.Field
		Name        respjson.Field
		Updated     respjson.Field
		Published   respjson.Field
		Settings    respjson.Field
		Snippets    respjson.Field
		Version     respjson.Field
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
	Inapp  BrandSettingsInapp  `json:"inapp,nullable"`
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
	Primary     string            `json:"primary,nullable"`
	Secondary   string            `json:"secondary,nullable"`
	ExtraFields map[string]string `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Primary     respjson.Field
		Secondary   respjson.Field
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
	Footer           BrandSettingsEmailFooter           `json:"footer,nullable"`
	Head             BrandSettingsEmailHead             `json:"head,nullable"`
	Header           BrandSettingsEmailHeader           `json:"header,nullable"`
	TemplateOverride BrandSettingsEmailTemplateOverride `json:"templateOverride,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Footer           respjson.Field
		Head             respjson.Field
		Header           respjson.Field
		TemplateOverride respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandSettingsEmail) RawJSON() string { return r.JSON.raw }
func (r *BrandSettingsEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandSettingsEmailFooter struct {
	Content        string `json:"content,nullable"`
	InheritDefault bool   `json:"inheritDefault,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content        respjson.Field
		InheritDefault respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandSettingsEmailFooter) RawJSON() string { return r.JSON.raw }
func (r *BrandSettingsEmailFooter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandSettingsEmailHead struct {
	InheritDefault bool   `json:"inheritDefault,required"`
	Content        string `json:"content,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InheritDefault respjson.Field
		Content        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandSettingsEmailHead) RawJSON() string { return r.JSON.raw }
func (r *BrandSettingsEmailHead) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandSettingsEmailHeader struct {
	Logo           BrandSettingsEmailHeaderLogo `json:"logo,required"`
	BarColor       string                       `json:"barColor,nullable"`
	InheritDefault bool                         `json:"inheritDefault,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Logo           respjson.Field
		BarColor       respjson.Field
		InheritDefault respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandSettingsEmailHeader) RawJSON() string { return r.JSON.raw }
func (r *BrandSettingsEmailHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandSettingsEmailHeaderLogo struct {
	Href  string `json:"href,nullable"`
	Image string `json:"image,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Href        respjson.Field
		Image       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandSettingsEmailHeaderLogo) RawJSON() string { return r.JSON.raw }
func (r *BrandSettingsEmailHeaderLogo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandSettingsEmailTemplateOverride struct {
	Mjml                  BrandTemplate `json:"mjml,required"`
	FooterBackgroundColor string        `json:"footerBackgroundColor,nullable"`
	FooterFullWidth       bool          `json:"footerFullWidth,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Mjml                  respjson.Field
		FooterBackgroundColor respjson.Field
		FooterFullWidth       respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
	BrandTemplate
}

// Returns the unmodified JSON received from the API
func (r BrandSettingsEmailTemplateOverride) RawJSON() string { return r.JSON.raw }
func (r *BrandSettingsEmailTemplateOverride) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandSettingsInapp struct {
	Colors             BrandSettingsInappColors           `json:"colors,required"`
	Icons              BrandSettingsInappIcons            `json:"icons,required"`
	WidgetBackground   BrandSettingsInappWidgetBackground `json:"widgetBackground,required"`
	BorderRadius       string                             `json:"borderRadius,nullable"`
	DisableMessageIcon bool                               `json:"disableMessageIcon,nullable"`
	FontFamily         string                             `json:"fontFamily,nullable"`
	// Any of "top", "bottom", "left", "right".
	Placement string `json:"placement,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Colors             respjson.Field
		Icons              respjson.Field
		WidgetBackground   respjson.Field
		BorderRadius       respjson.Field
		DisableMessageIcon respjson.Field
		FontFamily         respjson.Field
		Placement          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandSettingsInapp) RawJSON() string { return r.JSON.raw }
func (r *BrandSettingsInapp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandSettingsInappColors struct {
	Primary     string            `json:"primary,nullable"`
	Secondary   string            `json:"secondary,nullable"`
	ExtraFields map[string]string `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Primary     respjson.Field
		Secondary   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandSettingsInappColors) RawJSON() string { return r.JSON.raw }
func (r *BrandSettingsInappColors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandSettingsInappIcons struct {
	Bell    string `json:"bell,nullable"`
	Message string `json:"message,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bell        respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandSettingsInappIcons) RawJSON() string { return r.JSON.raw }
func (r *BrandSettingsInappIcons) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandSettingsInappWidgetBackground struct {
	BottomColor string `json:"bottomColor,nullable"`
	TopColor    string `json:"topColor,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BottomColor respjson.Field
		TopColor    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandSettingsInappWidgetBackground) RawJSON() string { return r.JSON.raw }
func (r *BrandSettingsInappWidgetBackground) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandSettingsParam struct {
	Colors BrandSettingsColorsParam `json:"colors,omitzero"`
	Email  BrandSettingsEmailParam  `json:"email,omitzero"`
	Inapp  BrandSettingsInappParam  `json:"inapp,omitzero"`
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
	Primary     param.Opt[string] `json:"primary,omitzero"`
	Secondary   param.Opt[string] `json:"secondary,omitzero"`
	ExtraFields map[string]string `json:"-"`
	paramObj
}

func (r BrandSettingsColorsParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandSettingsColorsParam
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *BrandSettingsColorsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandSettingsEmailParam struct {
	Footer           BrandSettingsEmailFooterParam           `json:"footer,omitzero"`
	Head             BrandSettingsEmailHeadParam             `json:"head,omitzero"`
	Header           BrandSettingsEmailHeaderParam           `json:"header,omitzero"`
	TemplateOverride BrandSettingsEmailTemplateOverrideParam `json:"templateOverride,omitzero"`
	paramObj
}

func (r BrandSettingsEmailParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandSettingsEmailParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandSettingsEmailParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandSettingsEmailFooterParam struct {
	Content        param.Opt[string] `json:"content,omitzero"`
	InheritDefault param.Opt[bool]   `json:"inheritDefault,omitzero"`
	paramObj
}

func (r BrandSettingsEmailFooterParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandSettingsEmailFooterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandSettingsEmailFooterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property InheritDefault is required.
type BrandSettingsEmailHeadParam struct {
	InheritDefault bool              `json:"inheritDefault,required"`
	Content        param.Opt[string] `json:"content,omitzero"`
	paramObj
}

func (r BrandSettingsEmailHeadParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandSettingsEmailHeadParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandSettingsEmailHeadParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Logo is required.
type BrandSettingsEmailHeaderParam struct {
	Logo           BrandSettingsEmailHeaderLogoParam `json:"logo,omitzero,required"`
	BarColor       param.Opt[string]                 `json:"barColor,omitzero"`
	InheritDefault param.Opt[bool]                   `json:"inheritDefault,omitzero"`
	paramObj
}

func (r BrandSettingsEmailHeaderParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandSettingsEmailHeaderParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandSettingsEmailHeaderParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandSettingsEmailHeaderLogoParam struct {
	Href  param.Opt[string] `json:"href,omitzero"`
	Image param.Opt[string] `json:"image,omitzero"`
	paramObj
}

func (r BrandSettingsEmailHeaderLogoParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandSettingsEmailHeaderLogoParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandSettingsEmailHeaderLogoParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandSettingsEmailTemplateOverrideParam struct {
	Mjml                  BrandTemplateParam `json:"mjml,omitzero,required"`
	FooterBackgroundColor param.Opt[string]  `json:"footerBackgroundColor,omitzero"`
	FooterFullWidth       param.Opt[bool]    `json:"footerFullWidth,omitzero"`
	BrandTemplateParam
}

func (r BrandSettingsEmailTemplateOverrideParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandSettingsEmailTemplateOverrideParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// The properties Colors, Icons, WidgetBackground are required.
type BrandSettingsInappParam struct {
	Colors             BrandSettingsInappColorsParam           `json:"colors,omitzero,required"`
	Icons              BrandSettingsInappIconsParam            `json:"icons,omitzero,required"`
	WidgetBackground   BrandSettingsInappWidgetBackgroundParam `json:"widgetBackground,omitzero,required"`
	BorderRadius       param.Opt[string]                       `json:"borderRadius,omitzero"`
	DisableMessageIcon param.Opt[bool]                         `json:"disableMessageIcon,omitzero"`
	FontFamily         param.Opt[string]                       `json:"fontFamily,omitzero"`
	// Any of "top", "bottom", "left", "right".
	Placement string `json:"placement,omitzero"`
	paramObj
}

func (r BrandSettingsInappParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandSettingsInappParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandSettingsInappParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BrandSettingsInappParam](
		"placement", "top", "bottom", "left", "right",
	)
}

type BrandSettingsInappColorsParam struct {
	Primary     param.Opt[string] `json:"primary,omitzero"`
	Secondary   param.Opt[string] `json:"secondary,omitzero"`
	ExtraFields map[string]string `json:"-"`
	paramObj
}

func (r BrandSettingsInappColorsParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandSettingsInappColorsParam
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *BrandSettingsInappColorsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandSettingsInappIconsParam struct {
	Bell    param.Opt[string] `json:"bell,omitzero"`
	Message param.Opt[string] `json:"message,omitzero"`
	paramObj
}

func (r BrandSettingsInappIconsParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandSettingsInappIconsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandSettingsInappIconsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandSettingsInappWidgetBackgroundParam struct {
	BottomColor param.Opt[string] `json:"bottomColor,omitzero"`
	TopColor    param.Opt[string] `json:"topColor,omitzero"`
	paramObj
}

func (r BrandSettingsInappWidgetBackgroundParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandSettingsInappWidgetBackgroundParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandSettingsInappWidgetBackgroundParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandSnippets struct {
	Items []BrandSnippetsItem `json:"items,nullable"`
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
	Name  string `json:"name,required"`
	Value string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
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

type BrandSnippetsParam struct {
	Items []BrandSnippetsItemParam `json:"items,omitzero"`
	paramObj
}

func (r BrandSnippetsParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandSnippetsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandSnippetsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Value are required.
type BrandSnippetsItemParam struct {
	Name  string `json:"name,required"`
	Value string `json:"value,required"`
	paramObj
}

func (r BrandSnippetsItemParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandSnippetsItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandSnippetsItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandTemplate struct {
	Enabled               bool   `json:"enabled,required"`
	BackgroundColor       string `json:"backgroundColor,nullable"`
	BlocksBackgroundColor string `json:"blocksBackgroundColor,nullable"`
	Footer                string `json:"footer,nullable"`
	Head                  string `json:"head,nullable"`
	Header                string `json:"header,nullable"`
	Width                 string `json:"width,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled               respjson.Field
		BackgroundColor       respjson.Field
		BlocksBackgroundColor respjson.Field
		Footer                respjson.Field
		Head                  respjson.Field
		Header                respjson.Field
		Width                 respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandTemplate) RawJSON() string { return r.JSON.raw }
func (r *BrandTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BrandTemplate to a BrandTemplateParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BrandTemplateParam.Overrides()
func (r BrandTemplate) ToParam() BrandTemplateParam {
	return param.Override[BrandTemplateParam](json.RawMessage(r.RawJSON()))
}

// The property Enabled is required.
type BrandTemplateParam struct {
	Enabled               bool              `json:"enabled,required"`
	BackgroundColor       param.Opt[string] `json:"backgroundColor,omitzero"`
	BlocksBackgroundColor param.Opt[string] `json:"blocksBackgroundColor,omitzero"`
	Footer                param.Opt[string] `json:"footer,omitzero"`
	Head                  param.Opt[string] `json:"head,omitzero"`
	Header                param.Opt[string] `json:"header,omitzero"`
	Width                 param.Opt[string] `json:"width,omitzero"`
	paramObj
}

func (r BrandTemplateParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandTemplateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandTemplateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	Name     string             `json:"name,required"`
	ID       param.Opt[string]  `json:"id,omitzero"`
	Settings BrandSettingsParam `json:"settings,omitzero"`
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
