// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"encoding/json"

	"github.com/trycourier/courier-go"
	"github.com/trycourier/courier-go/internal/apijson"
	"github.com/trycourier/courier-go/packages/param"
	"github.com/trycourier/courier-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Alignment string

const (
	AlignmentCenter Alignment = "center"
	AlignmentLeft   Alignment = "left"
	AlignmentRight  Alignment = "right"
	AlignmentFull   Alignment = "full"
)

type ChannelPreference struct {
	// Any of "direct_message", "email", "push", "sms", "webhook", "inbox".
	Channel courier.ChannelClassification `json:"channel,required"`
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
	Channel courier.ChannelClassification `json:"channel,omitzero,required"`
	paramObj
}

func (r ChannelPreferenceParam) MarshalJSON() (data []byte, err error) {
	type shadow ChannelPreferenceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChannelPreferenceParam) UnmarshalJSON(data []byte) error {
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
	Status             courier.PreferenceStatus `json:"status,required"`
	ChannelPreferences []ChannelPreference      `json:"channel_preferences,nullable"`
	Rules              []Rule                   `json:"rules,nullable"`
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
	Status             courier.PreferenceStatus `json:"status,omitzero,required"`
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
