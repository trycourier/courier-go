// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/trycourier/courier-go/v3/internal/apijson"
	"github.com/trycourier/courier-go/v3/internal/apiquery"
	"github.com/trycourier/courier-go/v3/internal/requestconfig"
	"github.com/trycourier/courier-go/v3/option"
	"github.com/trycourier/courier-go/v3/packages/param"
	"github.com/trycourier/courier-go/v3/packages/respjson"
	"github.com/trycourier/courier-go/v3/shared"
)

// AudienceService contains methods and other services that help with interacting
// with the courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAudienceService] method instead.
type AudienceService struct {
	Options []option.RequestOption
}

// NewAudienceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAudienceService(opts ...option.RequestOption) (r AudienceService) {
	r = AudienceService{}
	r.Options = opts
	return
}

// Returns the specified audience by id.
func (r *AudienceService) Get(ctx context.Context, audienceID string, opts ...option.RequestOption) (res *shared.Audience, err error) {
	opts = slices.Concat(r.Options, opts)
	if audienceID == "" {
		err = errors.New("missing required audience_id parameter")
		return
	}
	path := fmt.Sprintf("audiences/%s", audienceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Creates or updates audience.
func (r *AudienceService) Update(ctx context.Context, audienceID string, body AudienceUpdateParams, opts ...option.RequestOption) (res *AudienceUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if audienceID == "" {
		err = errors.New("missing required audience_id parameter")
		return
	}
	path := fmt.Sprintf("audiences/%s", audienceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Get the audiences associated with the authorization token.
func (r *AudienceService) List(ctx context.Context, query AudienceListParams, opts ...option.RequestOption) (res *AudienceListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "audiences"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes the specified audience.
func (r *AudienceService) Delete(ctx context.Context, audienceID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if audienceID == "" {
		err = errors.New("missing required audience_id parameter")
		return
	}
	path := fmt.Sprintf("audiences/%s", audienceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get list of members of an audience.
func (r *AudienceService) ListMembers(ctx context.Context, audienceID string, query AudienceListMembersParams, opts ...option.RequestOption) (res *AudienceListMembersResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if audienceID == "" {
		err = errors.New("missing required audience_id parameter")
		return
	}
	path := fmt.Sprintf("audiences/%s/members", audienceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AudienceUpdateResponse struct {
	Audience shared.Audience `json:"audience,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Audience    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudienceUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *AudienceUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudienceListResponse struct {
	Items  []shared.Audience `json:"items,required"`
	Paging shared.Paging     `json:"paging,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudienceListResponse) RawJSON() string { return r.JSON.raw }
func (r *AudienceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudienceListMembersResponse struct {
	Items  []AudienceListMembersResponseItem `json:"items,required"`
	Paging shared.Paging                     `json:"paging,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudienceListMembersResponse) RawJSON() string { return r.JSON.raw }
func (r *AudienceListMembersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudienceListMembersResponseItem struct {
	AddedAt         string `json:"added_at,required"`
	AudienceID      string `json:"audience_id,required"`
	AudienceVersion int64  `json:"audience_version,required"`
	MemberID        string `json:"member_id,required"`
	Reason          string `json:"reason,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddedAt         respjson.Field
		AudienceID      respjson.Field
		AudienceVersion respjson.Field
		MemberID        respjson.Field
		Reason          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudienceListMembersResponseItem) RawJSON() string { return r.JSON.raw }
func (r *AudienceListMembersResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudienceUpdateParams struct {
	// A description of the audience
	Description param.Opt[string] `json:"description,omitzero"`
	// The name of the audience
	Name param.Opt[string] `json:"name,omitzero"`
	// A single filter to use for filtering
	Filter shared.FilterParam `json:"filter,omitzero"`
	paramObj
}

func (r AudienceUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AudienceUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AudienceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudienceListParams struct {
	// A unique identifier that allows for fetching the next set of audiences
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AudienceListParams]'s query parameters as `url.Values`.
func (r AudienceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AudienceListMembersParams struct {
	// A unique identifier that allows for fetching the next set of members
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AudienceListMembersParams]'s query parameters as
// `url.Values`.
func (r AudienceListMembersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
