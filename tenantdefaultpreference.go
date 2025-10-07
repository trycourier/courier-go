// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"github.com/stainless-sdks/courier-go/option"
)

// TenantDefaultPreferenceService contains methods and other services that help
// with interacting with the courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTenantDefaultPreferenceService] method instead.
type TenantDefaultPreferenceService struct {
	Options []option.RequestOption
	Items   TenantDefaultPreferenceItemService
}

// NewTenantDefaultPreferenceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTenantDefaultPreferenceService(opts ...option.RequestOption) (r TenantDefaultPreferenceService) {
	r = TenantDefaultPreferenceService{}
	r.Options = opts
	r.Items = NewTenantDefaultPreferenceItemService(opts...)
	return
}
