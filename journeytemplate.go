// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/trycourier/courier-go/v4/internal/apijson"
	"github.com/trycourier/courier-go/v4/internal/apiquery"
	shimjson "github.com/trycourier/courier-go/v4/internal/encoding/json"
	"github.com/trycourier/courier-go/v4/internal/requestconfig"
	"github.com/trycourier/courier-go/v4/option"
	"github.com/trycourier/courier-go/v4/packages/param"
)

// JourneyTemplateService contains methods and other services that help with
// interacting with the Courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewJourneyTemplateService] method instead.
type JourneyTemplateService struct {
	Options []option.RequestOption
}

// NewJourneyTemplateService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewJourneyTemplateService(opts ...option.RequestOption) (r JourneyTemplateService) {
	r = JourneyTemplateService{}
	r.Options = opts
	return
}

// Create a notification template scoped to this journey. Defaults to `DRAFT`
// state; pass `state: "PUBLISHED"` to publish on create.
func (r *JourneyTemplateService) New(ctx context.Context, templateID string, body JourneyTemplateNewParams, opts ...option.RequestOption) (res *JourneyTemplateGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if templateID == "" {
		err = errors.New("missing required templateId parameter")
		return nil, err
	}
	path := fmt.Sprintf("journeys/%s/templates", templateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a journey's own notification template with its name, brand, subscription
// topic, and content. Defaults to the published version.
func (r *JourneyTemplateService) Get(ctx context.Context, notificationID string, query JourneyTemplateGetParams, opts ...option.RequestOption) (res *JourneyTemplateGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.TemplateID == "" {
		err = errors.New("missing required templateId parameter")
		return nil, err
	}
	if notificationID == "" {
		err = errors.New("missing required notificationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("journeys/%s/templates/%s", query.TemplateID, notificationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List notification templates scoped to this journey. Journey-scoped notification
// templates can only be referenced from `send` nodes within the same journey.
func (r *JourneyTemplateService) List(ctx context.Context, templateID string, query JourneyTemplateListParams, opts ...option.RequestOption) (res *JourneyTemplateListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if templateID == "" {
		err = errors.New("missing required templateId parameter")
		return nil, err
	}
	path := fmt.Sprintf("journeys/%s/templates", templateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Archives one journey's notification template, preventing further sends. Detach
// any send node referencing it beforehand.
func (r *JourneyTemplateService) Archive(ctx context.Context, notificationID string, body JourneyTemplateArchiveParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.TemplateID == "" {
		err = errors.New("missing required templateId parameter")
		return err
	}
	if notificationID == "" {
		err = errors.New("missing required notificationId parameter")
		return err
	}
	path := fmt.Sprintf("journeys/%s/templates/%s", body.TemplateID, notificationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Lists the published versions of a template that belongs to a journey, most
// recent first. Paged by cursor.
func (r *JourneyTemplateService) ListVersions(ctx context.Context, notificationID string, query JourneyTemplateListVersionsParams, opts ...option.RequestOption) (res *NotificationTemplateVersionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.TemplateID == "" {
		err = errors.New("missing required templateId parameter")
		return nil, err
	}
	if notificationID == "" {
		err = errors.New("missing required notificationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("journeys/%s/templates/%s/versions", query.TemplateID, notificationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Publishes a journey-scoped template's draft as a new version. Pass a version
// instead to roll back the template to an earlier publish.
func (r *JourneyTemplateService) Publish(ctx context.Context, notificationID string, params JourneyTemplatePublishParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.TemplateID == "" {
		err = errors.New("missing required templateId parameter")
		return err
	}
	if notificationID == "" {
		err = errors.New("missing required notificationId parameter")
		return err
	}
	path := fmt.Sprintf("journeys/%s/templates/%s/publish", params.TemplateID, notificationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return err
}

// Replace the elemental content of a journey-scoped notification template.
// Overwrites all elements in the template draft with the provided content.
func (r *JourneyTemplateService) PutContent(ctx context.Context, notificationID string, params JourneyTemplatePutContentParams, opts ...option.RequestOption) (res *NotificationContentMutationResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.TemplateID == "" {
		err = errors.New("missing required templateId parameter")
		return nil, err
	}
	if notificationID == "" {
		err = errors.New("missing required notificationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("journeys/%s/templates/%s/content", params.TemplateID, notificationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Set locale-specific content overrides for a journey-scoped notification
// template. Each element override must reference an existing element by ID.
func (r *JourneyTemplateService) PutLocale(ctx context.Context, localeID string, params JourneyTemplatePutLocaleParams, opts ...option.RequestOption) (res *NotificationContentMutationResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.TemplateID == "" {
		err = errors.New("missing required templateId parameter")
		return nil, err
	}
	if params.NotificationID == "" {
		err = errors.New("missing required notificationId parameter")
		return nil, err
	}
	if localeID == "" {
		err = errors.New("missing required localeId parameter")
		return nil, err
	}
	path := fmt.Sprintf("journeys/%s/templates/%s/locales/%s", params.TemplateID, params.NotificationID, localeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Replaces the draft content of one journey's notification template. Publish it
// before send nodes referencing it render the change.
func (r *JourneyTemplateService) Replace(ctx context.Context, notificationID string, params JourneyTemplateReplaceParams, opts ...option.RequestOption) (res *JourneyTemplateGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.TemplateID == "" {
		err = errors.New("missing required templateId parameter")
		return nil, err
	}
	if notificationID == "" {
		err = errors.New("missing required notificationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("journeys/%s/templates/%s", params.TemplateID, notificationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Returns the Elemental elements and version of a journey-scoped template's
// content. Compare versions to see what changed between publishes.
func (r *JourneyTemplateService) GetContent(ctx context.Context, notificationID string, params JourneyTemplateGetContentParams, opts ...option.RequestOption) (res *NotificationContentGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.TemplateID == "" {
		err = errors.New("missing required templateId parameter")
		return nil, err
	}
	if notificationID == "" {
		err = errors.New("missing required notificationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("journeys/%s/templates/%s/content", params.TemplateID, notificationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type JourneyTemplateNewParams struct {
	// Request body for creating a notification template scoped to a journey.
	JourneyTemplateCreateRequest JourneyTemplateCreateRequestParam
	paramObj
}

func (r JourneyTemplateNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.JourneyTemplateCreateRequest)
}
func (r *JourneyTemplateNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JourneyTemplateGetParams struct {
	TemplateID string `path:"templateId" api:"required" json:"-"`
	paramObj
}

type JourneyTemplateListParams struct {
	// Pagination cursor from a prior response.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Page size. Minimum 1, maximum 100.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [JourneyTemplateListParams]'s query parameters as
// `url.Values`.
func (r JourneyTemplateListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type JourneyTemplateArchiveParams struct {
	TemplateID string `path:"templateId" api:"required" json:"-"`
	paramObj
}

type JourneyTemplateListVersionsParams struct {
	TemplateID string `path:"templateId" api:"required" json:"-"`
	paramObj
}

type JourneyTemplatePublishParams struct {
	TemplateID string `path:"templateId" api:"required" json:"-"`
	// Request body for publishing a journey-scoped notification template. Pass
	// `version` to roll back to a prior version; omit to publish the current draft.
	JourneyTemplatePublishRequest JourneyTemplatePublishRequestParam
	paramObj
}

func (r JourneyTemplatePublishParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.JourneyTemplatePublishRequest)
}
func (r *JourneyTemplatePublishParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JourneyTemplatePutContentParams struct {
	TemplateID string `path:"templateId" api:"required" json:"-"`
	// Request body for replacing the elemental content of a notification template.
	NotificationContentPutRequest NotificationContentPutRequestParam
	paramObj
}

func (r JourneyTemplatePutContentParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.NotificationContentPutRequest)
}
func (r *JourneyTemplatePutContentParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JourneyTemplatePutLocaleParams struct {
	TemplateID     string `path:"templateId" api:"required" json:"-"`
	NotificationID string `path:"notificationId" api:"required" json:"-"`
	// Request body for setting locale-specific content overrides. Each element
	// override must include the target element ID.
	NotificationLocalePutRequest NotificationLocalePutRequestParam
	paramObj
}

func (r JourneyTemplatePutLocaleParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.NotificationLocalePutRequest)
}
func (r *JourneyTemplatePutLocaleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JourneyTemplateReplaceParams struct {
	TemplateID string `path:"templateId" api:"required" json:"-"`
	// Request body for replacing a journey-scoped notification template draft.
	JourneyTemplateReplaceRequest JourneyTemplateReplaceRequestParam
	paramObj
}

func (r JourneyTemplateReplaceParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.JourneyTemplateReplaceRequest)
}
func (r *JourneyTemplateReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JourneyTemplateGetContentParams struct {
	TemplateID string `path:"templateId" api:"required" json:"-"`
	// Accepts `draft`, `published`, or a version string (e.g., `v001`). Defaults to
	// `published`.
	Version param.Opt[string] `query:"version,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [JourneyTemplateGetContentParams]'s query parameters as
// `url.Values`.
func (r JourneyTemplateGetContentParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
