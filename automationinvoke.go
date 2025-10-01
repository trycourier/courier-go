// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"github.com/stainless-sdks/courier-go/internal/apijson"
	"github.com/stainless-sdks/courier-go/option"
	"github.com/stainless-sdks/courier-go/packages/param"
	"github.com/stainless-sdks/courier-go/packages/respjson"
)

// AutomationInvokeService contains methods and other services that help with
// interacting with the courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAutomationInvokeService] method instead.
type AutomationInvokeService struct {
	Options []option.RequestOption
}

// NewAutomationInvokeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAutomationInvokeService(opts ...option.RequestOption) (r AutomationInvokeService) {
	r = AutomationInvokeService{}
	r.Options = opts
	return
}

type AutomationInvokeParams struct {
	Brand     param.Opt[string] `json:"brand,omitzero"`
	Recipient param.Opt[string] `json:"recipient,omitzero"`
	Template  param.Opt[string] `json:"template,omitzero"`
	Data      map[string]any    `json:"data,omitzero"`
	Profile   any               `json:"profile,omitzero"`
	paramObj
}

func (r AutomationInvokeParams) MarshalJSON() (data []byte, err error) {
	type shadow AutomationInvokeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationInvokeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

type AutomationStepParam struct {
	If  param.Opt[string] `json:"if,omitzero"`
	Ref param.Opt[string] `json:"ref,omitzero"`
	paramObj
}

func (r AutomationStepParam) MarshalJSON() (data []byte, err error) {
	type shadow AutomationStepParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationStepParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergeAlgorithm string

const (
	MergeAlgorithmReplace   MergeAlgorithm = "replace"
	MergeAlgorithmNone      MergeAlgorithm = "none"
	MergeAlgorithmOverwrite MergeAlgorithm = "overwrite"
	MergeAlgorithmSoftMerge MergeAlgorithm = "soft-merge"
)
