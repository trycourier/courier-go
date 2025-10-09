// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
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
	"github.com/trycourier/courier-go/shared"
)

// ListService contains methods and other services that help with interacting with
// the courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewListService] method instead.
type ListService struct {
	Options       []option.RequestOption
	Subscriptions ListSubscriptionService
}

// NewListService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewListService(opts ...option.RequestOption) (r ListService) {
	r = ListService{}
	r.Options = opts
	r.Subscriptions = NewListSubscriptionService(opts...)
	return
}

// Returns a list based on the list ID provided.
func (r *ListService) Get(ctx context.Context, listID string, opts ...option.RequestOption) (res *shared.UserList, err error) {
	opts = slices.Concat(r.Options, opts)
	if listID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
	path := fmt.Sprintf("lists/%s", listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Create or replace an existing list with the supplied values.
func (r *ListService) Update(ctx context.Context, listID string, body ListUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if listID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
	path := fmt.Sprintf("lists/%s", listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Returns all of the lists, with the ability to filter based on a pattern.
func (r *ListService) List(ctx context.Context, query ListListParams, opts ...option.RequestOption) (res *ListListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "lists"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a list by list ID.
func (r *ListService) Delete(ctx context.Context, listID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if listID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
	path := fmt.Sprintf("lists/%s", listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Restore a previously deleted list.
func (r *ListService) Restore(ctx context.Context, listID string, body ListRestoreParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if listID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
	path := fmt.Sprintf("lists/%s/restore", listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

type ListListResponse struct {
	Items  []shared.UserList `json:"items,required"`
	Paging shared.Paging     `json:"paging,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListListResponse) RawJSON() string { return r.JSON.raw }
func (r *ListListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListUpdateParams struct {
	Name        string                           `json:"name,required"`
	Preferences shared.RecipientPreferencesParam `json:"preferences,omitzero"`
	paramObj
}

func (r ListUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ListUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ListUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListListParams struct {
	// A unique identifier that allows for fetching the next page of lists.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// "A pattern used to filter the list items returned. Pattern types supported:
	// exact match on `list_id` or a pattern of one or more pattern parts. you may
	// replace a part with either: `*` to match all parts in that position, or `**` to
	// signify a wildcard `endsWith` pattern match."
	Pattern param.Opt[string] `query:"pattern,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ListListParams]'s query parameters as `url.Values`.
func (r ListListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ListRestoreParams struct {
	paramObj
}

func (r ListRestoreParams) MarshalJSON() (data []byte, err error) {
	type shadow ListRestoreParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ListRestoreParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
