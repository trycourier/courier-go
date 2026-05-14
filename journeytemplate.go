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

// Create a notification template scoped to this journey. The template is created
// in DRAFT state.
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

// Fetch a journey-scoped notification template by id. Pass `?version=draft`
// (default `published`) to retrieve the working draft, or `?version=vN` for a
// historical version.
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

// List notification templates scoped to this journey. Templates scoped to a
// journey can only be referenced from `send` nodes of the same journey.
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

// Archive a journey-scoped notification template. Archived templates cannot be
// sent.
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

// List published versions of a journey-scoped notification template, ordered most
// recent first.
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

// Publish the current draft of a journey-scoped notification template.
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

// Replace a journey-scoped notification template draft.
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

type JourneyTemplateNewParams struct {
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
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
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
	TemplateID                    string `path:"templateId" api:"required" json:"-"`
	JourneyTemplatePublishRequest JourneyTemplatePublishRequestParam
	paramObj
}

func (r JourneyTemplatePublishParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.JourneyTemplatePublishRequest)
}
func (r *JourneyTemplatePublishParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JourneyTemplateReplaceParams struct {
	TemplateID                    string `path:"templateId" api:"required" json:"-"`
	JourneyTemplateReplaceRequest JourneyTemplateReplaceRequestParam
	paramObj
}

func (r JourneyTemplateReplaceParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.JourneyTemplateReplaceRequest)
}
func (r *JourneyTemplateReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
