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

type BaseMessage struct {
	BrandID string `json:"brand_id,nullable"`
	// "Define run-time configuration for one or more channels. If you don't specify
	// channels, the default configuration for each channel will be used. Valid
	// ChannelId's are: email, sms, push, inbox, direct_message, banner, and webhook."
	Channels map[string]BaseMessageChannel `json:"channels,nullable"`
	// Context to load with this recipient. Will override any context set on
	// message.context.
	Context MessageContext `json:"context,nullable"`
	// An arbitrary object that includes any data you want to pass to the message. The
	// data will populate the corresponding template or elements variables.
	Data map[string]any `json:"data,nullable"`
	// Defines the time to wait before delivering the message. You can specify one of
	// the following options. Duration with the number of milliseconds to delay. Until
	// with an ISO 8601 timestamp that specifies when it should be delivered. Until
	// with an OpenStreetMap opening_hours-like format that specifies the
	// [Delivery Window](https://www.courier.com/docs/platform/sending/failover/#delivery-window)
	// (e.g., 'Mo-Fr 08:00-18:00pm')
	Delay BaseMessageDelay `json:"delay,nullable"`
	// "Expiry allows you to set an absolute or relative time in which a message
	// expires. Note: This is only valid for the Courier Inbox channel as of
	// 12-08-2022."
	Expiry BaseMessageExpiry `json:"expiry,nullable"`
	// Metadata such as utm tracking attached with the notification through this
	// channel.
	Metadata    BaseMessageMetadata    `json:"metadata,nullable"`
	Preferences BaseMessagePreferences `json:"preferences,nullable"`
	// An object whose keys are valid provider identifiers which map to an object.
	Providers map[string]BaseMessageProvider `json:"providers,nullable"`
	// Allows you to customize which channel(s) Courier will potentially deliver the
	// message. If no routing key is specified, Courier will use the default routing
	// configuration or routing defined by the template.
	Routing BaseMessageRouting `json:"routing,nullable"`
	// Time in ms to attempt the channel before failing over to the next available
	// channel.
	Timeout BaseMessageTimeout `json:"timeout,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BrandID     respjson.Field
		Channels    respjson.Field
		Context     respjson.Field
		Data        respjson.Field
		Delay       respjson.Field
		Expiry      respjson.Field
		Metadata    respjson.Field
		Preferences respjson.Field
		Providers   respjson.Field
		Routing     respjson.Field
		Timeout     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaseMessage) RawJSON() string { return r.JSON.raw }
func (r *BaseMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BaseMessage to a BaseMessageParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BaseMessageParam.Overrides()
func (r BaseMessage) ToParam() BaseMessageParam {
	return param.Override[BaseMessageParam](json.RawMessage(r.RawJSON()))
}

type BaseMessageChannel struct {
	// Id of the brand that should be used for rendering the message. If not specified,
	// the brand configured as default brand will be used.
	BrandID string `json:"brand_id,nullable"`
	// A JavaScript conditional expression to determine if the message should be sent
	// through the channel. Has access to the data and profile object. Only applies
	// when a custom routing strategy is defined. For example,
	// `data.name === profile.name`.
	If       string                     `json:"if,nullable"`
	Metadata BaseMessageChannelMetadata `json:"metadata,nullable"`
	// Channel specific overrides.
	Override map[string]any `json:"override,nullable"`
	// A list of providers enabled for this channel. Courier will select one provider
	// to send through unless routing_method is set to all.
	Providers []string `json:"providers,nullable"`
	// The method for selecting the providers to send the message with. Single will
	// send to one of the available providers for this channel, all will send the
	// message through all channels. Defaults to `single`.
	//
	// Any of "all", "single".
	RoutingMethod string                     `json:"routing_method,nullable"`
	Timeouts      BaseMessageChannelTimeouts `json:"timeouts,nullable"`
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
func (r BaseMessageChannel) RawJSON() string { return r.JSON.raw }
func (r *BaseMessageChannel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BaseMessageChannelMetadata struct {
	Utm Utm `json:"utm,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Utm         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaseMessageChannelMetadata) RawJSON() string { return r.JSON.raw }
func (r *BaseMessageChannelMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BaseMessageChannelTimeouts struct {
	Channel  int64 `json:"channel,nullable"`
	Provider int64 `json:"provider,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Channel     respjson.Field
		Provider    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaseMessageChannelTimeouts) RawJSON() string { return r.JSON.raw }
func (r *BaseMessageChannelTimeouts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines the time to wait before delivering the message. You can specify one of
// the following options. Duration with the number of milliseconds to delay. Until
// with an ISO 8601 timestamp that specifies when it should be delivered. Until
// with an OpenStreetMap opening_hours-like format that specifies the
// [Delivery Window](https://www.courier.com/docs/platform/sending/failover/#delivery-window)
// (e.g., 'Mo-Fr 08:00-18:00pm')
type BaseMessageDelay struct {
	// The duration of the delay in milliseconds.
	Duration int64 `json:"duration,nullable"`
	// An ISO 8601 timestamp that specifies when it should be delivered or an
	// OpenStreetMap opening_hours-like format that specifies the
	// [Delivery Window](https://www.courier.com/docs/platform/sending/failover/#delivery-window)
	// (e.g., 'Mo-Fr 08:00-18:00pm')
	Until string `json:"until,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Duration    respjson.Field
		Until       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaseMessageDelay) RawJSON() string { return r.JSON.raw }
func (r *BaseMessageDelay) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// "Expiry allows you to set an absolute or relative time in which a message
// expires. Note: This is only valid for the Courier Inbox channel as of
// 12-08-2022."
type BaseMessageExpiry struct {
	// A duration in the form of milliseconds or an ISO8601 Duration format (i.e.
	// P1DT4H).
	ExpiresIn BaseMessageExpiryExpiresInUnion `json:"expires_in,required"`
	// An epoch timestamp or ISO8601 timestamp with timezone
	// `(YYYY-MM-DDThh:mm:ss.sTZD)` that describes the time in which a message expires.
	ExpiresAt string `json:"expires_at,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresIn   respjson.Field
		ExpiresAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaseMessageExpiry) RawJSON() string { return r.JSON.raw }
func (r *BaseMessageExpiry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BaseMessageExpiryExpiresInUnion contains all possible properties and values from
// [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type BaseMessageExpiryExpiresInUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u BaseMessageExpiryExpiresInUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BaseMessageExpiryExpiresInUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BaseMessageExpiryExpiresInUnion) RawJSON() string { return u.JSON.raw }

func (r *BaseMessageExpiryExpiresInUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata such as utm tracking attached with the notification through this
// channel.
type BaseMessageMetadata struct {
	// An arbitrary string to tracks the event that generated this request (e.g.
	// 'signup').
	Event string `json:"event,nullable"`
	// An array of up to 9 tags you wish to associate with this request (and
	// corresponding messages) for later analysis. Individual tags cannot be more than
	// 30 characters in length.
	Tags []string `json:"tags,nullable"`
	// A unique ID used to correlate this request to processing on your servers. Note:
	// Courier does not verify the uniqueness of this ID.
	TraceID string `json:"trace_id,nullable"`
	// Identify the campaign that refers traffic to a specific website, and attributes
	// the browser's website session.
	Utm Utm `json:"utm,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Event       respjson.Field
		Tags        respjson.Field
		TraceID     respjson.Field
		Utm         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaseMessageMetadata) RawJSON() string { return r.JSON.raw }
func (r *BaseMessageMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BaseMessagePreferences struct {
	// The ID of the subscription topic you want to apply to the message. If this is a
	// templated message, it will override the subscription topic if already associated
	SubscriptionTopicID string `json:"subscription_topic_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SubscriptionTopicID respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaseMessagePreferences) RawJSON() string { return r.JSON.raw }
func (r *BaseMessagePreferences) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BaseMessageProvider struct {
	// A JavaScript conditional expression to determine if the message should be sent
	// through the provider. Has access to the data and profile object. Only applies
	// when a custom routing strategy is defined. For example,
	// `data.name === profile.name`
	If       string                      `json:"if,nullable"`
	Metadata BaseMessageProviderMetadata `json:"metadata,nullable"`
	// Provider specific overrides.
	Override map[string]any `json:"override,nullable"`
	Timeouts int64          `json:"timeouts,nullable"`
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
func (r BaseMessageProvider) RawJSON() string { return r.JSON.raw }
func (r *BaseMessageProvider) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BaseMessageProviderMetadata struct {
	Utm Utm `json:"utm,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Utm         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaseMessageProviderMetadata) RawJSON() string { return r.JSON.raw }
func (r *BaseMessageProviderMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows you to customize which channel(s) Courier will potentially deliver the
// message. If no routing key is specified, Courier will use the default routing
// configuration or routing defined by the template.
type BaseMessageRouting struct {
	// A list of channels or providers to send the message through. Can also
	// recursively define sub-routing methods, which can be useful for defining
	// advanced push notification delivery strategies.
	Channels []MessageRoutingChannelUnion `json:"channels,required"`
	// Any of "all", "single".
	Method string `json:"method,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Channels    respjson.Field
		Method      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaseMessageRouting) RawJSON() string { return r.JSON.raw }
func (r *BaseMessageRouting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Time in ms to attempt the channel before failing over to the next available
// channel.
type BaseMessageTimeout struct {
	Channel map[string]int64 `json:"channel,nullable"`
	// Any of "no-escalation", "delivered", "viewed", "engaged".
	Criteria   string           `json:"criteria,nullable"`
	Escalation int64            `json:"escalation,nullable"`
	Message    int64            `json:"message,nullable"`
	Provider   map[string]int64 `json:"provider,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Channel     respjson.Field
		Criteria    respjson.Field
		Escalation  respjson.Field
		Message     respjson.Field
		Provider    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaseMessageTimeout) RawJSON() string { return r.JSON.raw }
func (r *BaseMessageTimeout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BaseMessageParam struct {
	BrandID param.Opt[string] `json:"brand_id,omitzero"`
	// "Define run-time configuration for one or more channels. If you don't specify
	// channels, the default configuration for each channel will be used. Valid
	// ChannelId's are: email, sms, push, inbox, direct_message, banner, and webhook."
	Channels map[string]BaseMessageChannelParam `json:"channels,omitzero"`
	// An arbitrary object that includes any data you want to pass to the message. The
	// data will populate the corresponding template or elements variables.
	Data map[string]any `json:"data,omitzero"`
	// Defines the time to wait before delivering the message. You can specify one of
	// the following options. Duration with the number of milliseconds to delay. Until
	// with an ISO 8601 timestamp that specifies when it should be delivered. Until
	// with an OpenStreetMap opening_hours-like format that specifies the
	// [Delivery Window](https://www.courier.com/docs/platform/sending/failover/#delivery-window)
	// (e.g., 'Mo-Fr 08:00-18:00pm')
	Delay BaseMessageDelayParam `json:"delay,omitzero"`
	// "Expiry allows you to set an absolute or relative time in which a message
	// expires. Note: This is only valid for the Courier Inbox channel as of
	// 12-08-2022."
	Expiry BaseMessageExpiryParam `json:"expiry,omitzero"`
	// Metadata such as utm tracking attached with the notification through this
	// channel.
	Metadata    BaseMessageMetadataParam    `json:"metadata,omitzero"`
	Preferences BaseMessagePreferencesParam `json:"preferences,omitzero"`
	// An object whose keys are valid provider identifiers which map to an object.
	Providers map[string]BaseMessageProviderParam `json:"providers,omitzero"`
	// Allows you to customize which channel(s) Courier will potentially deliver the
	// message. If no routing key is specified, Courier will use the default routing
	// configuration or routing defined by the template.
	Routing BaseMessageRoutingParam `json:"routing,omitzero"`
	// Time in ms to attempt the channel before failing over to the next available
	// channel.
	Timeout BaseMessageTimeoutParam `json:"timeout,omitzero"`
	// Context to load with this recipient. Will override any context set on
	// message.context.
	Context MessageContextParam `json:"context,omitzero"`
	paramObj
}

func (r BaseMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow BaseMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaseMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BaseMessageChannelParam struct {
	// Id of the brand that should be used for rendering the message. If not specified,
	// the brand configured as default brand will be used.
	BrandID param.Opt[string] `json:"brand_id,omitzero"`
	// A JavaScript conditional expression to determine if the message should be sent
	// through the channel. Has access to the data and profile object. Only applies
	// when a custom routing strategy is defined. For example,
	// `data.name === profile.name`.
	If       param.Opt[string]               `json:"if,omitzero"`
	Metadata BaseMessageChannelMetadataParam `json:"metadata,omitzero"`
	// Channel specific overrides.
	Override map[string]any `json:"override,omitzero"`
	// A list of providers enabled for this channel. Courier will select one provider
	// to send through unless routing_method is set to all.
	Providers []string `json:"providers,omitzero"`
	// The method for selecting the providers to send the message with. Single will
	// send to one of the available providers for this channel, all will send the
	// message through all channels. Defaults to `single`.
	//
	// Any of "all", "single".
	RoutingMethod string                          `json:"routing_method,omitzero"`
	Timeouts      BaseMessageChannelTimeoutsParam `json:"timeouts,omitzero"`
	paramObj
}

func (r BaseMessageChannelParam) MarshalJSON() (data []byte, err error) {
	type shadow BaseMessageChannelParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaseMessageChannelParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BaseMessageChannelParam](
		"routing_method", "all", "single",
	)
}

type BaseMessageChannelMetadataParam struct {
	Utm UtmParam `json:"utm,omitzero"`
	paramObj
}

func (r BaseMessageChannelMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow BaseMessageChannelMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaseMessageChannelMetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BaseMessageChannelTimeoutsParam struct {
	Channel  param.Opt[int64] `json:"channel,omitzero"`
	Provider param.Opt[int64] `json:"provider,omitzero"`
	paramObj
}

func (r BaseMessageChannelTimeoutsParam) MarshalJSON() (data []byte, err error) {
	type shadow BaseMessageChannelTimeoutsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaseMessageChannelTimeoutsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines the time to wait before delivering the message. You can specify one of
// the following options. Duration with the number of milliseconds to delay. Until
// with an ISO 8601 timestamp that specifies when it should be delivered. Until
// with an OpenStreetMap opening_hours-like format that specifies the
// [Delivery Window](https://www.courier.com/docs/platform/sending/failover/#delivery-window)
// (e.g., 'Mo-Fr 08:00-18:00pm')
type BaseMessageDelayParam struct {
	// The duration of the delay in milliseconds.
	Duration param.Opt[int64] `json:"duration,omitzero"`
	// An ISO 8601 timestamp that specifies when it should be delivered or an
	// OpenStreetMap opening_hours-like format that specifies the
	// [Delivery Window](https://www.courier.com/docs/platform/sending/failover/#delivery-window)
	// (e.g., 'Mo-Fr 08:00-18:00pm')
	Until param.Opt[string] `json:"until,omitzero"`
	paramObj
}

func (r BaseMessageDelayParam) MarshalJSON() (data []byte, err error) {
	type shadow BaseMessageDelayParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaseMessageDelayParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// "Expiry allows you to set an absolute or relative time in which a message
// expires. Note: This is only valid for the Courier Inbox channel as of
// 12-08-2022."
//
// The property ExpiresIn is required.
type BaseMessageExpiryParam struct {
	// A duration in the form of milliseconds or an ISO8601 Duration format (i.e.
	// P1DT4H).
	ExpiresIn BaseMessageExpiryExpiresInUnionParam `json:"expires_in,omitzero,required"`
	// An epoch timestamp or ISO8601 timestamp with timezone
	// `(YYYY-MM-DDThh:mm:ss.sTZD)` that describes the time in which a message expires.
	ExpiresAt param.Opt[string] `json:"expires_at,omitzero"`
	paramObj
}

func (r BaseMessageExpiryParam) MarshalJSON() (data []byte, err error) {
	type shadow BaseMessageExpiryParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaseMessageExpiryParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BaseMessageExpiryExpiresInUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u BaseMessageExpiryExpiresInUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *BaseMessageExpiryExpiresInUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BaseMessageExpiryExpiresInUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// Metadata such as utm tracking attached with the notification through this
// channel.
type BaseMessageMetadataParam struct {
	// An arbitrary string to tracks the event that generated this request (e.g.
	// 'signup').
	Event param.Opt[string] `json:"event,omitzero"`
	// A unique ID used to correlate this request to processing on your servers. Note:
	// Courier does not verify the uniqueness of this ID.
	TraceID param.Opt[string] `json:"trace_id,omitzero"`
	// An array of up to 9 tags you wish to associate with this request (and
	// corresponding messages) for later analysis. Individual tags cannot be more than
	// 30 characters in length.
	Tags []string `json:"tags,omitzero"`
	// Identify the campaign that refers traffic to a specific website, and attributes
	// the browser's website session.
	Utm UtmParam `json:"utm,omitzero"`
	paramObj
}

func (r BaseMessageMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow BaseMessageMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaseMessageMetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SubscriptionTopicID is required.
type BaseMessagePreferencesParam struct {
	// The ID of the subscription topic you want to apply to the message. If this is a
	// templated message, it will override the subscription topic if already associated
	SubscriptionTopicID string `json:"subscription_topic_id,required"`
	paramObj
}

func (r BaseMessagePreferencesParam) MarshalJSON() (data []byte, err error) {
	type shadow BaseMessagePreferencesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaseMessagePreferencesParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BaseMessageProviderParam struct {
	// A JavaScript conditional expression to determine if the message should be sent
	// through the provider. Has access to the data and profile object. Only applies
	// when a custom routing strategy is defined. For example,
	// `data.name === profile.name`
	If       param.Opt[string]                `json:"if,omitzero"`
	Timeouts param.Opt[int64]                 `json:"timeouts,omitzero"`
	Metadata BaseMessageProviderMetadataParam `json:"metadata,omitzero"`
	// Provider specific overrides.
	Override map[string]any `json:"override,omitzero"`
	paramObj
}

func (r BaseMessageProviderParam) MarshalJSON() (data []byte, err error) {
	type shadow BaseMessageProviderParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaseMessageProviderParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BaseMessageProviderMetadataParam struct {
	Utm UtmParam `json:"utm,omitzero"`
	paramObj
}

func (r BaseMessageProviderMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow BaseMessageProviderMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaseMessageProviderMetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows you to customize which channel(s) Courier will potentially deliver the
// message. If no routing key is specified, Courier will use the default routing
// configuration or routing defined by the template.
//
// The properties Channels, Method are required.
type BaseMessageRoutingParam struct {
	// A list of channels or providers to send the message through. Can also
	// recursively define sub-routing methods, which can be useful for defining
	// advanced push notification delivery strategies.
	Channels []MessageRoutingChannelUnionParam `json:"channels,omitzero,required"`
	// Any of "all", "single".
	Method string `json:"method,omitzero,required"`
	paramObj
}

func (r BaseMessageRoutingParam) MarshalJSON() (data []byte, err error) {
	type shadow BaseMessageRoutingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaseMessageRoutingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BaseMessageRoutingParam](
		"method", "all", "single",
	)
}

// Time in ms to attempt the channel before failing over to the next available
// channel.
type BaseMessageTimeoutParam struct {
	Escalation param.Opt[int64] `json:"escalation,omitzero"`
	Message    param.Opt[int64] `json:"message,omitzero"`
	Channel    map[string]int64 `json:"channel,omitzero"`
	// Any of "no-escalation", "delivered", "viewed", "engaged".
	Criteria string           `json:"criteria,omitzero"`
	Provider map[string]int64 `json:"provider,omitzero"`
	paramObj
}

func (r BaseMessageTimeoutParam) MarshalJSON() (data []byte, err error) {
	type shadow BaseMessageTimeoutParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaseMessageTimeoutParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BaseMessageTimeoutParam](
		"criteria", "no-escalation", "delivered", "viewed", "engaged",
	)
}

type BaseMessageSendToParam struct {
	// The recipient or a list of recipients of the message
	To BaseMessageSendToToUnionParam `json:"to,omitzero"`
	paramObj
}

func (r BaseMessageSendToParam) MarshalJSON() (data []byte, err error) {
	type shadow BaseMessageSendToParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaseMessageSendToParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BaseMessageSendToToUnionParam struct {
	OfAudienceRecipient         *BaseMessageSendToToAudienceRecipientParam  `json:",omitzero,inline"`
	OfBaseMessageSendToToObject *BaseMessageSendToToObjectParam             `json:",omitzero,inline"`
	OfVariant2                  *BaseMessageSendToToObjectParam             `json:",omitzero,inline"`
	OfUserRecipient             *UserRecipientParam                         `json:",omitzero,inline"`
	OfSlackRecipient            *BaseMessageSendToToSlackRecipientParam     `json:",omitzero,inline"`
	OfMsTeamsRecipient          *BaseMessageSendToToMsTeamsRecipientParam   `json:",omitzero,inline"`
	OfRecipientData             map[string]any                              `json:",omitzero,inline"`
	OfPagerdutyRecipient        *BaseMessageSendToToPagerdutyRecipientParam `json:",omitzero,inline"`
	OfWebhookRecipient          *BaseMessageSendToToWebhookRecipientParam   `json:",omitzero,inline"`
	OfRecipientArray            []RecipientUnionParam                       `json:",omitzero,inline"`
	paramUnion
}

func (u BaseMessageSendToToUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAudienceRecipient,
		u.OfBaseMessageSendToToObject,
		u.OfVariant2,
		u.OfUserRecipient,
		u.OfSlackRecipient,
		u.OfMsTeamsRecipient,
		u.OfRecipientData,
		u.OfPagerdutyRecipient,
		u.OfWebhookRecipient,
		u.OfRecipientArray)
}
func (u *BaseMessageSendToToUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BaseMessageSendToToUnionParam) asAny() any {
	if !param.IsOmitted(u.OfAudienceRecipient) {
		return u.OfAudienceRecipient
	} else if !param.IsOmitted(u.OfBaseMessageSendToToObject) {
		return u.OfBaseMessageSendToToObject
	} else if !param.IsOmitted(u.OfVariant2) {
		return u.OfVariant2
	} else if !param.IsOmitted(u.OfUserRecipient) {
		return u.OfUserRecipient
	} else if !param.IsOmitted(u.OfSlackRecipient) {
		return u.OfSlackRecipient
	} else if !param.IsOmitted(u.OfMsTeamsRecipient) {
		return u.OfMsTeamsRecipient
	} else if !param.IsOmitted(u.OfRecipientData) {
		return &u.OfRecipientData
	} else if !param.IsOmitted(u.OfPagerdutyRecipient) {
		return u.OfPagerdutyRecipient
	} else if !param.IsOmitted(u.OfWebhookRecipient) {
		return u.OfWebhookRecipient
	} else if !param.IsOmitted(u.OfRecipientArray) {
		return &u.OfRecipientArray
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaseMessageSendToToUnionParam) GetAudienceID() *string {
	if vt := u.OfAudienceRecipient; vt != nil {
		return &vt.AudienceID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaseMessageSendToToUnionParam) GetListID() *string {
	if vt := u.OfBaseMessageSendToToObject; vt != nil && vt.ListID.Valid() {
		return &vt.ListID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaseMessageSendToToUnionParam) GetListPattern() *string {
	if vt := u.OfVariant2; vt != nil && vt.ListPattern.Valid() {
		return &vt.ListPattern.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaseMessageSendToToUnionParam) GetAccountID() *string {
	if vt := u.OfUserRecipient; vt != nil && vt.AccountID.Valid() {
		return &vt.AccountID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaseMessageSendToToUnionParam) GetContext() *MessageContextParam {
	if vt := u.OfUserRecipient; vt != nil {
		return &vt.Context
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaseMessageSendToToUnionParam) GetEmail() *string {
	if vt := u.OfUserRecipient; vt != nil && vt.Email.Valid() {
		return &vt.Email.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaseMessageSendToToUnionParam) GetLocale() *string {
	if vt := u.OfUserRecipient; vt != nil && vt.Locale.Valid() {
		return &vt.Locale.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaseMessageSendToToUnionParam) GetPhoneNumber() *string {
	if vt := u.OfUserRecipient; vt != nil && vt.PhoneNumber.Valid() {
		return &vt.PhoneNumber.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaseMessageSendToToUnionParam) GetPreferences() *UserRecipientPreferencesParam {
	if vt := u.OfUserRecipient; vt != nil {
		return &vt.Preferences
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaseMessageSendToToUnionParam) GetTenantID() *string {
	if vt := u.OfUserRecipient; vt != nil && vt.TenantID.Valid() {
		return &vt.TenantID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaseMessageSendToToUnionParam) GetUserID() *string {
	if vt := u.OfUserRecipient; vt != nil && vt.UserID.Valid() {
		return &vt.UserID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaseMessageSendToToUnionParam) GetSlack() *BaseMessageSendToToSlackRecipientSlackUnionParam {
	if vt := u.OfSlackRecipient; vt != nil {
		return &vt.Slack
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaseMessageSendToToUnionParam) GetMsTeams() *BaseMessageSendToToMsTeamsRecipientMsTeamsUnionParam {
	if vt := u.OfMsTeamsRecipient; vt != nil {
		return &vt.MsTeams
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaseMessageSendToToUnionParam) GetPagerduty() *BaseMessageSendToToPagerdutyRecipientPagerdutyParam {
	if vt := u.OfPagerdutyRecipient; vt != nil {
		return &vt.Pagerduty
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaseMessageSendToToUnionParam) GetWebhook() *BaseMessageSendToToWebhookRecipientWebhookParam {
	if vt := u.OfWebhookRecipient; vt != nil {
		return &vt.Webhook
	}
	return nil
}

// Returns a pointer to the underlying variant's Data property, if present.
func (u BaseMessageSendToToUnionParam) GetData() map[string]any {
	if vt := u.OfAudienceRecipient; vt != nil {
		return vt.Data
	} else if vt := u.OfBaseMessageSendToToObject; vt != nil {
		return vt.Data
	} else if vt := u.OfVariant2; vt != nil {
		return vt.Data
	} else if vt := u.OfUserRecipient; vt != nil {
		return vt.Data
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u BaseMessageSendToToUnionParam) GetFilters() (res baseMessageSendToToUnionParamFilters) {
	if vt := u.OfAudienceRecipient; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfBaseMessageSendToToObject; vt != nil {
		res.any = &vt.Filters
	}
	return
}

// Can have the runtime types [_[]BaseMessageSendToToAudienceRecipientFilterParam],
// [_[]BaseMessageSendToToObjectFilterParam]
type baseMessageSendToToUnionParamFilters struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]courier.BaseMessageSendToToAudienceRecipientFilterParam:
//	case *[]courier.BaseMessageSendToToObjectFilterParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u baseMessageSendToToUnionParamFilters) AsAny() any { return u.any }

// The property AudienceID is required.
type BaseMessageSendToToAudienceRecipientParam struct {
	// A unique identifier associated with an Audience. A message will be sent to each
	// user in the audience.
	AudienceID string                                            `json:"audience_id,required"`
	Data       map[string]any                                    `json:"data,omitzero"`
	Filters    []BaseMessageSendToToAudienceRecipientFilterParam `json:"filters,omitzero"`
	paramObj
}

func (r BaseMessageSendToToAudienceRecipientParam) MarshalJSON() (data []byte, err error) {
	type shadow BaseMessageSendToToAudienceRecipientParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaseMessageSendToToAudienceRecipientParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Operator, Path, Value are required.
type BaseMessageSendToToAudienceRecipientFilterParam struct {
	// Send to users only if they are member of the account
	//
	// Any of "MEMBER_OF".
	Operator string `json:"operator,omitzero,required"`
	// Any of "account_id".
	Path  string `json:"path,omitzero,required"`
	Value string `json:"value,required"`
	paramObj
}

func (r BaseMessageSendToToAudienceRecipientFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow BaseMessageSendToToAudienceRecipientFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaseMessageSendToToAudienceRecipientFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BaseMessageSendToToAudienceRecipientFilterParam](
		"operator", "MEMBER_OF",
	)
	apijson.RegisterFieldValidator[BaseMessageSendToToAudienceRecipientFilterParam](
		"path", "account_id",
	)
}

type BaseMessageSendToToObjectParam struct {
	ListID  param.Opt[string]                      `json:"list_id,omitzero"`
	Data    map[string]any                         `json:"data,omitzero"`
	Filters []BaseMessageSendToToObjectFilterParam `json:"filters,omitzero"`
	paramObj
}

func (r BaseMessageSendToToObjectParam) MarshalJSON() (data []byte, err error) {
	type shadow BaseMessageSendToToObjectParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaseMessageSendToToObjectParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Operator, Path, Value are required.
type BaseMessageSendToToObjectFilterParam struct {
	// Send to users only if they are member of the account
	//
	// Any of "MEMBER_OF".
	Operator string `json:"operator,omitzero,required"`
	// Any of "account_id".
	Path  string `json:"path,omitzero,required"`
	Value string `json:"value,required"`
	paramObj
}

func (r BaseMessageSendToToObjectFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow BaseMessageSendToToObjectFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaseMessageSendToToObjectFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BaseMessageSendToToObjectFilterParam](
		"operator", "MEMBER_OF",
	)
	apijson.RegisterFieldValidator[BaseMessageSendToToObjectFilterParam](
		"path", "account_id",
	)
}

// The property Slack is required.
type BaseMessageSendToToSlackRecipientParam struct {
	Slack BaseMessageSendToToSlackRecipientSlackUnionParam `json:"slack,omitzero,required"`
	paramObj
}

func (r BaseMessageSendToToSlackRecipientParam) MarshalJSON() (data []byte, err error) {
	type shadow BaseMessageSendToToSlackRecipientParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaseMessageSendToToSlackRecipientParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BaseMessageSendToToSlackRecipientSlackUnionParam struct {
	OfSendToSlackChannel *BaseMessageSendToToSlackRecipientSlackSendToSlackChannelParam `json:",omitzero,inline"`
	OfSendToSlackEmail   *BaseMessageSendToToSlackRecipientSlackSendToSlackEmailParam   `json:",omitzero,inline"`
	OfSendToSlackUserID  *BaseMessageSendToToSlackRecipientSlackSendToSlackUserIDParam  `json:",omitzero,inline"`
	paramUnion
}

func (u BaseMessageSendToToSlackRecipientSlackUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSendToSlackChannel, u.OfSendToSlackEmail, u.OfSendToSlackUserID)
}
func (u *BaseMessageSendToToSlackRecipientSlackUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BaseMessageSendToToSlackRecipientSlackUnionParam) asAny() any {
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
func (u BaseMessageSendToToSlackRecipientSlackUnionParam) GetChannel() *string {
	if vt := u.OfSendToSlackChannel; vt != nil {
		return &vt.Channel
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaseMessageSendToToSlackRecipientSlackUnionParam) GetEmail() *string {
	if vt := u.OfSendToSlackEmail; vt != nil {
		return &vt.Email
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaseMessageSendToToSlackRecipientSlackUnionParam) GetUserID() *string {
	if vt := u.OfSendToSlackUserID; vt != nil {
		return &vt.UserID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaseMessageSendToToSlackRecipientSlackUnionParam) GetAccessToken() *string {
	if vt := u.OfSendToSlackChannel; vt != nil {
		return (*string)(&vt.AccessToken)
	} else if vt := u.OfSendToSlackEmail; vt != nil {
		return (*string)(&vt.AccessToken)
	} else if vt := u.OfSendToSlackUserID; vt != nil {
		return (*string)(&vt.AccessToken)
	}
	return nil
}

type BaseMessageSendToToSlackRecipientSlackSendToSlackChannelParam struct {
	Channel string `json:"channel,required"`
	SlackBasePropertiesParam
}

func (r BaseMessageSendToToSlackRecipientSlackSendToSlackChannelParam) MarshalJSON() (data []byte, err error) {
	type shadow BaseMessageSendToToSlackRecipientSlackSendToSlackChannelParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type BaseMessageSendToToSlackRecipientSlackSendToSlackEmailParam struct {
	Email string `json:"email,required"`
	SlackBasePropertiesParam
}

func (r BaseMessageSendToToSlackRecipientSlackSendToSlackEmailParam) MarshalJSON() (data []byte, err error) {
	type shadow BaseMessageSendToToSlackRecipientSlackSendToSlackEmailParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type BaseMessageSendToToSlackRecipientSlackSendToSlackUserIDParam struct {
	UserID string `json:"user_id,required"`
	SlackBasePropertiesParam
}

func (r BaseMessageSendToToSlackRecipientSlackSendToSlackUserIDParam) MarshalJSON() (data []byte, err error) {
	type shadow BaseMessageSendToToSlackRecipientSlackSendToSlackUserIDParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// The property MsTeams is required.
type BaseMessageSendToToMsTeamsRecipientParam struct {
	MsTeams BaseMessageSendToToMsTeamsRecipientMsTeamsUnionParam `json:"ms_teams,omitzero,required"`
	paramObj
}

func (r BaseMessageSendToToMsTeamsRecipientParam) MarshalJSON() (data []byte, err error) {
	type shadow BaseMessageSendToToMsTeamsRecipientParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaseMessageSendToToMsTeamsRecipientParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BaseMessageSendToToMsTeamsRecipientMsTeamsUnionParam struct {
	OfSendToMsTeamsUserID         *BaseMessageSendToToMsTeamsRecipientMsTeamsSendToMsTeamsUserIDParam         `json:",omitzero,inline"`
	OfSendToMsTeamsEmail          *BaseMessageSendToToMsTeamsRecipientMsTeamsSendToMsTeamsEmailParam          `json:",omitzero,inline"`
	OfSendToMsTeamsChannelID      *BaseMessageSendToToMsTeamsRecipientMsTeamsSendToMsTeamsChannelIDParam      `json:",omitzero,inline"`
	OfSendToMsTeamsConversationID *BaseMessageSendToToMsTeamsRecipientMsTeamsSendToMsTeamsConversationIDParam `json:",omitzero,inline"`
	OfSendToMsTeamsChannelName    *BaseMessageSendToToMsTeamsRecipientMsTeamsSendToMsTeamsChannelNameParam    `json:",omitzero,inline"`
	paramUnion
}

func (u BaseMessageSendToToMsTeamsRecipientMsTeamsUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSendToMsTeamsUserID,
		u.OfSendToMsTeamsEmail,
		u.OfSendToMsTeamsChannelID,
		u.OfSendToMsTeamsConversationID,
		u.OfSendToMsTeamsChannelName)
}
func (u *BaseMessageSendToToMsTeamsRecipientMsTeamsUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BaseMessageSendToToMsTeamsRecipientMsTeamsUnionParam) asAny() any {
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
func (u BaseMessageSendToToMsTeamsRecipientMsTeamsUnionParam) GetUserID() *string {
	if vt := u.OfSendToMsTeamsUserID; vt != nil {
		return &vt.UserID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaseMessageSendToToMsTeamsRecipientMsTeamsUnionParam) GetEmail() *string {
	if vt := u.OfSendToMsTeamsEmail; vt != nil {
		return &vt.Email
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaseMessageSendToToMsTeamsRecipientMsTeamsUnionParam) GetChannelID() *string {
	if vt := u.OfSendToMsTeamsChannelID; vt != nil {
		return &vt.ChannelID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaseMessageSendToToMsTeamsRecipientMsTeamsUnionParam) GetConversationID() *string {
	if vt := u.OfSendToMsTeamsConversationID; vt != nil {
		return &vt.ConversationID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaseMessageSendToToMsTeamsRecipientMsTeamsUnionParam) GetChannelName() *string {
	if vt := u.OfSendToMsTeamsChannelName; vt != nil {
		return &vt.ChannelName
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaseMessageSendToToMsTeamsRecipientMsTeamsUnionParam) GetTeamID() *string {
	if vt := u.OfSendToMsTeamsChannelName; vt != nil {
		return &vt.TeamID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaseMessageSendToToMsTeamsRecipientMsTeamsUnionParam) GetServiceURL() *string {
	if vt := u.OfSendToMsTeamsUserID; vt != nil {
		return (*string)(&vt.ServiceURL)
	} else if vt := u.OfSendToMsTeamsEmail; vt != nil {
		return (*string)(&vt.ServiceURL)
	} else if vt := u.OfSendToMsTeamsChannelID; vt != nil {
		return (*string)(&vt.ServiceURL)
	} else if vt := u.OfSendToMsTeamsConversationID; vt != nil {
		return (*string)(&vt.ServiceURL)
	} else if vt := u.OfSendToMsTeamsChannelName; vt != nil {
		return (*string)(&vt.ServiceURL)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaseMessageSendToToMsTeamsRecipientMsTeamsUnionParam) GetTenantID() *string {
	if vt := u.OfSendToMsTeamsUserID; vt != nil {
		return (*string)(&vt.TenantID)
	} else if vt := u.OfSendToMsTeamsEmail; vt != nil {
		return (*string)(&vt.TenantID)
	} else if vt := u.OfSendToMsTeamsChannelID; vt != nil {
		return (*string)(&vt.TenantID)
	} else if vt := u.OfSendToMsTeamsConversationID; vt != nil {
		return (*string)(&vt.TenantID)
	} else if vt := u.OfSendToMsTeamsChannelName; vt != nil {
		return (*string)(&vt.TenantID)
	}
	return nil
}

type BaseMessageSendToToMsTeamsRecipientMsTeamsSendToMsTeamsUserIDParam struct {
	UserID string `json:"user_id,required"`
	MsTeamsBasePropertiesParam
}

func (r BaseMessageSendToToMsTeamsRecipientMsTeamsSendToMsTeamsUserIDParam) MarshalJSON() (data []byte, err error) {
	type shadow BaseMessageSendToToMsTeamsRecipientMsTeamsSendToMsTeamsUserIDParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type BaseMessageSendToToMsTeamsRecipientMsTeamsSendToMsTeamsEmailParam struct {
	Email string `json:"email,required"`
	MsTeamsBasePropertiesParam
}

func (r BaseMessageSendToToMsTeamsRecipientMsTeamsSendToMsTeamsEmailParam) MarshalJSON() (data []byte, err error) {
	type shadow BaseMessageSendToToMsTeamsRecipientMsTeamsSendToMsTeamsEmailParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type BaseMessageSendToToMsTeamsRecipientMsTeamsSendToMsTeamsChannelIDParam struct {
	ChannelID string `json:"channel_id,required"`
	MsTeamsBasePropertiesParam
}

func (r BaseMessageSendToToMsTeamsRecipientMsTeamsSendToMsTeamsChannelIDParam) MarshalJSON() (data []byte, err error) {
	type shadow BaseMessageSendToToMsTeamsRecipientMsTeamsSendToMsTeamsChannelIDParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type BaseMessageSendToToMsTeamsRecipientMsTeamsSendToMsTeamsConversationIDParam struct {
	ConversationID string `json:"conversation_id,required"`
	MsTeamsBasePropertiesParam
}

func (r BaseMessageSendToToMsTeamsRecipientMsTeamsSendToMsTeamsConversationIDParam) MarshalJSON() (data []byte, err error) {
	type shadow BaseMessageSendToToMsTeamsRecipientMsTeamsSendToMsTeamsConversationIDParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type BaseMessageSendToToMsTeamsRecipientMsTeamsSendToMsTeamsChannelNameParam struct {
	ChannelName string `json:"channel_name,required"`
	TeamID      string `json:"team_id,required"`
	MsTeamsBasePropertiesParam
}

func (r BaseMessageSendToToMsTeamsRecipientMsTeamsSendToMsTeamsChannelNameParam) MarshalJSON() (data []byte, err error) {
	type shadow BaseMessageSendToToMsTeamsRecipientMsTeamsSendToMsTeamsChannelNameParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// The property Pagerduty is required.
type BaseMessageSendToToPagerdutyRecipientParam struct {
	Pagerduty BaseMessageSendToToPagerdutyRecipientPagerdutyParam `json:"pagerduty,omitzero,required"`
	paramObj
}

func (r BaseMessageSendToToPagerdutyRecipientParam) MarshalJSON() (data []byte, err error) {
	type shadow BaseMessageSendToToPagerdutyRecipientParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaseMessageSendToToPagerdutyRecipientParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BaseMessageSendToToPagerdutyRecipientPagerdutyParam struct {
	EventAction param.Opt[string] `json:"event_action,omitzero"`
	RoutingKey  param.Opt[string] `json:"routing_key,omitzero"`
	Severity    param.Opt[string] `json:"severity,omitzero"`
	Source      param.Opt[string] `json:"source,omitzero"`
	paramObj
}

func (r BaseMessageSendToToPagerdutyRecipientPagerdutyParam) MarshalJSON() (data []byte, err error) {
	type shadow BaseMessageSendToToPagerdutyRecipientPagerdutyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaseMessageSendToToPagerdutyRecipientPagerdutyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Webhook is required.
type BaseMessageSendToToWebhookRecipientParam struct {
	Webhook BaseMessageSendToToWebhookRecipientWebhookParam `json:"webhook,omitzero,required"`
	paramObj
}

func (r BaseMessageSendToToWebhookRecipientParam) MarshalJSON() (data []byte, err error) {
	type shadow BaseMessageSendToToWebhookRecipientParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaseMessageSendToToWebhookRecipientParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property URL is required.
type BaseMessageSendToToWebhookRecipientWebhookParam struct {
	// The URL to send the webhook request to.
	URL string `json:"url,required"`
	// Authentication configuration for the webhook request.
	Authentication BaseMessageSendToToWebhookRecipientWebhookAuthenticationParam `json:"authentication,omitzero"`
	// Custom headers to include in the webhook request.
	Headers map[string]string `json:"headers,omitzero"`
	// The HTTP method to use for the webhook request. Defaults to POST if not
	// specified.
	//
	// Any of "POST", "PUT".
	Method string `json:"method,omitzero"`
	// Specifies what profile information is included in the request payload. Defaults
	// to 'limited' if not specified.
	//
	// Any of "limited", "expanded".
	Profile string `json:"profile,omitzero"`
	paramObj
}

func (r BaseMessageSendToToWebhookRecipientWebhookParam) MarshalJSON() (data []byte, err error) {
	type shadow BaseMessageSendToToWebhookRecipientWebhookParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaseMessageSendToToWebhookRecipientWebhookParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BaseMessageSendToToWebhookRecipientWebhookParam](
		"method", "POST", "PUT",
	)
	apijson.RegisterFieldValidator[BaseMessageSendToToWebhookRecipientWebhookParam](
		"profile", "limited", "expanded",
	)
}

// Authentication configuration for the webhook request.
//
// The property Mode is required.
type BaseMessageSendToToWebhookRecipientWebhookAuthenticationParam struct {
	// The authentication mode to use. Defaults to 'none' if not specified.
	//
	// Any of "none", "basic", "bearer".
	Mode string `json:"mode,omitzero,required"`
	// Token for bearer authentication.
	Token param.Opt[string] `json:"token,omitzero"`
	// Password for basic authentication.
	Password param.Opt[string] `json:"password,omitzero"`
	// Username for basic authentication.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

func (r BaseMessageSendToToWebhookRecipientWebhookAuthenticationParam) MarshalJSON() (data []byte, err error) {
	type shadow BaseMessageSendToToWebhookRecipientWebhookAuthenticationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaseMessageSendToToWebhookRecipientWebhookAuthenticationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BaseMessageSendToToWebhookRecipientWebhookAuthenticationParam](
		"mode", "none", "basic", "bearer",
	)
}

// ContentUnion contains all possible properties and values from
// [ElementalContent], [ContentElementalContentSugar].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ContentUnion struct {
	// This field is from variant [ElementalContent].
	Elements []ElementalNodeUnion `json:"elements"`
	// This field is from variant [ElementalContent].
	Version string `json:"version"`
	// This field is from variant [ElementalContent].
	Brand any `json:"brand"`
	// This field is from variant [ContentElementalContentSugar].
	Body string `json:"body"`
	// This field is from variant [ContentElementalContentSugar].
	Title string `json:"title"`
	JSON  struct {
		Elements respjson.Field
		Version  respjson.Field
		Brand    respjson.Field
		Body     respjson.Field
		Title    respjson.Field
		raw      string
	} `json:"-"`
}

func (u ContentUnion) AsElementalContent() (v ElementalContent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContentUnion) AsElementalContentSugar() (v ContentElementalContentSugar) {
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

// Syntatic Sugar to provide a fast shorthand for Courier Elemental Blocks.
type ContentElementalContentSugar struct {
	// The text content displayed in the notification.
	Body string `json:"body,required"`
	// The title to be displayed by supported channels i.e. push, email (as subject)
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

func ContentParamOfElementalContent(elements []ElementalNodeUnionParam, version string) ContentUnionParam {
	var variant ElementalContentParam
	variant.Elements = elements
	variant.Version = version
	return ContentUnionParam{OfElementalContent: &variant}
}

func ContentParamOfElementalContentSugar(body string, title string) ContentUnionParam {
	var variant ContentElementalContentSugarParam
	variant.Body = body
	variant.Title = title
	return ContentUnionParam{OfElementalContentSugar: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ContentUnionParam struct {
	OfElementalContent      *ElementalContentParam             `json:",omitzero,inline"`
	OfElementalContentSugar *ContentElementalContentSugarParam `json:",omitzero,inline"`
	paramUnion
}

func (u ContentUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfElementalContent, u.OfElementalContentSugar)
}
func (u *ContentUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ContentUnionParam) asAny() any {
	if !param.IsOmitted(u.OfElementalContent) {
		return u.OfElementalContent
	} else if !param.IsOmitted(u.OfElementalContentSugar) {
		return u.OfElementalContentSugar
	}
	return nil
}

// Syntatic Sugar to provide a fast shorthand for Courier Elemental Blocks.
//
// The properties Body, Title are required.
type ContentElementalContentSugarParam struct {
	// The text content displayed in the notification.
	Body string `json:"body,required"`
	// The title to be displayed by supported channels i.e. push, email (as subject)
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

func MessageParamOfContentMessage[
	T ElementalContentParam | ContentElementalContentSugarParam,
](content T) MessageUnionParam {
	var variant MessageContentMessageParam
	switch v := any(content).(type) {
	case ElementalContentParam:
		variant.Content.OfElementalContent = &v
	case ContentElementalContentSugarParam:
		variant.Content.OfElementalContentSugar = &v
	}
	return MessageUnionParam{OfContentMessage: &variant}
}

func MessageParamOfTemplateMessage(template string) MessageUnionParam {
	var variant MessageTemplateMessageParam
	variant.Template = template
	return MessageUnionParam{OfTemplateMessage: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MessageUnionParam struct {
	OfContentMessage  *MessageContentMessageParam  `json:",omitzero,inline"`
	OfTemplateMessage *MessageTemplateMessageParam `json:",omitzero,inline"`
	paramUnion
}

func (u MessageUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfContentMessage, u.OfTemplateMessage)
}
func (u *MessageUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *MessageUnionParam) asAny() any {
	if !param.IsOmitted(u.OfContentMessage) {
		return u.OfContentMessage
	} else if !param.IsOmitted(u.OfTemplateMessage) {
		return u.OfTemplateMessage
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MessageUnionParam) GetContent() *ContentUnionParam {
	if vt := u.OfContentMessage; vt != nil {
		return &vt.Content
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MessageUnionParam) GetTemplate() *string {
	if vt := u.OfTemplateMessage; vt != nil {
		return &vt.Template
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MessageUnionParam) GetBrandID() *string {
	if vt := u.OfContentMessage; vt != nil && vt.BrandID.Valid() {
		return &vt.BrandID.Value
	} else if vt := u.OfTemplateMessage; vt != nil && vt.BrandID.Valid() {
		return &vt.BrandID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's Channels property, if present.
func (u MessageUnionParam) GetChannels() map[string]BaseMessageChannelParam {
	if vt := u.OfContentMessage; vt != nil {
		return vt.Channels
	} else if vt := u.OfTemplateMessage; vt != nil {
		return vt.Channels
	}
	return nil
}

// Returns a pointer to the underlying variant's Context property, if present.
func (u MessageUnionParam) GetContext() *MessageContextParam {
	if vt := u.OfContentMessage; vt != nil {
		return &vt.Context
	} else if vt := u.OfTemplateMessage; vt != nil {
		return &vt.Context
	}
	return nil
}

// Returns a pointer to the underlying variant's Data property, if present.
func (u MessageUnionParam) GetData() map[string]any {
	if vt := u.OfContentMessage; vt != nil {
		return vt.Data
	} else if vt := u.OfTemplateMessage; vt != nil {
		return vt.Data
	}
	return nil
}

// Returns a pointer to the underlying variant's Delay property, if present.
func (u MessageUnionParam) GetDelay() *BaseMessageDelayParam {
	if vt := u.OfContentMessage; vt != nil {
		return &vt.Delay
	} else if vt := u.OfTemplateMessage; vt != nil {
		return &vt.Delay
	}
	return nil
}

// Returns a pointer to the underlying variant's Expiry property, if present.
func (u MessageUnionParam) GetExpiry() *BaseMessageExpiryParam {
	if vt := u.OfContentMessage; vt != nil {
		return &vt.Expiry
	} else if vt := u.OfTemplateMessage; vt != nil {
		return &vt.Expiry
	}
	return nil
}

// Returns a pointer to the underlying variant's Metadata property, if present.
func (u MessageUnionParam) GetMetadata() *BaseMessageMetadataParam {
	if vt := u.OfContentMessage; vt != nil {
		return &vt.Metadata
	} else if vt := u.OfTemplateMessage; vt != nil {
		return &vt.Metadata
	}
	return nil
}

// Returns a pointer to the underlying variant's Preferences property, if present.
func (u MessageUnionParam) GetPreferences() *BaseMessagePreferencesParam {
	if vt := u.OfContentMessage; vt != nil {
		return &vt.Preferences
	} else if vt := u.OfTemplateMessage; vt != nil {
		return &vt.Preferences
	}
	return nil
}

// Returns a pointer to the underlying variant's Providers property, if present.
func (u MessageUnionParam) GetProviders() map[string]BaseMessageProviderParam {
	if vt := u.OfContentMessage; vt != nil {
		return vt.Providers
	} else if vt := u.OfTemplateMessage; vt != nil {
		return vt.Providers
	}
	return nil
}

// Returns a pointer to the underlying variant's Routing property, if present.
func (u MessageUnionParam) GetRouting() *BaseMessageRoutingParam {
	if vt := u.OfContentMessage; vt != nil {
		return &vt.Routing
	} else if vt := u.OfTemplateMessage; vt != nil {
		return &vt.Routing
	}
	return nil
}

// Returns a pointer to the underlying variant's Timeout property, if present.
func (u MessageUnionParam) GetTimeout() *BaseMessageTimeoutParam {
	if vt := u.OfContentMessage; vt != nil {
		return &vt.Timeout
	} else if vt := u.OfTemplateMessage; vt != nil {
		return &vt.Timeout
	}
	return nil
}

// Returns a pointer to the underlying variant's To property, if present.
func (u MessageUnionParam) GetTo() *BaseMessageSendToToUnionParam {
	if vt := u.OfContentMessage; vt != nil {
		return &vt.To
	} else if vt := u.OfTemplateMessage; vt != nil {
		return &vt.To
	}
	return nil
}

// Describes the content of the message in a way that will work for email, push,
// chat, or any channel.
type MessageContentMessageParam struct {
	// Describes the content of the message in a way that will work for email, push,
	// chat, or any channel. Either this or template must be specified.
	Content ContentUnionParam `json:"content,omitzero,required"`
	BaseMessageParam
	BaseMessageSendToParam
}

func (r MessageContentMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageContentMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// A template for a type of message that can be sent more than once. For example,
// you might create an "Appointment Reminder" Notification or “Reset Password”
// Notifications.
type MessageTemplateMessageParam struct {
	// The id of the notification template to be rendered and sent to the recipient(s).
	// This field or the content field must be supplied.
	Template string `json:"template,required"`
	BaseMessageParam
	BaseMessageSendToParam
}

func (r MessageTemplateMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTemplateMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type MessageContext struct {
	// An id of a tenant, see
	// [tenants api docs](https://www.courier.com/docs/reference/tenants/). Will load
	// brand, default preferences and any other base context data associated with this
	// tenant.
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
	// An id of a tenant, see
	// [tenants api docs](https://www.courier.com/docs/reference/tenants/). Will load
	// brand, default preferences and any other base context data associated with this
	// tenant.
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

// The properties ServiceURL, TenantID are required.
type MsTeamsBasePropertiesParam struct {
	ServiceURL string `json:"service_url,required"`
	TenantID   string `json:"tenant_id,required"`
	paramObj
}

func (r MsTeamsBasePropertiesParam) MarshalJSON() (data []byte, err error) {
	type shadow MsTeamsBasePropertiesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MsTeamsBasePropertiesParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func RecipientParamOfAudienceRecipient(audienceID string) RecipientUnionParam {
	var variant RecipientAudienceRecipientParam
	variant.AudienceID = audienceID
	return RecipientUnionParam{OfAudienceRecipient: &variant}
}

func RecipientParamOfSlackRecipient[
	T RecipientSlackRecipientSlackSendToSlackChannelParam | RecipientSlackRecipientSlackSendToSlackEmailParam | RecipientSlackRecipientSlackSendToSlackUserIDParam,
](slack T) RecipientUnionParam {
	var variant RecipientSlackRecipientParam
	switch v := any(slack).(type) {
	case RecipientSlackRecipientSlackSendToSlackChannelParam:
		variant.Slack.OfSendToSlackChannel = &v
	case RecipientSlackRecipientSlackSendToSlackEmailParam:
		variant.Slack.OfSendToSlackEmail = &v
	case RecipientSlackRecipientSlackSendToSlackUserIDParam:
		variant.Slack.OfSendToSlackUserID = &v
	}
	return RecipientUnionParam{OfSlackRecipient: &variant}
}

func RecipientParamOfMsTeamsRecipient[
	T RecipientMsTeamsRecipientMsTeamsSendToMsTeamsUserIDParam | RecipientMsTeamsRecipientMsTeamsSendToMsTeamsEmailParam | RecipientMsTeamsRecipientMsTeamsSendToMsTeamsChannelIDParam | RecipientMsTeamsRecipientMsTeamsSendToMsTeamsConversationIDParam | RecipientMsTeamsRecipientMsTeamsSendToMsTeamsChannelNameParam,
](msTeams T) RecipientUnionParam {
	var variant RecipientMsTeamsRecipientParam
	switch v := any(msTeams).(type) {
	case RecipientMsTeamsRecipientMsTeamsSendToMsTeamsUserIDParam:
		variant.MsTeams.OfSendToMsTeamsUserID = &v
	case RecipientMsTeamsRecipientMsTeamsSendToMsTeamsEmailParam:
		variant.MsTeams.OfSendToMsTeamsEmail = &v
	case RecipientMsTeamsRecipientMsTeamsSendToMsTeamsChannelIDParam:
		variant.MsTeams.OfSendToMsTeamsChannelID = &v
	case RecipientMsTeamsRecipientMsTeamsSendToMsTeamsConversationIDParam:
		variant.MsTeams.OfSendToMsTeamsConversationID = &v
	case RecipientMsTeamsRecipientMsTeamsSendToMsTeamsChannelNameParam:
		variant.MsTeams.OfSendToMsTeamsChannelName = &v
	}
	return RecipientUnionParam{OfMsTeamsRecipient: &variant}
}

func RecipientParamOfPagerdutyRecipient(pagerduty RecipientPagerdutyRecipientPagerdutyParam) RecipientUnionParam {
	var variant RecipientPagerdutyRecipientParam
	variant.Pagerduty = pagerduty
	return RecipientUnionParam{OfPagerdutyRecipient: &variant}
}

func RecipientParamOfWebhookRecipient(webhook RecipientWebhookRecipientWebhookParam) RecipientUnionParam {
	var variant RecipientWebhookRecipientParam
	variant.Webhook = webhook
	return RecipientUnionParam{OfWebhookRecipient: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type RecipientUnionParam struct {
	OfAudienceRecipient  *RecipientAudienceRecipientParam  `json:",omitzero,inline"`
	OfRecipientObject    *RecipientObjectParam             `json:",omitzero,inline"`
	OfVariant2           *RecipientObjectParam             `json:",omitzero,inline"`
	OfUserRecipient      *UserRecipientParam               `json:",omitzero,inline"`
	OfSlackRecipient     *RecipientSlackRecipientParam     `json:",omitzero,inline"`
	OfMsTeamsRecipient   *RecipientMsTeamsRecipientParam   `json:",omitzero,inline"`
	OfRecipientData      map[string]any                    `json:",omitzero,inline"`
	OfPagerdutyRecipient *RecipientPagerdutyRecipientParam `json:",omitzero,inline"`
	OfWebhookRecipient   *RecipientWebhookRecipientParam   `json:",omitzero,inline"`
	paramUnion
}

func (u RecipientUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAudienceRecipient,
		u.OfRecipientObject,
		u.OfVariant2,
		u.OfUserRecipient,
		u.OfSlackRecipient,
		u.OfMsTeamsRecipient,
		u.OfRecipientData,
		u.OfPagerdutyRecipient,
		u.OfWebhookRecipient)
}
func (u *RecipientUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *RecipientUnionParam) asAny() any {
	if !param.IsOmitted(u.OfAudienceRecipient) {
		return u.OfAudienceRecipient
	} else if !param.IsOmitted(u.OfRecipientObject) {
		return u.OfRecipientObject
	} else if !param.IsOmitted(u.OfVariant2) {
		return u.OfVariant2
	} else if !param.IsOmitted(u.OfUserRecipient) {
		return u.OfUserRecipient
	} else if !param.IsOmitted(u.OfSlackRecipient) {
		return u.OfSlackRecipient
	} else if !param.IsOmitted(u.OfMsTeamsRecipient) {
		return u.OfMsTeamsRecipient
	} else if !param.IsOmitted(u.OfRecipientData) {
		return &u.OfRecipientData
	} else if !param.IsOmitted(u.OfPagerdutyRecipient) {
		return u.OfPagerdutyRecipient
	} else if !param.IsOmitted(u.OfWebhookRecipient) {
		return u.OfWebhookRecipient
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u RecipientUnionParam) GetAudienceID() *string {
	if vt := u.OfAudienceRecipient; vt != nil {
		return &vt.AudienceID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u RecipientUnionParam) GetListID() *string {
	if vt := u.OfRecipientObject; vt != nil && vt.ListID.Valid() {
		return &vt.ListID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u RecipientUnionParam) GetListPattern() *string {
	if vt := u.OfVariant2; vt != nil && vt.ListPattern.Valid() {
		return &vt.ListPattern.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u RecipientUnionParam) GetAccountID() *string {
	if vt := u.OfUserRecipient; vt != nil && vt.AccountID.Valid() {
		return &vt.AccountID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u RecipientUnionParam) GetContext() *MessageContextParam {
	if vt := u.OfUserRecipient; vt != nil {
		return &vt.Context
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u RecipientUnionParam) GetEmail() *string {
	if vt := u.OfUserRecipient; vt != nil && vt.Email.Valid() {
		return &vt.Email.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u RecipientUnionParam) GetLocale() *string {
	if vt := u.OfUserRecipient; vt != nil && vt.Locale.Valid() {
		return &vt.Locale.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u RecipientUnionParam) GetPhoneNumber() *string {
	if vt := u.OfUserRecipient; vt != nil && vt.PhoneNumber.Valid() {
		return &vt.PhoneNumber.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u RecipientUnionParam) GetPreferences() *UserRecipientPreferencesParam {
	if vt := u.OfUserRecipient; vt != nil {
		return &vt.Preferences
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u RecipientUnionParam) GetTenantID() *string {
	if vt := u.OfUserRecipient; vt != nil && vt.TenantID.Valid() {
		return &vt.TenantID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u RecipientUnionParam) GetUserID() *string {
	if vt := u.OfUserRecipient; vt != nil && vt.UserID.Valid() {
		return &vt.UserID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u RecipientUnionParam) GetSlack() *RecipientSlackRecipientSlackUnionParam {
	if vt := u.OfSlackRecipient; vt != nil {
		return &vt.Slack
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u RecipientUnionParam) GetMsTeams() *RecipientMsTeamsRecipientMsTeamsUnionParam {
	if vt := u.OfMsTeamsRecipient; vt != nil {
		return &vt.MsTeams
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u RecipientUnionParam) GetPagerduty() *RecipientPagerdutyRecipientPagerdutyParam {
	if vt := u.OfPagerdutyRecipient; vt != nil {
		return &vt.Pagerduty
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u RecipientUnionParam) GetWebhook() *RecipientWebhookRecipientWebhookParam {
	if vt := u.OfWebhookRecipient; vt != nil {
		return &vt.Webhook
	}
	return nil
}

// Returns a pointer to the underlying variant's Data property, if present.
func (u RecipientUnionParam) GetData() map[string]any {
	if vt := u.OfAudienceRecipient; vt != nil {
		return vt.Data
	} else if vt := u.OfRecipientObject; vt != nil {
		return vt.Data
	} else if vt := u.OfVariant2; vt != nil {
		return vt.Data
	} else if vt := u.OfUserRecipient; vt != nil {
		return vt.Data
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u RecipientUnionParam) GetFilters() (res recipientUnionParamFilters) {
	if vt := u.OfAudienceRecipient; vt != nil {
		res.any = &vt.Filters
	} else if vt := u.OfRecipientObject; vt != nil {
		res.any = &vt.Filters
	}
	return
}

// Can have the runtime types [_[]RecipientAudienceRecipientFilterParam],
// [_[]RecipientObjectFilterParam]
type recipientUnionParamFilters struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]courier.RecipientAudienceRecipientFilterParam:
//	case *[]courier.RecipientObjectFilterParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u recipientUnionParamFilters) AsAny() any { return u.any }

// The property AudienceID is required.
type RecipientAudienceRecipientParam struct {
	// A unique identifier associated with an Audience. A message will be sent to each
	// user in the audience.
	AudienceID string                                  `json:"audience_id,required"`
	Data       map[string]any                          `json:"data,omitzero"`
	Filters    []RecipientAudienceRecipientFilterParam `json:"filters,omitzero"`
	paramObj
}

func (r RecipientAudienceRecipientParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientAudienceRecipientParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecipientAudienceRecipientParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Operator, Path, Value are required.
type RecipientAudienceRecipientFilterParam struct {
	// Send to users only if they are member of the account
	//
	// Any of "MEMBER_OF".
	Operator string `json:"operator,omitzero,required"`
	// Any of "account_id".
	Path  string `json:"path,omitzero,required"`
	Value string `json:"value,required"`
	paramObj
}

func (r RecipientAudienceRecipientFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientAudienceRecipientFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecipientAudienceRecipientFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RecipientAudienceRecipientFilterParam](
		"operator", "MEMBER_OF",
	)
	apijson.RegisterFieldValidator[RecipientAudienceRecipientFilterParam](
		"path", "account_id",
	)
}

type RecipientObjectParam struct {
	ListID  param.Opt[string]            `json:"list_id,omitzero"`
	Data    map[string]any               `json:"data,omitzero"`
	Filters []RecipientObjectFilterParam `json:"filters,omitzero"`
	paramObj
}

func (r RecipientObjectParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientObjectParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecipientObjectParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Operator, Path, Value are required.
type RecipientObjectFilterParam struct {
	// Send to users only if they are member of the account
	//
	// Any of "MEMBER_OF".
	Operator string `json:"operator,omitzero,required"`
	// Any of "account_id".
	Path  string `json:"path,omitzero,required"`
	Value string `json:"value,required"`
	paramObj
}

func (r RecipientObjectFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientObjectFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecipientObjectFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RecipientObjectFilterParam](
		"operator", "MEMBER_OF",
	)
	apijson.RegisterFieldValidator[RecipientObjectFilterParam](
		"path", "account_id",
	)
}

// The property Slack is required.
type RecipientSlackRecipientParam struct {
	Slack RecipientSlackRecipientSlackUnionParam `json:"slack,omitzero,required"`
	paramObj
}

func (r RecipientSlackRecipientParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientSlackRecipientParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecipientSlackRecipientParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type RecipientSlackRecipientSlackUnionParam struct {
	OfSendToSlackChannel *RecipientSlackRecipientSlackSendToSlackChannelParam `json:",omitzero,inline"`
	OfSendToSlackEmail   *RecipientSlackRecipientSlackSendToSlackEmailParam   `json:",omitzero,inline"`
	OfSendToSlackUserID  *RecipientSlackRecipientSlackSendToSlackUserIDParam  `json:",omitzero,inline"`
	paramUnion
}

func (u RecipientSlackRecipientSlackUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSendToSlackChannel, u.OfSendToSlackEmail, u.OfSendToSlackUserID)
}
func (u *RecipientSlackRecipientSlackUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *RecipientSlackRecipientSlackUnionParam) asAny() any {
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
func (u RecipientSlackRecipientSlackUnionParam) GetChannel() *string {
	if vt := u.OfSendToSlackChannel; vt != nil {
		return &vt.Channel
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u RecipientSlackRecipientSlackUnionParam) GetEmail() *string {
	if vt := u.OfSendToSlackEmail; vt != nil {
		return &vt.Email
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u RecipientSlackRecipientSlackUnionParam) GetUserID() *string {
	if vt := u.OfSendToSlackUserID; vt != nil {
		return &vt.UserID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u RecipientSlackRecipientSlackUnionParam) GetAccessToken() *string {
	if vt := u.OfSendToSlackChannel; vt != nil {
		return (*string)(&vt.AccessToken)
	} else if vt := u.OfSendToSlackEmail; vt != nil {
		return (*string)(&vt.AccessToken)
	} else if vt := u.OfSendToSlackUserID; vt != nil {
		return (*string)(&vt.AccessToken)
	}
	return nil
}

type RecipientSlackRecipientSlackSendToSlackChannelParam struct {
	Channel string `json:"channel,required"`
	SlackBasePropertiesParam
}

func (r RecipientSlackRecipientSlackSendToSlackChannelParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientSlackRecipientSlackSendToSlackChannelParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type RecipientSlackRecipientSlackSendToSlackEmailParam struct {
	Email string `json:"email,required"`
	SlackBasePropertiesParam
}

func (r RecipientSlackRecipientSlackSendToSlackEmailParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientSlackRecipientSlackSendToSlackEmailParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type RecipientSlackRecipientSlackSendToSlackUserIDParam struct {
	UserID string `json:"user_id,required"`
	SlackBasePropertiesParam
}

func (r RecipientSlackRecipientSlackSendToSlackUserIDParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientSlackRecipientSlackSendToSlackUserIDParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// The property MsTeams is required.
type RecipientMsTeamsRecipientParam struct {
	MsTeams RecipientMsTeamsRecipientMsTeamsUnionParam `json:"ms_teams,omitzero,required"`
	paramObj
}

func (r RecipientMsTeamsRecipientParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientMsTeamsRecipientParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecipientMsTeamsRecipientParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type RecipientMsTeamsRecipientMsTeamsUnionParam struct {
	OfSendToMsTeamsUserID         *RecipientMsTeamsRecipientMsTeamsSendToMsTeamsUserIDParam         `json:",omitzero,inline"`
	OfSendToMsTeamsEmail          *RecipientMsTeamsRecipientMsTeamsSendToMsTeamsEmailParam          `json:",omitzero,inline"`
	OfSendToMsTeamsChannelID      *RecipientMsTeamsRecipientMsTeamsSendToMsTeamsChannelIDParam      `json:",omitzero,inline"`
	OfSendToMsTeamsConversationID *RecipientMsTeamsRecipientMsTeamsSendToMsTeamsConversationIDParam `json:",omitzero,inline"`
	OfSendToMsTeamsChannelName    *RecipientMsTeamsRecipientMsTeamsSendToMsTeamsChannelNameParam    `json:",omitzero,inline"`
	paramUnion
}

func (u RecipientMsTeamsRecipientMsTeamsUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSendToMsTeamsUserID,
		u.OfSendToMsTeamsEmail,
		u.OfSendToMsTeamsChannelID,
		u.OfSendToMsTeamsConversationID,
		u.OfSendToMsTeamsChannelName)
}
func (u *RecipientMsTeamsRecipientMsTeamsUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *RecipientMsTeamsRecipientMsTeamsUnionParam) asAny() any {
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
func (u RecipientMsTeamsRecipientMsTeamsUnionParam) GetUserID() *string {
	if vt := u.OfSendToMsTeamsUserID; vt != nil {
		return &vt.UserID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u RecipientMsTeamsRecipientMsTeamsUnionParam) GetEmail() *string {
	if vt := u.OfSendToMsTeamsEmail; vt != nil {
		return &vt.Email
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u RecipientMsTeamsRecipientMsTeamsUnionParam) GetChannelID() *string {
	if vt := u.OfSendToMsTeamsChannelID; vt != nil {
		return &vt.ChannelID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u RecipientMsTeamsRecipientMsTeamsUnionParam) GetConversationID() *string {
	if vt := u.OfSendToMsTeamsConversationID; vt != nil {
		return &vt.ConversationID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u RecipientMsTeamsRecipientMsTeamsUnionParam) GetChannelName() *string {
	if vt := u.OfSendToMsTeamsChannelName; vt != nil {
		return &vt.ChannelName
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u RecipientMsTeamsRecipientMsTeamsUnionParam) GetTeamID() *string {
	if vt := u.OfSendToMsTeamsChannelName; vt != nil {
		return &vt.TeamID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u RecipientMsTeamsRecipientMsTeamsUnionParam) GetServiceURL() *string {
	if vt := u.OfSendToMsTeamsUserID; vt != nil {
		return (*string)(&vt.ServiceURL)
	} else if vt := u.OfSendToMsTeamsEmail; vt != nil {
		return (*string)(&vt.ServiceURL)
	} else if vt := u.OfSendToMsTeamsChannelID; vt != nil {
		return (*string)(&vt.ServiceURL)
	} else if vt := u.OfSendToMsTeamsConversationID; vt != nil {
		return (*string)(&vt.ServiceURL)
	} else if vt := u.OfSendToMsTeamsChannelName; vt != nil {
		return (*string)(&vt.ServiceURL)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u RecipientMsTeamsRecipientMsTeamsUnionParam) GetTenantID() *string {
	if vt := u.OfSendToMsTeamsUserID; vt != nil {
		return (*string)(&vt.TenantID)
	} else if vt := u.OfSendToMsTeamsEmail; vt != nil {
		return (*string)(&vt.TenantID)
	} else if vt := u.OfSendToMsTeamsChannelID; vt != nil {
		return (*string)(&vt.TenantID)
	} else if vt := u.OfSendToMsTeamsConversationID; vt != nil {
		return (*string)(&vt.TenantID)
	} else if vt := u.OfSendToMsTeamsChannelName; vt != nil {
		return (*string)(&vt.TenantID)
	}
	return nil
}

type RecipientMsTeamsRecipientMsTeamsSendToMsTeamsUserIDParam struct {
	UserID string `json:"user_id,required"`
	MsTeamsBasePropertiesParam
}

func (r RecipientMsTeamsRecipientMsTeamsSendToMsTeamsUserIDParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientMsTeamsRecipientMsTeamsSendToMsTeamsUserIDParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type RecipientMsTeamsRecipientMsTeamsSendToMsTeamsEmailParam struct {
	Email string `json:"email,required"`
	MsTeamsBasePropertiesParam
}

func (r RecipientMsTeamsRecipientMsTeamsSendToMsTeamsEmailParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientMsTeamsRecipientMsTeamsSendToMsTeamsEmailParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type RecipientMsTeamsRecipientMsTeamsSendToMsTeamsChannelIDParam struct {
	ChannelID string `json:"channel_id,required"`
	MsTeamsBasePropertiesParam
}

func (r RecipientMsTeamsRecipientMsTeamsSendToMsTeamsChannelIDParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientMsTeamsRecipientMsTeamsSendToMsTeamsChannelIDParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type RecipientMsTeamsRecipientMsTeamsSendToMsTeamsConversationIDParam struct {
	ConversationID string `json:"conversation_id,required"`
	MsTeamsBasePropertiesParam
}

func (r RecipientMsTeamsRecipientMsTeamsSendToMsTeamsConversationIDParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientMsTeamsRecipientMsTeamsSendToMsTeamsConversationIDParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type RecipientMsTeamsRecipientMsTeamsSendToMsTeamsChannelNameParam struct {
	ChannelName string `json:"channel_name,required"`
	TeamID      string `json:"team_id,required"`
	MsTeamsBasePropertiesParam
}

func (r RecipientMsTeamsRecipientMsTeamsSendToMsTeamsChannelNameParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientMsTeamsRecipientMsTeamsSendToMsTeamsChannelNameParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// The property Pagerduty is required.
type RecipientPagerdutyRecipientParam struct {
	Pagerduty RecipientPagerdutyRecipientPagerdutyParam `json:"pagerduty,omitzero,required"`
	paramObj
}

func (r RecipientPagerdutyRecipientParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientPagerdutyRecipientParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecipientPagerdutyRecipientParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecipientPagerdutyRecipientPagerdutyParam struct {
	EventAction param.Opt[string] `json:"event_action,omitzero"`
	RoutingKey  param.Opt[string] `json:"routing_key,omitzero"`
	Severity    param.Opt[string] `json:"severity,omitzero"`
	Source      param.Opt[string] `json:"source,omitzero"`
	paramObj
}

func (r RecipientPagerdutyRecipientPagerdutyParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientPagerdutyRecipientPagerdutyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecipientPagerdutyRecipientPagerdutyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Webhook is required.
type RecipientWebhookRecipientParam struct {
	Webhook RecipientWebhookRecipientWebhookParam `json:"webhook,omitzero,required"`
	paramObj
}

func (r RecipientWebhookRecipientParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientWebhookRecipientParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecipientWebhookRecipientParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property URL is required.
type RecipientWebhookRecipientWebhookParam struct {
	// The URL to send the webhook request to.
	URL string `json:"url,required"`
	// Authentication configuration for the webhook request.
	Authentication RecipientWebhookRecipientWebhookAuthenticationParam `json:"authentication,omitzero"`
	// Custom headers to include in the webhook request.
	Headers map[string]string `json:"headers,omitzero"`
	// The HTTP method to use for the webhook request. Defaults to POST if not
	// specified.
	//
	// Any of "POST", "PUT".
	Method string `json:"method,omitzero"`
	// Specifies what profile information is included in the request payload. Defaults
	// to 'limited' if not specified.
	//
	// Any of "limited", "expanded".
	Profile string `json:"profile,omitzero"`
	paramObj
}

func (r RecipientWebhookRecipientWebhookParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientWebhookRecipientWebhookParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecipientWebhookRecipientWebhookParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RecipientWebhookRecipientWebhookParam](
		"method", "POST", "PUT",
	)
	apijson.RegisterFieldValidator[RecipientWebhookRecipientWebhookParam](
		"profile", "limited", "expanded",
	)
}

// Authentication configuration for the webhook request.
//
// The property Mode is required.
type RecipientWebhookRecipientWebhookAuthenticationParam struct {
	// The authentication mode to use. Defaults to 'none' if not specified.
	//
	// Any of "none", "basic", "bearer".
	Mode string `json:"mode,omitzero,required"`
	// Token for bearer authentication.
	Token param.Opt[string] `json:"token,omitzero"`
	// Password for basic authentication.
	Password param.Opt[string] `json:"password,omitzero"`
	// Username for basic authentication.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

func (r RecipientWebhookRecipientWebhookAuthenticationParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientWebhookRecipientWebhookAuthenticationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecipientWebhookRecipientWebhookAuthenticationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RecipientWebhookRecipientWebhookAuthenticationParam](
		"mode", "none", "basic", "bearer",
	)
}

// The property AccessToken is required.
type SlackBasePropertiesParam struct {
	AccessToken string `json:"access_token,required"`
	paramObj
}

func (r SlackBasePropertiesParam) MarshalJSON() (data []byte, err error) {
	type shadow SlackBasePropertiesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SlackBasePropertiesParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Utm struct {
	Campaign string `json:"campaign,nullable"`
	Content  string `json:"content,nullable"`
	Medium   string `json:"medium,nullable"`
	Source   string `json:"source,nullable"`
	Term     string `json:"term,nullable"`
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

type SendMessageResponse struct {
	// A successful call to `POST /send` returns a `202` status code along with a
	// `requestId` in the response body.
	//
	// For send requests that have a single recipient, the `requestId` is assigned to
	// the derived message as its message_id. Therefore the `requestId` can be supplied
	// to the Message's API for single recipient messages.
	//
	// For send requests that have multiple recipients (accounts, audiences, lists,
	// etc.), Courier assigns a unique id to each derived message as its `message_id`.
	// Therefore the `requestId` cannot be supplied to the Message's API for
	// single-recipient messages.
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
	// Defines the message to be delivered
	Message MessageUnionParam `json:"message,omitzero,required"`
	paramObj
}

func (r SendMessageParams) MarshalJSON() (data []byte, err error) {
	type shadow SendMessageParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendMessageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
