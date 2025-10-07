// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
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

// AuditEventService contains methods and other services that help with interacting
// with the courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuditEventService] method instead.
type AuditEventService struct {
	Options []option.RequestOption
}

// NewAuditEventService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAuditEventService(opts ...option.RequestOption) (r AuditEventService) {
	r = AuditEventService{}
	r.Options = opts
	return
}

// Fetch a specific audit event by ID.
func (r *AuditEventService) Get(ctx context.Context, auditEventID string, opts ...option.RequestOption) (res *AuditEvent, err error) {
	opts = slices.Concat(r.Options, opts)
	if auditEventID == "" {
		err = errors.New("missing required audit-event-id parameter")
		return
	}
	path := fmt.Sprintf("audit-events/%s", auditEventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetch the list of audit events
func (r *AuditEventService) List(ctx context.Context, query AuditEventListParams, opts ...option.RequestOption) (res *AuditEventListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "audit-events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AuditEvent struct {
	Actor        AuditEventActor `json:"actor,required"`
	AuditEventID string          `json:"auditEventId,required"`
	Source       string          `json:"source,required"`
	Target       string          `json:"target,required"`
	Timestamp    string          `json:"timestamp,required"`
	Type         string          `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Actor        respjson.Field
		AuditEventID respjson.Field
		Source       respjson.Field
		Target       respjson.Field
		Timestamp    respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuditEvent) RawJSON() string { return r.JSON.raw }
func (r *AuditEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuditEventActor struct {
	ID    string `json:"id,required"`
	Email string `json:"email,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Email       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuditEventActor) RawJSON() string { return r.JSON.raw }
func (r *AuditEventActor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuditEventListResponse struct {
	Paging  Paging       `json:"paging,required"`
	Results []AuditEvent `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Paging      respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuditEventListResponse) RawJSON() string { return r.JSON.raw }
func (r *AuditEventListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuditEventListParams struct {
	// A unique identifier that allows for fetching the next set of audit events.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AuditEventListParams]'s query parameters as `url.Values`.
func (r AuditEventListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
