// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/trycourier/courier-go/v4/internal/apijson"
	"github.com/trycourier/courier-go/v4/internal/apiquery"
	"github.com/trycourier/courier-go/v4/internal/requestconfig"
	"github.com/trycourier/courier-go/v4/option"
	"github.com/trycourier/courier-go/v4/packages/param"
	"github.com/trycourier/courier-go/v4/packages/respjson"
)

// Invoke a stored automation template or an ad hoc automation defined in the
// request.
//
// AutomationService contains methods and other services that help with interacting
// with the Courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAutomationService] method instead.
type AutomationService struct {
	Options []option.RequestOption
	// Invoke a stored automation template or an ad hoc automation defined in the
	// request.
	Invoke AutomationInvokeService
	// Invoke a stored automation template or an ad hoc automation defined in the
	// request.
	Runs AutomationRunService
}

// NewAutomationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAutomationService(opts ...option.RequestOption) (r AutomationService) {
	r = AutomationService{}
	r.Options = opts
	r.Invoke = NewAutomationInvokeService(opts...)
	r.Runs = NewAutomationRunService(opts...)
	return
}

// Lists the workspace's saved automation templates, each with its id and a cursor
// for paging to the next page of results.
func (r *AutomationService) List(ctx context.Context, query AutomationListParams, opts ...option.RequestOption) (res *AutomationTemplateListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "automations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type AutomationInvokeResponse struct {
	RunID string `json:"runId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RunID       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationInvokeResponse) RawJSON() string { return r.JSON.raw }
func (r *AutomationInvokeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An Automation run as it appears in a list response.
type AutomationRunListItem struct {
	// A unique identifier representing the run.
	RunID string `json:"run_id" api:"required"`
	// Internal provenance strings describing what started the run, e.g.
	// `invoke/<template_id>` or `segment/page/Pricing Page`. Diagnostic only — the
	// format is unstable and should not be parsed.
	Source []string `json:"source" api:"required"`
	// When the run started, as an ISO 8601 timestamp.
	CreatedAt string `json:"created_at"`
	// The state of the run: `PROCESSING`, `PROCESSED`, `WAITING`, `CANCELED`, `ERROR`,
	// `THROTTLED`, or `NOT PROCESSED`. Not an enum — new values have been added
	// before.
	Status string `json:"status"`
	// The id of the Automation Template this run belongs to.
	TemplateID string `json:"template_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RunID       respjson.Field
		Source      respjson.Field
		CreatedAt   respjson.Field
		Status      respjson.Field
		TemplateID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationRunListItem) RawJSON() string { return r.JSON.raw }
func (r *AutomationRunListItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A page of Automation runs.
type AutomationRunListResponse struct {
	Runs []AutomationRunListItem `json:"runs" api:"required"`
	// Pass back as `cursor` to fetch the next page. Absent on the last page.
	NextCursor string `json:"next_cursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Runs        respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationRunListResponse) RawJSON() string { return r.JSON.raw }
func (r *AutomationRunListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One executed step of an Automation run.
type AutomationRunStep struct {
	// The kind of step that ran, e.g. `send`, `delay`, or `update-profile`.
	Action string `json:"action" api:"required"`
	// The state of the step: the seven run statuses, plus `SKIPPED` and `COMPUTING`.
	// Not an enum — new values have been added before.
	Status string `json:"status" api:"required"`
	// When the step started, as an ISO 8601 timestamp.
	CreatedAt string `json:"created_at"`
	// The message this step produced, present on send steps. Pass it to
	// `GET /messages/{message_id}` for delivery status. A send to a List or an
	// Audience yields one id for the request, not one per recipient.
	MessageID string `json:"message_id"`
	// A unique identifier representing the step.
	StepID string `json:"step_id"`
	// When the step last changed state, as an ISO 8601 timestamp.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Status      respjson.Field
		CreatedAt   respjson.Field
		MessageID   respjson.Field
		StepID      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationRunStep) RawJSON() string { return r.JSON.raw }
func (r *AutomationRunStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Every step of an Automation run. Not paginated.
type AutomationRunStepsResponse struct {
	Steps []AutomationRunStep `json:"steps" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Steps       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationRunStepsResponse) RawJSON() string { return r.JSON.raw }
func (r *AutomationRunStepsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationTemplate struct {
	// The unique identifier of the automation template.
	ID string `json:"id" api:"required"`
	// The name of the automation template.
	Name string `json:"name" api:"required"`
	// The version of the template published or drafted.
	//
	// Any of "published", "draft".
	Version AutomationTemplateVersion `json:"version" api:"required"`
	// ISO 8601 timestamp when the template was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// ISO 8601 timestamp when the template was last updated.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Version     respjson.Field
		CreatedAt   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationTemplate) RawJSON() string { return r.JSON.raw }
func (r *AutomationTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The version of the template published or drafted.
type AutomationTemplateVersion string

const (
	AutomationTemplateVersionPublished AutomationTemplateVersion = "published"
	AutomationTemplateVersionDraft     AutomationTemplateVersion = "draft"
)

type AutomationTemplateListResponse struct {
	// A cursor token for pagination. Present when there are more results available.
	Cursor    string               `json:"cursor"`
	Templates []AutomationTemplate `json:"templates"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cursor      respjson.Field
		Templates   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomationTemplateListResponse) RawJSON() string { return r.JSON.raw }
func (r *AutomationTemplateListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationListParams struct {
	// A cursor token for pagination. Use the cursor from the previous response to
	// fetch the next page of results.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// The version of templates to retrieve. Accepted values are published (for
	// published templates) or draft (for draft templates). Defaults to published.
	//
	// Any of "published", "draft".
	Version AutomationListParamsVersion `query:"version,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AutomationListParams]'s query parameters as `url.Values`.
func (r AutomationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The version of templates to retrieve. Accepted values are published (for
// published templates) or draft (for draft templates). Defaults to published.
type AutomationListParamsVersion string

const (
	AutomationListParamsVersionPublished AutomationListParamsVersion = "published"
	AutomationListParamsVersionDraft     AutomationListParamsVersion = "draft"
)
