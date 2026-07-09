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

// WorkspacePreferenceService contains methods and other services that help with
// interacting with the Courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkspacePreferenceService] method instead.
type WorkspacePreferenceService struct {
	Options []option.RequestOption
	Topics  WorkspacePreferenceTopicService
}

// NewWorkspacePreferenceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWorkspacePreferenceService(opts ...option.RequestOption) (r WorkspacePreferenceService) {
	r = WorkspacePreferenceService{}
	r.Options = opts
	r.Topics = NewWorkspacePreferenceTopicService(opts...)
	return
}

// Create a workspace preference. The workspace preference id is generated and
// returned. Topics are created inside a workspace preference via POST
// /preferences/sections/{section_id}/topics.
func (r *WorkspacePreferenceService) New(ctx context.Context, body WorkspacePreferenceNewParams, opts ...option.RequestOption) (res *WorkspacePreferenceGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "preferences/sections"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a workspace preference by id, including its topics.
func (r *WorkspacePreferenceService) Get(ctx context.Context, sectionID string, opts ...option.RequestOption) (res *WorkspacePreferenceGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if sectionID == "" {
		err = errors.New("missing required section_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("preferences/sections/%s", sectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List the workspace's preferences. Each workspace preference embeds its topics.
// Scoped to the workspace of the API key.
func (r *WorkspacePreferenceService) List(ctx context.Context, opts ...option.RequestOption) (res *WorkspacePreferenceListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "preferences/sections"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Archive a workspace preference. The workspace preference must be empty: delete
// its topics first, otherwise the request fails with 409.
func (r *WorkspacePreferenceService) Archive(ctx context.Context, sectionID string, opts ...option.RequestOption) (err error) {
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

// Publish the workspace's preferences page. Takes a snapshot of every workspace
// preference with its topics under a new published version, making the current
// state visible on the hosted preferences page (non-draft).
func (r *WorkspacePreferenceService) Publish(ctx context.Context, body WorkspacePreferencePublishParams, opts ...option.RequestOption) (res *PublishPreferencesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "preferences/publish"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Replace a workspace preference. Full document replacement; missing optional
// fields are cleared. Topics attached to the workspace preference are unaffected.
func (r *WorkspacePreferenceService) Replace(ctx context.Context, sectionID string, body WorkspacePreferenceReplaceParams, opts ...option.RequestOption) (res *WorkspacePreferenceGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if sectionID == "" {
		err = errors.New("missing required section_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("preferences/sections/%s", sectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Optional page metadata to apply when publishing the workspace's preferences
// page. All fields are optional; omitted fields fall back to the page defaults
// (and the workspace default brand).
type PublishPreferencesRequestParam struct {
	// Brand for the hosted page - "default" (workspace default brand), "none" (no
	// brand), or a specific brand id. Defaults to "default".
	BrandID param.Opt[string] `json:"brand_id,omitzero"`
	// Description shown under the heading on the hosted preferences page.
	Description param.Opt[string] `json:"description,omitzero"`
	// Heading shown at the top of the hosted preferences page.
	Heading param.Opt[string] `json:"heading,omitzero"`
	paramObj
}

func (r PublishPreferencesRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow PublishPreferencesRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublishPreferencesRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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

// Request body for creating a workspace preference.
//
// The property Name is required.
type WorkspacePreferenceCreateRequestParam struct {
	// Human-readable name for the workspace preference.
	Name string `json:"name" api:"required"`
	// Optional description shown under the section on the hosted preferences page.
	Description param.Opt[string] `json:"description,omitzero"`
	// Whether the workspace preference defines custom routing for its topics.
	HasCustomRouting param.Opt[bool] `json:"has_custom_routing,omitzero"`
	// Default channels for the workspace preference. Defaults to empty if omitted.
	RoutingOptions []shared.ChannelClassification `json:"routing_options,omitzero"`
	paramObj
}

func (r WorkspacePreferenceCreateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkspacePreferenceCreateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkspacePreferenceCreateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A workspace preference in your workspace, including its topics.
type WorkspacePreferenceGetResponse struct {
	// The workspace preference id.
	ID string `json:"id" api:"required"`
	// ISO-8601 timestamp of when the workspace preference was created.
	Created string `json:"created" api:"required"`
	// Whether the workspace preference defines custom routing for its topics.
	HasCustomRouting bool `json:"has_custom_routing" api:"required"`
	// Human-readable name.
	Name string `json:"name" api:"required"`
	// Default channels for the workspace preference. May be empty.
	RoutingOptions []shared.ChannelClassification `json:"routing_options" api:"required"`
	// The topics contained in this workspace preference.
	Topics []WorkspacePreferenceTopicGetResponse `json:"topics" api:"required"`
	// Id of the creator.
	Creator string `json:"creator" api:"nullable"`
	// Optional description shown under the section on the hosted preferences page.
	Description string `json:"description" api:"nullable"`
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
		Description      respjson.Field
		Updated          respjson.Field
		Updater          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkspacePreferenceGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkspacePreferenceGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The workspace's preferences, each with its topics.
type WorkspacePreferenceListResponse struct {
	Results []WorkspacePreferenceGetResponse `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkspacePreferenceListResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkspacePreferenceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body for replacing a workspace preference. Full document replacement;
// missing optional fields are cleared.
//
// The property Name is required.
type WorkspacePreferenceReplaceRequestParam struct {
	// Human-readable name for the workspace preference.
	Name string `json:"name" api:"required"`
	// Optional description shown under the section on the hosted preferences page.
	// Omit to clear.
	Description param.Opt[string] `json:"description,omitzero"`
	// Whether the workspace preference defines custom routing for its topics.
	HasCustomRouting param.Opt[bool] `json:"has_custom_routing,omitzero"`
	// Default channels for the workspace preference. Omit to clear.
	RoutingOptions []shared.ChannelClassification `json:"routing_options,omitzero"`
	paramObj
}

func (r WorkspacePreferenceReplaceRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkspacePreferenceReplaceRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkspacePreferenceReplaceRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body for creating a preference topic.
//
// The properties DefaultStatus, Name are required.
type WorkspacePreferenceTopicCreateRequestParam struct {
	// The default subscription status applied when a recipient has not set their own.
	//
	// Any of "OPTED_OUT", "OPTED_IN", "REQUIRED".
	DefaultStatus WorkspacePreferenceTopicCreateRequestDefaultStatus `json:"default_status,omitzero" api:"required"`
	// Human-readable name for the preference topic.
	Name string `json:"name" api:"required"`
	// Optional description shown under the topic on the hosted preferences page.
	Description param.Opt[string] `json:"description,omitzero"`
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

func (r WorkspacePreferenceTopicCreateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkspacePreferenceTopicCreateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkspacePreferenceTopicCreateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The default subscription status applied when a recipient has not set their own.
type WorkspacePreferenceTopicCreateRequestDefaultStatus string

const (
	WorkspacePreferenceTopicCreateRequestDefaultStatusOptedOut WorkspacePreferenceTopicCreateRequestDefaultStatus = "OPTED_OUT"
	WorkspacePreferenceTopicCreateRequestDefaultStatusOptedIn  WorkspacePreferenceTopicCreateRequestDefaultStatus = "OPTED_IN"
	WorkspacePreferenceTopicCreateRequestDefaultStatusRequired WorkspacePreferenceTopicCreateRequestDefaultStatus = "REQUIRED"
)

// A subscription preference topic in your workspace.
type WorkspacePreferenceTopicGetResponse struct {
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
	DefaultStatus WorkspacePreferenceTopicGetResponseDefaultStatus `json:"default_status" api:"required"`
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
	// Optional description shown under the topic on the hosted preferences page.
	Description string `json:"description" api:"nullable"`
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
		Description              respjson.Field
		Updater                  respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkspacePreferenceTopicGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkspacePreferenceTopicGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The default subscription status applied when a recipient has not set their own.
type WorkspacePreferenceTopicGetResponseDefaultStatus string

const (
	WorkspacePreferenceTopicGetResponseDefaultStatusOptedOut WorkspacePreferenceTopicGetResponseDefaultStatus = "OPTED_OUT"
	WorkspacePreferenceTopicGetResponseDefaultStatusOptedIn  WorkspacePreferenceTopicGetResponseDefaultStatus = "OPTED_IN"
	WorkspacePreferenceTopicGetResponseDefaultStatusRequired WorkspacePreferenceTopicGetResponseDefaultStatus = "REQUIRED"
)

// Topics contained in a workspace preference.
type WorkspacePreferenceTopicListResponse struct {
	Results []WorkspacePreferenceTopicGetResponse `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkspacePreferenceTopicListResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkspacePreferenceTopicListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body for replacing a preference topic. Full document replacement;
// missing optional fields are cleared.
//
// The properties DefaultStatus, Name are required.
type WorkspacePreferenceTopicReplaceRequestParam struct {
	// The default subscription status applied when a recipient has not set their own.
	//
	// Any of "OPTED_OUT", "OPTED_IN", "REQUIRED".
	DefaultStatus WorkspacePreferenceTopicReplaceRequestDefaultStatus `json:"default_status,omitzero" api:"required"`
	// Human-readable name for the preference topic.
	Name string `json:"name" api:"required"`
	// Optional description shown under the topic on the hosted preferences page. Omit
	// to clear.
	Description param.Opt[string] `json:"description,omitzero"`
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

func (r WorkspacePreferenceTopicReplaceRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkspacePreferenceTopicReplaceRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkspacePreferenceTopicReplaceRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The default subscription status applied when a recipient has not set their own.
type WorkspacePreferenceTopicReplaceRequestDefaultStatus string

const (
	WorkspacePreferenceTopicReplaceRequestDefaultStatusOptedOut WorkspacePreferenceTopicReplaceRequestDefaultStatus = "OPTED_OUT"
	WorkspacePreferenceTopicReplaceRequestDefaultStatusOptedIn  WorkspacePreferenceTopicReplaceRequestDefaultStatus = "OPTED_IN"
	WorkspacePreferenceTopicReplaceRequestDefaultStatusRequired WorkspacePreferenceTopicReplaceRequestDefaultStatus = "REQUIRED"
)

type WorkspacePreferenceNewParams struct {
	// Request body for creating a workspace preference.
	WorkspacePreferenceCreateRequest WorkspacePreferenceCreateRequestParam
	paramObj
}

func (r WorkspacePreferenceNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.WorkspacePreferenceCreateRequest)
}
func (r *WorkspacePreferenceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkspacePreferencePublishParams struct {
	// Optional page metadata to apply when publishing the workspace's preferences
	// page. All fields are optional; omitted fields fall back to the page defaults
	// (and the workspace default brand).
	PublishPreferencesRequest PublishPreferencesRequestParam
	paramObj
}

func (r WorkspacePreferencePublishParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublishPreferencesRequest)
}
func (r *WorkspacePreferencePublishParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkspacePreferenceReplaceParams struct {
	// Request body for replacing a workspace preference. Full document replacement;
	// missing optional fields are cleared.
	WorkspacePreferenceReplaceRequest WorkspacePreferenceReplaceRequestParam
	paramObj
}

func (r WorkspacePreferenceReplaceParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.WorkspacePreferenceReplaceRequest)
}
func (r *WorkspacePreferenceReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
