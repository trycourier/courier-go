// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/trycourier/courier-go/v4/internal/apijson"
	"github.com/trycourier/courier-go/v4/internal/requestconfig"
	"github.com/trycourier/courier-go/v4/option"
	"github.com/trycourier/courier-go/v4/packages/param"
	"github.com/trycourier/courier-go/v4/packages/respjson"
)

// UserTokenService contains methods and other services that help with interacting
// with the Courier API.
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
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
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
func (r *UserTokenService) List(ctx context.Context, userID string, opts ...option.RequestOption) (res *UserTokenListResponse, err error) {
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
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
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
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
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
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
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

type UserToken struct {
	// Full body of the token. Must match token in URL path parameter.
	Token string `json:"token" api:"required"`
	// Any of "firebase-fcm", "apn", "expo", "onesignal".
	ProviderKey UserTokenProviderKey `json:"provider_key" api:"required"`
	// Information about the device the token came from.
	Device UserTokenDevice `json:"device" api:"nullable"`
	// ISO 8601 formatted date the token expires. Defaults to 2 months. Set to false to
	// disable expiration.
	ExpiryDate UserTokenExpiryDateUnion `json:"expiry_date" api:"nullable"`
	// Properties about the token.
	Properties any `json:"properties"`
	// Tracking information about the device the token came from.
	Tracking UserTokenTracking `json:"tracking" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		ProviderKey respjson.Field
		Device      respjson.Field
		ExpiryDate  respjson.Field
		Properties  respjson.Field
		Tracking    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserToken) RawJSON() string { return r.JSON.raw }
func (r *UserToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenProviderKey string

const (
	UserTokenProviderKeyFirebaseFcm UserTokenProviderKey = "firebase-fcm"
	UserTokenProviderKeyApn         UserTokenProviderKey = "apn"
	UserTokenProviderKeyExpo        UserTokenProviderKey = "expo"
	UserTokenProviderKeyOnesignal   UserTokenProviderKey = "onesignal"
)

// Information about the device the token came from.
type UserTokenDevice struct {
	// Id of the advertising identifier
	AdID string `json:"ad_id" api:"nullable"`
	// Id of the application the token is used for
	AppID string `json:"app_id" api:"nullable"`
	// Id of the device the token is associated with
	DeviceID string `json:"device_id" api:"nullable"`
	// The device manufacturer
	Manufacturer string `json:"manufacturer" api:"nullable"`
	// The device model
	Model string `json:"model" api:"nullable"`
	// The device platform i.e. android, ios, web
	Platform string `json:"platform" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdID         respjson.Field
		AppID        respjson.Field
		DeviceID     respjson.Field
		Manufacturer respjson.Field
		Model        respjson.Field
		Platform     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserTokenDevice) RawJSON() string { return r.JSON.raw }
func (r *UserTokenDevice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UserTokenExpiryDateUnion contains all possible properties and values from
// [string], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfBool]
type UserTokenExpiryDateUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u UserTokenExpiryDateUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UserTokenExpiryDateUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u UserTokenExpiryDateUnion) RawJSON() string { return u.JSON.raw }

func (r *UserTokenExpiryDateUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tracking information about the device the token came from.
type UserTokenTracking struct {
	// The IP address of the device
	IP string `json:"ip" api:"nullable"`
	// The latitude of the device
	Lat string `json:"lat" api:"nullable"`
	// The longitude of the device
	Long string `json:"long" api:"nullable"`
	// The operating system version
	OsVersion string `json:"os_version" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IP          respjson.Field
		Lat         respjson.Field
		Long        respjson.Field
		OsVersion   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserTokenTracking) RawJSON() string { return r.JSON.raw }
func (r *UserTokenTracking) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenGetResponse struct {
	// Any of "active", "unknown", "failed", "revoked".
	Status string `json:"status" api:"nullable"`
	// The reason for the token status.
	StatusReason string `json:"status_reason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status       respjson.Field
		StatusReason respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
	UserToken
}

// Returns the unmodified JSON received from the API
func (r UserTokenGetResponse) RawJSON() string { return r.JSON.raw }
func (r *UserTokenGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A list of tokens registered with the user.
type UserTokenListResponse struct {
	Tokens []UserToken `json:"tokens" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Tokens      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserTokenListResponse) RawJSON() string { return r.JSON.raw }
func (r *UserTokenListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenGetParams struct {
	UserID string `path:"user_id" api:"required" json:"-"`
	paramObj
}

type UserTokenUpdateParams struct {
	UserID string                       `path:"user_id" api:"required" json:"-"`
	Patch  []UserTokenUpdateParamsPatch `json:"patch,omitzero" api:"required"`
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
	Op string `json:"op" api:"required"`
	// The JSON path specifying the part of the profile to operate on.
	Path string `json:"path" api:"required"`
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
	UserID string `path:"user_id" api:"required" json:"-"`
	paramObj
}

type UserTokenAddSingleParams struct {
	UserID string `path:"user_id" api:"required" json:"-"`
	// Any of "firebase-fcm", "apn", "expo", "onesignal".
	ProviderKey UserTokenAddSingleParamsProviderKey `json:"provider_key,omitzero" api:"required"`
	// Information about the device the token came from.
	Device UserTokenAddSingleParamsDevice `json:"device,omitzero"`
	// ISO 8601 formatted date the token expires. Defaults to 2 months. Set to false to
	// disable expiration.
	ExpiryDate UserTokenAddSingleParamsExpiryDateUnion `json:"expiry_date,omitzero"`
	// Tracking information about the device the token came from.
	Tracking UserTokenAddSingleParamsTracking `json:"tracking,omitzero"`
	// Properties about the token.
	Properties any `json:"properties,omitzero"`
	paramObj
}

func (r UserTokenAddSingleParams) MarshalJSON() (data []byte, err error) {
	type shadow UserTokenAddSingleParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserTokenAddSingleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenAddSingleParamsProviderKey string

const (
	UserTokenAddSingleParamsProviderKeyFirebaseFcm UserTokenAddSingleParamsProviderKey = "firebase-fcm"
	UserTokenAddSingleParamsProviderKeyApn         UserTokenAddSingleParamsProviderKey = "apn"
	UserTokenAddSingleParamsProviderKeyExpo        UserTokenAddSingleParamsProviderKey = "expo"
	UserTokenAddSingleParamsProviderKeyOnesignal   UserTokenAddSingleParamsProviderKey = "onesignal"
)

// Information about the device the token came from.
type UserTokenAddSingleParamsDevice struct {
	// Id of the advertising identifier
	AdID param.Opt[string] `json:"ad_id,omitzero"`
	// Id of the application the token is used for
	AppID param.Opt[string] `json:"app_id,omitzero"`
	// Id of the device the token is associated with
	DeviceID param.Opt[string] `json:"device_id,omitzero"`
	// The device manufacturer
	Manufacturer param.Opt[string] `json:"manufacturer,omitzero"`
	// The device model
	Model param.Opt[string] `json:"model,omitzero"`
	// The device platform i.e. android, ios, web
	Platform param.Opt[string] `json:"platform,omitzero"`
	paramObj
}

func (r UserTokenAddSingleParamsDevice) MarshalJSON() (data []byte, err error) {
	type shadow UserTokenAddSingleParamsDevice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserTokenAddSingleParamsDevice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type UserTokenAddSingleParamsExpiryDateUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]   `json:",omitzero,inline"`
	paramUnion
}

func (u UserTokenAddSingleParamsExpiryDateUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfBool)
}
func (u *UserTokenAddSingleParamsExpiryDateUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *UserTokenAddSingleParamsExpiryDateUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

// Tracking information about the device the token came from.
type UserTokenAddSingleParamsTracking struct {
	// The IP address of the device
	IP param.Opt[string] `json:"ip,omitzero"`
	// The latitude of the device
	Lat param.Opt[string] `json:"lat,omitzero"`
	// The longitude of the device
	Long param.Opt[string] `json:"long,omitzero"`
	// The operating system version
	OsVersion param.Opt[string] `json:"os_version,omitzero"`
	paramObj
}

func (r UserTokenAddSingleParamsTracking) MarshalJSON() (data []byte, err error) {
	type shadow UserTokenAddSingleParamsTracking
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserTokenAddSingleParamsTracking) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
