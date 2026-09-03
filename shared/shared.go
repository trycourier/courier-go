// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"encoding/json"

	"github.com/trycourier/courier-go/v4/internal/apijson"
	"github.com/trycourier/courier-go/v4/packages/param"
	"github.com/trycourier/courier-go/v4/packages/respjson"
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

// The properties Operator, Path, Value are required.
type AudienceFilterParam struct {
	// Send to users only if they are member of the account
	//
	// Any of "MEMBER_OF".
	Operator AudienceFilterOperator `json:"operator,omitzero" api:"required"`
	// Any of "account_id".
	Path  AudienceFilterPath `json:"path,omitzero" api:"required"`
	Value string             `json:"value" api:"required"`
	paramObj
}

func (r AudienceFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow AudienceFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AudienceFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Send to users only if they are member of the account
type AudienceFilterOperator string

const (
	AudienceFilterOperatorMemberOf AudienceFilterOperator = "MEMBER_OF"
)

type AudienceFilterPath string

const (
	AudienceFilterPathAccountID AudienceFilterPath = "account_id"
)

// Filter configuration for audience membership containing an array of filter rules
type AudienceFilterConfig struct {
	// Array of filter rules (single conditions or nested groups)
	Filters []FilterConfig `json:"filters" api:"required"`
	// The logical operator (AND/OR) combining the rules in `filters`. Required when
	// `filters` contains more than one rule. If omitted, the top-level `operator`
	// field on the request is used instead.
	//
	// Any of "AND", "OR".
	Operator AudienceFilterConfigOperator `json:"operator"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filters     respjson.Field
		Operator    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudienceFilterConfig) RawJSON() string { return r.JSON.raw }
func (r *AudienceFilterConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AudienceFilterConfig to a AudienceFilterConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AudienceFilterConfigParam.Overrides()
func (r AudienceFilterConfig) ToParam() AudienceFilterConfigParam {
	return param.Override[AudienceFilterConfigParam](json.RawMessage(r.RawJSON()))
}

// The logical operator (AND/OR) combining the rules in `filters`. Required when
// `filters` contains more than one rule. If omitted, the top-level `operator`
// field on the request is used instead.
type AudienceFilterConfigOperator string

const (
	AudienceFilterConfigOperatorAnd AudienceFilterConfigOperator = "AND"
	AudienceFilterConfigOperatorOr  AudienceFilterConfigOperator = "OR"
)

// Filter configuration for audience membership containing an array of filter rules
//
// The property Filters is required.
type AudienceFilterConfigParam struct {
	// Array of filter rules (single conditions or nested groups)
	Filters []FilterConfigParam `json:"filters,omitzero" api:"required"`
	// The logical operator (AND/OR) combining the rules in `filters`. Required when
	// `filters` contains more than one rule. If omitted, the top-level `operator`
	// field on the request is used instead.
	//
	// Any of "AND", "OR".
	Operator AudienceFilterConfigOperator `json:"operator,omitzero"`
	paramObj
}

func (r AudienceFilterConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow AudienceFilterConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AudienceFilterConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Send to all users in an audience
//
// The property AudienceID is required.
type AudienceRecipientParam struct {
	// A unique identifier associated with an Audience. A message will be sent to each
	// user in the audience.
	AudienceID string                `json:"audience_id" api:"required"`
	Data       map[string]any        `json:"data,omitzero"`
	Filters    []AudienceFilterParam `json:"filters,omitzero"`
	paramObj
}

func (r AudienceRecipientParam) MarshalJSON() (data []byte, err error) {
	type shadow AudienceRecipientParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AudienceRecipientParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Channel struct {
	// Brand id used for rendering.
	BrandID string `json:"brand_id" api:"nullable"`
	// JS conditional with access to data/profile.
	If       string          `json:"if" api:"nullable"`
	Metadata ChannelMetadata `json:"metadata" api:"nullable"`
	// Channel specific overrides.
	Override map[string]any `json:"override" api:"nullable"`
	// Providers enabled for this channel.
	Providers []string `json:"providers" api:"nullable"`
	// Defaults to `single`.
	//
	// Any of "all", "single".
	RoutingMethod ChannelRoutingMethod `json:"routing_method" api:"nullable"`
	Timeouts      Timeouts             `json:"timeouts" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BrandID       respjson.Field
		If            respjson.Field
		Metadata      respjson.Field
		Override      respjson.Field
		Providers     respjson.Field
		RoutingMethod respjson.Field
		Timeouts      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Channel) RawJSON() string { return r.JSON.raw }
func (r *Channel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Channel to a ChannelParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ChannelParam.Overrides()
func (r Channel) ToParam() ChannelParam {
	return param.Override[ChannelParam](json.RawMessage(r.RawJSON()))
}

// Defaults to `single`.
type ChannelRoutingMethod string

const (
	ChannelRoutingMethodAll    ChannelRoutingMethod = "all"
	ChannelRoutingMethodSingle ChannelRoutingMethod = "single"
)

type ChannelParam struct {
	// Brand id used for rendering.
	BrandID param.Opt[string] `json:"brand_id,omitzero"`
	// JS conditional with access to data/profile.
	If param.Opt[string] `json:"if,omitzero"`
	// Channel specific overrides.
	Override map[string]any `json:"override,omitzero"`
	// Providers enabled for this channel.
	Providers []string `json:"providers,omitzero"`
	// Defaults to `single`.
	//
	// Any of "all", "single".
	RoutingMethod ChannelRoutingMethod `json:"routing_method,omitzero"`
	Metadata      ChannelMetadataParam `json:"metadata,omitzero"`
	Timeouts      TimeoutsParam        `json:"timeouts,omitzero"`
	paramObj
}

func (r ChannelParam) MarshalJSON() (data []byte, err error) {
	type shadow ChannelParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChannelParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChannelClassification string

const (
	ChannelClassificationDirectMessage ChannelClassification = "direct_message"
	ChannelClassificationEmail         ChannelClassification = "email"
	ChannelClassificationPush          ChannelClassification = "push"
	ChannelClassificationSMS           ChannelClassification = "sms"
	ChannelClassificationWebhook       ChannelClassification = "webhook"
	ChannelClassificationInbox         ChannelClassification = "inbox"
)

type ChannelMetadata struct {
	Utm Utm `json:"utm" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Utm         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChannelMetadata) RawJSON() string { return r.JSON.raw }
func (r *ChannelMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ChannelMetadata to a ChannelMetadataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ChannelMetadataParam.Overrides()
func (r ChannelMetadata) ToParam() ChannelMetadataParam {
	return param.Override[ChannelMetadataParam](json.RawMessage(r.RawJSON()))
}

type ChannelMetadataParam struct {
	Utm UtmParam `json:"utm,omitzero"`
	paramObj
}

func (r ChannelMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow ChannelMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChannelMetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChannelPreference struct {
	// Any of "direct_message", "email", "push", "sms", "webhook", "inbox".
	Channel ChannelClassification `json:"channel" api:"required"`
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
	Channel ChannelClassification `json:"channel,omitzero" api:"required"`
	paramObj
}

func (r ChannelPreferenceParam) MarshalJSON() (data []byte, err error) {
	type shadow ChannelPreferenceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChannelPreferenceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows the user to execute an action. Can be a button or a link.
type ElementalActionNode struct {
	// The text content of the action shown to the user.
	Content string `json:"content" api:"required"`
	// The target URL of the action.
	Href string `json:"href" api:"required"`
	// A unique id used to identify the action when it is executed.
	ActionID string `json:"action_id" api:"nullable"`
	// The alignment of the action button. Defaults to "center".
	//
	// Any of "center", "left", "right", "full".
	Align Alignment `json:"align" api:"nullable"`
	// The background color of the action button.
	BackgroundColor string `json:"background_color" api:"nullable"`
	// CSS border-radius applied to the action button. For example, `4px`
	BorderRadius string `json:"border_radius" api:"nullable"`
	// CSS border width applied to the action button. For example, `1px`
	BorderSize string `json:"border_size" api:"nullable"`
	// When true, the action's href is not rewritten for click-through tracking, even
	// when click-through tracking is enabled for the workspace.
	DisableTracking bool `json:"disable_tracking" api:"nullable"`
	// CSS font-size applied to the action button label. For example, `14px`
	FontSize string `json:"font_size" api:"nullable"`
	// Region specific content. See
	// [locales docs](https://www.courier.com/docs/platform/content/elemental/locales/)
	// for more details.
	Locales Locales `json:"locales" api:"nullable"`
	// CSS padding applied to the action button. For example, `8px 16px`
	Padding string `json:"padding" api:"nullable"`
	// Defaults to `button`.
	//
	// Any of "button", "link".
	Style string `json:"style" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content         respjson.Field
		Href            respjson.Field
		ActionID        respjson.Field
		Align           respjson.Field
		BackgroundColor respjson.Field
		BorderRadius    respjson.Field
		BorderSize      respjson.Field
		DisableTracking respjson.Field
		FontSize        respjson.Field
		Locales         respjson.Field
		Padding         respjson.Field
		Style           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
	ElementalBaseNode
}

// Returns the unmodified JSON received from the API
func (r ElementalActionNode) RawJSON() string { return r.JSON.raw }
func (r *ElementalActionNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ElementalActionNode to a ElementalActionNodeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ElementalActionNodeParam.Overrides()
func (r ElementalActionNode) ToParam() ElementalActionNodeParam {
	return param.Override[ElementalActionNodeParam](json.RawMessage(r.RawJSON()))
}

// Allows the user to execute an action. Can be a button or a link.
type ElementalActionNodeParam struct {
	// The text content of the action shown to the user.
	Content string `json:"content" api:"required"`
	// The target URL of the action.
	Href string `json:"href" api:"required"`
	// A unique id used to identify the action when it is executed.
	ActionID param.Opt[string] `json:"action_id,omitzero"`
	// The alignment of the action button. Defaults to "center".
	Align Alignment `json:"align,omitzero"`
	// The background color of the action button.
	BackgroundColor param.Opt[string] `json:"background_color,omitzero"`
	// CSS border-radius applied to the action button. For example, `4px`
	BorderRadius param.Opt[string] `json:"border_radius,omitzero"`
	// CSS border width applied to the action button. For example, `1px`
	BorderSize param.Opt[string] `json:"border_size,omitzero"`
	// When true, the action's href is not rewritten for click-through tracking, even
	// when click-through tracking is enabled for the workspace.
	DisableTracking param.Opt[bool] `json:"disable_tracking,omitzero"`
	// CSS font-size applied to the action button label. For example, `14px`
	FontSize param.Opt[string] `json:"font_size,omitzero"`
	// Region specific content. See
	// [locales docs](https://www.courier.com/docs/platform/content/elemental/locales/)
	// for more details.
	Locales LocalesParam `json:"locales,omitzero"`
	// CSS padding applied to the action button. For example, `8px 16px`
	Padding param.Opt[string] `json:"padding,omitzero"`
	// Defaults to `button`.
	Style string `json:"style,omitzero"`
	ElementalBaseNodeParam
}

func (r ElementalActionNodeParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*ElementalActionNodeParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// Allows the user to execute an action. Can be a button or a link.
type ElementalActionNodeWithType struct {
	// Any of "action".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ElementalActionNode
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

// Allows the user to execute an action. Can be a button or a link.
type ElementalActionNodeWithTypeParam struct {
	Type string `json:"type,omitzero"`
	ElementalActionNodeParam
}

func (r ElementalActionNodeWithTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*ElementalActionNodeWithTypeParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

type ElementalBaseNode struct {
	Channels []string `json:"channels" api:"nullable"`
	If       string   `json:"if" api:"nullable"`
	Loop     string   `json:"loop" api:"nullable"`
	Ref      string   `json:"ref" api:"nullable"`
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
	Channel string `json:"channel"`
	// An array of elements to apply to the channel. If `raw` has not been specified,
	// `elements` is `required`. Channel elements cannot nest, so these are any node
	// except another channel block.
	Elements []ElementalNodeNonChannelUnion `json:"elements" api:"nullable"`
	// Email only. Document-level base font size (CSS px, e.g. `16px`) for body content
	// — text, quote, list and action button labels. Heading styles (`h1`/`h2`/`h3`)
	// and `subtext` keep their preset sizes.
	FontSize string `json:"font_size" api:"nullable"`
	// Email only. Document-level line height (CSS px or unitless multiplier, e.g.
	// `24px` or `1.5`) applied to all body content unless overridden per block.
	LineHeight string `json:"line_height" api:"nullable"`
	// Email only. Document-level body padding applied once around the email body, as a
	// CSS px shorthand (1–4 values), e.g. `48px 64px`.
	Padding string `json:"padding" api:"nullable"`
	// Raw data to apply to the channel. If `elements` has not been specified, `raw` is
	// required.
	Raw map[string]any `json:"raw" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Channel     respjson.Field
		Elements    respjson.Field
		FontSize    respjson.Field
		LineHeight  respjson.Field
		Padding     respjson.Field
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
	Channel param.Opt[string] `json:"channel,omitzero"`
	// An array of elements to apply to the channel. If `raw` has not been specified,
	// `elements` is `required`. Channel elements cannot nest, so these are any node
	// except another channel block.
	Elements []ElementalNodeNonChannelUnionParam `json:"elements,omitzero"`
	// Email only. Document-level base font size (CSS px, e.g. `16px`) for body content
	// — text, quote, list and action button labels. Heading styles (`h1`/`h2`/`h3`)
	// and `subtext` keep their preset sizes.
	FontSize param.Opt[string] `json:"font_size,omitzero"`
	// Email only. Document-level line height (CSS px or unitless multiplier, e.g.
	// `24px` or `1.5`) applied to all body content unless overridden per block.
	LineHeight param.Opt[string] `json:"line_height,omitzero"`
	// Email only. Document-level body padding applied once around the email body, as a
	// CSS px shorthand (1–4 values), e.g. `48px 64px`.
	Padding param.Opt[string] `json:"padding,omitzero"`
	// Raw data to apply to the channel. If `elements` has not been specified, `raw` is
	// required.
	Raw map[string]any `json:"raw,omitzero"`
	ElementalBaseNodeParam
}

func (r ElementalChannelNodeParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*ElementalChannelNodeParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
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
	type shadow struct {
		*ElementalChannelNodeWithTypeParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

type ElementalContent struct {
	Elements []ElementalNodeUnion `json:"elements" api:"required"`
	// For example, "2022-01-01"
	Version string `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Elements    respjson.Field
		Version     respjson.Field
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
	Elements []ElementalNodeUnionParam `json:"elements,omitzero" api:"required"`
	// For example, "2022-01-01"
	Version string `json:"version" api:"required"`
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
	Body string `json:"body" api:"required"`
	// Title/subject displayed by supported channels.
	Title string `json:"title" api:"required"`
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
	Body string `json:"body" api:"required"`
	// Title/subject displayed by supported channels.
	Title string `json:"title" api:"required"`
	paramObj
}

func (r ElementalContentSugarParam) MarshalJSON() (data []byte, err error) {
	type shadow ElementalContentSugarParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ElementalContentSugarParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Renders a dividing line between elements.
type ElementalDividerNode struct {
	// The CSS color to render the line with. For example, `#fff`
	Color string `json:"color" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Color       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ElementalBaseNode
}

// Returns the unmodified JSON received from the API
func (r ElementalDividerNode) RawJSON() string { return r.JSON.raw }
func (r *ElementalDividerNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ElementalDividerNode to a ElementalDividerNodeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ElementalDividerNodeParam.Overrides()
func (r ElementalDividerNode) ToParam() ElementalDividerNodeParam {
	return param.Override[ElementalDividerNodeParam](json.RawMessage(r.RawJSON()))
}

// Renders a dividing line between elements.
type ElementalDividerNodeParam struct {
	// The CSS color to render the line with. For example, `#fff`
	Color param.Opt[string] `json:"color,omitzero"`
	ElementalBaseNodeParam
}

func (r ElementalDividerNodeParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*ElementalDividerNodeParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// Renders a dividing line between elements.
type ElementalDividerNodeWithType struct {
	// Any of "divider".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ElementalDividerNode
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

// Renders a dividing line between elements.
type ElementalDividerNodeWithTypeParam struct {
	Type string `json:"type,omitzero"`
	ElementalDividerNodeParam
}

func (r ElementalDividerNodeWithTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*ElementalDividerNodeWithTypeParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// Raw HTML string inside an Elemental document. When rendering a message, this
// node is turned into output only for the email channel; for other channels it
// produces no blocks.
type ElementalHTMLNode struct {
	// Raw HTML string to render inside the notification.
	Content string `json:"content" api:"required"`
	// Region specific content. See
	// [locales docs](https://www.courier.com/docs/platform/content/elemental/locales/)
	// for more details.
	Locales Locales `json:"locales" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Locales     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ElementalBaseNode
}

// Returns the unmodified JSON received from the API
func (r ElementalHTMLNode) RawJSON() string { return r.JSON.raw }
func (r *ElementalHTMLNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ElementalHTMLNode to a ElementalHTMLNodeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ElementalHTMLNodeParam.Overrides()
func (r ElementalHTMLNode) ToParam() ElementalHTMLNodeParam {
	return param.Override[ElementalHTMLNodeParam](json.RawMessage(r.RawJSON()))
}

// Raw HTML string inside an Elemental document. When rendering a message, this
// node is turned into output only for the email channel; for other channels it
// produces no blocks.
type ElementalHTMLNodeParam struct {
	// Raw HTML string to render inside the notification.
	Content string `json:"content" api:"required"`
	// Region specific content. See
	// [locales docs](https://www.courier.com/docs/platform/content/elemental/locales/)
	// for more details.
	Locales LocalesParam `json:"locales,omitzero"`
	ElementalBaseNodeParam
}

func (r ElementalHTMLNodeParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*ElementalHTMLNodeParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// Raw HTML string inside an Elemental document. When rendering a message, this
// node is turned into output only for the email channel; for other channels it
// produces no blocks.
type ElementalHTMLNodeWithType struct {
	// Any of "html".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ElementalHTMLNode
}

// Returns the unmodified JSON received from the API
func (r ElementalHTMLNodeWithType) RawJSON() string { return r.JSON.raw }
func (r *ElementalHTMLNodeWithType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ElementalHTMLNodeWithType to a
// ElementalHTMLNodeWithTypeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ElementalHTMLNodeWithTypeParam.Overrides()
func (r ElementalHTMLNodeWithType) ToParam() ElementalHTMLNodeWithTypeParam {
	return param.Override[ElementalHTMLNodeWithTypeParam](json.RawMessage(r.RawJSON()))
}

// Raw HTML string inside an Elemental document. When rendering a message, this
// node is turned into output only for the email channel; for other channels it
// produces no blocks.
type ElementalHTMLNodeWithTypeParam struct {
	Type string `json:"type,omitzero"`
	ElementalHTMLNodeParam
}

func (r ElementalHTMLNodeWithTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*ElementalHTMLNodeWithTypeParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// Used to embed an image into the notification.
type ElementalImageNode struct {
	// The source of the image.
	Src string `json:"src" api:"required"`
	// The alignment of the image.
	//
	// Any of "center", "left", "right", "full".
	Align Alignment `json:"align" api:"nullable"`
	// Alternate text for the image.
	AltText string `json:"alt_text" api:"nullable"`
	// CSS border color applied to the image. For example, `#ccc`
	BorderColor string `json:"border_color" api:"nullable"`
	// CSS border width applied to the image. For example, `1px`
	BorderSize string `json:"border_size" api:"nullable"`
	// A URL to link to when the image is clicked.
	Href string `json:"href" api:"nullable"`
	// CSS padding applied around the image. For example, `10px`
	Padding string `json:"padding" api:"nullable"`
	// CSS width properties to apply to the image. For example, 50px
	Width string `json:"width" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Src         respjson.Field
		Align       respjson.Field
		AltText     respjson.Field
		BorderColor respjson.Field
		BorderSize  respjson.Field
		Href        respjson.Field
		Padding     respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ElementalBaseNode
}

// Returns the unmodified JSON received from the API
func (r ElementalImageNode) RawJSON() string { return r.JSON.raw }
func (r *ElementalImageNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ElementalImageNode to a ElementalImageNodeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ElementalImageNodeParam.Overrides()
func (r ElementalImageNode) ToParam() ElementalImageNodeParam {
	return param.Override[ElementalImageNodeParam](json.RawMessage(r.RawJSON()))
}

// Used to embed an image into the notification.
type ElementalImageNodeParam struct {
	// The source of the image.
	Src string `json:"src" api:"required"`
	// The alignment of the image.
	Align Alignment `json:"align,omitzero"`
	// Alternate text for the image.
	AltText param.Opt[string] `json:"alt_text,omitzero"`
	// CSS border color applied to the image. For example, `#ccc`
	BorderColor param.Opt[string] `json:"border_color,omitzero"`
	// CSS border width applied to the image. For example, `1px`
	BorderSize param.Opt[string] `json:"border_size,omitzero"`
	// A URL to link to when the image is clicked.
	Href param.Opt[string] `json:"href,omitzero"`
	// CSS padding applied around the image. For example, `10px`
	Padding param.Opt[string] `json:"padding,omitzero"`
	// CSS width properties to apply to the image. For example, 50px
	Width param.Opt[string] `json:"width,omitzero"`
	ElementalBaseNodeParam
}

func (r ElementalImageNodeParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*ElementalImageNodeParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// Used to embed an image into the notification.
type ElementalImageNodeWithType struct {
	// Any of "image".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ElementalImageNode
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

// Used to embed an image into the notification.
type ElementalImageNodeWithTypeParam struct {
	Type string `json:"type,omitzero"`
	ElementalImageNodeParam
}

func (r ElementalImageNodeWithTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*ElementalImageNodeWithTypeParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// The meta element contains information describing the notification that may be
// used by a particular channel or provider. One important field is the title field
// which will be used as the title for channels that support it.
type ElementalMetaNode struct {
	// The title to be displayed by supported channels. For example, the email subject.
	Title string `json:"title" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ElementalBaseNode
}

// Returns the unmodified JSON received from the API
func (r ElementalMetaNode) RawJSON() string { return r.JSON.raw }
func (r *ElementalMetaNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ElementalMetaNode to a ElementalMetaNodeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ElementalMetaNodeParam.Overrides()
func (r ElementalMetaNode) ToParam() ElementalMetaNodeParam {
	return param.Override[ElementalMetaNodeParam](json.RawMessage(r.RawJSON()))
}

// The meta element contains information describing the notification that may be
// used by a particular channel or provider. One important field is the title field
// which will be used as the title for channels that support it.
type ElementalMetaNodeParam struct {
	// The title to be displayed by supported channels. For example, the email subject.
	Title param.Opt[string] `json:"title,omitzero"`
	ElementalBaseNodeParam
}

func (r ElementalMetaNodeParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*ElementalMetaNodeParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// The meta element contains information describing the notification that may be
// used by a particular channel or provider. One important field is the title field
// which will be used as the title for channels that support it.
type ElementalMetaNodeWithType struct {
	// Any of "meta".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ElementalMetaNode
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

// The meta element contains information describing the notification that may be
// used by a particular channel or provider. One important field is the title field
// which will be used as the title for channels that support it.
type ElementalMetaNodeWithTypeParam struct {
	Type string `json:"type,omitzero"`
	ElementalMetaNodeParam
}

func (r ElementalMetaNodeWithTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*ElementalMetaNodeWithTypeParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// ElementalNodeUnion contains all possible properties and values from
// [ElementalTextNodeWithType], [ElementalMetaNodeWithType],
// [ElementalChannelNodeWithType], [ElementalImageNodeWithType],
// [ElementalActionNodeWithType], [ElementalDividerNodeWithType],
// [ElementalQuoteNodeWithType], [ElementalHTMLNodeWithType].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ElementalNodeUnion struct {
	// This field is from variant [ElementalTextNodeWithType],
	// [ElementalMetaNodeWithType], [ElementalChannelNodeWithType],
	// [ElementalImageNodeWithType], [ElementalActionNodeWithType],
	// [ElementalDividerNodeWithType], [ElementalQuoteNodeWithType],
	// [ElementalHTMLNodeWithType].
	Channels []string `json:"channels"`
	// This field is from variant [ElementalTextNodeWithType],
	// [ElementalMetaNodeWithType], [ElementalChannelNodeWithType],
	// [ElementalImageNodeWithType], [ElementalActionNodeWithType],
	// [ElementalDividerNodeWithType], [ElementalQuoteNodeWithType],
	// [ElementalHTMLNodeWithType].
	If string `json:"if"`
	// This field is from variant [ElementalTextNodeWithType],
	// [ElementalMetaNodeWithType], [ElementalChannelNodeWithType],
	// [ElementalImageNodeWithType], [ElementalActionNodeWithType],
	// [ElementalDividerNodeWithType], [ElementalQuoteNodeWithType],
	// [ElementalHTMLNodeWithType].
	Loop string `json:"loop"`
	// This field is from variant [ElementalTextNodeWithType],
	// [ElementalMetaNodeWithType], [ElementalChannelNodeWithType],
	// [ElementalImageNodeWithType], [ElementalActionNodeWithType],
	// [ElementalDividerNodeWithType], [ElementalQuoteNodeWithType],
	// [ElementalHTMLNodeWithType].
	Ref   string `json:"ref"`
	Align string `json:"align"`
	// This field is from variant [ElementalTextNodeWithType].
	Bold     string `json:"bold"`
	Color    string `json:"color"`
	Content  string `json:"content"`
	FontSize string `json:"font_size"`
	// This field is from variant [ElementalTextNodeWithType].
	Format string `json:"format"`
	// This field is from variant [ElementalTextNodeWithType].
	Italic     string `json:"italic"`
	LineHeight string `json:"line_height"`
	// This field is from variant [ElementalTextNodeWithType].
	Locales Locales `json:"locales"`
	// This field is from variant [ElementalTextNodeWithType].
	Strikethrough string `json:"strikethrough"`
	// This field is from variant [ElementalTextNodeWithType].
	TextStyle TextStyle `json:"text_style"`
	// This field is from variant [ElementalTextNodeWithType].
	Underline string `json:"underline"`
	Type      string `json:"type"`
	// This field is from variant [ElementalMetaNodeWithType].
	Title string `json:"title"`
	// This field is from variant [ElementalChannelNodeWithType].
	Channel string `json:"channel"`
	// This field is from variant [ElementalChannelNodeWithType].
	Elements []ElementalNodeNonChannelUnion `json:"elements"`
	Padding  string                         `json:"padding"`
	// This field is from variant [ElementalChannelNodeWithType].
	Raw map[string]any `json:"raw"`
	// This field is from variant [ElementalImageNodeWithType].
	Src string `json:"src"`
	// This field is from variant [ElementalImageNodeWithType].
	AltText     string `json:"alt_text"`
	BorderColor string `json:"border_color"`
	BorderSize  string `json:"border_size"`
	Href        string `json:"href"`
	// This field is from variant [ElementalImageNodeWithType].
	Width string `json:"width"`
	// This field is from variant [ElementalActionNodeWithType].
	ActionID string `json:"action_id"`
	// This field is from variant [ElementalActionNodeWithType].
	BackgroundColor string `json:"background_color"`
	// This field is from variant [ElementalActionNodeWithType].
	BorderRadius string `json:"border_radius"`
	// This field is from variant [ElementalActionNodeWithType].
	DisableTracking bool `json:"disable_tracking"`
	// This field is from variant [ElementalActionNodeWithType].
	Style string `json:"style"`
	JSON  struct {
		Channels        respjson.Field
		If              respjson.Field
		Loop            respjson.Field
		Ref             respjson.Field
		Align           respjson.Field
		Bold            respjson.Field
		Color           respjson.Field
		Content         respjson.Field
		FontSize        respjson.Field
		Format          respjson.Field
		Italic          respjson.Field
		LineHeight      respjson.Field
		Locales         respjson.Field
		Strikethrough   respjson.Field
		TextStyle       respjson.Field
		Underline       respjson.Field
		Type            respjson.Field
		Title           respjson.Field
		Channel         respjson.Field
		Elements        respjson.Field
		Padding         respjson.Field
		Raw             respjson.Field
		Src             respjson.Field
		AltText         respjson.Field
		BorderColor     respjson.Field
		BorderSize      respjson.Field
		Href            respjson.Field
		Width           respjson.Field
		ActionID        respjson.Field
		BackgroundColor respjson.Field
		BorderRadius    respjson.Field
		DisableTracking respjson.Field
		Style           respjson.Field
		raw             string
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

func (u ElementalNodeUnion) AsElementalHTMLNodeWithType() (v ElementalHTMLNodeWithType) {
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

func ElementalNodeParamOfElementalImageNodeWithType(src string) ElementalNodeUnionParam {
	var variant ElementalImageNodeWithTypeParam
	variant.Src = src
	return ElementalNodeUnionParam{OfElementalImageNodeWithType: &variant}
}

func ElementalNodeParamOfElementalActionNodeWithType(content string, href string) ElementalNodeUnionParam {
	var variant ElementalActionNodeWithTypeParam
	variant.Content = content
	variant.Href = href
	return ElementalNodeUnionParam{OfElementalActionNodeWithType: &variant}
}

func ElementalNodeParamOfElementalQuoteNodeWithType(content string) ElementalNodeUnionParam {
	var variant ElementalQuoteNodeWithTypeParam
	variant.Content = content
	return ElementalNodeUnionParam{OfElementalQuoteNodeWithType: &variant}
}

func ElementalNodeParamOfElementalHTMLNodeWithType(content string) ElementalNodeUnionParam {
	var variant ElementalHTMLNodeWithTypeParam
	variant.Content = content
	return ElementalNodeUnionParam{OfElementalHTMLNodeWithType: &variant}
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
	OfElementalHTMLNodeWithType    *ElementalHTMLNodeWithTypeParam    `json:",omitzero,inline"`
	paramUnion
}

func (u ElementalNodeUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfElementalTextNodeWithType,
		u.OfElementalMetaNodeWithType,
		u.OfElementalChannelNodeWithType,
		u.OfElementalImageNodeWithType,
		u.OfElementalActionNodeWithType,
		u.OfElementalDividerNodeWithType,
		u.OfElementalQuoteNodeWithType,
		u.OfElementalHTMLNodeWithType)
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
	} else if !param.IsOmitted(u.OfElementalHTMLNodeWithType) {
		return u.OfElementalHTMLNodeWithType
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetBold() *string {
	if vt := u.OfElementalTextNodeWithType; vt != nil && vt.Bold.Valid() {
		return &vt.Bold.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetFormat() *string {
	if vt := u.OfElementalTextNodeWithType; vt != nil {
		return &vt.Format
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetItalic() *string {
	if vt := u.OfElementalTextNodeWithType; vt != nil && vt.Italic.Valid() {
		return &vt.Italic.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetStrikethrough() *string {
	if vt := u.OfElementalTextNodeWithType; vt != nil && vt.Strikethrough.Valid() {
		return &vt.Strikethrough.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetUnderline() *string {
	if vt := u.OfElementalTextNodeWithType; vt != nil && vt.Underline.Valid() {
		return &vt.Underline.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetTitle() *string {
	if vt := u.OfElementalMetaNodeWithType; vt != nil && vt.Title.Valid() {
		return &vt.Title.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetChannel() *string {
	if vt := u.OfElementalChannelNodeWithType; vt != nil && vt.Channel.Valid() {
		return &vt.Channel.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetElements() []ElementalNodeNonChannelUnionParam {
	if vt := u.OfElementalChannelNodeWithType; vt != nil {
		return vt.Elements
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
func (u ElementalNodeUnionParam) GetSrc() *string {
	if vt := u.OfElementalImageNodeWithType; vt != nil {
		return &vt.Src
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetAltText() *string {
	if vt := u.OfElementalImageNodeWithType; vt != nil && vt.AltText.Valid() {
		return &vt.AltText.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetWidth() *string {
	if vt := u.OfElementalImageNodeWithType; vt != nil && vt.Width.Valid() {
		return &vt.Width.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetActionID() *string {
	if vt := u.OfElementalActionNodeWithType; vt != nil && vt.ActionID.Valid() {
		return &vt.ActionID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetBackgroundColor() *string {
	if vt := u.OfElementalActionNodeWithType; vt != nil && vt.BackgroundColor.Valid() {
		return &vt.BackgroundColor.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetBorderRadius() *string {
	if vt := u.OfElementalActionNodeWithType; vt != nil && vt.BorderRadius.Valid() {
		return &vt.BorderRadius.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetDisableTracking() *bool {
	if vt := u.OfElementalActionNodeWithType; vt != nil && vt.DisableTracking.Valid() {
		return &vt.DisableTracking.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetStyle() *string {
	if vt := u.OfElementalActionNodeWithType; vt != nil {
		return &vt.Style
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
	} else if vt := u.OfElementalHTMLNodeWithType; vt != nil && vt.If.Valid() {
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
	} else if vt := u.OfElementalHTMLNodeWithType; vt != nil && vt.Loop.Valid() {
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
	} else if vt := u.OfElementalHTMLNodeWithType; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetAlign() *string {
	if vt := u.OfElementalTextNodeWithType; vt != nil {
		return (*string)(&vt.Align)
	} else if vt := u.OfElementalImageNodeWithType; vt != nil {
		return (*string)(&vt.Align)
	} else if vt := u.OfElementalActionNodeWithType; vt != nil {
		return (*string)(&vt.Align)
	} else if vt := u.OfElementalQuoteNodeWithType; vt != nil {
		return (*string)(&vt.Align)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetColor() *string {
	if vt := u.OfElementalTextNodeWithType; vt != nil && vt.Color.Valid() {
		return &vt.Color.Value
	} else if vt := u.OfElementalDividerNodeWithType; vt != nil && vt.Color.Valid() {
		return &vt.Color.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetContent() *string {
	if vt := u.OfElementalTextNodeWithType; vt != nil && vt.Content.Valid() {
		return &vt.Content.Value
	} else if vt := u.OfElementalActionNodeWithType; vt != nil {
		return (*string)(&vt.Content)
	} else if vt := u.OfElementalQuoteNodeWithType; vt != nil {
		return (*string)(&vt.Content)
	} else if vt := u.OfElementalHTMLNodeWithType; vt != nil {
		return (*string)(&vt.Content)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetFontSize() *string {
	if vt := u.OfElementalTextNodeWithType; vt != nil && vt.FontSize.Valid() {
		return &vt.FontSize.Value
	} else if vt := u.OfElementalChannelNodeWithType; vt != nil && vt.FontSize.Valid() {
		return &vt.FontSize.Value
	} else if vt := u.OfElementalActionNodeWithType; vt != nil && vt.FontSize.Valid() {
		return &vt.FontSize.Value
	} else if vt := u.OfElementalQuoteNodeWithType; vt != nil && vt.FontSize.Valid() {
		return &vt.FontSize.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetLineHeight() *string {
	if vt := u.OfElementalTextNodeWithType; vt != nil && vt.LineHeight.Valid() {
		return &vt.LineHeight.Value
	} else if vt := u.OfElementalChannelNodeWithType; vt != nil && vt.LineHeight.Valid() {
		return &vt.LineHeight.Value
	} else if vt := u.OfElementalQuoteNodeWithType; vt != nil && vt.LineHeight.Valid() {
		return &vt.LineHeight.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetTextStyle() *string {
	if vt := u.OfElementalTextNodeWithType; vt != nil {
		return (*string)(&vt.TextStyle)
	} else if vt := u.OfElementalQuoteNodeWithType; vt != nil {
		return (*string)(&vt.TextStyle)
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
	} else if vt := u.OfElementalHTMLNodeWithType; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetPadding() *string {
	if vt := u.OfElementalChannelNodeWithType; vt != nil && vt.Padding.Valid() {
		return &vt.Padding.Value
	} else if vt := u.OfElementalImageNodeWithType; vt != nil && vt.Padding.Valid() {
		return &vt.Padding.Value
	} else if vt := u.OfElementalActionNodeWithType; vt != nil && vt.Padding.Valid() {
		return &vt.Padding.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetBorderColor() *string {
	if vt := u.OfElementalImageNodeWithType; vt != nil && vt.BorderColor.Valid() {
		return &vt.BorderColor.Value
	} else if vt := u.OfElementalQuoteNodeWithType; vt != nil && vt.BorderColor.Valid() {
		return &vt.BorderColor.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetBorderSize() *string {
	if vt := u.OfElementalImageNodeWithType; vt != nil && vt.BorderSize.Valid() {
		return &vt.BorderSize.Value
	} else if vt := u.OfElementalActionNodeWithType; vt != nil && vt.BorderSize.Valid() {
		return &vt.BorderSize.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetHref() *string {
	if vt := u.OfElementalImageNodeWithType; vt != nil && vt.Href.Valid() {
		return &vt.Href.Value
	} else if vt := u.OfElementalActionNodeWithType; vt != nil {
		return (*string)(&vt.Href)
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
	} else if vt := u.OfElementalHTMLNodeWithType; vt != nil {
		return vt.Channels
	}
	return nil
}

// Returns a pointer to the underlying variant's Locales property, if present.
func (u ElementalNodeUnionParam) GetLocales() LocalesParam {
	if vt := u.OfElementalTextNodeWithType; vt != nil {
		return vt.Locales
	} else if vt := u.OfElementalActionNodeWithType; vt != nil {
		return vt.Locales
	} else if vt := u.OfElementalQuoteNodeWithType; vt != nil {
		return vt.Locales
	} else if vt := u.OfElementalHTMLNodeWithType; vt != nil {
		return vt.Locales
	}
	return nil
}

// ElementalNodeNonChannelUnion contains all possible properties and values from
// [ElementalNodeNonChannelObject], [ElementalNodeNonChannelObject2],
// [ElementalNodeNonChannelObject3], [ElementalNodeNonChannelObject4],
// [ElementalNodeNonChannelObject5], [ElementalNodeNonChannelObject6],
// [ElementalNodeNonChannelObject7].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ElementalNodeNonChannelUnion struct {
	// This field is from variant [ElementalNodeNonChannelObject],
	// [ElementalNodeNonChannelObject2], [ElementalNodeNonChannelObject3],
	// [ElementalNodeNonChannelObject4], [ElementalNodeNonChannelObject5],
	// [ElementalNodeNonChannelObject6], [ElementalNodeNonChannelObject7].
	Channels []string `json:"channels"`
	// This field is from variant [ElementalNodeNonChannelObject],
	// [ElementalNodeNonChannelObject2], [ElementalNodeNonChannelObject3],
	// [ElementalNodeNonChannelObject4], [ElementalNodeNonChannelObject5],
	// [ElementalNodeNonChannelObject6], [ElementalNodeNonChannelObject7].
	If string `json:"if"`
	// This field is from variant [ElementalNodeNonChannelObject],
	// [ElementalNodeNonChannelObject2], [ElementalNodeNonChannelObject3],
	// [ElementalNodeNonChannelObject4], [ElementalNodeNonChannelObject5],
	// [ElementalNodeNonChannelObject6], [ElementalNodeNonChannelObject7].
	Loop string `json:"loop"`
	// This field is from variant [ElementalNodeNonChannelObject],
	// [ElementalNodeNonChannelObject2], [ElementalNodeNonChannelObject3],
	// [ElementalNodeNonChannelObject4], [ElementalNodeNonChannelObject5],
	// [ElementalNodeNonChannelObject6], [ElementalNodeNonChannelObject7].
	Ref   string `json:"ref"`
	Align string `json:"align"`
	// This field is from variant [ElementalNodeNonChannelObject].
	Bold     string `json:"bold"`
	Color    string `json:"color"`
	Content  string `json:"content"`
	FontSize string `json:"font_size"`
	// This field is from variant [ElementalNodeNonChannelObject].
	Format string `json:"format"`
	// This field is from variant [ElementalNodeNonChannelObject].
	Italic     string `json:"italic"`
	LineHeight string `json:"line_height"`
	// This field is from variant [ElementalNodeNonChannelObject].
	Locales Locales `json:"locales"`
	// This field is from variant [ElementalNodeNonChannelObject].
	Strikethrough string `json:"strikethrough"`
	// This field is from variant [ElementalNodeNonChannelObject].
	TextStyle TextStyle `json:"text_style"`
	// This field is from variant [ElementalNodeNonChannelObject].
	Underline string `json:"underline"`
	Type      string `json:"type"`
	// This field is from variant [ElementalNodeNonChannelObject2].
	Title string `json:"title"`
	// This field is from variant [ElementalNodeNonChannelObject3].
	Src string `json:"src"`
	// This field is from variant [ElementalNodeNonChannelObject3].
	AltText     string `json:"alt_text"`
	BorderColor string `json:"border_color"`
	BorderSize  string `json:"border_size"`
	Href        string `json:"href"`
	Padding     string `json:"padding"`
	// This field is from variant [ElementalNodeNonChannelObject3].
	Width string `json:"width"`
	// This field is from variant [ElementalNodeNonChannelObject4].
	ActionID string `json:"action_id"`
	// This field is from variant [ElementalNodeNonChannelObject4].
	BackgroundColor string `json:"background_color"`
	// This field is from variant [ElementalNodeNonChannelObject4].
	BorderRadius string `json:"border_radius"`
	// This field is from variant [ElementalNodeNonChannelObject4].
	DisableTracking bool `json:"disable_tracking"`
	// This field is from variant [ElementalNodeNonChannelObject4].
	Style string `json:"style"`
	JSON  struct {
		Channels        respjson.Field
		If              respjson.Field
		Loop            respjson.Field
		Ref             respjson.Field
		Align           respjson.Field
		Bold            respjson.Field
		Color           respjson.Field
		Content         respjson.Field
		FontSize        respjson.Field
		Format          respjson.Field
		Italic          respjson.Field
		LineHeight      respjson.Field
		Locales         respjson.Field
		Strikethrough   respjson.Field
		TextStyle       respjson.Field
		Underline       respjson.Field
		Type            respjson.Field
		Title           respjson.Field
		Src             respjson.Field
		AltText         respjson.Field
		BorderColor     respjson.Field
		BorderSize      respjson.Field
		Href            respjson.Field
		Padding         respjson.Field
		Width           respjson.Field
		ActionID        respjson.Field
		BackgroundColor respjson.Field
		BorderRadius    respjson.Field
		DisableTracking respjson.Field
		Style           respjson.Field
		raw             string
	} `json:"-"`
}

func (u ElementalNodeNonChannelUnion) AsElementalNodeNonChannelObject() (v ElementalNodeNonChannelObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ElementalNodeNonChannelUnion) AsElementalNodeNonChannelObject2() (v ElementalNodeNonChannelObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ElementalNodeNonChannelUnion) AsElementalNodeNonChannelObject3() (v ElementalNodeNonChannelObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ElementalNodeNonChannelUnion) AsElementalNodeNonChannelObject4() (v ElementalNodeNonChannelObject4) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ElementalNodeNonChannelUnion) AsElementalNodeNonChannelObject5() (v ElementalNodeNonChannelObject5) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ElementalNodeNonChannelUnion) AsElementalNodeNonChannelObject6() (v ElementalNodeNonChannelObject6) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ElementalNodeNonChannelUnion) AsElementalNodeNonChannelObject7() (v ElementalNodeNonChannelObject7) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ElementalNodeNonChannelUnion) RawJSON() string { return u.JSON.raw }

func (r *ElementalNodeNonChannelUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ElementalNodeNonChannelUnion to a
// ElementalNodeNonChannelUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ElementalNodeNonChannelUnionParam.Overrides()
func (r ElementalNodeNonChannelUnion) ToParam() ElementalNodeNonChannelUnionParam {
	return param.Override[ElementalNodeNonChannelUnionParam](json.RawMessage(r.RawJSON()))
}

// Represents a body of text to be rendered inside of the notification.
type ElementalNodeNonChannelObject struct {
	// Any of "text".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ElementalTextNode
}

// Returns the unmodified JSON received from the API
func (r ElementalNodeNonChannelObject) RawJSON() string { return r.JSON.raw }
func (r *ElementalNodeNonChannelObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The meta element contains information describing the notification that may be
// used by a particular channel or provider. One important field is the title field
// which will be used as the title for channels that support it.
type ElementalNodeNonChannelObject2 struct {
	// Any of "meta".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ElementalMetaNode
}

// Returns the unmodified JSON received from the API
func (r ElementalNodeNonChannelObject2) RawJSON() string { return r.JSON.raw }
func (r *ElementalNodeNonChannelObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Used to embed an image into the notification.
type ElementalNodeNonChannelObject3 struct {
	// Any of "image".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ElementalImageNode
}

// Returns the unmodified JSON received from the API
func (r ElementalNodeNonChannelObject3) RawJSON() string { return r.JSON.raw }
func (r *ElementalNodeNonChannelObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows the user to execute an action. Can be a button or a link.
type ElementalNodeNonChannelObject4 struct {
	// Any of "action".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ElementalActionNode
}

// Returns the unmodified JSON received from the API
func (r ElementalNodeNonChannelObject4) RawJSON() string { return r.JSON.raw }
func (r *ElementalNodeNonChannelObject4) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Renders a dividing line between elements.
type ElementalNodeNonChannelObject5 struct {
	// Any of "divider".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ElementalDividerNode
}

// Returns the unmodified JSON received from the API
func (r ElementalNodeNonChannelObject5) RawJSON() string { return r.JSON.raw }
func (r *ElementalNodeNonChannelObject5) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Renders a quote block.
type ElementalNodeNonChannelObject6 struct {
	// Any of "quote".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ElementalQuoteNode
}

// Returns the unmodified JSON received from the API
func (r ElementalNodeNonChannelObject6) RawJSON() string { return r.JSON.raw }
func (r *ElementalNodeNonChannelObject6) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Raw HTML string inside an Elemental document. When rendering a message, this
// node is turned into output only for the email channel; for other channels it
// produces no blocks.
type ElementalNodeNonChannelObject7 struct {
	// Any of "html".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ElementalHTMLNode
}

// Returns the unmodified JSON received from the API
func (r ElementalNodeNonChannelObject7) RawJSON() string { return r.JSON.raw }
func (r *ElementalNodeNonChannelObject7) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func ElementalNodeNonChannelParamOfElementalNodeNonChannelObject3(src string) ElementalNodeNonChannelUnionParam {
	var variant ElementalNodeNonChannelObject3Param
	variant.Src = src
	return ElementalNodeNonChannelUnionParam{OfElementalNodeNonChannelObject3: &variant}
}

func ElementalNodeNonChannelParamOfElementalNodeNonChannelObject4(content string, href string) ElementalNodeNonChannelUnionParam {
	var variant ElementalNodeNonChannelObject4Param
	variant.Content = content
	variant.Href = href
	return ElementalNodeNonChannelUnionParam{OfElementalNodeNonChannelObject4: &variant}
}

func ElementalNodeNonChannelParamOfElementalNodeNonChannelObject6(content string) ElementalNodeNonChannelUnionParam {
	var variant ElementalNodeNonChannelObject6Param
	variant.Content = content
	return ElementalNodeNonChannelUnionParam{OfElementalNodeNonChannelObject6: &variant}
}

func ElementalNodeNonChannelParamOfElementalNodeNonChannelObject7(content string) ElementalNodeNonChannelUnionParam {
	var variant ElementalNodeNonChannelObject7Param
	variant.Content = content
	return ElementalNodeNonChannelUnionParam{OfElementalNodeNonChannelObject7: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ElementalNodeNonChannelUnionParam struct {
	OfElementalNodeNonChannelObject  *ElementalNodeNonChannelObjectParam  `json:",omitzero,inline"`
	OfElementalNodeNonChannelObject2 *ElementalNodeNonChannelObject2Param `json:",omitzero,inline"`
	OfElementalNodeNonChannelObject3 *ElementalNodeNonChannelObject3Param `json:",omitzero,inline"`
	OfElementalNodeNonChannelObject4 *ElementalNodeNonChannelObject4Param `json:",omitzero,inline"`
	OfElementalNodeNonChannelObject5 *ElementalNodeNonChannelObject5Param `json:",omitzero,inline"`
	OfElementalNodeNonChannelObject6 *ElementalNodeNonChannelObject6Param `json:",omitzero,inline"`
	OfElementalNodeNonChannelObject7 *ElementalNodeNonChannelObject7Param `json:",omitzero,inline"`
	paramUnion
}

func (u ElementalNodeNonChannelUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfElementalNodeNonChannelObject,
		u.OfElementalNodeNonChannelObject2,
		u.OfElementalNodeNonChannelObject3,
		u.OfElementalNodeNonChannelObject4,
		u.OfElementalNodeNonChannelObject5,
		u.OfElementalNodeNonChannelObject6,
		u.OfElementalNodeNonChannelObject7)
}
func (u *ElementalNodeNonChannelUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ElementalNodeNonChannelUnionParam) asAny() any {
	if !param.IsOmitted(u.OfElementalNodeNonChannelObject) {
		return u.OfElementalNodeNonChannelObject
	} else if !param.IsOmitted(u.OfElementalNodeNonChannelObject2) {
		return u.OfElementalNodeNonChannelObject2
	} else if !param.IsOmitted(u.OfElementalNodeNonChannelObject3) {
		return u.OfElementalNodeNonChannelObject3
	} else if !param.IsOmitted(u.OfElementalNodeNonChannelObject4) {
		return u.OfElementalNodeNonChannelObject4
	} else if !param.IsOmitted(u.OfElementalNodeNonChannelObject5) {
		return u.OfElementalNodeNonChannelObject5
	} else if !param.IsOmitted(u.OfElementalNodeNonChannelObject6) {
		return u.OfElementalNodeNonChannelObject6
	} else if !param.IsOmitted(u.OfElementalNodeNonChannelObject7) {
		return u.OfElementalNodeNonChannelObject7
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeNonChannelUnionParam) GetBold() *string {
	if vt := u.OfElementalNodeNonChannelObject; vt != nil && vt.Bold.Valid() {
		return &vt.Bold.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeNonChannelUnionParam) GetFormat() *string {
	if vt := u.OfElementalNodeNonChannelObject; vt != nil {
		return &vt.Format
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeNonChannelUnionParam) GetItalic() *string {
	if vt := u.OfElementalNodeNonChannelObject; vt != nil && vt.Italic.Valid() {
		return &vt.Italic.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeNonChannelUnionParam) GetStrikethrough() *string {
	if vt := u.OfElementalNodeNonChannelObject; vt != nil && vt.Strikethrough.Valid() {
		return &vt.Strikethrough.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeNonChannelUnionParam) GetUnderline() *string {
	if vt := u.OfElementalNodeNonChannelObject; vt != nil && vt.Underline.Valid() {
		return &vt.Underline.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeNonChannelUnionParam) GetTitle() *string {
	if vt := u.OfElementalNodeNonChannelObject2; vt != nil && vt.Title.Valid() {
		return &vt.Title.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeNonChannelUnionParam) GetSrc() *string {
	if vt := u.OfElementalNodeNonChannelObject3; vt != nil {
		return &vt.Src
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeNonChannelUnionParam) GetAltText() *string {
	if vt := u.OfElementalNodeNonChannelObject3; vt != nil && vt.AltText.Valid() {
		return &vt.AltText.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeNonChannelUnionParam) GetWidth() *string {
	if vt := u.OfElementalNodeNonChannelObject3; vt != nil && vt.Width.Valid() {
		return &vt.Width.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeNonChannelUnionParam) GetActionID() *string {
	if vt := u.OfElementalNodeNonChannelObject4; vt != nil && vt.ActionID.Valid() {
		return &vt.ActionID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeNonChannelUnionParam) GetBackgroundColor() *string {
	if vt := u.OfElementalNodeNonChannelObject4; vt != nil && vt.BackgroundColor.Valid() {
		return &vt.BackgroundColor.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeNonChannelUnionParam) GetBorderRadius() *string {
	if vt := u.OfElementalNodeNonChannelObject4; vt != nil && vt.BorderRadius.Valid() {
		return &vt.BorderRadius.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeNonChannelUnionParam) GetDisableTracking() *bool {
	if vt := u.OfElementalNodeNonChannelObject4; vt != nil && vt.DisableTracking.Valid() {
		return &vt.DisableTracking.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeNonChannelUnionParam) GetStyle() *string {
	if vt := u.OfElementalNodeNonChannelObject4; vt != nil {
		return &vt.Style
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeNonChannelUnionParam) GetIf() *string {
	if vt := u.OfElementalNodeNonChannelObject; vt != nil && vt.If.Valid() {
		return &vt.If.Value
	} else if vt := u.OfElementalNodeNonChannelObject2; vt != nil && vt.If.Valid() {
		return &vt.If.Value
	} else if vt := u.OfElementalNodeNonChannelObject3; vt != nil && vt.If.Valid() {
		return &vt.If.Value
	} else if vt := u.OfElementalNodeNonChannelObject4; vt != nil && vt.If.Valid() {
		return &vt.If.Value
	} else if vt := u.OfElementalNodeNonChannelObject5; vt != nil && vt.If.Valid() {
		return &vt.If.Value
	} else if vt := u.OfElementalNodeNonChannelObject6; vt != nil && vt.If.Valid() {
		return &vt.If.Value
	} else if vt := u.OfElementalNodeNonChannelObject7; vt != nil && vt.If.Valid() {
		return &vt.If.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeNonChannelUnionParam) GetLoop() *string {
	if vt := u.OfElementalNodeNonChannelObject; vt != nil && vt.Loop.Valid() {
		return &vt.Loop.Value
	} else if vt := u.OfElementalNodeNonChannelObject2; vt != nil && vt.Loop.Valid() {
		return &vt.Loop.Value
	} else if vt := u.OfElementalNodeNonChannelObject3; vt != nil && vt.Loop.Valid() {
		return &vt.Loop.Value
	} else if vt := u.OfElementalNodeNonChannelObject4; vt != nil && vt.Loop.Valid() {
		return &vt.Loop.Value
	} else if vt := u.OfElementalNodeNonChannelObject5; vt != nil && vt.Loop.Valid() {
		return &vt.Loop.Value
	} else if vt := u.OfElementalNodeNonChannelObject6; vt != nil && vt.Loop.Valid() {
		return &vt.Loop.Value
	} else if vt := u.OfElementalNodeNonChannelObject7; vt != nil && vt.Loop.Valid() {
		return &vt.Loop.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeNonChannelUnionParam) GetRef() *string {
	if vt := u.OfElementalNodeNonChannelObject; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	} else if vt := u.OfElementalNodeNonChannelObject2; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	} else if vt := u.OfElementalNodeNonChannelObject3; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	} else if vt := u.OfElementalNodeNonChannelObject4; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	} else if vt := u.OfElementalNodeNonChannelObject5; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	} else if vt := u.OfElementalNodeNonChannelObject6; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	} else if vt := u.OfElementalNodeNonChannelObject7; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeNonChannelUnionParam) GetAlign() *string {
	if vt := u.OfElementalNodeNonChannelObject; vt != nil {
		return (*string)(&vt.Align)
	} else if vt := u.OfElementalNodeNonChannelObject3; vt != nil {
		return (*string)(&vt.Align)
	} else if vt := u.OfElementalNodeNonChannelObject4; vt != nil {
		return (*string)(&vt.Align)
	} else if vt := u.OfElementalNodeNonChannelObject6; vt != nil {
		return (*string)(&vt.Align)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeNonChannelUnionParam) GetColor() *string {
	if vt := u.OfElementalNodeNonChannelObject; vt != nil && vt.Color.Valid() {
		return &vt.Color.Value
	} else if vt := u.OfElementalNodeNonChannelObject5; vt != nil && vt.Color.Valid() {
		return &vt.Color.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeNonChannelUnionParam) GetContent() *string {
	if vt := u.OfElementalNodeNonChannelObject; vt != nil && vt.Content.Valid() {
		return &vt.Content.Value
	} else if vt := u.OfElementalNodeNonChannelObject4; vt != nil {
		return (*string)(&vt.Content)
	} else if vt := u.OfElementalNodeNonChannelObject6; vt != nil {
		return (*string)(&vt.Content)
	} else if vt := u.OfElementalNodeNonChannelObject7; vt != nil {
		return (*string)(&vt.Content)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeNonChannelUnionParam) GetFontSize() *string {
	if vt := u.OfElementalNodeNonChannelObject; vt != nil && vt.FontSize.Valid() {
		return &vt.FontSize.Value
	} else if vt := u.OfElementalNodeNonChannelObject4; vt != nil && vt.FontSize.Valid() {
		return &vt.FontSize.Value
	} else if vt := u.OfElementalNodeNonChannelObject6; vt != nil && vt.FontSize.Valid() {
		return &vt.FontSize.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeNonChannelUnionParam) GetLineHeight() *string {
	if vt := u.OfElementalNodeNonChannelObject; vt != nil && vt.LineHeight.Valid() {
		return &vt.LineHeight.Value
	} else if vt := u.OfElementalNodeNonChannelObject6; vt != nil && vt.LineHeight.Valid() {
		return &vt.LineHeight.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeNonChannelUnionParam) GetTextStyle() *string {
	if vt := u.OfElementalNodeNonChannelObject; vt != nil {
		return (*string)(&vt.TextStyle)
	} else if vt := u.OfElementalNodeNonChannelObject6; vt != nil {
		return (*string)(&vt.TextStyle)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeNonChannelUnionParam) GetType() *string {
	if vt := u.OfElementalNodeNonChannelObject; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfElementalNodeNonChannelObject2; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfElementalNodeNonChannelObject3; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfElementalNodeNonChannelObject4; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfElementalNodeNonChannelObject5; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfElementalNodeNonChannelObject6; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfElementalNodeNonChannelObject7; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeNonChannelUnionParam) GetBorderColor() *string {
	if vt := u.OfElementalNodeNonChannelObject3; vt != nil && vt.BorderColor.Valid() {
		return &vt.BorderColor.Value
	} else if vt := u.OfElementalNodeNonChannelObject6; vt != nil && vt.BorderColor.Valid() {
		return &vt.BorderColor.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeNonChannelUnionParam) GetBorderSize() *string {
	if vt := u.OfElementalNodeNonChannelObject3; vt != nil && vt.BorderSize.Valid() {
		return &vt.BorderSize.Value
	} else if vt := u.OfElementalNodeNonChannelObject4; vt != nil && vt.BorderSize.Valid() {
		return &vt.BorderSize.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeNonChannelUnionParam) GetHref() *string {
	if vt := u.OfElementalNodeNonChannelObject3; vt != nil && vt.Href.Valid() {
		return &vt.Href.Value
	} else if vt := u.OfElementalNodeNonChannelObject4; vt != nil {
		return (*string)(&vt.Href)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeNonChannelUnionParam) GetPadding() *string {
	if vt := u.OfElementalNodeNonChannelObject3; vt != nil && vt.Padding.Valid() {
		return &vt.Padding.Value
	} else if vt := u.OfElementalNodeNonChannelObject4; vt != nil && vt.Padding.Valid() {
		return &vt.Padding.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's Channels property, if present.
func (u ElementalNodeNonChannelUnionParam) GetChannels() []string {
	if vt := u.OfElementalNodeNonChannelObject; vt != nil {
		return vt.Channels
	} else if vt := u.OfElementalNodeNonChannelObject2; vt != nil {
		return vt.Channels
	} else if vt := u.OfElementalNodeNonChannelObject3; vt != nil {
		return vt.Channels
	} else if vt := u.OfElementalNodeNonChannelObject4; vt != nil {
		return vt.Channels
	} else if vt := u.OfElementalNodeNonChannelObject5; vt != nil {
		return vt.Channels
	} else if vt := u.OfElementalNodeNonChannelObject6; vt != nil {
		return vt.Channels
	} else if vt := u.OfElementalNodeNonChannelObject7; vt != nil {
		return vt.Channels
	}
	return nil
}

// Returns a pointer to the underlying variant's Locales property, if present.
func (u ElementalNodeNonChannelUnionParam) GetLocales() LocalesParam {
	if vt := u.OfElementalNodeNonChannelObject; vt != nil {
		return vt.Locales
	} else if vt := u.OfElementalNodeNonChannelObject4; vt != nil {
		return vt.Locales
	} else if vt := u.OfElementalNodeNonChannelObject6; vt != nil {
		return vt.Locales
	} else if vt := u.OfElementalNodeNonChannelObject7; vt != nil {
		return vt.Locales
	}
	return nil
}

// Represents a body of text to be rendered inside of the notification.
type ElementalNodeNonChannelObjectParam struct {
	Type string `json:"type,omitzero"`
	ElementalTextNodeParam
}

func (r ElementalNodeNonChannelObjectParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*ElementalNodeNonChannelObjectParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// The meta element contains information describing the notification that may be
// used by a particular channel or provider. One important field is the title field
// which will be used as the title for channels that support it.
type ElementalNodeNonChannelObject2Param struct {
	Type string `json:"type,omitzero"`
	ElementalMetaNodeParam
}

func (r ElementalNodeNonChannelObject2Param) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*ElementalNodeNonChannelObject2Param
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// Used to embed an image into the notification.
type ElementalNodeNonChannelObject3Param struct {
	Type string `json:"type,omitzero"`
	ElementalImageNodeParam
}

func (r ElementalNodeNonChannelObject3Param) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*ElementalNodeNonChannelObject3Param
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// Allows the user to execute an action. Can be a button or a link.
type ElementalNodeNonChannelObject4Param struct {
	Type string `json:"type,omitzero"`
	ElementalActionNodeParam
}

func (r ElementalNodeNonChannelObject4Param) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*ElementalNodeNonChannelObject4Param
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// Renders a dividing line between elements.
type ElementalNodeNonChannelObject5Param struct {
	Type string `json:"type,omitzero"`
	ElementalDividerNodeParam
}

func (r ElementalNodeNonChannelObject5Param) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*ElementalNodeNonChannelObject5Param
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// Renders a quote block.
type ElementalNodeNonChannelObject6Param struct {
	Type string `json:"type,omitzero"`
	ElementalQuoteNodeParam
}

func (r ElementalNodeNonChannelObject6Param) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*ElementalNodeNonChannelObject6Param
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// Raw HTML string inside an Elemental document. When rendering a message, this
// node is turned into output only for the email channel; for other channels it
// produces no blocks.
type ElementalNodeNonChannelObject7Param struct {
	Type string `json:"type,omitzero"`
	ElementalHTMLNodeParam
}

func (r ElementalNodeNonChannelObject7Param) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*ElementalNodeNonChannelObject7Param
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// Renders a quote block.
type ElementalQuoteNode struct {
	// The text value of the quote.
	Content string `json:"content" api:"required"`
	// Alignment of the quote.
	//
	// Any of "center", "left", "right", "full".
	Align Alignment `json:"align" api:"nullable"`
	// CSS border color property. For example, `#fff`
	BorderColor string `json:"border_color" api:"nullable"`
	// CSS px font size for this quote block, e.g. `16px`. Overrides the size of the
	// `text_style` preset. Email only.
	FontSize string `json:"font_size" api:"nullable"`
	// CSS line height for this quote block, as a px value or a unitless multiplier,
	// e.g. `24px` or `1.5`. Email only.
	LineHeight string `json:"line_height" api:"nullable"`
	// Region specific content. See
	// [locales docs](https://www.courier.com/docs/platform/content/elemental/locales/)
	// for more details.
	Locales Locales `json:"locales" api:"nullable"`
	// Any of "text", "h1", "h2", "subtext".
	TextStyle TextStyle `json:"text_style"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Align       respjson.Field
		BorderColor respjson.Field
		FontSize    respjson.Field
		LineHeight  respjson.Field
		Locales     respjson.Field
		TextStyle   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ElementalBaseNode
}

// Returns the unmodified JSON received from the API
func (r ElementalQuoteNode) RawJSON() string { return r.JSON.raw }
func (r *ElementalQuoteNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ElementalQuoteNode to a ElementalQuoteNodeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ElementalQuoteNodeParam.Overrides()
func (r ElementalQuoteNode) ToParam() ElementalQuoteNodeParam {
	return param.Override[ElementalQuoteNodeParam](json.RawMessage(r.RawJSON()))
}

// Renders a quote block.
type ElementalQuoteNodeParam struct {
	// The text value of the quote.
	Content string `json:"content" api:"required"`
	// Alignment of the quote.
	Align Alignment `json:"align,omitzero"`
	// CSS border color property. For example, `#fff`
	BorderColor param.Opt[string] `json:"border_color,omitzero"`
	// CSS px font size for this quote block, e.g. `16px`. Overrides the size of the
	// `text_style` preset. Email only.
	FontSize param.Opt[string] `json:"font_size,omitzero"`
	// CSS line height for this quote block, as a px value or a unitless multiplier,
	// e.g. `24px` or `1.5`. Email only.
	LineHeight param.Opt[string] `json:"line_height,omitzero"`
	// Region specific content. See
	// [locales docs](https://www.courier.com/docs/platform/content/elemental/locales/)
	// for more details.
	Locales   LocalesParam `json:"locales,omitzero"`
	TextStyle TextStyle    `json:"text_style,omitzero"`
	ElementalBaseNodeParam
}

func (r ElementalQuoteNodeParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*ElementalQuoteNodeParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// Renders a quote block.
type ElementalQuoteNodeWithType struct {
	// Any of "quote".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ElementalQuoteNode
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

// Renders a quote block.
type ElementalQuoteNodeWithTypeParam struct {
	Type string `json:"type,omitzero"`
	ElementalQuoteNodeParam
}

func (r ElementalQuoteNodeWithTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*ElementalQuoteNodeWithTypeParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// Represents a body of text to be rendered inside of the notification.
type ElementalTextNode struct {
	// Text alignment.
	//
	// Any of "left", "center", "right".
	Align string `json:"align"`
	// Apply bold to the text
	Bold string `json:"bold" api:"nullable"`
	// Specifies the color of text. Can be any valid css color value
	Color string `json:"color" api:"nullable"`
	// The text content displayed in the notification. Either this field must be
	// specified, or the elements field
	Content string `json:"content"`
	// CSS px font size for this text block, e.g. `16px`. Overrides the size of the
	// `text_style` preset. Email only.
	FontSize string `json:"font_size" api:"nullable"`
	// Any of "markdown".
	Format string `json:"format" api:"nullable"`
	// Apply italics to the text
	Italic string `json:"italic" api:"nullable"`
	// CSS line height for this text block, as a px value or a unitless multiplier,
	// e.g. `24px` or `1.5`. Email only.
	LineHeight string `json:"line_height" api:"nullable"`
	// Region specific content. See
	// [locales docs](https://www.courier.com/docs/platform/content/elemental/locales/)
	// for more details.
	Locales Locales `json:"locales" api:"nullable"`
	// Apply a strike through the text
	Strikethrough string `json:"strikethrough" api:"nullable"`
	// Allows the text to be rendered as a heading level.
	//
	// Any of "text", "h1", "h2", "subtext".
	TextStyle TextStyle `json:"text_style" api:"nullable"`
	// Apply an underline to the text
	Underline string `json:"underline" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Align         respjson.Field
		Bold          respjson.Field
		Color         respjson.Field
		Content       respjson.Field
		FontSize      respjson.Field
		Format        respjson.Field
		Italic        respjson.Field
		LineHeight    respjson.Field
		Locales       respjson.Field
		Strikethrough respjson.Field
		TextStyle     respjson.Field
		Underline     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
	ElementalBaseNode
}

// Returns the unmodified JSON received from the API
func (r ElementalTextNode) RawJSON() string { return r.JSON.raw }
func (r *ElementalTextNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ElementalTextNode to a ElementalTextNodeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ElementalTextNodeParam.Overrides()
func (r ElementalTextNode) ToParam() ElementalTextNodeParam {
	return param.Override[ElementalTextNodeParam](json.RawMessage(r.RawJSON()))
}

// Represents a body of text to be rendered inside of the notification.
type ElementalTextNodeParam struct {
	// Text alignment.
	Align string `json:"align,omitzero"`
	// Apply bold to the text
	Bold param.Opt[string] `json:"bold,omitzero"`
	// Specifies the color of text. Can be any valid css color value
	Color param.Opt[string] `json:"color,omitzero"`
	// The text content displayed in the notification. Either this field must be
	// specified, or the elements field
	Content param.Opt[string] `json:"content,omitzero"`
	// CSS px font size for this text block, e.g. `16px`. Overrides the size of the
	// `text_style` preset. Email only.
	FontSize param.Opt[string] `json:"font_size,omitzero"`
	Format   string            `json:"format,omitzero"`
	// Apply italics to the text
	Italic param.Opt[string] `json:"italic,omitzero"`
	// CSS line height for this text block, as a px value or a unitless multiplier,
	// e.g. `24px` or `1.5`. Email only.
	LineHeight param.Opt[string] `json:"line_height,omitzero"`
	// Region specific content. See
	// [locales docs](https://www.courier.com/docs/platform/content/elemental/locales/)
	// for more details.
	Locales LocalesParam `json:"locales,omitzero"`
	// Apply a strike through the text
	Strikethrough param.Opt[string] `json:"strikethrough,omitzero"`
	// Allows the text to be rendered as a heading level.
	TextStyle TextStyle `json:"text_style,omitzero"`
	// Apply an underline to the text
	Underline param.Opt[string] `json:"underline,omitzero"`
	ElementalBaseNodeParam
}

func (r ElementalTextNodeParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*ElementalTextNodeParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// Represents a body of text to be rendered inside of the notification.
type ElementalTextNodeWithType struct {
	// Any of "text".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ElementalTextNode
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

// Represents a body of text to be rendered inside of the notification.
type ElementalTextNodeWithTypeParam struct {
	Type string `json:"type,omitzero"`
	ElementalTextNodeParam
}

func (r ElementalTextNodeWithTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*ElementalTextNodeWithTypeParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// A filter rule that can be either a single condition (with path/value) or a
// nested group (with filters array). Use comparison operators (EQ, GT, etc.) for
// single conditions, and logical operators (AND, OR) for nested groups.
type FilterConfig struct {
	// The operator for this filter. Use comparison operators (EQ, GT, LT, GTE, LTE,
	// NEQ, EXISTS, INCLUDES, STARTS_WITH, ENDS_WITH, IS_BEFORE, IS_AFTER, OMIT) for
	// single conditions, or logical operators (AND, OR) for nested filter groups.
	Operator string `json:"operator" api:"required"`
	// Nested filter rules to combine with AND/OR. Required for nested filter groups,
	// not used for single filter conditions.
	Filters []FilterConfig `json:"filters"`
	// The attribute path from the user profile to filter on. Required for single
	// filter conditions, not used for nested filter groups.
	Path string `json:"path"`
	// The value to compare against. Required for single filter conditions, not used
	// for nested filter groups.
	Value string `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator    respjson.Field
		Filters     respjson.Field
		Path        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FilterConfig) RawJSON() string { return r.JSON.raw }
func (r *FilterConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this FilterConfig to a FilterConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// FilterConfigParam.Overrides()
func (r FilterConfig) ToParam() FilterConfigParam {
	return param.Override[FilterConfigParam](json.RawMessage(r.RawJSON()))
}

// A filter rule that can be either a single condition (with path/value) or a
// nested group (with filters array). Use comparison operators (EQ, GT, etc.) for
// single conditions, and logical operators (AND, OR) for nested groups.
//
// The property Operator is required.
type FilterConfigParam struct {
	// The operator for this filter. Use comparison operators (EQ, GT, LT, GTE, LTE,
	// NEQ, EXISTS, INCLUDES, STARTS_WITH, ENDS_WITH, IS_BEFORE, IS_AFTER, OMIT) for
	// single conditions, or logical operators (AND, OR) for nested filter groups.
	Operator string `json:"operator" api:"required"`
	// The attribute path from the user profile to filter on. Required for single
	// filter conditions, not used for nested filter groups.
	Path param.Opt[string] `json:"path,omitzero"`
	// The value to compare against. Required for single filter conditions, not used
	// for nested filter groups.
	Value param.Opt[string] `json:"value,omitzero"`
	// Nested filter rules to combine with AND/OR. Required for nested filter groups,
	// not used for single filter conditions.
	Filters []FilterConfigParam `json:"filters,omitzero"`
	paramObj
}

func (r FilterConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow FilterConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FilterConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Operator, Path, Value are required.
type ListFilterParam struct {
	// Send to users only if they are member of the account
	//
	// Any of "MEMBER_OF".
	Operator ListFilterOperator `json:"operator,omitzero" api:"required"`
	// Any of "account_id".
	Path  ListFilterPath `json:"path,omitzero" api:"required"`
	Value string         `json:"value" api:"required"`
	paramObj
}

func (r ListFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow ListFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ListFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Send to users only if they are member of the account
type ListFilterOperator string

const (
	ListFilterOperatorMemberOf ListFilterOperator = "MEMBER_OF"
)

type ListFilterPath string

const (
	ListFilterPathAccountID ListFilterPath = "account_id"
)

// Send to users in lists matching a pattern
type ListPatternRecipientParam struct {
	ListPattern param.Opt[string] `json:"list_pattern,omitzero"`
	Data        map[string]any    `json:"data,omitzero"`
	paramObj
}

func (r ListPatternRecipientParam) MarshalJSON() (data []byte, err error) {
	type shadow ListPatternRecipientParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ListPatternRecipientParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Send to all users in a specific list
type ListRecipientParam struct {
	ListID  param.Opt[string] `json:"list_id,omitzero"`
	Data    map[string]any    `json:"data,omitzero"`
	Filters []ListFilterParam `json:"filters,omitzero"`
	paramObj
}

func (r ListRecipientParam) MarshalJSON() (data []byte, err error) {
	type shadow ListRecipientParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ListRecipientParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Locales map[string]Locale

type Locale struct {
	Content string `json:"content" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Locale) RawJSON() string { return r.JSON.raw }
func (r *Locale) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LocalesParam map[string]LocaleParam

// The property Content is required.
type LocaleParam struct {
	Content string `json:"content" api:"required"`
	paramObj
}

func (r LocaleParam) MarshalJSON() (data []byte, err error) {
	type shadow LocaleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LocaleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageChannels map[string]Channel

type MessageChannelsParam map[string]ChannelParam

type MessageContext struct {
	// Tenant id used to load brand/default preferences/context.
	TenantID string `json:"tenant_id" api:"nullable"`
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

type MessageProviders map[string]MessageProvidersType

type MessageProvidersParam map[string]MessageProvidersTypeParam

type MessageProvidersType struct {
	// JS conditional with access to data/profile.
	If       string   `json:"if" api:"nullable"`
	Metadata Metadata `json:"metadata" api:"nullable"`
	// Provider-specific overrides.
	Override map[string]any `json:"override" api:"nullable"`
	Timeouts int64          `json:"timeouts" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		If          respjson.Field
		Metadata    respjson.Field
		Override    respjson.Field
		Timeouts    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageProvidersType) RawJSON() string { return r.JSON.raw }
func (r *MessageProvidersType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MessageProvidersType to a MessageProvidersTypeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MessageProvidersTypeParam.Overrides()
func (r MessageProvidersType) ToParam() MessageProvidersTypeParam {
	return param.Override[MessageProvidersTypeParam](json.RawMessage(r.RawJSON()))
}

type MessageProvidersTypeParam struct {
	// JS conditional with access to data/profile.
	If       param.Opt[string] `json:"if,omitzero"`
	Timeouts param.Opt[int64]  `json:"timeouts,omitzero"`
	// Provider-specific overrides.
	Override map[string]any `json:"override,omitzero"`
	Metadata MetadataParam  `json:"metadata,omitzero"`
	paramObj
}

func (r MessageProvidersTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageProvidersTypeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageProvidersTypeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageRouting struct {
	Channels []MessageRoutingChannelUnion `json:"channels" api:"required"`
	// Any of "all", "single".
	Method MessageRoutingMethod `json:"method" api:"required"`
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
	Channels []MessageRoutingChannelUnionParam `json:"channels,omitzero" api:"required"`
	// Any of "all", "single".
	Method MessageRoutingMethod `json:"method,omitzero" api:"required"`
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

type Metadata struct {
	Utm Utm `json:"utm" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Utm         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Metadata) RawJSON() string { return r.JSON.raw }
func (r *Metadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Metadata to a MetadataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MetadataParam.Overrides()
func (r Metadata) ToParam() MetadataParam {
	return param.Override[MetadataParam](json.RawMessage(r.RawJSON()))
}

type MetadataParam struct {
	Utm UtmParam `json:"utm,omitzero"`
	paramObj
}

func (r MetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow MetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func MsTeamsParamOfSendToMsTeamsUserID(userID string) MsTeamsUnionParam {
	var variant SendToMsTeamsUserIDParam
	variant.UserID = userID
	return MsTeamsUnionParam{OfSendToMsTeamsUserID: &variant}
}

func MsTeamsParamOfSendToMsTeamsEmail(email string) MsTeamsUnionParam {
	var variant SendToMsTeamsEmailParam
	variant.Email = email
	return MsTeamsUnionParam{OfSendToMsTeamsEmail: &variant}
}

func MsTeamsParamOfSendToMsTeamsChannelID(channelID string) MsTeamsUnionParam {
	var variant SendToMsTeamsChannelIDParam
	variant.ChannelID = channelID
	return MsTeamsUnionParam{OfSendToMsTeamsChannelID: &variant}
}

func MsTeamsParamOfSendToMsTeamsConversationID(conversationID string, serviceURL string, tenantID string) MsTeamsUnionParam {
	var variant SendToMsTeamsConversationIDParam
	variant.ConversationID = conversationID
	variant.ServiceURL = serviceURL
	variant.TenantID = tenantID
	return MsTeamsUnionParam{OfSendToMsTeamsConversationID: &variant}
}

func MsTeamsParamOfSendToMsTeamsChannelName(channelName string, teamID string) MsTeamsUnionParam {
	var variant SendToMsTeamsChannelNameParam
	variant.ChannelName = channelName
	variant.TeamID = teamID
	return MsTeamsUnionParam{OfSendToMsTeamsChannelName: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MsTeamsUnionParam struct {
	OfSendToMsTeamsUserID         *SendToMsTeamsUserIDParam         `json:",omitzero,inline"`
	OfSendToMsTeamsEmail          *SendToMsTeamsEmailParam          `json:",omitzero,inline"`
	OfSendToMsTeamsChannelID      *SendToMsTeamsChannelIDParam      `json:",omitzero,inline"`
	OfSendToMsTeamsConversationID *SendToMsTeamsConversationIDParam `json:",omitzero,inline"`
	OfSendToMsTeamsChannelName    *SendToMsTeamsChannelNameParam    `json:",omitzero,inline"`
	paramUnion
}

func (u MsTeamsUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSendToMsTeamsUserID,
		u.OfSendToMsTeamsEmail,
		u.OfSendToMsTeamsChannelID,
		u.OfSendToMsTeamsConversationID,
		u.OfSendToMsTeamsChannelName)
}
func (u *MsTeamsUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *MsTeamsUnionParam) asAny() any {
	if !param.IsOmitted(u.OfSendToMsTeamsUserID) {
		return u.OfSendToMsTeamsUserID
	} else if !param.IsOmitted(u.OfSendToMsTeamsEmail) {
		return u.OfSendToMsTeamsEmail
	} else if !param.IsOmitted(u.OfSendToMsTeamsChannelID) {
		return u.OfSendToMsTeamsChannelID
	} else if !param.IsOmitted(u.OfSendToMsTeamsConversationID) {
		return u.OfSendToMsTeamsConversationID
	} else if !param.IsOmitted(u.OfSendToMsTeamsChannelName) {
		return u.OfSendToMsTeamsChannelName
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MsTeamsUnionParam) GetUserID() *string {
	if vt := u.OfSendToMsTeamsUserID; vt != nil {
		return &vt.UserID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MsTeamsUnionParam) GetEmail() *string {
	if vt := u.OfSendToMsTeamsEmail; vt != nil {
		return &vt.Email
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MsTeamsUnionParam) GetChannelID() *string {
	if vt := u.OfSendToMsTeamsChannelID; vt != nil {
		return &vt.ChannelID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MsTeamsUnionParam) GetConversationID() *string {
	if vt := u.OfSendToMsTeamsConversationID; vt != nil {
		return &vt.ConversationID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MsTeamsUnionParam) GetChannelName() *string {
	if vt := u.OfSendToMsTeamsChannelName; vt != nil {
		return &vt.ChannelName
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MsTeamsUnionParam) GetTeamID() *string {
	if vt := u.OfSendToMsTeamsChannelName; vt != nil {
		return &vt.TeamID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MsTeamsUnionParam) GetServiceURL() *string {
	if vt := u.OfSendToMsTeamsUserID; vt != nil && vt.ServiceURL.Valid() {
		return &vt.ServiceURL.Value
	} else if vt := u.OfSendToMsTeamsEmail; vt != nil && vt.ServiceURL.Valid() {
		return &vt.ServiceURL.Value
	} else if vt := u.OfSendToMsTeamsChannelID; vt != nil && vt.ServiceURL.Valid() {
		return &vt.ServiceURL.Value
	} else if vt := u.OfSendToMsTeamsConversationID; vt != nil {
		return (*string)(&vt.ServiceURL)
	} else if vt := u.OfSendToMsTeamsChannelName; vt != nil && vt.ServiceURL.Valid() {
		return &vt.ServiceURL.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MsTeamsUnionParam) GetTenantID() *string {
	if vt := u.OfSendToMsTeamsUserID; vt != nil && vt.TenantID.Valid() {
		return &vt.TenantID.Value
	} else if vt := u.OfSendToMsTeamsEmail; vt != nil && vt.TenantID.Valid() {
		return &vt.TenantID.Value
	} else if vt := u.OfSendToMsTeamsChannelID; vt != nil && vt.TenantID.Valid() {
		return &vt.TenantID.Value
	} else if vt := u.OfSendToMsTeamsConversationID; vt != nil {
		return (*string)(&vt.TenantID)
	} else if vt := u.OfSendToMsTeamsChannelName; vt != nil && vt.TenantID.Valid() {
		return &vt.TenantID.Value
	}
	return nil
}

// Send via Microsoft Teams
//
// The property MsTeams is required.
type MsTeamsRecipientParam struct {
	// Provide at least one of `tenant_id` or `service_url`. If you provide both, they
	// must agree.
	MsTeams MsTeamsUnionParam `json:"ms_teams,omitzero" api:"required"`
	paramObj
}

func (r MsTeamsRecipientParam) MarshalJSON() (data []byte, err error) {
	type shadow MsTeamsRecipientParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MsTeamsRecipientParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationPreferenceDetails struct {
	// Any of "OPTED_IN", "OPTED_OUT", "REQUIRED".
	Status             PreferenceStatus    `json:"status" api:"required"`
	ChannelPreferences []ChannelPreference `json:"channel_preferences" api:"nullable"`
	Rules              []Rule              `json:"rules" api:"nullable"`
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
	Status             PreferenceStatus         `json:"status,omitzero" api:"required"`
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

type PagerdutyParam struct {
	EventAction param.Opt[string] `json:"event_action,omitzero"`
	RoutingKey  param.Opt[string] `json:"routing_key,omitzero"`
	Severity    param.Opt[string] `json:"severity,omitzero"`
	Source      param.Opt[string] `json:"source,omitzero"`
	paramObj
}

func (r PagerdutyParam) MarshalJSON() (data []byte, err error) {
	type shadow PagerdutyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PagerdutyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Send via PagerDuty
//
// The property Pagerduty is required.
type PagerdutyRecipientParam struct {
	Pagerduty PagerdutyParam `json:"pagerduty,omitzero" api:"required"`
	paramObj
}

func (r PagerdutyRecipientParam) MarshalJSON() (data []byte, err error) {
	type shadow PagerdutyRecipientParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PagerdutyRecipientParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Paging struct {
	More   bool   `json:"more" api:"required"`
	Cursor string `json:"cursor" api:"nullable"`
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
	Status             PreferenceStatus    `json:"status" api:"required"`
	ChannelPreferences []ChannelPreference `json:"channel_preferences" api:"nullable"`
	Rules              []Rule              `json:"rules" api:"nullable"`
	// Any of "subscription", "list", "recipient".
	Source PreferenceSource `json:"source" api:"nullable"`
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
	Status             PreferenceStatus         `json:"status,omitzero" api:"required"`
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

type RecipientPreferences struct {
	Categories    map[string]NotificationPreferenceDetails `json:"categories" api:"nullable"`
	Notifications map[string]NotificationPreferenceDetails `json:"notifications" api:"nullable"`
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

type RecipientPreferencesParam struct {
	Categories    map[string]NotificationPreferenceDetailsParam `json:"categories,omitzero"`
	Notifications map[string]NotificationPreferenceDetailsParam `json:"notifications,omitzero"`
	paramObj
}

func (r RecipientPreferencesParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientPreferencesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecipientPreferencesParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Rule struct {
	Until string `json:"until" api:"required"`
	Start string `json:"start" api:"nullable"`
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
	Until string            `json:"until" api:"required"`
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

// Sends directly to a Microsoft Teams channel by its Bot Framework ID. Still
// provide at least one of `tenant_id` or `service_url` — sends without either have
// failed Bot Framework authentication in testing.
//
// The property ChannelID is required.
type SendToMsTeamsChannelIDParam struct {
	ChannelID  string            `json:"channel_id" api:"required"`
	ServiceURL param.Opt[string] `json:"service_url,omitzero"`
	TenantID   param.Opt[string] `json:"tenant_id,omitzero"`
	paramObj
}

func (r SendToMsTeamsChannelIDParam) MarshalJSON() (data []byte, err error) {
	type shadow SendToMsTeamsChannelIDParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendToMsTeamsChannelIDParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// `team_id` is required alongside `channel_name`. Also provide at least one of
// `tenant_id` or `service_url`; if you provide both, they must agree.
//
// The properties ChannelName, TeamID are required.
type SendToMsTeamsChannelNameParam struct {
	ChannelName string            `json:"channel_name" api:"required"`
	TeamID      string            `json:"team_id" api:"required"`
	ServiceURL  param.Opt[string] `json:"service_url,omitzero"`
	TenantID    param.Opt[string] `json:"tenant_id,omitzero"`
	paramObj
}

func (r SendToMsTeamsChannelNameParam) MarshalJSON() (data []byte, err error) {
	type shadow SendToMsTeamsChannelNameParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendToMsTeamsChannelNameParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ConversationID, ServiceURL, TenantID are required.
type SendToMsTeamsConversationIDParam struct {
	ConversationID string `json:"conversation_id" api:"required"`
	ServiceURL     string `json:"service_url" api:"required"`
	TenantID       string `json:"tenant_id" api:"required"`
	paramObj
}

func (r SendToMsTeamsConversationIDParam) MarshalJSON() (data []byte, err error) {
	type shadow SendToMsTeamsConversationIDParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendToMsTeamsConversationIDParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provide at least one of `tenant_id` or `service_url`. If you provide both, they
// must agree.
//
// The property Email is required.
type SendToMsTeamsEmailParam struct {
	Email      string            `json:"email" api:"required"`
	ServiceURL param.Opt[string] `json:"service_url,omitzero"`
	TenantID   param.Opt[string] `json:"tenant_id,omitzero"`
	paramObj
}

func (r SendToMsTeamsEmailParam) MarshalJSON() (data []byte, err error) {
	type shadow SendToMsTeamsEmailParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendToMsTeamsEmailParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provide at least one of `tenant_id` or `service_url`. If you provide both, they
// must agree.
//
// The property UserID is required.
type SendToMsTeamsUserIDParam struct {
	UserID     string            `json:"user_id" api:"required"`
	ServiceURL param.Opt[string] `json:"service_url,omitzero"`
	TenantID   param.Opt[string] `json:"tenant_id,omitzero"`
	paramObj
}

func (r SendToMsTeamsUserIDParam) MarshalJSON() (data []byte, err error) {
	type shadow SendToMsTeamsUserIDParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendToMsTeamsUserIDParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessToken, Channel are required.
type SendToSlackChannelParam struct {
	AccessToken string `json:"access_token" api:"required"`
	Channel     string `json:"channel" api:"required"`
	paramObj
}

func (r SendToSlackChannelParam) MarshalJSON() (data []byte, err error) {
	type shadow SendToSlackChannelParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendToSlackChannelParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessToken, Email are required.
type SendToSlackEmailParam struct {
	AccessToken string `json:"access_token" api:"required"`
	Email       string `json:"email" api:"required"`
	paramObj
}

func (r SendToSlackEmailParam) MarshalJSON() (data []byte, err error) {
	type shadow SendToSlackEmailParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendToSlackEmailParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessToken, UserID are required.
type SendToSlackUserIDParam struct {
	AccessToken string `json:"access_token" api:"required"`
	UserID      string `json:"user_id" api:"required"`
	paramObj
}

func (r SendToSlackUserIDParam) MarshalJSON() (data []byte, err error) {
	type shadow SendToSlackUserIDParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendToSlackUserIDParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func SlackParamOfSendToSlackChannel(accessToken string, channel string) SlackUnionParam {
	var variant SendToSlackChannelParam
	variant.AccessToken = accessToken
	variant.Channel = channel
	return SlackUnionParam{OfSendToSlackChannel: &variant}
}

func SlackParamOfSendToSlackEmail(accessToken string, email string) SlackUnionParam {
	var variant SendToSlackEmailParam
	variant.AccessToken = accessToken
	variant.Email = email
	return SlackUnionParam{OfSendToSlackEmail: &variant}
}

func SlackParamOfSendToSlackUserID(accessToken string, userID string) SlackUnionParam {
	var variant SendToSlackUserIDParam
	variant.AccessToken = accessToken
	variant.UserID = userID
	return SlackUnionParam{OfSendToSlackUserID: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SlackUnionParam struct {
	OfSendToSlackChannel *SendToSlackChannelParam `json:",omitzero,inline"`
	OfSendToSlackEmail   *SendToSlackEmailParam   `json:",omitzero,inline"`
	OfSendToSlackUserID  *SendToSlackUserIDParam  `json:",omitzero,inline"`
	paramUnion
}

func (u SlackUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSendToSlackChannel, u.OfSendToSlackEmail, u.OfSendToSlackUserID)
}
func (u *SlackUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SlackUnionParam) asAny() any {
	if !param.IsOmitted(u.OfSendToSlackChannel) {
		return u.OfSendToSlackChannel
	} else if !param.IsOmitted(u.OfSendToSlackEmail) {
		return u.OfSendToSlackEmail
	} else if !param.IsOmitted(u.OfSendToSlackUserID) {
		return u.OfSendToSlackUserID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SlackUnionParam) GetChannel() *string {
	if vt := u.OfSendToSlackChannel; vt != nil {
		return &vt.Channel
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SlackUnionParam) GetEmail() *string {
	if vt := u.OfSendToSlackEmail; vt != nil {
		return &vt.Email
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SlackUnionParam) GetUserID() *string {
	if vt := u.OfSendToSlackUserID; vt != nil {
		return &vt.UserID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SlackUnionParam) GetAccessToken() *string {
	if vt := u.OfSendToSlackChannel; vt != nil {
		return (*string)(&vt.AccessToken)
	} else if vt := u.OfSendToSlackEmail; vt != nil {
		return (*string)(&vt.AccessToken)
	} else if vt := u.OfSendToSlackUserID; vt != nil {
		return (*string)(&vt.AccessToken)
	}
	return nil
}

// Send via Slack (channel, email, or user_id)
//
// The property Slack is required.
type SlackRecipientParam struct {
	Slack SlackUnionParam `json:"slack,omitzero" api:"required"`
	paramObj
}

func (r SlackRecipientParam) MarshalJSON() (data []byte, err error) {
	type shadow SlackRecipientParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SlackRecipientParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TextStyle string

const (
	TextStyleText    TextStyle = "text"
	TextStyleH1      TextStyle = "h1"
	TextStyleH2      TextStyle = "h2"
	TextStyleSubtext TextStyle = "subtext"
)

type Timeouts struct {
	Channel  int64 `json:"channel" api:"nullable"`
	Provider int64 `json:"provider" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Channel     respjson.Field
		Provider    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Timeouts) RawJSON() string { return r.JSON.raw }
func (r *Timeouts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Timeouts to a TimeoutsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TimeoutsParam.Overrides()
func (r Timeouts) ToParam() TimeoutsParam {
	return param.Override[TimeoutsParam](json.RawMessage(r.RawJSON()))
}

type TimeoutsParam struct {
	Channel  param.Opt[int64] `json:"channel,omitzero"`
	Provider param.Opt[int64] `json:"provider,omitzero"`
	paramObj
}

func (r TimeoutsParam) MarshalJSON() (data []byte, err error) {
	type shadow TimeoutsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TimeoutsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserRecipient struct {
	// Deprecated - Use `tenant_id` instead.
	AccountID string `json:"account_id" api:"nullable"`
	// Context such as tenant_id to send the notification with.
	Context MessageContext `json:"context" api:"nullable"`
	Data    map[string]any `json:"data" api:"nullable"`
	// The user's email address.
	Email string `json:"email" api:"nullable"`
	// The id of the list to send the message to.
	ListID string `json:"list_id" api:"nullable"`
	// The user's preferred ISO 639-1 language code.
	Locale string `json:"locale" api:"nullable"`
	// The user's phone number.
	PhoneNumber string                   `json:"phone_number" api:"nullable"`
	Preferences UserRecipientPreferences `json:"preferences" api:"nullable"`
	// The id of the tenant the user is associated with.
	TenantID string `json:"tenant_id" api:"nullable"`
	// The user's unique identifier. Typically, this will match the user id of a user
	// in your system.
	UserID string `json:"user_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountID   respjson.Field
		Context     respjson.Field
		Data        respjson.Field
		Email       respjson.Field
		ListID      respjson.Field
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
	Notifications map[string]Preference `json:"notifications" api:"required"`
	Categories    map[string]Preference `json:"categories" api:"nullable"`
	TemplateID    string                `json:"templateId" api:"nullable"`
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
	// The id of the list to send the message to.
	ListID param.Opt[string] `json:"list_id,omitzero"`
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
	Notifications map[string]PreferenceParam `json:"notifications,omitzero" api:"required"`
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

type Utm struct {
	Campaign string `json:"campaign" api:"nullable"`
	Content  string `json:"content" api:"nullable"`
	Medium   string `json:"medium" api:"nullable"`
	Source   string `json:"source" api:"nullable"`
	Term     string `json:"term" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Campaign    respjson.Field
		Content     respjson.Field
		Medium      respjson.Field
		Source      respjson.Field
		Term        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Utm) RawJSON() string { return r.JSON.raw }
func (r *Utm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Utm to a UtmParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// UtmParam.Overrides()
func (r Utm) ToParam() UtmParam {
	return param.Override[UtmParam](json.RawMessage(r.RawJSON()))
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

type WebhookAuthMode string

const (
	WebhookAuthModeNone   WebhookAuthMode = "none"
	WebhookAuthModeBasic  WebhookAuthMode = "basic"
	WebhookAuthModeBearer WebhookAuthMode = "bearer"
)

// The property Mode is required.
type WebhookAuthenticationParam struct {
	// The authentication mode to use. Defaults to 'none' if not specified.
	//
	// Any of "none", "basic", "bearer".
	Mode WebhookAuthMode `json:"mode,omitzero" api:"required"`
	// Token for bearer authentication.
	Token param.Opt[string] `json:"token,omitzero"`
	// Password for basic authentication.
	Password param.Opt[string] `json:"password,omitzero"`
	// Username for basic authentication.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

func (r WebhookAuthenticationParam) MarshalJSON() (data []byte, err error) {
	type shadow WebhookAuthenticationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookAuthenticationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookMethod string

const (
	WebhookMethodPost WebhookMethod = "POST"
	WebhookMethodPut  WebhookMethod = "PUT"
)

// The property URL is required.
type WebhookProfileParam struct {
	// The URL to send the webhook request to.
	URL string `json:"url" api:"required"`
	// Custom headers to include in the webhook request.
	Headers map[string]string `json:"headers,omitzero"`
	// Authentication configuration for the webhook request.
	Authentication WebhookAuthenticationParam `json:"authentication,omitzero"`
	// The HTTP method to use for the webhook request. Defaults to POST if not
	// specified.
	//
	// Any of "POST", "PUT".
	Method WebhookMethod `json:"method,omitzero"`
	// Specifies what profile information is included in the request payload. Defaults
	// to 'limited' if not specified.
	//
	// Any of "limited", "expanded".
	Profile WebhookProfileType `json:"profile,omitzero"`
	paramObj
}

func (r WebhookProfileParam) MarshalJSON() (data []byte, err error) {
	type shadow WebhookProfileParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookProfileParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookProfileType string

const (
	WebhookProfileTypeLimited  WebhookProfileType = "limited"
	WebhookProfileTypeExpanded WebhookProfileType = "expanded"
)

// Send via webhook
//
// The property Webhook is required.
type WebhookRecipientParam struct {
	Webhook WebhookProfileParam `json:"webhook,omitzero" api:"required"`
	paramObj
}

func (r WebhookRecipientParam) MarshalJSON() (data []byte, err error) {
	type shadow WebhookRecipientParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookRecipientParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
