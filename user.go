// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"github.com/trycourier/courier-go/v4/option"
)

// UserService contains methods and other services that help with interacting with
// the Courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options []option.RequestOption
	// Read and write a single user's notification preferences, per topic and per
	// channel.
	Preferences UserPreferenceService
	// Associate a user with one or more tenants, and read or remove those
	// associations.
	Tenants UserTenantService
	// Register and manage the APNS and FCM device tokens Courier delivers push
	// notifications to.
	Tokens UserTokenService
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
