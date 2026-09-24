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
	"github.com/trycourier/courier-go/v4/packages/param"
	"github.com/trycourier/courier-go/v4/packages/respjson"
)

// Render a template's email content on real email clients and read back the
// screenshots, so you can check how it looks before you send it.
//
// PreviewService contains methods and other services that help with interacting
// with the Courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPreviewService] method instead.
type PreviewService struct {
	Options []option.RequestOption
}

// NewPreviewService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPreviewService(opts ...option.RequestOption) (r PreviewService) {
	r = PreviewService{}
	r.Options = opts
	return
}

// Archive a device set. This is a soft delete — the archived set is returned and
// no longer appears in list results. Runs already created against it keep their
// own copy of the device list and are unaffected. The Courier-provided default set
// cannot be archived and returns 409.
func (r *PreviewService) ArchiveDeviceSet(ctx context.Context, deviceSetID string, opts ...option.RequestOption) (res *DeviceSet, err error) {
	opts = slices.Concat(r.Options, opts)
	if deviceSetID == "" {
		err = errors.New("missing required deviceSetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("previews/device-sets/%s", deviceSetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Create a named, reusable set of preview devices. Every id must be one listed by
// `GET /previews/devices`; any other is a 422.
func (r *PreviewService) NewDeviceSet(ctx context.Context, body PreviewNewDeviceSetParams, opts ...option.RequestOption) (res *DeviceSet, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "previews/device-sets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List the workspace's preview sets. Archived sets are not returned.
func (r *PreviewService) ListDeviceSets(ctx context.Context, opts ...option.RequestOption) (res *DeviceSetListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "previews/device-sets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List the devices a preview can be rendered on. Reference data, identical for
// every workspace — these ids are what a device set is built from and what a run
// reports results for.
func (r *PreviewService) ListDevices(ctx context.Context, opts ...option.RequestOption) (res *PreviewDeviceListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "previews/devices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve a preview set by ID. Archived sets return 404.
func (r *PreviewService) GetDeviceSet(ctx context.Context, deviceSetID string, opts ...option.RequestOption) (res *DeviceSet, err error) {
	opts = slices.Concat(r.Options, opts)
	if deviceSetID == "" {
		err = errors.New("missing required deviceSetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("previews/device-sets/%s", deviceSetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Replace a device set. This is a full replace, not a patch — both the name and
// the device list are always written. The Courier-provided default set cannot be
// changed and returns 409.
func (r *PreviewService) UpdateDeviceSet(ctx context.Context, deviceSetID string, body PreviewUpdateDeviceSetParams, opts ...option.RequestOption) (res *DeviceSet, err error) {
	opts = slices.Concat(r.Options, opts)
	if deviceSetID == "" {
		err = errors.New("missing required deviceSetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("previews/device-sets/%s", deviceSetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Request body for creating or replacing a device set. A full replace, not a patch
// — both fields are always written.
//
// The properties DeviceIDs, Name are required.
type CreateDeviceSetRequestParam struct {
	// The devices the set contains, by `PreviewDevice.id`. At least one is required.
	DeviceIDs []string `json:"device_ids,omitzero" api:"required"`
	// Human-readable name.
	Name string `json:"name" api:"required"`
	paramObj
}

func (r CreateDeviceSetRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateDeviceSetRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateDeviceSetRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A named, reusable list of preview devices.
type DeviceSet struct {
	// Unique identifier for the device set.
	ID string `json:"id" api:"required"`
	// ISO-8601 timestamp of when the set was created.
	CreatedAt string `json:"created_at" api:"required"`
	// The devices in this set, by `PreviewDevice.id`.
	DeviceIDs []string `json:"device_ids" api:"required"`
	// Human-readable name.
	Name string `json:"name" api:"required"`
	// ISO-8601 timestamp of when the set was last written.
	UpdatedAt string `json:"updated_at" api:"required"`
	// ISO-8601 timestamp of when the set was archived. Present only on the archive
	// response, which is the one place the state is observable.
	ArchivedAt string `json:"archived_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		DeviceIDs   respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		ArchivedAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceSet) RawJSON() string { return r.JSON.raw }
func (r *DeviceSet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The workspace's active device sets. Not paginated.
type DeviceSetListResponse struct {
	Results []DeviceSet `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceSetListResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceSetListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One mail app on one platform, operating system and theme that a preview can be
// rendered on. Reference data, identical for every workspace. Every field is
// always present; `platform` and `platform_version` are null where they do not
// apply.
type PreviewDevice struct {
	// The device's identifier, used in `device_ids` when creating a device set or a
	// run.
	ID string `json:"id" api:"required"`
	// The mail app. For webmail it is the service (`outlook_com`, `gmail_com`); for
	// mobile the app (`apple_mail`, `gmail`); for desktop the app together with the
	// version it is sold under (`outlook_2019`, `outlook_microsoft_365`,
	// `apple_mail_16`), because that version is what separates one desktop Outlook
	// from another.
	App string `json:"app" api:"required"`
	// Where the app runs.
	//
	// Any of "webmail", "mobile", "desktop".
	Category PreviewDeviceCategory `json:"category" api:"required"`
	// Display name. Render it as-is rather than parsing it. It is also what separates
	// the two 120-dpi Outlook renders from their 100% siblings, which are otherwise
	// identical field for field.
	Name string `json:"name" api:"required"`
	// The operating system.
	Os string `json:"os" api:"required"`
	// The operating system's version. Always set.
	OsVersion string `json:"os_version" api:"required"`
	// What the app runs on — the browser for webmail (`chrome`, `edge`, `firefox`),
	// the phone for mobile (`iphone`, `pixel`). Null for desktop, where the app runs
	// on nothing but the OS.
	Platform string `json:"platform" api:"required"`
	// Which one of the platform — the phone model for mobile (`15_pro_max`, `10`).
	// Null for webmail, which always renders in the current browser, and for desktop.
	PlatformVersion string `json:"platform_version" api:"required"`
	// Whether the email is rendered in light or dark mode.
	//
	// Any of "light", "dark".
	Theme PreviewDeviceTheme `json:"theme" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		App             respjson.Field
		Category        respjson.Field
		Name            respjson.Field
		Os              respjson.Field
		OsVersion       respjson.Field
		Platform        respjson.Field
		PlatformVersion respjson.Field
		Theme           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PreviewDevice) RawJSON() string { return r.JSON.raw }
func (r *PreviewDevice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Where the app runs.
type PreviewDeviceCategory string

const (
	PreviewDeviceCategoryWebmail PreviewDeviceCategory = "webmail"
	PreviewDeviceCategoryMobile  PreviewDeviceCategory = "mobile"
	PreviewDeviceCategoryDesktop PreviewDeviceCategory = "desktop"
)

// Whether the email is rendered in light or dark mode.
type PreviewDeviceTheme string

const (
	PreviewDeviceThemeLight PreviewDeviceTheme = "light"
	PreviewDeviceThemeDark  PreviewDeviceTheme = "dark"
)

// The full catalog of renderable devices. Not paginated.
type PreviewDeviceListResponse struct {
	Results []PreviewDevice `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PreviewDeviceListResponse) RawJSON() string { return r.JSON.raw }
func (r *PreviewDeviceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PreviewNewDeviceSetParams struct {
	// Request body for creating or replacing a device set. A full replace, not a patch
	// — both fields are always written.
	CreateDeviceSetRequest CreateDeviceSetRequestParam
	paramObj
}

func (r PreviewNewDeviceSetParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateDeviceSetRequest)
}
func (r *PreviewNewDeviceSetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PreviewUpdateDeviceSetParams struct {
	// Request body for creating or replacing a device set. A full replace, not a patch
	// — both fields are always written.
	CreateDeviceSetRequest CreateDeviceSetRequestParam
	paramObj
}

func (r PreviewUpdateDeviceSetParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateDeviceSetRequest)
}
func (r *PreviewUpdateDeviceSetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
