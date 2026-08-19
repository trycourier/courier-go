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

// Build, version, publish, invoke, and cancel multi-step notification workflows,
// along with the templates scoped to them.
//
// JourneyRunService contains methods and other services that help with interacting
// with the Courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewJourneyRunService] method instead.
type JourneyRunService struct {
	Options []option.RequestOption
}

// NewJourneyRunService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewJourneyRunService(opts ...option.RequestOption) (r JourneyRunService) {
	r = JourneyRunService{}
	r.Options = opts
	return
}

// Fetch one Journey run by id. Returns `404` for an unknown run, a run belonging
// to another workspace, a run past the 95-day retention window, or an Automation
// run id — the same body in every case, so the response never reveals whether a
// run exists elsewhere.
func (r *JourneyRunService) Get(ctx context.Context, runID string, opts ...option.RequestOption) (res *JourneyRunResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("journeys/runs/%s", runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List runs of the workspace's Journeys, newest first, filtered by status,
// Journey, or date range and paged by cursor. Runs of v2 Automations are listed by
// `GET /automations/runs` instead — the two surfaces never return each other's
// runs. Runs are retained for 95 days.
func (r *JourneyRunService) List(ctx context.Context, query JourneyRunListParams, opts ...option.RequestOption) (res *JourneyRunListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "journeys/runs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List the per-node state of one Journey run, in full — this endpoint is not
// paginated. Each step's `node_id` is the id of the node in the published Journey,
// so a step maps directly onto the Journey graph. `message_id` is present on send
// steps that produced a message; follow it to `GET /messages/{message_id}` for
// delivery status.
func (r *JourneyRunService) ListSteps(ctx context.Context, runID string, opts ...option.RequestOption) (res *JourneyRunStepsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("journeys/runs/%s/steps", runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type JourneyRunListParams struct {
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
	// A comma-separated list of Journey ids to filter on.
	TemplateID param.Opt[string] `query:"template_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [JourneyRunListParams]'s query parameters as `url.Values`.
func (r JourneyRunListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
