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
		return
	}
	path := fmt.Sprintf("users/%s/preferences", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Fetch user preferences for a specific subscription topic.
func (r *UserPreferenceService) GetTopic(ctx context.Context, topicID string, params UserPreferenceGetTopicParams, opts ...option.RequestOption) (res *UserPreferenceGetTopicResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.UserID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	if topicID == "" {
		err = errors.New("missing required topic_id parameter")
		return
	}
	path := fmt.Sprintf("users/%s/preferences/%s", params.UserID, topicID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Update or Create user preferences for a specific subscription topic.
func (r *UserPreferenceService) UpdateOrNewTopic(ctx context.Context, topicID string, params UserPreferenceUpdateOrNewTopicParams, opts ...option.RequestOption) (res *UserPreferenceUpdateOrNewTopicResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.UserID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	if topicID == "" {
		err = errors.New("missing required topic_id parameter")
		return
	}
	path := fmt.Sprintf("users/%s/preferences/%s", params.UserID, topicID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type TopicPreference struct {
	// Any of "OPTED_IN", "OPTED_OUT", "REQUIRED".
	DefaultStatus shared.PreferenceStatus `json:"default_status,required"`
	// Any of "OPTED_IN", "OPTED_OUT", "REQUIRED".
	Status    shared.PreferenceStatus `json:"status,required"`
	TopicID   string                  `json:"topic_id,required"`
	TopicName string                  `json:"topic_name,required"`
	// The Channels a user has chosen to receive notifications through for this topic
	CustomRouting    []shared.ChannelClassification `json:"custom_routing,nullable"`
	HasCustomRouting bool                           `json:"has_custom_routing,nullable"`
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
	Items []TopicPreference `json:"items,required"`
	// Deprecated - Paging not implemented on this endpoint
	Paging shared.Paging `json:"paging,required"`
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

type UserPreferenceGetTopicResponse struct {
	Topic TopicPreference `json:"topic,required"`
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
	Message string `json:"message,required"`
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

type UserPreferenceGetTopicParams struct {
	UserID string `path:"user_id,required" json:"-"`
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
	UserID string                                    `path:"user_id,required" json:"-"`
	Topic  UserPreferenceUpdateOrNewTopicParamsTopic `json:"topic,omitzero,required"`
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
	Status           shared.PreferenceStatus `json:"status,omitzero,required"`
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
