// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"github.com/stainless-sdks/courier-go/option"
)

// UserService contains methods and other services that help with interacting with
// the courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options     []option.RequestOption
	Preferences UserPreferenceService
	Tenants     UserTenantService
	Tokens      UserTokenService
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r UserService) {
	r = UserService{}
	r.Options = opts
	r.Preferences = NewUserPreferenceService(opts...)
	r.Tenants = NewUserTenantService(opts...)
	r.Tokens = NewUserTokenService(opts...)
	return
}
