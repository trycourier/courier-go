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

	"github.com/trycourier/courier-go/v3/internal/apijson"
	"github.com/trycourier/courier-go/v3/internal/apiquery"
	"github.com/trycourier/courier-go/v3/internal/requestconfig"
	"github.com/trycourier/courier-go/v3/option"
	"github.com/trycourier/courier-go/v3/packages/param"
	"github.com/trycourier/courier-go/v3/packages/respjson"
	"github.com/trycourier/courier-go/v3/shared"
)

// BrandService contains methods and other services that help with interacting with
// the Courier API.
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

type BrandColors struct {
	Primary     string            `json:"primary"`
	Secondary   string            `json:"secondary"`
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
func (r BrandColors) RawJSON() string { return r.JSON.raw }
func (r *BrandColors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BrandColors to a BrandColorsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BrandColorsParam.Overrides()
func (r BrandColors) ToParam() BrandColorsParam {
	return param.Override[BrandColorsParam](json.RawMessage(r.RawJSON()))
}

type BrandColorsParam struct {
	Primary     param.Opt[string] `json:"primary,omitzero"`
	Secondary   param.Opt[string] `json:"secondary,omitzero"`
	ExtraFields map[string]string `json:"-"`
	paramObj
}

func (r BrandColorsParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandColorsParam
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *BrandColorsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandSettings struct {
	Colors BrandColors        `json:"colors,nullable"`
	Email  BrandSettingsEmail `json:"email,nullable"`
	Inapp  BrandSettingsInApp `json:"inapp,nullable"`
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

type BrandSettingsParam struct {
	Colors BrandColorsParam        `json:"colors,omitzero"`
	Email  BrandSettingsEmailParam `json:"email,omitzero"`
	Inapp  BrandSettingsInAppParam `json:"inapp,omitzero"`
	paramObj
}

func (r BrandSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandSettingsEmail struct {
	Footer           EmailFooter                        `json:"footer,nullable"`
	Head             EmailHead                          `json:"head,nullable"`
	Header           EmailHeader                        `json:"header,nullable"`
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

// ToParam converts this BrandSettingsEmail to a BrandSettingsEmailParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BrandSettingsEmailParam.Overrides()
func (r BrandSettingsEmail) ToParam() BrandSettingsEmailParam {
	return param.Override[BrandSettingsEmailParam](json.RawMessage(r.RawJSON()))
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

type BrandSettingsEmailParam struct {
	TemplateOverride BrandSettingsEmailTemplateOverrideParam `json:"templateOverride,omitzero"`
	Footer           EmailFooterParam                        `json:"footer,omitzero"`
	Head             EmailHeadParam                          `json:"head,omitzero"`
	Header           EmailHeaderParam                        `json:"header,omitzero"`
	paramObj
}

func (r BrandSettingsEmailParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandSettingsEmailParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandSettingsEmailParam) UnmarshalJSON(data []byte) error {
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

type BrandSettingsInApp struct {
	Colors             BrandColors      `json:"colors,required"`
	Icons              Icons            `json:"icons,required"`
	WidgetBackground   WidgetBackground `json:"widgetBackground,required"`
	BorderRadius       string           `json:"borderRadius,nullable"`
	DisableMessageIcon bool             `json:"disableMessageIcon,nullable"`
	FontFamily         string           `json:"fontFamily,nullable"`
	// Any of "top", "bottom", "left", "right".
	Placement BrandSettingsInAppPlacement `json:"placement,nullable"`
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
func (r BrandSettingsInApp) RawJSON() string { return r.JSON.raw }
func (r *BrandSettingsInApp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BrandSettingsInApp to a BrandSettingsInAppParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BrandSettingsInAppParam.Overrides()
func (r BrandSettingsInApp) ToParam() BrandSettingsInAppParam {
	return param.Override[BrandSettingsInAppParam](json.RawMessage(r.RawJSON()))
}

type BrandSettingsInAppPlacement string

const (
	BrandSettingsInAppPlacementTop    BrandSettingsInAppPlacement = "top"
	BrandSettingsInAppPlacementBottom BrandSettingsInAppPlacement = "bottom"
	BrandSettingsInAppPlacementLeft   BrandSettingsInAppPlacement = "left"
	BrandSettingsInAppPlacementRight  BrandSettingsInAppPlacement = "right"
)

// The properties Colors, Icons, WidgetBackground are required.
type BrandSettingsInAppParam struct {
	Colors             BrandColorsParam      `json:"colors,omitzero,required"`
	Icons              IconsParam            `json:"icons,omitzero,required"`
	WidgetBackground   WidgetBackgroundParam `json:"widgetBackground,omitzero,required"`
	BorderRadius       param.Opt[string]     `json:"borderRadius,omitzero"`
	DisableMessageIcon param.Opt[bool]       `json:"disableMessageIcon,omitzero"`
	FontFamily         param.Opt[string]     `json:"fontFamily,omitzero"`
	// Any of "top", "bottom", "left", "right".
	Placement BrandSettingsInAppPlacement `json:"placement,omitzero"`
	paramObj
}

func (r BrandSettingsInAppParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandSettingsInAppParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandSettingsInAppParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandSnippet struct {
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
func (r BrandSnippet) RawJSON() string { return r.JSON.raw }
func (r *BrandSnippet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BrandSnippet to a BrandSnippetParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BrandSnippetParam.Overrides()
func (r BrandSnippet) ToParam() BrandSnippetParam {
	return param.Override[BrandSnippetParam](json.RawMessage(r.RawJSON()))
}

// The properties Name, Value are required.
type BrandSnippetParam struct {
	Name  string `json:"name,required"`
	Value string `json:"value,required"`
	paramObj
}

func (r BrandSnippetParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandSnippetParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandSnippetParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandSnippets struct {
	Items []BrandSnippet `json:"items,nullable"`
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

type BrandSnippetsParam struct {
	Items []BrandSnippetParam `json:"items,omitzero"`
	paramObj
}

func (r BrandSnippetsParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandSnippetsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandSnippetsParam) UnmarshalJSON(data []byte) error {
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

type EmailFooter struct {
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
func (r EmailFooter) RawJSON() string { return r.JSON.raw }
func (r *EmailFooter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EmailFooter to a EmailFooterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EmailFooterParam.Overrides()
func (r EmailFooter) ToParam() EmailFooterParam {
	return param.Override[EmailFooterParam](json.RawMessage(r.RawJSON()))
}

type EmailFooterParam struct {
	Content        param.Opt[string] `json:"content,omitzero"`
	InheritDefault param.Opt[bool]   `json:"inheritDefault,omitzero"`
	paramObj
}

func (r EmailFooterParam) MarshalJSON() (data []byte, err error) {
	type shadow EmailFooterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailFooterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailHead struct {
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
func (r EmailHead) RawJSON() string { return r.JSON.raw }
func (r *EmailHead) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EmailHead to a EmailHeadParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EmailHeadParam.Overrides()
func (r EmailHead) ToParam() EmailHeadParam {
	return param.Override[EmailHeadParam](json.RawMessage(r.RawJSON()))
}

// The property InheritDefault is required.
type EmailHeadParam struct {
	InheritDefault bool              `json:"inheritDefault,required"`
	Content        param.Opt[string] `json:"content,omitzero"`
	paramObj
}

func (r EmailHeadParam) MarshalJSON() (data []byte, err error) {
	type shadow EmailHeadParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailHeadParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailHeader struct {
	Logo           Logo   `json:"logo,required"`
	BarColor       string `json:"barColor,nullable"`
	InheritDefault bool   `json:"inheritDefault,nullable"`
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
func (r EmailHeader) RawJSON() string { return r.JSON.raw }
func (r *EmailHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EmailHeader to a EmailHeaderParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EmailHeaderParam.Overrides()
func (r EmailHeader) ToParam() EmailHeaderParam {
	return param.Override[EmailHeaderParam](json.RawMessage(r.RawJSON()))
}

// The property Logo is required.
type EmailHeaderParam struct {
	Logo           LogoParam         `json:"logo,omitzero,required"`
	BarColor       param.Opt[string] `json:"barColor,omitzero"`
	InheritDefault param.Opt[bool]   `json:"inheritDefault,omitzero"`
	paramObj
}

func (r EmailHeaderParam) MarshalJSON() (data []byte, err error) {
	type shadow EmailHeaderParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailHeaderParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Icons struct {
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
func (r Icons) RawJSON() string { return r.JSON.raw }
func (r *Icons) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Icons to a IconsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// IconsParam.Overrides()
func (r Icons) ToParam() IconsParam {
	return param.Override[IconsParam](json.RawMessage(r.RawJSON()))
}

type IconsParam struct {
	Bell    param.Opt[string] `json:"bell,omitzero"`
	Message param.Opt[string] `json:"message,omitzero"`
	paramObj
}

func (r IconsParam) MarshalJSON() (data []byte, err error) {
	type shadow IconsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IconsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Logo struct {
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
func (r Logo) RawJSON() string { return r.JSON.raw }
func (r *Logo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Logo to a LogoParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// LogoParam.Overrides()
func (r Logo) ToParam() LogoParam {
	return param.Override[LogoParam](json.RawMessage(r.RawJSON()))
}

type LogoParam struct {
	Href  param.Opt[string] `json:"href,omitzero"`
	Image param.Opt[string] `json:"image,omitzero"`
	paramObj
}

func (r LogoParam) MarshalJSON() (data []byte, err error) {
	type shadow LogoParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogoParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WidgetBackground struct {
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
func (r WidgetBackground) RawJSON() string { return r.JSON.raw }
func (r *WidgetBackground) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WidgetBackground to a WidgetBackgroundParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WidgetBackgroundParam.Overrides()
func (r WidgetBackground) ToParam() WidgetBackgroundParam {
	return param.Override[WidgetBackgroundParam](json.RawMessage(r.RawJSON()))
}

type WidgetBackgroundParam struct {
	BottomColor param.Opt[string] `json:"bottomColor,omitzero"`
	TopColor    param.Opt[string] `json:"topColor,omitzero"`
	paramObj
}

func (r WidgetBackgroundParam) MarshalJSON() (data []byte, err error) {
	type shadow WidgetBackgroundParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WidgetBackgroundParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandListResponse struct {
	Paging  shared.Paging `json:"paging,required"`
	Results []Brand       `json:"results,required"`
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
