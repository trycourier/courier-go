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

// AutomationService contains methods and other services that help with interacting
// with the Courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAutomationService] method instead.
type AutomationService struct {
	Options []option.RequestOption
	Invoke  AutomationInvokeService
}

// NewAutomationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAutomationService(opts ...option.RequestOption) (r AutomationService) {
	r = AutomationService{}
	r.Options = opts
	r.Invoke = NewAutomationInvokeService(opts...)
	return
}

// Get the list of automations.
func (r *AutomationService) List(ctx context.Context, query AutomationListParams, opts ...option.RequestOption) (res *AutomationTemplateListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "automations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AutomationInvokeResponse struct {
	RunID string `json:"runId,required"`
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

type AutomationTemplate struct {
	// The unique identifier of the automation template.
	ID string `json:"id,required"`
	// The name of the automation template.
	Name string `json:"name,required"`
	// The version of the template published, draft.
	//
	// Any of "published", "draft".
	Version AutomationTemplateVersion `json:"version,required"`
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

// The version of the template published, draft.
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
