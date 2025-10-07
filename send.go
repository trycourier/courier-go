// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/stainless-sdks/courier-go/internal/apijson"
	"github.com/stainless-sdks/courier-go/internal/requestconfig"
	"github.com/stainless-sdks/courier-go/option"
	"github.com/stainless-sdks/courier-go/packages/param"
	"github.com/stainless-sdks/courier-go/packages/respjson"
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
	Elements []ElementalContentElementUnion `json:"elements"`
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

func ContentParamOfElementalContent(elements []ElementalContentElementUnionParam, version string) ContentUnionParam {
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
	Status             PreferenceStatus                                         `json:"status,omitzero,required"`
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
		"source", "subscription", "list", "recipient",
	)
}

// The property Channel is required.
type RecipientPreferencesNotificationChannelPreferenceParam struct {
	// Any of "direct_message", "email", "push", "sms", "webhook", "inbox".
	Channel ChannelClassification `json:"channel,omitzero,required"`
	paramObj
}

func (r RecipientPreferencesNotificationChannelPreferenceParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientPreferencesNotificationChannelPreferenceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecipientPreferencesNotificationChannelPreferenceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	Status             PreferenceStatus                                     `json:"status,omitzero,required"`
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
		"source", "subscription", "list", "recipient",
	)
}

// The property Channel is required.
type RecipientPreferencesCategoryChannelPreferenceParam struct {
	// Any of "direct_message", "email", "push", "sms", "webhook", "inbox".
	Channel ChannelClassification `json:"channel,omitzero,required"`
	paramObj
}

func (r RecipientPreferencesCategoryChannelPreferenceParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientPreferencesCategoryChannelPreferenceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecipientPreferencesCategoryChannelPreferenceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
type SendSendMessageParamsMessage struct {
	BrandID param.Opt[string] `json:"brand_id,omitzero"`
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
	To SendSendMessageParamsMessageToUnion `json:"to,omitzero"`
	// Describes content that will work for email, inbox, push, chat, or any channel
	// id.
	Content ContentUnionParam   `json:"content,omitzero"`
	Context MessageContextParam `json:"context,omitzero"`
	paramObj
}

func (r SendSendMessageParamsMessage) MarshalJSON() (data []byte, err error) {
	type shadow SendSendMessageParamsMessage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendSendMessageParamsMessage) UnmarshalJSON(data []byte) error {
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
	OfUserRecipient  *UserRecipientParam `json:",omitzero,inline"`
	OfRecipientArray []RecipientParam    `json:",omitzero,inline"`
	paramUnion
}

func (u SendSendMessageParamsMessageToUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfUserRecipient, u.OfRecipientArray)
}
func (u *SendSendMessageParamsMessageToUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SendSendMessageParamsMessageToUnion) asAny() any {
	if !param.IsOmitted(u.OfUserRecipient) {
		return u.OfUserRecipient
	} else if !param.IsOmitted(u.OfRecipientArray) {
		return &u.OfRecipientArray
	}
	return nil
}
