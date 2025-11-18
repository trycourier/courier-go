// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
	"net/http"
	"slices"

	"github.com/trycourier/courier-go/v4/internal/apijson"
	"github.com/trycourier/courier-go/v4/internal/requestconfig"
	"github.com/trycourier/courier-go/v4/option"
	"github.com/trycourier/courier-go/v4/packages/param"
	"github.com/trycourier/courier-go/v4/packages/respjson"
)

// AuthService contains methods and other services that help with interacting with
// the Courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthService] method instead.
type AuthService struct {
	Options []option.RequestOption
}

// NewAuthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAuthService(opts ...option.RequestOption) (r AuthService) {
	r = AuthService{}
	r.Options = opts
	return
}

// Returns a new access token.
func (r *AuthService) IssueToken(ctx context.Context, body AuthIssueTokenParams, opts ...option.RequestOption) (res *AuthIssueTokenResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "auth/issue-token"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AuthIssueTokenResponse struct {
	Token string `json:"token,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthIssueTokenResponse) RawJSON() string { return r.JSON.raw }
func (r *AuthIssueTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthIssueTokenParams struct {
	// Duration for token expiration. Accepts various time formats:
	//
	// - "2 hours" - 2 hours from now
	// - "1d" - 1 day
	// - "3 days" - 3 days
	// - "10h" - 10 hours
	// - "2.5 hrs" - 2.5 hours
	// - "1m" - 1 minute
	// - "5s" - 5 seconds
	// - "1y" - 1 year
	ExpiresIn string `json:"expires_in,required"`
	// Available scopes:
	//
	//   - `user_id:<user-id>` - Defines which user the token will be scoped to. Multiple
	//     can be listed if needed. Ex `user_id:pigeon user_id:bluebird`.
	//   - `read:messages` - Read messages.
	//   - `read:user-tokens` - Read user push tokens.
	//   - `write:user-tokens` - Write user push tokens.
	//   - `read:brands[:<brand_id>]` - Read brands, optionally restricted to a specific
	//     brand_id. Examples `read:brands`, `read:brands:my_brand`.
	//   - `write:brands[:<brand_id>]` - Write brands, optionally restricted to a
	//     specific brand_id. Examples `write:brands`, `write:brands:my_brand`.
	//   - `inbox:read:messages` - Read inbox messages.
	//   - `inbox:write:events` - Write inbox events, such as mark message as read.
	//   - `read:preferences` - Read user preferences.
	//   - `write:preferences` - Write user preferences. Example:
	//     `user_id:user123 write:user-tokens inbox:read:messages inbox:write:events read:preferences write:preferences read:brands`
	Scope string `json:"scope,required"`
	paramObj
}

func (r AuthIssueTokenParams) MarshalJSON() (data []byte, err error) {
	type shadow AuthIssueTokenParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AuthIssueTokenParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
