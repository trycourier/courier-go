// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/trycourier/courier-go/v4/internal/apijson"
	"github.com/trycourier/courier-go/v4/internal/apiquery"
	shimjson "github.com/trycourier/courier-go/v4/internal/encoding/json"
	"github.com/trycourier/courier-go/v4/internal/requestconfig"
	"github.com/trycourier/courier-go/v4/option"
	"github.com/trycourier/courier-go/v4/packages/param"
	"github.com/trycourier/courier-go/v4/packages/respjson"
)

// JourneyService contains methods and other services that help with interacting
// with the Courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewJourneyService] method instead.
type JourneyService struct {
	Options []option.RequestOption
}

// NewJourneyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewJourneyService(opts ...option.RequestOption) (r JourneyService) {
	r = JourneyService{}
	r.Options = opts
	return
}

// Get the list of journeys.
func (r *JourneyService) List(ctx context.Context, query JourneyListParams, opts ...option.RequestOption) (res *JourneysListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "journeys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Invoke a journey run from a journey template.
func (r *JourneyService) Invoke(ctx context.Context, templateID string, body JourneyInvokeParams, opts ...option.RequestOption) (res *JourneysInvokeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if templateID == "" {
		err = errors.New("missing required templateId parameter")
		return nil, err
	}
	path := fmt.Sprintf("journeys/%s/invoke", templateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// A journey template representing an automation workflow.
type Journey struct {
	// The unique identifier of the journey.
	ID string `json:"id" api:"required"`
	// The name of the journey.
	Name string `json:"name" api:"required"`
	// The version of the journey (published or draft).
	//
	// Any of "published", "draft".
	Version JourneyVersion `json:"version" api:"required"`
	// ISO 8601 timestamp when the journey was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// ISO 8601 timestamp when the journey was last updated.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Version     respjson.Field
		CreatedAt   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Journey) RawJSON() string { return r.JSON.raw }
func (r *Journey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The version of the journey (published or draft).
type JourneyVersion string

const (
	JourneyVersionPublished JourneyVersion = "published"
	JourneyVersionDraft     JourneyVersion = "draft"
)

// Request body for invoking a journey. Requires either a user identifier or a
// profile with contact information. User identifiers can be provided via user_id
// field, or resolved from profile/data objects (user_id, userId, or anonymousId
// fields).
type JourneysInvokeRequestParam struct {
	// A unique identifier for the user. If not provided, the system will attempt to
	// resolve the user identifier from profile or data objects.
	UserID param.Opt[string] `json:"user_id,omitzero"`
	// Data payload passed to the journey. The expected shape can be predefined using
	// the schema builder in the journey editor. This data is available in journey
	// steps for condition evaluation and template variable interpolation. Can also
	// contain user identifiers (user_id, userId, anonymousId) if not provided
	// elsewhere.
	Data map[string]any `json:"data,omitzero"`
	// Profile data for the user. Can contain contact information (email,
	// phone_number), user identifiers (user_id, userId, anonymousId), or any custom
	// profile fields. Profile fields are merged with any existing stored profile for
	// the user. Include context.tenant_id to load a tenant-scoped profile for
	// multi-tenant scenarios.
	Profile map[string]any `json:"profile,omitzero"`
	paramObj
}

func (r JourneysInvokeRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow JourneysInvokeRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JourneysInvokeRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JourneysInvokeResponse struct {
	// A unique identifier for the journey run that was created.
	RunID string `json:"runId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RunID       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JourneysInvokeResponse) RawJSON() string { return r.JSON.raw }
func (r *JourneysInvokeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JourneysListResponse struct {
	// A cursor token for pagination. Present when there are more results available.
	Cursor    string    `json:"cursor"`
	Templates []Journey `json:"templates"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cursor      respjson.Field
		Templates   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JourneysListResponse) RawJSON() string { return r.JSON.raw }
func (r *JourneysListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JourneyListParams struct {
	// A cursor token for pagination. Use the cursor from the previous response to
	// fetch the next page of results.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// The version of journeys to retrieve. Accepted values are published (for
	// published journeys) or draft (for draft journeys). Defaults to published.
	//
	// Any of "published", "draft".
	Version JourneyListParamsVersion `query:"version,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [JourneyListParams]'s query parameters as `url.Values`.
func (r JourneyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The version of journeys to retrieve. Accepted values are published (for
// published journeys) or draft (for draft journeys). Defaults to published.
type JourneyListParamsVersion string

const (
	JourneyListParamsVersionPublished JourneyListParamsVersion = "published"
	JourneyListParamsVersionDraft     JourneyListParamsVersion = "draft"
)

type JourneyInvokeParams struct {
	// Request body for invoking a journey. Requires either a user identifier or a
	// profile with contact information. User identifiers can be provided via user_id
	// field, or resolved from profile/data objects (user_id, userId, or anonymousId
	// fields).
	JourneysInvokeRequest JourneysInvokeRequestParam
	paramObj
}

func (r JourneyInvokeParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.JourneysInvokeRequest)
}
func (r *JourneyInvokeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
