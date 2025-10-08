// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/trycourier/courier-go/internal/apijson"
	shimjson "github.com/trycourier/courier-go/internal/encoding/json"
	"github.com/trycourier/courier-go/internal/requestconfig"
	"github.com/trycourier/courier-go/option"
	"github.com/trycourier/courier-go/packages/param"
	"github.com/trycourier/courier-go/packages/respjson"
	"github.com/trycourier/courier-go/shared"
)

// UserTokenService contains methods and other services that help with interacting
// with the courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserTokenService] method instead.
type UserTokenService struct {
	Options []option.RequestOption
}

// NewUserTokenService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserTokenService(opts ...option.RequestOption) (r UserTokenService) {
	r = UserTokenService{}
	r.Options = opts
	return
}

// Get single token available for a `:token`
func (r *UserTokenService) Get(ctx context.Context, token string, query UserTokenGetParams, opts ...option.RequestOption) (res *UserTokenGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.UserID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	if token == "" {
		err = errors.New("missing required token parameter")
		return
	}
	path := fmt.Sprintf("users/%s/tokens/%s", query.UserID, token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Apply a JSON Patch (RFC 6902) to the specified token.
func (r *UserTokenService) Update(ctx context.Context, token string, params UserTokenUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if params.UserID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	if token == "" {
		err = errors.New("missing required token parameter")
		return
	}
	path := fmt.Sprintf("users/%s/tokens/%s", params.UserID, token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, nil, opts...)
	return
}

// Gets all tokens available for a :user_id
func (r *UserTokenService) List(ctx context.Context, userID string, opts ...option.RequestOption) (res *[]shared.UserToken, err error) {
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("users/%s/tokens", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete User Token
func (r *UserTokenService) Delete(ctx context.Context, token string, body UserTokenDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.UserID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	if token == "" {
		err = errors.New("missing required token parameter")
		return
	}
	path := fmt.Sprintf("users/%s/tokens/%s", body.UserID, token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Adds multiple tokens to a user and overwrites matching existing tokens.
func (r *UserTokenService) AddMultiple(ctx context.Context, userID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("users/%s/tokens", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, nil, opts...)
	return
}

// Adds a single token to a user and overwrites a matching existing token.
func (r *UserTokenService) AddSingle(ctx context.Context, token string, params UserTokenAddSingleParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if params.UserID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	if token == "" {
		err = errors.New("missing required token parameter")
		return
	}
	path := fmt.Sprintf("users/%s/tokens/%s", params.UserID, token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, nil, opts...)
	return
}

type UserTokenGetResponse struct {
	// Any of "active", "unknown", "failed", "revoked".
	Status string `json:"status,nullable"`
	// The reason for the token status.
	StatusReason string `json:"status_reason,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status       respjson.Field
		StatusReason respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
	shared.UserToken
}

// Returns the unmodified JSON received from the API
func (r UserTokenGetResponse) RawJSON() string { return r.JSON.raw }
func (r *UserTokenGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenGetParams struct {
	UserID string `path:"user_id,required" json:"-"`
	paramObj
}

type UserTokenUpdateParams struct {
	UserID string                       `path:"user_id,required" json:"-"`
	Patch  []UserTokenUpdateParamsPatch `json:"patch,omitzero,required"`
	paramObj
}

func (r UserTokenUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow UserTokenUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserTokenUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Op, Path are required.
type UserTokenUpdateParamsPatch struct {
	// The operation to perform.
	Op string `json:"op,required"`
	// The JSON path specifying the part of the profile to operate on.
	Path string `json:"path,required"`
	// The value for the operation.
	Value param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r UserTokenUpdateParamsPatch) MarshalJSON() (data []byte, err error) {
	type shadow UserTokenUpdateParamsPatch
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserTokenUpdateParamsPatch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenDeleteParams struct {
	UserID string `path:"user_id,required" json:"-"`
	paramObj
}

type UserTokenAddSingleParams struct {
	UserID    string `path:"user_id,required" json:"-"`
	UserToken shared.UserTokenParam
	paramObj
}

func (r UserTokenAddSingleParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UserToken)
}
func (r *UserTokenAddSingleParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.UserToken)
}
