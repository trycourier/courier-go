// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/courier-go/internal/requestconfig"
	"github.com/stainless-sdks/courier-go/option"
)

// NotificationDraftService contains methods and other services that help with
// interacting with the courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNotificationDraftService] method instead.
type NotificationDraftService struct {
	Options []option.RequestOption
}

// NewNotificationDraftService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNotificationDraftService(opts ...option.RequestOption) (r NotificationDraftService) {
	r = NotificationDraftService{}
	r.Options = opts
	return
}

func (r *NotificationDraftService) GetContent(ctx context.Context, id string, opts ...option.RequestOption) (res *NotificationGetContent, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("notifications/%s/draft/content", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
