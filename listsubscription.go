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

	"github.com/stainless-sdks/courier-go/internal/apijson"
	"github.com/stainless-sdks/courier-go/internal/apiquery"
	"github.com/stainless-sdks/courier-go/internal/requestconfig"
	"github.com/stainless-sdks/courier-go/option"
	"github.com/stainless-sdks/courier-go/packages/param"
	"github.com/stainless-sdks/courier-go/packages/respjson"
)

// ListSubscriptionService contains methods and other services that help with
// interacting with the courier API.
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
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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

// The property RecipientID is required.
type PutSubscriptionsRecipientParam struct {
	RecipientID string                    `json:"recipientId,required"`
	Preferences RecipientPreferencesParam `json:"preferences,omitzero"`
	paramObj
}

func (r PutSubscriptionsRecipientParam) MarshalJSON() (data []byte, err error) {
	type shadow PutSubscriptionsRecipientParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PutSubscriptionsRecipientParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecipientPreferences struct {
	Categories    map[string]RecipientPreferencesCategory     `json:"categories,nullable"`
	Notifications map[string]RecipientPreferencesNotification `json:"notifications,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Categories    respjson.Field
		Notifications respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecipientPreferences) RawJSON() string { return r.JSON.raw }
func (r *RecipientPreferences) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this RecipientPreferences to a RecipientPreferencesParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RecipientPreferencesParam.Overrides()
func (r RecipientPreferences) ToParam() RecipientPreferencesParam {
	return param.Override[RecipientPreferencesParam](json.RawMessage(r.RawJSON()))
}

type RecipientPreferencesCategory struct {
	// Any of "OPTED_IN", "OPTED_OUT", "REQUIRED".
	Status             PreferenceStatus                                `json:"status,required"`
	ChannelPreferences []RecipientPreferencesCategoryChannelPreference `json:"channel_preferences,nullable"`
	Rules              []RecipientPreferencesCategoryRule              `json:"rules,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status             respjson.Field
		ChannelPreferences respjson.Field
		Rules              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecipientPreferencesCategory) RawJSON() string { return r.JSON.raw }
func (r *RecipientPreferencesCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecipientPreferencesCategoryChannelPreference struct {
	// Any of "direct_message", "email", "push", "sms", "webhook", "inbox".
	Channel ChannelClassification `json:"channel,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Channel     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecipientPreferencesCategoryChannelPreference) RawJSON() string { return r.JSON.raw }
func (r *RecipientPreferencesCategoryChannelPreference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecipientPreferencesCategoryRule struct {
	Until string `json:"until,required"`
	Start string `json:"start,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Until       respjson.Field
		Start       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecipientPreferencesCategoryRule) RawJSON() string { return r.JSON.raw }
func (r *RecipientPreferencesCategoryRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecipientPreferencesNotification struct {
	// Any of "OPTED_IN", "OPTED_OUT", "REQUIRED".
	Status             PreferenceStatus                                    `json:"status,required"`
	ChannelPreferences []RecipientPreferencesNotificationChannelPreference `json:"channel_preferences,nullable"`
	Rules              []RecipientPreferencesNotificationRule              `json:"rules,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status             respjson.Field
		ChannelPreferences respjson.Field
		Rules              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecipientPreferencesNotification) RawJSON() string { return r.JSON.raw }
func (r *RecipientPreferencesNotification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecipientPreferencesNotificationChannelPreference struct {
	// Any of "direct_message", "email", "push", "sms", "webhook", "inbox".
	Channel ChannelClassification `json:"channel,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Channel     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecipientPreferencesNotificationChannelPreference) RawJSON() string { return r.JSON.raw }
func (r *RecipientPreferencesNotificationChannelPreference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecipientPreferencesNotificationRule struct {
	Until string `json:"until,required"`
	Start string `json:"start,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Until       respjson.Field
		Start       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecipientPreferencesNotificationRule) RawJSON() string { return r.JSON.raw }
func (r *RecipientPreferencesNotificationRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecipientPreferencesParam struct {
	Categories    map[string]RecipientPreferencesCategoryParam     `json:"categories,omitzero"`
	Notifications map[string]RecipientPreferencesNotificationParam `json:"notifications,omitzero"`
	paramObj
}

func (r RecipientPreferencesParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientPreferencesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecipientPreferencesParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Status is required.
type RecipientPreferencesCategoryParam struct {
	// Any of "OPTED_IN", "OPTED_OUT", "REQUIRED".
	Status             PreferenceStatus                                     `json:"status,omitzero,required"`
	ChannelPreferences []RecipientPreferencesCategoryChannelPreferenceParam `json:"channel_preferences,omitzero"`
	Rules              []RecipientPreferencesCategoryRuleParam              `json:"rules,omitzero"`
	paramObj
}

func (r RecipientPreferencesCategoryParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientPreferencesCategoryParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecipientPreferencesCategoryParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Channel is required.
type RecipientPreferencesCategoryChannelPreferenceParam struct {
	// Any of "direct_message", "email", "push", "sms", "webhook", "inbox".
	Channel ChannelClassification `json:"channel,omitzero,required"`
	paramObj
}

func (r RecipientPreferencesCategoryChannelPreferenceParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientPreferencesCategoryChannelPreferenceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecipientPreferencesCategoryChannelPreferenceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Until is required.
type RecipientPreferencesCategoryRuleParam struct {
	Until string            `json:"until,required"`
	Start param.Opt[string] `json:"start,omitzero"`
	paramObj
}

func (r RecipientPreferencesCategoryRuleParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientPreferencesCategoryRuleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecipientPreferencesCategoryRuleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Status is required.
type RecipientPreferencesNotificationParam struct {
	// Any of "OPTED_IN", "OPTED_OUT", "REQUIRED".
	Status             PreferenceStatus                                         `json:"status,omitzero,required"`
	ChannelPreferences []RecipientPreferencesNotificationChannelPreferenceParam `json:"channel_preferences,omitzero"`
	Rules              []RecipientPreferencesNotificationRuleParam              `json:"rules,omitzero"`
	paramObj
}

func (r RecipientPreferencesNotificationParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientPreferencesNotificationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecipientPreferencesNotificationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Channel is required.
type RecipientPreferencesNotificationChannelPreferenceParam struct {
	// Any of "direct_message", "email", "push", "sms", "webhook", "inbox".
	Channel ChannelClassification `json:"channel,omitzero,required"`
	paramObj
}

func (r RecipientPreferencesNotificationChannelPreferenceParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientPreferencesNotificationChannelPreferenceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecipientPreferencesNotificationChannelPreferenceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Until is required.
type RecipientPreferencesNotificationRuleParam struct {
	Until string            `json:"until,required"`
	Start param.Opt[string] `json:"start,omitzero"`
	paramObj
}

func (r RecipientPreferencesNotificationRuleParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientPreferencesNotificationRuleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecipientPreferencesNotificationRuleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListSubscriptionListResponse struct {
	Items  []ListSubscriptionListResponseItem `json:"items,required"`
	Paging Paging                             `json:"paging,required"`
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
	RecipientID string               `json:"recipientId,required"`
	Created     string               `json:"created,nullable"`
	Preferences RecipientPreferences `json:"preferences,nullable"`
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
	ListID      string                    `path:"list_id,required" json:"-"`
	Preferences RecipientPreferencesParam `json:"preferences,omitzero"`
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
