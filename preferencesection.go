// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/trycourier/courier-go/v4/internal/apijson"
	shimjson "github.com/trycourier/courier-go/v4/internal/encoding/json"
	"github.com/trycourier/courier-go/v4/internal/requestconfig"
	"github.com/trycourier/courier-go/v4/option"
	"github.com/trycourier/courier-go/v4/packages/param"
	"github.com/trycourier/courier-go/v4/packages/respjson"
	"github.com/trycourier/courier-go/v4/shared"
)

// PreferenceSectionService contains methods and other services that help with
// interacting with the Courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPreferenceSectionService] method instead.
type PreferenceSectionService struct {
	Options []option.RequestOption
	Topics  PreferenceSectionTopicService
}

// NewPreferenceSectionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPreferenceSectionService(opts ...option.RequestOption) (r PreferenceSectionService) {
	r = PreferenceSectionService{}
	r.Options = opts
	r.Topics = NewPreferenceSectionTopicService(opts...)
	return
}

// Create a preference section in your workspace. The section id is generated and
// returned. Topics are created inside a section via POST
// /preferences/sections/{section_id}/topics.
func (r *PreferenceSectionService) New(ctx context.Context, body PreferenceSectionNewParams, opts ...option.RequestOption) (res *PreferenceSectionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "preferences/sections"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a preference section by id, including its topics.
func (r *PreferenceSectionService) Get(ctx context.Context, sectionID string, opts ...option.RequestOption) (res *PreferenceSectionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if sectionID == "" {
		err = errors.New("missing required section_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("preferences/sections/%s", sectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List the workspace's preference sections. Each section embeds its topics. Scoped
// to the workspace of the API key.
func (r *PreferenceSectionService) List(ctx context.Context, opts ...option.RequestOption) (res *PreferenceSectionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "preferences/sections"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Archive a preference section. The section must be empty: delete its topics
// first, otherwise the request fails with 409.
func (r *PreferenceSectionService) Archive(ctx context.Context, sectionID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if sectionID == "" {
		err = errors.New("missing required section_id parameter")
		return err
	}
	path := fmt.Sprintf("preferences/sections/%s", sectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Publish the workspace's preferences page. Takes a snapshot of every section with
// its topics under a new published version, making the current state visible on
// the hosted preferences page (non-draft).
func (r *PreferenceSectionService) Publish(ctx context.Context, opts ...option.RequestOption) (res *PublishPreferencesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "preferences/publish"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Replace a preference section. Full document replacement; missing optional fields
// are cleared. Topics attached to the section are unaffected.
func (r *PreferenceSectionService) Replace(ctx context.Context, sectionID string, body PreferenceSectionReplaceParams, opts ...option.RequestOption) (res *PreferenceSectionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if sectionID == "" {
		err = errors.New("missing required section_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("preferences/sections/%s", sectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Request body for creating a preference section.
//
// The property Name is required.
type PreferenceSectionCreateRequestParam struct {
	// Human-readable name for the section.
	Name string `json:"name" api:"required"`
	// Whether the section defines custom routing for its topics.
	HasCustomRouting param.Opt[bool] `json:"has_custom_routing,omitzero"`
	// Default channels for the section. Defaults to empty if omitted.
	RoutingOptions []shared.ChannelClassification `json:"routing_options,omitzero"`
	paramObj
}

func (r PreferenceSectionCreateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow PreferenceSectionCreateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PreferenceSectionCreateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A preference section in your workspace, including its topics.
type PreferenceSectionGetResponse struct {
	// The preference section id.
	ID string `json:"id" api:"required"`
	// ISO-8601 timestamp of when the section was created.
	Created string `json:"created" api:"required"`
	// Whether the section defines custom routing for its topics.
	HasCustomRouting bool `json:"has_custom_routing" api:"required"`
	// Human-readable name.
	Name string `json:"name" api:"required"`
	// Default channels for the section. May be empty.
	RoutingOptions []shared.ChannelClassification `json:"routing_options" api:"required"`
	// The topics contained in this section.
	Topics []PreferenceTopicGetResponse `json:"topics" api:"required"`
	// Id of the creator.
	Creator string `json:"creator" api:"nullable"`
	// ISO-8601 timestamp of the last update.
	Updated string `json:"updated" api:"nullable"`
	// Id of the last updater.
	Updater string `json:"updater" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Created          respjson.Field
		HasCustomRouting respjson.Field
		Name             respjson.Field
		RoutingOptions   respjson.Field
		Topics           respjson.Field
		Creator          respjson.Field
		Updated          respjson.Field
		Updater          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PreferenceSectionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PreferenceSectionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The workspace's preference sections, each with its topics.
type PreferenceSectionListResponse struct {
	Results []PreferenceSectionGetResponse `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PreferenceSectionListResponse) RawJSON() string { return r.JSON.raw }
func (r *PreferenceSectionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body for replacing a preference section. Full document replacement;
// missing optional fields are cleared.
//
// The property Name is required.
type PreferenceSectionReplaceRequestParam struct {
	// Human-readable name for the section.
	Name string `json:"name" api:"required"`
	// Whether the section defines custom routing for its topics.
	HasCustomRouting param.Opt[bool] `json:"has_custom_routing,omitzero"`
	// Default channels for the section. Omit to clear.
	RoutingOptions []shared.ChannelClassification `json:"routing_options,omitzero"`
	paramObj
}

func (r PreferenceSectionReplaceRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow PreferenceSectionReplaceRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PreferenceSectionReplaceRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body for creating a preference topic.
//
// The properties DefaultStatus, Name are required.
type PreferenceTopicCreateRequestParam struct {
	// The default subscription status applied when a recipient has not set their own.
	//
	// Any of "OPTED_OUT", "OPTED_IN", "REQUIRED".
	DefaultStatus PreferenceTopicCreateRequestDefaultStatus `json:"default_status,omitzero" api:"required"`
	// Human-readable name for the preference topic.
	Name string `json:"name" api:"required"`
	// Whether to include a list-unsubscribe header on emails for this topic.
	IncludeUnsubscribeHeader param.Opt[bool] `json:"include_unsubscribe_header,omitzero"`
	// Preference controls a recipient may customize for this topic. Defaults to empty
	// if omitted.
	//
	// Any of "snooze", "channel_preferences".
	AllowedPreferences []string `json:"allowed_preferences,omitzero"`
	// Default channels delivered for this topic. Defaults to empty if omitted.
	RoutingOptions []shared.ChannelClassification `json:"routing_options,omitzero"`
	// Arbitrary metadata associated with the topic.
	TopicData map[string]any `json:"topic_data,omitzero"`
	paramObj
}

func (r PreferenceTopicCreateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow PreferenceTopicCreateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PreferenceTopicCreateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The default subscription status applied when a recipient has not set their own.
type PreferenceTopicCreateRequestDefaultStatus string

const (
	PreferenceTopicCreateRequestDefaultStatusOptedOut PreferenceTopicCreateRequestDefaultStatus = "OPTED_OUT"
	PreferenceTopicCreateRequestDefaultStatusOptedIn  PreferenceTopicCreateRequestDefaultStatus = "OPTED_IN"
	PreferenceTopicCreateRequestDefaultStatusRequired PreferenceTopicCreateRequestDefaultStatus = "REQUIRED"
)

// A subscription preference topic in your workspace.
type PreferenceTopicGetResponse struct {
	// The preference topic id.
	ID string `json:"id" api:"required"`
	// Preference controls a recipient may customize. May be empty.
	//
	// Any of "snooze", "channel_preferences".
	AllowedPreferences []string `json:"allowed_preferences" api:"required"`
	// ISO-8601 timestamp of when the topic was created.
	Created string `json:"created" api:"required"`
	// The default subscription status applied when a recipient has not set their own.
	//
	// Any of "OPTED_OUT", "OPTED_IN", "REQUIRED".
	DefaultStatus PreferenceTopicGetResponseDefaultStatus `json:"default_status" api:"required"`
	// Whether a list-unsubscribe header is included on emails for this topic.
	IncludeUnsubscribeHeader bool `json:"include_unsubscribe_header" api:"required"`
	// Human-readable name.
	Name string `json:"name" api:"required"`
	// Default channels delivered for this topic. May be empty.
	RoutingOptions []shared.ChannelClassification `json:"routing_options" api:"required"`
	// Arbitrary metadata associated with the topic.
	TopicData map[string]any `json:"topic_data" api:"required"`
	// ISO-8601 timestamp of the last update.
	Updated string `json:"updated" api:"required"`
	// Id of the creator.
	Creator string `json:"creator" api:"nullable"`
	// Id of the last updater.
	Updater string `json:"updater" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		AllowedPreferences       respjson.Field
		Created                  respjson.Field
		DefaultStatus            respjson.Field
		IncludeUnsubscribeHeader respjson.Field
		Name                     respjson.Field
		RoutingOptions           respjson.Field
		TopicData                respjson.Field
		Updated                  respjson.Field
		Creator                  respjson.Field
		Updater                  respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PreferenceTopicGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PreferenceTopicGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The default subscription status applied when a recipient has not set their own.
type PreferenceTopicGetResponseDefaultStatus string

const (
	PreferenceTopicGetResponseDefaultStatusOptedOut PreferenceTopicGetResponseDefaultStatus = "OPTED_OUT"
	PreferenceTopicGetResponseDefaultStatusOptedIn  PreferenceTopicGetResponseDefaultStatus = "OPTED_IN"
	PreferenceTopicGetResponseDefaultStatusRequired PreferenceTopicGetResponseDefaultStatus = "REQUIRED"
)

// Topics contained in a preference section.
type PreferenceTopicListResponse struct {
	Results []PreferenceTopicGetResponse `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PreferenceTopicListResponse) RawJSON() string { return r.JSON.raw }
func (r *PreferenceTopicListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body for replacing a preference topic. Full document replacement;
// missing optional fields are cleared.
//
// The properties DefaultStatus, Name are required.
type PreferenceTopicReplaceRequestParam struct {
	// The default subscription status applied when a recipient has not set their own.
	//
	// Any of "OPTED_OUT", "OPTED_IN", "REQUIRED".
	DefaultStatus PreferenceTopicReplaceRequestDefaultStatus `json:"default_status,omitzero" api:"required"`
	// Human-readable name for the preference topic.
	Name string `json:"name" api:"required"`
	// Whether to include a list-unsubscribe header on emails for this topic.
	IncludeUnsubscribeHeader param.Opt[bool] `json:"include_unsubscribe_header,omitzero"`
	// Preference controls a recipient may customize. Omit to clear.
	//
	// Any of "snooze", "channel_preferences".
	AllowedPreferences []string `json:"allowed_preferences,omitzero"`
	// Default channels delivered for this topic. Omit to clear.
	RoutingOptions []shared.ChannelClassification `json:"routing_options,omitzero"`
	// Arbitrary metadata associated with the topic. Omit to clear.
	TopicData map[string]any `json:"topic_data,omitzero"`
	paramObj
}

func (r PreferenceTopicReplaceRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow PreferenceTopicReplaceRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PreferenceTopicReplaceRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The default subscription status applied when a recipient has not set their own.
type PreferenceTopicReplaceRequestDefaultStatus string

const (
	PreferenceTopicReplaceRequestDefaultStatusOptedOut PreferenceTopicReplaceRequestDefaultStatus = "OPTED_OUT"
	PreferenceTopicReplaceRequestDefaultStatusOptedIn  PreferenceTopicReplaceRequestDefaultStatus = "OPTED_IN"
	PreferenceTopicReplaceRequestDefaultStatusRequired PreferenceTopicReplaceRequestDefaultStatus = "REQUIRED"
)

// Result of publishing the workspace's preferences page.
type PublishPreferencesResponse struct {
	// Id of the published page snapshot.
	PageID string `json:"page_id" api:"required"`
	// ISO-8601 timestamp of the publish.
	PublishedAt string `json:"published_at" api:"required"`
	// Monotonic published version (epoch milliseconds).
	PublishedVersion float64 `json:"published_version" api:"required"`
	// Draft-mode hosted preferences page URL for previewing.
	PreviewURL string `json:"preview_url" api:"nullable"`
	// Id of the publisher.
	PublishedBy string `json:"published_by" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageID           respjson.Field
		PublishedAt      respjson.Field
		PublishedVersion respjson.Field
		PreviewURL       respjson.Field
		PublishedBy      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublishPreferencesResponse) RawJSON() string { return r.JSON.raw }
func (r *PublishPreferencesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PreferenceSectionNewParams struct {
	// Request body for creating a preference section.
	PreferenceSectionCreateRequest PreferenceSectionCreateRequestParam
	paramObj
}

func (r PreferenceSectionNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PreferenceSectionCreateRequest)
}
func (r *PreferenceSectionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PreferenceSectionReplaceParams struct {
	// Request body for replacing a preference section. Full document replacement;
	// missing optional fields are cleared.
	PreferenceSectionReplaceRequest PreferenceSectionReplaceRequestParam
	paramObj
}

func (r PreferenceSectionReplaceParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PreferenceSectionReplaceRequest)
}
func (r *PreferenceSectionReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
