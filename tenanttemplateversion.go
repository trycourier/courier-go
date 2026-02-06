// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/trycourier/courier-go/v4/internal/requestconfig"
	"github.com/trycourier/courier-go/v4/option"
)

// TenantTemplateVersionService contains methods and other services that help with
// interacting with the Courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTenantTemplateVersionService] method instead.
type TenantTemplateVersionService struct {
	Options []option.RequestOption
}

// NewTenantTemplateVersionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTenantTemplateVersionService(opts ...option.RequestOption) (r TenantTemplateVersionService) {
	r = TenantTemplateVersionService{}
	r.Options = opts
	return
}

// Fetches a specific version of a tenant template.
//
// Supports the following version formats:
//
// - `latest` - The most recent version of the template
// - `published` - The currently published version
// - `v{version}` - A specific version (e.g., "v1", "v2", "v1.0.0")
func (r *TenantTemplateVersionService) Get(ctx context.Context, version string, query TenantTemplateVersionGetParams, opts ...option.RequestOption) (res *BaseTemplateTenantAssociation, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.TenantID == "" {
		err = errors.New("missing required tenant_id parameter")
		return
	}
	if query.TemplateID == "" {
		err = errors.New("missing required template_id parameter")
		return
	}
	if version == "" {
		err = errors.New("missing required version parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s/templates/%s/versions/%s", query.TenantID, query.TemplateID, version)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TenantTemplateVersionGetParams struct {
	TenantID   string `path:"tenant_id,required" json:"-"`
	TemplateID string `path:"template_id,required" json:"-"`
	paramObj
}
