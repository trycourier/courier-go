// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
	"net/http"
	"slices"

	"github.com/trycourier/courier-go/internal/apijson"
	"github.com/trycourier/courier-go/internal/requestconfig"
	"github.com/trycourier/courier-go/option"
	"github.com/trycourier/courier-go/packages/param"
	"github.com/trycourier/courier-go/packages/respjson"
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

type MessageRoutingMethod string

const (
	MessageRoutingMethodAll    MessageRoutingMethod = "all"
	MessageRoutingMethodSingle MessageRoutingMethod = "single"
)

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

// The property Status is required.
type PreferenceParam struct {
	// Any of "OPTED_IN", "OPTED_OUT", "REQUIRED".
	Status             PreferenceStatus                   `json:"status,omitzero,required"`
	ChannelPreferences []PreferenceChannelPreferenceParam `json:"channel_preferences,omitzero"`
	Rules              []PreferenceRuleParam              `json:"rules,omitzero"`
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

// The property Channel is required.
type PreferenceChannelPreferenceParam struct {
	// Any of "direct_message", "email", "push", "sms", "webhook", "inbox".
	Channel string `json:"channel,omitzero,required"`
	paramObj
}

func (r PreferenceChannelPreferenceParam) MarshalJSON() (data []byte, err error) {
	type shadow PreferenceChannelPreferenceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PreferenceChannelPreferenceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PreferenceChannelPreferenceParam](
		"channel", "direct_message", "email", "push", "sms", "webhook", "inbox",
	)
}

// The property Until is required.
type PreferenceRuleParam struct {
	Until string            `json:"until,required"`
	Start param.Opt[string] `json:"start,omitzero"`
	paramObj
}

func (r PreferenceRuleParam) MarshalJSON() (data []byte, err error) {
	type shadow PreferenceRuleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PreferenceRuleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PreferenceSource string

const (
	PreferenceSourceSubscription PreferenceSource = "subscription"
	PreferenceSourceList         PreferenceSource = "list"
	PreferenceSourceRecipient    PreferenceSource = "recipient"
)

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
//
// The property Content is required.
type SendMessageParamsMessage struct {
	// Describes content that will work for email, inbox, push, chat, or any channel
	// id.
	Content SendMessageParamsMessageContentUnion `json:"content,omitzero,required"`
	BrandID param.Opt[string]                    `json:"brand_id,omitzero"`
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
	To      SendMessageParamsMessageToUnion `json:"to,omitzero"`
	Context MessageContextParam             `json:"context,omitzero"`
	paramObj
}

func (r SendMessageParamsMessage) MarshalJSON() (data []byte, err error) {
	type shadow SendMessageParamsMessage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendMessageParamsMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SendMessageParamsMessageContentUnion struct {
	OfElementalContentSugar *SendMessageParamsMessageContentElementalContentSugar `json:",omitzero,inline"`
	OfElementalContent      *SendMessageParamsMessageContentElementalContent      `json:",omitzero,inline"`
	paramUnion
}

func (u SendMessageParamsMessageContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfElementalContentSugar, u.OfElementalContent)
}
func (u *SendMessageParamsMessageContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SendMessageParamsMessageContentUnion) asAny() any {
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
type SendMessageParamsMessageContentElementalContentSugar struct {
	// The text content displayed in the notification.
	Body string `json:"body,required"`
	// Title/subject displayed by supported channels.
	Title string `json:"title,required"`
	paramObj
}

func (r SendMessageParamsMessageContentElementalContentSugar) MarshalJSON() (data []byte, err error) {
	type shadow SendMessageParamsMessageContentElementalContentSugar
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendMessageParamsMessageContentElementalContentSugar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Elements, Version are required.
type SendMessageParamsMessageContentElementalContent struct {
	Elements []ElementalNodeUnionParam `json:"elements,omitzero,required"`
	// For example, "2022-01-01"
	Version string `json:"version,required"`
	Brand   any    `json:"brand,omitzero"`
	paramObj
}

func (r SendMessageParamsMessageContentElementalContent) MarshalJSON() (data []byte, err error) {
	type shadow SendMessageParamsMessageContentElementalContent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendMessageParamsMessageContentElementalContent) UnmarshalJSON(data []byte) error {
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
	OfSendMessagesMessageToObject *SendMessageParamsMessageToObject `json:",omitzero,inline"`
	OfRecipientArray              []RecipientParam                  `json:",omitzero,inline"`
	paramUnion
}

func (u SendMessageParamsMessageToUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSendMessagesMessageToObject, u.OfRecipientArray)
}
func (u *SendMessageParamsMessageToUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SendMessageParamsMessageToUnion) asAny() any {
	if !param.IsOmitted(u.OfSendMessagesMessageToObject) {
		return u.OfSendMessagesMessageToObject
	} else if !param.IsOmitted(u.OfRecipientArray) {
		return &u.OfRecipientArray
	}
	return nil
}

type SendMessageParamsMessageToObject struct {
	// Use `tenant_id` instead.
	AccountID param.Opt[string] `json:"account_id,omitzero"`
	Email     param.Opt[string] `json:"email,omitzero"`
	// The user's preferred ISO 639-1 language code.
	Locale      param.Opt[string] `json:"locale,omitzero"`
	PhoneNumber param.Opt[string] `json:"phone_number,omitzero"`
	// Tenant id. Will load brand, default preferences and base context data.
	TenantID    param.Opt[string]                           `json:"tenant_id,omitzero"`
	UserID      param.Opt[string]                           `json:"user_id,omitzero"`
	Data        map[string]any                              `json:"data,omitzero"`
	Preferences SendMessageParamsMessageToObjectPreferences `json:"preferences,omitzero"`
	// Context such as tenant_id to send the notification with.
	Context MessageContextParam `json:"context,omitzero"`
	paramObj
}

func (r SendMessageParamsMessageToObject) MarshalJSON() (data []byte, err error) {
	type shadow SendMessageParamsMessageToObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendMessageParamsMessageToObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Notifications is required.
type SendMessageParamsMessageToObjectPreferences struct {
	Notifications map[string]PreferenceParam `json:"notifications,omitzero,required"`
	TemplateID    param.Opt[string]          `json:"templateId,omitzero"`
	Categories    map[string]PreferenceParam `json:"categories,omitzero"`
	paramObj
}

func (r SendMessageParamsMessageToObjectPreferences) MarshalJSON() (data []byte, err error) {
	type shadow SendMessageParamsMessageToObjectPreferences
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendMessageParamsMessageToObjectPreferences) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
