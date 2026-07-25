// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/trycourier/courier-go/v4/internal/apijson"
	"github.com/trycourier/courier-go/v4/internal/apiquery"
	"github.com/trycourier/courier-go/v4/internal/requestconfig"
	"github.com/trycourier/courier-go/v4/option"
	"github.com/trycourier/courier-go/v4/packages/param"
	"github.com/trycourier/courier-go/v4/packages/respjson"
	"github.com/trycourier/courier-go/v4/shared"
)

// UserPreferenceService contains methods and other services that help with
// interacting with the Courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserPreferenceService] method instead.
type UserPreferenceService struct {
	Options []option.RequestOption
}

// NewUserPreferenceService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserPreferenceService(opts ...option.RequestOption) (r UserPreferenceService) {
	r = UserPreferenceService{}
	r.Options = opts
	return
}

// Returns a user's preference overrides with paging, one entry per subscription
// topic they have set a choice for.
func (r *UserPreferenceService) Get(ctx context.Context, userID string, query UserPreferenceGetParams, opts ...option.RequestOption) (res *UserPreferenceGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("users/%s/preferences", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Replaces a user's entire set of preference overrides. Any topic you leave out is
// reset to its default, so send the full set rather than a subset.
func (r *UserPreferenceService) BulkReplace(ctx context.Context, userID string, params UserPreferenceBulkReplaceParams, opts ...option.RequestOption) (res *UserPreferenceBulkReplaceResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("users/%s/preferences", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Adds or updates a user's preferences for several subscription topics at once.
// Topics you leave out keep whatever they were set to before.
func (r *UserPreferenceService) BulkUpdate(ctx context.Context, userID string, params UserPreferenceBulkUpdateParams, opts ...option.RequestOption) (res *UserPreferenceBulkUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("users/%s/preferences", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Removes a user's override for one subscription topic, resetting it to the
// effective default from the tenant or workspace.
func (r *UserPreferenceService) DeleteTopic(ctx context.Context, topicID string, params UserPreferenceDeleteTopicParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.UserID == "" {
		err = errors.New("missing required user_id parameter")
		return err
	}
	if topicID == "" {
		err = errors.New("missing required topic_id parameter")
		return err
	}
	path := fmt.Sprintf("users/%s/preferences/%s", params.UserID, topicID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return err
}

// Returns a user's opt-in status and channel choices for one subscription topic,
// or the effective default if they have set no override.
func (r *UserPreferenceService) GetTopic(ctx context.Context, topicID string, params UserPreferenceGetTopicParams, opts ...option.RequestOption) (res *UserPreferenceGetTopicResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.UserID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	if topicID == "" {
		err = errors.New("missing required topic_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("users/%s/preferences/%s", params.UserID, topicID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Sets a user's opt-in status and channel choices for one subscription topic,
// overriding the tenant default for that topic only.
func (r *UserPreferenceService) UpdateOrNewTopic(ctx context.Context, topicID string, params UserPreferenceUpdateOrNewTopicParams, opts ...option.RequestOption) (res *UserPreferenceUpdateOrNewTopicResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.UserID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	if topicID == "" {
		err = errors.New("missing required topic_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("users/%s/preferences/%s", params.UserID, topicID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// A single topic override echoed in a bulk preference response.
type BulkPreferenceTopic struct {
	CustomRouting    []shared.ChannelClassification `json:"custom_routing" api:"required"`
	HasCustomRouting bool                           `json:"has_custom_routing" api:"required"`
	// The applied subscription status. Echoes the requested value, so it is always
	// OPTED_IN or OPTED_OUT.
	//
	// Any of "OPTED_IN", "OPTED_OUT".
	Status  BulkPreferenceTopicStatus `json:"status" api:"required"`
	TopicID string                    `json:"topic_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomRouting    respjson.Field
		HasCustomRouting respjson.Field
		Status           respjson.Field
		TopicID          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BulkPreferenceTopic) RawJSON() string { return r.JSON.raw }
func (r *BulkPreferenceTopic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The applied subscription status. Echoes the requested value, so it is always
// OPTED_IN or OPTED_OUT.
type BulkPreferenceTopicStatus string

const (
	BulkPreferenceTopicStatusOptedIn  BulkPreferenceTopicStatus = "OPTED_IN"
	BulkPreferenceTopicStatusOptedOut BulkPreferenceTopicStatus = "OPTED_OUT"
)

type TopicPreference struct {
	// The topic's default status, returned on reads. It applies whenever the user has
	// no override of their own (status equals this value).
	//
	// Any of "OPTED_IN", "OPTED_OUT", "REQUIRED".
	DefaultStatus shared.PreferenceStatus `json:"default_status" api:"required"`
	// The user's subscription status for this topic. OPTED_IN or OPTED_OUT reflect the
	// user's own choice; REQUIRED is a topic-level default set in the preferences
	// editor, not a user choice.
	//
	// Any of "OPTED_IN", "OPTED_OUT", "REQUIRED".
	Status shared.PreferenceStatus `json:"status" api:"required"`
	// The unique identifier of the subscription topic this preference applies to.
	TopicID string `json:"topic_id" api:"required"`
	// The display name of the subscription topic, returned on reads.
	TopicName string `json:"topic_name" api:"required"`
	// The channels the user has chosen to receive this topic on, present only when
	// has_custom_routing is true. One or more of: direct_message, email, push, sms,
	// webhook, inbox.
	CustomRouting []shared.ChannelClassification `json:"custom_routing" api:"nullable"`
	// Whether the user has chosen specific delivery channels for this topic (listed in
	// custom_routing) rather than the topic's default routing.
	HasCustomRouting bool `json:"has_custom_routing" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DefaultStatus    respjson.Field
		Status           respjson.Field
		TopicID          respjson.Field
		TopicName        respjson.Field
		CustomRouting    respjson.Field
		HasCustomRouting respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TopicPreference) RawJSON() string { return r.JSON.raw }
func (r *TopicPreference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserPreferenceGetResponse struct {
	// The Preferences associated with the user_id.
	Items []TopicPreference `json:"items" api:"required"`
	// Deprecated - Paging not implemented on this endpoint
	Paging shared.Paging `json:"paging" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserPreferenceGetResponse) RawJSON() string { return r.JSON.raw }
func (r *UserPreferenceGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserPreferenceBulkReplaceResponse struct {
	// The ids of the overrides that were reset to their topic default.
	Deleted []string `json:"deleted" api:"required"`
	// The complete resulting set of topic overrides for the user.
	Items []BulkPreferenceTopic `json:"items" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deleted     respjson.Field
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserPreferenceBulkReplaceResponse) RawJSON() string { return r.JSON.raw }
func (r *UserPreferenceBulkReplaceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserPreferenceBulkUpdateResponse struct {
	// The topics that could not be applied, each with a reason.
	Errors []UserPreferenceBulkUpdateResponseError `json:"errors" api:"required"`
	// The topics that were successfully created or updated.
	Items []BulkPreferenceTopic `json:"items" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Errors      respjson.Field
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserPreferenceBulkUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *UserPreferenceBulkUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single topic that could not be applied in a bulk preference request.
type UserPreferenceBulkUpdateResponseError struct {
	// A human-readable explanation of why the topic could not be applied.
	Reason  string `json:"reason" api:"required"`
	TopicID string `json:"topic_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Reason      respjson.Field
		TopicID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserPreferenceBulkUpdateResponseError) RawJSON() string { return r.JSON.raw }
func (r *UserPreferenceBulkUpdateResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserPreferenceGetTopicResponse struct {
	Topic TopicPreference `json:"topic" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Topic       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserPreferenceGetTopicResponse) RawJSON() string { return r.JSON.raw }
func (r *UserPreferenceGetTopicResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserPreferenceUpdateOrNewTopicResponse struct {
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserPreferenceUpdateOrNewTopicResponse) RawJSON() string { return r.JSON.raw }
func (r *UserPreferenceUpdateOrNewTopicResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserPreferenceGetParams struct {
	// Query the preferences of a user for this specific tenant context.
	TenantID param.Opt[string] `query:"tenant_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UserPreferenceGetParams]'s query parameters as
// `url.Values`.
func (r UserPreferenceGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserPreferenceBulkReplaceParams struct {
	// The complete set of topic overrides for the user. Up to 50 topics may be
	// provided. Any existing override not listed here is reset to its topic default;
	// an empty array resets every existing override.
	Topics []UserPreferenceBulkReplaceParamsTopic `json:"topics,omitzero" api:"required"`
	// Replace the preferences of a user for this specific tenant context.
	TenantID param.Opt[string] `query:"tenant_id,omitzero" json:"-"`
	paramObj
}

func (r UserPreferenceBulkReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow UserPreferenceBulkReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserPreferenceBulkReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [UserPreferenceBulkReplaceParams]'s query parameters as
// `url.Values`.
func (r UserPreferenceBulkReplaceParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The properties Status, TopicID are required.
type UserPreferenceBulkReplaceParamsTopic struct {
	// The subscription status to apply for this topic.
	//
	// Any of "OPTED_IN", "OPTED_OUT".
	Status string `json:"status,omitzero" api:"required"`
	// A unique identifier associated with a subscription topic.
	TopicID string `json:"topic_id" api:"required"`
	// Whether the recipient has chosen specific delivery channels for this topic.
	HasCustomRouting param.Opt[bool] `json:"has_custom_routing,omitzero"`
	// The channels a user has chosen to receive notifications through for this topic.
	CustomRouting []shared.ChannelClassification `json:"custom_routing,omitzero"`
	paramObj
}

func (r UserPreferenceBulkReplaceParamsTopic) MarshalJSON() (data []byte, err error) {
	type shadow UserPreferenceBulkReplaceParamsTopic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserPreferenceBulkReplaceParamsTopic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UserPreferenceBulkReplaceParamsTopic](
		"status", "OPTED_IN", "OPTED_OUT",
	)
}

type UserPreferenceBulkUpdateParams struct {
	// The topics to create or update. Between 1 and 50 topics may be provided in a
	// single request.
	Topics []UserPreferenceBulkUpdateParamsTopic `json:"topics,omitzero" api:"required"`
	// Update the preferences of a user for this specific tenant context.
	TenantID param.Opt[string] `query:"tenant_id,omitzero" json:"-"`
	paramObj
}

func (r UserPreferenceBulkUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow UserPreferenceBulkUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserPreferenceBulkUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [UserPreferenceBulkUpdateParams]'s query parameters as
// `url.Values`.
func (r UserPreferenceBulkUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The properties Status, TopicID are required.
type UserPreferenceBulkUpdateParamsTopic struct {
	// The subscription status to apply for this topic.
	//
	// Any of "OPTED_IN", "OPTED_OUT".
	Status string `json:"status,omitzero" api:"required"`
	// A unique identifier associated with a subscription topic.
	TopicID string `json:"topic_id" api:"required"`
	// Whether the recipient has chosen specific delivery channels for this topic.
	HasCustomRouting param.Opt[bool] `json:"has_custom_routing,omitzero"`
	// The channels a user has chosen to receive notifications through for this topic.
	CustomRouting []shared.ChannelClassification `json:"custom_routing,omitzero"`
	paramObj
}

func (r UserPreferenceBulkUpdateParamsTopic) MarshalJSON() (data []byte, err error) {
	type shadow UserPreferenceBulkUpdateParamsTopic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserPreferenceBulkUpdateParamsTopic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UserPreferenceBulkUpdateParamsTopic](
		"status", "OPTED_IN", "OPTED_OUT",
	)
}

type UserPreferenceDeleteTopicParams struct {
	UserID string `path:"user_id" api:"required" json:"-"`
	// Delete the preferences of a user for this specific tenant context.
	TenantID param.Opt[string] `query:"tenant_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UserPreferenceDeleteTopicParams]'s query parameters as
// `url.Values`.
func (r UserPreferenceDeleteTopicParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserPreferenceGetTopicParams struct {
	UserID string `path:"user_id" api:"required" json:"-"`
	// Query the preferences of a user for this specific tenant context.
	TenantID param.Opt[string] `query:"tenant_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UserPreferenceGetTopicParams]'s query parameters as
// `url.Values`.
func (r UserPreferenceGetTopicParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserPreferenceUpdateOrNewTopicParams struct {
	UserID string                                    `path:"user_id" api:"required" json:"-"`
	Topic  UserPreferenceUpdateOrNewTopicParamsTopic `json:"topic,omitzero" api:"required"`
	// Update the preferences of a user for this specific tenant context.
	TenantID param.Opt[string] `query:"tenant_id,omitzero" json:"-"`
	paramObj
}

func (r UserPreferenceUpdateOrNewTopicParams) MarshalJSON() (data []byte, err error) {
	type shadow UserPreferenceUpdateOrNewTopicParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserPreferenceUpdateOrNewTopicParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [UserPreferenceUpdateOrNewTopicParams]'s query parameters as
// `url.Values`.
func (r UserPreferenceUpdateOrNewTopicParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The property Status is required.
type UserPreferenceUpdateOrNewTopicParamsTopic struct {
	// The subscription status to set: OPTED_IN or OPTED_OUT. REQUIRED is a topic-level
	// default, not a user choice; the API rejects opting a user out of a REQUIRED
	// topic.
	//
	// Any of "OPTED_IN", "OPTED_OUT", "REQUIRED".
	Status shared.PreferenceStatus `json:"status,omitzero" api:"required"`
	// Set to true to route this topic to the channels in custom_routing instead of the
	// topic's default routing.
	HasCustomRouting param.Opt[bool] `json:"has_custom_routing,omitzero"`
	// The channels to deliver this topic on when has_custom_routing is true. One or
	// more of: direct_message, email, push, sms, webhook, inbox.
	CustomRouting []shared.ChannelClassification `json:"custom_routing,omitzero"`
	paramObj
}

func (r UserPreferenceUpdateOrNewTopicParamsTopic) MarshalJSON() (data []byte, err error) {
	type shadow UserPreferenceUpdateOrNewTopicParamsTopic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserPreferenceUpdateOrNewTopicParamsTopic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
