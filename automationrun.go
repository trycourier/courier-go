// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/trycourier/courier-go/v4/internal/apiquery"
	"github.com/trycourier/courier-go/v4/internal/requestconfig"
	"github.com/trycourier/courier-go/v4/option"
	"github.com/trycourier/courier-go/v4/packages/param"
)

// Invoke a stored automation template or an ad hoc automation defined in the
// request.
//
// AutomationRunService contains methods and other services that help with
// interacting with the Courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAutomationRunService] method instead.
type AutomationRunService struct {
	Options []option.RequestOption
}

// NewAutomationRunService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAutomationRunService(opts ...option.RequestOption) (r AutomationRunService) {
	r = AutomationRunService{}
	r.Options = opts
	return
}

// List runs of the workspace's v2 Automations, newest first, filtered by status,
// Template, or date range and paged by cursor. Journey (v3) runs are listed by
// `GET /journeys/runs` instead — the two surfaces never return each other's runs.
// Runs are retained for 95 days.
func (r *AutomationRunService) List(ctx context.Context, query AutomationRunListParams, opts ...option.RequestOption) (res *AutomationRunListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "automations/runs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List the per-step state of one Automation run, in full — this endpoint is not
// paginated. `message_id` is present on send steps that produced a message; follow
// it to `GET /messages/{message_id}` for delivery status. A send to a List or an
// Audience yields one `message_id` for the request, not one per recipient.
func (r *AutomationRunService) ListSteps(ctx context.Context, id string, opts ...option.RequestOption) (res *AutomationRunStepsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("automations/runs/%s/steps", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type AutomationRunListParams struct {
	// A cursor token for pagination. Use the `next_cursor` from the previous response
	// to fetch the next page of results. Treat it as opaque.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// An inclusive upper bound on `created_at`, in the same format as `start_date`.
	EndDate param.Opt[string] `query:"end_date,omitzero" json:"-"`
	// The number of runs to return per page, between `1` and `50`. Defaults to `20`.
	// Values outside the range are clamped, and a non-numeric value falls back to
	// `20`.
	Limit param.Opt[string] `query:"limit,omitzero" json:"-"`
	// An inclusive lower bound on `created_at`, as an ISO 8601 date or timestamp (e.g.
	// `2026-08-18` or `2026-08-18T20:06:36.259Z`). Any other format returns `400`.
	StartDate param.Opt[string] `query:"start_date,omitzero" json:"-"`
	// A comma-separated list of run statuses to filter on, e.g. `PROCESSED,ERROR`.
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	// A comma-separated list of Automation Template ids to filter on.
	TemplateID param.Opt[string] `query:"template_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AutomationRunListParams]'s query parameters as
// `url.Values`.
func (r AutomationRunListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
