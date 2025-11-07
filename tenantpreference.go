// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"github.com/trycourier/courier-go/v3/option"
)

// TenantPreferenceService contains methods and other services that help with
// interacting with the Courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTenantPreferenceService] method instead.
type TenantPreferenceService struct {
	Options []option.RequestOption
	Items   TenantPreferenceItemService
}

// NewTenantPreferenceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTenantPreferenceService(opts ...option.RequestOption) (r TenantPreferenceService) {
	r = TenantPreferenceService{}
	r.Options = opts
	r.Items = NewTenantPreferenceItemService(opts...)
	return
}
