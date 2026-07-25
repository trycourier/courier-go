// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/trycourier/courier-go/v4/internal/apijson"
	"github.com/trycourier/courier-go/v4/internal/apiquery"
	shimjson "github.com/trycourier/courier-go/v4/internal/encoding/json"
	"github.com/trycourier/courier-go/v4/internal/requestconfig"
	"github.com/trycourier/courier-go/v4/option"
	"github.com/trycourier/courier-go/v4/packages/param"
	"github.com/trycourier/courier-go/v4/packages/respjson"
	"github.com/trycourier/courier-go/v4/shared"
)

// JourneyService contains methods and other services that help with interacting
// with the Courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewJourneyService] method instead.
type JourneyService struct {
	Options   []option.RequestOption
	Templates JourneyTemplateService
}

// NewJourneyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewJourneyService(opts ...option.RequestOption) (r JourneyService) {
	r = JourneyService{}
	r.Options = opts
	r.Templates = NewJourneyTemplateService(opts...)
	return
}

// Creates a journey from a set of nodes, in draft state unless you pass a
// published state. Send nodes cannot be included until their templates exist.
func (r *JourneyService) New(ctx context.Context, body JourneyNewParams, opts ...option.RequestOption) (res *JourneyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "journeys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Fetch a journey by id. Pass `?version=draft` (default `published`) to retrieve
// the working draft, or `?version=vN` to retrieve a historical version.
func (r *JourneyService) Get(ctx context.Context, templateID string, query JourneyGetParams, opts ...option.RequestOption) (res *JourneyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if templateID == "" {
		err = errors.New("missing required templateId parameter")
		return nil, err
	}
	path := fmt.Sprintf("journeys/%s", templateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Lists the workspace's journeys, each carrying a name, state, and enabled flag.
// Paged by cursor.
func (r *JourneyService) List(ctx context.Context, query JourneyListParams, opts ...option.RequestOption) (res *JourneysListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "journeys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Archives a journey so it can no longer be invoked. Runs already in flight
// continue to completion, so archiving never strands a user mid-sequence.
func (r *JourneyService) Archive(ctx context.Context, templateID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if templateID == "" {
		err = errors.New("missing required templateId parameter")
		return err
	}
	path := fmt.Sprintf("journeys/%s", templateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Cancels in-flight journey runs, either every run sharing a cancelation token or
// one run by id. Use it to stop a sequence when the event resolves.
func (r *JourneyService) Cancel(ctx context.Context, body JourneyCancelParams, opts ...option.RequestOption) (res *CancelJourneyResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "journeys/cancel"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Starts a journey run for one user and returns a runId. Runs execute
// asynchronously, so the response arrives before any message is sent.
func (r *JourneyService) Invoke(ctx context.Context, templateID string, body JourneyInvokeParams, opts ...option.RequestOption) (res *JourneysInvokeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if templateID == "" {
		err = errors.New("missing required templateId parameter")
		return nil, err
	}
	path := fmt.Sprintf("journeys/%s/invoke", templateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Lists a journey's published versions, most recent first, so you have a version
// id to roll back to. Paged by cursor.
func (r *JourneyService) ListVersions(ctx context.Context, templateID string, opts ...option.RequestOption) (res *JourneyVersionsListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if templateID == "" {
		err = errors.New("missing required templateId parameter")
		return nil, err
	}
	path := fmt.Sprintf("journeys/%s/versions", templateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Publishes a journey's current draft as a new version, making it live for new
// runs. Pass a version instead to roll back to an earlier one.
func (r *JourneyService) Publish(ctx context.Context, templateID string, body JourneyPublishParams, opts ...option.RequestOption) (res *JourneyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if templateID == "" {
		err = errors.New("missing required templateId parameter")
		return nil, err
	}
	path := fmt.Sprintf("journeys/%s/publish", templateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Replaces a journey's working draft, leaving the published version live until you
// publish. Reach for this when editing a journey already running.
func (r *JourneyService) Replace(ctx context.Context, templateID string, body JourneyReplaceParams, opts ...option.RequestOption) (res *JourneyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if templateID == "" {
		err = errors.New("missing required templateId parameter")
		return nil, err
	}
	path := fmt.Sprintf("journeys/%s", templateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

func CancelJourneyRequestParamOfByCancelationToken(cancelationToken string) CancelJourneyRequestUnionParam {
	var variant CancelJourneyRequestByCancelationTokenParam
	variant.CancelationToken = cancelationToken
	return CancelJourneyRequestUnionParam{OfByCancelationToken: &variant}
}

func CancelJourneyRequestParamOfByRunID(runID string) CancelJourneyRequestUnionParam {
	var variant CancelJourneyRequestByRunIDParam
	variant.RunID = runID
	return CancelJourneyRequestUnionParam{OfByRunID: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CancelJourneyRequestUnionParam struct {
	OfByCancelationToken *CancelJourneyRequestByCancelationTokenParam `json:",omitzero,inline"`
	OfByRunID            *CancelJourneyRequestByRunIDParam            `json:",omitzero,inline"`
	paramUnion
}

func (u CancelJourneyRequestUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfByCancelationToken, u.OfByRunID)
}
func (u *CancelJourneyRequestUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CancelJourneyRequestUnionParam) asAny() any {
	if !param.IsOmitted(u.OfByCancelationToken) {
		return u.OfByCancelationToken
	} else if !param.IsOmitted(u.OfByRunID) {
		return u.OfByRunID
	}
	return nil
}

// The property CancelationToken is required.
type CancelJourneyRequestByCancelationTokenParam struct {
	CancelationToken string `json:"cancelation_token" api:"required"`
	paramObj
}

func (r CancelJourneyRequestByCancelationTokenParam) MarshalJSON() (data []byte, err error) {
	type shadow CancelJourneyRequestByCancelationTokenParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CancelJourneyRequestByCancelationTokenParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property RunID is required.
type CancelJourneyRequestByRunIDParam struct {
	RunID string `json:"run_id" api:"required"`
	paramObj
}

func (r CancelJourneyRequestByRunIDParam) MarshalJSON() (data []byte, err error) {
	type shadow CancelJourneyRequestByRunIDParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CancelJourneyRequestByRunIDParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CancelJourneyResponseUnion contains all possible properties and values from
// [CancelJourneyResponseTokenBranch], [CancelJourneyResponseRunIDBranch].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type CancelJourneyResponseUnion struct {
	// This field is from variant [CancelJourneyResponseTokenBranch].
	CancelationToken string `json:"cancelation_token"`
	// This field is from variant [CancelJourneyResponseRunIDBranch].
	RunID string `json:"run_id"`
	// This field is from variant [CancelJourneyResponseRunIDBranch].
	Status string `json:"status"`
	JSON   struct {
		CancelationToken respjson.Field
		RunID            respjson.Field
		Status           respjson.Field
		raw              string
	} `json:"-"`
}

func (u CancelJourneyResponseUnion) AsTokenBranch() (v CancelJourneyResponseTokenBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CancelJourneyResponseUnion) AsRunIDBranch() (v CancelJourneyResponseRunIDBranch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CancelJourneyResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *CancelJourneyResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CancelJourneyResponseTokenBranch struct {
	CancelationToken string `json:"cancelation_token" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CancelationToken respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CancelJourneyResponseTokenBranch) RawJSON() string { return r.JSON.raw }
func (r *CancelJourneyResponseTokenBranch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CancelJourneyResponseRunIDBranch struct {
	RunID string `json:"run_id" api:"required"`
	// The run's resulting status. `CANCELED` when the run was active and has been
	// canceled; `PROCESSED` or `ERROR` when the run had already finished and was left
	// unchanged; `CANCELED` for an already-canceled run.
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RunID       respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CancelJourneyResponseRunIDBranch) RawJSON() string { return r.JSON.raw }
func (r *CancelJourneyResponseRunIDBranch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body for creating a journey.
//
// The properties Name, Nodes are required.
type CreateJourneyRequestParam struct {
	Name    string                  `json:"name" api:"required"`
	Nodes   []JourneyNodeUnionParam `json:"nodes,omitzero" api:"required"`
	Enabled param.Opt[bool]         `json:"enabled,omitzero"`
	// Lifecycle state of a journey.
	//
	// Any of "DRAFT", "PUBLISHED".
	State JourneyState `json:"state,omitzero"`
	paramObj
}

func (r CreateJourneyRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateJourneyRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateJourneyRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A journey template representing an automation workflow.
type Journey struct {
	// The unique identifier of the journey.
	ID string `json:"id" api:"required"`
	// The name of the journey.
	Name string `json:"name" api:"required"`
	// The version of the journey (published or draft).
	//
	// Any of "published", "draft".
	Version JourneyVersion `json:"version" api:"required"`
	// ISO 8601 timestamp when the journey was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// ISO 8601 timestamp when the journey was last updated.
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
func (r Journey) RawJSON() string { return r.JSON.raw }
func (r *Journey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The version of the journey (published or draft).
type JourneyVersion string

const (
	JourneyVersionPublished JourneyVersion = "published"
	JourneyVersionDraft     JourneyVersion = "draft"
)

// Invoke an AI step with `user_prompt` and optional `web_search`. Returns a
// structured response conforming to `output_schema`.
type JourneyAINode struct {
	// A JSONSchema object (Draft-07-compatible). Validated at runtime by Ajv.
	OutputSchema map[string]any `json:"output_schema" api:"required"`
	// Any of "ai".
	Type JourneyAINodeType `json:"type" api:"required"`
	ID   string            `json:"id"`
	// Condition spec for a journey node. Accepts a single condition atom, an AND/OR
	// group, or an AND/OR nested group. Omit the `conditions` property entirely to
	// express "no conditions".
	Conditions JourneyConditionsFieldUnion `json:"conditions"`
	Model      string                      `json:"model"`
	UserPrompt string                      `json:"user_prompt"`
	WebSearch  bool                        `json:"web_search"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OutputSchema respjson.Field
		Type         respjson.Field
		ID           respjson.Field
		Conditions   respjson.Field
		Model        respjson.Field
		UserPrompt   respjson.Field
		WebSearch    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JourneyAINode) RawJSON() string { return r.JSON.raw }
func (r *JourneyAINode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this JourneyAINode to a JourneyAINodeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// JourneyAINodeParam.Overrides()
func (r JourneyAINode) ToParam() JourneyAINodeParam {
	return param.Override[JourneyAINodeParam](json.RawMessage(r.RawJSON()))
}

type JourneyAINodeType string

const (
	JourneyAINodeTypeAI JourneyAINodeType = "ai"
)

// Invoke an AI step with `user_prompt` and optional `web_search`. Returns a
// structured response conforming to `output_schema`.
//
// The properties OutputSchema, Type are required.
type JourneyAINodeParam struct {
	// A JSONSchema object (Draft-07-compatible). Validated at runtime by Ajv.
	OutputSchema map[string]any `json:"output_schema,omitzero" api:"required"`
	// Any of "ai".
	Type       JourneyAINodeType `json:"type,omitzero" api:"required"`
	ID         param.Opt[string] `json:"id,omitzero"`
	Model      param.Opt[string] `json:"model,omitzero"`
	UserPrompt param.Opt[string] `json:"user_prompt,omitzero"`
	WebSearch  param.Opt[bool]   `json:"web_search,omitzero"`
	// Condition spec for a journey node. Accepts a single condition atom, an AND/OR
	// group, or an AND/OR nested group. Omit the `conditions` property entirely to
	// express "no conditions".
	Conditions JourneyConditionsFieldUnionParam `json:"conditions,omitzero"`
	paramObj
}

func (r JourneyAINodeParam) MarshalJSON() (data []byte, err error) {
	type shadow JourneyAINodeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JourneyAINodeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Trigger fired when the journey is invoked via the API. The optional `schema`
// field is a JSON Schema that validates the invocation payload.
type JourneyAPIInvokeTriggerNode struct {
	// Any of "api-invoke".
	TriggerType JourneyAPIInvokeTriggerNodeTriggerType `json:"trigger_type" api:"required"`
	// Any of "trigger".
	Type JourneyAPIInvokeTriggerNodeType `json:"type" api:"required"`
	ID   string                          `json:"id"`
	// Condition spec for a journey node. Accepts a single condition atom, an AND/OR
	// group, or an AND/OR nested group. Omit the `conditions` property entirely to
	// express "no conditions".
	Conditions JourneyConditionsFieldUnion `json:"conditions"`
	// A JSONSchema object (Draft-07-compatible). Validated at runtime by Ajv.
	Schema map[string]any `json:"schema"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TriggerType respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		Conditions  respjson.Field
		Schema      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JourneyAPIInvokeTriggerNode) RawJSON() string { return r.JSON.raw }
func (r *JourneyAPIInvokeTriggerNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this JourneyAPIInvokeTriggerNode to a
// JourneyAPIInvokeTriggerNodeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// JourneyAPIInvokeTriggerNodeParam.Overrides()
func (r JourneyAPIInvokeTriggerNode) ToParam() JourneyAPIInvokeTriggerNodeParam {
	return param.Override[JourneyAPIInvokeTriggerNodeParam](json.RawMessage(r.RawJSON()))
}

type JourneyAPIInvokeTriggerNodeTriggerType string

const (
	JourneyAPIInvokeTriggerNodeTriggerTypeAPIInvoke JourneyAPIInvokeTriggerNodeTriggerType = "api-invoke"
)

type JourneyAPIInvokeTriggerNodeType string

const (
	JourneyAPIInvokeTriggerNodeTypeTrigger JourneyAPIInvokeTriggerNodeType = "trigger"
)

// Trigger fired when the journey is invoked via the API. The optional `schema`
// field is a JSON Schema that validates the invocation payload.
//
// The properties TriggerType, Type are required.
type JourneyAPIInvokeTriggerNodeParam struct {
	// Any of "api-invoke".
	TriggerType JourneyAPIInvokeTriggerNodeTriggerType `json:"trigger_type,omitzero" api:"required"`
	// Any of "trigger".
	Type JourneyAPIInvokeTriggerNodeType `json:"type,omitzero" api:"required"`
	ID   param.Opt[string]               `json:"id,omitzero"`
	// Condition spec for a journey node. Accepts a single condition atom, an AND/OR
	// group, or an AND/OR nested group. Omit the `conditions` property entirely to
	// express "no conditions".
	Conditions JourneyConditionsFieldUnionParam `json:"conditions,omitzero"`
	// A JSONSchema object (Draft-07-compatible). Validated at runtime by Ajv.
	Schema map[string]any `json:"schema,omitzero"`
	paramObj
}

func (r JourneyAPIInvokeTriggerNodeParam) MarshalJSON() (data []byte, err error) {
	type shadow JourneyAPIInvokeTriggerNodeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JourneyAPIInvokeTriggerNodeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JourneyConditionAtom []string

// A leaf condition group. Exactly one of `AND` or `OR` must be present at runtime;
// each is a list of `JourneyConditionAtom` tuples.
type JourneyConditionGroup struct {
	And []JourneyConditionAtom `json:"AND"`
	Or  []JourneyConditionAtom `json:"OR"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		And         respjson.Field
		Or          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JourneyConditionGroup) RawJSON() string { return r.JSON.raw }
func (r *JourneyConditionGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this JourneyConditionGroup to a JourneyConditionGroupParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// JourneyConditionGroupParam.Overrides()
func (r JourneyConditionGroup) ToParam() JourneyConditionGroupParam {
	return param.Override[JourneyConditionGroupParam](json.RawMessage(r.RawJSON()))
}

// A leaf condition group. Exactly one of `AND` or `OR` must be present at runtime;
// each is a list of `JourneyConditionAtom` tuples.
type JourneyConditionGroupParam struct {
	And []JourneyConditionAtom `json:"AND,omitzero"`
	Or  []JourneyConditionAtom `json:"OR,omitzero"`
	paramObj
}

func (r JourneyConditionGroupParam) MarshalJSON() (data []byte, err error) {
	type shadow JourneyConditionGroupParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JourneyConditionGroupParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A nested condition group. Exactly one of `AND` or `OR` must be present at
// runtime; each is a list of `JourneyConditionGroup` items.
type JourneyConditionNestedGroup struct {
	And []JourneyConditionGroup `json:"AND"`
	Or  []JourneyConditionGroup `json:"OR"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		And         respjson.Field
		Or          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JourneyConditionNestedGroup) RawJSON() string { return r.JSON.raw }
func (r *JourneyConditionNestedGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this JourneyConditionNestedGroup to a
// JourneyConditionNestedGroupParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// JourneyConditionNestedGroupParam.Overrides()
func (r JourneyConditionNestedGroup) ToParam() JourneyConditionNestedGroupParam {
	return param.Override[JourneyConditionNestedGroupParam](json.RawMessage(r.RawJSON()))
}

// A nested condition group. Exactly one of `AND` or `OR` must be present at
// runtime; each is a list of `JourneyConditionGroup` items.
type JourneyConditionNestedGroupParam struct {
	And []JourneyConditionGroupParam `json:"AND,omitzero"`
	Or  []JourneyConditionGroupParam `json:"OR,omitzero"`
	paramObj
}

func (r JourneyConditionNestedGroupParam) MarshalJSON() (data []byte, err error) {
	type shadow JourneyConditionNestedGroupParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JourneyConditionNestedGroupParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// JourneyConditionsFieldUnion contains all possible properties and values from
// [JourneyConditionAtom], [JourneyConditionGroup], [JourneyConditionNestedGroup].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfSingleCondition]
type JourneyConditionsFieldUnion struct {
	// This field will be present if the value is a [JourneyConditionAtom] instead of
	// an object.
	OfSingleCondition JourneyConditionAtom `json:",inline"`
	// This field is a union of [[]JourneyConditionAtom], [[]JourneyConditionGroup]
	And JourneyConditionsFieldUnionAnd `json:"AND"`
	// This field is a union of [[]JourneyConditionAtom], [[]JourneyConditionGroup]
	Or   JourneyConditionsFieldUnionOr `json:"OR"`
	JSON struct {
		OfSingleCondition respjson.Field
		And               respjson.Field
		Or                respjson.Field
		raw               string
	} `json:"-"`
}

func (u JourneyConditionsFieldUnion) AsSingleCondition() (v JourneyConditionAtom) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u JourneyConditionsFieldUnion) AsConditionGroup() (v JourneyConditionGroup) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u JourneyConditionsFieldUnion) AsNestedConditionGroup() (v JourneyConditionNestedGroup) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u JourneyConditionsFieldUnion) RawJSON() string { return u.JSON.raw }

func (r *JourneyConditionsFieldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// JourneyConditionsFieldUnionAnd is an implicit subunion of
// [JourneyConditionsFieldUnion]. JourneyConditionsFieldUnionAnd provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [JourneyConditionsFieldUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfArrayOfStringArray OfJourneyConditionGroupArray]
type JourneyConditionsFieldUnionAnd struct {
	// This field will be present if the value is a [[]JourneyConditionAtom] instead of
	// an object.
	OfArrayOfStringArray []JourneyConditionAtom `json:",inline"`
	// This field will be present if the value is a [[]JourneyConditionGroup] instead
	// of an object.
	OfJourneyConditionGroupArray []JourneyConditionGroup `json:",inline"`
	JSON                         struct {
		OfArrayOfStringArray         respjson.Field
		OfJourneyConditionGroupArray respjson.Field
		raw                          string
	} `json:"-"`
}

func (r *JourneyConditionsFieldUnionAnd) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// JourneyConditionsFieldUnionOr is an implicit subunion of
// [JourneyConditionsFieldUnion]. JourneyConditionsFieldUnionOr provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [JourneyConditionsFieldUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfArrayOfStringArray OfJourneyConditionGroupArray]
type JourneyConditionsFieldUnionOr struct {
	// This field will be present if the value is a [[]JourneyConditionAtom] instead of
	// an object.
	OfArrayOfStringArray []JourneyConditionAtom `json:",inline"`
	// This field will be present if the value is a [[]JourneyConditionGroup] instead
	// of an object.
	OfJourneyConditionGroupArray []JourneyConditionGroup `json:",inline"`
	JSON                         struct {
		OfArrayOfStringArray         respjson.Field
		OfJourneyConditionGroupArray respjson.Field
		raw                          string
	} `json:"-"`
}

func (r *JourneyConditionsFieldUnionOr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this JourneyConditionsFieldUnion to a
// JourneyConditionsFieldUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// JourneyConditionsFieldUnionParam.Overrides()
func (r JourneyConditionsFieldUnion) ToParam() JourneyConditionsFieldUnionParam {
	return param.Override[JourneyConditionsFieldUnionParam](json.RawMessage(r.RawJSON()))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type JourneyConditionsFieldUnionParam struct {
	OfSingleCondition      JourneyConditionAtom              `json:",omitzero,inline"`
	OfConditionGroup       *JourneyConditionGroupParam       `json:",omitzero,inline"`
	OfNestedConditionGroup *JourneyConditionNestedGroupParam `json:",omitzero,inline"`
	paramUnion
}

func (u JourneyConditionsFieldUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSingleCondition, u.OfConditionGroup, u.OfNestedConditionGroup)
}
func (u *JourneyConditionsFieldUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *JourneyConditionsFieldUnionParam) asAny() any {
	if !param.IsOmitted(u.OfSingleCondition) {
		return &u.OfSingleCondition
	} else if !param.IsOmitted(u.OfConditionGroup) {
		return u.OfConditionGroup
	} else if !param.IsOmitted(u.OfNestedConditionGroup) {
		return u.OfNestedConditionGroup
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u JourneyConditionsFieldUnionParam) GetAnd() (res journeyConditionsFieldUnionParamAnd) {
	if vt := u.OfConditionGroup; vt != nil {
		res.any = &vt.And
	} else if vt := u.OfNestedConditionGroup; vt != nil {
		res.any = &vt.And
	}
	return
}

// Can have the runtime types [_[]JourneyConditionAtom],
// [_[]JourneyConditionGroupParam]
type journeyConditionsFieldUnionParamAnd struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]courier.JourneyConditionAtom:
//	case *[]courier.JourneyConditionGroupParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u journeyConditionsFieldUnionParamAnd) AsAny() any { return u.any }

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u JourneyConditionsFieldUnionParam) GetOr() (res journeyConditionsFieldUnionParamOr) {
	if vt := u.OfConditionGroup; vt != nil {
		res.any = &vt.Or
	} else if vt := u.OfNestedConditionGroup; vt != nil {
		res.any = &vt.Or
	}
	return
}

// Can have the runtime types [_[]JourneyConditionAtom],
// [_[]JourneyConditionGroupParam]
type journeyConditionsFieldUnionParamOr struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]courier.JourneyConditionAtom:
//	case *[]courier.JourneyConditionGroupParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u journeyConditionsFieldUnionParamOr) AsAny() any { return u.any }

// Pause the journey run for a fixed `duration`.
type JourneyDelayDurationNode struct {
	Duration string `json:"duration" api:"required"`
	// Any of "duration".
	Mode JourneyDelayDurationNodeMode `json:"mode" api:"required"`
	// Any of "delay".
	Type JourneyDelayDurationNodeType `json:"type" api:"required"`
	ID   string                       `json:"id"`
	// Condition spec for a journey node. Accepts a single condition atom, an AND/OR
	// group, or an AND/OR nested group. Omit the `conditions` property entirely to
	// express "no conditions".
	Conditions JourneyConditionsFieldUnion `json:"conditions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Duration    respjson.Field
		Mode        respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		Conditions  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JourneyDelayDurationNode) RawJSON() string { return r.JSON.raw }
func (r *JourneyDelayDurationNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this JourneyDelayDurationNode to a
// JourneyDelayDurationNodeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// JourneyDelayDurationNodeParam.Overrides()
func (r JourneyDelayDurationNode) ToParam() JourneyDelayDurationNodeParam {
	return param.Override[JourneyDelayDurationNodeParam](json.RawMessage(r.RawJSON()))
}

type JourneyDelayDurationNodeMode string

const (
	JourneyDelayDurationNodeModeDuration JourneyDelayDurationNodeMode = "duration"
)

type JourneyDelayDurationNodeType string

const (
	JourneyDelayDurationNodeTypeDelay JourneyDelayDurationNodeType = "delay"
)

// Pause the journey run for a fixed `duration`.
//
// The properties Duration, Mode, Type are required.
type JourneyDelayDurationNodeParam struct {
	Duration string `json:"duration" api:"required"`
	// Any of "duration".
	Mode JourneyDelayDurationNodeMode `json:"mode,omitzero" api:"required"`
	// Any of "delay".
	Type JourneyDelayDurationNodeType `json:"type,omitzero" api:"required"`
	ID   param.Opt[string]            `json:"id,omitzero"`
	// Condition spec for a journey node. Accepts a single condition atom, an AND/OR
	// group, or an AND/OR nested group. Omit the `conditions` property entirely to
	// express "no conditions".
	Conditions JourneyConditionsFieldUnionParam `json:"conditions,omitzero"`
	paramObj
}

func (r JourneyDelayDurationNodeParam) MarshalJSON() (data []byte, err error) {
	type shadow JourneyDelayDurationNodeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JourneyDelayDurationNodeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pause the journey run `until` a specific time.
type JourneyDelayUntilNode struct {
	// Any of "until".
	Mode JourneyDelayUntilNodeMode `json:"mode" api:"required"`
	// Any of "delay".
	Type  JourneyDelayUntilNodeType `json:"type" api:"required"`
	Until string                    `json:"until" api:"required"`
	ID    string                    `json:"id"`
	// Condition spec for a journey node. Accepts a single condition atom, an AND/OR
	// group, or an AND/OR nested group. Omit the `conditions` property entirely to
	// express "no conditions".
	Conditions JourneyConditionsFieldUnion `json:"conditions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Mode        respjson.Field
		Type        respjson.Field
		Until       respjson.Field
		ID          respjson.Field
		Conditions  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JourneyDelayUntilNode) RawJSON() string { return r.JSON.raw }
func (r *JourneyDelayUntilNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this JourneyDelayUntilNode to a JourneyDelayUntilNodeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// JourneyDelayUntilNodeParam.Overrides()
func (r JourneyDelayUntilNode) ToParam() JourneyDelayUntilNodeParam {
	return param.Override[JourneyDelayUntilNodeParam](json.RawMessage(r.RawJSON()))
}

type JourneyDelayUntilNodeMode string

const (
	JourneyDelayUntilNodeModeUntil JourneyDelayUntilNodeMode = "until"
)

type JourneyDelayUntilNodeType string

const (
	JourneyDelayUntilNodeTypeDelay JourneyDelayUntilNodeType = "delay"
)

// Pause the journey run `until` a specific time.
//
// The properties Mode, Type, Until are required.
type JourneyDelayUntilNodeParam struct {
	// Any of "until".
	Mode JourneyDelayUntilNodeMode `json:"mode,omitzero" api:"required"`
	// Any of "delay".
	Type  JourneyDelayUntilNodeType `json:"type,omitzero" api:"required"`
	Until string                    `json:"until" api:"required"`
	ID    param.Opt[string]         `json:"id,omitzero"`
	// Condition spec for a journey node. Accepts a single condition atom, an AND/OR
	// group, or an AND/OR nested group. Omit the `conditions` property entirely to
	// express "no conditions".
	Conditions JourneyConditionsFieldUnionParam `json:"conditions,omitzero"`
	paramObj
}

func (r JourneyDelayUntilNodeParam) MarshalJSON() (data []byte, err error) {
	type shadow JourneyDelayUntilNodeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JourneyDelayUntilNodeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Terminate the journey run.
type JourneyExitNode struct {
	// Any of "exit".
	Type JourneyExitNodeType `json:"type" api:"required"`
	ID   string              `json:"id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JourneyExitNode) RawJSON() string { return r.JSON.raw }
func (r *JourneyExitNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this JourneyExitNode to a JourneyExitNodeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// JourneyExitNodeParam.Overrides()
func (r JourneyExitNode) ToParam() JourneyExitNodeParam {
	return param.Override[JourneyExitNodeParam](json.RawMessage(r.RawJSON()))
}

type JourneyExitNodeType string

const (
	JourneyExitNodeTypeExit JourneyExitNodeType = "exit"
)

// Terminate the journey run.
//
// The property Type is required.
type JourneyExitNodeParam struct {
	// Any of "exit".
	Type JourneyExitNodeType `json:"type,omitzero" api:"required"`
	ID   param.Opt[string]   `json:"id,omitzero"`
	paramObj
}

func (r JourneyExitNodeParam) MarshalJSON() (data []byte, err error) {
	type shadow JourneyExitNodeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JourneyExitNodeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A/B experiment config for a send node. The recipient is deterministically
// bucketed by `bucketingKey` and routed to one of the `variants` in proportion to
// its `weight`. Present on a send node INSTEAD OF `message.template`.
type JourneyExperiment struct {
	// The value used to deterministically assign a recipient to a variant. Must be
	// non-empty with no leading or trailing whitespace.
	BucketingKey string `json:"bucketingKey" api:"required"`
	// Between 2 and 10 weighted template variants.
	Variants []JourneyExperimentVariant `json:"variants" api:"required"`
	// Unique experiment id (prefixed `exp_`). Omit to have one generated
	// automatically; when supplied it must be a valid `exp_` id.
	ID string `json:"id"`
	// Optional display name for the experiment.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BucketingKey respjson.Field
		Variants     respjson.Field
		ID           respjson.Field
		Name         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JourneyExperiment) RawJSON() string { return r.JSON.raw }
func (r *JourneyExperiment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this JourneyExperiment to a JourneyExperimentParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// JourneyExperimentParam.Overrides()
func (r JourneyExperiment) ToParam() JourneyExperimentParam {
	return param.Override[JourneyExperimentParam](json.RawMessage(r.RawJSON()))
}

// A/B experiment config for a send node. The recipient is deterministically
// bucketed by `bucketingKey` and routed to one of the `variants` in proportion to
// its `weight`. Present on a send node INSTEAD OF `message.template`.
//
// The properties BucketingKey, Variants are required.
type JourneyExperimentParam struct {
	// The value used to deterministically assign a recipient to a variant. Must be
	// non-empty with no leading or trailing whitespace.
	BucketingKey string `json:"bucketingKey" api:"required"`
	// Between 2 and 10 weighted template variants.
	Variants []JourneyExperimentVariantParam `json:"variants,omitzero" api:"required"`
	// Unique experiment id (prefixed `exp_`). Omit to have one generated
	// automatically; when supplied it must be a valid `exp_` id.
	ID param.Opt[string] `json:"id,omitzero"`
	// Optional display name for the experiment.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r JourneyExperimentParam) MarshalJSON() (data []byte, err error) {
	type shadow JourneyExperimentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JourneyExperimentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single weighted variant of an experiment. Variant ids must be unique within
// the experiment and the sum of all variant weights must be greater than 0.
// Weights are relative (no sum-to-100 requirement) — routing normalizes them
// proportionally.
type JourneyExperimentVariant struct {
	ID string `json:"id" api:"required"`
	// The notification template sent for this variant.
	TemplateID string `json:"templateId" api:"required"`
	// Relative routing weight. Must be non-negative.
	Weight float64 `json:"weight" api:"required"`
	// Optional display name for the variant.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		TemplateID  respjson.Field
		Weight      respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JourneyExperimentVariant) RawJSON() string { return r.JSON.raw }
func (r *JourneyExperimentVariant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this JourneyExperimentVariant to a
// JourneyExperimentVariantParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// JourneyExperimentVariantParam.Overrides()
func (r JourneyExperimentVariant) ToParam() JourneyExperimentVariantParam {
	return param.Override[JourneyExperimentVariantParam](json.RawMessage(r.RawJSON()))
}

// A single weighted variant of an experiment. Variant ids must be unique within
// the experiment and the sum of all variant weights must be greater than 0.
// Weights are relative (no sum-to-100 requirement) — routing normalizes them
// proportionally.
//
// The properties ID, TemplateID, Weight are required.
type JourneyExperimentVariantParam struct {
	ID string `json:"id" api:"required"`
	// The notification template sent for this variant.
	TemplateID string `json:"templateId" api:"required"`
	// Relative routing weight. Must be non-negative.
	Weight float64 `json:"weight" api:"required"`
	// Optional display name for the variant.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r JourneyExperimentVariantParam) MarshalJSON() (data []byte, err error) {
	type shadow JourneyExperimentVariantParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JourneyExperimentVariantParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Issue an HTTP GET or DELETE request and merge the response into the journey
// state per `merge_strategy`.
type JourneyFetchGetDeleteNode struct {
	// Strategy for merging a fetch response into the journey run state.
	//
	// Any of "overwrite", "soft-merge", "replace", "none".
	MergeStrategy JourneyMergeStrategy `json:"merge_strategy" api:"required"`
	// Any of "get", "delete".
	Method JourneyFetchGetDeleteNodeMethod `json:"method" api:"required"`
	// Any of "fetch".
	Type JourneyFetchGetDeleteNodeType `json:"type" api:"required"`
	URL  string                        `json:"url" api:"required"`
	ID   string                        `json:"id"`
	// Condition spec for a journey node. Accepts a single condition atom, an AND/OR
	// group, or an AND/OR nested group. Omit the `conditions` property entirely to
	// express "no conditions".
	Conditions  JourneyConditionsFieldUnion `json:"conditions"`
	Headers     map[string]string           `json:"headers"`
	QueryParams map[string]string           `json:"query_params"`
	// A JSONSchema object (Draft-07-compatible). Validated at runtime by Ajv.
	ResponseSchema map[string]any `json:"response_schema"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MergeStrategy  respjson.Field
		Method         respjson.Field
		Type           respjson.Field
		URL            respjson.Field
		ID             respjson.Field
		Conditions     respjson.Field
		Headers        respjson.Field
		QueryParams    respjson.Field
		ResponseSchema respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JourneyFetchGetDeleteNode) RawJSON() string { return r.JSON.raw }
func (r *JourneyFetchGetDeleteNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this JourneyFetchGetDeleteNode to a
// JourneyFetchGetDeleteNodeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// JourneyFetchGetDeleteNodeParam.Overrides()
func (r JourneyFetchGetDeleteNode) ToParam() JourneyFetchGetDeleteNodeParam {
	return param.Override[JourneyFetchGetDeleteNodeParam](json.RawMessage(r.RawJSON()))
}

type JourneyFetchGetDeleteNodeMethod string

const (
	JourneyFetchGetDeleteNodeMethodGet    JourneyFetchGetDeleteNodeMethod = "get"
	JourneyFetchGetDeleteNodeMethodDelete JourneyFetchGetDeleteNodeMethod = "delete"
)

type JourneyFetchGetDeleteNodeType string

const (
	JourneyFetchGetDeleteNodeTypeFetch JourneyFetchGetDeleteNodeType = "fetch"
)

// Issue an HTTP GET or DELETE request and merge the response into the journey
// state per `merge_strategy`.
//
// The properties MergeStrategy, Method, Type, URL are required.
type JourneyFetchGetDeleteNodeParam struct {
	// Strategy for merging a fetch response into the journey run state.
	//
	// Any of "overwrite", "soft-merge", "replace", "none".
	MergeStrategy JourneyMergeStrategy `json:"merge_strategy,omitzero" api:"required"`
	// Any of "get", "delete".
	Method JourneyFetchGetDeleteNodeMethod `json:"method,omitzero" api:"required"`
	// Any of "fetch".
	Type JourneyFetchGetDeleteNodeType `json:"type,omitzero" api:"required"`
	URL  string                        `json:"url" api:"required"`
	ID   param.Opt[string]             `json:"id,omitzero"`
	// Condition spec for a journey node. Accepts a single condition atom, an AND/OR
	// group, or an AND/OR nested group. Omit the `conditions` property entirely to
	// express "no conditions".
	Conditions  JourneyConditionsFieldUnionParam `json:"conditions,omitzero"`
	Headers     map[string]string                `json:"headers,omitzero"`
	QueryParams map[string]string                `json:"query_params,omitzero"`
	// A JSONSchema object (Draft-07-compatible). Validated at runtime by Ajv.
	ResponseSchema map[string]any `json:"response_schema,omitzero"`
	paramObj
}

func (r JourneyFetchGetDeleteNodeParam) MarshalJSON() (data []byte, err error) {
	type shadow JourneyFetchGetDeleteNodeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JourneyFetchGetDeleteNodeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Issue an HTTP POST or PUT request with a `body` and merge the response into the
// journey state per `merge_strategy`.
type JourneyFetchPostPutNode struct {
	// Strategy for merging a fetch response into the journey run state.
	//
	// Any of "overwrite", "soft-merge", "replace", "none".
	MergeStrategy JourneyMergeStrategy `json:"merge_strategy" api:"required"`
	// Any of "post", "put".
	Method JourneyFetchPostPutNodeMethod `json:"method" api:"required"`
	// Any of "fetch".
	Type JourneyFetchPostPutNodeType `json:"type" api:"required"`
	URL  string                      `json:"url" api:"required"`
	ID   string                      `json:"id"`
	Body string                      `json:"body"`
	// Condition spec for a journey node. Accepts a single condition atom, an AND/OR
	// group, or an AND/OR nested group. Omit the `conditions` property entirely to
	// express "no conditions".
	Conditions  JourneyConditionsFieldUnion `json:"conditions"`
	Headers     map[string]string           `json:"headers"`
	QueryParams map[string]string           `json:"query_params"`
	// A JSONSchema object (Draft-07-compatible). Validated at runtime by Ajv.
	ResponseSchema map[string]any `json:"response_schema"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MergeStrategy  respjson.Field
		Method         respjson.Field
		Type           respjson.Field
		URL            respjson.Field
		ID             respjson.Field
		Body           respjson.Field
		Conditions     respjson.Field
		Headers        respjson.Field
		QueryParams    respjson.Field
		ResponseSchema respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JourneyFetchPostPutNode) RawJSON() string { return r.JSON.raw }
func (r *JourneyFetchPostPutNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this JourneyFetchPostPutNode to a JourneyFetchPostPutNodeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// JourneyFetchPostPutNodeParam.Overrides()
func (r JourneyFetchPostPutNode) ToParam() JourneyFetchPostPutNodeParam {
	return param.Override[JourneyFetchPostPutNodeParam](json.RawMessage(r.RawJSON()))
}

type JourneyFetchPostPutNodeMethod string

const (
	JourneyFetchPostPutNodeMethodPost JourneyFetchPostPutNodeMethod = "post"
	JourneyFetchPostPutNodeMethodPut  JourneyFetchPostPutNodeMethod = "put"
)

type JourneyFetchPostPutNodeType string

const (
	JourneyFetchPostPutNodeTypeFetch JourneyFetchPostPutNodeType = "fetch"
)

// Issue an HTTP POST or PUT request with a `body` and merge the response into the
// journey state per `merge_strategy`.
//
// The properties MergeStrategy, Method, Type, URL are required.
type JourneyFetchPostPutNodeParam struct {
	// Strategy for merging a fetch response into the journey run state.
	//
	// Any of "overwrite", "soft-merge", "replace", "none".
	MergeStrategy JourneyMergeStrategy `json:"merge_strategy,omitzero" api:"required"`
	// Any of "post", "put".
	Method JourneyFetchPostPutNodeMethod `json:"method,omitzero" api:"required"`
	// Any of "fetch".
	Type JourneyFetchPostPutNodeType `json:"type,omitzero" api:"required"`
	URL  string                      `json:"url" api:"required"`
	ID   param.Opt[string]           `json:"id,omitzero"`
	Body param.Opt[string]           `json:"body,omitzero"`
	// Condition spec for a journey node. Accepts a single condition atom, an AND/OR
	// group, or an AND/OR nested group. Omit the `conditions` property entirely to
	// express "no conditions".
	Conditions  JourneyConditionsFieldUnionParam `json:"conditions,omitzero"`
	Headers     map[string]string                `json:"headers,omitzero"`
	QueryParams map[string]string                `json:"query_params,omitzero"`
	// A JSONSchema object (Draft-07-compatible). Validated at runtime by Ajv.
	ResponseSchema map[string]any `json:"response_schema,omitzero"`
	paramObj
}

func (r JourneyFetchPostPutNodeParam) MarshalJSON() (data []byte, err error) {
	type shadow JourneyFetchPostPutNodeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JourneyFetchPostPutNodeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Strategy for merging a fetch response into the journey run state.
type JourneyMergeStrategy string

const (
	JourneyMergeStrategyOverwrite JourneyMergeStrategy = "overwrite"
	JourneyMergeStrategySoftMerge JourneyMergeStrategy = "soft-merge"
	JourneyMergeStrategyReplace   JourneyMergeStrategy = "replace"
	JourneyMergeStrategyNone      JourneyMergeStrategy = "none"
)

// JourneyNodeUnion contains all possible properties and values from
// [JourneyAPIInvokeTriggerNode], [JourneySegmentTriggerNode], [JourneySendNode],
// [JourneyDelayDurationNode], [JourneyDelayUntilNode],
// [JourneyFetchGetDeleteNode], [JourneyFetchPostPutNode], [JourneyAINode],
// [JourneyThrottleStaticNode], [JourneyThrottleDynamicNode], [JourneyNodeBatch],
// [JourneyNodeAddToDigest], [JourneyExitNode], [JourneyNodeJourneyBranchNode].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type JourneyNodeUnion struct {
	TriggerType string `json:"trigger_type"`
	Type        string `json:"type"`
	ID          string `json:"id"`
	// This field is from variant [JourneyAPIInvokeTriggerNode].
	Conditions JourneyConditionsFieldUnion `json:"conditions"`
	// This field is from variant [JourneyAPIInvokeTriggerNode].
	Schema map[string]any `json:"schema"`
	// This field is from variant [JourneySegmentTriggerNode].
	RequestType JourneySegmentTriggerNodeRequestType `json:"request_type"`
	// This field is from variant [JourneySegmentTriggerNode].
	EventID string `json:"event_id"`
	// This field is from variant [JourneySendNode].
	Message JourneySendNodeMessage `json:"message"`
	// This field is from variant [JourneySendNode].
	Experiment JourneyExperiment `json:"experiment"`
	// This field is from variant [JourneyDelayDurationNode].
	Duration string `json:"duration"`
	Mode     string `json:"mode"`
	// This field is from variant [JourneyDelayUntilNode].
	Until string `json:"until"`
	// This field is from variant [JourneyFetchGetDeleteNode].
	MergeStrategy  JourneyMergeStrategy `json:"merge_strategy"`
	Method         string               `json:"method"`
	URL            string               `json:"url"`
	Headers        string               `json:"headers"`
	QueryParams    string               `json:"query_params"`
	ResponseSchema any                  `json:"response_schema"`
	// This field is from variant [JourneyFetchPostPutNode].
	Body string `json:"body"`
	// This field is from variant [JourneyAINode].
	OutputSchema map[string]any `json:"output_schema"`
	// This field is from variant [JourneyAINode].
	Model string `json:"model"`
	// This field is from variant [JourneyAINode].
	UserPrompt string `json:"user_prompt"`
	// This field is from variant [JourneyAINode].
	WebSearch  bool   `json:"web_search"`
	MaxAllowed int64  `json:"max_allowed"`
	Period     string `json:"period"`
	Scope      string `json:"scope"`
	// This field is from variant [JourneyThrottleDynamicNode].
	ThrottleKey string `json:"throttle_key"`
	// This field is from variant [JourneyNodeBatch].
	MaxWaitPeriod string `json:"max_wait_period"`
	// This field is from variant [JourneyNodeBatch].
	Retain JourneyNodeBatchRetain `json:"retain"`
	// This field is from variant [JourneyNodeBatch].
	WaitPeriod string `json:"wait_period"`
	// This field is from variant [JourneyNodeBatch].
	CategoryKey string `json:"category_key"`
	// This field is from variant [JourneyNodeBatch].
	MaxItems int64 `json:"max_items"`
	// This field is from variant [JourneyNodeAddToDigest].
	SubscriptionTopicID string `json:"subscription_topic_id"`
	// This field is from variant [JourneyNodeJourneyBranchNode].
	Default JourneyNodeJourneyBranchNodeDefault `json:"default"`
	// This field is from variant [JourneyNodeJourneyBranchNode].
	Paths []JourneyNodeJourneyBranchNodePath `json:"paths"`
	JSON  struct {
		TriggerType         respjson.Field
		Type                respjson.Field
		ID                  respjson.Field
		Conditions          respjson.Field
		Schema              respjson.Field
		RequestType         respjson.Field
		EventID             respjson.Field
		Message             respjson.Field
		Experiment          respjson.Field
		Duration            respjson.Field
		Mode                respjson.Field
		Until               respjson.Field
		MergeStrategy       respjson.Field
		Method              respjson.Field
		URL                 respjson.Field
		Headers             respjson.Field
		QueryParams         respjson.Field
		ResponseSchema      respjson.Field
		Body                respjson.Field
		OutputSchema        respjson.Field
		Model               respjson.Field
		UserPrompt          respjson.Field
		WebSearch           respjson.Field
		MaxAllowed          respjson.Field
		Period              respjson.Field
		Scope               respjson.Field
		ThrottleKey         respjson.Field
		MaxWaitPeriod       respjson.Field
		Retain              respjson.Field
		WaitPeriod          respjson.Field
		CategoryKey         respjson.Field
		MaxItems            respjson.Field
		SubscriptionTopicID respjson.Field
		Default             respjson.Field
		Paths               respjson.Field
		raw                 string
	} `json:"-"`
}

func (u JourneyNodeUnion) AsAPIInvokeTrigger() (v JourneyAPIInvokeTriggerNode) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u JourneyNodeUnion) AsSegmentTrigger() (v JourneySegmentTriggerNode) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u JourneyNodeUnion) AsSend() (v JourneySendNode) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u JourneyNodeUnion) AsDelayForDuration() (v JourneyDelayDurationNode) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u JourneyNodeUnion) AsDelayUntil() (v JourneyDelayUntilNode) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u JourneyNodeUnion) AsFetchGetDelete() (v JourneyFetchGetDeleteNode) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u JourneyNodeUnion) AsFetchPostPut() (v JourneyFetchPostPutNode) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u JourneyNodeUnion) AsAI() (v JourneyAINode) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u JourneyNodeUnion) AsThrottleStatic() (v JourneyThrottleStaticNode) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u JourneyNodeUnion) AsThrottleDynamic() (v JourneyThrottleDynamicNode) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u JourneyNodeUnion) AsBatch() (v JourneyNodeBatch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u JourneyNodeUnion) AsAddToDigest() (v JourneyNodeAddToDigest) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u JourneyNodeUnion) AsExit() (v JourneyExitNode) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u JourneyNodeUnion) AsJourneyBranchNode() (v JourneyNodeJourneyBranchNode) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u JourneyNodeUnion) RawJSON() string { return u.JSON.raw }

func (r *JourneyNodeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this JourneyNodeUnion to a JourneyNodeUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// JourneyNodeUnionParam.Overrides()
func (r JourneyNodeUnion) ToParam() JourneyNodeUnionParam {
	return param.Override[JourneyNodeUnionParam](json.RawMessage(r.RawJSON()))
}

// Collect events arriving at the node into a single batch and fire one downstream
// step with the aggregated payload. The first event into a batch owns the run;
// later contributing events terminate at the batch step. The batch releases when
// any of `max_items` is reached, a quiet window of `wait_period` elapses, or the
// `max_wait_period` ceiling hits.
type JourneyNodeBatch struct {
	// ISO 8601 duration. Hard ceiling from the first event into the batch; releases
	// the batch unconditionally when it elapses.
	MaxWaitPeriod string `json:"max_wait_period" api:"required"`
	// How to select which collected events to retain in the aggregated payload when
	// the batch releases.
	Retain JourneyNodeBatchRetain `json:"retain" api:"required"`
	// Any of "user".
	Scope string `json:"scope" api:"required"`
	// Any of "batch".
	Type string `json:"type" api:"required"`
	// ISO 8601 duration. Quiet window that releases the batch when it elapses with no
	// new contributing events. Must be less than `max_wait_period`.
	WaitPeriod string `json:"wait_period" api:"required"`
	ID         string `json:"id"`
	// Optional partition key. Events with the same `category_key` are batched
	// together; events with different values are batched separately.
	CategoryKey string `json:"category_key"`
	// Condition spec for a journey node. Accepts a single condition atom, an AND/OR
	// group, or an AND/OR nested group. Omit the `conditions` property entirely to
	// express "no conditions".
	Conditions JourneyConditionsFieldUnion `json:"conditions"`
	// Releases the batch once this many events have been collected.
	MaxItems int64 `json:"max_items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxWaitPeriod respjson.Field
		Retain        respjson.Field
		Scope         respjson.Field
		Type          respjson.Field
		WaitPeriod    respjson.Field
		ID            respjson.Field
		CategoryKey   respjson.Field
		Conditions    respjson.Field
		MaxItems      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JourneyNodeBatch) RawJSON() string { return r.JSON.raw }
func (r *JourneyNodeBatch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How to select which collected events to retain in the aggregated payload when
// the batch releases.
type JourneyNodeBatchRetain struct {
	Count int64 `json:"count" api:"required"`
	// Any of "first", "last", "highest", "lowest".
	Type string `json:"type" api:"required"`
	// Dot-path into the event payload (e.g. `data.priority`). Required when `type` is
	// `highest` or `lowest`.
	SortKey string `json:"sort_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Type        respjson.Field
		SortKey     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JourneyNodeBatchRetain) RawJSON() string { return r.JSON.raw }
func (r *JourneyNodeBatchRetain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Add the current event to a digest keyed by the given subscription topic. The
// digest accumulates events and releases them on the schedule configured for the
// topic.
type JourneyNodeAddToDigest struct {
	// The subscription topic that owns the digest the event is added to.
	SubscriptionTopicID string `json:"subscription_topic_id" api:"required"`
	// Any of "add-to-digest".
	Type string `json:"type" api:"required"`
	ID   string `json:"id"`
	// Condition spec for a journey node. Accepts a single condition atom, an AND/OR
	// group, or an AND/OR nested group. Omit the `conditions` property entirely to
	// express "no conditions".
	Conditions JourneyConditionsFieldUnion `json:"conditions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SubscriptionTopicID respjson.Field
		Type                respjson.Field
		ID                  respjson.Field
		Conditions          respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JourneyNodeAddToDigest) RawJSON() string { return r.JSON.raw }
func (r *JourneyNodeAddToDigest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Branch node. Routes to the first entry in `paths[]` whose `conditions` match,
// else falls through to `default.nodes`.
type JourneyNodeJourneyBranchNode struct {
	Default JourneyNodeJourneyBranchNodeDefault `json:"default" api:"required"`
	Paths   []JourneyNodeJourneyBranchNodePath  `json:"paths" api:"required"`
	// Any of "branch".
	Type string `json:"type" api:"required"`
	ID   string `json:"id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Default     respjson.Field
		Paths       respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JourneyNodeJourneyBranchNode) RawJSON() string { return r.JSON.raw }
func (r *JourneyNodeJourneyBranchNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JourneyNodeJourneyBranchNodeDefault struct {
	Nodes []JourneyNodeUnion `json:"nodes" api:"required"`
	Label string             `json:"label"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Nodes       respjson.Field
		Label       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JourneyNodeJourneyBranchNodeDefault) RawJSON() string { return r.JSON.raw }
func (r *JourneyNodeJourneyBranchNodeDefault) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JourneyNodeJourneyBranchNodePath struct {
	// Condition spec for a journey node. Accepts a single condition atom, an AND/OR
	// group, or an AND/OR nested group. Omit the `conditions` property entirely to
	// express "no conditions".
	Conditions JourneyConditionsFieldUnion `json:"conditions" api:"required"`
	Nodes      []JourneyNodeUnion          `json:"nodes" api:"required"`
	Label      string                      `json:"label"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Conditions  respjson.Field
		Nodes       respjson.Field
		Label       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JourneyNodeJourneyBranchNodePath) RawJSON() string { return r.JSON.raw }
func (r *JourneyNodeJourneyBranchNodePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func JourneyNodeParamOfAPIInvokeTrigger(triggerType JourneyAPIInvokeTriggerNodeTriggerType, type_ JourneyAPIInvokeTriggerNodeType) JourneyNodeUnionParam {
	var variant JourneyAPIInvokeTriggerNodeParam
	variant.TriggerType = triggerType
	variant.Type = type_
	return JourneyNodeUnionParam{OfAPIInvokeTrigger: &variant}
}

func JourneyNodeParamOfSegmentTrigger(requestType JourneySegmentTriggerNodeRequestType, triggerType JourneySegmentTriggerNodeTriggerType, type_ JourneySegmentTriggerNodeType) JourneyNodeUnionParam {
	var variant JourneySegmentTriggerNodeParam
	variant.RequestType = requestType
	variant.TriggerType = triggerType
	variant.Type = type_
	return JourneyNodeUnionParam{OfSegmentTrigger: &variant}
}

func JourneyNodeParamOfSend(message JourneySendNodeMessageParam, type_ JourneySendNodeType) JourneyNodeUnionParam {
	var variant JourneySendNodeParam
	variant.Message = message
	variant.Type = type_
	return JourneyNodeUnionParam{OfSend: &variant}
}

func JourneyNodeParamOfDelayForDuration(duration string, mode JourneyDelayDurationNodeMode, type_ JourneyDelayDurationNodeType) JourneyNodeUnionParam {
	var variant JourneyDelayDurationNodeParam
	variant.Duration = duration
	variant.Mode = mode
	variant.Type = type_
	return JourneyNodeUnionParam{OfDelayForDuration: &variant}
}

func JourneyNodeParamOfDelayUntil(mode JourneyDelayUntilNodeMode, type_ JourneyDelayUntilNodeType, until string) JourneyNodeUnionParam {
	var variant JourneyDelayUntilNodeParam
	variant.Mode = mode
	variant.Type = type_
	variant.Until = until
	return JourneyNodeUnionParam{OfDelayUntil: &variant}
}

func JourneyNodeParamOfAI(outputSchema map[string]any, type_ JourneyAINodeType) JourneyNodeUnionParam {
	var variant JourneyAINodeParam
	variant.OutputSchema = outputSchema
	variant.Type = type_
	return JourneyNodeUnionParam{OfAI: &variant}
}

func JourneyNodeParamOfAddToDigest(subscriptionTopicID string, type_ string) JourneyNodeUnionParam {
	var variant JourneyNodeAddToDigestParam
	variant.SubscriptionTopicID = subscriptionTopicID
	variant.Type = type_
	return JourneyNodeUnionParam{OfAddToDigest: &variant}
}

func JourneyNodeParamOfExit(type_ JourneyExitNodeType) JourneyNodeUnionParam {
	var variant JourneyExitNodeParam
	variant.Type = type_
	return JourneyNodeUnionParam{OfExit: &variant}
}

func JourneyNodeParamOfJourneyBranchNode(default_ JourneyNodeJourneyBranchNodeDefaultParam, paths []JourneyNodeJourneyBranchNodePathParam, type_ string) JourneyNodeUnionParam {
	var variant JourneyNodeJourneyBranchNodeParam
	variant.Default = default_
	variant.Paths = paths
	variant.Type = type_
	return JourneyNodeUnionParam{OfJourneyBranchNode: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type JourneyNodeUnionParam struct {
	OfAPIInvokeTrigger  *JourneyAPIInvokeTriggerNodeParam  `json:",omitzero,inline"`
	OfSegmentTrigger    *JourneySegmentTriggerNodeParam    `json:",omitzero,inline"`
	OfSend              *JourneySendNodeParam              `json:",omitzero,inline"`
	OfDelayForDuration  *JourneyDelayDurationNodeParam     `json:",omitzero,inline"`
	OfDelayUntil        *JourneyDelayUntilNodeParam        `json:",omitzero,inline"`
	OfFetchGetDelete    *JourneyFetchGetDeleteNodeParam    `json:",omitzero,inline"`
	OfFetchPostPut      *JourneyFetchPostPutNodeParam      `json:",omitzero,inline"`
	OfAI                *JourneyAINodeParam                `json:",omitzero,inline"`
	OfThrottleStatic    *JourneyThrottleStaticNodeParam    `json:",omitzero,inline"`
	OfThrottleDynamic   *JourneyThrottleDynamicNodeParam   `json:",omitzero,inline"`
	OfBatch             *JourneyNodeBatchParam             `json:",omitzero,inline"`
	OfAddToDigest       *JourneyNodeAddToDigestParam       `json:",omitzero,inline"`
	OfExit              *JourneyExitNodeParam              `json:",omitzero,inline"`
	OfJourneyBranchNode *JourneyNodeJourneyBranchNodeParam `json:",omitzero,inline"`
	paramUnion
}

func (u JourneyNodeUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAPIInvokeTrigger,
		u.OfSegmentTrigger,
		u.OfSend,
		u.OfDelayForDuration,
		u.OfDelayUntil,
		u.OfFetchGetDelete,
		u.OfFetchPostPut,
		u.OfAI,
		u.OfThrottleStatic,
		u.OfThrottleDynamic,
		u.OfBatch,
		u.OfAddToDigest,
		u.OfExit,
		u.OfJourneyBranchNode)
}
func (u *JourneyNodeUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *JourneyNodeUnionParam) asAny() any {
	if !param.IsOmitted(u.OfAPIInvokeTrigger) {
		return u.OfAPIInvokeTrigger
	} else if !param.IsOmitted(u.OfSegmentTrigger) {
		return u.OfSegmentTrigger
	} else if !param.IsOmitted(u.OfSend) {
		return u.OfSend
	} else if !param.IsOmitted(u.OfDelayForDuration) {
		return u.OfDelayForDuration
	} else if !param.IsOmitted(u.OfDelayUntil) {
		return u.OfDelayUntil
	} else if !param.IsOmitted(u.OfFetchGetDelete) {
		return u.OfFetchGetDelete
	} else if !param.IsOmitted(u.OfFetchPostPut) {
		return u.OfFetchPostPut
	} else if !param.IsOmitted(u.OfAI) {
		return u.OfAI
	} else if !param.IsOmitted(u.OfThrottleStatic) {
		return u.OfThrottleStatic
	} else if !param.IsOmitted(u.OfThrottleDynamic) {
		return u.OfThrottleDynamic
	} else if !param.IsOmitted(u.OfBatch) {
		return u.OfBatch
	} else if !param.IsOmitted(u.OfAddToDigest) {
		return u.OfAddToDigest
	} else if !param.IsOmitted(u.OfExit) {
		return u.OfExit
	} else if !param.IsOmitted(u.OfJourneyBranchNode) {
		return u.OfJourneyBranchNode
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JourneyNodeUnionParam) GetSchema() map[string]any {
	if vt := u.OfAPIInvokeTrigger; vt != nil {
		return vt.Schema
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JourneyNodeUnionParam) GetRequestType() *string {
	if vt := u.OfSegmentTrigger; vt != nil {
		return (*string)(&vt.RequestType)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JourneyNodeUnionParam) GetEventID() *string {
	if vt := u.OfSegmentTrigger; vt != nil && vt.EventID.Valid() {
		return &vt.EventID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JourneyNodeUnionParam) GetMessage() *JourneySendNodeMessageParam {
	if vt := u.OfSend; vt != nil {
		return &vt.Message
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JourneyNodeUnionParam) GetExperiment() *JourneyExperimentParam {
	if vt := u.OfSend; vt != nil {
		return &vt.Experiment
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JourneyNodeUnionParam) GetDuration() *string {
	if vt := u.OfDelayForDuration; vt != nil {
		return &vt.Duration
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JourneyNodeUnionParam) GetUntil() *string {
	if vt := u.OfDelayUntil; vt != nil {
		return &vt.Until
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JourneyNodeUnionParam) GetBody() *string {
	if vt := u.OfFetchPostPut; vt != nil && vt.Body.Valid() {
		return &vt.Body.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JourneyNodeUnionParam) GetOutputSchema() map[string]any {
	if vt := u.OfAI; vt != nil {
		return vt.OutputSchema
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JourneyNodeUnionParam) GetModel() *string {
	if vt := u.OfAI; vt != nil && vt.Model.Valid() {
		return &vt.Model.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JourneyNodeUnionParam) GetUserPrompt() *string {
	if vt := u.OfAI; vt != nil && vt.UserPrompt.Valid() {
		return &vt.UserPrompt.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JourneyNodeUnionParam) GetWebSearch() *bool {
	if vt := u.OfAI; vt != nil && vt.WebSearch.Valid() {
		return &vt.WebSearch.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JourneyNodeUnionParam) GetThrottleKey() *string {
	if vt := u.OfThrottleDynamic; vt != nil {
		return &vt.ThrottleKey
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JourneyNodeUnionParam) GetMaxWaitPeriod() *string {
	if vt := u.OfBatch; vt != nil {
		return &vt.MaxWaitPeriod
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JourneyNodeUnionParam) GetRetain() *JourneyNodeBatchRetainParam {
	if vt := u.OfBatch; vt != nil {
		return &vt.Retain
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JourneyNodeUnionParam) GetWaitPeriod() *string {
	if vt := u.OfBatch; vt != nil {
		return &vt.WaitPeriod
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JourneyNodeUnionParam) GetCategoryKey() *string {
	if vt := u.OfBatch; vt != nil && vt.CategoryKey.Valid() {
		return &vt.CategoryKey.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JourneyNodeUnionParam) GetMaxItems() *int64 {
	if vt := u.OfBatch; vt != nil && vt.MaxItems.Valid() {
		return &vt.MaxItems.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JourneyNodeUnionParam) GetSubscriptionTopicID() *string {
	if vt := u.OfAddToDigest; vt != nil {
		return &vt.SubscriptionTopicID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JourneyNodeUnionParam) GetDefault() *JourneyNodeJourneyBranchNodeDefaultParam {
	if vt := u.OfJourneyBranchNode; vt != nil {
		return &vt.Default
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JourneyNodeUnionParam) GetPaths() []JourneyNodeJourneyBranchNodePathParam {
	if vt := u.OfJourneyBranchNode; vt != nil {
		return vt.Paths
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JourneyNodeUnionParam) GetTriggerType() *string {
	if vt := u.OfAPIInvokeTrigger; vt != nil {
		return (*string)(&vt.TriggerType)
	} else if vt := u.OfSegmentTrigger; vt != nil {
		return (*string)(&vt.TriggerType)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JourneyNodeUnionParam) GetType() *string {
	if vt := u.OfAPIInvokeTrigger; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfSegmentTrigger; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfSend; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfDelayForDuration; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfDelayUntil; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfFetchGetDelete; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfFetchPostPut; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAI; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfThrottleStatic; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfThrottleDynamic; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfBatch; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAddToDigest; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfExit; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfJourneyBranchNode; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JourneyNodeUnionParam) GetID() *string {
	if vt := u.OfAPIInvokeTrigger; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfSegmentTrigger; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfSend; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfDelayForDuration; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfDelayUntil; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfFetchGetDelete; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfFetchPostPut; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfAI; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfThrottleStatic; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfThrottleDynamic; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfBatch; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfAddToDigest; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfExit; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfJourneyBranchNode; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JourneyNodeUnionParam) GetMode() *string {
	if vt := u.OfDelayForDuration; vt != nil {
		return (*string)(&vt.Mode)
	} else if vt := u.OfDelayUntil; vt != nil {
		return (*string)(&vt.Mode)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JourneyNodeUnionParam) GetMergeStrategy() *string {
	if vt := u.OfFetchGetDelete; vt != nil {
		return (*string)(&vt.MergeStrategy)
	} else if vt := u.OfFetchPostPut; vt != nil {
		return (*string)(&vt.MergeStrategy)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JourneyNodeUnionParam) GetMethod() *string {
	if vt := u.OfFetchGetDelete; vt != nil {
		return (*string)(&vt.Method)
	} else if vt := u.OfFetchPostPut; vt != nil {
		return (*string)(&vt.Method)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JourneyNodeUnionParam) GetURL() *string {
	if vt := u.OfFetchGetDelete; vt != nil {
		return (*string)(&vt.URL)
	} else if vt := u.OfFetchPostPut; vt != nil {
		return (*string)(&vt.URL)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JourneyNodeUnionParam) GetMaxAllowed() *int64 {
	if vt := u.OfThrottleStatic; vt != nil {
		return (*int64)(&vt.MaxAllowed)
	} else if vt := u.OfThrottleDynamic; vt != nil {
		return (*int64)(&vt.MaxAllowed)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JourneyNodeUnionParam) GetPeriod() *string {
	if vt := u.OfThrottleStatic; vt != nil {
		return (*string)(&vt.Period)
	} else if vt := u.OfThrottleDynamic; vt != nil {
		return (*string)(&vt.Period)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u JourneyNodeUnionParam) GetScope() *string {
	if vt := u.OfThrottleStatic; vt != nil {
		return (*string)(&vt.Scope)
	} else if vt := u.OfThrottleDynamic; vt != nil {
		return (*string)(&vt.Scope)
	} else if vt := u.OfBatch; vt != nil {
		return (*string)(&vt.Scope)
	}
	return nil
}

// Returns a pointer to the underlying variant's Conditions property, if present.
func (u JourneyNodeUnionParam) GetConditions() *JourneyConditionsFieldUnionParam {
	if vt := u.OfAPIInvokeTrigger; vt != nil {
		return &vt.Conditions
	} else if vt := u.OfSegmentTrigger; vt != nil {
		return &vt.Conditions
	} else if vt := u.OfSend; vt != nil {
		return &vt.Conditions
	} else if vt := u.OfDelayForDuration; vt != nil {
		return &vt.Conditions
	} else if vt := u.OfDelayUntil; vt != nil {
		return &vt.Conditions
	} else if vt := u.OfFetchGetDelete; vt != nil {
		return &vt.Conditions
	} else if vt := u.OfFetchPostPut; vt != nil {
		return &vt.Conditions
	} else if vt := u.OfAI; vt != nil {
		return &vt.Conditions
	} else if vt := u.OfThrottleStatic; vt != nil {
		return &vt.Conditions
	} else if vt := u.OfThrottleDynamic; vt != nil {
		return &vt.Conditions
	} else if vt := u.OfBatch; vt != nil {
		return &vt.Conditions
	} else if vt := u.OfAddToDigest; vt != nil {
		return &vt.Conditions
	}
	return nil
}

// Returns a pointer to the underlying variant's Headers property, if present.
func (u JourneyNodeUnionParam) GetHeaders() map[string]string {
	if vt := u.OfFetchGetDelete; vt != nil {
		return vt.Headers
	} else if vt := u.OfFetchPostPut; vt != nil {
		return vt.Headers
	}
	return nil
}

// Returns a pointer to the underlying variant's QueryParams property, if present.
func (u JourneyNodeUnionParam) GetQueryParams() map[string]string {
	if vt := u.OfFetchGetDelete; vt != nil {
		return vt.QueryParams
	} else if vt := u.OfFetchPostPut; vt != nil {
		return vt.QueryParams
	}
	return nil
}

// Returns a pointer to the underlying variant's ResponseSchema property, if
// present.
func (u JourneyNodeUnionParam) GetResponseSchema() map[string]any {
	if vt := u.OfFetchGetDelete; vt != nil {
		return vt.ResponseSchema
	} else if vt := u.OfFetchPostPut; vt != nil {
		return vt.ResponseSchema
	}
	return nil
}

// Collect events arriving at the node into a single batch and fire one downstream
// step with the aggregated payload. The first event into a batch owns the run;
// later contributing events terminate at the batch step. The batch releases when
// any of `max_items` is reached, a quiet window of `wait_period` elapses, or the
// `max_wait_period` ceiling hits.
//
// The properties MaxWaitPeriod, Retain, Scope, Type, WaitPeriod are required.
type JourneyNodeBatchParam struct {
	// ISO 8601 duration. Hard ceiling from the first event into the batch; releases
	// the batch unconditionally when it elapses.
	MaxWaitPeriod string `json:"max_wait_period" api:"required"`
	// How to select which collected events to retain in the aggregated payload when
	// the batch releases.
	Retain JourneyNodeBatchRetainParam `json:"retain,omitzero" api:"required"`
	// Any of "user".
	Scope string `json:"scope,omitzero" api:"required"`
	// Any of "batch".
	Type string `json:"type,omitzero" api:"required"`
	// ISO 8601 duration. Quiet window that releases the batch when it elapses with no
	// new contributing events. Must be less than `max_wait_period`.
	WaitPeriod string            `json:"wait_period" api:"required"`
	ID         param.Opt[string] `json:"id,omitzero"`
	// Optional partition key. Events with the same `category_key` are batched
	// together; events with different values are batched separately.
	CategoryKey param.Opt[string] `json:"category_key,omitzero"`
	// Releases the batch once this many events have been collected.
	MaxItems param.Opt[int64] `json:"max_items,omitzero"`
	// Condition spec for a journey node. Accepts a single condition atom, an AND/OR
	// group, or an AND/OR nested group. Omit the `conditions` property entirely to
	// express "no conditions".
	Conditions JourneyConditionsFieldUnionParam `json:"conditions,omitzero"`
	paramObj
}

func (r JourneyNodeBatchParam) MarshalJSON() (data []byte, err error) {
	type shadow JourneyNodeBatchParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JourneyNodeBatchParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[JourneyNodeBatchParam](
		"scope", "user",
	)
	apijson.RegisterFieldValidator[JourneyNodeBatchParam](
		"type", "batch",
	)
}

// How to select which collected events to retain in the aggregated payload when
// the batch releases.
//
// The properties Count, Type are required.
type JourneyNodeBatchRetainParam struct {
	Count int64 `json:"count" api:"required"`
	// Any of "first", "last", "highest", "lowest".
	Type string `json:"type,omitzero" api:"required"`
	// Dot-path into the event payload (e.g. `data.priority`). Required when `type` is
	// `highest` or `lowest`.
	SortKey param.Opt[string] `json:"sort_key,omitzero"`
	paramObj
}

func (r JourneyNodeBatchRetainParam) MarshalJSON() (data []byte, err error) {
	type shadow JourneyNodeBatchRetainParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JourneyNodeBatchRetainParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[JourneyNodeBatchRetainParam](
		"type", "first", "last", "highest", "lowest",
	)
}

// Add the current event to a digest keyed by the given subscription topic. The
// digest accumulates events and releases them on the schedule configured for the
// topic.
//
// The properties SubscriptionTopicID, Type are required.
type JourneyNodeAddToDigestParam struct {
	// The subscription topic that owns the digest the event is added to.
	SubscriptionTopicID string `json:"subscription_topic_id" api:"required"`
	// Any of "add-to-digest".
	Type string            `json:"type,omitzero" api:"required"`
	ID   param.Opt[string] `json:"id,omitzero"`
	// Condition spec for a journey node. Accepts a single condition atom, an AND/OR
	// group, or an AND/OR nested group. Omit the `conditions` property entirely to
	// express "no conditions".
	Conditions JourneyConditionsFieldUnionParam `json:"conditions,omitzero"`
	paramObj
}

func (r JourneyNodeAddToDigestParam) MarshalJSON() (data []byte, err error) {
	type shadow JourneyNodeAddToDigestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JourneyNodeAddToDigestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[JourneyNodeAddToDigestParam](
		"type", "add-to-digest",
	)
}

// Branch node. Routes to the first entry in `paths[]` whose `conditions` match,
// else falls through to `default.nodes`.
//
// The properties Default, Paths, Type are required.
type JourneyNodeJourneyBranchNodeParam struct {
	Default JourneyNodeJourneyBranchNodeDefaultParam `json:"default,omitzero" api:"required"`
	Paths   []JourneyNodeJourneyBranchNodePathParam  `json:"paths,omitzero" api:"required"`
	// Any of "branch".
	Type string            `json:"type,omitzero" api:"required"`
	ID   param.Opt[string] `json:"id,omitzero"`
	paramObj
}

func (r JourneyNodeJourneyBranchNodeParam) MarshalJSON() (data []byte, err error) {
	type shadow JourneyNodeJourneyBranchNodeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JourneyNodeJourneyBranchNodeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[JourneyNodeJourneyBranchNodeParam](
		"type", "branch",
	)
}

// The property Nodes is required.
type JourneyNodeJourneyBranchNodeDefaultParam struct {
	Nodes []JourneyNodeUnionParam `json:"nodes,omitzero" api:"required"`
	Label param.Opt[string]       `json:"label,omitzero"`
	paramObj
}

func (r JourneyNodeJourneyBranchNodeDefaultParam) MarshalJSON() (data []byte, err error) {
	type shadow JourneyNodeJourneyBranchNodeDefaultParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JourneyNodeJourneyBranchNodeDefaultParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Conditions, Nodes are required.
type JourneyNodeJourneyBranchNodePathParam struct {
	// Condition spec for a journey node. Accepts a single condition atom, an AND/OR
	// group, or an AND/OR nested group. Omit the `conditions` property entirely to
	// express "no conditions".
	Conditions JourneyConditionsFieldUnionParam `json:"conditions,omitzero" api:"required"`
	Nodes      []JourneyNodeUnionParam          `json:"nodes,omitzero" api:"required"`
	Label      param.Opt[string]                `json:"label,omitzero"`
	paramObj
}

func (r JourneyNodeJourneyBranchNodePathParam) MarshalJSON() (data []byte, err error) {
	type shadow JourneyNodeJourneyBranchNodePathParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JourneyNodeJourneyBranchNodePathParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body for publishing a journey. Pass `version` to roll back to a prior
// version; omit to publish the current draft.
type JourneyPublishRequestParam struct {
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

func (r JourneyPublishRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow JourneyPublishRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JourneyPublishRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A journey, with its current draft or published nodes and metadata.
type JourneyResponse struct {
	ID        string             `json:"id" api:"required"`
	Created   int64              `json:"created" api:"required"`
	Creator   string             `json:"creator" api:"required"`
	Enabled   bool               `json:"enabled" api:"required"`
	Name      string             `json:"name" api:"required"`
	Nodes     []JourneyNodeUnion `json:"nodes" api:"required"`
	Published int64              `json:"published" api:"required"`
	// Lifecycle state of a journey.
	//
	// Any of "DRAFT", "PUBLISHED".
	State   JourneyState `json:"state" api:"required"`
	Updated int64        `json:"updated" api:"required"`
	Updater string       `json:"updater" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Created     respjson.Field
		Creator     respjson.Field
		Enabled     respjson.Field
		Name        respjson.Field
		Nodes       respjson.Field
		Published   respjson.Field
		State       respjson.Field
		Updated     respjson.Field
		Updater     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JourneyResponse) RawJSON() string { return r.JSON.raw }
func (r *JourneyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Trigger fired by a segment event (`identify`, `group`, or `track`).
type JourneySegmentTriggerNode struct {
	// Any of "identify", "group", "track".
	RequestType JourneySegmentTriggerNodeRequestType `json:"request_type" api:"required"`
	// Any of "segment".
	TriggerType JourneySegmentTriggerNodeTriggerType `json:"trigger_type" api:"required"`
	// Any of "trigger".
	Type JourneySegmentTriggerNodeType `json:"type" api:"required"`
	ID   string                        `json:"id"`
	// Condition spec for a journey node. Accepts a single condition atom, an AND/OR
	// group, or an AND/OR nested group. Omit the `conditions` property entirely to
	// express "no conditions".
	Conditions JourneyConditionsFieldUnion `json:"conditions"`
	EventID    string                      `json:"event_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RequestType respjson.Field
		TriggerType respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		Conditions  respjson.Field
		EventID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JourneySegmentTriggerNode) RawJSON() string { return r.JSON.raw }
func (r *JourneySegmentTriggerNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this JourneySegmentTriggerNode to a
// JourneySegmentTriggerNodeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// JourneySegmentTriggerNodeParam.Overrides()
func (r JourneySegmentTriggerNode) ToParam() JourneySegmentTriggerNodeParam {
	return param.Override[JourneySegmentTriggerNodeParam](json.RawMessage(r.RawJSON()))
}

type JourneySegmentTriggerNodeRequestType string

const (
	JourneySegmentTriggerNodeRequestTypeIdentify JourneySegmentTriggerNodeRequestType = "identify"
	JourneySegmentTriggerNodeRequestTypeGroup    JourneySegmentTriggerNodeRequestType = "group"
	JourneySegmentTriggerNodeRequestTypeTrack    JourneySegmentTriggerNodeRequestType = "track"
)

type JourneySegmentTriggerNodeTriggerType string

const (
	JourneySegmentTriggerNodeTriggerTypeSegment JourneySegmentTriggerNodeTriggerType = "segment"
)

type JourneySegmentTriggerNodeType string

const (
	JourneySegmentTriggerNodeTypeTrigger JourneySegmentTriggerNodeType = "trigger"
)

// Trigger fired by a segment event (`identify`, `group`, or `track`).
//
// The properties RequestType, TriggerType, Type are required.
type JourneySegmentTriggerNodeParam struct {
	// Any of "identify", "group", "track".
	RequestType JourneySegmentTriggerNodeRequestType `json:"request_type,omitzero" api:"required"`
	// Any of "segment".
	TriggerType JourneySegmentTriggerNodeTriggerType `json:"trigger_type,omitzero" api:"required"`
	// Any of "trigger".
	Type    JourneySegmentTriggerNodeType `json:"type,omitzero" api:"required"`
	ID      param.Opt[string]             `json:"id,omitzero"`
	EventID param.Opt[string]             `json:"event_id,omitzero"`
	// Condition spec for a journey node. Accepts a single condition atom, an AND/OR
	// group, or an AND/OR nested group. Omit the `conditions` property entirely to
	// express "no conditions".
	Conditions JourneyConditionsFieldUnionParam `json:"conditions,omitzero"`
	paramObj
}

func (r JourneySegmentTriggerNodeParam) MarshalJSON() (data []byte, err error) {
	type shadow JourneySegmentTriggerNodeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JourneySegmentTriggerNodeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Send to the recipient. A send node sources its content from EXACTLY ONE of
// `message.template` (a single notification template) or `experiment` (an A/B
// split across weighted template variants) — supplying both, or neither, is
// rejected. Optionally override the recipient address, delay the send, or attach
// `data`.
type JourneySendNode struct {
	Message JourneySendNodeMessage `json:"message" api:"required"`
	// Any of "send".
	Type JourneySendNodeType `json:"type" api:"required"`
	ID   string              `json:"id"`
	// Condition spec for a journey node. Accepts a single condition atom, an AND/OR
	// group, or an AND/OR nested group. Omit the `conditions` property entirely to
	// express "no conditions".
	Conditions JourneyConditionsFieldUnion `json:"conditions"`
	// A/B experiment config for a send node. The recipient is deterministically
	// bucketed by `bucketingKey` and routed to one of the `variants` in proportion to
	// its `weight`. Present on a send node INSTEAD OF `message.template`.
	Experiment JourneyExperiment `json:"experiment"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		Conditions  respjson.Field
		Experiment  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JourneySendNode) RawJSON() string { return r.JSON.raw }
func (r *JourneySendNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this JourneySendNode to a JourneySendNodeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// JourneySendNodeParam.Overrides()
func (r JourneySendNode) ToParam() JourneySendNodeParam {
	return param.Override[JourneySendNodeParam](json.RawMessage(r.RawJSON()))
}

type JourneySendNodeMessage struct {
	Data     map[string]any              `json:"data"`
	Delay    JourneySendNodeMessageDelay `json:"delay"`
	Template string                      `json:"template"`
	To       JourneySendNodeMessageTo    `json:"to"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Delay       respjson.Field
		Template    respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JourneySendNodeMessage) RawJSON() string { return r.JSON.raw }
func (r *JourneySendNodeMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JourneySendNodeMessageDelay struct {
	Until    string `json:"until" api:"required"`
	Timezone string `json:"timezone"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Until       respjson.Field
		Timezone    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JourneySendNodeMessageDelay) RawJSON() string { return r.JSON.raw }
func (r *JourneySendNodeMessageDelay) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JourneySendNodeMessageTo struct {
	EmailOverride       string `json:"email_override"`
	PhoneNumberOverride string `json:"phone_number_override"`
	UserIDOverride      string `json:"user_id_override"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EmailOverride       respjson.Field
		PhoneNumberOverride respjson.Field
		UserIDOverride      respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JourneySendNodeMessageTo) RawJSON() string { return r.JSON.raw }
func (r *JourneySendNodeMessageTo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JourneySendNodeType string

const (
	JourneySendNodeTypeSend JourneySendNodeType = "send"
)

// Send to the recipient. A send node sources its content from EXACTLY ONE of
// `message.template` (a single notification template) or `experiment` (an A/B
// split across weighted template variants) — supplying both, or neither, is
// rejected. Optionally override the recipient address, delay the send, or attach
// `data`.
//
// The properties Message, Type are required.
type JourneySendNodeParam struct {
	Message JourneySendNodeMessageParam `json:"message,omitzero" api:"required"`
	// Any of "send".
	Type JourneySendNodeType `json:"type,omitzero" api:"required"`
	ID   param.Opt[string]   `json:"id,omitzero"`
	// Condition spec for a journey node. Accepts a single condition atom, an AND/OR
	// group, or an AND/OR nested group. Omit the `conditions` property entirely to
	// express "no conditions".
	Conditions JourneyConditionsFieldUnionParam `json:"conditions,omitzero"`
	// A/B experiment config for a send node. The recipient is deterministically
	// bucketed by `bucketingKey` and routed to one of the `variants` in proportion to
	// its `weight`. Present on a send node INSTEAD OF `message.template`.
	Experiment JourneyExperimentParam `json:"experiment,omitzero"`
	paramObj
}

func (r JourneySendNodeParam) MarshalJSON() (data []byte, err error) {
	type shadow JourneySendNodeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JourneySendNodeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JourneySendNodeMessageParam struct {
	Template param.Opt[string]                `json:"template,omitzero"`
	Data     map[string]any                   `json:"data,omitzero"`
	Delay    JourneySendNodeMessageDelayParam `json:"delay,omitzero"`
	To       JourneySendNodeMessageToParam    `json:"to,omitzero"`
	paramObj
}

func (r JourneySendNodeMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow JourneySendNodeMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JourneySendNodeMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Until is required.
type JourneySendNodeMessageDelayParam struct {
	Until    string            `json:"until" api:"required"`
	Timezone param.Opt[string] `json:"timezone,omitzero"`
	paramObj
}

func (r JourneySendNodeMessageDelayParam) MarshalJSON() (data []byte, err error) {
	type shadow JourneySendNodeMessageDelayParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JourneySendNodeMessageDelayParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JourneySendNodeMessageToParam struct {
	EmailOverride       param.Opt[string] `json:"email_override,omitzero"`
	PhoneNumberOverride param.Opt[string] `json:"phone_number_override,omitzero"`
	UserIDOverride      param.Opt[string] `json:"user_id_override,omitzero"`
	paramObj
}

func (r JourneySendNodeMessageToParam) MarshalJSON() (data []byte, err error) {
	type shadow JourneySendNodeMessageToParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JourneySendNodeMessageToParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lifecycle state of a journey.
type JourneyState string

const (
	JourneyStateDraft     JourneyState = "DRAFT"
	JourneyStatePublished JourneyState = "PUBLISHED"
)

// Request body for creating a notification template scoped to a journey.
//
// The properties Channel, Notification are required.
type JourneyTemplateCreateRequestParam struct {
	Channel      string                                        `json:"channel" api:"required"`
	Notification JourneyTemplateCreateRequestNotificationParam `json:"notification,omitzero" api:"required"`
	ProviderKey  param.Opt[string]                             `json:"providerKey,omitzero"`
	State        param.Opt[string]                             `json:"state,omitzero"`
	paramObj
}

func (r JourneyTemplateCreateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow JourneyTemplateCreateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JourneyTemplateCreateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Brand, Content, Name, Subscription, Tags are required.
type JourneyTemplateCreateRequestNotificationParam struct {
	Brand        JourneyTemplateCreateRequestNotificationBrandParam        `json:"brand,omitzero" api:"required"`
	Subscription JourneyTemplateCreateRequestNotificationSubscriptionParam `json:"subscription,omitzero" api:"required"`
	Content      JourneyTemplateCreateRequestNotificationContentParam      `json:"content,omitzero" api:"required"`
	Name         string                                                    `json:"name" api:"required"`
	Tags         []string                                                  `json:"tags,omitzero" api:"required"`
	paramObj
}

func (r JourneyTemplateCreateRequestNotificationParam) MarshalJSON() (data []byte, err error) {
	type shadow JourneyTemplateCreateRequestNotificationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JourneyTemplateCreateRequestNotificationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type JourneyTemplateCreateRequestNotificationBrandParam struct {
	ID string `json:"id" api:"required"`
	paramObj
}

func (r JourneyTemplateCreateRequestNotificationBrandParam) MarshalJSON() (data []byte, err error) {
	type shadow JourneyTemplateCreateRequestNotificationBrandParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JourneyTemplateCreateRequestNotificationBrandParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Elements, Version are required.
type JourneyTemplateCreateRequestNotificationContentParam struct {
	Elements []shared.ElementalNodeUnionParam `json:"elements,omitzero" api:"required"`
	// Any of "2022-01-01".
	Version string `json:"version,omitzero" api:"required"`
	// Any of "default", "strict".
	Scope string `json:"scope,omitzero"`
	paramObj
}

func (r JourneyTemplateCreateRequestNotificationContentParam) MarshalJSON() (data []byte, err error) {
	type shadow JourneyTemplateCreateRequestNotificationContentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JourneyTemplateCreateRequestNotificationContentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[JourneyTemplateCreateRequestNotificationContentParam](
		"version", "2022-01-01",
	)
	apijson.RegisterFieldValidator[JourneyTemplateCreateRequestNotificationContentParam](
		"scope", "default", "strict",
	)
}

// The property TopicID is required.
type JourneyTemplateCreateRequestNotificationSubscriptionParam struct {
	TopicID string `json:"topic_id" api:"required"`
	paramObj
}

func (r JourneyTemplateCreateRequestNotificationSubscriptionParam) MarshalJSON() (data []byte, err error) {
	type shadow JourneyTemplateCreateRequestNotificationSubscriptionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JourneyTemplateCreateRequestNotificationSubscriptionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A journey-scoped notification template.
type JourneyTemplateGetResponse struct {
	ID      string                            `json:"id" api:"required"`
	Brand   JourneyTemplateGetResponseBrand   `json:"brand" api:"required"`
	Content JourneyTemplateGetResponseContent `json:"content" api:"required"`
	Created int64                             `json:"created" api:"required"`
	Creator string                            `json:"creator" api:"required"`
	Name    string                            `json:"name" api:"required"`
	// Any of "DRAFT", "PUBLISHED".
	State        JourneyTemplateGetResponseState        `json:"state" api:"required"`
	Subscription JourneyTemplateGetResponseSubscription `json:"subscription" api:"required"`
	Tags         []string                               `json:"tags" api:"required"`
	Updated      int64                                  `json:"updated"`
	Updater      string                                 `json:"updater"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Brand        respjson.Field
		Content      respjson.Field
		Created      respjson.Field
		Creator      respjson.Field
		Name         respjson.Field
		State        respjson.Field
		Subscription respjson.Field
		Tags         respjson.Field
		Updated      respjson.Field
		Updater      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JourneyTemplateGetResponse) RawJSON() string { return r.JSON.raw }
func (r *JourneyTemplateGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JourneyTemplateGetResponseBrand struct {
	ID string `json:"id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JourneyTemplateGetResponseBrand) RawJSON() string { return r.JSON.raw }
func (r *JourneyTemplateGetResponseBrand) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JourneyTemplateGetResponseContent struct {
	Elements []shared.ElementalNodeUnion `json:"elements" api:"required"`
	// Any of "2022-01-01".
	Version string `json:"version" api:"required"`
	// Any of "default", "strict".
	Scope string `json:"scope"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Elements    respjson.Field
		Version     respjson.Field
		Scope       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JourneyTemplateGetResponseContent) RawJSON() string { return r.JSON.raw }
func (r *JourneyTemplateGetResponseContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JourneyTemplateGetResponseState string

const (
	JourneyTemplateGetResponseStateDraft     JourneyTemplateGetResponseState = "DRAFT"
	JourneyTemplateGetResponseStatePublished JourneyTemplateGetResponseState = "PUBLISHED"
)

type JourneyTemplateGetResponseSubscription struct {
	TopicID string `json:"topic_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TopicID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JourneyTemplateGetResponseSubscription) RawJSON() string { return r.JSON.raw }
func (r *JourneyTemplateGetResponseSubscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Paged list of journey-scoped notification templates.
type JourneyTemplateListResponse struct {
	Paging  shared.Paging            `json:"paging" api:"required"`
	Results []JourneyTemplateSummary `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Paging      respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JourneyTemplateListResponse) RawJSON() string { return r.JSON.raw }
func (r *JourneyTemplateListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body for publishing a journey-scoped notification template. Pass
// `version` to roll back to a prior version; omit to publish the current draft.
type JourneyTemplatePublishRequestParam struct {
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

func (r JourneyTemplatePublishRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow JourneyTemplatePublishRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JourneyTemplatePublishRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body for replacing a journey-scoped notification template draft.
//
// The property Notification is required.
type JourneyTemplateReplaceRequestParam struct {
	Notification JourneyTemplateReplaceRequestNotificationParam `json:"notification,omitzero" api:"required"`
	State        param.Opt[string]                              `json:"state,omitzero"`
	paramObj
}

func (r JourneyTemplateReplaceRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow JourneyTemplateReplaceRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JourneyTemplateReplaceRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Brand, Content, Name, Subscription, Tags are required.
type JourneyTemplateReplaceRequestNotificationParam struct {
	Brand        JourneyTemplateReplaceRequestNotificationBrandParam        `json:"brand,omitzero" api:"required"`
	Subscription JourneyTemplateReplaceRequestNotificationSubscriptionParam `json:"subscription,omitzero" api:"required"`
	Content      JourneyTemplateReplaceRequestNotificationContentParam      `json:"content,omitzero" api:"required"`
	Name         string                                                     `json:"name" api:"required"`
	Tags         []string                                                   `json:"tags,omitzero" api:"required"`
	paramObj
}

func (r JourneyTemplateReplaceRequestNotificationParam) MarshalJSON() (data []byte, err error) {
	type shadow JourneyTemplateReplaceRequestNotificationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JourneyTemplateReplaceRequestNotificationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type JourneyTemplateReplaceRequestNotificationBrandParam struct {
	ID string `json:"id" api:"required"`
	paramObj
}

func (r JourneyTemplateReplaceRequestNotificationBrandParam) MarshalJSON() (data []byte, err error) {
	type shadow JourneyTemplateReplaceRequestNotificationBrandParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JourneyTemplateReplaceRequestNotificationBrandParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Elements, Version are required.
type JourneyTemplateReplaceRequestNotificationContentParam struct {
	Elements []shared.ElementalNodeUnionParam `json:"elements,omitzero" api:"required"`
	// Any of "2022-01-01".
	Version string `json:"version,omitzero" api:"required"`
	// Any of "default", "strict".
	Scope string `json:"scope,omitzero"`
	paramObj
}

func (r JourneyTemplateReplaceRequestNotificationContentParam) MarshalJSON() (data []byte, err error) {
	type shadow JourneyTemplateReplaceRequestNotificationContentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JourneyTemplateReplaceRequestNotificationContentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[JourneyTemplateReplaceRequestNotificationContentParam](
		"version", "2022-01-01",
	)
	apijson.RegisterFieldValidator[JourneyTemplateReplaceRequestNotificationContentParam](
		"scope", "default", "strict",
	)
}

// The property TopicID is required.
type JourneyTemplateReplaceRequestNotificationSubscriptionParam struct {
	TopicID string `json:"topic_id" api:"required"`
	paramObj
}

func (r JourneyTemplateReplaceRequestNotificationSubscriptionParam) MarshalJSON() (data []byte, err error) {
	type shadow JourneyTemplateReplaceRequestNotificationSubscriptionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JourneyTemplateReplaceRequestNotificationSubscriptionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Summary fields of a journey-scoped notification template returned in list
// responses.
type JourneyTemplateSummary struct {
	ID      string   `json:"id" api:"required"`
	Created int64    `json:"created" api:"required"`
	Creator string   `json:"creator" api:"required"`
	Name    string   `json:"name" api:"required"`
	State   string   `json:"state" api:"required"`
	Tags    []string `json:"tags" api:"required"`
	Updated int64    `json:"updated"`
	Updater string   `json:"updater"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Created     respjson.Field
		Creator     respjson.Field
		Name        respjson.Field
		State       respjson.Field
		Tags        respjson.Field
		Updated     respjson.Field
		Updater     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JourneyTemplateSummary) RawJSON() string { return r.JSON.raw }
func (r *JourneyTemplateSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Throttle the journey by a dynamic `throttle_key`, allowing at most `max_allowed`
// invocations per `period`.
type JourneyThrottleDynamicNode struct {
	MaxAllowed int64  `json:"max_allowed" api:"required"`
	Period     string `json:"period" api:"required"`
	// Any of "dynamic".
	Scope       JourneyThrottleDynamicNodeScope `json:"scope" api:"required"`
	ThrottleKey string                          `json:"throttle_key" api:"required"`
	// Any of "throttle".
	Type JourneyThrottleDynamicNodeType `json:"type" api:"required"`
	ID   string                         `json:"id"`
	// Condition spec for a journey node. Accepts a single condition atom, an AND/OR
	// group, or an AND/OR nested group. Omit the `conditions` property entirely to
	// express "no conditions".
	Conditions JourneyConditionsFieldUnion `json:"conditions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxAllowed  respjson.Field
		Period      respjson.Field
		Scope       respjson.Field
		ThrottleKey respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		Conditions  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JourneyThrottleDynamicNode) RawJSON() string { return r.JSON.raw }
func (r *JourneyThrottleDynamicNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this JourneyThrottleDynamicNode to a
// JourneyThrottleDynamicNodeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// JourneyThrottleDynamicNodeParam.Overrides()
func (r JourneyThrottleDynamicNode) ToParam() JourneyThrottleDynamicNodeParam {
	return param.Override[JourneyThrottleDynamicNodeParam](json.RawMessage(r.RawJSON()))
}

type JourneyThrottleDynamicNodeScope string

const (
	JourneyThrottleDynamicNodeScopeDynamic JourneyThrottleDynamicNodeScope = "dynamic"
)

type JourneyThrottleDynamicNodeType string

const (
	JourneyThrottleDynamicNodeTypeThrottle JourneyThrottleDynamicNodeType = "throttle"
)

// Throttle the journey by a dynamic `throttle_key`, allowing at most `max_allowed`
// invocations per `period`.
//
// The properties MaxAllowed, Period, Scope, ThrottleKey, Type are required.
type JourneyThrottleDynamicNodeParam struct {
	MaxAllowed int64  `json:"max_allowed" api:"required"`
	Period     string `json:"period" api:"required"`
	// Any of "dynamic".
	Scope       JourneyThrottleDynamicNodeScope `json:"scope,omitzero" api:"required"`
	ThrottleKey string                          `json:"throttle_key" api:"required"`
	// Any of "throttle".
	Type JourneyThrottleDynamicNodeType `json:"type,omitzero" api:"required"`
	ID   param.Opt[string]              `json:"id,omitzero"`
	// Condition spec for a journey node. Accepts a single condition atom, an AND/OR
	// group, or an AND/OR nested group. Omit the `conditions` property entirely to
	// express "no conditions".
	Conditions JourneyConditionsFieldUnionParam `json:"conditions,omitzero"`
	paramObj
}

func (r JourneyThrottleDynamicNodeParam) MarshalJSON() (data []byte, err error) {
	type shadow JourneyThrottleDynamicNodeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JourneyThrottleDynamicNodeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Throttle the journey by a static `scope` (`user` or `global`), allowing at most
// `max_allowed` invocations per `period`.
type JourneyThrottleStaticNode struct {
	MaxAllowed int64  `json:"max_allowed" api:"required"`
	Period     string `json:"period" api:"required"`
	// Any of "user", "global".
	Scope JourneyThrottleStaticNodeScope `json:"scope" api:"required"`
	// Any of "throttle".
	Type JourneyThrottleStaticNodeType `json:"type" api:"required"`
	ID   string                        `json:"id"`
	// Condition spec for a journey node. Accepts a single condition atom, an AND/OR
	// group, or an AND/OR nested group. Omit the `conditions` property entirely to
	// express "no conditions".
	Conditions JourneyConditionsFieldUnion `json:"conditions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxAllowed  respjson.Field
		Period      respjson.Field
		Scope       respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		Conditions  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JourneyThrottleStaticNode) RawJSON() string { return r.JSON.raw }
func (r *JourneyThrottleStaticNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this JourneyThrottleStaticNode to a
// JourneyThrottleStaticNodeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// JourneyThrottleStaticNodeParam.Overrides()
func (r JourneyThrottleStaticNode) ToParam() JourneyThrottleStaticNodeParam {
	return param.Override[JourneyThrottleStaticNodeParam](json.RawMessage(r.RawJSON()))
}

type JourneyThrottleStaticNodeScope string

const (
	JourneyThrottleStaticNodeScopeUser   JourneyThrottleStaticNodeScope = "user"
	JourneyThrottleStaticNodeScopeGlobal JourneyThrottleStaticNodeScope = "global"
)

type JourneyThrottleStaticNodeType string

const (
	JourneyThrottleStaticNodeTypeThrottle JourneyThrottleStaticNodeType = "throttle"
)

// Throttle the journey by a static `scope` (`user` or `global`), allowing at most
// `max_allowed` invocations per `period`.
//
// The properties MaxAllowed, Period, Scope, Type are required.
type JourneyThrottleStaticNodeParam struct {
	MaxAllowed int64  `json:"max_allowed" api:"required"`
	Period     string `json:"period" api:"required"`
	// Any of "user", "global".
	Scope JourneyThrottleStaticNodeScope `json:"scope,omitzero" api:"required"`
	// Any of "throttle".
	Type JourneyThrottleStaticNodeType `json:"type,omitzero" api:"required"`
	ID   param.Opt[string]             `json:"id,omitzero"`
	// Condition spec for a journey node. Accepts a single condition atom, an AND/OR
	// group, or an AND/OR nested group. Omit the `conditions` property entirely to
	// express "no conditions".
	Conditions JourneyConditionsFieldUnionParam `json:"conditions,omitzero"`
	paramObj
}

func (r JourneyThrottleStaticNodeParam) MarshalJSON() (data []byte, err error) {
	type shadow JourneyThrottleStaticNodeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JourneyThrottleStaticNodeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A published version of a journey.
type JourneyVersionItem struct {
	Created   int64  `json:"created" api:"required"`
	Creator   string `json:"creator" api:"required"`
	Name      string `json:"name" api:"required"`
	Published int64  `json:"published" api:"required"`
	Version   string `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Created     respjson.Field
		Creator     respjson.Field
		Name        respjson.Field
		Published   respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JourneyVersionItem) RawJSON() string { return r.JSON.raw }
func (r *JourneyVersionItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Paged list of published journey versions, most recent first.
type JourneyVersionsListResponse struct {
	Paging  shared.Paging        `json:"paging" api:"required"`
	Results []JourneyVersionItem `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Paging      respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JourneyVersionsListResponse) RawJSON() string { return r.JSON.raw }
func (r *JourneyVersionsListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body for invoking a journey. Requires either a user identifier or a
// profile with contact information. User identifiers can be provided via user_id
// field, or resolved from profile/data objects (user_id, userId, or anonymousId
// fields).
type JourneysInvokeRequestParam struct {
	// A unique identifier for the user. If not provided, the system will attempt to
	// resolve the user identifier from profile or data objects.
	UserID param.Opt[string] `json:"user_id,omitzero"`
	// Data payload passed to the journey. The expected shape can be predefined using
	// the schema builder in the journey editor. This data is available in journey
	// steps for condition evaluation and template variable interpolation. Can also
	// contain user identifiers (user_id, userId, anonymousId) if not provided
	// elsewhere.
	Data map[string]any `json:"data,omitzero"`
	// Profile data for the user. Can contain contact information (email,
	// phone_number), user identifiers (user_id, userId, anonymousId), or any custom
	// profile fields. Profile fields are merged with any existing stored profile for
	// the user. Include context.tenant_id to load a tenant-scoped profile for
	// multi-tenant scenarios.
	Profile map[string]any `json:"profile,omitzero"`
	paramObj
}

func (r JourneysInvokeRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow JourneysInvokeRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JourneysInvokeRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JourneysInvokeResponse struct {
	// A unique identifier for the journey run that was created.
	RunID string `json:"runId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RunID       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JourneysInvokeResponse) RawJSON() string { return r.JSON.raw }
func (r *JourneysInvokeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JourneysListResponse struct {
	// A cursor token for pagination. Present when there are more results available.
	Cursor    string    `json:"cursor"`
	Templates []Journey `json:"templates"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cursor      respjson.Field
		Templates   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JourneysListResponse) RawJSON() string { return r.JSON.raw }
func (r *JourneysListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JourneyNewParams struct {
	// Request body for creating a journey.
	CreateJourneyRequest CreateJourneyRequestParam
	paramObj
}

func (r JourneyNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateJourneyRequest)
}
func (r *JourneyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JourneyGetParams struct {
	// Version selector: `draft`, `published` (default), or `vN`.
	Version param.Opt[string] `query:"version,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [JourneyGetParams]'s query parameters as `url.Values`.
func (r JourneyGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type JourneyListParams struct {
	// A cursor token for pagination. Use the cursor from the previous response to
	// fetch the next page of results.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// The version of journeys to retrieve. Accepted values are published (for
	// published journeys) or draft (for draft journeys). Defaults to published.
	//
	// Any of "published", "draft".
	Version JourneyListParamsVersion `query:"version,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [JourneyListParams]'s query parameters as `url.Values`.
func (r JourneyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The version of journeys to retrieve. Accepted values are published (for
// published journeys) or draft (for draft journeys). Defaults to published.
type JourneyListParamsVersion string

const (
	JourneyListParamsVersionPublished JourneyListParamsVersion = "published"
	JourneyListParamsVersionDraft     JourneyListParamsVersion = "draft"
)

type JourneyCancelParams struct {
	// Request body for `POST /journeys/cancel`. Provide EXACTLY ONE of
	// `cancelation_token` (cancels every run associated with the token) or `run_id`
	// (cancels a single tenant-scoped run).
	CancelJourneyRequest CancelJourneyRequestUnionParam
	paramObj
}

func (r JourneyCancelParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CancelJourneyRequest)
}
func (r *JourneyCancelParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JourneyInvokeParams struct {
	// Request body for invoking a journey. Requires either a user identifier or a
	// profile with contact information. User identifiers can be provided via user_id
	// field, or resolved from profile/data objects (user_id, userId, or anonymousId
	// fields).
	JourneysInvokeRequest JourneysInvokeRequestParam
	paramObj
}

func (r JourneyInvokeParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.JourneysInvokeRequest)
}
func (r *JourneyInvokeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JourneyPublishParams struct {
	// Request body for publishing a journey. Pass `version` to roll back to a prior
	// version; omit to publish the current draft.
	JourneyPublishRequest JourneyPublishRequestParam
	paramObj
}

func (r JourneyPublishParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.JourneyPublishRequest)
}
func (r *JourneyPublishParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JourneyReplaceParams struct {
	// Request body for creating a journey.
	CreateJourneyRequest CreateJourneyRequestParam
	paramObj
}

func (r JourneyReplaceParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateJourneyRequest)
}
func (r *JourneyReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
