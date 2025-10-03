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
func (r *SendService) SendMessage(ctx context.Context, body SendSendMessageParams, opts ...option.RequestOption) (res *SendSendMessageResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "send"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
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
	Notifications map[string]RecipientPreferencesNotificationParam `json:"notifications,omitzero,required"`
	TemplateID    param.Opt[string]                                `json:"templateId,omitzero"`
	Categories    map[string]RecipientPreferencesCategoryParam     `json:"categories,omitzero"`
	paramObj
}

func (r RecipientPreferencesParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientPreferencesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecipientPreferencesParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Status is required.
type RecipientPreferencesNotificationParam struct {
	// Any of "OPTED_IN", "OPTED_OUT", "REQUIRED".
	Status             string                                                   `json:"status,omitzero,required"`
	ChannelPreferences []RecipientPreferencesNotificationChannelPreferenceParam `json:"channel_preferences,omitzero"`
	Rules              []RecipientPreferencesNotificationRuleParam              `json:"rules,omitzero"`
	// Any of "subscription", "list", "recipient".
	Source string `json:"source,omitzero"`
	paramObj
}

func (r RecipientPreferencesNotificationParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientPreferencesNotificationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecipientPreferencesNotificationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RecipientPreferencesNotificationParam](
		"status", "OPTED_IN", "OPTED_OUT", "REQUIRED",
	)
	apijson.RegisterFieldValidator[RecipientPreferencesNotificationParam](
		"source", "subscription", "list", "recipient",
	)
}

// The property Channel is required.
type RecipientPreferencesNotificationChannelPreferenceParam struct {
	// Any of "direct_message", "email", "push", "sms", "webhook", "inbox".
	Channel string `json:"channel,omitzero,required"`
	paramObj
}

func (r RecipientPreferencesNotificationChannelPreferenceParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientPreferencesNotificationChannelPreferenceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecipientPreferencesNotificationChannelPreferenceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RecipientPreferencesNotificationChannelPreferenceParam](
		"channel", "direct_message", "email", "push", "sms", "webhook", "inbox",
	)
}

// The property Until is required.
type RecipientPreferencesNotificationRuleParam struct {
	Until string            `json:"until,required"`
	Start param.Opt[string] `json:"start,omitzero"`
	paramObj
}

func (r RecipientPreferencesNotificationRuleParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientPreferencesNotificationRuleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecipientPreferencesNotificationRuleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Status is required.
type RecipientPreferencesCategoryParam struct {
	// Any of "OPTED_IN", "OPTED_OUT", "REQUIRED".
	Status             string                                               `json:"status,omitzero,required"`
	ChannelPreferences []RecipientPreferencesCategoryChannelPreferenceParam `json:"channel_preferences,omitzero"`
	Rules              []RecipientPreferencesCategoryRuleParam              `json:"rules,omitzero"`
	// Any of "subscription", "list", "recipient".
	Source string `json:"source,omitzero"`
	paramObj
}

func (r RecipientPreferencesCategoryParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientPreferencesCategoryParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecipientPreferencesCategoryParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RecipientPreferencesCategoryParam](
		"status", "OPTED_IN", "OPTED_OUT", "REQUIRED",
	)
	apijson.RegisterFieldValidator[RecipientPreferencesCategoryParam](
		"source", "subscription", "list", "recipient",
	)
}

// The property Channel is required.
type RecipientPreferencesCategoryChannelPreferenceParam struct {
	// Any of "direct_message", "email", "push", "sms", "webhook", "inbox".
	Channel string `json:"channel,omitzero,required"`
	paramObj
}

func (r RecipientPreferencesCategoryChannelPreferenceParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientPreferencesCategoryChannelPreferenceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecipientPreferencesCategoryChannelPreferenceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RecipientPreferencesCategoryChannelPreferenceParam](
		"channel", "direct_message", "email", "push", "sms", "webhook", "inbox",
	)
}

// The property Until is required.
type RecipientPreferencesCategoryRuleParam struct {
	Until string            `json:"until,required"`
	Start param.Opt[string] `json:"start,omitzero"`
	paramObj
}

func (r RecipientPreferencesCategoryRuleParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientPreferencesCategoryRuleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecipientPreferencesCategoryRuleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SendSendMessageResponse struct {
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
func (r SendSendMessageResponse) RawJSON() string { return r.JSON.raw }
func (r *SendSendMessageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SendSendMessageParams struct {
	// The message property has the following primary top-level properties. They define
	// the destination and content of the message.
	Message SendSendMessageParamsMessage `json:"message,omitzero,required"`
	paramObj
}

func (r SendSendMessageParams) MarshalJSON() (data []byte, err error) {
	type shadow SendSendMessageParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendSendMessageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The message property has the following primary top-level properties. They define
// the destination and content of the message.
//
// The property Content is required.
type SendSendMessageParamsMessage struct {
	// Syntactic sugar to provide a fast shorthand for Courier Elemental Blocks.
	Content SendSendMessageParamsMessageContent `json:"content,omitzero,required"`
	BrandID param.Opt[string]                   `json:"brand_id,omitzero"`
	// Define run-time configuration for channels. Valid ChannelId's: email, sms, push,
	// inbox, direct_message, banner, webhook.
	Channels    map[string]SendSendMessageParamsMessageChannel  `json:"channels,omitzero"`
	Data        map[string]any                                  `json:"data,omitzero"`
	Delay       SendSendMessageParamsMessageDelay               `json:"delay,omitzero"`
	Expiry      SendSendMessageParamsMessageExpiry              `json:"expiry,omitzero"`
	Metadata    SendSendMessageParamsMessageMetadata            `json:"metadata,omitzero"`
	Preferences SendSendMessageParamsMessagePreferences         `json:"preferences,omitzero"`
	Providers   map[string]SendSendMessageParamsMessageProvider `json:"providers,omitzero"`
	// Customize which channels/providers Courier may deliver the message through.
	Routing SendSendMessageParamsMessageRouting `json:"routing,omitzero"`
	Timeout SendSendMessageParamsMessageTimeout `json:"timeout,omitzero"`
	// The recipient or a list of recipients of the message
	To      SendSendMessageParamsMessageToUnion `json:"to,omitzero"`
	Context MessageContextParam                 `json:"context,omitzero"`
	paramObj
}

func (r SendSendMessageParamsMessage) MarshalJSON() (data []byte, err error) {
	type shadow SendSendMessageParamsMessage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendSendMessageParamsMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Syntactic sugar to provide a fast shorthand for Courier Elemental Blocks.
//
// The properties Body, Title are required.
type SendSendMessageParamsMessageContent struct {
	// The text content displayed in the notification.
	Body string `json:"body,required"`
	// Title/subject displayed by supported channels.
	Title string `json:"title,required"`
	paramObj
}

func (r SendSendMessageParamsMessageContent) MarshalJSON() (data []byte, err error) {
	type shadow SendSendMessageParamsMessageContent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendSendMessageParamsMessageContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SendSendMessageParamsMessageChannel struct {
	// Brand id used for rendering.
	BrandID param.Opt[string] `json:"brand_id,omitzero"`
	// JS conditional with access to data/profile.
	If       param.Opt[string]                           `json:"if,omitzero"`
	Metadata SendSendMessageParamsMessageChannelMetadata `json:"metadata,omitzero"`
	// Channel specific overrides.
	Override map[string]any `json:"override,omitzero"`
	// Providers enabled for this channel.
	Providers []string `json:"providers,omitzero"`
	// Defaults to `single`.
	//
	// Any of "all", "single".
	RoutingMethod string                                      `json:"routing_method,omitzero"`
	Timeouts      SendSendMessageParamsMessageChannelTimeouts `json:"timeouts,omitzero"`
	paramObj
}

func (r SendSendMessageParamsMessageChannel) MarshalJSON() (data []byte, err error) {
	type shadow SendSendMessageParamsMessageChannel
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendSendMessageParamsMessageChannel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SendSendMessageParamsMessageChannel](
		"routing_method", "all", "single",
	)
}

type SendSendMessageParamsMessageChannelMetadata struct {
	Utm SendSendMessageParamsMessageChannelMetadataUtm `json:"utm,omitzero"`
	paramObj
}

func (r SendSendMessageParamsMessageChannelMetadata) MarshalJSON() (data []byte, err error) {
	type shadow SendSendMessageParamsMessageChannelMetadata
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendSendMessageParamsMessageChannelMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SendSendMessageParamsMessageChannelMetadataUtm struct {
	Campaign param.Opt[string] `json:"campaign,omitzero"`
	Content  param.Opt[string] `json:"content,omitzero"`
	Medium   param.Opt[string] `json:"medium,omitzero"`
	Source   param.Opt[string] `json:"source,omitzero"`
	Term     param.Opt[string] `json:"term,omitzero"`
	paramObj
}

func (r SendSendMessageParamsMessageChannelMetadataUtm) MarshalJSON() (data []byte, err error) {
	type shadow SendSendMessageParamsMessageChannelMetadataUtm
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendSendMessageParamsMessageChannelMetadataUtm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SendSendMessageParamsMessageChannelTimeouts struct {
	Channel  param.Opt[int64] `json:"channel,omitzero"`
	Provider param.Opt[int64] `json:"provider,omitzero"`
	paramObj
}

func (r SendSendMessageParamsMessageChannelTimeouts) MarshalJSON() (data []byte, err error) {
	type shadow SendSendMessageParamsMessageChannelTimeouts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendSendMessageParamsMessageChannelTimeouts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SendSendMessageParamsMessageDelay struct {
	// The duration of the delay in milliseconds.
	Duration param.Opt[int64] `json:"duration,omitzero"`
	// ISO 8601 timestamp or opening_hours-like format.
	Until param.Opt[string] `json:"until,omitzero"`
	paramObj
}

func (r SendSendMessageParamsMessageDelay) MarshalJSON() (data []byte, err error) {
	type shadow SendSendMessageParamsMessageDelay
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendSendMessageParamsMessageDelay) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ExpiresIn is required.
type SendSendMessageParamsMessageExpiry struct {
	// Duration in ms or ISO8601 duration (e.g. P1DT4H).
	ExpiresIn SendSendMessageParamsMessageExpiryExpiresInUnion `json:"expires_in,omitzero,required"`
	// Epoch or ISO8601 timestamp with timezone.
	ExpiresAt param.Opt[string] `json:"expires_at,omitzero"`
	paramObj
}

func (r SendSendMessageParamsMessageExpiry) MarshalJSON() (data []byte, err error) {
	type shadow SendSendMessageParamsMessageExpiry
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendSendMessageParamsMessageExpiry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SendSendMessageParamsMessageExpiryExpiresInUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u SendSendMessageParamsMessageExpiryExpiresInUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *SendSendMessageParamsMessageExpiryExpiresInUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SendSendMessageParamsMessageExpiryExpiresInUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

type SendSendMessageParamsMessageMetadata struct {
	Event   param.Opt[string]                       `json:"event,omitzero"`
	TraceID param.Opt[string]                       `json:"trace_id,omitzero"`
	Tags    []string                                `json:"tags,omitzero"`
	Utm     SendSendMessageParamsMessageMetadataUtm `json:"utm,omitzero"`
	paramObj
}

func (r SendSendMessageParamsMessageMetadata) MarshalJSON() (data []byte, err error) {
	type shadow SendSendMessageParamsMessageMetadata
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendSendMessageParamsMessageMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SendSendMessageParamsMessageMetadataUtm struct {
	Campaign param.Opt[string] `json:"campaign,omitzero"`
	Content  param.Opt[string] `json:"content,omitzero"`
	Medium   param.Opt[string] `json:"medium,omitzero"`
	Source   param.Opt[string] `json:"source,omitzero"`
	Term     param.Opt[string] `json:"term,omitzero"`
	paramObj
}

func (r SendSendMessageParamsMessageMetadataUtm) MarshalJSON() (data []byte, err error) {
	type shadow SendSendMessageParamsMessageMetadataUtm
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendSendMessageParamsMessageMetadataUtm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SubscriptionTopicID is required.
type SendSendMessageParamsMessagePreferences struct {
	// The subscription topic to apply to the message.
	SubscriptionTopicID string `json:"subscription_topic_id,required"`
	paramObj
}

func (r SendSendMessageParamsMessagePreferences) MarshalJSON() (data []byte, err error) {
	type shadow SendSendMessageParamsMessagePreferences
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendSendMessageParamsMessagePreferences) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SendSendMessageParamsMessageProvider struct {
	// JS conditional with access to data/profile.
	If       param.Opt[string]                            `json:"if,omitzero"`
	Timeouts param.Opt[int64]                             `json:"timeouts,omitzero"`
	Metadata SendSendMessageParamsMessageProviderMetadata `json:"metadata,omitzero"`
	// Provider-specific overrides.
	Override map[string]any `json:"override,omitzero"`
	paramObj
}

func (r SendSendMessageParamsMessageProvider) MarshalJSON() (data []byte, err error) {
	type shadow SendSendMessageParamsMessageProvider
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendSendMessageParamsMessageProvider) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SendSendMessageParamsMessageProviderMetadata struct {
	Utm SendSendMessageParamsMessageProviderMetadataUtm `json:"utm,omitzero"`
	paramObj
}

func (r SendSendMessageParamsMessageProviderMetadata) MarshalJSON() (data []byte, err error) {
	type shadow SendSendMessageParamsMessageProviderMetadata
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendSendMessageParamsMessageProviderMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SendSendMessageParamsMessageProviderMetadataUtm struct {
	Campaign param.Opt[string] `json:"campaign,omitzero"`
	Content  param.Opt[string] `json:"content,omitzero"`
	Medium   param.Opt[string] `json:"medium,omitzero"`
	Source   param.Opt[string] `json:"source,omitzero"`
	Term     param.Opt[string] `json:"term,omitzero"`
	paramObj
}

func (r SendSendMessageParamsMessageProviderMetadataUtm) MarshalJSON() (data []byte, err error) {
	type shadow SendSendMessageParamsMessageProviderMetadataUtm
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendSendMessageParamsMessageProviderMetadataUtm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Customize which channels/providers Courier may deliver the message through.
//
// The properties Channels, Method are required.
type SendSendMessageParamsMessageRouting struct {
	// A list of channels or providers (or nested routing rules).
	Channels []MessageRoutingChannelUnionParam `json:"channels,omitzero,required"`
	// Any of "all", "single".
	Method string `json:"method,omitzero,required"`
	paramObj
}

func (r SendSendMessageParamsMessageRouting) MarshalJSON() (data []byte, err error) {
	type shadow SendSendMessageParamsMessageRouting
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendSendMessageParamsMessageRouting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SendSendMessageParamsMessageRouting](
		"method", "all", "single",
	)
}

type SendSendMessageParamsMessageTimeout struct {
	Escalation param.Opt[int64] `json:"escalation,omitzero"`
	Message    param.Opt[int64] `json:"message,omitzero"`
	Channel    map[string]int64 `json:"channel,omitzero"`
	// Any of "no-escalation", "delivered", "viewed", "engaged".
	Criteria string           `json:"criteria,omitzero"`
	Provider map[string]int64 `json:"provider,omitzero"`
	paramObj
}

func (r SendSendMessageParamsMessageTimeout) MarshalJSON() (data []byte, err error) {
	type shadow SendSendMessageParamsMessageTimeout
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendSendMessageParamsMessageTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SendSendMessageParamsMessageTimeout](
		"criteria", "no-escalation", "delivered", "viewed", "engaged",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SendSendMessageParamsMessageToUnion struct {
	OfSendSendMessagesMessageToObject *SendSendMessageParamsMessageToObject `json:",omitzero,inline"`
	OfRecipientArray                  []RecipientParam                      `json:",omitzero,inline"`
	paramUnion
}

func (u SendSendMessageParamsMessageToUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSendSendMessagesMessageToObject, u.OfRecipientArray)
}
func (u *SendSendMessageParamsMessageToUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SendSendMessageParamsMessageToUnion) asAny() any {
	if !param.IsOmitted(u.OfSendSendMessagesMessageToObject) {
		return u.OfSendSendMessagesMessageToObject
	} else if !param.IsOmitted(u.OfRecipientArray) {
		return &u.OfRecipientArray
	}
	return nil
}

type SendSendMessageParamsMessageToObject struct {
	// Use `tenant_id` instead.
	AccountID param.Opt[string] `json:"account_id,omitzero"`
	Email     param.Opt[string] `json:"email,omitzero"`
	// The user's preferred ISO 639-1 language code.
	Locale      param.Opt[string] `json:"locale,omitzero"`
	PhoneNumber param.Opt[string] `json:"phone_number,omitzero"`
	// Tenant id. Will load brand, default preferences and base context data.
	TenantID    param.Opt[string]                               `json:"tenant_id,omitzero"`
	UserID      param.Opt[string]                               `json:"user_id,omitzero"`
	Data        map[string]any                                  `json:"data,omitzero"`
	Preferences SendSendMessageParamsMessageToObjectPreferences `json:"preferences,omitzero"`
	// Context such as tenant_id to send the notification with.
	Context MessageContextParam `json:"context,omitzero"`
	paramObj
}

func (r SendSendMessageParamsMessageToObject) MarshalJSON() (data []byte, err error) {
	type shadow SendSendMessageParamsMessageToObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendSendMessageParamsMessageToObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Notifications is required.
type SendSendMessageParamsMessageToObjectPreferences struct {
	Notifications map[string]SendSendMessageParamsMessageToObjectPreferencesNotification `json:"notifications,omitzero,required"`
	TemplateID    param.Opt[string]                                                      `json:"templateId,omitzero"`
	Categories    map[string]SendSendMessageParamsMessageToObjectPreferencesCategory     `json:"categories,omitzero"`
	paramObj
}

func (r SendSendMessageParamsMessageToObjectPreferences) MarshalJSON() (data []byte, err error) {
	type shadow SendSendMessageParamsMessageToObjectPreferences
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendSendMessageParamsMessageToObjectPreferences) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Status is required.
type SendSendMessageParamsMessageToObjectPreferencesNotification struct {
	// Any of "OPTED_IN", "OPTED_OUT", "REQUIRED".
	Status             string                                                                         `json:"status,omitzero,required"`
	ChannelPreferences []SendSendMessageParamsMessageToObjectPreferencesNotificationChannelPreference `json:"channel_preferences,omitzero"`
	Rules              []SendSendMessageParamsMessageToObjectPreferencesNotificationRule              `json:"rules,omitzero"`
	// Any of "subscription", "list", "recipient".
	Source string `json:"source,omitzero"`
	paramObj
}

func (r SendSendMessageParamsMessageToObjectPreferencesNotification) MarshalJSON() (data []byte, err error) {
	type shadow SendSendMessageParamsMessageToObjectPreferencesNotification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendSendMessageParamsMessageToObjectPreferencesNotification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SendSendMessageParamsMessageToObjectPreferencesNotification](
		"status", "OPTED_IN", "OPTED_OUT", "REQUIRED",
	)
	apijson.RegisterFieldValidator[SendSendMessageParamsMessageToObjectPreferencesNotification](
		"source", "subscription", "list", "recipient",
	)
}

// The property Channel is required.
type SendSendMessageParamsMessageToObjectPreferencesNotificationChannelPreference struct {
	// Any of "direct_message", "email", "push", "sms", "webhook", "inbox".
	Channel string `json:"channel,omitzero,required"`
	paramObj
}

func (r SendSendMessageParamsMessageToObjectPreferencesNotificationChannelPreference) MarshalJSON() (data []byte, err error) {
	type shadow SendSendMessageParamsMessageToObjectPreferencesNotificationChannelPreference
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendSendMessageParamsMessageToObjectPreferencesNotificationChannelPreference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SendSendMessageParamsMessageToObjectPreferencesNotificationChannelPreference](
		"channel", "direct_message", "email", "push", "sms", "webhook", "inbox",
	)
}

// The property Until is required.
type SendSendMessageParamsMessageToObjectPreferencesNotificationRule struct {
	Until string            `json:"until,required"`
	Start param.Opt[string] `json:"start,omitzero"`
	paramObj
}

func (r SendSendMessageParamsMessageToObjectPreferencesNotificationRule) MarshalJSON() (data []byte, err error) {
	type shadow SendSendMessageParamsMessageToObjectPreferencesNotificationRule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendSendMessageParamsMessageToObjectPreferencesNotificationRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Status is required.
type SendSendMessageParamsMessageToObjectPreferencesCategory struct {
	// Any of "OPTED_IN", "OPTED_OUT", "REQUIRED".
	Status             string                                                                     `json:"status,omitzero,required"`
	ChannelPreferences []SendSendMessageParamsMessageToObjectPreferencesCategoryChannelPreference `json:"channel_preferences,omitzero"`
	Rules              []SendSendMessageParamsMessageToObjectPreferencesCategoryRule              `json:"rules,omitzero"`
	// Any of "subscription", "list", "recipient".
	Source string `json:"source,omitzero"`
	paramObj
}

func (r SendSendMessageParamsMessageToObjectPreferencesCategory) MarshalJSON() (data []byte, err error) {
	type shadow SendSendMessageParamsMessageToObjectPreferencesCategory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendSendMessageParamsMessageToObjectPreferencesCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SendSendMessageParamsMessageToObjectPreferencesCategory](
		"status", "OPTED_IN", "OPTED_OUT", "REQUIRED",
	)
	apijson.RegisterFieldValidator[SendSendMessageParamsMessageToObjectPreferencesCategory](
		"source", "subscription", "list", "recipient",
	)
}

// The property Channel is required.
type SendSendMessageParamsMessageToObjectPreferencesCategoryChannelPreference struct {
	// Any of "direct_message", "email", "push", "sms", "webhook", "inbox".
	Channel string `json:"channel,omitzero,required"`
	paramObj
}

func (r SendSendMessageParamsMessageToObjectPreferencesCategoryChannelPreference) MarshalJSON() (data []byte, err error) {
	type shadow SendSendMessageParamsMessageToObjectPreferencesCategoryChannelPreference
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendSendMessageParamsMessageToObjectPreferencesCategoryChannelPreference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SendSendMessageParamsMessageToObjectPreferencesCategoryChannelPreference](
		"channel", "direct_message", "email", "push", "sms", "webhook", "inbox",
	)
}

// The property Until is required.
type SendSendMessageParamsMessageToObjectPreferencesCategoryRule struct {
	Until string            `json:"until,required"`
	Start param.Opt[string] `json:"start,omitzero"`
	paramObj
}

func (r SendSendMessageParamsMessageToObjectPreferencesCategoryRule) MarshalJSON() (data []byte, err error) {
	type shadow SendSendMessageParamsMessageToObjectPreferencesCategoryRule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendSendMessageParamsMessageToObjectPreferencesCategoryRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
