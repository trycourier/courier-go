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

// PreferenceSectionTopicService contains methods and other services that help with
// interacting with the Courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPreferenceSectionTopicService] method instead.
type PreferenceSectionTopicService struct {
	Options []option.RequestOption
}

// NewPreferenceSectionTopicService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPreferenceSectionTopicService(opts ...option.RequestOption) (r PreferenceSectionTopicService) {
	r = PreferenceSectionTopicService{}
	r.Options = opts
	return
}

// Create a subscription preference topic inside a section. Fails with 404 if the
// section does not exist. The topic id is generated and returned.
func (r *PreferenceSectionTopicService) New(ctx context.Context, sectionID string, body PreferenceSectionTopicNewParams, opts ...option.RequestOption) (res *PreferenceTopicGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if sectionID == "" {
		err = errors.New("missing required section_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("preferences/sections/%s/topics", sectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a topic within a section. Returns 404 if the section does not exist,
// the topic does not exist, or the topic belongs to a different section.
func (r *PreferenceSectionTopicService) Get(ctx context.Context, topicID string, query PreferenceSectionTopicGetParams, opts ...option.RequestOption) (res *PreferenceTopicGetResponse, err error) {
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

// List the topics in a preference section.
func (r *PreferenceSectionTopicService) List(ctx context.Context, sectionID string, opts ...option.RequestOption) (res *PreferenceTopicListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if sectionID == "" {
		err = errors.New("missing required section_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("preferences/sections/%s/topics", sectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Archive a topic and remove it from its section. Same 404 rules as GET.
func (r *PreferenceSectionTopicService) Archive(ctx context.Context, topicID string, body PreferenceSectionTopicArchiveParams, opts ...option.RequestOption) (err error) {
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

// Replace a topic within a section. Full document replacement; missing optional
// fields are cleared. Same 404 rules as GET.
func (r *PreferenceSectionTopicService) Replace(ctx context.Context, topicID string, params PreferenceSectionTopicReplaceParams, opts ...option.RequestOption) (res *PreferenceTopicGetResponse, err error) {
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

type PreferenceSectionTopicNewParams struct {
	// Request body for creating a preference topic.
	PreferenceTopicCreateRequest PreferenceTopicCreateRequestParam
	paramObj
}

func (r PreferenceSectionTopicNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PreferenceTopicCreateRequest)
}
func (r *PreferenceSectionTopicNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PreferenceSectionTopicGetParams struct {
	SectionID string `path:"section_id" api:"required" json:"-"`
	paramObj
}

type PreferenceSectionTopicArchiveParams struct {
	SectionID string `path:"section_id" api:"required" json:"-"`
	paramObj
}

type PreferenceSectionTopicReplaceParams struct {
	SectionID string `path:"section_id" api:"required" json:"-"`
	// Request body for replacing a preference topic. Full document replacement;
	// missing optional fields are cleared.
	PreferenceTopicReplaceRequest PreferenceTopicReplaceRequestParam
	paramObj
}

func (r PreferenceSectionTopicReplaceParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PreferenceTopicReplaceRequest)
}
func (r *PreferenceSectionTopicReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
