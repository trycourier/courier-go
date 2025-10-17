// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"encoding/json"

	"github.com/trycourier/courier-go/v3/internal/apijson"
	"github.com/trycourier/courier-go/v3/packages/param"
	"github.com/trycourier/courier-go/v3/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type ChannelClassification string

const (
	ChannelClassificationDirectMessage ChannelClassification = "direct_message"
	ChannelClassificationEmail         ChannelClassification = "email"
	ChannelClassificationPush          ChannelClassification = "push"
	ChannelClassificationSMS           ChannelClassification = "sms"
	ChannelClassificationWebhook       ChannelClassification = "webhook"
	ChannelClassificationInbox         ChannelClassification = "inbox"
)

type ChannelPreference struct {
	// Any of "direct_message", "email", "push", "sms", "webhook", "inbox".
	Channel ChannelClassification `json:"channel,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Channel     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChannelPreference) RawJSON() string { return r.JSON.raw }
func (r *ChannelPreference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ChannelPreference to a ChannelPreferenceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ChannelPreferenceParam.Overrides()
func (r ChannelPreference) ToParam() ChannelPreferenceParam {
	return param.Override[ChannelPreferenceParam](json.RawMessage(r.RawJSON()))
}

// The property Channel is required.
type ChannelPreferenceParam struct {
	// Any of "direct_message", "email", "push", "sms", "webhook", "inbox".
	Channel ChannelClassification `json:"channel,omitzero,required"`
	paramObj
}

func (r ChannelPreferenceParam) MarshalJSON() (data []byte, err error) {
	type shadow ChannelPreferenceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChannelPreferenceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ElementalActionNodeWithType struct {
	// Any of "action".
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
func (r ElementalActionNodeWithType) RawJSON() string { return r.JSON.raw }
func (r *ElementalActionNodeWithType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ElementalActionNodeWithType to a
// ElementalActionNodeWithTypeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ElementalActionNodeWithTypeParam.Overrides()
func (r ElementalActionNodeWithType) ToParam() ElementalActionNodeWithTypeParam {
	return param.Override[ElementalActionNodeWithTypeParam](json.RawMessage(r.RawJSON()))
}

type ElementalActionNodeWithTypeParam struct {
	Type string `json:"type,omitzero"`
	ElementalBaseNodeParam
}

func (r ElementalActionNodeWithTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow ElementalActionNodeWithTypeParam
	return param.MarshalObject(r, (*shadow)(&r))
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
type ElementalChannelNodeWithType struct {
	// Any of "channel".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ElementalChannelNode
}

// Returns the unmodified JSON received from the API
func (r ElementalChannelNodeWithType) RawJSON() string { return r.JSON.raw }
func (r *ElementalChannelNodeWithType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ElementalChannelNodeWithType to a
// ElementalChannelNodeWithTypeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ElementalChannelNodeWithTypeParam.Overrides()
func (r ElementalChannelNodeWithType) ToParam() ElementalChannelNodeWithTypeParam {
	return param.Override[ElementalChannelNodeWithTypeParam](json.RawMessage(r.RawJSON()))
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
type ElementalChannelNodeWithTypeParam struct {
	Type string `json:"type,omitzero"`
	ElementalChannelNodeParam
}

func (r ElementalChannelNodeWithTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow ElementalChannelNodeWithTypeParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type ElementalContent struct {
	Elements []ElementalNodeUnion `json:"elements,required"`
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

// The properties Elements, Version are required.
type ElementalContentParam struct {
	Elements []ElementalNodeUnionParam `json:"elements,omitzero,required"`
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

// Syntactic sugar to provide a fast shorthand for Courier Elemental Blocks.
type ElementalContentSugar struct {
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
func (r ElementalContentSugar) RawJSON() string { return r.JSON.raw }
func (r *ElementalContentSugar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ElementalContentSugar to a ElementalContentSugarParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ElementalContentSugarParam.Overrides()
func (r ElementalContentSugar) ToParam() ElementalContentSugarParam {
	return param.Override[ElementalContentSugarParam](json.RawMessage(r.RawJSON()))
}

// Syntactic sugar to provide a fast shorthand for Courier Elemental Blocks.
//
// The properties Body, Title are required.
type ElementalContentSugarParam struct {
	// The text content displayed in the notification.
	Body string `json:"body,required"`
	// Title/subject displayed by supported channels.
	Title string `json:"title,required"`
	paramObj
}

func (r ElementalContentSugarParam) MarshalJSON() (data []byte, err error) {
	type shadow ElementalContentSugarParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ElementalContentSugarParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ElementalDividerNodeWithType struct {
	// Any of "divider".
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
func (r ElementalDividerNodeWithType) RawJSON() string { return r.JSON.raw }
func (r *ElementalDividerNodeWithType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ElementalDividerNodeWithType to a
// ElementalDividerNodeWithTypeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ElementalDividerNodeWithTypeParam.Overrides()
func (r ElementalDividerNodeWithType) ToParam() ElementalDividerNodeWithTypeParam {
	return param.Override[ElementalDividerNodeWithTypeParam](json.RawMessage(r.RawJSON()))
}

type ElementalDividerNodeWithTypeParam struct {
	Type string `json:"type,omitzero"`
	ElementalBaseNodeParam
}

func (r ElementalDividerNodeWithTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow ElementalDividerNodeWithTypeParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type ElementalImageNodeWithType struct {
	// Any of "image".
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
func (r ElementalImageNodeWithType) RawJSON() string { return r.JSON.raw }
func (r *ElementalImageNodeWithType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ElementalImageNodeWithType to a
// ElementalImageNodeWithTypeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ElementalImageNodeWithTypeParam.Overrides()
func (r ElementalImageNodeWithType) ToParam() ElementalImageNodeWithTypeParam {
	return param.Override[ElementalImageNodeWithTypeParam](json.RawMessage(r.RawJSON()))
}

type ElementalImageNodeWithTypeParam struct {
	Type string `json:"type,omitzero"`
	ElementalBaseNodeParam
}

func (r ElementalImageNodeWithTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow ElementalImageNodeWithTypeParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type ElementalMetaNodeWithType struct {
	// Any of "meta".
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
func (r ElementalMetaNodeWithType) RawJSON() string { return r.JSON.raw }
func (r *ElementalMetaNodeWithType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ElementalMetaNodeWithType to a
// ElementalMetaNodeWithTypeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ElementalMetaNodeWithTypeParam.Overrides()
func (r ElementalMetaNodeWithType) ToParam() ElementalMetaNodeWithTypeParam {
	return param.Override[ElementalMetaNodeWithTypeParam](json.RawMessage(r.RawJSON()))
}

type ElementalMetaNodeWithTypeParam struct {
	Type string `json:"type,omitzero"`
	ElementalBaseNodeParam
}

func (r ElementalMetaNodeWithTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow ElementalMetaNodeWithTypeParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// ElementalNodeUnion contains all possible properties and values from
// [ElementalTextNodeWithType], [ElementalMetaNodeWithType],
// [ElementalChannelNodeWithType], [ElementalImageNodeWithType],
// [ElementalActionNodeWithType], [ElementalDividerNodeWithType],
// [ElementalQuoteNodeWithType].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ElementalNodeUnion struct {
	// This field is from variant [ElementalTextNodeWithType],
	// [ElementalMetaNodeWithType], [ElementalChannelNodeWithType],
	// [ElementalImageNodeWithType], [ElementalActionNodeWithType],
	// [ElementalDividerNodeWithType], [ElementalQuoteNodeWithType].
	Channels []string `json:"channels"`
	// This field is from variant [ElementalTextNodeWithType],
	// [ElementalMetaNodeWithType], [ElementalChannelNodeWithType],
	// [ElementalImageNodeWithType], [ElementalActionNodeWithType],
	// [ElementalDividerNodeWithType], [ElementalQuoteNodeWithType].
	If string `json:"if"`
	// This field is from variant [ElementalTextNodeWithType],
	// [ElementalMetaNodeWithType], [ElementalChannelNodeWithType],
	// [ElementalImageNodeWithType], [ElementalActionNodeWithType],
	// [ElementalDividerNodeWithType], [ElementalQuoteNodeWithType].
	Loop string `json:"loop"`
	// This field is from variant [ElementalTextNodeWithType],
	// [ElementalMetaNodeWithType], [ElementalChannelNodeWithType],
	// [ElementalImageNodeWithType], [ElementalActionNodeWithType],
	// [ElementalDividerNodeWithType], [ElementalQuoteNodeWithType].
	Ref  string `json:"ref"`
	Type string `json:"type"`
	// This field is from variant [ElementalChannelNodeWithType].
	Channel string `json:"channel"`
	// This field is from variant [ElementalChannelNodeWithType].
	Raw  map[string]any `json:"raw"`
	JSON struct {
		Channels respjson.Field
		If       respjson.Field
		Loop     respjson.Field
		Ref      respjson.Field
		Type     respjson.Field
		Channel  respjson.Field
		Raw      respjson.Field
		raw      string
	} `json:"-"`
}

func (u ElementalNodeUnion) AsElementalTextNodeWithType() (v ElementalTextNodeWithType) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ElementalNodeUnion) AsElementalMetaNodeWithType() (v ElementalMetaNodeWithType) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ElementalNodeUnion) AsElementalChannelNodeWithType() (v ElementalChannelNodeWithType) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ElementalNodeUnion) AsElementalImageNodeWithType() (v ElementalImageNodeWithType) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ElementalNodeUnion) AsElementalActionNodeWithType() (v ElementalActionNodeWithType) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ElementalNodeUnion) AsElementalDividerNodeWithType() (v ElementalDividerNodeWithType) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ElementalNodeUnion) AsElementalQuoteNodeWithType() (v ElementalQuoteNodeWithType) {
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

func ElementalNodeParamOfElementalChannelNodeWithType(channel string) ElementalNodeUnionParam {
	var variant ElementalChannelNodeWithTypeParam
	variant.Channel = channel
	return ElementalNodeUnionParam{OfElementalChannelNodeWithType: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ElementalNodeUnionParam struct {
	OfElementalTextNodeWithType    *ElementalTextNodeWithTypeParam    `json:",omitzero,inline"`
	OfElementalMetaNodeWithType    *ElementalMetaNodeWithTypeParam    `json:",omitzero,inline"`
	OfElementalChannelNodeWithType *ElementalChannelNodeWithTypeParam `json:",omitzero,inline"`
	OfElementalImageNodeWithType   *ElementalImageNodeWithTypeParam   `json:",omitzero,inline"`
	OfElementalActionNodeWithType  *ElementalActionNodeWithTypeParam  `json:",omitzero,inline"`
	OfElementalDividerNodeWithType *ElementalDividerNodeWithTypeParam `json:",omitzero,inline"`
	OfElementalQuoteNodeWithType   *ElementalQuoteNodeWithTypeParam   `json:",omitzero,inline"`
	paramUnion
}

func (u ElementalNodeUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfElementalTextNodeWithType,
		u.OfElementalMetaNodeWithType,
		u.OfElementalChannelNodeWithType,
		u.OfElementalImageNodeWithType,
		u.OfElementalActionNodeWithType,
		u.OfElementalDividerNodeWithType,
		u.OfElementalQuoteNodeWithType)
}
func (u *ElementalNodeUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ElementalNodeUnionParam) asAny() any {
	if !param.IsOmitted(u.OfElementalTextNodeWithType) {
		return u.OfElementalTextNodeWithType
	} else if !param.IsOmitted(u.OfElementalMetaNodeWithType) {
		return u.OfElementalMetaNodeWithType
	} else if !param.IsOmitted(u.OfElementalChannelNodeWithType) {
		return u.OfElementalChannelNodeWithType
	} else if !param.IsOmitted(u.OfElementalImageNodeWithType) {
		return u.OfElementalImageNodeWithType
	} else if !param.IsOmitted(u.OfElementalActionNodeWithType) {
		return u.OfElementalActionNodeWithType
	} else if !param.IsOmitted(u.OfElementalDividerNodeWithType) {
		return u.OfElementalDividerNodeWithType
	} else if !param.IsOmitted(u.OfElementalQuoteNodeWithType) {
		return u.OfElementalQuoteNodeWithType
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetChannel() *string {
	if vt := u.OfElementalChannelNodeWithType; vt != nil {
		return &vt.Channel
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetRaw() map[string]any {
	if vt := u.OfElementalChannelNodeWithType; vt != nil {
		return vt.Raw
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetIf() *string {
	if vt := u.OfElementalTextNodeWithType; vt != nil && vt.If.Valid() {
		return &vt.If.Value
	} else if vt := u.OfElementalMetaNodeWithType; vt != nil && vt.If.Valid() {
		return &vt.If.Value
	} else if vt := u.OfElementalChannelNodeWithType; vt != nil && vt.If.Valid() {
		return &vt.If.Value
	} else if vt := u.OfElementalImageNodeWithType; vt != nil && vt.If.Valid() {
		return &vt.If.Value
	} else if vt := u.OfElementalActionNodeWithType; vt != nil && vt.If.Valid() {
		return &vt.If.Value
	} else if vt := u.OfElementalDividerNodeWithType; vt != nil && vt.If.Valid() {
		return &vt.If.Value
	} else if vt := u.OfElementalQuoteNodeWithType; vt != nil && vt.If.Valid() {
		return &vt.If.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetLoop() *string {
	if vt := u.OfElementalTextNodeWithType; vt != nil && vt.Loop.Valid() {
		return &vt.Loop.Value
	} else if vt := u.OfElementalMetaNodeWithType; vt != nil && vt.Loop.Valid() {
		return &vt.Loop.Value
	} else if vt := u.OfElementalChannelNodeWithType; vt != nil && vt.Loop.Valid() {
		return &vt.Loop.Value
	} else if vt := u.OfElementalImageNodeWithType; vt != nil && vt.Loop.Valid() {
		return &vt.Loop.Value
	} else if vt := u.OfElementalActionNodeWithType; vt != nil && vt.Loop.Valid() {
		return &vt.Loop.Value
	} else if vt := u.OfElementalDividerNodeWithType; vt != nil && vt.Loop.Valid() {
		return &vt.Loop.Value
	} else if vt := u.OfElementalQuoteNodeWithType; vt != nil && vt.Loop.Valid() {
		return &vt.Loop.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetRef() *string {
	if vt := u.OfElementalTextNodeWithType; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	} else if vt := u.OfElementalMetaNodeWithType; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	} else if vt := u.OfElementalChannelNodeWithType; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	} else if vt := u.OfElementalImageNodeWithType; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	} else if vt := u.OfElementalActionNodeWithType; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	} else if vt := u.OfElementalDividerNodeWithType; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	} else if vt := u.OfElementalQuoteNodeWithType; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetType() *string {
	if vt := u.OfElementalTextNodeWithType; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfElementalMetaNodeWithType; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfElementalChannelNodeWithType; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfElementalImageNodeWithType; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfElementalActionNodeWithType; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfElementalDividerNodeWithType; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfElementalQuoteNodeWithType; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's Channels property, if present.
func (u ElementalNodeUnionParam) GetChannels() []string {
	if vt := u.OfElementalTextNodeWithType; vt != nil {
		return vt.Channels
	} else if vt := u.OfElementalMetaNodeWithType; vt != nil {
		return vt.Channels
	} else if vt := u.OfElementalChannelNodeWithType; vt != nil {
		return vt.Channels
	} else if vt := u.OfElementalImageNodeWithType; vt != nil {
		return vt.Channels
	} else if vt := u.OfElementalActionNodeWithType; vt != nil {
		return vt.Channels
	} else if vt := u.OfElementalDividerNodeWithType; vt != nil {
		return vt.Channels
	} else if vt := u.OfElementalQuoteNodeWithType; vt != nil {
		return vt.Channels
	}
	return nil
}

type ElementalQuoteNodeWithType struct {
	// Any of "quote".
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
func (r ElementalQuoteNodeWithType) RawJSON() string { return r.JSON.raw }
func (r *ElementalQuoteNodeWithType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ElementalQuoteNodeWithType to a
// ElementalQuoteNodeWithTypeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ElementalQuoteNodeWithTypeParam.Overrides()
func (r ElementalQuoteNodeWithType) ToParam() ElementalQuoteNodeWithTypeParam {
	return param.Override[ElementalQuoteNodeWithTypeParam](json.RawMessage(r.RawJSON()))
}

type ElementalQuoteNodeWithTypeParam struct {
	Type string `json:"type,omitzero"`
	ElementalBaseNodeParam
}

func (r ElementalQuoteNodeWithTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow ElementalQuoteNodeWithTypeParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type ElementalTextNodeWithType struct {
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
func (r ElementalTextNodeWithType) RawJSON() string { return r.JSON.raw }
func (r *ElementalTextNodeWithType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ElementalTextNodeWithType to a
// ElementalTextNodeWithTypeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ElementalTextNodeWithTypeParam.Overrides()
func (r ElementalTextNodeWithType) ToParam() ElementalTextNodeWithTypeParam {
	return param.Override[ElementalTextNodeWithTypeParam](json.RawMessage(r.RawJSON()))
}

type ElementalTextNodeWithTypeParam struct {
	Type string `json:"type,omitzero"`
	ElementalBaseNodeParam
}

func (r ElementalTextNodeWithTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow ElementalTextNodeWithTypeParam
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

type NotificationPreferenceDetails struct {
	// Any of "OPTED_IN", "OPTED_OUT", "REQUIRED".
	Status             PreferenceStatus    `json:"status,required"`
	ChannelPreferences []ChannelPreference `json:"channel_preferences,nullable"`
	Rules              []Rule              `json:"rules,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status             respjson.Field
		ChannelPreferences respjson.Field
		Rules              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationPreferenceDetails) RawJSON() string { return r.JSON.raw }
func (r *NotificationPreferenceDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this NotificationPreferenceDetails to a
// NotificationPreferenceDetailsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// NotificationPreferenceDetailsParam.Overrides()
func (r NotificationPreferenceDetails) ToParam() NotificationPreferenceDetailsParam {
	return param.Override[NotificationPreferenceDetailsParam](json.RawMessage(r.RawJSON()))
}

// The property Status is required.
type NotificationPreferenceDetailsParam struct {
	// Any of "OPTED_IN", "OPTED_OUT", "REQUIRED".
	Status             PreferenceStatus         `json:"status,omitzero,required"`
	ChannelPreferences []ChannelPreferenceParam `json:"channel_preferences,omitzero"`
	Rules              []RuleParam              `json:"rules,omitzero"`
	paramObj
}

func (r NotificationPreferenceDetailsParam) MarshalJSON() (data []byte, err error) {
	type shadow NotificationPreferenceDetailsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NotificationPreferenceDetailsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Paging struct {
	More   bool   `json:"more,required"`
	Cursor string `json:"cursor,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		More        respjson.Field
		Cursor      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Paging) RawJSON() string { return r.JSON.raw }
func (r *Paging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Preference struct {
	// Any of "OPTED_IN", "OPTED_OUT", "REQUIRED".
	Status             PreferenceStatus    `json:"status,required"`
	ChannelPreferences []ChannelPreference `json:"channel_preferences,nullable"`
	Rules              []Rule              `json:"rules,nullable"`
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
	Status             PreferenceStatus         `json:"status,omitzero,required"`
	ChannelPreferences []ChannelPreferenceParam `json:"channel_preferences,omitzero"`
	Rules              []RuleParam              `json:"rules,omitzero"`
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

type PreferenceStatus string

const (
	PreferenceStatusOptedIn  PreferenceStatus = "OPTED_IN"
	PreferenceStatusOptedOut PreferenceStatus = "OPTED_OUT"
	PreferenceStatusRequired PreferenceStatus = "REQUIRED"
)

type RecipientParam struct {
	// Deprecated - Use `tenant_id` instead.
	AccountID param.Opt[string] `json:"account_id,omitzero"`
	// The user's email address.
	Email param.Opt[string] `json:"email,omitzero"`
	// The user's preferred ISO 639-1 language code.
	Locale param.Opt[string] `json:"locale,omitzero"`
	// The user's phone number.
	PhoneNumber param.Opt[string] `json:"phone_number,omitzero"`
	// The id of the tenant the user is associated with.
	TenantID param.Opt[string] `json:"tenant_id,omitzero"`
	// The user's unique identifier. Typically, this will match the user id of a user
	// in your system.
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

type RecipientPreferences struct {
	Categories    map[string]NotificationPreferenceDetails `json:"categories,nullable"`
	Notifications map[string]NotificationPreferenceDetails `json:"notifications,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Categories    respjson.Field
		Notifications respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecipientPreferences) RawJSON() string { return r.JSON.raw }
func (r *RecipientPreferences) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this RecipientPreferences to a RecipientPreferencesParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RecipientPreferencesParam.Overrides()
func (r RecipientPreferences) ToParam() RecipientPreferencesParam {
	return param.Override[RecipientPreferencesParam](json.RawMessage(r.RawJSON()))
}

type Rule struct {
	Until string `json:"until,required"`
	Start string `json:"start,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Until       respjson.Field
		Start       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Rule) RawJSON() string { return r.JSON.raw }
func (r *Rule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Rule to a RuleParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RuleParam.Overrides()
func (r Rule) ToParam() RuleParam {
	return param.Override[RuleParam](json.RawMessage(r.RawJSON()))
}

// The property Until is required.
type RuleParam struct {
	Until string            `json:"until,required"`
	Start param.Opt[string] `json:"start,omitzero"`
	paramObj
}

func (r RuleParam) MarshalJSON() (data []byte, err error) {
	type shadow RuleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserRecipient struct {
	// Deprecated - Use `tenant_id` instead.
	AccountID string `json:"account_id,nullable"`
	// Context such as tenant_id to send the notification with.
	Context MessageContext `json:"context,nullable"`
	Data    map[string]any `json:"data,nullable"`
	// The user's email address.
	Email string `json:"email,nullable"`
	// The user's preferred ISO 639-1 language code.
	Locale string `json:"locale,nullable"`
	// The user's phone number.
	PhoneNumber string                   `json:"phone_number,nullable"`
	Preferences UserRecipientPreferences `json:"preferences,nullable"`
	// The id of the tenant the user is associated with.
	TenantID string `json:"tenant_id,nullable"`
	// The user's unique identifier. Typically, this will match the user id of a user
	// in your system.
	UserID string `json:"user_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountID   respjson.Field
		Context     respjson.Field
		Data        respjson.Field
		Email       respjson.Field
		Locale      respjson.Field
		PhoneNumber respjson.Field
		Preferences respjson.Field
		TenantID    respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserRecipient) RawJSON() string { return r.JSON.raw }
func (r *UserRecipient) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this UserRecipient to a UserRecipientParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// UserRecipientParam.Overrides()
func (r UserRecipient) ToParam() UserRecipientParam {
	return param.Override[UserRecipientParam](json.RawMessage(r.RawJSON()))
}

type UserRecipientPreferences struct {
	Notifications map[string]Preference `json:"notifications,required"`
	Categories    map[string]Preference `json:"categories,nullable"`
	TemplateID    string                `json:"templateId,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Notifications respjson.Field
		Categories    respjson.Field
		TemplateID    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserRecipientPreferences) RawJSON() string { return r.JSON.raw }
func (r *UserRecipientPreferences) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserRecipientParam struct {
	// Deprecated - Use `tenant_id` instead.
	AccountID param.Opt[string] `json:"account_id,omitzero"`
	// The user's email address.
	Email param.Opt[string] `json:"email,omitzero"`
	// The user's preferred ISO 639-1 language code.
	Locale param.Opt[string] `json:"locale,omitzero"`
	// The user's phone number.
	PhoneNumber param.Opt[string] `json:"phone_number,omitzero"`
	// The id of the tenant the user is associated with.
	TenantID param.Opt[string] `json:"tenant_id,omitzero"`
	// The user's unique identifier. Typically, this will match the user id of a user
	// in your system.
	UserID      param.Opt[string]             `json:"user_id,omitzero"`
	Data        map[string]any                `json:"data,omitzero"`
	Preferences UserRecipientPreferencesParam `json:"preferences,omitzero"`
	// Context such as tenant_id to send the notification with.
	Context MessageContextParam `json:"context,omitzero"`
	paramObj
}

func (r UserRecipientParam) MarshalJSON() (data []byte, err error) {
	type shadow UserRecipientParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserRecipientParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Notifications is required.
type UserRecipientPreferencesParam struct {
	Notifications map[string]PreferenceParam `json:"notifications,omitzero,required"`
	TemplateID    param.Opt[string]          `json:"templateId,omitzero"`
	Categories    map[string]PreferenceParam `json:"categories,omitzero"`
	paramObj
}

func (r UserRecipientPreferencesParam) MarshalJSON() (data []byte, err error) {
	type shadow UserRecipientPreferencesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserRecipientPreferencesParam) UnmarshalJSON(data []byte) error {
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
