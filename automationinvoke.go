// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/courier-go/internal/apijson"
	"github.com/stainless-sdks/courier-go/internal/requestconfig"
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

// Invoke an ad hoc automation run. This endpoint accepts a JSON payload with a
// series of automation steps. For information about what steps are available,
// checkout the ad hoc automation guide
// [here](https://www.courier.com/docs/automations/steps/).
func (r *AutomationInvokeService) InvokeAdHoc(ctx context.Context, body AutomationInvokeInvokeAdHocParams, opts ...option.RequestOption) (res *AutomationInvokeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "automations/invoke"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Invoke an automation run from an automation template.
func (r *AutomationInvokeService) InvokeByTemplate(ctx context.Context, templateID string, body AutomationInvokeInvokeByTemplateParams, opts ...option.RequestOption) (res *AutomationInvokeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if templateID == "" {
		err = errors.New("missing required templateId parameter")
		return
	}
	path := fmt.Sprintf("automations/%s/invoke", templateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
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

type AutomationInvokeInvokeAdHocParams struct {
	Automation AutomationInvokeInvokeAdHocParamsAutomation `json:"automation,omitzero,required"`
	Brand      param.Opt[string]                           `json:"brand,omitzero"`
	Recipient  param.Opt[string]                           `json:"recipient,omitzero"`
	Template   param.Opt[string]                           `json:"template,omitzero"`
	Data       map[string]any                              `json:"data,omitzero"`
	Profile    map[string]any                              `json:"profile,omitzero"`
	paramObj
}

func (r AutomationInvokeInvokeAdHocParams) MarshalJSON() (data []byte, err error) {
	type shadow AutomationInvokeInvokeAdHocParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationInvokeInvokeAdHocParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Steps is required.
type AutomationInvokeInvokeAdHocParamsAutomation struct {
	Steps            []AutomationInvokeInvokeAdHocParamsAutomationStepUnion `json:"steps,omitzero,required"`
	CancelationToken param.Opt[string]                                      `json:"cancelation_token,omitzero"`
	paramObj
}

func (r AutomationInvokeInvokeAdHocParamsAutomation) MarshalJSON() (data []byte, err error) {
	type shadow AutomationInvokeInvokeAdHocParamsAutomation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationInvokeInvokeAdHocParamsAutomation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutomationInvokeInvokeAdHocParamsAutomationStepUnion struct {
	OfAutomationDelayStep         *AutomationInvokeInvokeAdHocParamsAutomationStepAutomationDelayStep         `json:",omitzero,inline"`
	OfAutomationSendStep          *AutomationInvokeInvokeAdHocParamsAutomationStepAutomationSendStep          `json:",omitzero,inline"`
	OfAutomationSendListStep      *AutomationInvokeInvokeAdHocParamsAutomationStepAutomationSendListStep      `json:",omitzero,inline"`
	OfAutomationUpdateProfileStep *AutomationInvokeInvokeAdHocParamsAutomationStepAutomationUpdateProfileStep `json:",omitzero,inline"`
	OfAutomationCancelStep        *AutomationInvokeInvokeAdHocParamsAutomationStepAutomationCancelStep        `json:",omitzero,inline"`
	OfAutomationFetchDataStep     *AutomationInvokeInvokeAdHocParamsAutomationStepAutomationFetchDataStep     `json:",omitzero,inline"`
	OfAutomationInvokeStep        *AutomationInvokeInvokeAdHocParamsAutomationStepAutomationInvokeStep        `json:",omitzero,inline"`
	paramUnion
}

func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAutomationDelayStep,
		u.OfAutomationSendStep,
		u.OfAutomationSendListStep,
		u.OfAutomationUpdateProfileStep,
		u.OfAutomationCancelStep,
		u.OfAutomationFetchDataStep,
		u.OfAutomationInvokeStep)
}
func (u *AutomationInvokeInvokeAdHocParamsAutomationStepUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AutomationInvokeInvokeAdHocParamsAutomationStepUnion) asAny() any {
	if !param.IsOmitted(u.OfAutomationDelayStep) {
		return u.OfAutomationDelayStep
	} else if !param.IsOmitted(u.OfAutomationSendStep) {
		return u.OfAutomationSendStep
	} else if !param.IsOmitted(u.OfAutomationSendListStep) {
		return u.OfAutomationSendListStep
	} else if !param.IsOmitted(u.OfAutomationUpdateProfileStep) {
		return u.OfAutomationUpdateProfileStep
	} else if !param.IsOmitted(u.OfAutomationCancelStep) {
		return u.OfAutomationCancelStep
	} else if !param.IsOmitted(u.OfAutomationFetchDataStep) {
		return u.OfAutomationFetchDataStep
	} else if !param.IsOmitted(u.OfAutomationInvokeStep) {
		return u.OfAutomationInvokeStep
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetDuration() *string {
	if vt := u.OfAutomationDelayStep; vt != nil && vt.Duration.Valid() {
		return &vt.Duration.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetUntil() *string {
	if vt := u.OfAutomationDelayStep; vt != nil && vt.Until.Valid() {
		return &vt.Until.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetRecipient() *string {
	if vt := u.OfAutomationSendStep; vt != nil && vt.Recipient.Valid() {
		return &vt.Recipient.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetList() *string {
	if vt := u.OfAutomationSendListStep; vt != nil {
		return &vt.List
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetMerge() *string {
	if vt := u.OfAutomationUpdateProfileStep; vt != nil {
		return &vt.Merge
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetRecipientID() *string {
	if vt := u.OfAutomationUpdateProfileStep; vt != nil && vt.RecipientID.Valid() {
		return &vt.RecipientID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetCancelationToken() *string {
	if vt := u.OfAutomationCancelStep; vt != nil {
		return &vt.CancelationToken
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetWebhook() *AutomationInvokeInvokeAdHocParamsAutomationStepAutomationFetchDataStepWebhook {
	if vt := u.OfAutomationFetchDataStep; vt != nil {
		return &vt.Webhook
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetMergeStrategy() *string {
	if vt := u.OfAutomationFetchDataStep; vt != nil {
		return &vt.MergeStrategy
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetAction() *string {
	if vt := u.OfAutomationDelayStep; vt != nil {
		return (*string)(&vt.Action)
	} else if vt := u.OfAutomationSendStep; vt != nil {
		return (*string)(&vt.Action)
	} else if vt := u.OfAutomationSendListStep; vt != nil {
		return (*string)(&vt.Action)
	} else if vt := u.OfAutomationUpdateProfileStep; vt != nil {
		return (*string)(&vt.Action)
	} else if vt := u.OfAutomationCancelStep; vt != nil {
		return (*string)(&vt.Action)
	} else if vt := u.OfAutomationFetchDataStep; vt != nil {
		return (*string)(&vt.Action)
	} else if vt := u.OfAutomationInvokeStep; vt != nil {
		return (*string)(&vt.Action)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetBrand() *string {
	if vt := u.OfAutomationSendStep; vt != nil && vt.Brand.Valid() {
		return &vt.Brand.Value
	} else if vt := u.OfAutomationSendListStep; vt != nil && vt.Brand.Valid() {
		return &vt.Brand.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetTemplate() *string {
	if vt := u.OfAutomationSendStep; vt != nil && vt.Template.Valid() {
		return &vt.Template.Value
	} else if vt := u.OfAutomationInvokeStep; vt != nil {
		return (*string)(&vt.Template)
	}
	return nil
}

// Returns a pointer to the underlying variant's Data property, if present.
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetData() map[string]any {
	if vt := u.OfAutomationSendStep; vt != nil {
		return vt.Data
	} else if vt := u.OfAutomationSendListStep; vt != nil {
		return vt.Data
	}
	return nil
}

// Returns a pointer to the underlying variant's Profile property, if present.
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetProfile() map[string]any {
	if vt := u.OfAutomationSendStep; vt != nil {
		return vt.Profile
	} else if vt := u.OfAutomationUpdateProfileStep; vt != nil {
		return vt.Profile
	}
	return nil
}

// The property Action is required.
type AutomationInvokeInvokeAdHocParamsAutomationStepAutomationDelayStep struct {
	// Any of "delay".
	Action   string            `json:"action,omitzero,required"`
	Duration param.Opt[string] `json:"duration,omitzero"`
	Until    param.Opt[string] `json:"until,omitzero"`
	paramObj
}

func (r AutomationInvokeInvokeAdHocParamsAutomationStepAutomationDelayStep) MarshalJSON() (data []byte, err error) {
	type shadow AutomationInvokeInvokeAdHocParamsAutomationStepAutomationDelayStep
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationInvokeInvokeAdHocParamsAutomationStepAutomationDelayStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationInvokeInvokeAdHocParamsAutomationStepAutomationDelayStep](
		"action", "delay",
	)
}

// The property Action is required.
type AutomationInvokeInvokeAdHocParamsAutomationStepAutomationSendStep struct {
	// Any of "send".
	Action    string            `json:"action,omitzero,required"`
	Brand     param.Opt[string] `json:"brand,omitzero"`
	Recipient param.Opt[string] `json:"recipient,omitzero"`
	Template  param.Opt[string] `json:"template,omitzero"`
	Data      map[string]any    `json:"data,omitzero"`
	Profile   map[string]any    `json:"profile,omitzero"`
	paramObj
}

func (r AutomationInvokeInvokeAdHocParamsAutomationStepAutomationSendStep) MarshalJSON() (data []byte, err error) {
	type shadow AutomationInvokeInvokeAdHocParamsAutomationStepAutomationSendStep
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationInvokeInvokeAdHocParamsAutomationStepAutomationSendStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationInvokeInvokeAdHocParamsAutomationStepAutomationSendStep](
		"action", "send",
	)
}

// The properties Action, List are required.
type AutomationInvokeInvokeAdHocParamsAutomationStepAutomationSendListStep struct {
	// Any of "send-list".
	Action string            `json:"action,omitzero,required"`
	List   string            `json:"list,required"`
	Brand  param.Opt[string] `json:"brand,omitzero"`
	Data   map[string]any    `json:"data,omitzero"`
	paramObj
}

func (r AutomationInvokeInvokeAdHocParamsAutomationStepAutomationSendListStep) MarshalJSON() (data []byte, err error) {
	type shadow AutomationInvokeInvokeAdHocParamsAutomationStepAutomationSendListStep
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationInvokeInvokeAdHocParamsAutomationStepAutomationSendListStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationInvokeInvokeAdHocParamsAutomationStepAutomationSendListStep](
		"action", "send-list",
	)
}

// The properties Action, Profile are required.
type AutomationInvokeInvokeAdHocParamsAutomationStepAutomationUpdateProfileStep struct {
	// Any of "update-profile".
	Action      string            `json:"action,omitzero,required"`
	Profile     map[string]any    `json:"profile,omitzero,required"`
	RecipientID param.Opt[string] `json:"recipient_id,omitzero"`
	// Any of "none", "overwrite", "soft-merge", "replace".
	Merge string `json:"merge,omitzero"`
	paramObj
}

func (r AutomationInvokeInvokeAdHocParamsAutomationStepAutomationUpdateProfileStep) MarshalJSON() (data []byte, err error) {
	type shadow AutomationInvokeInvokeAdHocParamsAutomationStepAutomationUpdateProfileStep
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationInvokeInvokeAdHocParamsAutomationStepAutomationUpdateProfileStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationInvokeInvokeAdHocParamsAutomationStepAutomationUpdateProfileStep](
		"action", "update-profile",
	)
	apijson.RegisterFieldValidator[AutomationInvokeInvokeAdHocParamsAutomationStepAutomationUpdateProfileStep](
		"merge", "none", "overwrite", "soft-merge", "replace",
	)
}

// The properties Action, CancelationToken are required.
type AutomationInvokeInvokeAdHocParamsAutomationStepAutomationCancelStep struct {
	// Any of "cancel".
	Action           string `json:"action,omitzero,required"`
	CancelationToken string `json:"cancelation_token,required"`
	paramObj
}

func (r AutomationInvokeInvokeAdHocParamsAutomationStepAutomationCancelStep) MarshalJSON() (data []byte, err error) {
	type shadow AutomationInvokeInvokeAdHocParamsAutomationStepAutomationCancelStep
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationInvokeInvokeAdHocParamsAutomationStepAutomationCancelStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationInvokeInvokeAdHocParamsAutomationStepAutomationCancelStep](
		"action", "cancel",
	)
}

// The properties Action, Webhook are required.
type AutomationInvokeInvokeAdHocParamsAutomationStepAutomationFetchDataStep struct {
	// Any of "fetch-data".
	Action  string                                                                        `json:"action,omitzero,required"`
	Webhook AutomationInvokeInvokeAdHocParamsAutomationStepAutomationFetchDataStepWebhook `json:"webhook,omitzero,required"`
	// Any of "replace", "overwrite", "soft-merge".
	MergeStrategy string `json:"merge_strategy,omitzero"`
	paramObj
}

func (r AutomationInvokeInvokeAdHocParamsAutomationStepAutomationFetchDataStep) MarshalJSON() (data []byte, err error) {
	type shadow AutomationInvokeInvokeAdHocParamsAutomationStepAutomationFetchDataStep
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationInvokeInvokeAdHocParamsAutomationStepAutomationFetchDataStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationInvokeInvokeAdHocParamsAutomationStepAutomationFetchDataStep](
		"action", "fetch-data",
	)
	apijson.RegisterFieldValidator[AutomationInvokeInvokeAdHocParamsAutomationStepAutomationFetchDataStep](
		"merge_strategy", "replace", "overwrite", "soft-merge",
	)
}

// The properties Method, URL are required.
type AutomationInvokeInvokeAdHocParamsAutomationStepAutomationFetchDataStepWebhook struct {
	// Any of "GET", "POST", "PUT", "PATCH", "DELETE".
	Method  string            `json:"method,omitzero,required"`
	URL     string            `json:"url,required"`
	Body    param.Opt[string] `json:"body,omitzero"`
	Headers map[string]string `json:"headers,omitzero"`
	paramObj
}

func (r AutomationInvokeInvokeAdHocParamsAutomationStepAutomationFetchDataStepWebhook) MarshalJSON() (data []byte, err error) {
	type shadow AutomationInvokeInvokeAdHocParamsAutomationStepAutomationFetchDataStepWebhook
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationInvokeInvokeAdHocParamsAutomationStepAutomationFetchDataStepWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationInvokeInvokeAdHocParamsAutomationStepAutomationFetchDataStepWebhook](
		"method", "GET", "POST", "PUT", "PATCH", "DELETE",
	)
}

// The properties Action, Template are required.
type AutomationInvokeInvokeAdHocParamsAutomationStepAutomationInvokeStep struct {
	// Any of "invoke".
	Action   string `json:"action,omitzero,required"`
	Template string `json:"template,required"`
	paramObj
}

func (r AutomationInvokeInvokeAdHocParamsAutomationStepAutomationInvokeStep) MarshalJSON() (data []byte, err error) {
	type shadow AutomationInvokeInvokeAdHocParamsAutomationStepAutomationInvokeStep
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationInvokeInvokeAdHocParamsAutomationStepAutomationInvokeStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationInvokeInvokeAdHocParamsAutomationStepAutomationInvokeStep](
		"action", "invoke",
	)
}

type AutomationInvokeInvokeByTemplateParams struct {
	Recipient param.Opt[string] `json:"recipient,omitzero,required"`
	Brand     param.Opt[string] `json:"brand,omitzero"`
	Template  param.Opt[string] `json:"template,omitzero"`
	Data      map[string]any    `json:"data,omitzero"`
	Profile   map[string]any    `json:"profile,omitzero"`
	paramObj
}

func (r AutomationInvokeInvokeByTemplateParams) MarshalJSON() (data []byte, err error) {
	type shadow AutomationInvokeInvokeByTemplateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationInvokeInvokeByTemplateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
