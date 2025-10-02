// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
	"net/http"
	"slices"

	"github.com/trycourier/courier-go/internal/apijson"
	"github.com/trycourier/courier-go/internal/requestconfig"
	"github.com/trycourier/courier-go/option"
	"github.com/trycourier/courier-go/packages/param"
	"github.com/trycourier/courier-go/packages/respjson"
)

// AuthService contains methods and other services that help with interacting with
// the courier API.
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
	Token string `json:"token,nullable"`
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
	ExpiresIn string `json:"expires_in,required"`
	Scope     string `json:"scope,required"`
	paramObj
}

func (r AuthIssueTokenParams) MarshalJSON() (data []byte, err error) {
	type shadow AuthIssueTokenParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AuthIssueTokenParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
