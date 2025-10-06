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

	"github.com/trycourier/courier-go/internal/apijson"
	"github.com/trycourier/courier-go/internal/apiquery"
	"github.com/trycourier/courier-go/internal/requestconfig"
	"github.com/trycourier/courier-go/option"
	"github.com/trycourier/courier-go/packages/param"
	"github.com/trycourier/courier-go/packages/respjson"
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

// InboundBulkMessageUnion contains all possible properties and values from
// [InboundBulkMessageInboundBulkTemplateMessage],
// [InboundBulkMessageInboundBulkContentMessage].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type InboundBulkMessageUnion struct {
	// This field is from variant [InboundBulkMessageInboundBulkTemplateMessage].
	Template string `json:"template"`
	Brand    string `json:"brand"`
	Data     any    `json:"data"`
	Event    string `json:"event"`
	Locale   any    `json:"locale"`
	Override any    `json:"override"`
	// This field is from variant [InboundBulkMessageInboundBulkContentMessage].
	Content ContentUnion `json:"content"`
	JSON    struct {
		Template respjson.Field
		Brand    respjson.Field
		Data     respjson.Field
		Event    respjson.Field
		Locale   respjson.Field
		Override respjson.Field
		Content  respjson.Field
		raw      string
	} `json:"-"`
}

func (u InboundBulkMessageUnion) AsInboundBulkTemplateMessage() (v InboundBulkMessageInboundBulkTemplateMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InboundBulkMessageUnion) AsInboundBulkContentMessage() (v InboundBulkMessageInboundBulkContentMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InboundBulkMessageUnion) RawJSON() string { return u.JSON.raw }

func (r *InboundBulkMessageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this InboundBulkMessageUnion to a InboundBulkMessageUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// InboundBulkMessageUnionParam.Overrides()
func (r InboundBulkMessageUnion) ToParam() InboundBulkMessageUnionParam {
	return param.Override[InboundBulkMessageUnionParam](json.RawMessage(r.RawJSON()))
}

type InboundBulkMessageInboundBulkTemplateMessage struct {
	Template string                    `json:"template,required"`
	Brand    string                    `json:"brand,nullable"`
	Data     map[string]any            `json:"data,nullable"`
	Event    string                    `json:"event,nullable"`
	Locale   map[string]map[string]any `json:"locale,nullable"`
	Override map[string]any            `json:"override,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Template    respjson.Field
		Brand       respjson.Field
		Data        respjson.Field
		Event       respjson.Field
		Locale      respjson.Field
		Override    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboundBulkMessageInboundBulkTemplateMessage) RawJSON() string { return r.JSON.raw }
func (r *InboundBulkMessageInboundBulkTemplateMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboundBulkMessageInboundBulkContentMessage struct {
	// Syntactic sugar to provide a fast shorthand for Courier Elemental Blocks.
	Content  ContentUnion              `json:"content,required"`
	Brand    string                    `json:"brand,nullable"`
	Data     map[string]any            `json:"data,nullable"`
	Event    string                    `json:"event,nullable"`
	Locale   map[string]map[string]any `json:"locale,nullable"`
	Override map[string]any            `json:"override,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Brand       respjson.Field
		Data        respjson.Field
		Event       respjson.Field
		Locale      respjson.Field
		Override    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboundBulkMessageInboundBulkContentMessage) RawJSON() string { return r.JSON.raw }
func (r *InboundBulkMessageInboundBulkContentMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func InboundBulkMessageParamOfInboundBulkTemplateMessage(template string) InboundBulkMessageUnionParam {
	var variant InboundBulkMessageInboundBulkTemplateMessageParam
	variant.Template = template
	return InboundBulkMessageUnionParam{OfInboundBulkTemplateMessage: &variant}
}

func InboundBulkMessageParamOfInboundBulkContentMessage[
	T ContentElementalContentSugarParam | ElementalContentParam,
](content T) InboundBulkMessageUnionParam {
	var variant InboundBulkMessageInboundBulkContentMessageParam
	switch v := any(content).(type) {
	case ContentElementalContentSugarParam:
		variant.Content.OfElementalContentSugar = &v
	case ElementalContentParam:
		variant.Content.OfElementalContent = &v
	}
	return InboundBulkMessageUnionParam{OfInboundBulkContentMessage: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InboundBulkMessageUnionParam struct {
	OfInboundBulkTemplateMessage *InboundBulkMessageInboundBulkTemplateMessageParam `json:",omitzero,inline"`
	OfInboundBulkContentMessage  *InboundBulkMessageInboundBulkContentMessageParam  `json:",omitzero,inline"`
	paramUnion
}

func (u InboundBulkMessageUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInboundBulkTemplateMessage, u.OfInboundBulkContentMessage)
}
func (u *InboundBulkMessageUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *InboundBulkMessageUnionParam) asAny() any {
	if !param.IsOmitted(u.OfInboundBulkTemplateMessage) {
		return u.OfInboundBulkTemplateMessage
	} else if !param.IsOmitted(u.OfInboundBulkContentMessage) {
		return u.OfInboundBulkContentMessage
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InboundBulkMessageUnionParam) GetTemplate() *string {
	if vt := u.OfInboundBulkTemplateMessage; vt != nil {
		return &vt.Template
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InboundBulkMessageUnionParam) GetContent() *ContentUnionParam {
	if vt := u.OfInboundBulkContentMessage; vt != nil {
		return &vt.Content
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InboundBulkMessageUnionParam) GetBrand() *string {
	if vt := u.OfInboundBulkTemplateMessage; vt != nil && vt.Brand.Valid() {
		return &vt.Brand.Value
	} else if vt := u.OfInboundBulkContentMessage; vt != nil && vt.Brand.Valid() {
		return &vt.Brand.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InboundBulkMessageUnionParam) GetEvent() *string {
	if vt := u.OfInboundBulkTemplateMessage; vt != nil && vt.Event.Valid() {
		return &vt.Event.Value
	} else if vt := u.OfInboundBulkContentMessage; vt != nil && vt.Event.Valid() {
		return &vt.Event.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's Data property, if present.
func (u InboundBulkMessageUnionParam) GetData() map[string]any {
	if vt := u.OfInboundBulkTemplateMessage; vt != nil {
		return vt.Data
	} else if vt := u.OfInboundBulkContentMessage; vt != nil {
		return vt.Data
	}
	return nil
}

// Returns a pointer to the underlying variant's Locale property, if present.
func (u InboundBulkMessageUnionParam) GetLocale() map[string]map[string]any {
	if vt := u.OfInboundBulkTemplateMessage; vt != nil {
		return vt.Locale
	} else if vt := u.OfInboundBulkContentMessage; vt != nil {
		return vt.Locale
	}
	return nil
}

// Returns a pointer to the underlying variant's Override property, if present.
func (u InboundBulkMessageUnionParam) GetOverride() map[string]any {
	if vt := u.OfInboundBulkTemplateMessage; vt != nil {
		return vt.Override
	} else if vt := u.OfInboundBulkContentMessage; vt != nil {
		return vt.Override
	}
	return nil
}

// The property Template is required.
type InboundBulkMessageInboundBulkTemplateMessageParam struct {
	Template string                    `json:"template,required"`
	Brand    param.Opt[string]         `json:"brand,omitzero"`
	Event    param.Opt[string]         `json:"event,omitzero"`
	Data     map[string]any            `json:"data,omitzero"`
	Locale   map[string]map[string]any `json:"locale,omitzero"`
	Override map[string]any            `json:"override,omitzero"`
	paramObj
}

func (r InboundBulkMessageInboundBulkTemplateMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow InboundBulkMessageInboundBulkTemplateMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InboundBulkMessageInboundBulkTemplateMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Content is required.
type InboundBulkMessageInboundBulkContentMessageParam struct {
	// Syntactic sugar to provide a fast shorthand for Courier Elemental Blocks.
	Content  ContentUnionParam         `json:"content,omitzero,required"`
	Brand    param.Opt[string]         `json:"brand,omitzero"`
	Event    param.Opt[string]         `json:"event,omitzero"`
	Data     map[string]any            `json:"data,omitzero"`
	Locale   map[string]map[string]any `json:"locale,omitzero"`
	Override map[string]any            `json:"override,omitzero"`
	paramObj
}

func (r InboundBulkMessageInboundBulkContentMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow InboundBulkMessageInboundBulkContentMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InboundBulkMessageInboundBulkContentMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	// Use `tenant_id` instead.
	AccountID string `json:"account_id,nullable"`
	// Context such as tenant_id to send the notification with.
	Context MessageContext `json:"context,nullable"`
	Data    map[string]any `json:"data,nullable"`
	Email   string         `json:"email,nullable"`
	// The user's preferred ISO 639-1 language code.
	Locale      string                   `json:"locale,nullable"`
	PhoneNumber string                   `json:"phone_number,nullable"`
	Preferences UserRecipientPreferences `json:"preferences,nullable"`
	// Tenant id. Will load brand, default preferences and base context data.
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
	// Use `tenant_id` instead.
	AccountID param.Opt[string] `json:"account_id,omitzero"`
	Email     param.Opt[string] `json:"email,omitzero"`
	// The user's preferred ISO 639-1 language code.
	Locale      param.Opt[string] `json:"locale,omitzero"`
	PhoneNumber param.Opt[string] `json:"phone_number,omitzero"`
	// Tenant id. Will load brand, default preferences and base context data.
	TenantID    param.Opt[string]             `json:"tenant_id,omitzero"`
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
	Definition InboundBulkMessageUnion `json:"definition,required"`
	Enqueued   int64                   `json:"enqueued,required"`
	Failures   int64                   `json:"failures,required"`
	Received   int64                   `json:"received,required"`
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
	Message InboundBulkMessageUnionParam `json:"message,omitzero,required"`
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
