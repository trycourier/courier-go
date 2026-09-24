// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"github.com/trycourier/courier-go/v4/option"
)

// NotificationPreviewService contains methods and other services that help with
// interacting with the Courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNotificationPreviewService] method instead.
type NotificationPreviewService struct {
	Options []option.RequestOption
	// Render a template's email content on real email clients and read back the
	// screenshots, so you can check how it looks before you send it.
	Runs NotificationPreviewRunService
}

// NewNotificationPreviewService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNotificationPreviewService(opts ...option.RequestOption) (r NotificationPreviewService) {
	r = NotificationPreviewService{}
	r.Options = opts
	r.Runs = NewNotificationPreviewRunService(opts...)
	return
}
