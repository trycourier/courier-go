// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"encoding/json"

	"github.com/trycourier/courier-go/v3/internal/apijson"
	"github.com/trycourier/courier-go/v3/packages/param"
	"github.com/trycourier/courier-go/v3/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Audience struct {
	// A unique identifier representing the audience_id
	ID        string `json:"id,required"`
	CreatedAt string `json:"created_at,required"`
	// A description of the audience
	Description string `json:"description,required"`
	// A single filter to use for filtering
	Filter Filter `json:"filter,required"`
	// The name of the audience
	Name      string `json:"name,required"`
	UpdatedAt string `json:"updated_at,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		Filter      respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Audience) RawJSON() string { return r.JSON.raw }
func (r *Audience) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuditEvent struct {
	Actor        AuditEventActor `json:"actor,required"`
	AuditEventID string          `json:"auditEventId,required"`
	Source       string          `json:"source,required"`
	Target       string          `json:"target,required"`
	Timestamp    string          `json:"timestamp,required"`
	Type         string          `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Actor        respjson.Field
		AuditEventID respjson.Field
		Source       respjson.Field
		Target       respjson.Field
		Timestamp    respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuditEvent) RawJSON() string { return r.JSON.raw }
func (r *AuditEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuditEventActor struct {
	ID    string `json:"id,required"`
	Email string `json:"email,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Email       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuditEventActor) RawJSON() string { return r.JSON.raw }
func (r *AuditEventActor) UnmarshalJSON(data []byte) error {
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

type BaseCheck struct {
	ID string `json:"id,required"`
	// Any of "RESOLVED", "FAILED", "PENDING".
	Status BaseCheckStatus `json:"status,required"`
	// Any of "custom".
	Type BaseCheckType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaseCheck) RawJSON() string { return r.JSON.raw }
func (r *BaseCheck) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BaseCheck to a BaseCheckParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BaseCheckParam.Overrides()
func (r BaseCheck) ToParam() BaseCheckParam {
	return param.Override[BaseCheckParam](json.RawMessage(r.RawJSON()))
}

type BaseCheckStatus string

const (
	BaseCheckStatusResolved BaseCheckStatus = "RESOLVED"
	BaseCheckStatusFailed   BaseCheckStatus = "FAILED"
	BaseCheckStatusPending  BaseCheckStatus = "PENDING"
)

type BaseCheckType string

const (
	BaseCheckTypeCustom BaseCheckType = "custom"
)

// The properties ID, Status, Type are required.
type BaseCheckParam struct {
	ID string `json:"id,required"`
	// Any of "RESOLVED", "FAILED", "PENDING".
	Status BaseCheckStatus `json:"status,omitzero,required"`
	// Any of "custom".
	Type BaseCheckType `json:"type,omitzero,required"`
	paramObj
}

func (r BaseCheckParam) MarshalJSON() (data []byte, err error) {
	type shadow BaseCheckParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaseCheckParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BaseTemplateTenantAssociation struct {
	// The template's id
	ID string `json:"id,required"`
	// The timestamp at which the template was created
	CreatedAt string `json:"created_at,required"`
	// The timestamp at which the template was published
	PublishedAt string `json:"published_at,required"`
	// The timestamp at which the template was last updated
	UpdatedAt string `json:"updated_at,required"`
	// The version of the template
	Version string `json:"version,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		PublishedAt respjson.Field
		UpdatedAt   respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaseTemplateTenantAssociation) RawJSON() string { return r.JSON.raw }
func (r *BaseTemplateTenantAssociation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Brand struct {
	ID        string        `json:"id,required"`
	Created   int64         `json:"created,required"`
	Name      string        `json:"name,required"`
	Updated   int64         `json:"updated,required"`
	Published int64         `json:"published,nullable"`
	Settings  BrandSettings `json:"settings,nullable"`
	Snippets  BrandSnippets `json:"snippets,nullable"`
	Version   string        `json:"version,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Created     respjson.Field
		Name        respjson.Field
		Updated     respjson.Field
		Published   respjson.Field
		Settings    respjson.Field
		Snippets    respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Brand) RawJSON() string { return r.JSON.raw }
func (r *Brand) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandColors struct {
	Primary     string            `json:"primary"`
	Secondary   string            `json:"secondary"`
	ExtraFields map[string]string `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Primary     respjson.Field
		Secondary   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandColors) RawJSON() string { return r.JSON.raw }
func (r *BrandColors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BrandColors to a BrandColorsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BrandColorsParam.Overrides()
func (r BrandColors) ToParam() BrandColorsParam {
	return param.Override[BrandColorsParam](json.RawMessage(r.RawJSON()))
}

type BrandColorsParam struct {
	Primary     param.Opt[string] `json:"primary,omitzero"`
	Secondary   param.Opt[string] `json:"secondary,omitzero"`
	ExtraFields map[string]string `json:"-"`
	paramObj
}

func (r BrandColorsParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandColorsParam
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *BrandColorsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandSettings struct {
	Colors BrandColors        `json:"colors,nullable"`
	Email  BrandSettingsEmail `json:"email,nullable"`
	Inapp  BrandSettingsInApp `json:"inapp,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Colors      respjson.Field
		Email       respjson.Field
		Inapp       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandSettings) RawJSON() string { return r.JSON.raw }
func (r *BrandSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BrandSettings to a BrandSettingsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BrandSettingsParam.Overrides()
func (r BrandSettings) ToParam() BrandSettingsParam {
	return param.Override[BrandSettingsParam](json.RawMessage(r.RawJSON()))
}

type BrandSettingsParam struct {
	Colors BrandColorsParam        `json:"colors,omitzero"`
	Email  BrandSettingsEmailParam `json:"email,omitzero"`
	Inapp  BrandSettingsInAppParam `json:"inapp,omitzero"`
	paramObj
}

func (r BrandSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandSettingsEmail struct {
	Footer           EmailFooter                        `json:"footer,nullable"`
	Head             EmailHead                          `json:"head,nullable"`
	Header           EmailHeader                        `json:"header,nullable"`
	TemplateOverride BrandSettingsEmailTemplateOverride `json:"templateOverride,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Footer           respjson.Field
		Head             respjson.Field
		Header           respjson.Field
		TemplateOverride respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandSettingsEmail) RawJSON() string { return r.JSON.raw }
func (r *BrandSettingsEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BrandSettingsEmail to a BrandSettingsEmailParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BrandSettingsEmailParam.Overrides()
func (r BrandSettingsEmail) ToParam() BrandSettingsEmailParam {
	return param.Override[BrandSettingsEmailParam](json.RawMessage(r.RawJSON()))
}

type BrandSettingsEmailTemplateOverride struct {
	Mjml                  BrandTemplate `json:"mjml,required"`
	FooterBackgroundColor string        `json:"footerBackgroundColor,nullable"`
	FooterFullWidth       bool          `json:"footerFullWidth,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Mjml                  respjson.Field
		FooterBackgroundColor respjson.Field
		FooterFullWidth       respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
	BrandTemplate
}

// Returns the unmodified JSON received from the API
func (r BrandSettingsEmailTemplateOverride) RawJSON() string { return r.JSON.raw }
func (r *BrandSettingsEmailTemplateOverride) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandSettingsEmailParam struct {
	TemplateOverride BrandSettingsEmailTemplateOverrideParam `json:"templateOverride,omitzero"`
	Footer           EmailFooterParam                        `json:"footer,omitzero"`
	Head             EmailHeadParam                          `json:"head,omitzero"`
	Header           EmailHeaderParam                        `json:"header,omitzero"`
	paramObj
}

func (r BrandSettingsEmailParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandSettingsEmailParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandSettingsEmailParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandSettingsEmailTemplateOverrideParam struct {
	Mjml                  BrandTemplateParam `json:"mjml,omitzero,required"`
	FooterBackgroundColor param.Opt[string]  `json:"footerBackgroundColor,omitzero"`
	FooterFullWidth       param.Opt[bool]    `json:"footerFullWidth,omitzero"`
	BrandTemplateParam
}

func (r BrandSettingsEmailTemplateOverrideParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandSettingsEmailTemplateOverrideParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type BrandSettingsInApp struct {
	Colors             BrandColors      `json:"colors,required"`
	Icons              Icons            `json:"icons,required"`
	WidgetBackground   WidgetBackground `json:"widgetBackground,required"`
	BorderRadius       string           `json:"borderRadius,nullable"`
	DisableMessageIcon bool             `json:"disableMessageIcon,nullable"`
	FontFamily         string           `json:"fontFamily,nullable"`
	// Any of "top", "bottom", "left", "right".
	Placement BrandSettingsInAppPlacement `json:"placement,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Colors             respjson.Field
		Icons              respjson.Field
		WidgetBackground   respjson.Field
		BorderRadius       respjson.Field
		DisableMessageIcon respjson.Field
		FontFamily         respjson.Field
		Placement          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandSettingsInApp) RawJSON() string { return r.JSON.raw }
func (r *BrandSettingsInApp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BrandSettingsInApp to a BrandSettingsInAppParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BrandSettingsInAppParam.Overrides()
func (r BrandSettingsInApp) ToParam() BrandSettingsInAppParam {
	return param.Override[BrandSettingsInAppParam](json.RawMessage(r.RawJSON()))
}

type BrandSettingsInAppPlacement string

const (
	BrandSettingsInAppPlacementTop    BrandSettingsInAppPlacement = "top"
	BrandSettingsInAppPlacementBottom BrandSettingsInAppPlacement = "bottom"
	BrandSettingsInAppPlacementLeft   BrandSettingsInAppPlacement = "left"
	BrandSettingsInAppPlacementRight  BrandSettingsInAppPlacement = "right"
)

// The properties Colors, Icons, WidgetBackground are required.
type BrandSettingsInAppParam struct {
	Colors             BrandColorsParam      `json:"colors,omitzero,required"`
	Icons              IconsParam            `json:"icons,omitzero,required"`
	WidgetBackground   WidgetBackgroundParam `json:"widgetBackground,omitzero,required"`
	BorderRadius       param.Opt[string]     `json:"borderRadius,omitzero"`
	DisableMessageIcon param.Opt[bool]       `json:"disableMessageIcon,omitzero"`
	FontFamily         param.Opt[string]     `json:"fontFamily,omitzero"`
	// Any of "top", "bottom", "left", "right".
	Placement BrandSettingsInAppPlacement `json:"placement,omitzero"`
	paramObj
}

func (r BrandSettingsInAppParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandSettingsInAppParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandSettingsInAppParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandSnippet struct {
	Name  string `json:"name,required"`
	Value string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandSnippet) RawJSON() string { return r.JSON.raw }
func (r *BrandSnippet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BrandSnippet to a BrandSnippetParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BrandSnippetParam.Overrides()
func (r BrandSnippet) ToParam() BrandSnippetParam {
	return param.Override[BrandSnippetParam](json.RawMessage(r.RawJSON()))
}

// The properties Name, Value are required.
type BrandSnippetParam struct {
	Name  string `json:"name,required"`
	Value string `json:"value,required"`
	paramObj
}

func (r BrandSnippetParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandSnippetParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandSnippetParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandSnippets struct {
	Items []BrandSnippet `json:"items,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandSnippets) RawJSON() string { return r.JSON.raw }
func (r *BrandSnippets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BrandSnippets to a BrandSnippetsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BrandSnippetsParam.Overrides()
func (r BrandSnippets) ToParam() BrandSnippetsParam {
	return param.Override[BrandSnippetsParam](json.RawMessage(r.RawJSON()))
}

type BrandSnippetsParam struct {
	Items []BrandSnippetParam `json:"items,omitzero"`
	paramObj
}

func (r BrandSnippetsParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandSnippetsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandSnippetsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandTemplate struct {
	Enabled               bool   `json:"enabled,required"`
	BackgroundColor       string `json:"backgroundColor,nullable"`
	BlocksBackgroundColor string `json:"blocksBackgroundColor,nullable"`
	Footer                string `json:"footer,nullable"`
	Head                  string `json:"head,nullable"`
	Header                string `json:"header,nullable"`
	Width                 string `json:"width,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled               respjson.Field
		BackgroundColor       respjson.Field
		BlocksBackgroundColor respjson.Field
		Footer                respjson.Field
		Head                  respjson.Field
		Header                respjson.Field
		Width                 respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandTemplate) RawJSON() string { return r.JSON.raw }
func (r *BrandTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BrandTemplate to a BrandTemplateParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BrandTemplateParam.Overrides()
func (r BrandTemplate) ToParam() BrandTemplateParam {
	return param.Override[BrandTemplateParam](json.RawMessage(r.RawJSON()))
}

// The property Enabled is required.
type BrandTemplateParam struct {
	Enabled               bool              `json:"enabled,required"`
	BackgroundColor       param.Opt[string] `json:"backgroundColor,omitzero"`
	BlocksBackgroundColor param.Opt[string] `json:"blocksBackgroundColor,omitzero"`
	Footer                param.Opt[string] `json:"footer,omitzero"`
	Head                  param.Opt[string] `json:"head,omitzero"`
	Header                param.Opt[string] `json:"header,omitzero"`
	Width                 param.Opt[string] `json:"width,omitzero"`
	paramObj
}

func (r BrandTemplateParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandTemplateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandTemplateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChannelClassification string

const (
	ChannelClassificationDirectMessage ChannelClassification = "direct_message"
	ChannelClassificationEmail         ChannelClassification = "email"
	ChannelClassificationPush          ChannelClassification = "push"
	ChannelClassificationSMS           ChannelClassification = "sms"
	ChannelClassificationWebhook       ChannelClassification = "webhook"
	ChannelClassificationInbox         ChannelClassification = "inbox"
)

type ChannelPreference struct {
	// Any of "direct_message", "email", "push", "sms", "webhook", "inbox".
	Channel ChannelClassification `json:"channel,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Channel     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChannelPreference) RawJSON() string { return r.JSON.raw }
func (r *ChannelPreference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ChannelPreference to a ChannelPreferenceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ChannelPreferenceParam.Overrides()
func (r ChannelPreference) ToParam() ChannelPreferenceParam {
	return param.Override[ChannelPreferenceParam](json.RawMessage(r.RawJSON()))
}

// The property Channel is required.
type ChannelPreferenceParam struct {
	// Any of "direct_message", "email", "push", "sms", "webhook", "inbox".
	Channel ChannelClassification `json:"channel,omitzero,required"`
	paramObj
}

func (r ChannelPreferenceParam) MarshalJSON() (data []byte, err error) {
	type shadow ChannelPreferenceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChannelPreferenceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Check struct {
	Updated int64 `json:"updated,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Updated     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BaseCheck
}

// Returns the unmodified JSON received from the API
func (r Check) RawJSON() string { return r.JSON.raw }
func (r *Check) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DefaultPreferences struct {
	Items []DefaultPreferencesItem `json:"items,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DefaultPreferences) RawJSON() string { return r.JSON.raw }
func (r *DefaultPreferences) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DefaultPreferences to a DefaultPreferencesParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DefaultPreferencesParam.Overrides()
func (r DefaultPreferences) ToParam() DefaultPreferencesParam {
	return param.Override[DefaultPreferencesParam](json.RawMessage(r.RawJSON()))
}

type DefaultPreferencesItem struct {
	// Topic ID
	ID string `json:"id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	SubscriptionTopicNew
}

// Returns the unmodified JSON received from the API
func (r DefaultPreferencesItem) RawJSON() string { return r.JSON.raw }
func (r *DefaultPreferencesItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DefaultPreferencesParam struct {
	Items []DefaultPreferencesItemParam `json:"items,omitzero"`
	paramObj
}

func (r DefaultPreferencesParam) MarshalJSON() (data []byte, err error) {
	type shadow DefaultPreferencesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DefaultPreferencesParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DefaultPreferencesItemParam struct {
	// Topic ID
	ID string `json:"id,required"`
	SubscriptionTopicNewParam
}

func (r DefaultPreferencesItemParam) MarshalJSON() (data []byte, err error) {
	type shadow DefaultPreferencesItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type ElementalActionNodeWithType struct {
	// Any of "action".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ElementalBaseNode
}

// Returns the unmodified JSON received from the API
func (r ElementalActionNodeWithType) RawJSON() string { return r.JSON.raw }
func (r *ElementalActionNodeWithType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ElementalActionNodeWithType to a
// ElementalActionNodeWithTypeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ElementalActionNodeWithTypeParam.Overrides()
func (r ElementalActionNodeWithType) ToParam() ElementalActionNodeWithTypeParam {
	return param.Override[ElementalActionNodeWithTypeParam](json.RawMessage(r.RawJSON()))
}

type ElementalActionNodeWithTypeParam struct {
	Type string `json:"type,omitzero"`
	ElementalBaseNodeParam
}

func (r ElementalActionNodeWithTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow ElementalActionNodeWithTypeParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type ElementalBaseNode struct {
	Channels []string `json:"channels,nullable"`
	If       string   `json:"if,nullable"`
	Loop     string   `json:"loop,nullable"`
	Ref      string   `json:"ref,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Channels    respjson.Field
		If          respjson.Field
		Loop        respjson.Field
		Ref         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ElementalBaseNode) RawJSON() string { return r.JSON.raw }
func (r *ElementalBaseNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ElementalBaseNode to a ElementalBaseNodeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ElementalBaseNodeParam.Overrides()
func (r ElementalBaseNode) ToParam() ElementalBaseNodeParam {
	return param.Override[ElementalBaseNodeParam](json.RawMessage(r.RawJSON()))
}

type ElementalBaseNodeParam struct {
	If       param.Opt[string] `json:"if,omitzero"`
	Loop     param.Opt[string] `json:"loop,omitzero"`
	Ref      param.Opt[string] `json:"ref,omitzero"`
	Channels []string          `json:"channels,omitzero"`
	paramObj
}

func (r ElementalBaseNodeParam) MarshalJSON() (data []byte, err error) {
	type shadow ElementalBaseNodeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ElementalBaseNodeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The channel element allows a notification to be customized based on which
// channel it is sent through. For example, you may want to display a detailed
// message when the notification is sent through email, and a more concise message
// in a push notification. Channel elements are only valid as top-level elements;
// you cannot nest channel elements. If there is a channel element specified at the
// top-level of the document, all sibling elements must be channel elements. Note:
// As an alternative, most elements support a `channel` property. Which allows you
// to selectively display an individual element on a per channel basis. See the
// [control flow docs](https://www.courier.com/docs/platform/content/elemental/control-flow/)
// for more details.
type ElementalChannelNode struct {
	// The channel the contents of this element should be applied to. Can be `email`,
	// `push`, `direct_message`, `sms` or a provider such as slack
	Channel string `json:"channel,required"`
	// Raw data to apply to the channel. If `elements` has not been specified, `raw` is
	// `required`.
	Raw map[string]any `json:"raw,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Channel     respjson.Field
		Raw         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ElementalBaseNode
}

// Returns the unmodified JSON received from the API
func (r ElementalChannelNode) RawJSON() string { return r.JSON.raw }
func (r *ElementalChannelNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ElementalChannelNode to a ElementalChannelNodeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ElementalChannelNodeParam.Overrides()
func (r ElementalChannelNode) ToParam() ElementalChannelNodeParam {
	return param.Override[ElementalChannelNodeParam](json.RawMessage(r.RawJSON()))
}

// The channel element allows a notification to be customized based on which
// channel it is sent through. For example, you may want to display a detailed
// message when the notification is sent through email, and a more concise message
// in a push notification. Channel elements are only valid as top-level elements;
// you cannot nest channel elements. If there is a channel element specified at the
// top-level of the document, all sibling elements must be channel elements. Note:
// As an alternative, most elements support a `channel` property. Which allows you
// to selectively display an individual element on a per channel basis. See the
// [control flow docs](https://www.courier.com/docs/platform/content/elemental/control-flow/)
// for more details.
type ElementalChannelNodeParam struct {
	// The channel the contents of this element should be applied to. Can be `email`,
	// `push`, `direct_message`, `sms` or a provider such as slack
	Channel string `json:"channel,required"`
	// Raw data to apply to the channel. If `elements` has not been specified, `raw` is
	// `required`.
	Raw map[string]any `json:"raw,omitzero"`
	ElementalBaseNodeParam
}

func (r ElementalChannelNodeParam) MarshalJSON() (data []byte, err error) {
	type shadow ElementalChannelNodeParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// The channel element allows a notification to be customized based on which
// channel it is sent through. For example, you may want to display a detailed
// message when the notification is sent through email, and a more concise message
// in a push notification. Channel elements are only valid as top-level elements;
// you cannot nest channel elements. If there is a channel element specified at the
// top-level of the document, all sibling elements must be channel elements. Note:
// As an alternative, most elements support a `channel` property. Which allows you
// to selectively display an individual element on a per channel basis. See the
// [control flow docs](https://www.courier.com/docs/platform/content/elemental/control-flow/)
// for more details.
type ElementalChannelNodeWithType struct {
	// Any of "channel".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ElementalChannelNode
}

// Returns the unmodified JSON received from the API
func (r ElementalChannelNodeWithType) RawJSON() string { return r.JSON.raw }
func (r *ElementalChannelNodeWithType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ElementalChannelNodeWithType to a
// ElementalChannelNodeWithTypeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ElementalChannelNodeWithTypeParam.Overrides()
func (r ElementalChannelNodeWithType) ToParam() ElementalChannelNodeWithTypeParam {
	return param.Override[ElementalChannelNodeWithTypeParam](json.RawMessage(r.RawJSON()))
}

// The channel element allows a notification to be customized based on which
// channel it is sent through. For example, you may want to display a detailed
// message when the notification is sent through email, and a more concise message
// in a push notification. Channel elements are only valid as top-level elements;
// you cannot nest channel elements. If there is a channel element specified at the
// top-level of the document, all sibling elements must be channel elements. Note:
// As an alternative, most elements support a `channel` property. Which allows you
// to selectively display an individual element on a per channel basis. See the
// [control flow docs](https://www.courier.com/docs/platform/content/elemental/control-flow/)
// for more details.
type ElementalChannelNodeWithTypeParam struct {
	Type string `json:"type,omitzero"`
	ElementalChannelNodeParam
}

func (r ElementalChannelNodeWithTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow ElementalChannelNodeWithTypeParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type ElementalContent struct {
	Elements []ElementalNodeUnion `json:"elements,required"`
	// For example, "2022-01-01"
	Version string `json:"version,required"`
	Brand   string `json:"brand,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Elements    respjson.Field
		Version     respjson.Field
		Brand       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ElementalContent) RawJSON() string { return r.JSON.raw }
func (r *ElementalContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ElementalContent to a ElementalContentParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ElementalContentParam.Overrides()
func (r ElementalContent) ToParam() ElementalContentParam {
	return param.Override[ElementalContentParam](json.RawMessage(r.RawJSON()))
}

// The properties Elements, Version are required.
type ElementalContentParam struct {
	Elements []ElementalNodeUnionParam `json:"elements,omitzero,required"`
	// For example, "2022-01-01"
	Version string            `json:"version,required"`
	Brand   param.Opt[string] `json:"brand,omitzero"`
	paramObj
}

func (r ElementalContentParam) MarshalJSON() (data []byte, err error) {
	type shadow ElementalContentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ElementalContentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Syntactic sugar to provide a fast shorthand for Courier Elemental Blocks.
type ElementalContentSugar struct {
	// The text content displayed in the notification.
	Body string `json:"body,required"`
	// Title/subject displayed by supported channels.
	Title string `json:"title,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ElementalContentSugar) RawJSON() string { return r.JSON.raw }
func (r *ElementalContentSugar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ElementalContentSugar to a ElementalContentSugarParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ElementalContentSugarParam.Overrides()
func (r ElementalContentSugar) ToParam() ElementalContentSugarParam {
	return param.Override[ElementalContentSugarParam](json.RawMessage(r.RawJSON()))
}

// Syntactic sugar to provide a fast shorthand for Courier Elemental Blocks.
//
// The properties Body, Title are required.
type ElementalContentSugarParam struct {
	// The text content displayed in the notification.
	Body string `json:"body,required"`
	// Title/subject displayed by supported channels.
	Title string `json:"title,required"`
	paramObj
}

func (r ElementalContentSugarParam) MarshalJSON() (data []byte, err error) {
	type shadow ElementalContentSugarParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ElementalContentSugarParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ElementalDividerNodeWithType struct {
	// Any of "divider".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ElementalBaseNode
}

// Returns the unmodified JSON received from the API
func (r ElementalDividerNodeWithType) RawJSON() string { return r.JSON.raw }
func (r *ElementalDividerNodeWithType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ElementalDividerNodeWithType to a
// ElementalDividerNodeWithTypeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ElementalDividerNodeWithTypeParam.Overrides()
func (r ElementalDividerNodeWithType) ToParam() ElementalDividerNodeWithTypeParam {
	return param.Override[ElementalDividerNodeWithTypeParam](json.RawMessage(r.RawJSON()))
}

type ElementalDividerNodeWithTypeParam struct {
	Type string `json:"type,omitzero"`
	ElementalBaseNodeParam
}

func (r ElementalDividerNodeWithTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow ElementalDividerNodeWithTypeParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type ElementalImageNodeWithType struct {
	// Any of "image".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ElementalBaseNode
}

// Returns the unmodified JSON received from the API
func (r ElementalImageNodeWithType) RawJSON() string { return r.JSON.raw }
func (r *ElementalImageNodeWithType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ElementalImageNodeWithType to a
// ElementalImageNodeWithTypeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ElementalImageNodeWithTypeParam.Overrides()
func (r ElementalImageNodeWithType) ToParam() ElementalImageNodeWithTypeParam {
	return param.Override[ElementalImageNodeWithTypeParam](json.RawMessage(r.RawJSON()))
}

type ElementalImageNodeWithTypeParam struct {
	Type string `json:"type,omitzero"`
	ElementalBaseNodeParam
}

func (r ElementalImageNodeWithTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow ElementalImageNodeWithTypeParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type ElementalMetaNodeWithType struct {
	// Any of "meta".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ElementalBaseNode
}

// Returns the unmodified JSON received from the API
func (r ElementalMetaNodeWithType) RawJSON() string { return r.JSON.raw }
func (r *ElementalMetaNodeWithType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ElementalMetaNodeWithType to a
// ElementalMetaNodeWithTypeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ElementalMetaNodeWithTypeParam.Overrides()
func (r ElementalMetaNodeWithType) ToParam() ElementalMetaNodeWithTypeParam {
	return param.Override[ElementalMetaNodeWithTypeParam](json.RawMessage(r.RawJSON()))
}

type ElementalMetaNodeWithTypeParam struct {
	Type string `json:"type,omitzero"`
	ElementalBaseNodeParam
}

func (r ElementalMetaNodeWithTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow ElementalMetaNodeWithTypeParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// ElementalNodeUnion contains all possible properties and values from
// [ElementalTextNodeWithType], [ElementalMetaNodeWithType],
// [ElementalChannelNodeWithType], [ElementalImageNodeWithType],
// [ElementalActionNodeWithType], [ElementalDividerNodeWithType],
// [ElementalQuoteNodeWithType].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ElementalNodeUnion struct {
	// This field is from variant [ElementalTextNodeWithType],
	// [ElementalMetaNodeWithType], [ElementalChannelNodeWithType],
	// [ElementalImageNodeWithType], [ElementalActionNodeWithType],
	// [ElementalDividerNodeWithType], [ElementalQuoteNodeWithType].
	Channels []string `json:"channels"`
	// This field is from variant [ElementalTextNodeWithType],
	// [ElementalMetaNodeWithType], [ElementalChannelNodeWithType],
	// [ElementalImageNodeWithType], [ElementalActionNodeWithType],
	// [ElementalDividerNodeWithType], [ElementalQuoteNodeWithType].
	If string `json:"if"`
	// This field is from variant [ElementalTextNodeWithType],
	// [ElementalMetaNodeWithType], [ElementalChannelNodeWithType],
	// [ElementalImageNodeWithType], [ElementalActionNodeWithType],
	// [ElementalDividerNodeWithType], [ElementalQuoteNodeWithType].
	Loop string `json:"loop"`
	// This field is from variant [ElementalTextNodeWithType],
	// [ElementalMetaNodeWithType], [ElementalChannelNodeWithType],
	// [ElementalImageNodeWithType], [ElementalActionNodeWithType],
	// [ElementalDividerNodeWithType], [ElementalQuoteNodeWithType].
	Ref  string `json:"ref"`
	Type string `json:"type"`
	// This field is from variant [ElementalChannelNodeWithType].
	Channel string `json:"channel"`
	// This field is from variant [ElementalChannelNodeWithType].
	Raw  map[string]any `json:"raw"`
	JSON struct {
		Channels respjson.Field
		If       respjson.Field
		Loop     respjson.Field
		Ref      respjson.Field
		Type     respjson.Field
		Channel  respjson.Field
		Raw      respjson.Field
		raw      string
	} `json:"-"`
}

func (u ElementalNodeUnion) AsElementalTextNodeWithType() (v ElementalTextNodeWithType) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ElementalNodeUnion) AsElementalMetaNodeWithType() (v ElementalMetaNodeWithType) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ElementalNodeUnion) AsElementalChannelNodeWithType() (v ElementalChannelNodeWithType) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ElementalNodeUnion) AsElementalImageNodeWithType() (v ElementalImageNodeWithType) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ElementalNodeUnion) AsElementalActionNodeWithType() (v ElementalActionNodeWithType) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ElementalNodeUnion) AsElementalDividerNodeWithType() (v ElementalDividerNodeWithType) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ElementalNodeUnion) AsElementalQuoteNodeWithType() (v ElementalQuoteNodeWithType) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ElementalNodeUnion) RawJSON() string { return u.JSON.raw }

func (r *ElementalNodeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ElementalNodeUnion to a ElementalNodeUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ElementalNodeUnionParam.Overrides()
func (r ElementalNodeUnion) ToParam() ElementalNodeUnionParam {
	return param.Override[ElementalNodeUnionParam](json.RawMessage(r.RawJSON()))
}

func ElementalNodeParamOfElementalChannelNodeWithType(channel string) ElementalNodeUnionParam {
	var variant ElementalChannelNodeWithTypeParam
	variant.Channel = channel
	return ElementalNodeUnionParam{OfElementalChannelNodeWithType: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ElementalNodeUnionParam struct {
	OfElementalTextNodeWithType    *ElementalTextNodeWithTypeParam    `json:",omitzero,inline"`
	OfElementalMetaNodeWithType    *ElementalMetaNodeWithTypeParam    `json:",omitzero,inline"`
	OfElementalChannelNodeWithType *ElementalChannelNodeWithTypeParam `json:",omitzero,inline"`
	OfElementalImageNodeWithType   *ElementalImageNodeWithTypeParam   `json:",omitzero,inline"`
	OfElementalActionNodeWithType  *ElementalActionNodeWithTypeParam  `json:",omitzero,inline"`
	OfElementalDividerNodeWithType *ElementalDividerNodeWithTypeParam `json:",omitzero,inline"`
	OfElementalQuoteNodeWithType   *ElementalQuoteNodeWithTypeParam   `json:",omitzero,inline"`
	paramUnion
}

func (u ElementalNodeUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfElementalTextNodeWithType,
		u.OfElementalMetaNodeWithType,
		u.OfElementalChannelNodeWithType,
		u.OfElementalImageNodeWithType,
		u.OfElementalActionNodeWithType,
		u.OfElementalDividerNodeWithType,
		u.OfElementalQuoteNodeWithType)
}
func (u *ElementalNodeUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ElementalNodeUnionParam) asAny() any {
	if !param.IsOmitted(u.OfElementalTextNodeWithType) {
		return u.OfElementalTextNodeWithType
	} else if !param.IsOmitted(u.OfElementalMetaNodeWithType) {
		return u.OfElementalMetaNodeWithType
	} else if !param.IsOmitted(u.OfElementalChannelNodeWithType) {
		return u.OfElementalChannelNodeWithType
	} else if !param.IsOmitted(u.OfElementalImageNodeWithType) {
		return u.OfElementalImageNodeWithType
	} else if !param.IsOmitted(u.OfElementalActionNodeWithType) {
		return u.OfElementalActionNodeWithType
	} else if !param.IsOmitted(u.OfElementalDividerNodeWithType) {
		return u.OfElementalDividerNodeWithType
	} else if !param.IsOmitted(u.OfElementalQuoteNodeWithType) {
		return u.OfElementalQuoteNodeWithType
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetChannel() *string {
	if vt := u.OfElementalChannelNodeWithType; vt != nil {
		return &vt.Channel
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetRaw() map[string]any {
	if vt := u.OfElementalChannelNodeWithType; vt != nil {
		return vt.Raw
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetIf() *string {
	if vt := u.OfElementalTextNodeWithType; vt != nil && vt.If.Valid() {
		return &vt.If.Value
	} else if vt := u.OfElementalMetaNodeWithType; vt != nil && vt.If.Valid() {
		return &vt.If.Value
	} else if vt := u.OfElementalChannelNodeWithType; vt != nil && vt.If.Valid() {
		return &vt.If.Value
	} else if vt := u.OfElementalImageNodeWithType; vt != nil && vt.If.Valid() {
		return &vt.If.Value
	} else if vt := u.OfElementalActionNodeWithType; vt != nil && vt.If.Valid() {
		return &vt.If.Value
	} else if vt := u.OfElementalDividerNodeWithType; vt != nil && vt.If.Valid() {
		return &vt.If.Value
	} else if vt := u.OfElementalQuoteNodeWithType; vt != nil && vt.If.Valid() {
		return &vt.If.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetLoop() *string {
	if vt := u.OfElementalTextNodeWithType; vt != nil && vt.Loop.Valid() {
		return &vt.Loop.Value
	} else if vt := u.OfElementalMetaNodeWithType; vt != nil && vt.Loop.Valid() {
		return &vt.Loop.Value
	} else if vt := u.OfElementalChannelNodeWithType; vt != nil && vt.Loop.Valid() {
		return &vt.Loop.Value
	} else if vt := u.OfElementalImageNodeWithType; vt != nil && vt.Loop.Valid() {
		return &vt.Loop.Value
	} else if vt := u.OfElementalActionNodeWithType; vt != nil && vt.Loop.Valid() {
		return &vt.Loop.Value
	} else if vt := u.OfElementalDividerNodeWithType; vt != nil && vt.Loop.Valid() {
		return &vt.Loop.Value
	} else if vt := u.OfElementalQuoteNodeWithType; vt != nil && vt.Loop.Valid() {
		return &vt.Loop.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetRef() *string {
	if vt := u.OfElementalTextNodeWithType; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	} else if vt := u.OfElementalMetaNodeWithType; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	} else if vt := u.OfElementalChannelNodeWithType; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	} else if vt := u.OfElementalImageNodeWithType; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	} else if vt := u.OfElementalActionNodeWithType; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	} else if vt := u.OfElementalDividerNodeWithType; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	} else if vt := u.OfElementalQuoteNodeWithType; vt != nil && vt.Ref.Valid() {
		return &vt.Ref.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ElementalNodeUnionParam) GetType() *string {
	if vt := u.OfElementalTextNodeWithType; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfElementalMetaNodeWithType; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfElementalChannelNodeWithType; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfElementalImageNodeWithType; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfElementalActionNodeWithType; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfElementalDividerNodeWithType; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfElementalQuoteNodeWithType; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's Channels property, if present.
func (u ElementalNodeUnionParam) GetChannels() []string {
	if vt := u.OfElementalTextNodeWithType; vt != nil {
		return vt.Channels
	} else if vt := u.OfElementalMetaNodeWithType; vt != nil {
		return vt.Channels
	} else if vt := u.OfElementalChannelNodeWithType; vt != nil {
		return vt.Channels
	} else if vt := u.OfElementalImageNodeWithType; vt != nil {
		return vt.Channels
	} else if vt := u.OfElementalActionNodeWithType; vt != nil {
		return vt.Channels
	} else if vt := u.OfElementalDividerNodeWithType; vt != nil {
		return vt.Channels
	} else if vt := u.OfElementalQuoteNodeWithType; vt != nil {
		return vt.Channels
	}
	return nil
}

type ElementalQuoteNodeWithType struct {
	// Any of "quote".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ElementalBaseNode
}

// Returns the unmodified JSON received from the API
func (r ElementalQuoteNodeWithType) RawJSON() string { return r.JSON.raw }
func (r *ElementalQuoteNodeWithType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ElementalQuoteNodeWithType to a
// ElementalQuoteNodeWithTypeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ElementalQuoteNodeWithTypeParam.Overrides()
func (r ElementalQuoteNodeWithType) ToParam() ElementalQuoteNodeWithTypeParam {
	return param.Override[ElementalQuoteNodeWithTypeParam](json.RawMessage(r.RawJSON()))
}

type ElementalQuoteNodeWithTypeParam struct {
	Type string `json:"type,omitzero"`
	ElementalBaseNodeParam
}

func (r ElementalQuoteNodeWithTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow ElementalQuoteNodeWithTypeParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type ElementalTextNodeWithType struct {
	// Any of "text".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ElementalBaseNode
}

// Returns the unmodified JSON received from the API
func (r ElementalTextNodeWithType) RawJSON() string { return r.JSON.raw }
func (r *ElementalTextNodeWithType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ElementalTextNodeWithType to a
// ElementalTextNodeWithTypeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ElementalTextNodeWithTypeParam.Overrides()
func (r ElementalTextNodeWithType) ToParam() ElementalTextNodeWithTypeParam {
	return param.Override[ElementalTextNodeWithTypeParam](json.RawMessage(r.RawJSON()))
}

type ElementalTextNodeWithTypeParam struct {
	Type string `json:"type,omitzero"`
	ElementalBaseNodeParam
}

func (r ElementalTextNodeWithTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow ElementalTextNodeWithTypeParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type EmailFooter struct {
	Content        string `json:"content,nullable"`
	InheritDefault bool   `json:"inheritDefault,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content        respjson.Field
		InheritDefault respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailFooter) RawJSON() string { return r.JSON.raw }
func (r *EmailFooter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EmailFooter to a EmailFooterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EmailFooterParam.Overrides()
func (r EmailFooter) ToParam() EmailFooterParam {
	return param.Override[EmailFooterParam](json.RawMessage(r.RawJSON()))
}

type EmailFooterParam struct {
	Content        param.Opt[string] `json:"content,omitzero"`
	InheritDefault param.Opt[bool]   `json:"inheritDefault,omitzero"`
	paramObj
}

func (r EmailFooterParam) MarshalJSON() (data []byte, err error) {
	type shadow EmailFooterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailFooterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailHead struct {
	InheritDefault bool   `json:"inheritDefault,required"`
	Content        string `json:"content,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InheritDefault respjson.Field
		Content        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailHead) RawJSON() string { return r.JSON.raw }
func (r *EmailHead) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EmailHead to a EmailHeadParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EmailHeadParam.Overrides()
func (r EmailHead) ToParam() EmailHeadParam {
	return param.Override[EmailHeadParam](json.RawMessage(r.RawJSON()))
}

// The property InheritDefault is required.
type EmailHeadParam struct {
	InheritDefault bool              `json:"inheritDefault,required"`
	Content        param.Opt[string] `json:"content,omitzero"`
	paramObj
}

func (r EmailHeadParam) MarshalJSON() (data []byte, err error) {
	type shadow EmailHeadParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailHeadParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailHeader struct {
	Logo           Logo   `json:"logo,required"`
	BarColor       string `json:"barColor,nullable"`
	InheritDefault bool   `json:"inheritDefault,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Logo           respjson.Field
		BarColor       respjson.Field
		InheritDefault respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailHeader) RawJSON() string { return r.JSON.raw }
func (r *EmailHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EmailHeader to a EmailHeaderParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EmailHeaderParam.Overrides()
func (r EmailHeader) ToParam() EmailHeaderParam {
	return param.Override[EmailHeaderParam](json.RawMessage(r.RawJSON()))
}

// The property Logo is required.
type EmailHeaderParam struct {
	Logo           LogoParam         `json:"logo,omitzero,required"`
	BarColor       param.Opt[string] `json:"barColor,omitzero"`
	InheritDefault param.Opt[bool]   `json:"inheritDefault,omitzero"`
	paramObj
}

func (r EmailHeaderParam) MarshalJSON() (data []byte, err error) {
	type shadow EmailHeaderParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailHeaderParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Filter struct {
	// The operator to use for filtering
	//
	// Any of "ENDS_WITH", "EQ", "EXISTS", "GT", "GTE", "INCLUDES", "IS_AFTER",
	// "IS_BEFORE", "LT", "LTE", "NEQ", "OMIT", "STARTS_WITH", "AND", "OR".
	Operator FilterOperator `json:"operator,required"`
	// The attribe name from profile whose value will be operated against the filter
	// value
	Path string `json:"path,required"`
	// The value to use for filtering
	Value string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator    respjson.Field
		Path        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Filter) RawJSON() string { return r.JSON.raw }
func (r *Filter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Filter to a FilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// FilterParam.Overrides()
func (r Filter) ToParam() FilterParam {
	return param.Override[FilterParam](json.RawMessage(r.RawJSON()))
}

// The operator to use for filtering
type FilterOperator string

const (
	FilterOperatorEndsWith   FilterOperator = "ENDS_WITH"
	FilterOperatorEq         FilterOperator = "EQ"
	FilterOperatorExists     FilterOperator = "EXISTS"
	FilterOperatorGt         FilterOperator = "GT"
	FilterOperatorGte        FilterOperator = "GTE"
	FilterOperatorIncludes   FilterOperator = "INCLUDES"
	FilterOperatorIsAfter    FilterOperator = "IS_AFTER"
	FilterOperatorIsBefore   FilterOperator = "IS_BEFORE"
	FilterOperatorLt         FilterOperator = "LT"
	FilterOperatorLte        FilterOperator = "LTE"
	FilterOperatorNeq        FilterOperator = "NEQ"
	FilterOperatorOmit       FilterOperator = "OMIT"
	FilterOperatorStartsWith FilterOperator = "STARTS_WITH"
	FilterOperatorAnd        FilterOperator = "AND"
	FilterOperatorOr         FilterOperator = "OR"
)

// The properties Operator, Path, Value are required.
type FilterParam struct {
	// The operator to use for filtering
	//
	// Any of "ENDS_WITH", "EQ", "EXISTS", "GT", "GTE", "INCLUDES", "IS_AFTER",
	// "IS_BEFORE", "LT", "LTE", "NEQ", "OMIT", "STARTS_WITH", "AND", "OR".
	Operator FilterOperator `json:"operator,omitzero,required"`
	// The attribe name from profile whose value will be operated against the filter
	// value
	Path string `json:"path,required"`
	// The value to use for filtering
	Value string `json:"value,required"`
	paramObj
}

func (r FilterParam) MarshalJSON() (data []byte, err error) {
	type shadow FilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Icons struct {
	Bell    string `json:"bell,nullable"`
	Message string `json:"message,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bell        respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Icons) RawJSON() string { return r.JSON.raw }
func (r *Icons) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Icons to a IconsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// IconsParam.Overrides()
func (r Icons) ToParam() IconsParam {
	return param.Override[IconsParam](json.RawMessage(r.RawJSON()))
}

type IconsParam struct {
	Bell    param.Opt[string] `json:"bell,omitzero"`
	Message param.Opt[string] `json:"message,omitzero"`
	paramObj
}

func (r IconsParam) MarshalJSON() (data []byte, err error) {
	type shadow IconsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IconsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InboundBulkMessageUnion contains all possible properties and values from
// [InboundBulkMessageInboundBulkTemplateMessage],
// [InboundBulkMessageInboundBulkContentMessage].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type InboundBulkMessageUnion struct {
	// This field is from variant [InboundBulkMessageInboundBulkTemplateMessage].
	Template string `json:"template"`
	Brand    string `json:"brand"`
	Data     any    `json:"data"`
	Event    string `json:"event"`
	Locale   any    `json:"locale"`
	Override any    `json:"override"`
	// This field is from variant [InboundBulkMessageInboundBulkContentMessage].
	Content InboundBulkMessageInboundBulkContentMessageContentUnion `json:"content"`
	JSON    struct {
		Template respjson.Field
		Brand    respjson.Field
		Data     respjson.Field
		Event    respjson.Field
		Locale   respjson.Field
		Override respjson.Field
		Content  respjson.Field
		raw      string
	} `json:"-"`
}

func (u InboundBulkMessageUnion) AsInboundBulkTemplateMessage() (v InboundBulkMessageInboundBulkTemplateMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InboundBulkMessageUnion) AsInboundBulkContentMessage() (v InboundBulkMessageInboundBulkContentMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InboundBulkMessageUnion) RawJSON() string { return u.JSON.raw }

func (r *InboundBulkMessageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this InboundBulkMessageUnion to a InboundBulkMessageUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// InboundBulkMessageUnionParam.Overrides()
func (r InboundBulkMessageUnion) ToParam() InboundBulkMessageUnionParam {
	return param.Override[InboundBulkMessageUnionParam](json.RawMessage(r.RawJSON()))
}

type InboundBulkMessageInboundBulkTemplateMessage struct {
	Template string                    `json:"template,required"`
	Brand    string                    `json:"brand,nullable"`
	Data     map[string]any            `json:"data,nullable"`
	Event    string                    `json:"event,nullable"`
	Locale   map[string]map[string]any `json:"locale,nullable"`
	Override map[string]any            `json:"override,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Template    respjson.Field
		Brand       respjson.Field
		Data        respjson.Field
		Event       respjson.Field
		Locale      respjson.Field
		Override    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboundBulkMessageInboundBulkTemplateMessage) RawJSON() string { return r.JSON.raw }
func (r *InboundBulkMessageInboundBulkTemplateMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboundBulkMessageInboundBulkContentMessage struct {
	// Syntactic sugar to provide a fast shorthand for Courier Elemental Blocks.
	Content  InboundBulkMessageInboundBulkContentMessageContentUnion `json:"content,required"`
	Brand    string                                                  `json:"brand,nullable"`
	Data     map[string]any                                          `json:"data,nullable"`
	Event    string                                                  `json:"event,nullable"`
	Locale   map[string]map[string]any                               `json:"locale,nullable"`
	Override map[string]any                                          `json:"override,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Brand       respjson.Field
		Data        respjson.Field
		Event       respjson.Field
		Locale      respjson.Field
		Override    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboundBulkMessageInboundBulkContentMessage) RawJSON() string { return r.JSON.raw }
func (r *InboundBulkMessageInboundBulkContentMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InboundBulkMessageInboundBulkContentMessageContentUnion contains all possible
// properties and values from [ElementalContentSugar], [ElementalContent].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type InboundBulkMessageInboundBulkContentMessageContentUnion struct {
	// This field is from variant [ElementalContentSugar].
	Body string `json:"body"`
	// This field is from variant [ElementalContentSugar].
	Title string `json:"title"`
	// This field is from variant [ElementalContent].
	Elements []ElementalNodeUnion `json:"elements"`
	// This field is from variant [ElementalContent].
	Version string `json:"version"`
	// This field is from variant [ElementalContent].
	Brand string `json:"brand"`
	JSON  struct {
		Body     respjson.Field
		Title    respjson.Field
		Elements respjson.Field
		Version  respjson.Field
		Brand    respjson.Field
		raw      string
	} `json:"-"`
}

func (u InboundBulkMessageInboundBulkContentMessageContentUnion) AsElementalContentSugar() (v ElementalContentSugar) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InboundBulkMessageInboundBulkContentMessageContentUnion) AsElementalContent() (v ElementalContent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InboundBulkMessageInboundBulkContentMessageContentUnion) RawJSON() string { return u.JSON.raw }

func (r *InboundBulkMessageInboundBulkContentMessageContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func InboundBulkMessageParamOfInboundBulkTemplateMessage(template string) InboundBulkMessageUnionParam {
	var variant InboundBulkMessageInboundBulkTemplateMessageParam
	variant.Template = template
	return InboundBulkMessageUnionParam{OfInboundBulkTemplateMessage: &variant}
}

func InboundBulkMessageParamOfInboundBulkContentMessage[
	T ElementalContentSugarParam | ElementalContentParam,
](content T) InboundBulkMessageUnionParam {
	var variant InboundBulkMessageInboundBulkContentMessageParam
	switch v := any(content).(type) {
	case ElementalContentSugarParam:
		variant.Content.OfElementalContentSugar = &v
	case ElementalContentParam:
		variant.Content.OfElementalContent = &v
	}
	return InboundBulkMessageUnionParam{OfInboundBulkContentMessage: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InboundBulkMessageUnionParam struct {
	OfInboundBulkTemplateMessage *InboundBulkMessageInboundBulkTemplateMessageParam `json:",omitzero,inline"`
	OfInboundBulkContentMessage  *InboundBulkMessageInboundBulkContentMessageParam  `json:",omitzero,inline"`
	paramUnion
}

func (u InboundBulkMessageUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInboundBulkTemplateMessage, u.OfInboundBulkContentMessage)
}
func (u *InboundBulkMessageUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *InboundBulkMessageUnionParam) asAny() any {
	if !param.IsOmitted(u.OfInboundBulkTemplateMessage) {
		return u.OfInboundBulkTemplateMessage
	} else if !param.IsOmitted(u.OfInboundBulkContentMessage) {
		return u.OfInboundBulkContentMessage
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InboundBulkMessageUnionParam) GetTemplate() *string {
	if vt := u.OfInboundBulkTemplateMessage; vt != nil {
		return &vt.Template
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InboundBulkMessageUnionParam) GetContent() *InboundBulkMessageInboundBulkContentMessageContentUnionParam {
	if vt := u.OfInboundBulkContentMessage; vt != nil {
		return &vt.Content
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InboundBulkMessageUnionParam) GetBrand() *string {
	if vt := u.OfInboundBulkTemplateMessage; vt != nil && vt.Brand.Valid() {
		return &vt.Brand.Value
	} else if vt := u.OfInboundBulkContentMessage; vt != nil && vt.Brand.Valid() {
		return &vt.Brand.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InboundBulkMessageUnionParam) GetEvent() *string {
	if vt := u.OfInboundBulkTemplateMessage; vt != nil && vt.Event.Valid() {
		return &vt.Event.Value
	} else if vt := u.OfInboundBulkContentMessage; vt != nil && vt.Event.Valid() {
		return &vt.Event.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's Data property, if present.
func (u InboundBulkMessageUnionParam) GetData() map[string]any {
	if vt := u.OfInboundBulkTemplateMessage; vt != nil {
		return vt.Data
	} else if vt := u.OfInboundBulkContentMessage; vt != nil {
		return vt.Data
	}
	return nil
}

// Returns a pointer to the underlying variant's Locale property, if present.
func (u InboundBulkMessageUnionParam) GetLocale() map[string]map[string]any {
	if vt := u.OfInboundBulkTemplateMessage; vt != nil {
		return vt.Locale
	} else if vt := u.OfInboundBulkContentMessage; vt != nil {
		return vt.Locale
	}
	return nil
}

// Returns a pointer to the underlying variant's Override property, if present.
func (u InboundBulkMessageUnionParam) GetOverride() map[string]any {
	if vt := u.OfInboundBulkTemplateMessage; vt != nil {
		return vt.Override
	} else if vt := u.OfInboundBulkContentMessage; vt != nil {
		return vt.Override
	}
	return nil
}

// The property Template is required.
type InboundBulkMessageInboundBulkTemplateMessageParam struct {
	Template string                    `json:"template,required"`
	Brand    param.Opt[string]         `json:"brand,omitzero"`
	Event    param.Opt[string]         `json:"event,omitzero"`
	Data     map[string]any            `json:"data,omitzero"`
	Locale   map[string]map[string]any `json:"locale,omitzero"`
	Override map[string]any            `json:"override,omitzero"`
	paramObj
}

func (r InboundBulkMessageInboundBulkTemplateMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow InboundBulkMessageInboundBulkTemplateMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InboundBulkMessageInboundBulkTemplateMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Content is required.
type InboundBulkMessageInboundBulkContentMessageParam struct {
	// Syntactic sugar to provide a fast shorthand for Courier Elemental Blocks.
	Content  InboundBulkMessageInboundBulkContentMessageContentUnionParam `json:"content,omitzero,required"`
	Brand    param.Opt[string]                                            `json:"brand,omitzero"`
	Event    param.Opt[string]                                            `json:"event,omitzero"`
	Data     map[string]any                                               `json:"data,omitzero"`
	Locale   map[string]map[string]any                                    `json:"locale,omitzero"`
	Override map[string]any                                               `json:"override,omitzero"`
	paramObj
}

func (r InboundBulkMessageInboundBulkContentMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow InboundBulkMessageInboundBulkContentMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InboundBulkMessageInboundBulkContentMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InboundBulkMessageInboundBulkContentMessageContentUnionParam struct {
	OfElementalContentSugar *ElementalContentSugarParam `json:",omitzero,inline"`
	OfElementalContent      *ElementalContentParam      `json:",omitzero,inline"`
	paramUnion
}

func (u InboundBulkMessageInboundBulkContentMessageContentUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfElementalContentSugar, u.OfElementalContent)
}
func (u *InboundBulkMessageInboundBulkContentMessageContentUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *InboundBulkMessageInboundBulkContentMessageContentUnionParam) asAny() any {
	if !param.IsOmitted(u.OfElementalContentSugar) {
		return u.OfElementalContentSugar
	} else if !param.IsOmitted(u.OfElementalContent) {
		return u.OfElementalContent
	}
	return nil
}

type InboundBulkMessageUser struct {
	Data        any                  `json:"data"`
	Preferences RecipientPreferences `json:"preferences,nullable"`
	Profile     any                  `json:"profile"`
	Recipient   string               `json:"recipient,nullable"`
	To          UserRecipient        `json:"to,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Preferences respjson.Field
		Profile     respjson.Field
		Recipient   respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboundBulkMessageUser) RawJSON() string { return r.JSON.raw }
func (r *InboundBulkMessageUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this InboundBulkMessageUser to a InboundBulkMessageUserParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// InboundBulkMessageUserParam.Overrides()
func (r InboundBulkMessageUser) ToParam() InboundBulkMessageUserParam {
	return param.Override[InboundBulkMessageUserParam](json.RawMessage(r.RawJSON()))
}

type InboundBulkMessageUserParam struct {
	Recipient   param.Opt[string]         `json:"recipient,omitzero"`
	Data        any                       `json:"data,omitzero"`
	Preferences RecipientPreferencesParam `json:"preferences,omitzero"`
	Profile     any                       `json:"profile,omitzero"`
	To          UserRecipientParam        `json:"to,omitzero"`
	paramObj
}

func (r InboundBulkMessageUserParam) MarshalJSON() (data []byte, err error) {
	type shadow InboundBulkMessageUserParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InboundBulkMessageUserParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Logo struct {
	Href  string `json:"href,nullable"`
	Image string `json:"image,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Href        respjson.Field
		Image       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Logo) RawJSON() string { return r.JSON.raw }
func (r *Logo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Logo to a LogoParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// LogoParam.Overrides()
func (r Logo) ToParam() LogoParam {
	return param.Override[LogoParam](json.RawMessage(r.RawJSON()))
}

type LogoParam struct {
	Href  param.Opt[string] `json:"href,omitzero"`
	Image param.Opt[string] `json:"image,omitzero"`
	paramObj
}

func (r LogoParam) MarshalJSON() (data []byte, err error) {
	type shadow LogoParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogoParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageContext struct {
	// Tenant id used to load brand/default preferences/context.
	TenantID string `json:"tenant_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TenantID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageContext) RawJSON() string { return r.JSON.raw }
func (r *MessageContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MessageContext to a MessageContextParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MessageContextParam.Overrides()
func (r MessageContext) ToParam() MessageContextParam {
	return param.Override[MessageContextParam](json.RawMessage(r.RawJSON()))
}

type MessageContextParam struct {
	// Tenant id used to load brand/default preferences/context.
	TenantID param.Opt[string] `json:"tenant_id,omitzero"`
	paramObj
}

func (r MessageContextParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageContextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageContextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageDetails struct {
	// A unique identifier associated with the message you wish to retrieve (results
	// from a send).
	ID string `json:"id,required"`
	// A UTC timestamp at which the recipient clicked on a tracked link for the first
	// time. Stored as a millisecond representation of the Unix epoch.
	Clicked int64 `json:"clicked,required"`
	// A UTC timestamp at which the Integration provider delivered the message. Stored
	// as a millisecond representation of the Unix epoch.
	Delivered int64 `json:"delivered,required"`
	// A UTC timestamp at which Courier received the message request. Stored as a
	// millisecond representation of the Unix epoch.
	Enqueued int64 `json:"enqueued,required"`
	// A unique identifier associated with the event of the delivered message.
	Event string `json:"event,required"`
	// A unique identifier associated with the notification of the delivered message.
	Notification string `json:"notification,required"`
	// A UTC timestamp at which the recipient opened a message for the first time.
	// Stored as a millisecond representation of the Unix epoch.
	Opened int64 `json:"opened,required"`
	// A unique identifier associated with the recipient of the delivered message.
	Recipient string `json:"recipient,required"`
	// A UTC timestamp at which Courier passed the message to the Integration provider.
	// Stored as a millisecond representation of the Unix epoch.
	Sent int64 `json:"sent,required"`
	// The current status of the message.
	//
	// Any of "CANCELED", "CLICKED", "DELAYED", "DELIVERED", "DIGESTED", "ENQUEUED",
	// "FILTERED", "OPENED", "ROUTED", "SENT", "SIMULATED", "THROTTLED",
	// "UNDELIVERABLE", "UNMAPPED", "UNROUTABLE".
	Status MessageDetailsStatus `json:"status,required"`
	// A message describing the error that occurred.
	Error string `json:"error,nullable"`
	// The reason for the current status of the message.
	//
	// Any of "BOUNCED", "FAILED", "FILTERED", "NO_CHANNELS", "NO_PROVIDERS",
	// "OPT_IN_REQUIRED", "PROVIDER_ERROR", "UNPUBLISHED", "UNSUBSCRIBED".
	Reason MessageDetailsReason `json:"reason,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Clicked      respjson.Field
		Delivered    respjson.Field
		Enqueued     respjson.Field
		Event        respjson.Field
		Notification respjson.Field
		Opened       respjson.Field
		Recipient    respjson.Field
		Sent         respjson.Field
		Status       respjson.Field
		Error        respjson.Field
		Reason       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageDetails) RawJSON() string { return r.JSON.raw }
func (r *MessageDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current status of the message.
type MessageDetailsStatus string

const (
	MessageDetailsStatusCanceled      MessageDetailsStatus = "CANCELED"
	MessageDetailsStatusClicked       MessageDetailsStatus = "CLICKED"
	MessageDetailsStatusDelayed       MessageDetailsStatus = "DELAYED"
	MessageDetailsStatusDelivered     MessageDetailsStatus = "DELIVERED"
	MessageDetailsStatusDigested      MessageDetailsStatus = "DIGESTED"
	MessageDetailsStatusEnqueued      MessageDetailsStatus = "ENQUEUED"
	MessageDetailsStatusFiltered      MessageDetailsStatus = "FILTERED"
	MessageDetailsStatusOpened        MessageDetailsStatus = "OPENED"
	MessageDetailsStatusRouted        MessageDetailsStatus = "ROUTED"
	MessageDetailsStatusSent          MessageDetailsStatus = "SENT"
	MessageDetailsStatusSimulated     MessageDetailsStatus = "SIMULATED"
	MessageDetailsStatusThrottled     MessageDetailsStatus = "THROTTLED"
	MessageDetailsStatusUndeliverable MessageDetailsStatus = "UNDELIVERABLE"
	MessageDetailsStatusUnmapped      MessageDetailsStatus = "UNMAPPED"
	MessageDetailsStatusUnroutable    MessageDetailsStatus = "UNROUTABLE"
)

// The reason for the current status of the message.
type MessageDetailsReason string

const (
	MessageDetailsReasonBounced       MessageDetailsReason = "BOUNCED"
	MessageDetailsReasonFailed        MessageDetailsReason = "FAILED"
	MessageDetailsReasonFiltered      MessageDetailsReason = "FILTERED"
	MessageDetailsReasonNoChannels    MessageDetailsReason = "NO_CHANNELS"
	MessageDetailsReasonNoProviders   MessageDetailsReason = "NO_PROVIDERS"
	MessageDetailsReasonOptInRequired MessageDetailsReason = "OPT_IN_REQUIRED"
	MessageDetailsReasonProviderError MessageDetailsReason = "PROVIDER_ERROR"
	MessageDetailsReasonUnpublished   MessageDetailsReason = "UNPUBLISHED"
	MessageDetailsReasonUnsubscribed  MessageDetailsReason = "UNSUBSCRIBED"
)

type MessageRouting struct {
	Channels []MessageRoutingChannelUnion `json:"channels,required"`
	// Any of "all", "single".
	Method MessageRoutingMethod `json:"method,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Channels    respjson.Field
		Method      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageRouting) RawJSON() string { return r.JSON.raw }
func (r *MessageRouting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MessageRouting to a MessageRoutingParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MessageRoutingParam.Overrides()
func (r MessageRouting) ToParam() MessageRoutingParam {
	return param.Override[MessageRoutingParam](json.RawMessage(r.RawJSON()))
}

type MessageRoutingMethod string

const (
	MessageRoutingMethodAll    MessageRoutingMethod = "all"
	MessageRoutingMethodSingle MessageRoutingMethod = "single"
)

// The properties Channels, Method are required.
type MessageRoutingParam struct {
	Channels []MessageRoutingChannelUnionParam `json:"channels,omitzero,required"`
	// Any of "all", "single".
	Method MessageRoutingMethod `json:"method,omitzero,required"`
	paramObj
}

func (r MessageRoutingParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageRoutingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageRoutingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MessageRoutingChannelUnion contains all possible properties and values from
// [string], [MessageRouting].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type MessageRoutingChannelUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is from variant [MessageRouting].
	Channels []MessageRoutingChannelUnion `json:"channels"`
	// This field is from variant [MessageRouting].
	Method MessageRoutingMethod `json:"method"`
	JSON   struct {
		OfString respjson.Field
		Channels respjson.Field
		Method   respjson.Field
		raw      string
	} `json:"-"`
}

func (u MessageRoutingChannelUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageRoutingChannelUnion) AsMessageRouting() (v MessageRouting) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MessageRoutingChannelUnion) RawJSON() string { return u.JSON.raw }

func (r *MessageRoutingChannelUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MessageRoutingChannelUnion to a
// MessageRoutingChannelUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MessageRoutingChannelUnionParam.Overrides()
func (r MessageRoutingChannelUnion) ToParam() MessageRoutingChannelUnionParam {
	return param.Override[MessageRoutingChannelUnionParam](json.RawMessage(r.RawJSON()))
}

func MessageRoutingChannelParamOfMessageRouting(channels []MessageRoutingChannelUnionParam, method MessageRoutingMethod) MessageRoutingChannelUnionParam {
	var variant MessageRoutingParam
	variant.Channels = channels
	variant.Method = method
	return MessageRoutingChannelUnionParam{OfMessageRouting: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MessageRoutingChannelUnionParam struct {
	OfString         param.Opt[string]    `json:",omitzero,inline"`
	OfMessageRouting *MessageRoutingParam `json:",omitzero,inline"`
	paramUnion
}

func (u MessageRoutingChannelUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfMessageRouting)
}
func (u *MessageRoutingChannelUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *MessageRoutingChannelUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfMessageRouting) {
		return u.OfMessageRouting
	}
	return nil
}

type NotificationGetContent struct {
	Blocks   []NotificationGetContentBlock   `json:"blocks,nullable"`
	Channels []NotificationGetContentChannel `json:"channels,nullable"`
	Checksum string                          `json:"checksum,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Blocks      respjson.Field
		Channels    respjson.Field
		Checksum    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationGetContent) RawJSON() string { return r.JSON.raw }
func (r *NotificationGetContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationGetContentBlock struct {
	ID string `json:"id,required"`
	// Any of "action", "divider", "image", "jsonnet", "list", "markdown", "quote",
	// "template", "text".
	Type     string                                            `json:"type,required"`
	Alias    string                                            `json:"alias,nullable"`
	Checksum string                                            `json:"checksum,nullable"`
	Content  NotificationGetContentBlockContentUnion           `json:"content,nullable"`
	Context  string                                            `json:"context,nullable"`
	Locales  map[string]NotificationGetContentBlockLocaleUnion `json:"locales,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		Alias       respjson.Field
		Checksum    respjson.Field
		Content     respjson.Field
		Context     respjson.Field
		Locales     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationGetContentBlock) RawJSON() string { return r.JSON.raw }
func (r *NotificationGetContentBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// NotificationGetContentBlockContentUnion contains all possible properties and
// values from [string],
// [NotificationGetContentBlockContentNotificationContentHierarchy].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type NotificationGetContentBlockContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is from variant
	// [NotificationGetContentBlockContentNotificationContentHierarchy].
	Children string `json:"children"`
	// This field is from variant
	// [NotificationGetContentBlockContentNotificationContentHierarchy].
	Parent string `json:"parent"`
	JSON   struct {
		OfString respjson.Field
		Children respjson.Field
		Parent   respjson.Field
		raw      string
	} `json:"-"`
}

func (u NotificationGetContentBlockContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NotificationGetContentBlockContentUnion) AsNotificationContentHierarchy() (v NotificationGetContentBlockContentNotificationContentHierarchy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u NotificationGetContentBlockContentUnion) RawJSON() string { return u.JSON.raw }

func (r *NotificationGetContentBlockContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationGetContentBlockContentNotificationContentHierarchy struct {
	Children string `json:"children,nullable"`
	Parent   string `json:"parent,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Children    respjson.Field
		Parent      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationGetContentBlockContentNotificationContentHierarchy) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationGetContentBlockContentNotificationContentHierarchy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// NotificationGetContentBlockLocaleUnion contains all possible properties and
// values from [string],
// [NotificationGetContentBlockLocaleNotificationContentHierarchy].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type NotificationGetContentBlockLocaleUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is from variant
	// [NotificationGetContentBlockLocaleNotificationContentHierarchy].
	Children string `json:"children"`
	// This field is from variant
	// [NotificationGetContentBlockLocaleNotificationContentHierarchy].
	Parent string `json:"parent"`
	JSON   struct {
		OfString respjson.Field
		Children respjson.Field
		Parent   respjson.Field
		raw      string
	} `json:"-"`
}

func (u NotificationGetContentBlockLocaleUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NotificationGetContentBlockLocaleUnion) AsNotificationContentHierarchy() (v NotificationGetContentBlockLocaleNotificationContentHierarchy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u NotificationGetContentBlockLocaleUnion) RawJSON() string { return u.JSON.raw }

func (r *NotificationGetContentBlockLocaleUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationGetContentBlockLocaleNotificationContentHierarchy struct {
	Children string `json:"children,nullable"`
	Parent   string `json:"parent,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Children    respjson.Field
		Parent      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationGetContentBlockLocaleNotificationContentHierarchy) RawJSON() string {
	return r.JSON.raw
}
func (r *NotificationGetContentBlockLocaleNotificationContentHierarchy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationGetContentChannel struct {
	ID       string                                         `json:"id,required"`
	Checksum string                                         `json:"checksum,nullable"`
	Content  NotificationGetContentChannelContent           `json:"content,nullable"`
	Locales  map[string]NotificationGetContentChannelLocale `json:"locales,nullable"`
	Type     string                                         `json:"type,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Checksum    respjson.Field
		Content     respjson.Field
		Locales     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationGetContentChannel) RawJSON() string { return r.JSON.raw }
func (r *NotificationGetContentChannel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationGetContentChannelContent struct {
	Subject string `json:"subject,nullable"`
	Title   string `json:"title,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Subject     respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationGetContentChannelContent) RawJSON() string { return r.JSON.raw }
func (r *NotificationGetContentChannelContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationGetContentChannelLocale struct {
	Subject string `json:"subject,nullable"`
	Title   string `json:"title,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Subject     respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationGetContentChannelLocale) RawJSON() string { return r.JSON.raw }
func (r *NotificationGetContentChannelLocale) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationPreferenceDetails struct {
	// Any of "OPTED_IN", "OPTED_OUT", "REQUIRED".
	Status             PreferenceStatus    `json:"status,required"`
	ChannelPreferences []ChannelPreference `json:"channel_preferences,nullable"`
	Rules              []Rule              `json:"rules,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status             respjson.Field
		ChannelPreferences respjson.Field
		Rules              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationPreferenceDetails) RawJSON() string { return r.JSON.raw }
func (r *NotificationPreferenceDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this NotificationPreferenceDetails to a
// NotificationPreferenceDetailsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// NotificationPreferenceDetailsParam.Overrides()
func (r NotificationPreferenceDetails) ToParam() NotificationPreferenceDetailsParam {
	return param.Override[NotificationPreferenceDetailsParam](json.RawMessage(r.RawJSON()))
}

// The property Status is required.
type NotificationPreferenceDetailsParam struct {
	// Any of "OPTED_IN", "OPTED_OUT", "REQUIRED".
	Status             PreferenceStatus         `json:"status,omitzero,required"`
	ChannelPreferences []ChannelPreferenceParam `json:"channel_preferences,omitzero"`
	Rules              []RuleParam              `json:"rules,omitzero"`
	paramObj
}

func (r NotificationPreferenceDetailsParam) MarshalJSON() (data []byte, err error) {
	type shadow NotificationPreferenceDetailsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NotificationPreferenceDetailsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Paging struct {
	More   bool   `json:"more,required"`
	Cursor string `json:"cursor,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		More        respjson.Field
		Cursor      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Paging) RawJSON() string { return r.JSON.raw }
func (r *Paging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Preference struct {
	// Any of "OPTED_IN", "OPTED_OUT", "REQUIRED".
	Status             PreferenceStatus    `json:"status,required"`
	ChannelPreferences []ChannelPreference `json:"channel_preferences,nullable"`
	Rules              []Rule              `json:"rules,nullable"`
	// Any of "subscription", "list", "recipient".
	Source PreferenceSource `json:"source,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status             respjson.Field
		ChannelPreferences respjson.Field
		Rules              respjson.Field
		Source             respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Preference) RawJSON() string { return r.JSON.raw }
func (r *Preference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Preference to a PreferenceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PreferenceParam.Overrides()
func (r Preference) ToParam() PreferenceParam {
	return param.Override[PreferenceParam](json.RawMessage(r.RawJSON()))
}

type PreferenceSource string

const (
	PreferenceSourceSubscription PreferenceSource = "subscription"
	PreferenceSourceList         PreferenceSource = "list"
	PreferenceSourceRecipient    PreferenceSource = "recipient"
)

// The property Status is required.
type PreferenceParam struct {
	// Any of "OPTED_IN", "OPTED_OUT", "REQUIRED".
	Status             PreferenceStatus         `json:"status,omitzero,required"`
	ChannelPreferences []ChannelPreferenceParam `json:"channel_preferences,omitzero"`
	Rules              []RuleParam              `json:"rules,omitzero"`
	// Any of "subscription", "list", "recipient".
	Source PreferenceSource `json:"source,omitzero"`
	paramObj
}

func (r PreferenceParam) MarshalJSON() (data []byte, err error) {
	type shadow PreferenceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PreferenceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PreferenceStatus string

const (
	PreferenceStatusOptedIn  PreferenceStatus = "OPTED_IN"
	PreferenceStatusOptedOut PreferenceStatus = "OPTED_OUT"
	PreferenceStatusRequired PreferenceStatus = "REQUIRED"
)

// The property RecipientID is required.
type PutSubscriptionsRecipientParam struct {
	RecipientID string                    `json:"recipientId,required"`
	Preferences RecipientPreferencesParam `json:"preferences,omitzero"`
	paramObj
}

func (r PutSubscriptionsRecipientParam) MarshalJSON() (data []byte, err error) {
	type shadow PutSubscriptionsRecipientParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PutSubscriptionsRecipientParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecipientParam struct {
	// Use `tenant_id` instead.
	AccountID param.Opt[string] `json:"account_id,omitzero"`
	// The user's email address.
	Email param.Opt[string] `json:"email,omitzero"`
	// The user's preferred ISO 639-1 language code.
	Locale param.Opt[string] `json:"locale,omitzero"`
	// The user's phone number.
	PhoneNumber param.Opt[string] `json:"phone_number,omitzero"`
	// The id of the tenant the user is associated with.
	TenantID param.Opt[string] `json:"tenant_id,omitzero"`
	// The user's unique identifier. Typically, this will match the user id of a user
	// in your system.
	UserID      param.Opt[string]         `json:"user_id,omitzero"`
	Data        map[string]any            `json:"data,omitzero"`
	Preferences RecipientPreferencesParam `json:"preferences,omitzero"`
	// Context such as tenant_id to send the notification with.
	Context MessageContextParam `json:"context,omitzero"`
	paramObj
}

func (r RecipientParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecipientParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Notifications is required.
type RecipientPreferencesParam struct {
	Notifications map[string]PreferenceParam `json:"notifications,omitzero,required"`
	TemplateID    param.Opt[string]          `json:"templateId,omitzero"`
	Categories    map[string]PreferenceParam `json:"categories,omitzero"`
	paramObj
}

func (r RecipientPreferencesParam) MarshalJSON() (data []byte, err error) {
	type shadow RecipientPreferencesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecipientPreferencesParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecipientPreferences struct {
	Categories    map[string]NotificationPreferenceDetails `json:"categories,nullable"`
	Notifications map[string]NotificationPreferenceDetails `json:"notifications,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Categories    respjson.Field
		Notifications respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecipientPreferences) RawJSON() string { return r.JSON.raw }
func (r *RecipientPreferences) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this RecipientPreferences to a RecipientPreferencesParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RecipientPreferencesParam.Overrides()
func (r RecipientPreferences) ToParam() RecipientPreferencesParam {
	return param.Override[RecipientPreferencesParam](json.RawMessage(r.RawJSON()))
}

type Rule struct {
	Until string `json:"until,required"`
	Start string `json:"start,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Until       respjson.Field
		Start       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Rule) RawJSON() string { return r.JSON.raw }
func (r *Rule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Rule to a RuleParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RuleParam.Overrides()
func (r Rule) ToParam() RuleParam {
	return param.Override[RuleParam](json.RawMessage(r.RawJSON()))
}

// The property Until is required.
type RuleParam struct {
	Until string            `json:"until,required"`
	Start param.Opt[string] `json:"start,omitzero"`
	paramObj
}

func (r RuleParam) MarshalJSON() (data []byte, err error) {
	type shadow RuleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionList struct {
	ID      string `json:"id,required"`
	Name    string `json:"name,required"`
	Created string `json:"created,nullable"`
	Updated string `json:"updated,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Created     respjson.Field
		Updated     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionList) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionTopicNew struct {
	// Any of "OPTED_OUT", "OPTED_IN", "REQUIRED".
	Status SubscriptionTopicNewStatus `json:"status,required"`
	// The default channels to send to this tenant when has_custom_routing is enabled
	CustomRouting []ChannelClassification `json:"custom_routing,nullable"`
	// Override channel routing with custom preferences. This will override any
	// template prefernces that are set, but a user can still customize their
	// preferences
	HasCustomRouting bool `json:"has_custom_routing,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status           respjson.Field
		CustomRouting    respjson.Field
		HasCustomRouting respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionTopicNew) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionTopicNew) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SubscriptionTopicNew to a SubscriptionTopicNewParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SubscriptionTopicNewParam.Overrides()
func (r SubscriptionTopicNew) ToParam() SubscriptionTopicNewParam {
	return param.Override[SubscriptionTopicNewParam](json.RawMessage(r.RawJSON()))
}

type SubscriptionTopicNewStatus string

const (
	SubscriptionTopicNewStatusOptedOut SubscriptionTopicNewStatus = "OPTED_OUT"
	SubscriptionTopicNewStatusOptedIn  SubscriptionTopicNewStatus = "OPTED_IN"
	SubscriptionTopicNewStatusRequired SubscriptionTopicNewStatus = "REQUIRED"
)

// The property Status is required.
type SubscriptionTopicNewParam struct {
	// Any of "OPTED_OUT", "OPTED_IN", "REQUIRED".
	Status SubscriptionTopicNewStatus `json:"status,omitzero,required"`
	// Override channel routing with custom preferences. This will override any
	// template prefernces that are set, but a user can still customize their
	// preferences
	HasCustomRouting param.Opt[bool] `json:"has_custom_routing,omitzero"`
	// The default channels to send to this tenant when has_custom_routing is enabled
	CustomRouting []ChannelClassification `json:"custom_routing,omitzero"`
	paramObj
}

func (r SubscriptionTopicNewParam) MarshalJSON() (data []byte, err error) {
	type shadow SubscriptionTopicNewParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubscriptionTopicNewParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Tenant struct {
	// Id of the tenant.
	ID string `json:"id,required"`
	// Name of the tenant.
	Name string `json:"name,required"`
	// Brand to be used for the account when one is not specified by the send call.
	BrandID string `json:"brand_id,nullable"`
	// Defines the preferences used for the account when the user hasn't specified
	// their own.
	DefaultPreferences DefaultPreferences `json:"default_preferences,nullable"`
	// Tenant's parent id (if any).
	ParentTenantID string `json:"parent_tenant_id,nullable"`
	// Arbitrary properties accessible to a template.
	Properties map[string]any `json:"properties,nullable"`
	// A user profile object merged with user profile on send.
	UserProfile map[string]any `json:"user_profile,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Name               respjson.Field
		BrandID            respjson.Field
		DefaultPreferences respjson.Field
		ParentTenantID     respjson.Field
		Properties         respjson.Field
		UserProfile        respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Tenant) RawJSON() string { return r.JSON.raw }
func (r *Tenant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantAssociation struct {
	// Tenant ID for the association between tenant and user
	TenantID string `json:"tenant_id,required"`
	// Additional metadata to be applied to a user profile when used in a tenant
	// context
	Profile map[string]any `json:"profile,nullable"`
	// Any of "user".
	Type TenantAssociationType `json:"type,nullable"`
	// User ID for the association between tenant and user
	UserID string `json:"user_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TenantID    respjson.Field
		Profile     respjson.Field
		Type        respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenantAssociation) RawJSON() string { return r.JSON.raw }
func (r *TenantAssociation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TenantAssociation to a TenantAssociationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TenantAssociationParam.Overrides()
func (r TenantAssociation) ToParam() TenantAssociationParam {
	return param.Override[TenantAssociationParam](json.RawMessage(r.RawJSON()))
}

type TenantAssociationType string

const (
	TenantAssociationTypeUser TenantAssociationType = "user"
)

// The property TenantID is required.
type TenantAssociationParam struct {
	// Tenant ID for the association between tenant and user
	TenantID string `json:"tenant_id,required"`
	// User ID for the association between tenant and user
	UserID param.Opt[string] `json:"user_id,omitzero"`
	// Additional metadata to be applied to a user profile when used in a tenant
	// context
	Profile map[string]any `json:"profile,omitzero"`
	// Any of "user".
	Type TenantAssociationType `json:"type,omitzero"`
	paramObj
}

func (r TenantAssociationParam) MarshalJSON() (data []byte, err error) {
	type shadow TenantAssociationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TenantAssociationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TopicPreference struct {
	// Any of "OPTED_IN", "OPTED_OUT", "REQUIRED".
	DefaultStatus PreferenceStatus `json:"default_status,required"`
	// Any of "OPTED_IN", "OPTED_OUT", "REQUIRED".
	Status    PreferenceStatus `json:"status,required"`
	TopicID   string           `json:"topic_id,required"`
	TopicName string           `json:"topic_name,required"`
	// The Channels a user has chosen to receive notifications through for this topic
	CustomRouting    []ChannelClassification `json:"custom_routing,nullable"`
	HasCustomRouting bool                    `json:"has_custom_routing,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DefaultStatus    respjson.Field
		Status           respjson.Field
		TopicID          respjson.Field
		TopicName        respjson.Field
		CustomRouting    respjson.Field
		HasCustomRouting respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TopicPreference) RawJSON() string { return r.JSON.raw }
func (r *TopicPreference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserRecipient struct {
	// Use `tenant_id` instead.
	AccountID string `json:"account_id,nullable"`
	// Context such as tenant_id to send the notification with.
	Context MessageContext `json:"context,nullable"`
	Data    map[string]any `json:"data,nullable"`
	// The user's email address.
	Email string `json:"email,nullable"`
	// The user's preferred ISO 639-1 language code.
	Locale string `json:"locale,nullable"`
	// The user's phone number.
	PhoneNumber string                   `json:"phone_number,nullable"`
	Preferences UserRecipientPreferences `json:"preferences,nullable"`
	// The id of the tenant the user is associated with.
	TenantID string `json:"tenant_id,nullable"`
	// The user's unique identifier. Typically, this will match the user id of a user
	// in your system.
	UserID string `json:"user_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountID   respjson.Field
		Context     respjson.Field
		Data        respjson.Field
		Email       respjson.Field
		Locale      respjson.Field
		PhoneNumber respjson.Field
		Preferences respjson.Field
		TenantID    respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserRecipient) RawJSON() string { return r.JSON.raw }
func (r *UserRecipient) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this UserRecipient to a UserRecipientParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// UserRecipientParam.Overrides()
func (r UserRecipient) ToParam() UserRecipientParam {
	return param.Override[UserRecipientParam](json.RawMessage(r.RawJSON()))
}

type UserRecipientPreferences struct {
	Notifications map[string]Preference `json:"notifications,required"`
	Categories    map[string]Preference `json:"categories,nullable"`
	TemplateID    string                `json:"templateId,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Notifications respjson.Field
		Categories    respjson.Field
		TemplateID    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserRecipientPreferences) RawJSON() string { return r.JSON.raw }
func (r *UserRecipientPreferences) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserRecipientParam struct {
	// Use `tenant_id` instead.
	AccountID param.Opt[string] `json:"account_id,omitzero"`
	// The user's email address.
	Email param.Opt[string] `json:"email,omitzero"`
	// The user's preferred ISO 639-1 language code.
	Locale param.Opt[string] `json:"locale,omitzero"`
	// The user's phone number.
	PhoneNumber param.Opt[string] `json:"phone_number,omitzero"`
	// The id of the tenant the user is associated with.
	TenantID param.Opt[string] `json:"tenant_id,omitzero"`
	// The user's unique identifier. Typically, this will match the user id of a user
	// in your system.
	UserID      param.Opt[string]             `json:"user_id,omitzero"`
	Data        map[string]any                `json:"data,omitzero"`
	Preferences UserRecipientPreferencesParam `json:"preferences,omitzero"`
	// Context such as tenant_id to send the notification with.
	Context MessageContextParam `json:"context,omitzero"`
	paramObj
}

func (r UserRecipientParam) MarshalJSON() (data []byte, err error) {
	type shadow UserRecipientParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserRecipientParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Notifications is required.
type UserRecipientPreferencesParam struct {
	Notifications map[string]PreferenceParam `json:"notifications,omitzero,required"`
	TemplateID    param.Opt[string]          `json:"templateId,omitzero"`
	Categories    map[string]PreferenceParam `json:"categories,omitzero"`
	paramObj
}

func (r UserRecipientPreferencesParam) MarshalJSON() (data []byte, err error) {
	type shadow UserRecipientPreferencesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserRecipientPreferencesParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserToken struct {
	// Any of "firebase-fcm", "apn", "expo", "onesignal".
	ProviderKey UserTokenProviderKey `json:"provider_key,required"`
	// Full body of the token. Must match token in URL.
	Token string `json:"token,nullable"`
	// Information about the device the token is associated with.
	Device UserTokenDevice `json:"device,nullable"`
	// ISO 8601 formatted date the token expires. Defaults to 2 months. Set to false to
	// disable expiration.
	ExpiryDate UserTokenExpiryDateUnion `json:"expiry_date,nullable"`
	// Properties sent to the provider along with the token
	Properties any `json:"properties"`
	// Information about the device the token is associated with.
	Tracking UserTokenTracking `json:"tracking,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ProviderKey respjson.Field
		Token       respjson.Field
		Device      respjson.Field
		ExpiryDate  respjson.Field
		Properties  respjson.Field
		Tracking    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserToken) RawJSON() string { return r.JSON.raw }
func (r *UserToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this UserToken to a UserTokenParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// UserTokenParam.Overrides()
func (r UserToken) ToParam() UserTokenParam {
	return param.Override[UserTokenParam](json.RawMessage(r.RawJSON()))
}

type UserTokenProviderKey string

const (
	UserTokenProviderKeyFirebaseFcm UserTokenProviderKey = "firebase-fcm"
	UserTokenProviderKeyApn         UserTokenProviderKey = "apn"
	UserTokenProviderKeyExpo        UserTokenProviderKey = "expo"
	UserTokenProviderKeyOnesignal   UserTokenProviderKey = "onesignal"
)

// Information about the device the token is associated with.
type UserTokenDevice struct {
	// Id of the advertising identifier
	AdID string `json:"ad_id,nullable"`
	// Id of the application the token is used for
	AppID string `json:"app_id,nullable"`
	// Id of the device the token is associated with
	DeviceID string `json:"device_id,nullable"`
	// The device manufacturer
	Manufacturer string `json:"manufacturer,nullable"`
	// The device model
	Model string `json:"model,nullable"`
	// The device platform i.e. android, ios, web
	Platform string `json:"platform,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdID         respjson.Field
		AppID        respjson.Field
		DeviceID     respjson.Field
		Manufacturer respjson.Field
		Model        respjson.Field
		Platform     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserTokenDevice) RawJSON() string { return r.JSON.raw }
func (r *UserTokenDevice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UserTokenExpiryDateUnion contains all possible properties and values from
// [string], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfBool]
type UserTokenExpiryDateUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u UserTokenExpiryDateUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UserTokenExpiryDateUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u UserTokenExpiryDateUnion) RawJSON() string { return u.JSON.raw }

func (r *UserTokenExpiryDateUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information about the device the token is associated with.
type UserTokenTracking struct {
	// The IP address of the device
	IP string `json:"ip,nullable"`
	// The latitude of the device
	Lat string `json:"lat,nullable"`
	// The longitude of the device
	Long string `json:"long,nullable"`
	// The operating system version
	OsVersion string `json:"os_version,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IP          respjson.Field
		Lat         respjson.Field
		Long        respjson.Field
		OsVersion   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserTokenTracking) RawJSON() string { return r.JSON.raw }
func (r *UserTokenTracking) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ProviderKey is required.
type UserTokenParam struct {
	// Any of "firebase-fcm", "apn", "expo", "onesignal".
	ProviderKey UserTokenProviderKey `json:"provider_key,omitzero,required"`
	// Full body of the token. Must match token in URL.
	Token param.Opt[string] `json:"token,omitzero"`
	// Information about the device the token is associated with.
	Device UserTokenDeviceParam `json:"device,omitzero"`
	// ISO 8601 formatted date the token expires. Defaults to 2 months. Set to false to
	// disable expiration.
	ExpiryDate UserTokenExpiryDateUnionParam `json:"expiry_date,omitzero"`
	// Information about the device the token is associated with.
	Tracking UserTokenTrackingParam `json:"tracking,omitzero"`
	// Properties sent to the provider along with the token
	Properties any `json:"properties,omitzero"`
	paramObj
}

func (r UserTokenParam) MarshalJSON() (data []byte, err error) {
	type shadow UserTokenParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserTokenParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information about the device the token is associated with.
type UserTokenDeviceParam struct {
	// Id of the advertising identifier
	AdID param.Opt[string] `json:"ad_id,omitzero"`
	// Id of the application the token is used for
	AppID param.Opt[string] `json:"app_id,omitzero"`
	// Id of the device the token is associated with
	DeviceID param.Opt[string] `json:"device_id,omitzero"`
	// The device manufacturer
	Manufacturer param.Opt[string] `json:"manufacturer,omitzero"`
	// The device model
	Model param.Opt[string] `json:"model,omitzero"`
	// The device platform i.e. android, ios, web
	Platform param.Opt[string] `json:"platform,omitzero"`
	paramObj
}

func (r UserTokenDeviceParam) MarshalJSON() (data []byte, err error) {
	type shadow UserTokenDeviceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserTokenDeviceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type UserTokenExpiryDateUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]   `json:",omitzero,inline"`
	paramUnion
}

func (u UserTokenExpiryDateUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfBool)
}
func (u *UserTokenExpiryDateUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *UserTokenExpiryDateUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

// Information about the device the token is associated with.
type UserTokenTrackingParam struct {
	// The IP address of the device
	IP param.Opt[string] `json:"ip,omitzero"`
	// The latitude of the device
	Lat param.Opt[string] `json:"lat,omitzero"`
	// The longitude of the device
	Long param.Opt[string] `json:"long,omitzero"`
	// The operating system version
	OsVersion param.Opt[string] `json:"os_version,omitzero"`
	paramObj
}

func (r UserTokenTrackingParam) MarshalJSON() (data []byte, err error) {
	type shadow UserTokenTrackingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserTokenTrackingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UtmParam struct {
	Campaign param.Opt[string] `json:"campaign,omitzero"`
	Content  param.Opt[string] `json:"content,omitzero"`
	Medium   param.Opt[string] `json:"medium,omitzero"`
	Source   param.Opt[string] `json:"source,omitzero"`
	Term     param.Opt[string] `json:"term,omitzero"`
	paramObj
}

func (r UtmParam) MarshalJSON() (data []byte, err error) {
	type shadow UtmParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UtmParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WidgetBackground struct {
	BottomColor string `json:"bottomColor,nullable"`
	TopColor    string `json:"topColor,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BottomColor respjson.Field
		TopColor    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WidgetBackground) RawJSON() string { return r.JSON.raw }
func (r *WidgetBackground) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WidgetBackground to a WidgetBackgroundParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WidgetBackgroundParam.Overrides()
func (r WidgetBackground) ToParam() WidgetBackgroundParam {
	return param.Override[WidgetBackgroundParam](json.RawMessage(r.RawJSON()))
}

type WidgetBackgroundParam struct {
	BottomColor param.Opt[string] `json:"bottomColor,omitzero"`
	TopColor    param.Opt[string] `json:"topColor,omitzero"`
	paramObj
}

func (r WidgetBackgroundParam) MarshalJSON() (data []byte, err error) {
	type shadow WidgetBackgroundParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WidgetBackgroundParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
