// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
	"net/http"
	"slices"

	"github.com/trycourier/courier-go/v4/internal/apijson"
	"github.com/trycourier/courier-go/v4/internal/requestconfig"
	"github.com/trycourier/courier-go/v4/option"
	"github.com/trycourier/courier-go/v4/packages/param"
	"github.com/trycourier/courier-go/v4/packages/respjson"
)

// Report suspected bugs, confusing behavior and missing features in Courier, from
// an AI agent, a developer tool or by hand.
//
// FeedbackService contains methods and other services that help with interacting
// with the Courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFeedbackService] method instead.
type FeedbackService struct {
	Options []option.RequestOption
}

// NewFeedbackService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFeedbackService(opts ...option.RequestOption) (r FeedbackService) {
	r = FeedbackService{}
	r.Options = opts
	return
}

// Reports a suspected bug, confusing behavior or a missing feature in Courier. It
// is built for AI agents and developer tools to call when Courier does something
// unexpected, and people can call it too.
//
// A report is not a support ticket. A `201` means Courier stored the report; it
// does not promise a reply or a fix. For help with your account, contact Courier
// support.
//
// Never include secrets, API keys or customer data such as recipient names, email
// addresses or message content.
//
// Requires an API key with no scope restrictions; a JWT or a scoped API key is
// refused. Each workspace can submit 200 reports per UTC day, and requests that
// fail validation count toward that limit.
func (r *FeedbackService) New(ctx context.Context, body FeedbackNewParams, opts ...option.RequestOption) (res *FeedbackNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "feedback"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type FeedbackNewResponse struct {
	// The ID of the stored report, `fbk_` followed by a ULID.
	FeedbackID string `json:"feedback_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FeedbackID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FeedbackNewResponse) RawJSON() string { return r.JSON.raw }
func (r *FeedbackNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FeedbackNewParams struct {
	// What kind of report this is: `possible_bug` for behavior that looks wrong,
	// `confusing_behavior` for behavior that works but was hard to understand,
	// `feature_request` for something Courier does not do yet, and `other`.
	//
	// Any of "possible_bug", "confusing_behavior", "feature_request", "other".
	Category FeedbackNewParamsCategory `json:"category,omitzero" api:"required"`
	// What happened, in enough detail for someone at Courier to follow it.
	Description string `json:"description" api:"required"`
	// Where you met the problem: the REST `api`, an `sdk`, the Courier `mcp` server,
	// the Courier app `ui`, the `docs`, or `other`.
	//
	// Any of "api", "sdk", "mcp", "ui", "docs", "other".
	Surface FeedbackNewParamsSurface `json:"surface,omitzero" api:"required"`
	// A one-line summary of the report.
	Title string `json:"title" api:"required"`
	// What you expected to happen instead.
	ExpectedBehavior param.Opt[string] `json:"expected_behavior,omitzero"`
	// The agent or tool that submitted the report, such as `Claude Code`.
	Harness param.Opt[string] `json:"harness,omitzero"`
	// Where it happened, such as an endpoint (`POST /send`), an SDK method, an MCP
	// tool name or a docs URL.
	Location param.Opt[string] `json:"location,omitzero"`
	// The AI model that submitted the report, such as `claude-opus-5-5`.
	Model param.Opt[string] `json:"model,omitzero"`
	// The request ID of an API call that shows the problem, if there is one.
	RequestID param.Opt[string] `json:"request_id,omitzero"`
	paramObj
}

func (r FeedbackNewParams) MarshalJSON() (data []byte, err error) {
	type shadow FeedbackNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FeedbackNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// What kind of report this is: `possible_bug` for behavior that looks wrong,
// `confusing_behavior` for behavior that works but was hard to understand,
// `feature_request` for something Courier does not do yet, and `other`.
type FeedbackNewParamsCategory string

const (
	FeedbackNewParamsCategoryPossibleBug       FeedbackNewParamsCategory = "possible_bug"
	FeedbackNewParamsCategoryConfusingBehavior FeedbackNewParamsCategory = "confusing_behavior"
	FeedbackNewParamsCategoryFeatureRequest    FeedbackNewParamsCategory = "feature_request"
	FeedbackNewParamsCategoryOther             FeedbackNewParamsCategory = "other"
)

// Where you met the problem: the REST `api`, an `sdk`, the Courier `mcp` server,
// the Courier app `ui`, the `docs`, or `other`.
type FeedbackNewParamsSurface string

const (
	FeedbackNewParamsSurfaceAPI   FeedbackNewParamsSurface = "api"
	FeedbackNewParamsSurfaceSDK   FeedbackNewParamsSurface = "sdk"
	FeedbackNewParamsSurfaceMcp   FeedbackNewParamsSurface = "mcp"
	FeedbackNewParamsSurfaceUi    FeedbackNewParamsSurface = "ui"
	FeedbackNewParamsSurfaceDocs  FeedbackNewParamsSurface = "docs"
	FeedbackNewParamsSurfaceOther FeedbackNewParamsSurface = "other"
)
