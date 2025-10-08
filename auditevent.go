// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/trycourier/courier-go/internal/apijson"
	"github.com/trycourier/courier-go/internal/apiquery"
	"github.com/trycourier/courier-go/internal/requestconfig"
	"github.com/trycourier/courier-go/option"
	"github.com/trycourier/courier-go/packages/param"
	"github.com/trycourier/courier-go/packages/respjson"
	"github.com/trycourier/courier-go/shared"
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
func (r *AuditEventService) Get(ctx context.Context, auditEventID string, opts ...option.RequestOption) (res *shared.AuditEvent, err error) {
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

type AuditEventListResponse struct {
	Paging  shared.Paging       `json:"paging,required"`
	Results []shared.AuditEvent `json:"results,required"`
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
