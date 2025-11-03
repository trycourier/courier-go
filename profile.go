// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/trycourier/courier-go/v3/internal/apijson"
	"github.com/trycourier/courier-go/v3/internal/requestconfig"
	"github.com/trycourier/courier-go/v3/option"
	"github.com/trycourier/courier-go/v3/packages/param"
	"github.com/trycourier/courier-go/v3/packages/respjson"
	"github.com/trycourier/courier-go/v3/shared"
)

// ProfileService contains methods and other services that help with interacting
// with the Courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProfileService] method instead.
type ProfileService struct {
	Options []option.RequestOption
	Lists   ProfileListService
}

// NewProfileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProfileService(opts ...option.RequestOption) (r ProfileService) {
	r = ProfileService{}
	r.Options = opts
	r.Lists = NewProfileListService(opts...)
	return
}

// Merge the supplied values with an existing profile or create a new profile if
// one doesn't already exist.
func (r *ProfileService) New(ctx context.Context, userID string, body ProfileNewParams, opts ...option.RequestOption) (res *ProfileNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("profiles/%s", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns the specified user profile.
func (r *ProfileService) Get(ctx context.Context, userID string, opts ...option.RequestOption) (res *ProfileGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("profiles/%s", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a profile
func (r *ProfileService) Update(ctx context.Context, userID string, body ProfileUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("profiles/%s", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return
}

// Deletes the specified user profile.
func (r *ProfileService) Delete(ctx context.Context, userID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("profiles/%s", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// When using `PUT`, be sure to include all the key-value pairs required by the
// recipient's profile. Any key-value pairs that exist in the profile but fail to
// be included in the `PUT` request will be removed from the profile. Remember, a
// `PUT` update is a full replacement of the data. For partial updates, use the
// [Patch](https://www.courier.com/docs/reference/profiles/patch/) request.
func (r *ProfileService) Replace(ctx context.Context, userID string, body ProfileReplaceParams, opts ...option.RequestOption) (res *ProfileReplaceResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("profiles/%s", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// The property ListID is required.
type SubscribeToListsRequestItemParam struct {
	ListID      string                           `json:"listId,required"`
	Preferences shared.RecipientPreferencesParam `json:"preferences,omitzero"`
	paramObj
}

func (r SubscribeToListsRequestItemParam) MarshalJSON() (data []byte, err error) {
	type shadow SubscribeToListsRequestItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubscribeToListsRequestItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileNewResponse struct {
	// Any of "SUCCESS".
	Status ProfileNewResponseStatus `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ProfileNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileNewResponseStatus string

const (
	ProfileNewResponseStatusSuccess ProfileNewResponseStatus = "SUCCESS"
)

type ProfileGetResponse struct {
	Profile     map[string]any              `json:"profile,required"`
	Preferences shared.RecipientPreferences `json:"preferences,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Profile     respjson.Field
		Preferences respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ProfileGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileReplaceResponse struct {
	// Any of "SUCCESS".
	Status ProfileReplaceResponseStatus `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileReplaceResponse) RawJSON() string { return r.JSON.raw }
func (r *ProfileReplaceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileReplaceResponseStatus string

const (
	ProfileReplaceResponseStatusSuccess ProfileReplaceResponseStatus = "SUCCESS"
)

type ProfileNewParams struct {
	Profile map[string]any `json:"profile,omitzero,required"`
	paramObj
}

func (r ProfileNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ProfileNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileUpdateParams struct {
	// List of patch operations to apply to the profile.
	Patch []ProfileUpdateParamsPatch `json:"patch,omitzero,required"`
	paramObj
}

func (r ProfileUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ProfileUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Op, Path, Value are required.
type ProfileUpdateParamsPatch struct {
	// The operation to perform.
	Op string `json:"op,required"`
	// The JSON path specifying the part of the profile to operate on.
	Path string `json:"path,required"`
	// The value for the operation.
	Value string `json:"value,required"`
	paramObj
}

func (r ProfileUpdateParamsPatch) MarshalJSON() (data []byte, err error) {
	type shadow ProfileUpdateParamsPatch
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileUpdateParamsPatch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileReplaceParams struct {
	Profile map[string]any `json:"profile,omitzero,required"`
	paramObj
}

func (r ProfileReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow ProfileReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
