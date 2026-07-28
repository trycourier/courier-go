// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/trycourier/courier-go/v4/internal/apijson"
	"github.com/trycourier/courier-go/v4/internal/requestconfig"
	"github.com/trycourier/courier-go/v4/option"
	"github.com/trycourier/courier-go/v4/packages/param"
)

// InboxMessageService contains methods and other services that help with
// interacting with the Courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInboxMessageService] method instead.
type InboxMessageService struct {
	Options []option.RequestOption
}

// NewInboxMessageService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInboxMessageService(opts ...option.RequestOption) (r InboxMessageService) {
	r = InboxMessageService{}
	r.Options = opts
	return
}

// Delete a user's inbox message. The message is removed from every inbox read (it
// stops appearing in the recipient's Inbox); it can be restored.
func (r *InboxMessageService) Delete(ctx context.Context, messageID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return err
	}
	path := fmt.Sprintf("inbox/messages/%s", messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Restore a previously deleted inbox message.
func (r *InboxMessageService) Restore(ctx context.Context, messageID string, body InboxMessageRestoreParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return err
	}
	path := fmt.Sprintf("inbox/messages/%s/restore", messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return err
}

type InboxMessageRestoreParams struct {
	paramObj
}

func (r InboxMessageRestoreParams) MarshalJSON() (data []byte, err error) {
	type shadow InboxMessageRestoreParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InboxMessageRestoreParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
