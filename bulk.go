// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/courier-go/internal/apijson"
	"github.com/stainless-sdks/courier-go/internal/apiquery"
	"github.com/stainless-sdks/courier-go/internal/requestconfig"
	"github.com/stainless-sdks/courier-go/option"
	"github.com/stainless-sdks/courier-go/packages/param"
	"github.com/stainless-sdks/courier-go/packages/respjson"
)

// BulkService contains methods and other services that help with interacting with
// the courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBulkService] method instead.
type BulkService struct {
	Options []option.RequestOption
}

// NewBulkService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBulkService(opts ...option.RequestOption) (r BulkService) {
	r = BulkService{}
	r.Options = opts
	return
}

// Ingest user data into a Bulk Job
func (r *BulkService) AddUsers(ctx context.Context, jobID string, body BulkAddUsersParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return
	}
	path := fmt.Sprintf("bulk/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Create a bulk job
func (r *BulkService) NewJob(ctx context.Context, body BulkNewJobParams, opts ...option.RequestOption) (res *BulkNewJobResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "bulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get Bulk Job Users
func (r *BulkService) ListUsers(ctx context.Context, jobID string, query BulkListUsersParams, opts ...option.RequestOption) (res *BulkListUsersResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return
	}
	path := fmt.Sprintf("bulk/%s/users", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get a bulk job
func (r *BulkService) GetJob(ctx context.Context, jobID string, opts ...option.RequestOption) (res *BulkGetJobResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return
	}
	path := fmt.Sprintf("bulk/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Run a bulk job
func (r *BulkService) RunJob(ctx context.Context, jobID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return
	}
	path := fmt.Sprintf("bulk/%s/run", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

type InboundBulkMessage struct {
	// A unique identifier that represents the brand that should be used for rendering
	// the notification.
	Brand string `json:"brand,nullable"`
	// JSON that includes any data you want to pass to a message template. The data
	// will populate the corresponding template variables.
	Data   map[string]any `json:"data,nullable"`
	Event  string         `json:"event,nullable"`
	Locale map[string]any `json:"locale,nullable"`
	// Describes the content of the message in a way that will work for email, push,
	// chat, or any channel.
	Message InboundBulkMessageMessageUnion `json:"message,nullable"`
	// JSON that is merged into the request sent by Courier to the provider to override
	// properties or to gain access to features in the provider API that are not
	// natively supported by Courier.
	Override any `json:"override"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Brand       respjson.Field
		Data        respjson.Field
		Event       respjson.Field
		Locale      respjson.Field
		Message     respjson.Field
		Override    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboundBulkMessage) RawJSON() string { return r.JSON.raw }
func (r *InboundBulkMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this InboundBulkMessage to a InboundBulkMessageParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// InboundBulkMessageParam.Overrides()
func (r InboundBulkMessage) ToParam() InboundBulkMessageParam {
	return param.Override[InboundBulkMessageParam](json.RawMessage(r.RawJSON()))
}

// InboundBulkMessageMessageUnion contains all possible properties and values from
// [InboundBulkMessageMessageInboundBulkTemplateMessage],
// [InboundBulkMessageMessageInboundBulkContentMessage].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type InboundBulkMessageMessageUnion struct {
	// This field is from variant
	// [InboundBulkMessageMessageInboundBulkTemplateMessage],
	// [InboundBulkMessageMessageInboundBulkContentMessage].
	BrandID string `json:"brand_id"`
	// This field is from variant
	// [InboundBulkMessageMessageInboundBulkTemplateMessage],
	// [InboundBulkMessageMessageInboundBulkContentMessage].
	Channels map[string]BaseMessageChannel `json:"channels"`
	// This field is from variant
	// [InboundBulkMessageMessageInboundBulkTemplateMessage],
	// [InboundBulkMessageMessageInboundBulkContentMessage].
	Context MessageContext `json:"context"`
	// This field is from variant
	// [InboundBulkMessageMessageInboundBulkTemplateMessage],
	// [InboundBulkMessageMessageInboundBulkContentMessage].
	Data map[string]any `json:"data"`
	// This field is from variant
	// [InboundBulkMessageMessageInboundBulkTemplateMessage],
	// [InboundBulkMessageMessageInboundBulkContentMessage].
	Delay BaseMessageDelay `json:"delay"`
	// This field is from variant
	// [InboundBulkMessageMessageInboundBulkTemplateMessage],
	// [InboundBulkMessageMessageInboundBulkContentMessage].
	Expiry BaseMessageExpiry `json:"expiry"`
	// This field is from variant
	// [InboundBulkMessageMessageInboundBulkTemplateMessage],
	// [InboundBulkMessageMessageInboundBulkContentMessage].
	Metadata BaseMessageMetadata `json:"metadata"`
	// This field is from variant
	// [InboundBulkMessageMessageInboundBulkTemplateMessage],
	// [InboundBulkMessageMessageInboundBulkContentMessage].
	Preferences BaseMessagePreferences `json:"preferences"`
	// This field is from variant
	// [InboundBulkMessageMessageInboundBulkTemplateMessage],
	// [InboundBulkMessageMessageInboundBulkContentMessage].
	Providers map[string]BaseMessageProvider `json:"providers"`
	// This field is from variant
	// [InboundBulkMessageMessageInboundBulkTemplateMessage],
	// [InboundBulkMessageMessageInboundBulkContentMessage].
	Routing BaseMessageRouting `json:"routing"`
	// This field is from variant
	// [InboundBulkMessageMessageInboundBulkTemplateMessage],
	// [InboundBulkMessageMessageInboundBulkContentMessage].
	Timeout BaseMessageTimeout `json:"timeout"`
	// This field is from variant
	// [InboundBulkMessageMessageInboundBulkTemplateMessage].
	Template string `json:"template"`
	// This field is from variant [InboundBulkMessageMessageInboundBulkContentMessage].
	Content ContentUnion `json:"content"`
	JSON    struct {
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
		Template    respjson.Field
		Content     respjson.Field
		raw         string
	} `json:"-"`
}

func (u InboundBulkMessageMessageUnion) AsInboundBulkTemplateMessage() (v InboundBulkMessageMessageInboundBulkTemplateMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InboundBulkMessageMessageUnion) AsInboundBulkContentMessage() (v InboundBulkMessageMessageInboundBulkContentMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InboundBulkMessageMessageUnion) RawJSON() string { return u.JSON.raw }

func (r *InboundBulkMessageMessageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes the content of the message in a way that will work for email, push,
// chat, or any channel.
type InboundBulkMessageMessageInboundBulkTemplateMessage struct {
	// The id of the notification template to be rendered and sent to the recipient(s).
	// This field or the content field must be supplied.
	Template string `json:"template,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Template    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BaseMessage
}

// Returns the unmodified JSON received from the API
func (r InboundBulkMessageMessageInboundBulkTemplateMessage) RawJSON() string { return r.JSON.raw }
func (r *InboundBulkMessageMessageInboundBulkTemplateMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A template for a type of message that can be sent more than once. For example,
// you might create an "Appointment Reminder" Notification or “Reset Password”
// Notifications.
type InboundBulkMessageMessageInboundBulkContentMessage struct {
	// Describes the content of the message in a way that will work for email, push,
	// chat, or any channel. Either this or template must be specified.
	Content ContentUnion `json:"content,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BaseMessage
}

// Returns the unmodified JSON received from the API
func (r InboundBulkMessageMessageInboundBulkContentMessage) RawJSON() string { return r.JSON.raw }
func (r *InboundBulkMessageMessageInboundBulkContentMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboundBulkMessageParam struct {
	// A unique identifier that represents the brand that should be used for rendering
	// the notification.
	Brand param.Opt[string] `json:"brand,omitzero"`
	Event param.Opt[string] `json:"event,omitzero"`
	// JSON that includes any data you want to pass to a message template. The data
	// will populate the corresponding template variables.
	Data   map[string]any `json:"data,omitzero"`
	Locale map[string]any `json:"locale,omitzero"`
	// Describes the content of the message in a way that will work for email, push,
	// chat, or any channel.
	Message InboundBulkMessageMessageUnionParam `json:"message,omitzero"`
	// JSON that is merged into the request sent by Courier to the provider to override
	// properties or to gain access to features in the provider API that are not
	// natively supported by Courier.
	Override any `json:"override,omitzero"`
	paramObj
}

func (r InboundBulkMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow InboundBulkMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InboundBulkMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InboundBulkMessageMessageUnionParam struct {
	OfInboundBulkTemplateMessage *InboundBulkMessageMessageInboundBulkTemplateMessageParam `json:",omitzero,inline"`
	OfInboundBulkContentMessage  *InboundBulkMessageMessageInboundBulkContentMessageParam  `json:",omitzero,inline"`
	paramUnion
}

func (u InboundBulkMessageMessageUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInboundBulkTemplateMessage, u.OfInboundBulkContentMessage)
}
func (u *InboundBulkMessageMessageUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *InboundBulkMessageMessageUnionParam) asAny() any {
	if !param.IsOmitted(u.OfInboundBulkTemplateMessage) {
		return u.OfInboundBulkTemplateMessage
	} else if !param.IsOmitted(u.OfInboundBulkContentMessage) {
		return u.OfInboundBulkContentMessage
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InboundBulkMessageMessageUnionParam) GetTemplate() *string {
	if vt := u.OfInboundBulkTemplateMessage; vt != nil {
		return &vt.Template
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InboundBulkMessageMessageUnionParam) GetContent() *ContentUnionParam {
	if vt := u.OfInboundBulkContentMessage; vt != nil {
		return &vt.Content
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InboundBulkMessageMessageUnionParam) GetBrandID() *string {
	if vt := u.OfInboundBulkTemplateMessage; vt != nil && vt.BrandID.Valid() {
		return &vt.BrandID.Value
	} else if vt := u.OfInboundBulkContentMessage; vt != nil && vt.BrandID.Valid() {
		return &vt.BrandID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's Channels property, if present.
func (u InboundBulkMessageMessageUnionParam) GetChannels() map[string]BaseMessageChannelParam {
	if vt := u.OfInboundBulkTemplateMessage; vt != nil {
		return vt.Channels
	} else if vt := u.OfInboundBulkContentMessage; vt != nil {
		return vt.Channels
	}
	return nil
}

// Returns a pointer to the underlying variant's Context property, if present.
func (u InboundBulkMessageMessageUnionParam) GetContext() *MessageContextParam {
	if vt := u.OfInboundBulkTemplateMessage; vt != nil {
		return &vt.Context
	} else if vt := u.OfInboundBulkContentMessage; vt != nil {
		return &vt.Context
	}
	return nil
}

// Returns a pointer to the underlying variant's Data property, if present.
func (u InboundBulkMessageMessageUnionParam) GetData() map[string]any {
	if vt := u.OfInboundBulkTemplateMessage; vt != nil {
		return vt.Data
	} else if vt := u.OfInboundBulkContentMessage; vt != nil {
		return vt.Data
	}
	return nil
}

// Returns a pointer to the underlying variant's Delay property, if present.
func (u InboundBulkMessageMessageUnionParam) GetDelay() *BaseMessageDelayParam {
	if vt := u.OfInboundBulkTemplateMessage; vt != nil {
		return &vt.Delay
	} else if vt := u.OfInboundBulkContentMessage; vt != nil {
		return &vt.Delay
	}
	return nil
}

// Returns a pointer to the underlying variant's Expiry property, if present.
func (u InboundBulkMessageMessageUnionParam) GetExpiry() *BaseMessageExpiryParam {
	if vt := u.OfInboundBulkTemplateMessage; vt != nil {
		return &vt.Expiry
	} else if vt := u.OfInboundBulkContentMessage; vt != nil {
		return &vt.Expiry
	}
	return nil
}

// Returns a pointer to the underlying variant's Metadata property, if present.
func (u InboundBulkMessageMessageUnionParam) GetMetadata() *BaseMessageMetadataParam {
	if vt := u.OfInboundBulkTemplateMessage; vt != nil {
		return &vt.Metadata
	} else if vt := u.OfInboundBulkContentMessage; vt != nil {
		return &vt.Metadata
	}
	return nil
}

// Returns a pointer to the underlying variant's Preferences property, if present.
func (u InboundBulkMessageMessageUnionParam) GetPreferences() *BaseMessagePreferencesParam {
	if vt := u.OfInboundBulkTemplateMessage; vt != nil {
		return &vt.Preferences
	} else if vt := u.OfInboundBulkContentMessage; vt != nil {
		return &vt.Preferences
	}
	return nil
}

// Returns a pointer to the underlying variant's Providers property, if present.
func (u InboundBulkMessageMessageUnionParam) GetProviders() map[string]BaseMessageProviderParam {
	if vt := u.OfInboundBulkTemplateMessage; vt != nil {
		return vt.Providers
	} else if vt := u.OfInboundBulkContentMessage; vt != nil {
		return vt.Providers
	}
	return nil
}

// Returns a pointer to the underlying variant's Routing property, if present.
func (u InboundBulkMessageMessageUnionParam) GetRouting() *BaseMessageRoutingParam {
	if vt := u.OfInboundBulkTemplateMessage; vt != nil {
		return &vt.Routing
	} else if vt := u.OfInboundBulkContentMessage; vt != nil {
		return &vt.Routing
	}
	return nil
}

// Returns a pointer to the underlying variant's Timeout property, if present.
func (u InboundBulkMessageMessageUnionParam) GetTimeout() *BaseMessageTimeoutParam {
	if vt := u.OfInboundBulkTemplateMessage; vt != nil {
		return &vt.Timeout
	} else if vt := u.OfInboundBulkContentMessage; vt != nil {
		return &vt.Timeout
	}
	return nil
}

// Describes the content of the message in a way that will work for email, push,
// chat, or any channel.
type InboundBulkMessageMessageInboundBulkTemplateMessageParam struct {
	// The id of the notification template to be rendered and sent to the recipient(s).
	// This field or the content field must be supplied.
	Template string `json:"template,required"`
	BaseMessageParam
}

func (r InboundBulkMessageMessageInboundBulkTemplateMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow InboundBulkMessageMessageInboundBulkTemplateMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// A template for a type of message that can be sent more than once. For example,
// you might create an "Appointment Reminder" Notification or “Reset Password”
// Notifications.
type InboundBulkMessageMessageInboundBulkContentMessageParam struct {
	// Describes the content of the message in a way that will work for email, push,
	// chat, or any channel. Either this or template must be specified.
	Content ContentUnionParam `json:"content,omitzero,required"`
	BaseMessageParam
}

func (r InboundBulkMessageMessageInboundBulkContentMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow InboundBulkMessageMessageInboundBulkContentMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type InboundBulkMessageUser struct {
	Data        any                  `json:"data"`
	Preferences RecipientPreferences `json:"preferences,nullable"`
	Profile     any                  `json:"profile"`
	Recipient   string               `json:"recipient,nullable"`
	To          UserRecipient        `json:"to,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Preferences respjson.Field
		Profile     respjson.Field
		Recipient   respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboundBulkMessageUser) RawJSON() string { return r.JSON.raw }
func (r *InboundBulkMessageUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this InboundBulkMessageUser to a InboundBulkMessageUserParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// InboundBulkMessageUserParam.Overrides()
func (r InboundBulkMessageUser) ToParam() InboundBulkMessageUserParam {
	return param.Override[InboundBulkMessageUserParam](json.RawMessage(r.RawJSON()))
}

type InboundBulkMessageUserParam struct {
	Recipient   param.Opt[string]         `json:"recipient,omitzero"`
	Data        any                       `json:"data,omitzero"`
	Preferences RecipientPreferencesParam `json:"preferences,omitzero"`
	Profile     any                       `json:"profile,omitzero"`
	To          UserRecipientParam        `json:"to,omitzero"`
	paramObj
}

func (r InboundBulkMessageUserParam) MarshalJSON() (data []byte, err error) {
	type shadow InboundBulkMessageUserParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InboundBulkMessageUserParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserRecipient struct {
	// Use `tenant_id` instad.
	AccountID string `json:"account_id,nullable"`
	// Context information such as tenant_id to send the notification with.
	Context MessageContext `json:"context,nullable"`
	Data    map[string]any `json:"data,nullable"`
	Email   string         `json:"email,nullable"`
	// The user's preferred ISO 639-1 language code.
	Locale      string                   `json:"locale,nullable"`
	PhoneNumber string                   `json:"phone_number,nullable"`
	Preferences UserRecipientPreferences `json:"preferences,nullable"`
	// An id of a tenant,
	// [see tenants api docs](https://www.courier.com/docs/reference/tenants). Will
	// load brand, default preferences and any other base context data associated with
	// this tenant.
	TenantID string `json:"tenant_id,nullable"`
	UserID   string `json:"user_id,nullable"`
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
	Notifications map[string]UserRecipientPreferencesNotification `json:"notifications,required"`
	Categories    map[string]UserRecipientPreferencesCategory     `json:"categories,nullable"`
	TemplateID    string                                          `json:"templateId,nullable"`
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

type UserRecipientPreferencesNotification struct {
	// Any of "OPTED_IN", "OPTED_OUT", "REQUIRED".
	Status             PreferenceStatus                                        `json:"status,required"`
	ChannelPreferences []UserRecipientPreferencesNotificationChannelPreference `json:"channel_preferences,nullable"`
	Rules              []UserRecipientPreferencesNotificationRule              `json:"rules,nullable"`
	// Any of "subscription", "list", "recipient".
	Source string `json:"source,nullable"`
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
func (r UserRecipientPreferencesNotification) RawJSON() string { return r.JSON.raw }
func (r *UserRecipientPreferencesNotification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserRecipientPreferencesNotificationChannelPreference struct {
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
func (r UserRecipientPreferencesNotificationChannelPreference) RawJSON() string { return r.JSON.raw }
func (r *UserRecipientPreferencesNotificationChannelPreference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserRecipientPreferencesNotificationRule struct {
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
func (r UserRecipientPreferencesNotificationRule) RawJSON() string { return r.JSON.raw }
func (r *UserRecipientPreferencesNotificationRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserRecipientPreferencesCategory struct {
	// Any of "OPTED_IN", "OPTED_OUT", "REQUIRED".
	Status             PreferenceStatus                                    `json:"status,required"`
	ChannelPreferences []UserRecipientPreferencesCategoryChannelPreference `json:"channel_preferences,nullable"`
	Rules              []UserRecipientPreferencesCategoryRule              `json:"rules,nullable"`
	// Any of "subscription", "list", "recipient".
	Source string `json:"source,nullable"`
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
func (r UserRecipientPreferencesCategory) RawJSON() string { return r.JSON.raw }
func (r *UserRecipientPreferencesCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserRecipientPreferencesCategoryChannelPreference struct {
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
func (r UserRecipientPreferencesCategoryChannelPreference) RawJSON() string { return r.JSON.raw }
func (r *UserRecipientPreferencesCategoryChannelPreference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserRecipientPreferencesCategoryRule struct {
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
func (r UserRecipientPreferencesCategoryRule) RawJSON() string { return r.JSON.raw }
func (r *UserRecipientPreferencesCategoryRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserRecipientParam struct {
	// Use `tenant_id` instad.
	AccountID param.Opt[string] `json:"account_id,omitzero"`
	Email     param.Opt[string] `json:"email,omitzero"`
	// The user's preferred ISO 639-1 language code.
	Locale      param.Opt[string] `json:"locale,omitzero"`
	PhoneNumber param.Opt[string] `json:"phone_number,omitzero"`
	// An id of a tenant,
	// [see tenants api docs](https://www.courier.com/docs/reference/tenants). Will
	// load brand, default preferences and any other base context data associated with
	// this tenant.
	TenantID    param.Opt[string]             `json:"tenant_id,omitzero"`
	UserID      param.Opt[string]             `json:"user_id,omitzero"`
	Data        map[string]any                `json:"data,omitzero"`
	Preferences UserRecipientPreferencesParam `json:"preferences,omitzero"`
	// Context information such as tenant_id to send the notification with.
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
	Notifications map[string]UserRecipientPreferencesNotificationParam `json:"notifications,omitzero,required"`
	TemplateID    param.Opt[string]                                    `json:"templateId,omitzero"`
	Categories    map[string]UserRecipientPreferencesCategoryParam     `json:"categories,omitzero"`
	paramObj
}

func (r UserRecipientPreferencesParam) MarshalJSON() (data []byte, err error) {
	type shadow UserRecipientPreferencesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserRecipientPreferencesParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Status is required.
type UserRecipientPreferencesNotificationParam struct {
	// Any of "OPTED_IN", "OPTED_OUT", "REQUIRED".
	Status             PreferenceStatus                                             `json:"status,omitzero,required"`
	ChannelPreferences []UserRecipientPreferencesNotificationChannelPreferenceParam `json:"channel_preferences,omitzero"`
	Rules              []UserRecipientPreferencesNotificationRuleParam              `json:"rules,omitzero"`
	// Any of "subscription", "list", "recipient".
	Source string `json:"source,omitzero"`
	paramObj
}

func (r UserRecipientPreferencesNotificationParam) MarshalJSON() (data []byte, err error) {
	type shadow UserRecipientPreferencesNotificationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserRecipientPreferencesNotificationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UserRecipientPreferencesNotificationParam](
		"source", "subscription", "list", "recipient",
	)
}

// The property Channel is required.
type UserRecipientPreferencesNotificationChannelPreferenceParam struct {
	// Any of "direct_message", "email", "push", "sms", "webhook", "inbox".
	Channel ChannelClassification `json:"channel,omitzero,required"`
	paramObj
}

func (r UserRecipientPreferencesNotificationChannelPreferenceParam) MarshalJSON() (data []byte, err error) {
	type shadow UserRecipientPreferencesNotificationChannelPreferenceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserRecipientPreferencesNotificationChannelPreferenceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Until is required.
type UserRecipientPreferencesNotificationRuleParam struct {
	Until string            `json:"until,required"`
	Start param.Opt[string] `json:"start,omitzero"`
	paramObj
}

func (r UserRecipientPreferencesNotificationRuleParam) MarshalJSON() (data []byte, err error) {
	type shadow UserRecipientPreferencesNotificationRuleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserRecipientPreferencesNotificationRuleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Status is required.
type UserRecipientPreferencesCategoryParam struct {
	// Any of "OPTED_IN", "OPTED_OUT", "REQUIRED".
	Status             PreferenceStatus                                         `json:"status,omitzero,required"`
	ChannelPreferences []UserRecipientPreferencesCategoryChannelPreferenceParam `json:"channel_preferences,omitzero"`
	Rules              []UserRecipientPreferencesCategoryRuleParam              `json:"rules,omitzero"`
	// Any of "subscription", "list", "recipient".
	Source string `json:"source,omitzero"`
	paramObj
}

func (r UserRecipientPreferencesCategoryParam) MarshalJSON() (data []byte, err error) {
	type shadow UserRecipientPreferencesCategoryParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserRecipientPreferencesCategoryParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UserRecipientPreferencesCategoryParam](
		"source", "subscription", "list", "recipient",
	)
}

// The property Channel is required.
type UserRecipientPreferencesCategoryChannelPreferenceParam struct {
	// Any of "direct_message", "email", "push", "sms", "webhook", "inbox".
	Channel ChannelClassification `json:"channel,omitzero,required"`
	paramObj
}

func (r UserRecipientPreferencesCategoryChannelPreferenceParam) MarshalJSON() (data []byte, err error) {
	type shadow UserRecipientPreferencesCategoryChannelPreferenceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserRecipientPreferencesCategoryChannelPreferenceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Until is required.
type UserRecipientPreferencesCategoryRuleParam struct {
	Until string            `json:"until,required"`
	Start param.Opt[string] `json:"start,omitzero"`
	paramObj
}

func (r UserRecipientPreferencesCategoryRuleParam) MarshalJSON() (data []byte, err error) {
	type shadow UserRecipientPreferencesCategoryRuleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserRecipientPreferencesCategoryRuleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BulkNewJobResponse struct {
	JobID string `json:"jobId,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		JobID       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BulkNewJobResponse) RawJSON() string { return r.JSON.raw }
func (r *BulkNewJobResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BulkListUsersResponse struct {
	Items  []BulkListUsersResponseItem `json:"items,required"`
	Paging Paging                      `json:"paging,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BulkListUsersResponse) RawJSON() string { return r.JSON.raw }
func (r *BulkListUsersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BulkListUsersResponseItem struct {
	// Any of "PENDING", "ENQUEUED", "ERROR".
	Status    string `json:"status,required"`
	MessageID string `json:"messageId,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		MessageID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	InboundBulkMessageUser
}

// Returns the unmodified JSON received from the API
func (r BulkListUsersResponseItem) RawJSON() string { return r.JSON.raw }
func (r *BulkListUsersResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BulkGetJobResponse struct {
	Job BulkGetJobResponseJob `json:"job,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Job         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BulkGetJobResponse) RawJSON() string { return r.JSON.raw }
func (r *BulkGetJobResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BulkGetJobResponseJob struct {
	Definition InboundBulkMessage `json:"definition,required"`
	Enqueued   int64              `json:"enqueued,required"`
	Failures   int64              `json:"failures,required"`
	Received   int64              `json:"received,required"`
	// Any of "CREATED", "PROCESSING", "COMPLETED", "ERROR".
	Status string `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Definition  respjson.Field
		Enqueued    respjson.Field
		Failures    respjson.Field
		Received    respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BulkGetJobResponseJob) RawJSON() string { return r.JSON.raw }
func (r *BulkGetJobResponseJob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BulkAddUsersParams struct {
	Users []InboundBulkMessageUserParam `json:"users,omitzero,required"`
	paramObj
}

func (r BulkAddUsersParams) MarshalJSON() (data []byte, err error) {
	type shadow BulkAddUsersParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BulkAddUsersParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BulkNewJobParams struct {
	Message InboundBulkMessageParam `json:"message,omitzero,required"`
	paramObj
}

func (r BulkNewJobParams) MarshalJSON() (data []byte, err error) {
	type shadow BulkNewJobParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BulkNewJobParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BulkListUsersParams struct {
	// A unique identifier that allows for fetching the next set of users added to the
	// bulk job
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BulkListUsersParams]'s query parameters as `url.Values`.
func (r BulkListUsersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
