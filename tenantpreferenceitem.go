// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	shimjson "github.com/trycourier/courier-go/v4/internal/encoding/json"
	"github.com/trycourier/courier-go/v4/internal/requestconfig"
	"github.com/trycourier/courier-go/v4/option"
)

// TenantPreferenceItemService contains methods and other services that help with
// interacting with the Courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTenantPreferenceItemService] method instead.
type TenantPreferenceItemService struct {
	Options []option.RequestOption
}

// NewTenantPreferenceItemService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTenantPreferenceItemService(opts ...option.RequestOption) (r TenantPreferenceItemService) {
	r = TenantPreferenceItemService{}
	r.Options = opts
	return
}

// Create or Replace Default Preferences For Topic
func (r *TenantPreferenceItemService) Update(ctx context.Context, topicID string, params TenantPreferenceItemUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
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
func (r *TenantPreferenceItemService) Delete(ctx context.Context, topicID string, body TenantPreferenceItemDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
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

type TenantPreferenceItemUpdateParams struct {
	TenantID             string `path:"tenant_id,required" json:"-"`
	SubscriptionTopicNew SubscriptionTopicNewParam
	paramObj
}

func (r TenantPreferenceItemUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SubscriptionTopicNew)
}
func (r *TenantPreferenceItemUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.SubscriptionTopicNew)
}

type TenantPreferenceItemDeleteParams struct {
	TenantID string `path:"tenant_id,required" json:"-"`
	paramObj
}
