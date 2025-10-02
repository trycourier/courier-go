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

// TenantTemplateService contains methods and other services that help with
// interacting with the courier API.
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

type ElementalChannelNode struct {
	// The channel the contents of this element should be applied to. Can be `email`,
	// `push`, `direct_message`, `sms` or a provider such as slack
	Channel  string   `json:"channel,required"`
	Channels []string `json:"channels,nullable"`
	// An array of elements to apply to the channel. If `raw` has not been specified,
	// `elements` is `required`.
	Elements []ElementalNodeUnion `json:"elements,nullable"`
	If       string               `json:"if,nullable"`
	Loop     string               `json:"loop,nullable"`
	// Raw data to apply to the channel. If `elements` has not been specified, `raw` is
	// `required`.
	Raw map[string]any `json:"raw,nullable"`
	Ref string         `json:"ref,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Channel     respjson.Field
		Channels    respjson.Field
		Elements    respjson.Field
		If          respjson.Field
		Loop        respjson.Field
		Raw         respjson.Field
		Ref         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ElementalChannelNode) RawJSON() string { return r.JSON.raw }
func (r *ElementalChannelNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ElementalChannelNode to a ElementalChannelNodeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ElementalChannelNodeParam.Overrides()
func (r ElementalChannelNode) ToParam() ElementalChannelNodeParam {
	return param.Override[ElementalChannelNodeParam](json.RawMessage(r.RawJSON()))
}

// The property Channel is required.
type ElementalChannelNodeParam struct {
	// The channel the contents of this element should be applied to. Can be `email`,
	// `push`, `direct_message`, `sms` or a provider such as slack
	Channel  string            `json:"channel,required"`
	If       param.Opt[string] `json:"if,omitzero"`
	Loop     param.Opt[string] `json:"loop,omitzero"`
	Ref      param.Opt[string] `json:"ref,omitzero"`
	Channels []string          `json:"channels,omitzero"`
	// An array of elements to apply to the channel. If `raw` has not been specified,
	// `elements` is `required`.
	Elements []ElementalNodeUnionParam `json:"elements,omitzero"`
	// Raw data to apply to the channel. If `elements` has not been specified, `raw` is
	// `required`.
	Raw map[string]any `json:"raw,omitzero"`
	paramObj
}

func (r ElementalChannelNodeParam) MarshalJSON() (data []byte, err error) {
	type shadow ElementalChannelNodeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ElementalChannelNodeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ElementalContent struct {
	Elements []ElementalNodeUnion `json:"elements,required"`
	// For example, "2022-01-01"
	Version string `json:"version,required"`
	Brand   any    `json:"brand"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Elements    respjson.Field
		Version     respjson.Field
		Brand       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ElementalContent) RawJSON() string { return r.JSON.raw }
func (r *ElementalContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ElementalContent to a ElementalContentParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ElementalContentParam.Overrides()
func (r ElementalContent) ToParam() ElementalContentParam {
	return param.Override[ElementalContentParam](json.RawMessage(r.RawJSON()))
}

// The properties Elements, Version are required.
type ElementalContentParam struct {
	Elements []ElementalNodeUnionParam `json:"elements,omitzero,required"`
	// For example, "2022-01-01"
	Version string `json:"version,required"`
	Brand   any    `json:"brand,omitzero"`
	paramObj
}

func (r ElementalContentParam) MarshalJSON() (data []byte, err error) {
	type shadow ElementalContentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ElementalContentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ElementalGroupNode struct {
	// Sub elements to render.
	Elements []ElementalNodeUnion `json:"elements,required"`
	Channels []string             `json:"channels,nullable"`
	If       string               `json:"if,nullable"`
	Loop     string               `json:"loop,nullable"`
	Ref      string               `json:"ref,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Elements    respjson.Field
		Channels    respjson.Field
		If          respjson.Field
		Loop        respjson.Field
		Ref         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ElementalGroupNode) RawJSON() string { return r.JSON.raw }
func (r *ElementalGroupNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ElementalGroupNode to a ElementalGroupNodeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ElementalGroupNodeParam.Overrides()
func (r ElementalGroupNode) ToParam() ElementalGroupNodeParam {
	return param.Override[ElementalGroupNodeParam](json.RawMessage(r.RawJSON()))
}

// The property Elements is required.
type ElementalGroupNodeParam struct {
	// Sub elements to render.
	Elements []ElementalNodeUnionParam `json:"elements,omitzero,required"`
	If       param.Opt[string]         `json:"if,omitzero"`
	Loop     param.Opt[string]         `json:"loop,omitzero"`
	Ref      param.Opt[string]         `json:"ref,omitzero"`
	Channels []string                  `json:"channels,omitzero"`
	paramObj
}

func (r ElementalGroupNodeParam) MarshalJSON() (data []byte, err error) {
	type shadow ElementalGroupNodeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ElementalGroupNodeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ElementalNodeUnion contains all possible properties and values from
// [ElementalNodeObject], [ElementalNodeObject], [ElementalNodeObject],
// [ElementalNodeObject], [ElementalNodeObject], [ElementalNodeObject],
// [ElementalNodeObject], [ElementalNodeObject].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ElementalNodeUnion struct {
	Channels []string `json:"channels"`
	If       string   `json:"if"`
	Loop     string   `json:"loop"`
	Ref      string   `json:"ref"`
	// This field is from variant [ElementalNodeObject].
	Type string `json:"type"`
	// This field is from variant [ElementalNodeObject].
	Channel  string               `json:"channel"`
	Elements []ElementalNodeUnion `json:"elements"`
	// This field is from variant [ElementalNodeObject].
	Raw  map[string]any `json:"raw"`
	JSON struct {
		Channels respjson.Field
		If       respjson.Field
		Loop     respjson.Field
		Ref      respjson.Field
		Type     respjson.Field
		Channel  respjson.Field
		Elements respjson.Field
		Raw      respjson.Field
		raw      string
	} `json:"-"`
}

func (u ElementalNodeUnion) AsElementalNodeObject() (v ElementalNodeObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ElementalNodeUnion) AsVariant2() (v ElementalNodeObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ElementalNodeUnion) AsVariant3() (v ElementalNodeObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ElementalNodeUnion) RawJSON() string { return u.JSON.raw }

func (r *ElementalNodeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ElementalNodeUnion to a ElementalNodeUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ElementalNodeUnionParam.Overrides()
func (r ElementalNodeUnion) ToParam() ElementalNodeUnionParam {
	return param.Override[ElementalNodeUnionParam](json.RawMessage(r.RawJSON()))
}

type ElementalNodeObject struct {
	Channels []string `json:"channels,nullable"`
	If       string   `json:"if,nullable"`
	Loop     string   `json:"loop,nullable"`
	Ref      string   `json:"ref,nullable"`
	// Any of "text".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Channels    respjson.Field
		If          respjson.Field
		Loop        respjson.Field
		Ref         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ElementalNodeObject) RawJSON() string { return r.JSON.raw }
func (r *ElementalNodeObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func ElementalNodeParamOfVariant3(channel string) ElementalNodeUnionParam {
	var variant ElementalNodeObjectParam
	variant.Channel = channel
	return ElementalNodeUnionParam{OfVariant3: &variant}
}

func ElementalNodeParamOfVariant3(elements []ElementalNodeUnionParam) ElementalNodeUnionParam {
	var variant ElementalNodeObjectParam
	variant.Elements = elements
	return ElementalNodeUnionParam{OfVariant3: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ElementalNodeUnionParam struct {
	OfElementalNodeObject *ElementalNodeObjectParam `json:",omitzero,inline"`
	OfVariant2            *ElementalNodeObjectParam `json:",omitzero,inline"`
	OfVariant3            *ElementalNodeObjectParam `json:",omitzero,inline"`
	paramUnion
}

func (u ElementalNodeUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfElementalNodeObject,
		u.OfVariant2,
		u.OfVariant3,
		u.OfVariant3,
		u.OfVariant3,
		u.OfVariant3,
		u.OfVariant3,
		u.OfVariant3)
}
func (u *ElementalNodeUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ElementalNodeUnionParam) asAny() any {
	if !param.IsOmitted(u.OfElementalNodeObject) {
		return u.OfElementalNodeObject
	} else if !param.IsOmitted(u.OfVariant2) {
		return u.OfVariant2
	} else if !param.IsOmitted(u.OfVariant3) {
		return u.OfVariant3
	} else if !param.IsOmitted(u.OfVariant3) {
		return u.OfVariant3
	} else if !param.IsOmitted(u.OfVariant3) {
		return u.OfVariant3
	} else if !param.IsOmitted(u.OfVariant3) {
		return u.OfVariant3
	} else if !param.IsOmitted(u.OfVariant3) {
		return u.OfVariant3
	} else if !param.IsOmitted(u.OfVariant3) {
		return u.OfVariant3
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetChannel() *string {
	if vt := u.OfVariant3; vt != nil {
		return &vt.Channel
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetRaw() map[string]any {
	if vt := u.OfVariant3; vt != nil {
		return vt.Raw
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetIf() *string {
	if vt := u.OfElementalNodeObject; vt != nil && vt.If.Valid() {
		return &vt.If.Value
	} else if vt := u.OfVariant2; vt != nil && vt.If.Valid() {
		return &vt.If.Value
	} else if vt := u.OfVariant3; vt != nil && vt.If.Valid() {
		return &vt.If.Value
	} else if vt := u.OfVariant3; vt != nil && vt.If.Valid() {
		return &vt.If.Value
	} else if vt := u.OfVariant3; vt != nil && vt.If.Valid() {
		return &vt.If.Value
	} else if vt := u.OfVariant3; vt != nil && vt.If.Valid() {
		return &vt.If.Value
	} else if vt := u.OfVariant3; vt != nil && vt.If.Valid() {
		return &vt.If.Value
	} else if vt := u.OfVariant3; vt != nil && vt.If.Valid() {
		return &vt.If.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetLoop() *string {
	if vt := u.OfElementalNodeObject; vt != nil && vt.Loop.Valid() {
		return &vt.Loop.Value
	} else if vt := u.OfVariant2; vt != nil && vt.Loop.Valid() {
		return &vt.Loop.Value
	} else if vt := u.OfVariant3; vt != nil && vt.Loop.Valid() {
		return &vt.Loop.Value
	} else if vt := u.OfVariant3; vt != nil && vt.Loop.Valid() {
		return &vt.Loop.Value
	} else if vt := u.OfVariant3; vt != nil && vt.Loop.Valid() {
		return &vt.Loop.Value
	} else if vt := u.OfVariant3; vt != nil && vt.Loop.Valid() {
		return &vt.Loop.Value
	} else if vt := u.OfVariant3; vt != nil && vt.Loop.Valid() {
		return &vt.Loop.Value
	} else if vt := u.OfVariant3; vt != nil && vt.Loop.Valid() {
		return &vt.Loop.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetRef() *string {
	if vt := u.OfElementalNodeObject; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	} else if vt := u.OfVariant2; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	} else if vt := u.OfVariant3; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	} else if vt := u.OfVariant3; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	} else if vt := u.OfVariant3; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	} else if vt := u.OfVariant3; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	} else if vt := u.OfVariant3; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	} else if vt := u.OfVariant3; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetType() *string {
	if vt := u.OfElementalNodeObject; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfVariant2; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfVariant3; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfVariant3; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfVariant3; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfVariant3; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfVariant3; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfVariant3; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's Channels property, if present.
func (u ElementalNodeUnionParam) GetChannels() []string {
	if vt := u.OfElementalNodeObject; vt != nil {
		return vt.Channels
	} else if vt := u.OfVariant2; vt != nil {
		return vt.Channels
	} else if vt := u.OfVariant3; vt != nil {
		return vt.Channels
	} else if vt := u.OfVariant3; vt != nil {
		return vt.Channels
	} else if vt := u.OfVariant3; vt != nil {
		return vt.Channels
	} else if vt := u.OfVariant3; vt != nil {
		return vt.Channels
	} else if vt := u.OfVariant3; vt != nil {
		return vt.Channels
	} else if vt := u.OfVariant3; vt != nil {
		return vt.Channels
	}
	return nil
}

// Returns a pointer to the underlying variant's Elements property, if present.
func (u ElementalNodeUnionParam) GetElements() []ElementalNodeUnionParam {
	if vt := u.OfVariant3; vt != nil {
		return vt.Elements
	} else if vt := u.OfVariant3; vt != nil {
		return vt.Elements
	}
	return nil
}

type ElementalNodeObjectParam struct {
	If       param.Opt[string] `json:"if,omitzero"`
	Loop     param.Opt[string] `json:"loop,omitzero"`
	Ref      param.Opt[string] `json:"ref,omitzero"`
	Channels []string          `json:"channels,omitzero"`
	// Any of "text".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ElementalNodeObjectParam) MarshalJSON() (data []byte, err error) {
	type shadow ElementalNodeObjectParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ElementalNodeObjectParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ElementalNodeObjectParam](
		"type", "text",
	)
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
	Routing MessageRouting `json:"routing,required"`
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
