// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	shimjson "github.com/trycourier/courier-go/v3/internal/encoding/json"
	"github.com/trycourier/courier-go/v3/internal/requestconfig"
	"github.com/trycourier/courier-go/v3/option"
	"github.com/trycourier/courier-go/v3/shared"
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

type TenantDefaultPreferenceItemUpdateParams struct {
	TenantID             string `path:"tenant_id,required" json:"-"`
	SubscriptionTopicNew shared.SubscriptionTopicNewParam
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
