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

// ListSubscriptionService contains methods and other services that help with
// interacting with the Courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewListSubscriptionService] method instead.
type ListSubscriptionService struct {
	Options []option.RequestOption
}

// NewListSubscriptionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewListSubscriptionService(opts ...option.RequestOption) (r ListSubscriptionService) {
	r = ListSubscriptionService{}
	r.Options = opts
	return
}

// Get the list's subscriptions.
func (r *ListSubscriptionService) List(ctx context.Context, listID string, query ListSubscriptionListParams, opts ...option.RequestOption) (res *ListSubscriptionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if listID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
	path := fmt.Sprintf("lists/%s/subscriptions", listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Subscribes additional users to the list, without modifying existing
// subscriptions. If the list does not exist, it will be automatically created.
func (r *ListSubscriptionService) Add(ctx context.Context, listID string, body ListSubscriptionAddParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if listID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
	path := fmt.Sprintf("lists/%s/subscriptions", listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Subscribes the users to the list, overwriting existing subscriptions. If the
// list does not exist, it will be automatically created.
func (r *ListSubscriptionService) Subscribe(ctx context.Context, listID string, body ListSubscriptionSubscribeParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if listID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
	path := fmt.Sprintf("lists/%s/subscriptions", listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Subscribe a user to an existing list (note: if the List does not exist, it will
// be automatically created).
func (r *ListSubscriptionService) SubscribeUser(ctx context.Context, userID string, params ListSubscriptionSubscribeUserParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.ListID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("lists/%s/subscriptions/%s", params.ListID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, nil, opts...)
	return
}

// Delete a subscription to a list by list ID and user ID.
func (r *ListSubscriptionService) UnsubscribeUser(ctx context.Context, userID string, body ListSubscriptionUnsubscribeUserParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ListID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("lists/%s/subscriptions/%s", body.ListID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type ListSubscriptionListResponse struct {
	Items  []ListSubscriptionListResponseItem `json:"items,required"`
	Paging shared.Paging                      `json:"paging,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListSubscriptionListResponse) RawJSON() string { return r.JSON.raw }
func (r *ListSubscriptionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListSubscriptionListResponseItem struct {
	RecipientID string                      `json:"recipientId,required"`
	Created     string                      `json:"created,nullable"`
	Preferences shared.RecipientPreferences `json:"preferences,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RecipientID respjson.Field
		Created     respjson.Field
		Preferences respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListSubscriptionListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *ListSubscriptionListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListSubscriptionListParams struct {
	// A unique identifier that allows for fetching the next set of list subscriptions
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ListSubscriptionListParams]'s query parameters as
// `url.Values`.
func (r ListSubscriptionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ListSubscriptionAddParams struct {
	Recipients []PutSubscriptionsRecipientParam `json:"recipients,omitzero,required"`
	paramObj
}

func (r ListSubscriptionAddParams) MarshalJSON() (data []byte, err error) {
	type shadow ListSubscriptionAddParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ListSubscriptionAddParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListSubscriptionSubscribeParams struct {
	Recipients []PutSubscriptionsRecipientParam `json:"recipients,omitzero,required"`
	paramObj
}

func (r ListSubscriptionSubscribeParams) MarshalJSON() (data []byte, err error) {
	type shadow ListSubscriptionSubscribeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ListSubscriptionSubscribeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListSubscriptionSubscribeUserParams struct {
	ListID      string                           `path:"list_id,required" json:"-"`
	Preferences shared.RecipientPreferencesParam `json:"preferences,omitzero"`
	paramObj
}

func (r ListSubscriptionSubscribeUserParams) MarshalJSON() (data []byte, err error) {
	type shadow ListSubscriptionSubscribeUserParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ListSubscriptionSubscribeUserParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListSubscriptionUnsubscribeUserParams struct {
	ListID string `path:"list_id,required" json:"-"`
	paramObj
}
