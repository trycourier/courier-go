// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/trycourier/courier-go/internal/apijson"
	shimjson "github.com/trycourier/courier-go/internal/encoding/json"
	"github.com/trycourier/courier-go/internal/requestconfig"
	"github.com/trycourier/courier-go/option"
	"github.com/trycourier/courier-go/packages/param"
	"github.com/trycourier/courier-go/packages/respjson"
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

type AutomationInvokeInvokeAdHocParams struct {
	Automation AutomationInvokeInvokeAdHocParamsAutomation `json:"automation,omitzero,required"`
	Brand      param.Opt[string]                           `json:"brand,omitzero"`
	Recipient  param.Opt[string]                           `json:"recipient,omitzero"`
	Template   param.Opt[string]                           `json:"template,omitzero"`
	Data       map[string]any                              `json:"data,omitzero"`
	Profile    any                                         `json:"profile,omitzero"`
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
	OfAutomationAddToDigestStep   *AutomationInvokeInvokeAdHocParamsAutomationStepAutomationAddToDigestStep   `json:",omitzero,inline"`
	OfAutomationAddToBatchStep    *AutomationInvokeInvokeAdHocParamsAutomationStepAutomationAddToBatchStep    `json:",omitzero,inline"`
	OfAutomationThrottleStep      *AutomationInvokeInvokeAdHocParamsAutomationStepAutomationThrottleStep      `json:",omitzero,inline"`
	OfAutomationCancelStep        *AutomationInvokeInvokeAdHocParamsAutomationStepAutomationCancelStep        `json:",omitzero,inline"`
	OfAutomationDelayStep         *AutomationInvokeInvokeAdHocParamsAutomationStepAutomationDelayStep         `json:",omitzero,inline"`
	OfAutomationFetchDataStep     *AutomationInvokeInvokeAdHocParamsAutomationStepAutomationFetchDataStep     `json:",omitzero,inline"`
	OfAutomationInvokeStep        *AutomationInvokeInvokeAdHocParamsAutomationStepAutomationInvokeStep        `json:",omitzero,inline"`
	OfAutomationSendStep          *AutomationInvokeInvokeAdHocParamsAutomationStepAutomationSendStep          `json:",omitzero,inline"`
	OfAutomationV2SendStep        *AutomationInvokeInvokeAdHocParamsAutomationStepAutomationV2SendStep        `json:",omitzero,inline"`
	OfAutomationSendListStep      *AutomationInvokeInvokeAdHocParamsAutomationStepAutomationSendListStep      `json:",omitzero,inline"`
	OfAutomationUpdateProfileStep *AutomationInvokeInvokeAdHocParamsAutomationStepAutomationUpdateProfileStep `json:",omitzero,inline"`
	paramUnion
}

func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAutomationAddToDigestStep,
		u.OfAutomationAddToBatchStep,
		u.OfAutomationThrottleStep,
		u.OfAutomationCancelStep,
		u.OfAutomationDelayStep,
		u.OfAutomationFetchDataStep,
		u.OfAutomationInvokeStep,
		u.OfAutomationSendStep,
		u.OfAutomationV2SendStep,
		u.OfAutomationSendListStep,
		u.OfAutomationUpdateProfileStep)
}
func (u *AutomationInvokeInvokeAdHocParamsAutomationStepUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AutomationInvokeInvokeAdHocParamsAutomationStepUnion) asAny() any {
	if !param.IsOmitted(u.OfAutomationAddToDigestStep) {
		return u.OfAutomationAddToDigestStep
	} else if !param.IsOmitted(u.OfAutomationAddToBatchStep) {
		return u.OfAutomationAddToBatchStep
	} else if !param.IsOmitted(u.OfAutomationThrottleStep) {
		return u.OfAutomationThrottleStep
	} else if !param.IsOmitted(u.OfAutomationCancelStep) {
		return u.OfAutomationCancelStep
	} else if !param.IsOmitted(u.OfAutomationDelayStep) {
		return u.OfAutomationDelayStep
	} else if !param.IsOmitted(u.OfAutomationFetchDataStep) {
		return u.OfAutomationFetchDataStep
	} else if !param.IsOmitted(u.OfAutomationInvokeStep) {
		return u.OfAutomationInvokeStep
	} else if !param.IsOmitted(u.OfAutomationSendStep) {
		return u.OfAutomationSendStep
	} else if !param.IsOmitted(u.OfAutomationV2SendStep) {
		return u.OfAutomationV2SendStep
	} else if !param.IsOmitted(u.OfAutomationSendListStep) {
		return u.OfAutomationSendListStep
	} else if !param.IsOmitted(u.OfAutomationUpdateProfileStep) {
		return u.OfAutomationUpdateProfileStep
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetSubscriptionTopicID() *string {
	if vt := u.OfAutomationAddToDigestStep; vt != nil {
		return &vt.SubscriptionTopicID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetMaxWaitPeriod() *string {
	if vt := u.OfAutomationAddToBatchStep; vt != nil {
		return &vt.MaxWaitPeriod
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetRetain() *AutomationInvokeInvokeAdHocParamsAutomationStepAutomationAddToBatchStepRetain {
	if vt := u.OfAutomationAddToBatchStep; vt != nil {
		return &vt.Retain
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetWaitPeriod() *string {
	if vt := u.OfAutomationAddToBatchStep; vt != nil {
		return &vt.WaitPeriod
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetBatchID() *string {
	if vt := u.OfAutomationAddToBatchStep; vt != nil && vt.BatchID.Valid() {
		return &vt.BatchID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetBatchKey() *string {
	if vt := u.OfAutomationAddToBatchStep; vt != nil && vt.BatchKey.Valid() {
		return &vt.BatchKey.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetCategoryKey() *string {
	if vt := u.OfAutomationAddToBatchStep; vt != nil && vt.CategoryKey.Valid() {
		return &vt.CategoryKey.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetMaxItems() *AutomationInvokeInvokeAdHocParamsAutomationStepAutomationAddToBatchStepMaxItemsUnion {
	if vt := u.OfAutomationAddToBatchStep; vt != nil {
		return &vt.MaxItems
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetMaxAllowed() *int64 {
	if vt := u.OfAutomationThrottleStep; vt != nil {
		return &vt.MaxAllowed
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetOnThrottle() *AutomationInvokeInvokeAdHocParamsAutomationStepAutomationThrottleStepOnThrottle {
	if vt := u.OfAutomationThrottleStep; vt != nil {
		return &vt.OnThrottle
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetPeriod() *string {
	if vt := u.OfAutomationThrottleStep; vt != nil {
		return &vt.Period
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetShouldAlert() *bool {
	if vt := u.OfAutomationThrottleStep; vt != nil {
		return &vt.ShouldAlert
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetThrottleKey() *string {
	if vt := u.OfAutomationThrottleStep; vt != nil && vt.ThrottleKey.Valid() {
		return &vt.ThrottleKey.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetCancelationToken() *string {
	if vt := u.OfAutomationCancelStep; vt != nil && vt.CancelationToken.Valid() {
		return &vt.CancelationToken.Value
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
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetMergeStrategy() *string {
	if vt := u.OfAutomationFetchDataStep; vt != nil {
		return (*string)(&vt.MergeStrategy)
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
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetIdempotencyExpiry() *string {
	if vt := u.OfAutomationFetchDataStep; vt != nil && vt.IdempotencyExpiry.Valid() {
		return &vt.IdempotencyExpiry.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetIdempotencyKey() *string {
	if vt := u.OfAutomationFetchDataStep; vt != nil && vt.IdempotencyKey.Valid() {
		return &vt.IdempotencyKey.Value
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
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetMessage() *MessageUnionParam {
	if vt := u.OfAutomationV2SendStep; vt != nil {
		return &vt.Message
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
		return (*string)(&vt.Merge)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetRecipientID() *string {
	if vt := u.OfAutomationUpdateProfileStep; vt != nil {
		return &vt.RecipientID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetIf() *string {
	if vt := u.OfAutomationAddToDigestStep; vt != nil && vt.If.Valid() {
		return &vt.If.Value
	} else if vt := u.OfAutomationAddToBatchStep; vt != nil && vt.If.Valid() {
		return &vt.If.Value
	} else if vt := u.OfAutomationThrottleStep; vt != nil && vt.If.Valid() {
		return &vt.If.Value
	} else if vt := u.OfAutomationCancelStep; vt != nil && vt.If.Valid() {
		return &vt.If.Value
	} else if vt := u.OfAutomationDelayStep; vt != nil && vt.If.Valid() {
		return &vt.If.Value
	} else if vt := u.OfAutomationFetchDataStep; vt != nil && vt.If.Valid() {
		return &vt.If.Value
	} else if vt := u.OfAutomationInvokeStep; vt != nil && vt.If.Valid() {
		return &vt.If.Value
	} else if vt := u.OfAutomationSendStep; vt != nil && vt.If.Valid() {
		return &vt.If.Value
	} else if vt := u.OfAutomationV2SendStep; vt != nil && vt.If.Valid() {
		return &vt.If.Value
	} else if vt := u.OfAutomationSendListStep; vt != nil && vt.If.Valid() {
		return &vt.If.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetRef() *string {
	if vt := u.OfAutomationAddToDigestStep; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	} else if vt := u.OfAutomationAddToBatchStep; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	} else if vt := u.OfAutomationThrottleStep; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	} else if vt := u.OfAutomationCancelStep; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	} else if vt := u.OfAutomationDelayStep; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	} else if vt := u.OfAutomationFetchDataStep; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	} else if vt := u.OfAutomationInvokeStep; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	} else if vt := u.OfAutomationSendStep; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	} else if vt := u.OfAutomationV2SendStep; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	} else if vt := u.OfAutomationSendListStep; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetAction() *string {
	if vt := u.OfAutomationAddToDigestStep; vt != nil {
		return (*string)(&vt.Action)
	} else if vt := u.OfAutomationAddToBatchStep; vt != nil {
		return (*string)(&vt.Action)
	} else if vt := u.OfAutomationThrottleStep; vt != nil {
		return (*string)(&vt.Action)
	} else if vt := u.OfAutomationCancelStep; vt != nil {
		return (*string)(&vt.Action)
	} else if vt := u.OfAutomationDelayStep; vt != nil {
		return (*string)(&vt.Action)
	} else if vt := u.OfAutomationFetchDataStep; vt != nil {
		return (*string)(&vt.Action)
	} else if vt := u.OfAutomationInvokeStep; vt != nil {
		return (*string)(&vt.Action)
	} else if vt := u.OfAutomationSendStep; vt != nil {
		return (*string)(&vt.Action)
	} else if vt := u.OfAutomationV2SendStep; vt != nil {
		return (*string)(&vt.Action)
	} else if vt := u.OfAutomationSendListStep; vt != nil {
		return (*string)(&vt.Action)
	} else if vt := u.OfAutomationUpdateProfileStep; vt != nil {
		return (*string)(&vt.Action)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetScope() *string {
	if vt := u.OfAutomationAddToBatchStep; vt != nil {
		return (*string)(&vt.Scope)
	} else if vt := u.OfAutomationThrottleStep; vt != nil {
		return (*string)(&vt.Scope)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetTemplate() *string {
	if vt := u.OfAutomationInvokeStep; vt != nil {
		return (*string)(&vt.Template)
	} else if vt := u.OfAutomationSendStep; vt != nil && vt.Template.Valid() {
		return &vt.Template.Value
	} else if vt := u.OfAutomationSendListStep; vt != nil && vt.Template.Valid() {
		return &vt.Template.Value
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

// Returns a pointer to the underlying variant's Data property, if present.
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetData() map[string]any {
	if vt := u.OfAutomationSendStep; vt != nil {
		return vt.Data
	} else if vt := u.OfAutomationSendListStep; vt != nil {
		return vt.Data
	}
	return nil
}

// Returns a pointer to the underlying variant's Override property, if present.
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetOverride() map[string]any {
	if vt := u.OfAutomationSendStep; vt != nil {
		return vt.Override
	} else if vt := u.OfAutomationSendListStep; vt != nil {
		return vt.Override
	}
	return nil
}

// Returns a pointer to the underlying variant's Profile property, if present.
func (u AutomationInvokeInvokeAdHocParamsAutomationStepUnion) GetProfile() *any {
	if vt := u.OfAutomationSendStep; vt != nil {
		return &vt.Profile
	} else if vt := u.OfAutomationUpdateProfileStep; vt != nil {
		return &vt.Profile
	}
	return nil
}

type AutomationInvokeInvokeAdHocParamsAutomationStepAutomationAddToDigestStep struct {
	Action string `json:"action,omitzero,required"`
	// The subscription topic that has digests enabled
	SubscriptionTopicID string `json:"subscription_topic_id,required"`
	AutomationStepParam
}

func (r AutomationInvokeInvokeAdHocParamsAutomationStepAutomationAddToDigestStep) MarshalJSON() (data []byte, err error) {
	type shadow AutomationInvokeInvokeAdHocParamsAutomationStepAutomationAddToDigestStep
	return param.MarshalObject(r, (*shadow)(&r))
}

type AutomationInvokeInvokeAdHocParamsAutomationStepAutomationAddToBatchStep struct {
	Action string `json:"action,omitzero,required"`
	// Defines the maximum wait time before the batch should be released. Must be less
	// than wait period. Maximum of 60 days. Specified as an
	// [ISO 8601 duration](https://en.wikipedia.org/wiki/ISO_8601#Durations)
	MaxWaitPeriod string `json:"max_wait_period,required"`
	// Defines what items should be retained and passed along to the next steps when
	// the batch is released
	Retain AutomationInvokeInvokeAdHocParamsAutomationStepAutomationAddToBatchStepRetain `json:"retain,omitzero,required"`
	// Defines the period of inactivity before the batch is released. Specified as an
	// [ISO 8601 duration](https://en.wikipedia.org/wiki/ISO_8601#Durations)
	WaitPeriod string            `json:"wait_period,required"`
	BatchID    param.Opt[string] `json:"batch_id,omitzero"`
	// If using scope=dynamic, provide the key or a reference (e.g.,
	// refs.data.batch_key)
	BatchKey param.Opt[string] `json:"batch_key,omitzero"`
	// Defines the field of the data object the batch is set to when complete. Defaults
	// to `batch`
	CategoryKey param.Opt[string] `json:"category_key,omitzero"`
	// If specified, the batch will release as soon as this number is reached
	MaxItems AutomationInvokeInvokeAdHocParamsAutomationStepAutomationAddToBatchStepMaxItemsUnion `json:"max_items,omitzero"`
	// Determine the scope of the batching. If user, chosen in this order: recipient,
	// profile.user_id, data.user_id, data.userId. If dynamic, then specify where the
	// batch_key or a reference to the batch_key
	Scope string `json:"scope,omitzero"`
	AutomationStepParam
}

func (r AutomationInvokeInvokeAdHocParamsAutomationStepAutomationAddToBatchStep) MarshalJSON() (data []byte, err error) {
	type shadow AutomationInvokeInvokeAdHocParamsAutomationStepAutomationAddToBatchStep
	return param.MarshalObject(r, (*shadow)(&r))
}

// Defines what items should be retained and passed along to the next steps when
// the batch is released
//
// The properties Count, Type are required.
type AutomationInvokeInvokeAdHocParamsAutomationStepAutomationAddToBatchStepRetain struct {
	// The number of records to keep in batch. Default is 10 and only configurable by
	// requesting from support. When configurable minimum is 2 and maximum is 100.
	Count int64 `json:"count,required"`
	// Keep N number of notifications based on the type. First/Last N based on
	// notification received. highest/lowest based on a scoring key providing in the
	// data accessed by sort_key
	//
	// Any of "first", "last", "highest", "lowest".
	Type string `json:"type,omitzero,required"`
	// Defines the data value data[sort_key] that is used to sort the stored items.
	// Required when type is set to highest or lowest.
	SortKey param.Opt[string] `json:"sort_key,omitzero"`
	paramObj
}

func (r AutomationInvokeInvokeAdHocParamsAutomationStepAutomationAddToBatchStepRetain) MarshalJSON() (data []byte, err error) {
	type shadow AutomationInvokeInvokeAdHocParamsAutomationStepAutomationAddToBatchStepRetain
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationInvokeInvokeAdHocParamsAutomationStepAutomationAddToBatchStepRetain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AutomationInvokeInvokeAdHocParamsAutomationStepAutomationAddToBatchStepRetain](
		"type", "first", "last", "highest", "lowest",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AutomationInvokeInvokeAdHocParamsAutomationStepAutomationAddToBatchStepMaxItemsUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u AutomationInvokeInvokeAdHocParamsAutomationStepAutomationAddToBatchStepMaxItemsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *AutomationInvokeInvokeAdHocParamsAutomationStepAutomationAddToBatchStepMaxItemsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AutomationInvokeInvokeAdHocParamsAutomationStepAutomationAddToBatchStepMaxItemsUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

type AutomationInvokeInvokeAdHocParamsAutomationStepAutomationThrottleStep struct {
	Action string `json:"action,omitzero,required"`
	// Maximum number of allowed notifications in that timeframe
	MaxAllowed int64                                                                           `json:"max_allowed,required"`
	OnThrottle AutomationInvokeInvokeAdHocParamsAutomationStepAutomationThrottleStepOnThrottle `json:"on_throttle,omitzero,required"`
	// Defines the throttle period which corresponds to the max_allowed. Specified as
	// an ISO 8601 duration, https://en.wikipedia.org/wiki/ISO_8601#Durations
	Period string `json:"period,required"`
	Scope  string `json:"scope,omitzero,required"`
	// Value must be true
	ShouldAlert bool `json:"should_alert,required"`
	// If using scope=dynamic, provide the reference (e.g., refs.data.throttle_key) to
	// the how the throttle should be identified
	ThrottleKey param.Opt[string] `json:"throttle_key,omitzero"`
	AutomationStepParam
}

func (r AutomationInvokeInvokeAdHocParamsAutomationStepAutomationThrottleStep) MarshalJSON() (data []byte, err error) {
	type shadow AutomationInvokeInvokeAdHocParamsAutomationStepAutomationThrottleStep
	return param.MarshalObject(r, (*shadow)(&r))
}

// The property NodeID is required.
type AutomationInvokeInvokeAdHocParamsAutomationStepAutomationThrottleStepOnThrottle struct {
	// The node to go to if the request is throttled
	NodeID string `json:"$node_id,required"`
	paramObj
}

func (r AutomationInvokeInvokeAdHocParamsAutomationStepAutomationThrottleStepOnThrottle) MarshalJSON() (data []byte, err error) {
	type shadow AutomationInvokeInvokeAdHocParamsAutomationStepAutomationThrottleStepOnThrottle
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomationInvokeInvokeAdHocParamsAutomationStepAutomationThrottleStepOnThrottle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomationInvokeInvokeAdHocParamsAutomationStepAutomationCancelStep struct {
	Action           string            `json:"action,omitzero,required"`
	CancelationToken param.Opt[string] `json:"cancelation_token,omitzero"`
	AutomationStepParam
}

func (r AutomationInvokeInvokeAdHocParamsAutomationStepAutomationCancelStep) MarshalJSON() (data []byte, err error) {
	type shadow AutomationInvokeInvokeAdHocParamsAutomationStepAutomationCancelStep
	return param.MarshalObject(r, (*shadow)(&r))
}

type AutomationInvokeInvokeAdHocParamsAutomationStepAutomationDelayStep struct {
	Action string `json:"action,omitzero,required"`
	// The [ISO 8601 duration](https://en.wikipedia.org/wiki/ISO_8601#Durations) string
	// for how long to delay for
	Duration param.Opt[string] `json:"duration,omitzero"`
	// The ISO 8601 timestamp for when the delay should end
	Until param.Opt[string] `json:"until,omitzero"`
	AutomationStepParam
}

func (r AutomationInvokeInvokeAdHocParamsAutomationStepAutomationDelayStep) MarshalJSON() (data []byte, err error) {
	type shadow AutomationInvokeInvokeAdHocParamsAutomationStepAutomationDelayStep
	return param.MarshalObject(r, (*shadow)(&r))
}

type AutomationInvokeInvokeAdHocParamsAutomationStepAutomationFetchDataStep struct {
	Action            string                                                                        `json:"action,omitzero,required"`
	MergeStrategy     MergeAlgorithm                                                                `json:"merge_strategy,omitzero,required"`
	Webhook           AutomationInvokeInvokeAdHocParamsAutomationStepAutomationFetchDataStepWebhook `json:"webhook,omitzero,required"`
	IdempotencyExpiry param.Opt[string]                                                             `json:"idempotency_expiry,omitzero"`
	IdempotencyKey    param.Opt[string]                                                             `json:"idempotency_key,omitzero"`
	AutomationStepParam
}

func (r AutomationInvokeInvokeAdHocParamsAutomationStepAutomationFetchDataStep) MarshalJSON() (data []byte, err error) {
	type shadow AutomationInvokeInvokeAdHocParamsAutomationStepAutomationFetchDataStep
	return param.MarshalObject(r, (*shadow)(&r))
}

// The properties Method, URL are required.
type AutomationInvokeInvokeAdHocParamsAutomationStepAutomationFetchDataStepWebhook struct {
	// Any of "GET", "POST".
	Method  string         `json:"method,omitzero,required"`
	URL     string         `json:"url,required"`
	Body    map[string]any `json:"body,omitzero"`
	Headers map[string]any `json:"headers,omitzero"`
	Params  map[string]any `json:"params,omitzero"`
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
		"method", "GET", "POST",
	)
}

type AutomationInvokeInvokeAdHocParamsAutomationStepAutomationInvokeStep struct {
	Action   string `json:"action,omitzero,required"`
	Template string `json:"template,required"`
	AutomationStepParam
}

func (r AutomationInvokeInvokeAdHocParamsAutomationStepAutomationInvokeStep) MarshalJSON() (data []byte, err error) {
	type shadow AutomationInvokeInvokeAdHocParamsAutomationStepAutomationInvokeStep
	return param.MarshalObject(r, (*shadow)(&r))
}

type AutomationInvokeInvokeAdHocParamsAutomationStepAutomationSendStep struct {
	Action    string            `json:"action,omitzero,required"`
	Brand     param.Opt[string] `json:"brand,omitzero"`
	Data      map[string]any    `json:"data,omitzero"`
	Override  map[string]any    `json:"override,omitzero"`
	Profile   any               `json:"profile,omitzero"`
	Recipient param.Opt[string] `json:"recipient,omitzero"`
	Template  param.Opt[string] `json:"template,omitzero"`
	AutomationStepParam
}

func (r AutomationInvokeInvokeAdHocParamsAutomationStepAutomationSendStep) MarshalJSON() (data []byte, err error) {
	type shadow AutomationInvokeInvokeAdHocParamsAutomationStepAutomationSendStep
	return param.MarshalObject(r, (*shadow)(&r))
}

type AutomationInvokeInvokeAdHocParamsAutomationStepAutomationV2SendStep struct {
	Action string `json:"action,omitzero,required"`
	// Describes the content of the message in a way that will work for email, push,
	// chat, or any channel.
	Message MessageUnionParam `json:"message,omitzero,required"`
	AutomationStepParam
}

func (r AutomationInvokeInvokeAdHocParamsAutomationStepAutomationV2SendStep) MarshalJSON() (data []byte, err error) {
	type shadow AutomationInvokeInvokeAdHocParamsAutomationStepAutomationV2SendStep
	return param.MarshalObject(r, (*shadow)(&r))
}

type AutomationInvokeInvokeAdHocParamsAutomationStepAutomationSendListStep struct {
	Action   string            `json:"action,omitzero,required"`
	List     string            `json:"list,required"`
	Brand    param.Opt[string] `json:"brand,omitzero"`
	Data     map[string]any    `json:"data,omitzero"`
	Override map[string]any    `json:"override,omitzero"`
	Template param.Opt[string] `json:"template,omitzero"`
	AutomationStepParam
}

func (r AutomationInvokeInvokeAdHocParamsAutomationStepAutomationSendListStep) MarshalJSON() (data []byte, err error) {
	type shadow AutomationInvokeInvokeAdHocParamsAutomationStepAutomationSendListStep
	return param.MarshalObject(r, (*shadow)(&r))
}

// The properties Action, Merge, Profile, RecipientID are required.
type AutomationInvokeInvokeAdHocParamsAutomationStepAutomationUpdateProfileStep struct {
	// Any of "update-profile".
	Action string `json:"action,omitzero,required"`
	// Any of "replace", "none", "overwrite", "soft-merge".
	Merge       MergeAlgorithm `json:"merge,omitzero,required"`
	Profile     any            `json:"profile,omitzero,required"`
	RecipientID string         `json:"recipient_id,required"`
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
}

type AutomationInvokeInvokeByTemplateParams struct {
	AutomationInvokeParams AutomationInvokeParams
	paramObj
}

func (r AutomationInvokeInvokeByTemplateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AutomationInvokeParams)
}
func (r *AutomationInvokeInvokeByTemplateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.AutomationInvokeParams)
}
