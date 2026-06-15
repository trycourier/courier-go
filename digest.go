// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"github.com/trycourier/courier-go/v4/internal/apijson"
	"github.com/trycourier/courier-go/v4/option"
	"github.com/trycourier/courier-go/v4/packages/respjson"
)

// DigestService contains methods and other services that help with interacting
// with the Courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDigestService] method instead.
type DigestService struct {
	Options   []option.RequestOption
	Schedules DigestScheduleService
}

// NewDigestService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDigestService(opts ...option.RequestOption) (r DigestService) {
	r = DigestService{}
	r.Options = opts
	r.Schedules = NewDigestScheduleService(opts...)
	return
}

type DigestCategory struct {
	// The key that identifies the category within the digest.
	CategoryKey string `json:"category_key" api:"required"`
	// Which events to keep when the number of events exceeds the retention limit for
	// the category.
	//
	// Any of "first", "last", "highest", "lowest", "none".
	Retain DigestCategoryRetain `json:"retain" api:"required"`
	// The data key used to order events when `retain` is `highest` or `lowest`.
	SortKey string `json:"sort_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CategoryKey respjson.Field
		Retain      respjson.Field
		SortKey     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DigestCategory) RawJSON() string { return r.JSON.raw }
func (r *DigestCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Which events to keep when the number of events exceeds the retention limit for
// the category.
type DigestCategoryRetain string

const (
	DigestCategoryRetainFirst   DigestCategoryRetain = "first"
	DigestCategoryRetainLast    DigestCategoryRetain = "last"
	DigestCategoryRetainHighest DigestCategoryRetain = "highest"
	DigestCategoryRetainLowest  DigestCategoryRetain = "lowest"
	DigestCategoryRetainNone    DigestCategoryRetain = "none"
)

type DigestInstance struct {
	// A unique identifier for the digest instance.
	DigestInstanceID string `json:"digest_instance_id" api:"required"`
	// The total number of events received for this instance.
	EventCount int64 `json:"event_count" api:"required"`
	// The status of the digest instance. `IN_PROGRESS` instances are still
	// accumulating events; `COMPLETED` instances have been released.
	//
	// Any of "IN_PROGRESS", "COMPLETED".
	Status DigestInstanceStatus `json:"status" api:"required"`
	// The ID of the user this digest instance belongs to.
	UserID string `json:"user_id" api:"required"`
	// The categories configured for the digest.
	Categories []DigestCategory `json:"categories"`
	// A map of category key to the number of events received for that category.
	CategoryKeyCounts map[string]int64 `json:"category_key_counts"`
	// An ISO 8601 timestamp of when the digest instance was created.
	CreatedAt string `json:"created_at"`
	// Whether the digest instance has been disabled.
	Disabled bool `json:"disabled"`
	// The ID of the tenant this digest instance belongs to, if any.
	TenantID string `json:"tenant_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DigestInstanceID  respjson.Field
		EventCount        respjson.Field
		Status            respjson.Field
		UserID            respjson.Field
		Categories        respjson.Field
		CategoryKeyCounts respjson.Field
		CreatedAt         respjson.Field
		Disabled          respjson.Field
		TenantID          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DigestInstance) RawJSON() string { return r.JSON.raw }
func (r *DigestInstance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the digest instance. `IN_PROGRESS` instances are still
// accumulating events; `COMPLETED` instances have been released.
type DigestInstanceStatus string

const (
	DigestInstanceStatusInProgress DigestInstanceStatus = "IN_PROGRESS"
	DigestInstanceStatusCompleted  DigestInstanceStatus = "COMPLETED"
)

type DigestInstanceListResponse struct {
	// Whether there are more digest instances to fetch using the cursor.
	HasMore bool `json:"has_more" api:"required"`
	// The digest instances for this page of results.
	Items []DigestInstance `json:"items" api:"required"`
	// Always `list` for a paginated list response.
	//
	// Any of "list".
	Type DigestInstanceListResponseType `json:"type" api:"required"`
	// A cursor token for fetching the next page of results, or null when there are
	// none.
	Cursor string `json:"cursor" api:"nullable"`
	// The path to fetch the next page of results, or null when there are none.
	NextURL string `json:"next_url" api:"nullable"`
	// The path of the current request.
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore     respjson.Field
		Items       respjson.Field
		Type        respjson.Field
		Cursor      respjson.Field
		NextURL     respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DigestInstanceListResponse) RawJSON() string { return r.JSON.raw }
func (r *DigestInstanceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Always `list` for a paginated list response.
type DigestInstanceListResponseType string

const (
	DigestInstanceListResponseTypeList DigestInstanceListResponseType = "list"
)
