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

	"github.com/trycourier/courier-go/v4/internal/apijson"
	"github.com/trycourier/courier-go/v4/internal/apiquery"
	shimjson "github.com/trycourier/courier-go/v4/internal/encoding/json"
	"github.com/trycourier/courier-go/v4/internal/requestconfig"
	"github.com/trycourier/courier-go/v4/option"
	"github.com/trycourier/courier-go/v4/packages/param"
	"github.com/trycourier/courier-go/v4/packages/respjson"
	"github.com/trycourier/courier-go/v4/shared"
)

// RoutingStrategyService contains methods and other services that help with
// interacting with the Courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRoutingStrategyService] method instead.
type RoutingStrategyService struct {
	Options []option.RequestOption
}

// NewRoutingStrategyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRoutingStrategyService(opts ...option.RequestOption) (r RoutingStrategyService) {
	r = RoutingStrategyService{}
	r.Options = opts
	return
}

// Create a routing strategy. Requires a name and routing configuration at minimum.
// Channels and providers default to empty if omitted.
func (r *RoutingStrategyService) New(ctx context.Context, body RoutingStrategyNewParams, opts ...option.RequestOption) (res *RoutingStrategyMutationResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "routing-strategies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a routing strategy by ID. Returns the full entity including routing
// content and metadata.
func (r *RoutingStrategyService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *RoutingStrategyGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("routing-strategies/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List routing strategies in your workspace. Returns metadata only (no
// routing/channels/providers content). Use GET /routing-strategies/{id} for full
// details.
func (r *RoutingStrategyService) List(ctx context.Context, query RoutingStrategyListParams, opts ...option.RequestOption) (res *RoutingStrategyListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "routing-strategies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Archive a routing strategy. The strategy must not have associated notification
// templates. Unlink all templates before archiving.
func (r *RoutingStrategyService) Archive(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("routing-strategies/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Replace a routing strategy. Full document replacement; the caller must send the
// complete desired state. Missing optional fields are cleared.
func (r *RoutingStrategyService) Replace(ctx context.Context, id string, body RoutingStrategyReplaceParams, opts ...option.RequestOption) (res *RoutingStrategyMutationResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("routing-strategies/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Request body for creating a routing strategy.
//
// The properties Name, Routing are required.
type RoutingStrategyCreateRequestParam struct {
	// Human-readable name for the routing strategy.
	Name string `json:"name" api:"required"`
	// Routing tree defining channel selection method and order.
	Routing shared.MessageRoutingParam `json:"routing,omitzero" api:"required"`
	// Optional description of the routing strategy.
	Description param.Opt[string] `json:"description,omitzero"`
	// Optional tags for categorization.
	Tags []string `json:"tags,omitzero"`
	// Per-channel delivery configuration. Defaults to empty if omitted.
	Channels shared.MessageChannelsParam `json:"channels,omitzero"`
	// Per-provider delivery configuration. Defaults to empty if omitted.
	Providers shared.MessageProvidersParam `json:"providers,omitzero"`
	paramObj
}

func (r RoutingStrategyCreateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow RoutingStrategyCreateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RoutingStrategyCreateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Full routing strategy entity returned by GET.
type RoutingStrategyGetResponse struct {
	// The routing strategy ID (rs\_ prefix).
	ID string `json:"id" api:"required"`
	// Per-channel delivery configuration. May be empty.
	Channels shared.MessageChannels `json:"channels" api:"required"`
	// Epoch milliseconds when the strategy was created.
	Created int64 `json:"created" api:"required"`
	// User ID of the creator.
	Creator string `json:"creator" api:"required"`
	// Human-readable name.
	Name string `json:"name" api:"required"`
	// Per-provider delivery configuration. May be empty.
	Providers shared.MessageProviders `json:"providers" api:"required"`
	// Routing tree defining channel selection method and order.
	Routing shared.MessageRouting `json:"routing" api:"required"`
	// Description of the routing strategy.
	Description string `json:"description" api:"nullable"`
	// Tags for categorization.
	Tags []string `json:"tags" api:"nullable"`
	// Epoch milliseconds of last update.
	Updated int64 `json:"updated" api:"nullable"`
	// User ID of the last updater.
	Updater string `json:"updater" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Channels    respjson.Field
		Created     respjson.Field
		Creator     respjson.Field
		Name        respjson.Field
		Providers   respjson.Field
		Routing     respjson.Field
		Description respjson.Field
		Tags        respjson.Field
		Updated     respjson.Field
		Updater     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoutingStrategyGetResponse) RawJSON() string { return r.JSON.raw }
func (r *RoutingStrategyGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Paginated list of routing strategy summaries.
type RoutingStrategyListResponse struct {
	Paging  shared.Paging            `json:"paging" api:"required"`
	Results []RoutingStrategySummary `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Paging      respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoutingStrategyListResponse) RawJSON() string { return r.JSON.raw }
func (r *RoutingStrategyListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response returned by create and replace operations.
type RoutingStrategyMutationResponse struct {
	// The routing strategy ID (rs\_ prefix).
	ID string `json:"id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoutingStrategyMutationResponse) RawJSON() string { return r.JSON.raw }
func (r *RoutingStrategyMutationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body for replacing a routing strategy. Full document replacement;
// missing optional fields are cleared.
//
// The properties Name, Routing are required.
type RoutingStrategyReplaceRequestParam struct {
	// Human-readable name for the routing strategy.
	Name string `json:"name" api:"required"`
	// Routing tree defining channel selection method and order.
	Routing shared.MessageRoutingParam `json:"routing,omitzero" api:"required"`
	// Optional description. Omit or null to clear.
	Description param.Opt[string] `json:"description,omitzero"`
	// Optional tags. Omit or null to clear.
	Tags []string `json:"tags,omitzero"`
	// Per-channel delivery configuration. Omit to clear.
	Channels shared.MessageChannelsParam `json:"channels,omitzero"`
	// Per-provider delivery configuration. Omit to clear.
	Providers shared.MessageProvidersParam `json:"providers,omitzero"`
	paramObj
}

func (r RoutingStrategyReplaceRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow RoutingStrategyReplaceRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RoutingStrategyReplaceRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Routing strategy metadata returned in list responses. Does not include
// routing/channels/providers content.
type RoutingStrategySummary struct {
	// The routing strategy ID (rs\_ prefix).
	ID string `json:"id" api:"required"`
	// Epoch milliseconds when the strategy was created.
	Created int64 `json:"created" api:"required"`
	// User ID of the creator.
	Creator string `json:"creator" api:"required"`
	// Human-readable name.
	Name string `json:"name" api:"required"`
	// Description of the routing strategy.
	Description string `json:"description" api:"nullable"`
	// Tags for categorization.
	Tags []string `json:"tags" api:"nullable"`
	// Epoch milliseconds of last update.
	Updated int64 `json:"updated" api:"nullable"`
	// User ID of the last updater.
	Updater string `json:"updater" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Created     respjson.Field
		Creator     respjson.Field
		Name        respjson.Field
		Description respjson.Field
		Tags        respjson.Field
		Updated     respjson.Field
		Updater     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoutingStrategySummary) RawJSON() string { return r.JSON.raw }
func (r *RoutingStrategySummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoutingStrategyNewParams struct {
	// Request body for creating a routing strategy.
	RoutingStrategyCreateRequest RoutingStrategyCreateRequestParam
	paramObj
}

func (r RoutingStrategyNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.RoutingStrategyCreateRequest)
}
func (r *RoutingStrategyNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.RoutingStrategyCreateRequest)
}

type RoutingStrategyListParams struct {
	// Opaque pagination cursor from a previous response. Omit for the first page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of results per page. Default 20, max 100.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RoutingStrategyListParams]'s query parameters as
// `url.Values`.
func (r RoutingStrategyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RoutingStrategyReplaceParams struct {
	// Request body for replacing a routing strategy. Full document replacement;
	// missing optional fields are cleared.
	RoutingStrategyReplaceRequest RoutingStrategyReplaceRequestParam
	paramObj
}

func (r RoutingStrategyReplaceParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.RoutingStrategyReplaceRequest)
}
func (r *RoutingStrategyReplaceParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.RoutingStrategyReplaceRequest)
}
