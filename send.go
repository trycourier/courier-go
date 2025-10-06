// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/trycourier/courier-go/internal/apijson"
	"github.com/trycourier/courier-go/internal/requestconfig"
	"github.com/trycourier/courier-go/option"
	"github.com/trycourier/courier-go/packages/param"
	"github.com/trycourier/courier-go/packages/respjson"
	"github.com/trycourier/courier-go/shared"
)

// SendService contains methods and other services that help with interacting with
// the courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSendService] method instead.
type SendService struct {
	Options []option.RequestOption
}

// NewSendService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSendService(opts ...option.RequestOption) (r SendService) {
	r = SendService{}
	r.Options = opts
	return
}

// Use the send API to send a message to one or more recipients.
func (r *SendService) Message(ctx context.Context, body SendMessageParams, opts ...option.RequestOption) (res *SendMessageResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "send"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Alignment string

const (
	AlignmentCenter Alignment = "center"
	AlignmentLeft   Alignment = "left"
	AlignmentRight  Alignment = "right"
	AlignmentFull   Alignment = "full"
)

// ContentUnion contains all possible properties and values from
// [ContentElementalContentSugar], [ElementalContent].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ContentUnion struct {
	// This field is from variant [ContentElementalContentSugar].
	Body string `json:"body"`
	// This field is from variant [ContentElementalContentSugar].
	Title string `json:"title"`
	// This field is from variant [ElementalContent].
	Elements []ElementalNodeUnion `json:"elements"`
	// This field is from variant [ElementalContent].
	Version string `json:"version"`
	// This field is from variant [ElementalContent].
	Brand string `json:"brand"`
	JSON  struct {
		Body     respjson.Field
		Title    respjson.Field
		Elements respjson.Field
		Version  respjson.Field
		Brand    respjson.Field
		raw      string
	} `json:"-"`
}

func (u ContentUnion) AsElementalContentSugar() (v ContentElementalContentSugar) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContentUnion) AsElementalContent() (v ElementalContent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ContentUnion) RawJSON() string { return u.JSON.raw }

func (r *ContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ContentUnion to a ContentUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ContentUnionParam.Overrides()
func (r ContentUnion) ToParam() ContentUnionParam {
	return param.Override[ContentUnionParam](json.RawMessage(r.RawJSON()))
}

// Syntactic sugar to provide a fast shorthand for Courier Elemental Blocks.
type ContentElementalContentSugar struct {
	// The text content displayed in the notification.
	Body string `json:"body,required"`
	// Title/subject displayed by supported channels.
	Title string `json:"title,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContentElementalContentSugar) RawJSON() string { return r.JSON.raw }
func (r *ContentElementalContentSugar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func ContentParamOfElementalContentSugar(body string, title string) ContentUnionParam {
	var variant ContentElementalContentSugarParam
	variant.Body = body
	variant.Title = title
	return ContentUnionParam{OfElementalContentSugar: &variant}
}

func ContentParamOfElementalContent(elements []ElementalNodeUnionParam, version string) ContentUnionParam {
	var variant ElementalContentParam
	variant.Elements = elements
	variant.Version = version
	return ContentUnionParam{OfElementalContent: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ContentUnionParam struct {
	OfElementalContentSugar *ContentElementalContentSugarParam `json:",omitzero,inline"`
	OfElementalContent      *ElementalContentParam             `json:",omitzero,inline"`
	paramUnion
}

func (u ContentUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfElementalContentSugar, u.OfElementalContent)
}
func (u *ContentUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ContentUnionParam) asAny() any {
	if !param.IsOmitted(u.OfElementalContentSugar) {
		return u.OfElementalContentSugar
	} else if !param.IsOmitted(u.OfElementalContent) {
		return u.OfElementalContent
	}
	return nil
}

// Syntactic sugar to provide a fast shorthand for Courier Elemental Blocks.
//
// The properties Body, Title are required.
type ContentElementalContentSugarParam struct {
	// The text content displayed in the notification.
	Body string `json:"body,required"`
	// Title/subject displayed by supported channels.
	Title string `json:"title,required"`
	paramObj
}

func (r ContentElementalContentSugarParam) MarshalJSON() (data []byte, err error) {
	type shadow ContentElementalContentSugarParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContentElementalContentSugarParam) UnmarshalJSON(data []byte) error {
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

// The channel element allows a notification to be customized based on which
// channel it is sent through. For example, you may want to display a detailed
// message when the notification is sent through email, and a more concise message
// in a push notification. Channel elements are only valid as top-level elements;
// you cannot nest channel elements. If there is a channel element specified at the
// top-level of the document, all sibling elements must be channel elements. Note:
// As an alternative, most elements support a `channel` property. Which allows you
// to selectively display an individual element on a per channel basis. See the
// [control flow docs](https://www.courier.com/docs/platform/content/elemental/control-flow/)
// for more details.
type ElementalChannelNode struct {
	// The channel the contents of this element should be applied to. Can be `email`,
	// `push`, `direct_message`, `sms` or a provider such as slack
	Channel string `json:"channel,required"`
	// Raw data to apply to the channel. If `elements` has not been specified, `raw` is
	// `required`.
	Raw map[string]any `json:"raw,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Channel     respjson.Field
		Raw         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ElementalBaseNode
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

// The channel element allows a notification to be customized based on which
// channel it is sent through. For example, you may want to display a detailed
// message when the notification is sent through email, and a more concise message
// in a push notification. Channel elements are only valid as top-level elements;
// you cannot nest channel elements. If there is a channel element specified at the
// top-level of the document, all sibling elements must be channel elements. Note:
// As an alternative, most elements support a `channel` property. Which allows you
// to selectively display an individual element on a per channel basis. See the
// [control flow docs](https://www.courier.com/docs/platform/content/elemental/control-flow/)
// for more details.
type ElementalChannelNodeParam struct {
	// The channel the contents of this element should be applied to. Can be `email`,
	// `push`, `direct_message`, `sms` or a provider such as slack
	Channel string `json:"channel,required"`
	// Raw data to apply to the channel. If `elements` has not been specified, `raw` is
	// `required`.
	Raw map[string]any `json:"raw,omitzero"`
	ElementalBaseNodeParam
}

func (r ElementalChannelNodeParam) MarshalJSON() (data []byte, err error) {
	type shadow ElementalChannelNodeParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// ElementalNodeUnion contains all possible properties and values from
// [ElementalNodeObject], [ElementalNodeObject], [ElementalNodeObject],
// [ElementalNodeObject], [ElementalNodeObject], [ElementalNodeObject],
// [ElementalNodeObject].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ElementalNodeUnion struct {
	// This field is from variant [ElementalNodeObject], [ElementalNodeObject],
	// [ElementalNodeObject], [ElementalNodeObject], [ElementalNodeObject],
	// [ElementalNodeObject].
	Channels []string `json:"channels"`
	// This field is from variant [ElementalNodeObject], [ElementalNodeObject],
	// [ElementalNodeObject], [ElementalNodeObject], [ElementalNodeObject],
	// [ElementalNodeObject].
	If string `json:"if"`
	// This field is from variant [ElementalNodeObject], [ElementalNodeObject],
	// [ElementalNodeObject], [ElementalNodeObject], [ElementalNodeObject],
	// [ElementalNodeObject].
	Loop string `json:"loop"`
	// This field is from variant [ElementalNodeObject], [ElementalNodeObject],
	// [ElementalNodeObject], [ElementalNodeObject], [ElementalNodeObject],
	// [ElementalNodeObject].
	Ref string `json:"ref"`
	// This field is from variant [ElementalNodeObject].
	Type string `json:"type"`
	// This field is from variant [ElementalNodeObject].
	Channel string `json:"channel"`
	// This field is from variant [ElementalNodeObject].
	Raw map[string]any `json:"raw"`
	// This field is from variant [ElementalNodeObject].
	ActionID string `json:"action_id"`
	// This field is from variant [ElementalNodeObject].
	Align Alignment `json:"align"`
	// This field is from variant [ElementalNodeObject].
	BackgroundColor string `json:"background_color"`
	// This field is from variant [ElementalNodeObject].
	Content string `json:"content"`
	// This field is from variant [ElementalNodeObject].
	Href string `json:"href"`
	// This field is from variant [ElementalNodeObject].
	Locales map[string]ElementalNodeObjectLocale `json:"locales"`
	// This field is from variant [ElementalNodeObject].
	Style string `json:"style"`
	JSON  struct {
		Channels        respjson.Field
		If              respjson.Field
		Loop            respjson.Field
		Ref             respjson.Field
		Type            respjson.Field
		Channel         respjson.Field
		Raw             respjson.Field
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
func (r ElementalNodeObject) RawJSON() string { return r.JSON.raw }
func (r *ElementalNodeObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func ElementalNodeParamOfVariant3(channel string) ElementalNodeUnionParam {
	var variant ElementalNodeObjectParam
	variant.Channel = channel
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
func (u ElementalNodeUnionParam) GetActionID() *string {
	if vt := u.OfVariant3; vt != nil && vt.ActionID.Valid() {
		return &vt.ActionID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetAlign() *string {
	if vt := u.OfVariant3; vt != nil {
		return (*string)(&vt.Align)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetBackgroundColor() *string {
	if vt := u.OfVariant3; vt != nil && vt.BackgroundColor.Valid() {
		return &vt.BackgroundColor.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetContent() *string {
	if vt := u.OfVariant3; vt != nil && vt.Content.Valid() {
		return &vt.Content.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetHref() *string {
	if vt := u.OfVariant3; vt != nil && vt.Href.Valid() {
		return &vt.Href.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetLocales() map[string]ElementalNodeObjectLocaleParam {
	if vt := u.OfVariant3; vt != nil {
		return vt.Locales
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetStyle() *string {
	if vt := u.OfVariant3; vt != nil {
		return &vt.Style
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
	}
	return nil
}

type ElementalNodeObjectParam struct {
	Type string `json:"type,omitzero"`
	ElementalBaseNodeParam
}

func (r ElementalNodeObjectParam) MarshalJSON() (data []byte, err error) {
	type shadow ElementalNodeObjectParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type MessageContext struct {
	// Tenant id used to load brand/default preferences/context.
	TenantID string `json:"tenant_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TenantID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageContext) RawJSON() string { return r.JSON.raw }
func (r *MessageContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MessageContext to a MessageContextParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MessageContextParam.Overrides()
func (r MessageContext) ToParam() MessageContextParam {
	return param.Override[MessageContextParam](json.RawMessage(r.RawJSON()))
}

type MessageContextParam struct {
	// Tenant id used to load brand/default preferences/context.
	TenantID param.Opt[string] `json:"tenant_id,omitzero"`
	paramObj
}

func (r MessageContextParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageContextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageContextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageRouting struct {
	Channels []MessageRoutingChannelUnion `json:"channels,required"`
	// Any of "all", "single".
	Method MessageRoutingMethod `json:"method,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Channels    respjson.Field
		Method      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageRouting) RawJSON() string { return r.JSON.raw }
func (r *MessageRouting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MessageRouting to a MessageRoutingParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MessageRoutingParam.Overrides()
func (r MessageRouting) ToParam() MessageRoutingParam {
	return param.Override[MessageRoutingParam](json.RawMessage(r.RawJSON()))
}

type MessageRoutingMethod string

const (
	MessageRoutingMethodAll    MessageRoutingMethod = "all"
	MessageRoutingMethodSingle MessageRoutingMethod = "single"
)

// The properties Channels, Method are required.
type MessageRoutingParam struct {
	Channels []MessageRoutingChannelUnionParam `json:"channels,omitzero,required"`
	// Any of "all", "single".
	Method MessageRoutingMethod `json:"method,omitzero,required"`
	paramObj
}

func (r MessageRoutingParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageRoutingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageRoutingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MessageRoutingChannelUnion contains all possible properties and values from
// [string], [MessageRouting].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type MessageRoutingChannelUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is from variant [MessageRouting].
	Channels []MessageRoutingChannelUnion `json:"channels"`
	// This field is from variant [MessageRouting].
	Method MessageRoutingMethod `json:"method"`
	JSON   struct {
		OfString respjson.Field
		Channels respjson.Field
		Method   respjson.Field
		raw      string
	} `json:"-"`
}

func (u MessageRoutingChannelUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageRoutingChannelUnion) AsMessageRouting() (v MessageRouting) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MessageRoutingChannelUnion) RawJSON() string { return u.JSON.raw }

func (r *MessageRoutingChannelUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MessageRoutingChannelUnion to a
// MessageRoutingChannelUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MessageRoutingChannelUnionParam.Overrides()
func (r MessageRoutingChannelUnion) ToParam() MessageRoutingChannelUnionParam {
	return param.Override[MessageRoutingChannelUnionParam](json.RawMessage(r.RawJSON()))
}

func MessageRoutingChannelParamOfMessageRouting(channels []MessageRoutingChannelUnionParam, method MessageRoutingMethod) MessageRoutingChannelUnionParam {
	var variant MessageRoutingParam
	variant.Channels = channels
	variant.Method = method
	return MessageRoutingChannelUnionParam{OfMessageRouting: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MessageRoutingChannelUnionParam struct {
	OfString         param.Opt[string]    `json:",omitzero,inline"`
	OfMessageRouting *MessageRoutingParam `json:",omitzero,inline"`
	paramUnion
}

func (u MessageRoutingChannelUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfMessageRouting)
}
func (u *MessageRoutingChannelUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *MessageRoutingChannelUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfMessageRouting) {
		return u.OfMessageRouting
	}
	return nil
}

type Preference struct {
	// Any of "OPTED_IN", "OPTED_OUT", "REQUIRED".
	Status             PreferenceStatus           `json:"status,required"`
	ChannelPreferences []shared.ChannelPreference `json:"channel_preferences,nullable"`
	Rules              []shared.Rule              `json:"rules,nullable"`
	// Any of "subscription", "list", "recipient".
	Source PreferenceSource `json:"source,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status             respjson.Field
		ChannelPreferences respjson.Field
		Rules              respjson.Field
		Source             respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Preference) RawJSON() string { return r.JSON.raw }
func (r *Preference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Preference to a PreferenceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PreferenceParam.Overrides()
func (r Preference) ToParam() PreferenceParam {
	return param.Override[PreferenceParam](json.RawMessage(r.RawJSON()))
}

type PreferenceSource string

const (
	PreferenceSourceSubscription PreferenceSource = "subscription"
	PreferenceSourceList         PreferenceSource = "list"
	PreferenceSourceRecipient    PreferenceSource = "recipient"
)

// The property Status is required.
type PreferenceParam struct {
	// Any of "OPTED_IN", "OPTED_OUT", "REQUIRED".
	Status             PreferenceStatus                `json:"status,omitzero,required"`
	ChannelPreferences []shared.ChannelPreferenceParam `json:"channel_preferences,omitzero"`
	Rules              []shared.RuleParam              `json:"rules,omitzero"`
	// Any of "subscription", "list", "recipient".
	Source PreferenceSource `json:"source,omitzero"`
	paramObj
}

func (r PreferenceParam) MarshalJSON() (data []byte, err error) {
	type shadow PreferenceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PreferenceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecipientParam struct {
	// Use `tenant_id` instead.
	AccountID param.Opt[string] `json:"account_id,omitzero"`
	Email     param.Opt[string] `json:"email,omitzero"`
	// The user's preferred ISO 639-1 language code.
	Locale      param.Opt[string] `json:"locale,omitzero"`
	PhoneNumber param.Opt[string] `json:"phone_number,omitzero"`
	// Tenant id. Will load brand, default preferences and base context data.
	TenantID    param.Opt[string]         `json:"tenant_id,omitzero"`
	UserID      param.Opt[string]         `json:"user_id,omitzero"`
	Data        map[string]any            `json:"data,omitzero"`
	Preferences RecipientPreferencesParam `json:"preferences,omitzero"`
	// Context such as tenant_id to send the notification with.
	Context MessageContextParam `json:"context,omitzero"`
	paramObj
}

func (r RecipientParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecipientParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Notifications is required.
type RecipientPreferencesParam struct {
	Notifications map[string]PreferenceParam `json:"notifications,omitzero,required"`
	TemplateID    param.Opt[string]          `json:"templateId,omitzero"`
	Categories    map[string]PreferenceParam `json:"categories,omitzero"`
	paramObj
}

func (r RecipientPreferencesParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientPreferencesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecipientPreferencesParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UtmParam struct {
	Campaign param.Opt[string] `json:"campaign,omitzero"`
	Content  param.Opt[string] `json:"content,omitzero"`
	Medium   param.Opt[string] `json:"medium,omitzero"`
	Source   param.Opt[string] `json:"source,omitzero"`
	Term     param.Opt[string] `json:"term,omitzero"`
	paramObj
}

func (r UtmParam) MarshalJSON() (data []byte, err error) {
	type shadow UtmParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UtmParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SendMessageResponse struct {
	// A successful call to `POST /send` returns a `202` status code along with a
	// `requestId` in the response body. For single-recipient requests, the `requestId`
	// is the derived message_id. For multiple recipients, Courier assigns a unique
	// message_id to each derived message.
	RequestID string `json:"requestId,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RequestID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SendMessageResponse) RawJSON() string { return r.JSON.raw }
func (r *SendMessageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SendMessageParams struct {
	// The message property has the following primary top-level properties. They define
	// the destination and content of the message.
	Message SendMessageParamsMessage `json:"message,omitzero,required"`
	paramObj
}

func (r SendMessageParams) MarshalJSON() (data []byte, err error) {
	type shadow SendMessageParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendMessageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The message property has the following primary top-level properties. They define
// the destination and content of the message.
type SendMessageParamsMessage struct {
	BrandID param.Opt[string] `json:"brand_id,omitzero"`
	// Define run-time configuration for channels. Valid ChannelId's: email, sms, push,
	// inbox, direct_message, banner, webhook.
	Channels    map[string]SendMessageParamsMessageChannel  `json:"channels,omitzero"`
	Data        map[string]any                              `json:"data,omitzero"`
	Delay       SendMessageParamsMessageDelay               `json:"delay,omitzero"`
	Expiry      SendMessageParamsMessageExpiry              `json:"expiry,omitzero"`
	Metadata    SendMessageParamsMessageMetadata            `json:"metadata,omitzero"`
	Preferences SendMessageParamsMessagePreferences         `json:"preferences,omitzero"`
	Providers   map[string]SendMessageParamsMessageProvider `json:"providers,omitzero"`
	// Customize which channels/providers Courier may deliver the message through.
	Routing SendMessageParamsMessageRouting `json:"routing,omitzero"`
	Timeout SendMessageParamsMessageTimeout `json:"timeout,omitzero"`
	// The recipient or a list of recipients of the message
	To SendMessageParamsMessageToUnion `json:"to,omitzero"`
	// Describes content that will work for email, inbox, push, chat, or any channel
	// id.
	Content ContentUnionParam   `json:"content,omitzero"`
	Context MessageContextParam `json:"context,omitzero"`
	paramObj
}

func (r SendMessageParamsMessage) MarshalJSON() (data []byte, err error) {
	type shadow SendMessageParamsMessage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendMessageParamsMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SendMessageParamsMessageChannel struct {
	// Brand id used for rendering.
	BrandID param.Opt[string] `json:"brand_id,omitzero"`
	// JS conditional with access to data/profile.
	If       param.Opt[string]                       `json:"if,omitzero"`
	Metadata SendMessageParamsMessageChannelMetadata `json:"metadata,omitzero"`
	// Channel specific overrides.
	Override map[string]any `json:"override,omitzero"`
	// Providers enabled for this channel.
	Providers []string `json:"providers,omitzero"`
	// Defaults to `single`.
	//
	// Any of "all", "single".
	RoutingMethod string                                  `json:"routing_method,omitzero"`
	Timeouts      SendMessageParamsMessageChannelTimeouts `json:"timeouts,omitzero"`
	paramObj
}

func (r SendMessageParamsMessageChannel) MarshalJSON() (data []byte, err error) {
	type shadow SendMessageParamsMessageChannel
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendMessageParamsMessageChannel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SendMessageParamsMessageChannel](
		"routing_method", "all", "single",
	)
}

type SendMessageParamsMessageChannelMetadata struct {
	Utm UtmParam `json:"utm,omitzero"`
	paramObj
}

func (r SendMessageParamsMessageChannelMetadata) MarshalJSON() (data []byte, err error) {
	type shadow SendMessageParamsMessageChannelMetadata
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendMessageParamsMessageChannelMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SendMessageParamsMessageChannelTimeouts struct {
	Channel  param.Opt[int64] `json:"channel,omitzero"`
	Provider param.Opt[int64] `json:"provider,omitzero"`
	paramObj
}

func (r SendMessageParamsMessageChannelTimeouts) MarshalJSON() (data []byte, err error) {
	type shadow SendMessageParamsMessageChannelTimeouts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendMessageParamsMessageChannelTimeouts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SendMessageParamsMessageDelay struct {
	// The duration of the delay in milliseconds.
	Duration param.Opt[int64] `json:"duration,omitzero"`
	// ISO 8601 timestamp or opening_hours-like format.
	Until param.Opt[string] `json:"until,omitzero"`
	paramObj
}

func (r SendMessageParamsMessageDelay) MarshalJSON() (data []byte, err error) {
	type shadow SendMessageParamsMessageDelay
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendMessageParamsMessageDelay) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ExpiresIn is required.
type SendMessageParamsMessageExpiry struct {
	// Duration in ms or ISO8601 duration (e.g. P1DT4H).
	ExpiresIn SendMessageParamsMessageExpiryExpiresInUnion `json:"expires_in,omitzero,required"`
	// Epoch or ISO8601 timestamp with timezone.
	ExpiresAt param.Opt[string] `json:"expires_at,omitzero"`
	paramObj
}

func (r SendMessageParamsMessageExpiry) MarshalJSON() (data []byte, err error) {
	type shadow SendMessageParamsMessageExpiry
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendMessageParamsMessageExpiry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SendMessageParamsMessageExpiryExpiresInUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u SendMessageParamsMessageExpiryExpiresInUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *SendMessageParamsMessageExpiryExpiresInUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SendMessageParamsMessageExpiryExpiresInUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

type SendMessageParamsMessageMetadata struct {
	Event   param.Opt[string] `json:"event,omitzero"`
	TraceID param.Opt[string] `json:"trace_id,omitzero"`
	Tags    []string          `json:"tags,omitzero"`
	Utm     UtmParam          `json:"utm,omitzero"`
	paramObj
}

func (r SendMessageParamsMessageMetadata) MarshalJSON() (data []byte, err error) {
	type shadow SendMessageParamsMessageMetadata
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendMessageParamsMessageMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SubscriptionTopicID is required.
type SendMessageParamsMessagePreferences struct {
	// The subscription topic to apply to the message.
	SubscriptionTopicID string `json:"subscription_topic_id,required"`
	paramObj
}

func (r SendMessageParamsMessagePreferences) MarshalJSON() (data []byte, err error) {
	type shadow SendMessageParamsMessagePreferences
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendMessageParamsMessagePreferences) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SendMessageParamsMessageProvider struct {
	// JS conditional with access to data/profile.
	If       param.Opt[string]                        `json:"if,omitzero"`
	Timeouts param.Opt[int64]                         `json:"timeouts,omitzero"`
	Metadata SendMessageParamsMessageProviderMetadata `json:"metadata,omitzero"`
	// Provider-specific overrides.
	Override map[string]any `json:"override,omitzero"`
	paramObj
}

func (r SendMessageParamsMessageProvider) MarshalJSON() (data []byte, err error) {
	type shadow SendMessageParamsMessageProvider
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendMessageParamsMessageProvider) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SendMessageParamsMessageProviderMetadata struct {
	Utm UtmParam `json:"utm,omitzero"`
	paramObj
}

func (r SendMessageParamsMessageProviderMetadata) MarshalJSON() (data []byte, err error) {
	type shadow SendMessageParamsMessageProviderMetadata
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendMessageParamsMessageProviderMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Customize which channels/providers Courier may deliver the message through.
//
// The properties Channels, Method are required.
type SendMessageParamsMessageRouting struct {
	// A list of channels or providers (or nested routing rules).
	Channels []MessageRoutingChannelUnionParam `json:"channels,omitzero,required"`
	// Any of "all", "single".
	Method string `json:"method,omitzero,required"`
	paramObj
}

func (r SendMessageParamsMessageRouting) MarshalJSON() (data []byte, err error) {
	type shadow SendMessageParamsMessageRouting
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendMessageParamsMessageRouting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SendMessageParamsMessageRouting](
		"method", "all", "single",
	)
}

type SendMessageParamsMessageTimeout struct {
	Escalation param.Opt[int64] `json:"escalation,omitzero"`
	Message    param.Opt[int64] `json:"message,omitzero"`
	Channel    map[string]int64 `json:"channel,omitzero"`
	// Any of "no-escalation", "delivered", "viewed", "engaged".
	Criteria string           `json:"criteria,omitzero"`
	Provider map[string]int64 `json:"provider,omitzero"`
	paramObj
}

func (r SendMessageParamsMessageTimeout) MarshalJSON() (data []byte, err error) {
	type shadow SendMessageParamsMessageTimeout
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendMessageParamsMessageTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SendMessageParamsMessageTimeout](
		"criteria", "no-escalation", "delivered", "viewed", "engaged",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SendMessageParamsMessageToUnion struct {
	OfUserRecipient  *UserRecipientParam `json:",omitzero,inline"`
	OfRecipientArray []RecipientParam    `json:",omitzero,inline"`
	paramUnion
}

func (u SendMessageParamsMessageToUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfUserRecipient, u.OfRecipientArray)
}
func (u *SendMessageParamsMessageToUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SendMessageParamsMessageToUnion) asAny() any {
	if !param.IsOmitted(u.OfUserRecipient) {
		return u.OfUserRecipient
	} else if !param.IsOmitted(u.OfRecipientArray) {
		return &u.OfRecipientArray
	}
	return nil
}
