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

type Alignment string

const (
	AlignmentCenter Alignment = "center"
	AlignmentLeft   Alignment = "left"
	AlignmentRight  Alignment = "right"
	AlignmentFull   Alignment = "full"
)

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

type ElementalBaseNode struct {
	Channels []string `json:"channels,nullable"`
	If       string   `json:"if,nullable"`
	Loop     string   `json:"loop,nullable"`
	Ref      string   `json:"ref,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Channels    respjson.Field
		If          respjson.Field
		Loop        respjson.Field
		Ref         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ElementalBaseNode) RawJSON() string { return r.JSON.raw }
func (r *ElementalBaseNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ElementalBaseNode to a ElementalBaseNodeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ElementalBaseNodeParam.Overrides()
func (r ElementalBaseNode) ToParam() ElementalBaseNodeParam {
	return param.Override[ElementalBaseNodeParam](json.RawMessage(r.RawJSON()))
}

type ElementalBaseNodeParam struct {
	If       param.Opt[string] `json:"if,omitzero"`
	Loop     param.Opt[string] `json:"loop,omitzero"`
	Ref      param.Opt[string] `json:"ref,omitzero"`
	Channels []string          `json:"channels,omitzero"`
	paramObj
}

func (r ElementalBaseNodeParam) MarshalJSON() (data []byte, err error) {
	type shadow ElementalBaseNodeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ElementalBaseNodeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ElementalContent struct {
	Elements []ElementalContentElementUnion `json:"elements,required"`
	// For example, "2022-01-01"
	Version string `json:"version,required"`
	Brand   string `json:"brand,nullable"`
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

// ElementalContentElementUnion contains all possible properties and values from
// [ElementalContentElementObject], [ElementalContentElementObject],
// [ElementalContentElementObject], [ElementalContentElementObject],
// [ElementalContentElementObject], [ElementalContentElementObject],
// [ElementalContentElementObject].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ElementalContentElementUnion struct {
	// This field is from variant [ElementalContentElementObject],
	// [ElementalContentElementObject], [ElementalContentElementObject],
	// [ElementalContentElementObject], [ElementalContentElementObject],
	// [ElementalContentElementObject].
	Channels []string `json:"channels"`
	// This field is from variant [ElementalContentElementObject],
	// [ElementalContentElementObject], [ElementalContentElementObject],
	// [ElementalContentElementObject], [ElementalContentElementObject],
	// [ElementalContentElementObject].
	If string `json:"if"`
	// This field is from variant [ElementalContentElementObject],
	// [ElementalContentElementObject], [ElementalContentElementObject],
	// [ElementalContentElementObject], [ElementalContentElementObject],
	// [ElementalContentElementObject].
	Loop string `json:"loop"`
	// This field is from variant [ElementalContentElementObject],
	// [ElementalContentElementObject], [ElementalContentElementObject],
	// [ElementalContentElementObject], [ElementalContentElementObject],
	// [ElementalContentElementObject].
	Ref string `json:"ref"`
	// This field is from variant [ElementalContentElementObject].
	Type string `json:"type"`
	// This field is from variant [ElementalContentElementObject].
	ActionID string `json:"action_id"`
	// This field is from variant [ElementalContentElementObject].
	Align Alignment `json:"align"`
	// This field is from variant [ElementalContentElementObject].
	BackgroundColor string `json:"background_color"`
	// This field is from variant [ElementalContentElementObject].
	Content string `json:"content"`
	// This field is from variant [ElementalContentElementObject].
	Href string `json:"href"`
	// This field is from variant [ElementalContentElementObject].
	Locales map[string]ElementalContentElementObjectLocale `json:"locales"`
	// This field is from variant [ElementalContentElementObject].
	Style string `json:"style"`
	JSON  struct {
		Channels        respjson.Field
		If              respjson.Field
		Loop            respjson.Field
		Ref             respjson.Field
		Type            respjson.Field
		ActionID        respjson.Field
		Align           respjson.Field
		BackgroundColor respjson.Field
		Content         respjson.Field
		Href            respjson.Field
		Locales         respjson.Field
		Style           respjson.Field
		raw             string
	} `json:"-"`
}

func (u ElementalContentElementUnion) AsElementalContentElementObject() (v ElementalContentElementObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ElementalContentElementUnion) AsVariant2() (v ElementalContentElementObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ElementalContentElementUnion) AsVariant3() (v ElementalContentElementObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ElementalContentElementUnion) RawJSON() string { return u.JSON.raw }

func (r *ElementalContentElementUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ElementalContentElementObject struct {
	// Any of "text".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ElementalBaseNode
}

// Returns the unmodified JSON received from the API
func (r ElementalContentElementObject) RawJSON() string { return r.JSON.raw }
func (r *ElementalContentElementObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Elements, Version are required.
type ElementalContentParam struct {
	Elements []ElementalContentElementUnionParam `json:"elements,omitzero,required"`
	// For example, "2022-01-01"
	Version string            `json:"version,required"`
	Brand   param.Opt[string] `json:"brand,omitzero"`
	paramObj
}

func (r ElementalContentParam) MarshalJSON() (data []byte, err error) {
	type shadow ElementalContentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ElementalContentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ElementalContentElementUnionParam struct {
	OfElementalContentElementObject *ElementalContentElementObjectParam `json:",omitzero,inline"`
	OfVariant2                      *ElementalContentElementObjectParam `json:",omitzero,inline"`
	OfVariant3                      *ElementalContentElementObjectParam `json:",omitzero,inline"`
	paramUnion
}

func (u ElementalContentElementUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfElementalContentElementObject,
		u.OfVariant2,
		u.OfVariant3,
		u.OfVariant3,
		u.OfVariant3,
		u.OfVariant3,
		u.OfVariant3)
}
func (u *ElementalContentElementUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ElementalContentElementUnionParam) asAny() any {
	if !param.IsOmitted(u.OfElementalContentElementObject) {
		return u.OfElementalContentElementObject
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
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalContentElementUnionParam) GetActionID() *string {
	if vt := u.OfVariant3; vt != nil && vt.ActionID.Valid() {
		return &vt.ActionID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalContentElementUnionParam) GetAlign() *string {
	if vt := u.OfVariant3; vt != nil {
		return (*string)(&vt.Align)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalContentElementUnionParam) GetBackgroundColor() *string {
	if vt := u.OfVariant3; vt != nil && vt.BackgroundColor.Valid() {
		return &vt.BackgroundColor.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalContentElementUnionParam) GetContent() *string {
	if vt := u.OfVariant3; vt != nil && vt.Content.Valid() {
		return &vt.Content.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalContentElementUnionParam) GetHref() *string {
	if vt := u.OfVariant3; vt != nil && vt.Href.Valid() {
		return &vt.Href.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalContentElementUnionParam) GetLocales() map[string]ElementalContentElementObjectLocaleParam {
	if vt := u.OfVariant3; vt != nil {
		return vt.Locales
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalContentElementUnionParam) GetStyle() *string {
	if vt := u.OfVariant3; vt != nil {
		return &vt.Style
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalContentElementUnionParam) GetIf() *string {
	if vt := u.OfElementalContentElementObject; vt != nil && vt.If.Valid() {
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
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalContentElementUnionParam) GetLoop() *string {
	if vt := u.OfElementalContentElementObject; vt != nil && vt.Loop.Valid() {
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
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalContentElementUnionParam) GetRef() *string {
	if vt := u.OfElementalContentElementObject; vt != nil && vt.Ref.Valid() {
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
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalContentElementUnionParam) GetType() *string {
	if vt := u.OfElementalContentElementObject; vt != nil {
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
	}
	return nil
}

// Returns a pointer to the underlying variant's Channels property, if present.
func (u ElementalContentElementUnionParam) GetChannels() []string {
	if vt := u.OfElementalContentElementObject; vt != nil {
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
	}
	return nil
}

type ElementalContentElementObjectParam struct {
	Type string `json:"type,omitzero"`
	ElementalBaseNodeParam
}

func (r ElementalContentElementObjectParam) MarshalJSON() (data []byte, err error) {
	type shadow ElementalContentElementObjectParam
	return param.MarshalObject(r, (*shadow)(&r))
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
