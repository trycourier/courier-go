// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	shimjson "github.com/trycourier/courier-go/v4/internal/encoding/json"
	"github.com/trycourier/courier-go/v4/internal/requestconfig"
	"github.com/trycourier/courier-go/v4/option"
)

// TranslationService contains methods and other services that help with
// interacting with the Courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTranslationService] method instead.
type TranslationService struct {
	Options []option.RequestOption
}

// NewTranslationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTranslationService(opts ...option.RequestOption) (r TranslationService) {
	r = TranslationService{}
	r.Options = opts
	return
}

// Get translations by locale
func (r *TranslationService) Get(ctx context.Context, locale string, query TranslationGetParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.Domain == "" {
		err = errors.New("missing required domain parameter")
		return
	}
	if locale == "" {
		err = errors.New("missing required locale parameter")
		return
	}
	path := fmt.Sprintf("translations/%s/%s", query.Domain, locale)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a translation
func (r *TranslationService) Update(ctx context.Context, locale string, params TranslationUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if params.Domain == "" {
		err = errors.New("missing required domain parameter")
		return
	}
	if locale == "" {
		err = errors.New("missing required locale parameter")
		return
	}
	path := fmt.Sprintf("translations/%s/%s", params.Domain, locale)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, nil, opts...)
	return
}

type TranslationGetParams struct {
	Domain string `path:"domain,required" json:"-"`
	paramObj
}

type TranslationUpdateParams struct {
	Domain string `path:"domain,required" json:"-"`
	Body   string
	paramObj
}

func (r TranslationUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *TranslationUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}
