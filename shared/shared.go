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
