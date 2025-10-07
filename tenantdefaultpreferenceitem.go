// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/courier-go/internal/apijson"
	shimjson "github.com/stainless-sdks/courier-go/internal/encoding/json"
	"github.com/stainless-sdks/courier-go/internal/requestconfig"
	"github.com/stainless-sdks/courier-go/option"
	"github.com/stainless-sdks/courier-go/packages/param"
	"github.com/stainless-sdks/courier-go/packages/respjson"
)

// TenantDefaultPreferenceItemService contains methods and other services that help
// with interacting with the courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTenantDefaultPreferenceItemService] method instead.
type TenantDefaultPreferenceItemService struct {
	Options []option.RequestOption
}

// NewTenantDefaultPreferenceItemService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewTenantDefaultPreferenceItemService(opts ...option.RequestOption) (r TenantDefaultPreferenceItemService) {
	r = TenantDefaultPreferenceItemService{}
	r.Options = opts
	return
}

// Create or Replace Default Preferences For Topic
func (r *TenantDefaultPreferenceItemService) Update(ctx context.Context, topicID string, params TenantDefaultPreferenceItemUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if params.TenantID == "" {
		err = errors.New("missing required tenant_id parameter")
		return
	}
	if topicID == "" {
		err = errors.New("missing required topic_id parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s/default_preferences/items/%s", params.TenantID, topicID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, nil, opts...)
	return
}

// Remove Default Preferences For Topic
func (r *TenantDefaultPreferenceItemService) Delete(ctx context.Context, topicID string, body TenantDefaultPreferenceItemDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.TenantID == "" {
		err = errors.New("missing required tenant_id parameter")
		return
	}
	if topicID == "" {
		err = errors.New("missing required topic_id parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s/default_preferences/items/%s", body.TenantID, topicID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type ChannelClassification string

const (
	ChannelClassificationDirectMessage ChannelClassification = "direct_message"
	ChannelClassificationEmail         ChannelClassification = "email"
	ChannelClassificationPush          ChannelClassification = "push"
	ChannelClassificationSMS           ChannelClassification = "sms"
	ChannelClassificationWebhook       ChannelClassification = "webhook"
	ChannelClassificationInbox         ChannelClassification = "inbox"
)

type SubscriptionTopicNew struct {
	// Any of "OPTED_OUT", "OPTED_IN", "REQUIRED".
	Status SubscriptionTopicNewStatus `json:"status,required"`
	// The default channels to send to this tenant when has_custom_routing is enabled
	CustomRouting []ChannelClassification `json:"custom_routing,nullable"`
	// Override channel routing with custom preferences. This will override any
	// template prefernces that are set, but a user can still customize their
	// preferences
	HasCustomRouting bool `json:"has_custom_routing,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status           respjson.Field
		CustomRouting    respjson.Field
		HasCustomRouting respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionTopicNew) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionTopicNew) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SubscriptionTopicNew to a SubscriptionTopicNewParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SubscriptionTopicNewParam.Overrides()
func (r SubscriptionTopicNew) ToParam() SubscriptionTopicNewParam {
	return param.Override[SubscriptionTopicNewParam](json.RawMessage(r.RawJSON()))
}

type SubscriptionTopicNewStatus string

const (
	SubscriptionTopicNewStatusOptedOut SubscriptionTopicNewStatus = "OPTED_OUT"
	SubscriptionTopicNewStatusOptedIn  SubscriptionTopicNewStatus = "OPTED_IN"
	SubscriptionTopicNewStatusRequired SubscriptionTopicNewStatus = "REQUIRED"
)

// The property Status is required.
type SubscriptionTopicNewParam struct {
	// Any of "OPTED_OUT", "OPTED_IN", "REQUIRED".
	Status SubscriptionTopicNewStatus `json:"status,omitzero,required"`
	// Override channel routing with custom preferences. This will override any
	// template prefernces that are set, but a user can still customize their
	// preferences
	HasCustomRouting param.Opt[bool] `json:"has_custom_routing,omitzero"`
	// The default channels to send to this tenant when has_custom_routing is enabled
	CustomRouting []ChannelClassification `json:"custom_routing,omitzero"`
	paramObj
}

func (r SubscriptionTopicNewParam) MarshalJSON() (data []byte, err error) {
	type shadow SubscriptionTopicNewParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubscriptionTopicNewParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantDefaultPreferenceItemUpdateParams struct {
	TenantID             string `path:"tenant_id,required" json:"-"`
	SubscriptionTopicNew SubscriptionTopicNewParam
	paramObj
}

func (r TenantDefaultPreferenceItemUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SubscriptionTopicNew)
}
func (r *TenantDefaultPreferenceItemUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.SubscriptionTopicNew)
}

type TenantDefaultPreferenceItemDeleteParams struct {
	TenantID string `path:"tenant_id,required" json:"-"`
	paramObj
}
