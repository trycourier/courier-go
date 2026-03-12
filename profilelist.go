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

// ProfileListService contains methods and other services that help with
// interacting with the Courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProfileListService] method instead.
type ProfileListService struct {
	Options []option.RequestOption
}

// NewProfileListService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProfileListService(opts ...option.RequestOption) (r ProfileListService) {
	r = ProfileListService{}
	r.Options = opts
	return
}

// Returns the subscribed lists for a specified user.
func (r *ProfileListService) Get(ctx context.Context, userID string, query ProfileListGetParams, opts ...option.RequestOption) (res *ProfileListGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("profiles/%s/lists", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Removes all list subscriptions for given user.
func (r *ProfileListService) Delete(ctx context.Context, userID string, opts ...option.RequestOption) (res *ProfileListDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("profiles/%s/lists", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Subscribes the given user to one or more lists. If the list does not exist, it
// will be created.
func (r *ProfileListService) Subscribe(ctx context.Context, userID string, body ProfileListSubscribeParams, opts ...option.RequestOption) (res *ProfileListSubscribeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("profiles/%s/lists", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ProfileListGetResponse struct {
	Paging shared.Paging `json:"paging" api:"required"`
	// An array of lists
	Results []ProfileListGetResponseResult `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Paging      respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileListGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ProfileListGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileListGetResponseResult struct {
	ID string `json:"id" api:"required"`
	// The date/time of when the list was created. Represented as a string in ISO
	// format.
	Created string `json:"created" api:"required"`
	// List name
	Name string `json:"name" api:"required"`
	// The date/time of when the list was updated. Represented as a string in ISO
	// format.
	Updated     string                      `json:"updated" api:"required"`
	Preferences shared.RecipientPreferences `json:"preferences" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Created     respjson.Field
		Name        respjson.Field
		Updated     respjson.Field
		Preferences respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileListGetResponseResult) RawJSON() string { return r.JSON.raw }
func (r *ProfileListGetResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileListDeleteResponse struct {
	// Any of "SUCCESS".
	Status ProfileListDeleteResponseStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileListDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *ProfileListDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileListDeleteResponseStatus string

const (
	ProfileListDeleteResponseStatusSuccess ProfileListDeleteResponseStatus = "SUCCESS"
)

type ProfileListSubscribeResponse struct {
	// Any of "SUCCESS".
	Status ProfileListSubscribeResponseStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileListSubscribeResponse) RawJSON() string { return r.JSON.raw }
func (r *ProfileListSubscribeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileListSubscribeResponseStatus string

const (
	ProfileListSubscribeResponseStatusSuccess ProfileListSubscribeResponseStatus = "SUCCESS"
)

type ProfileListGetParams struct {
	// A unique identifier that allows for fetching the next set of message statuses.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ProfileListGetParams]'s query parameters as `url.Values`.
func (r ProfileListGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProfileListSubscribeParams struct {
	Lists []SubscribeToListsRequestItemParam `json:"lists,omitzero" api:"required"`
	paramObj
}

func (r ProfileListSubscribeParams) MarshalJSON() (data []byte, err error) {
	type shadow ProfileListSubscribeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileListSubscribeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
