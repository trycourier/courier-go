// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"github.com/trycourier/courier-go/v3/option"
)

// TenantTenantDefaultPreferenceService contains methods and other services that
// help with interacting with the courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTenantTenantDefaultPreferenceService] method instead.
type TenantTenantDefaultPreferenceService struct {
	Options []option.RequestOption
	Items   TenantTenantDefaultPreferenceItemService
}

// NewTenantTenantDefaultPreferenceService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewTenantTenantDefaultPreferenceService(opts ...option.RequestOption) (r TenantTenantDefaultPreferenceService) {
	r = TenantTenantDefaultPreferenceService{}
	r.Options = opts
	r.Items = NewTenantTenantDefaultPreferenceItemService(opts...)
	return
}
