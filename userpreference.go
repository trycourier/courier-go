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

// Fetch all user preferences.
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

// Replace a user's complete set of preference overrides in a single request. The
// topics in the request body become the recipient's entire set of overrides:
// listed topics are created or updated, and every existing override that is not
// included in the body is reset to its topic default. Submitting an empty `topics`
// array is a valid clear-all that resets every existing override.
//
// This operation is validation-atomic (all-or-nothing): structural validation
// fails fast with a single `400`, and if any topic is semantically invalid (an
// unknown topic, a `REQUIRED` topic that cannot be opted out, or a custom routing
// request that is not available on the workspace's plan) the request returns a
// single `400` aggregating every failure in `errors` and writes nothing. On
// success it returns `200` with `items` (the complete resulting override set) and
// `deleted` (the ids of the overrides that were reset to default).
//
// Every `topic_id` in the response — in `items`, `deleted`, and any `errors` — is
// returned in Courier's canonical topic id form, regardless of the form supplied
// in the request.
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

// Additively create or update a user's preferences for one or more subscription
// topics in a single request. Only the topics included in the request body are
// created or updated; any existing overrides for topics not listed are left
// untouched.
//
// Structural validation of the request body fails fast with a single `400`. Beyond
// that, each topic is processed independently (partial-success, not
// all-or-nothing): valid topics are written and returned in `items`, while topics
// that cannot be applied are collected in `errors` with a per-topic `reason` (for
// example an unknown topic, a `REQUIRED` topic that cannot be opted out, a custom
// routing request that is not available on the workspace's plan, or a write
// failure). The request therefore returns `200` with both lists whenever the body
// is structurally valid.
//
// Every `topic_id` in the response — in both `items` and `errors` — is returned in
// Courier's canonical topic id form, regardless of the form supplied in the
// request.
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

// Remove a user's preferences for a specific subscription topic, resetting the
// topic to its effective default. This operation is idempotent: deleting a
// preference that does not exist succeeds with no error.
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

// Fetch user preferences for a specific subscription topic.
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

// Update or Create user preferences for a specific subscription topic.
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
	// Any of "OPTED_IN", "OPTED_OUT", "REQUIRED".
	DefaultStatus shared.PreferenceStatus `json:"default_status" api:"required"`
	// Any of "OPTED_IN", "OPTED_OUT", "REQUIRED".
	Status    shared.PreferenceStatus `json:"status" api:"required"`
	TopicID   string                  `json:"topic_id" api:"required"`
	TopicName string                  `json:"topic_name" api:"required"`
	// The Channels a user has chosen to receive notifications through for this topic
	CustomRouting    []shared.ChannelClassification `json:"custom_routing" api:"nullable"`
	HasCustomRouting bool                           `json:"has_custom_routing" api:"nullable"`
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
	// Update the preferences of a user for this specific tenant context.
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
	// Any of "OPTED_IN", "OPTED_OUT", "REQUIRED".
	Status           shared.PreferenceStatus `json:"status,omitzero" api:"required"`
	HasCustomRouting param.Opt[bool]         `json:"has_custom_routing,omitzero"`
	// The Channels a user has chosen to receive notifications through for this topic
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
