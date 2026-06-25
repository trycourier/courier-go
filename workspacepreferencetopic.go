// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/trycourier/courier-go/v4/internal/apijson"
	shimjson "github.com/trycourier/courier-go/v4/internal/encoding/json"
	"github.com/trycourier/courier-go/v4/internal/requestconfig"
	"github.com/trycourier/courier-go/v4/option"
)

// WorkspacePreferenceTopicService contains methods and other services that help
// with interacting with the Courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkspacePreferenceTopicService] method instead.
type WorkspacePreferenceTopicService struct {
	Options []option.RequestOption
}

// NewWorkspacePreferenceTopicService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewWorkspacePreferenceTopicService(opts ...option.RequestOption) (r WorkspacePreferenceTopicService) {
	r = WorkspacePreferenceTopicService{}
	r.Options = opts
	return
}

// Create a subscription preference topic inside a workspace preference. Fails with
// 404 if the workspace preference does not exist. The topic id is generated and
// returned.
func (r *WorkspacePreferenceTopicService) New(ctx context.Context, sectionID string, body WorkspacePreferenceTopicNewParams, opts ...option.RequestOption) (res *WorkspacePreferenceTopicGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if sectionID == "" {
		err = errors.New("missing required section_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("preferences/sections/%s/topics", sectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a topic within a workspace preference. Returns 404 if the workspace
// preference does not exist, the topic does not exist, or the topic belongs to a
// different workspace preference.
func (r *WorkspacePreferenceTopicService) Get(ctx context.Context, topicID string, query WorkspacePreferenceTopicGetParams, opts ...option.RequestOption) (res *WorkspacePreferenceTopicGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.SectionID == "" {
		err = errors.New("missing required section_id parameter")
		return nil, err
	}
	if topicID == "" {
		err = errors.New("missing required topic_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("preferences/sections/%s/topics/%s", query.SectionID, topicID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List the topics in a workspace preference.
func (r *WorkspacePreferenceTopicService) List(ctx context.Context, sectionID string, opts ...option.RequestOption) (res *WorkspacePreferenceTopicListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if sectionID == "" {
		err = errors.New("missing required section_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("preferences/sections/%s/topics", sectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Archive a topic and remove it from its workspace preference. Same 404 rules as
// GET.
func (r *WorkspacePreferenceTopicService) Archive(ctx context.Context, topicID string, body WorkspacePreferenceTopicArchiveParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.SectionID == "" {
		err = errors.New("missing required section_id parameter")
		return err
	}
	if topicID == "" {
		err = errors.New("missing required topic_id parameter")
		return err
	}
	path := fmt.Sprintf("preferences/sections/%s/topics/%s", body.SectionID, topicID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Replace a topic within a workspace preference. Full document replacement;
// missing optional fields are cleared. Same 404 rules as GET.
func (r *WorkspacePreferenceTopicService) Replace(ctx context.Context, topicID string, params WorkspacePreferenceTopicReplaceParams, opts ...option.RequestOption) (res *WorkspacePreferenceTopicGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.SectionID == "" {
		err = errors.New("missing required section_id parameter")
		return nil, err
	}
	if topicID == "" {
		err = errors.New("missing required topic_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("preferences/sections/%s/topics/%s", params.SectionID, topicID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

type WorkspacePreferenceTopicNewParams struct {
	// Request body for creating a preference topic.
	WorkspacePreferenceTopicCreateRequest WorkspacePreferenceTopicCreateRequestParam
	paramObj
}

func (r WorkspacePreferenceTopicNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.WorkspacePreferenceTopicCreateRequest)
}
func (r *WorkspacePreferenceTopicNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkspacePreferenceTopicGetParams struct {
	SectionID string `path:"section_id" api:"required" json:"-"`
	paramObj
}

type WorkspacePreferenceTopicArchiveParams struct {
	SectionID string `path:"section_id" api:"required" json:"-"`
	paramObj
}

type WorkspacePreferenceTopicReplaceParams struct {
	SectionID string `path:"section_id" api:"required" json:"-"`
	// Request body for replacing a preference topic. Full document replacement;
	// missing optional fields are cleared.
	WorkspacePreferenceTopicReplaceRequest WorkspacePreferenceTopicReplaceRequestParam
	paramObj
}

func (r WorkspacePreferenceTopicReplaceParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.WorkspacePreferenceTopicReplaceRequest)
}
func (r *WorkspacePreferenceTopicReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
