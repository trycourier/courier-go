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
	"github.com/trycourier/courier-go/v4/internal/requestconfig"
	"github.com/trycourier/courier-go/v4/option"
	"github.com/trycourier/courier-go/v4/packages/param"
	"github.com/trycourier/courier-go/v4/packages/respjson"
	"github.com/trycourier/courier-go/v4/shared"
)

// ProviderService contains methods and other services that help with interacting
// with the Courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProviderService] method instead.
type ProviderService struct {
	Options []option.RequestOption
	Catalog ProviderCatalogService
}

// NewProviderService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProviderService(opts ...option.RequestOption) (r ProviderService) {
	r = ProviderService{}
	r.Options = opts
	r.Catalog = NewProviderCatalogService(opts...)
	return
}

// Create a new provider configuration. The `provider` field must be a known
// Courier provider key (see catalog).
func (r *ProviderService) New(ctx context.Context, body ProviderNewParams, opts ...option.RequestOption) (res *Provider, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "providers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Fetch a single provider configuration by ID.
func (r *ProviderService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Provider, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("providers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update an existing provider configuration. The `provider` key is required. All
// other fields are optional — omitted fields are cleared from the stored
// configuration (this is a full replacement, not a partial merge).
func (r *ProviderService) Update(ctx context.Context, id string, body ProviderUpdateParams, opts ...option.RequestOption) (res *Provider, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("providers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List configured provider integrations for the current workspace. Supports
// cursor-based pagination.
func (r *ProviderService) List(ctx context.Context, query ProviderListParams, opts ...option.RequestOption) (res *ProviderListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "providers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete a provider configuration. Returns 409 if the provider is still referenced
// by routing or notifications.
func (r *ProviderService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("providers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// A configured provider in the workspace.
type Provider struct {
	// A unique identifier for the provider configuration.
	ID string `json:"id" api:"required"`
	// Unix timestamp (ms) of when the provider was created.
	Created int64 `json:"created" api:"required"`
	// The provider key (e.g. "sendgrid", "twilio", "slack").
	Provider string `json:"provider" api:"required"`
	// Provider-specific settings (snake_case keys on the wire).
	Settings map[string]any `json:"settings" api:"required"`
	// Display title. Defaults to "Default Configuration" when not explicitly set.
	Title string `json:"title" api:"required"`
	// Optional alias for this configuration.
	Alias string `json:"alias"`
	// Unix timestamp (ms) of when the provider was last updated.
	Updated int64 `json:"updated" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Created     respjson.Field
		Provider    respjson.Field
		Settings    respjson.Field
		Title       respjson.Field
		Alias       respjson.Field
		Updated     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Provider) RawJSON() string { return r.JSON.raw }
func (r *Provider) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A provider type from the catalog. Contains the key, display name, description,
// and a `settings` object describing configuration schema fields.
type ProvidersCatalogEntry struct {
	// Courier taxonomy channel (e.g. email, push, sms, direct_message, inbox,
	// webhook).
	Channel string `json:"channel" api:"required"`
	// Short description of the provider.
	Description string `json:"description" api:"required"`
	// Human-readable display name.
	Name string `json:"name" api:"required"`
	// The provider key (e.g. "sendgrid", "twilio").
	Provider string `json:"provider" api:"required"`
	// Map of setting field names (snake_case) to their schema descriptors. Each
	// descriptor has `type` and `required`. Empty when the provider has no
	// configurable schema.
	Settings map[string]ProvidersCatalogEntrySetting `json:"settings" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Channel     respjson.Field
		Description respjson.Field
		Name        respjson.Field
		Provider    respjson.Field
		Settings    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProvidersCatalogEntry) RawJSON() string { return r.JSON.raw }
func (r *ProvidersCatalogEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes a single configuration field in the provider catalog.
type ProvidersCatalogEntrySetting struct {
	// Whether this field is required when configuring the provider.
	Required bool `json:"required" api:"required"`
	// The field's data type (e.g. "string", "boolean", "enum").
	Type string `json:"type" api:"required"`
	// Allowed values when `type` is "enum".
	Values []string `json:"values"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Required    respjson.Field
		Type        respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProvidersCatalogEntrySetting) RawJSON() string { return r.JSON.raw }
func (r *ProvidersCatalogEntrySetting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Paginated list of provider configurations.
type ProviderListResponse struct {
	Paging  shared.Paging `json:"paging" api:"required"`
	Results []Provider    `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Paging      respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProviderListResponse) RawJSON() string { return r.JSON.raw }
func (r *ProviderListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProviderNewParams struct {
	// The provider key identifying the type (e.g. "sendgrid", "twilio"). Must be a
	// known Courier provider — see the catalog endpoint for valid keys.
	Provider string `json:"provider" api:"required"`
	// Optional alias for this configuration.
	Alias param.Opt[string] `json:"alias,omitzero"`
	// Optional display title. Omit to use "Default Configuration".
	Title param.Opt[string] `json:"title,omitzero"`
	// Provider-specific settings (snake_case keys). Defaults to an empty object when
	// omitted. Use the catalog endpoint to discover required fields for a given
	// provider — omitting a required field returns a 400 validation error.
	Settings map[string]any `json:"settings,omitzero"`
	paramObj
}

func (r ProviderNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ProviderNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProviderNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProviderUpdateParams struct {
	// The provider key identifying the type.
	Provider string `json:"provider" api:"required"`
	// Updated alias. Omit to clear.
	Alias param.Opt[string] `json:"alias,omitzero"`
	// Updated display title.
	Title param.Opt[string] `json:"title,omitzero"`
	// Provider-specific settings (snake_case keys). Replaces the full settings object
	// — omitted settings fields are removed. Use the catalog endpoint to check
	// required fields.
	Settings map[string]any `json:"settings,omitzero"`
	paramObj
}

func (r ProviderUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ProviderUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProviderUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProviderListParams struct {
	// Opaque cursor for fetching the next page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ProviderListParams]'s query parameters as `url.Values`.
func (r ProviderListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
