// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier

import (
	"context"
	"net/http"
	"slices"

	"github.com/trycourier/courier-go/v4/internal/apijson"
	"github.com/trycourier/courier-go/v4/internal/requestconfig"
	"github.com/trycourier/courier-go/v4/option"
	"github.com/trycourier/courier-go/v4/packages/param"
	"github.com/trycourier/courier-go/v4/packages/respjson"
)

// InboundService contains methods and other services that help with interacting
// with the Courier API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInboundService] method instead.
type InboundService struct {
	Options []option.RequestOption
}

// NewInboundService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewInboundService(opts ...option.RequestOption) (r InboundService) {
	r = InboundService{}
	r.Options = opts
	return
}

// Courier Track Event
func (r *InboundService) TrackEvent(ctx context.Context, body InboundTrackEventParams, opts ...option.RequestOption) (res *InboundTrackEventResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "inbound/courier"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type InboundTrackEventResponse struct {
	// A successful call returns a `202` status code along with a `requestId` in the
	// response body.
	MessageID string `json:"messageId,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MessageID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboundTrackEventResponse) RawJSON() string { return r.JSON.raw }
func (r *InboundTrackEventResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboundTrackEventParams struct {
	// A descriptive name of the event. This name will appear as a trigger in the
	// Courier Automation Trigger node.
	Event string `json:"event,required"`
	// A required unique identifier that will be used to de-duplicate requests. If not
	// unique, will respond with 409 Conflict status
	MessageID  string         `json:"messageId,required"`
	Properties map[string]any `json:"properties,omitzero,required"`
	// Any of "track".
	Type InboundTrackEventParamsType `json:"type,omitzero,required"`
	// The user id associatiated with the track
	UserID param.Opt[string] `json:"userId,omitzero"`
	paramObj
}

func (r InboundTrackEventParams) MarshalJSON() (data []byte, err error) {
	type shadow InboundTrackEventParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InboundTrackEventParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboundTrackEventParamsType string

const (
	InboundTrackEventParamsTypeTrack InboundTrackEventParamsType = "track"
)
