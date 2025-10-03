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
	// Syntactic sugar to provide a fast shorthand for Courier Elemental Blocks.
	Content SendMessageParamsMessageContent `json:"content,omitzero,required"`
	BrandID param.Opt[string]               `json:"brand_id,omitzero"`
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

// Syntactic sugar to provide a fast shorthand for Courier Elemental Blocks.
//
// The properties Body, Title are required.
type SendMessageParamsMessageContent struct {
	// The text content displayed in the notification.
	Body string `json:"body,required"`
	// Title/subject displayed by supported channels.
	Title string `json:"title,required"`
	paramObj
}

func (r SendMessageParamsMessageContent) MarshalJSON() (data []byte, err error) {
	type shadow SendMessageParamsMessageContent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendMessageParamsMessageContent) UnmarshalJSON(data []byte) error {
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
	Utm SendMessageParamsMessageChannelMetadataUtm `json:"utm,omitzero"`
	paramObj
}

func (r SendMessageParamsMessageChannelMetadata) MarshalJSON() (data []byte, err error) {
	type shadow SendMessageParamsMessageChannelMetadata
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendMessageParamsMessageChannelMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SendMessageParamsMessageChannelMetadataUtm struct {
	Campaign param.Opt[string] `json:"campaign,omitzero"`
	Content  param.Opt[string] `json:"content,omitzero"`
	Medium   param.Opt[string] `json:"medium,omitzero"`
	Source   param.Opt[string] `json:"source,omitzero"`
	Term     param.Opt[string] `json:"term,omitzero"`
	paramObj
}

func (r SendMessageParamsMessageChannelMetadataUtm) MarshalJSON() (data []byte, err error) {
	type shadow SendMessageParamsMessageChannelMetadataUtm
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendMessageParamsMessageChannelMetadataUtm) UnmarshalJSON(data []byte) error {
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
	Event   param.Opt[string]                   `json:"event,omitzero"`
	TraceID param.Opt[string]                   `json:"trace_id,omitzero"`
	Tags    []string                            `json:"tags,omitzero"`
	Utm     SendMessageParamsMessageMetadataUtm `json:"utm,omitzero"`
	paramObj
}

func (r SendMessageParamsMessageMetadata) MarshalJSON() (data []byte, err error) {
	type shadow SendMessageParamsMessageMetadata
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendMessageParamsMessageMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SendMessageParamsMessageMetadataUtm struct {
	Campaign param.Opt[string] `json:"campaign,omitzero"`
	Content  param.Opt[string] `json:"content,omitzero"`
	Medium   param.Opt[string] `json:"medium,omitzero"`
	Source   param.Opt[string] `json:"source,omitzero"`
	Term     param.Opt[string] `json:"term,omitzero"`
	paramObj
}

func (r SendMessageParamsMessageMetadataUtm) MarshalJSON() (data []byte, err error) {
	type shadow SendMessageParamsMessageMetadataUtm
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendMessageParamsMessageMetadataUtm) UnmarshalJSON(data []byte) error {
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
	Utm SendMessageParamsMessageProviderMetadataUtm `json:"utm,omitzero"`
	paramObj
}

func (r SendMessageParamsMessageProviderMetadata) MarshalJSON() (data []byte, err error) {
	type shadow SendMessageParamsMessageProviderMetadata
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendMessageParamsMessageProviderMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SendMessageParamsMessageProviderMetadataUtm struct {
	Campaign param.Opt[string] `json:"campaign,omitzero"`
	Content  param.Opt[string] `json:"content,omitzero"`
	Medium   param.Opt[string] `json:"medium,omitzero"`
	Source   param.Opt[string] `json:"source,omitzero"`
	Term     param.Opt[string] `json:"term,omitzero"`
	paramObj
}

func (r SendMessageParamsMessageProviderMetadataUtm) MarshalJSON() (data []byte, err error) {
	type shadow SendMessageParamsMessageProviderMetadataUtm
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendMessageParamsMessageProviderMetadataUtm) UnmarshalJSON(data []byte) error {
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
	Notifications map[string]SendMessageParamsMessageToObjectPreferencesNotification `json:"notifications,omitzero,required"`
	TemplateID    param.Opt[string]                                                  `json:"templateId,omitzero"`
	Categories    map[string]SendMessageParamsMessageToObjectPreferencesCategory     `json:"categories,omitzero"`
	paramObj
}

func (r SendMessageParamsMessageToObjectPreferences) MarshalJSON() (data []byte, err error) {
	type shadow SendMessageParamsMessageToObjectPreferences
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendMessageParamsMessageToObjectPreferences) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Status is required.
type SendMessageParamsMessageToObjectPreferencesNotification struct {
	// Any of "OPTED_IN", "OPTED_OUT", "REQUIRED".
	Status             string                                                                     `json:"status,omitzero,required"`
	ChannelPreferences []SendMessageParamsMessageToObjectPreferencesNotificationChannelPreference `json:"channel_preferences,omitzero"`
	Rules              []SendMessageParamsMessageToObjectPreferencesNotificationRule              `json:"rules,omitzero"`
	// Any of "subscription", "list", "recipient".
	Source string `json:"source,omitzero"`
	paramObj
}

func (r SendMessageParamsMessageToObjectPreferencesNotification) MarshalJSON() (data []byte, err error) {
	type shadow SendMessageParamsMessageToObjectPreferencesNotification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendMessageParamsMessageToObjectPreferencesNotification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SendMessageParamsMessageToObjectPreferencesNotification](
		"status", "OPTED_IN", "OPTED_OUT", "REQUIRED",
	)
	apijson.RegisterFieldValidator[SendMessageParamsMessageToObjectPreferencesNotification](
		"source", "subscription", "list", "recipient",
	)
}

// The property Channel is required.
type SendMessageParamsMessageToObjectPreferencesNotificationChannelPreference struct {
	// Any of "direct_message", "email", "push", "sms", "webhook", "inbox".
	Channel string `json:"channel,omitzero,required"`
	paramObj
}

func (r SendMessageParamsMessageToObjectPreferencesNotificationChannelPreference) MarshalJSON() (data []byte, err error) {
	type shadow SendMessageParamsMessageToObjectPreferencesNotificationChannelPreference
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendMessageParamsMessageToObjectPreferencesNotificationChannelPreference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SendMessageParamsMessageToObjectPreferencesNotificationChannelPreference](
		"channel", "direct_message", "email", "push", "sms", "webhook", "inbox",
	)
}

// The property Until is required.
type SendMessageParamsMessageToObjectPreferencesNotificationRule struct {
	Until string            `json:"until,required"`
	Start param.Opt[string] `json:"start,omitzero"`
	paramObj
}

func (r SendMessageParamsMessageToObjectPreferencesNotificationRule) MarshalJSON() (data []byte, err error) {
	type shadow SendMessageParamsMessageToObjectPreferencesNotificationRule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendMessageParamsMessageToObjectPreferencesNotificationRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Status is required.
type SendMessageParamsMessageToObjectPreferencesCategory struct {
	// Any of "OPTED_IN", "OPTED_OUT", "REQUIRED".
	Status             string                                                                 `json:"status,omitzero,required"`
	ChannelPreferences []SendMessageParamsMessageToObjectPreferencesCategoryChannelPreference `json:"channel_preferences,omitzero"`
	Rules              []SendMessageParamsMessageToObjectPreferencesCategoryRule              `json:"rules,omitzero"`
	// Any of "subscription", "list", "recipient".
	Source string `json:"source,omitzero"`
	paramObj
}

func (r SendMessageParamsMessageToObjectPreferencesCategory) MarshalJSON() (data []byte, err error) {
	type shadow SendMessageParamsMessageToObjectPreferencesCategory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendMessageParamsMessageToObjectPreferencesCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SendMessageParamsMessageToObjectPreferencesCategory](
		"status", "OPTED_IN", "OPTED_OUT", "REQUIRED",
	)
	apijson.RegisterFieldValidator[SendMessageParamsMessageToObjectPreferencesCategory](
		"source", "subscription", "list", "recipient",
	)
}

// The property Channel is required.
type SendMessageParamsMessageToObjectPreferencesCategoryChannelPreference struct {
	// Any of "direct_message", "email", "push", "sms", "webhook", "inbox".
	Channel string `json:"channel,omitzero,required"`
	paramObj
}

func (r SendMessageParamsMessageToObjectPreferencesCategoryChannelPreference) MarshalJSON() (data []byte, err error) {
	type shadow SendMessageParamsMessageToObjectPreferencesCategoryChannelPreference
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendMessageParamsMessageToObjectPreferencesCategoryChannelPreference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SendMessageParamsMessageToObjectPreferencesCategoryChannelPreference](
		"channel", "direct_message", "email", "push", "sms", "webhook", "inbox",
	)
}

// The property Until is required.
type SendMessageParamsMessageToObjectPreferencesCategoryRule struct {
	Until string            `json:"until,required"`
	Start param.Opt[string] `json:"start,omitzero"`
	paramObj
}

func (r SendMessageParamsMessageToObjectPreferencesCategoryRule) MarshalJSON() (data []byte, err error) {
	type shadow SendMessageParamsMessageToObjectPreferencesCategoryRule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendMessageParamsMessageToObjectPreferencesCategoryRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
