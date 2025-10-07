// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/courier-go/internal/apijson"
	"github.com/stainless-sdks/courier-go/internal/apiquery"
	"github.com/stainless-sdks/courier-go/internal/requestconfig"
	"github.com/stainless-sdks/courier-go/option"
	"github.com/stainless-sdks/courier-go/packages/param"
	"github.com/stainless-sdks/courier-go/packages/respjson"
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
func (r *AudienceService) Get(ctx context.Context, audienceID string, opts ...option.RequestOption) (res *Audience, err error) {
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

type Audience struct {
	// A unique identifier representing the audience_id
	ID        string `json:"id,required"`
	CreatedAt string `json:"created_at,required"`
	// A description of the audience
	Description string `json:"description,required"`
	// The operator to use for filtering
	Filter FilterUnion `json:"filter,required"`
	// The name of the audience
	Name      string `json:"name,required"`
	UpdatedAt string `json:"updated_at,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		Filter      respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Audience) RawJSON() string { return r.JSON.raw }
func (r *Audience) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FilterUnion contains all possible properties and values from [FilterObject],
// [NestedFilterConfig].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FilterUnion struct {
	Operator string `json:"operator"`
	// This field is from variant [FilterObject].
	Path string `json:"path"`
	// This field is from variant [FilterObject].
	Value string `json:"value"`
	// This field is from variant [NestedFilterConfig].
	Rules []FilterConfigUnion `json:"rules"`
	JSON  struct {
		Operator respjson.Field
		Path     respjson.Field
		Value    respjson.Field
		Rules    respjson.Field
		raw      string
	} `json:"-"`
}

func (u FilterUnion) AsFilterObject() (v FilterObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FilterUnion) AsNestedFilterConfig() (v NestedFilterConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FilterUnion) RawJSON() string { return u.JSON.raw }

func (r *FilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this FilterUnion to a FilterUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// FilterUnionParam.Overrides()
func (r FilterUnion) ToParam() FilterUnionParam {
	return param.Override[FilterUnionParam](json.RawMessage(r.RawJSON()))
}

type FilterObject struct {
	// The operator to use for filtering
	//
	// Any of "ENDS_WITH", "EQ", "EXISTS", "GT", "GTE", "INCLUDES", "IS_AFTER",
	// "IS_BEFORE", "LT", "LTE", "NEQ", "OMIT", "STARTS_WITH", "AND", "OR".
	Operator FilterObjectOperator `json:"operator,required"`
	// The attribe name from profile whose value will be operated against the filter
	// value
	Path string `json:"path,required"`
	// The value to use for filtering
	Value string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator    respjson.Field
		Path        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FilterObject) RawJSON() string { return r.JSON.raw }
func (r *FilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The operator to use for filtering
type FilterObjectOperator string

const (
	FilterObjectOperatorEndsWith   FilterObjectOperator = "ENDS_WITH"
	FilterObjectOperatorEq         FilterObjectOperator = "EQ"
	FilterObjectOperatorExists     FilterObjectOperator = "EXISTS"
	FilterObjectOperatorGt         FilterObjectOperator = "GT"
	FilterObjectOperatorGte        FilterObjectOperator = "GTE"
	FilterObjectOperatorIncludes   FilterObjectOperator = "INCLUDES"
	FilterObjectOperatorIsAfter    FilterObjectOperator = "IS_AFTER"
	FilterObjectOperatorIsBefore   FilterObjectOperator = "IS_BEFORE"
	FilterObjectOperatorLt         FilterObjectOperator = "LT"
	FilterObjectOperatorLte        FilterObjectOperator = "LTE"
	FilterObjectOperatorNeq        FilterObjectOperator = "NEQ"
	FilterObjectOperatorOmit       FilterObjectOperator = "OMIT"
	FilterObjectOperatorStartsWith FilterObjectOperator = "STARTS_WITH"
	FilterObjectOperatorAnd        FilterObjectOperator = "AND"
	FilterObjectOperatorOr         FilterObjectOperator = "OR"
)

// The operator to use for filtering
type FilterOperator string

const (
	FilterOperatorEndsWith   FilterOperator = "ENDS_WITH"
	FilterOperatorEq         FilterOperator = "EQ"
	FilterOperatorExists     FilterOperator = "EXISTS"
	FilterOperatorGt         FilterOperator = "GT"
	FilterOperatorGte        FilterOperator = "GTE"
	FilterOperatorIncludes   FilterOperator = "INCLUDES"
	FilterOperatorIsAfter    FilterOperator = "IS_AFTER"
	FilterOperatorIsBefore   FilterOperator = "IS_BEFORE"
	FilterOperatorLt         FilterOperator = "LT"
	FilterOperatorLte        FilterOperator = "LTE"
	FilterOperatorNeq        FilterOperator = "NEQ"
	FilterOperatorOmit       FilterOperator = "OMIT"
	FilterOperatorStartsWith FilterOperator = "STARTS_WITH"
	FilterOperatorAnd        FilterOperator = "AND"
	FilterOperatorOr         FilterOperator = "OR"
)

func FilterParamOfFilterObject(operator FilterObjectOperator, path string, value string) FilterUnionParam {
	var variant FilterObjectParam
	variant.Operator = operator
	variant.Path = path
	variant.Value = value
	return FilterUnionParam{OfFilterObject: &variant}
}

func FilterParamOfNestedFilterConfig(operator NestedFilterConfigOperator, rules []FilterConfigUnionParam) FilterUnionParam {
	var variant NestedFilterConfigParam
	variant.Operator = operator
	variant.Rules = rules
	return FilterUnionParam{OfNestedFilterConfig: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FilterUnionParam struct {
	OfFilterObject       *FilterObjectParam       `json:",omitzero,inline"`
	OfNestedFilterConfig *NestedFilterConfigParam `json:",omitzero,inline"`
	paramUnion
}

func (u FilterUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFilterObject, u.OfNestedFilterConfig)
}
func (u *FilterUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FilterUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFilterObject) {
		return u.OfFilterObject
	} else if !param.IsOmitted(u.OfNestedFilterConfig) {
		return u.OfNestedFilterConfig
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FilterUnionParam) GetPath() *string {
	if vt := u.OfFilterObject; vt != nil {
		return &vt.Path
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FilterUnionParam) GetValue() *string {
	if vt := u.OfFilterObject; vt != nil {
		return &vt.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FilterUnionParam) GetRules() []FilterConfigUnionParam {
	if vt := u.OfNestedFilterConfig; vt != nil {
		return vt.Rules
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FilterUnionParam) GetOperator() *string {
	if vt := u.OfFilterObject; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfNestedFilterConfig; vt != nil {
		return (*string)(&vt.Operator)
	}
	return nil
}

// The properties Operator, Path, Value are required.
type FilterObjectParam struct {
	// The operator to use for filtering
	//
	// Any of "ENDS_WITH", "EQ", "EXISTS", "GT", "GTE", "INCLUDES", "IS_AFTER",
	// "IS_BEFORE", "LT", "LTE", "NEQ", "OMIT", "STARTS_WITH", "AND", "OR".
	Operator FilterObjectOperator `json:"operator,omitzero,required"`
	// The attribe name from profile whose value will be operated against the filter
	// value
	Path string `json:"path,required"`
	// The value to use for filtering
	Value string `json:"value,required"`
	paramObj
}

func (r FilterObjectParam) MarshalJSON() (data []byte, err error) {
	type shadow FilterObjectParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FilterObjectParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FilterConfigUnion contains all possible properties and values from
// [FilterConfigObject], [NestedFilterConfig].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FilterConfigUnion struct {
	Operator string `json:"operator"`
	// This field is from variant [FilterConfigObject].
	Path string `json:"path"`
	// This field is from variant [FilterConfigObject].
	Value string `json:"value"`
	// This field is from variant [NestedFilterConfig].
	Rules []FilterConfigUnion `json:"rules"`
	JSON  struct {
		Operator respjson.Field
		Path     respjson.Field
		Value    respjson.Field
		Rules    respjson.Field
		raw      string
	} `json:"-"`
}

func (u FilterConfigUnion) AsFilterConfigObject() (v FilterConfigObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FilterConfigUnion) AsNestedFilterConfig() (v NestedFilterConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FilterConfigUnion) RawJSON() string { return u.JSON.raw }

func (r *FilterConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this FilterConfigUnion to a FilterConfigUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// FilterConfigUnionParam.Overrides()
func (r FilterConfigUnion) ToParam() FilterConfigUnionParam {
	return param.Override[FilterConfigUnionParam](json.RawMessage(r.RawJSON()))
}

type FilterConfigObject struct {
	// The operator to use for filtering
	//
	// Any of "ENDS_WITH", "EQ", "EXISTS", "GT", "GTE", "INCLUDES", "IS_AFTER",
	// "IS_BEFORE", "LT", "LTE", "NEQ", "OMIT", "STARTS_WITH", "AND", "OR".
	Operator FilterConfigObjectOperator `json:"operator,required"`
	// The attribe name from profile whose value will be operated against the filter
	// value
	Path string `json:"path,required"`
	// The value to use for filtering
	Value string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator    respjson.Field
		Path        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FilterConfigObject) RawJSON() string { return r.JSON.raw }
func (r *FilterConfigObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The operator to use for filtering
type FilterConfigObjectOperator string

const (
	FilterConfigObjectOperatorEndsWith   FilterConfigObjectOperator = "ENDS_WITH"
	FilterConfigObjectOperatorEq         FilterConfigObjectOperator = "EQ"
	FilterConfigObjectOperatorExists     FilterConfigObjectOperator = "EXISTS"
	FilterConfigObjectOperatorGt         FilterConfigObjectOperator = "GT"
	FilterConfigObjectOperatorGte        FilterConfigObjectOperator = "GTE"
	FilterConfigObjectOperatorIncludes   FilterConfigObjectOperator = "INCLUDES"
	FilterConfigObjectOperatorIsAfter    FilterConfigObjectOperator = "IS_AFTER"
	FilterConfigObjectOperatorIsBefore   FilterConfigObjectOperator = "IS_BEFORE"
	FilterConfigObjectOperatorLt         FilterConfigObjectOperator = "LT"
	FilterConfigObjectOperatorLte        FilterConfigObjectOperator = "LTE"
	FilterConfigObjectOperatorNeq        FilterConfigObjectOperator = "NEQ"
	FilterConfigObjectOperatorOmit       FilterConfigObjectOperator = "OMIT"
	FilterConfigObjectOperatorStartsWith FilterConfigObjectOperator = "STARTS_WITH"
	FilterConfigObjectOperatorAnd        FilterConfigObjectOperator = "AND"
	FilterConfigObjectOperatorOr         FilterConfigObjectOperator = "OR"
)

// The operator to use for filtering
type FilterConfigOperator string

const (
	FilterConfigOperatorEndsWith   FilterConfigOperator = "ENDS_WITH"
	FilterConfigOperatorEq         FilterConfigOperator = "EQ"
	FilterConfigOperatorExists     FilterConfigOperator = "EXISTS"
	FilterConfigOperatorGt         FilterConfigOperator = "GT"
	FilterConfigOperatorGte        FilterConfigOperator = "GTE"
	FilterConfigOperatorIncludes   FilterConfigOperator = "INCLUDES"
	FilterConfigOperatorIsAfter    FilterConfigOperator = "IS_AFTER"
	FilterConfigOperatorIsBefore   FilterConfigOperator = "IS_BEFORE"
	FilterConfigOperatorLt         FilterConfigOperator = "LT"
	FilterConfigOperatorLte        FilterConfigOperator = "LTE"
	FilterConfigOperatorNeq        FilterConfigOperator = "NEQ"
	FilterConfigOperatorOmit       FilterConfigOperator = "OMIT"
	FilterConfigOperatorStartsWith FilterConfigOperator = "STARTS_WITH"
	FilterConfigOperatorAnd        FilterConfigOperator = "AND"
	FilterConfigOperatorOr         FilterConfigOperator = "OR"
)

func FilterConfigParamOfFilterConfigObject(operator FilterConfigObjectOperator, path string, value string) FilterConfigUnionParam {
	var variant FilterConfigObjectParam
	variant.Operator = operator
	variant.Path = path
	variant.Value = value
	return FilterConfigUnionParam{OfFilterConfigObject: &variant}
}

func FilterConfigParamOfNestedFilterConfig(operator NestedFilterConfigOperator, rules []FilterConfigUnionParam) FilterConfigUnionParam {
	var variant NestedFilterConfigParam
	variant.Operator = operator
	variant.Rules = rules
	return FilterConfigUnionParam{OfNestedFilterConfig: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FilterConfigUnionParam struct {
	OfFilterConfigObject *FilterConfigObjectParam `json:",omitzero,inline"`
	OfNestedFilterConfig *NestedFilterConfigParam `json:",omitzero,inline"`
	paramUnion
}

func (u FilterConfigUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFilterConfigObject, u.OfNestedFilterConfig)
}
func (u *FilterConfigUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FilterConfigUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFilterConfigObject) {
		return u.OfFilterConfigObject
	} else if !param.IsOmitted(u.OfNestedFilterConfig) {
		return u.OfNestedFilterConfig
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FilterConfigUnionParam) GetPath() *string {
	if vt := u.OfFilterConfigObject; vt != nil {
		return &vt.Path
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FilterConfigUnionParam) GetValue() *string {
	if vt := u.OfFilterConfigObject; vt != nil {
		return &vt.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FilterConfigUnionParam) GetRules() []FilterConfigUnionParam {
	if vt := u.OfNestedFilterConfig; vt != nil {
		return vt.Rules
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FilterConfigUnionParam) GetOperator() *string {
	if vt := u.OfFilterConfigObject; vt != nil {
		return (*string)(&vt.Operator)
	} else if vt := u.OfNestedFilterConfig; vt != nil {
		return (*string)(&vt.Operator)
	}
	return nil
}

// The properties Operator, Path, Value are required.
type FilterConfigObjectParam struct {
	// The operator to use for filtering
	//
	// Any of "ENDS_WITH", "EQ", "EXISTS", "GT", "GTE", "INCLUDES", "IS_AFTER",
	// "IS_BEFORE", "LT", "LTE", "NEQ", "OMIT", "STARTS_WITH", "AND", "OR".
	Operator FilterConfigObjectOperator `json:"operator,omitzero,required"`
	// The attribe name from profile whose value will be operated against the filter
	// value
	Path string `json:"path,required"`
	// The value to use for filtering
	Value string `json:"value,required"`
	paramObj
}

func (r FilterConfigObjectParam) MarshalJSON() (data []byte, err error) {
	type shadow FilterConfigObjectParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FilterConfigObjectParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NestedFilterConfig struct {
	// The operator to use for filtering
	//
	// Any of "ENDS_WITH", "EQ", "EXISTS", "GT", "GTE", "INCLUDES", "IS_AFTER",
	// "IS_BEFORE", "LT", "LTE", "NEQ", "OMIT", "STARTS_WITH", "AND", "OR".
	Operator NestedFilterConfigOperator `json:"operator,required"`
	Rules    []FilterConfigUnion        `json:"rules,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator    respjson.Field
		Rules       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NestedFilterConfig) RawJSON() string { return r.JSON.raw }
func (r *NestedFilterConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this NestedFilterConfig to a NestedFilterConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// NestedFilterConfigParam.Overrides()
func (r NestedFilterConfig) ToParam() NestedFilterConfigParam {
	return param.Override[NestedFilterConfigParam](json.RawMessage(r.RawJSON()))
}

// The operator to use for filtering
type NestedFilterConfigOperator string

const (
	NestedFilterConfigOperatorEndsWith   NestedFilterConfigOperator = "ENDS_WITH"
	NestedFilterConfigOperatorEq         NestedFilterConfigOperator = "EQ"
	NestedFilterConfigOperatorExists     NestedFilterConfigOperator = "EXISTS"
	NestedFilterConfigOperatorGt         NestedFilterConfigOperator = "GT"
	NestedFilterConfigOperatorGte        NestedFilterConfigOperator = "GTE"
	NestedFilterConfigOperatorIncludes   NestedFilterConfigOperator = "INCLUDES"
	NestedFilterConfigOperatorIsAfter    NestedFilterConfigOperator = "IS_AFTER"
	NestedFilterConfigOperatorIsBefore   NestedFilterConfigOperator = "IS_BEFORE"
	NestedFilterConfigOperatorLt         NestedFilterConfigOperator = "LT"
	NestedFilterConfigOperatorLte        NestedFilterConfigOperator = "LTE"
	NestedFilterConfigOperatorNeq        NestedFilterConfigOperator = "NEQ"
	NestedFilterConfigOperatorOmit       NestedFilterConfigOperator = "OMIT"
	NestedFilterConfigOperatorStartsWith NestedFilterConfigOperator = "STARTS_WITH"
	NestedFilterConfigOperatorAnd        NestedFilterConfigOperator = "AND"
	NestedFilterConfigOperatorOr         NestedFilterConfigOperator = "OR"
)

// The properties Operator, Rules are required.
type NestedFilterConfigParam struct {
	// The operator to use for filtering
	//
	// Any of "ENDS_WITH", "EQ", "EXISTS", "GT", "GTE", "INCLUDES", "IS_AFTER",
	// "IS_BEFORE", "LT", "LTE", "NEQ", "OMIT", "STARTS_WITH", "AND", "OR".
	Operator NestedFilterConfigOperator `json:"operator,omitzero,required"`
	Rules    []FilterConfigUnionParam   `json:"rules,omitzero,required"`
	paramObj
}

func (r NestedFilterConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow NestedFilterConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NestedFilterConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Paging struct {
	More   bool   `json:"more,required"`
	Cursor string `json:"cursor,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		More        respjson.Field
		Cursor      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Paging) RawJSON() string { return r.JSON.raw }
func (r *Paging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudienceUpdateResponse struct {
	Audience Audience `json:"audience,required"`
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
	Items  []Audience `json:"items,required"`
	Paging Paging     `json:"paging,required"`
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
	Paging Paging                            `json:"paging,required"`
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
	// The operator to use for filtering
	Filter FilterUnionParam `json:"filter,omitzero"`
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
