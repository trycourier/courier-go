// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/trycourier/courier-go/v3/core"
)

type GetMessageHistoryRequest struct {
	// A supported Message History type that will filter the events returned.
	Type *string `json:"-" url:"type,omitempty"`
}

type ListMessagesRequest struct {
	// A boolean value that indicates whether archived messages should be included in the response.
	Archived *bool `json:"-" url:"archived,omitempty"`
	// A unique identifier that allows for fetching the next set of messages.
	Cursor *string `json:"-" url:"cursor,omitempty"`
	// A unique identifier representing the event that was used to send the event.
	Event *string `json:"-" url:"event,omitempty"`
	// A unique identifier representing the list the message was sent to.
	List *string `json:"-" url:"list,omitempty"`
	// A unique identifier representing the message_id returned from either /send or /send/list.
	MessageId *string `json:"-" url:"messageId,omitempty"`
	// A unique identifier representing the notification that was used to send the event.
	Notification *string `json:"-" url:"notification,omitempty"`
	// The key assocated to the provider you want to filter on. E.g., sendgrid, inbox, twilio, slack, msteams, etc. Allows multiple values to be set in query parameters.
	Provider []*string `json:"-" url:"provider,omitempty"`
	// A unique identifier representing the recipient associated with the requested profile.
	Recipient *string `json:"-" url:"recipient,omitempty"`
	// An indicator of the current status of the message. Allows multiple values to be set in query parameters.
	Status []*string `json:"-" url:"status,omitempty"`
	// A tag placed in the metadata.tags during a notification send. Allows multiple values to be set in query parameters.
	Tag []*string `json:"-" url:"tag,omitempty"`
	// A comma delimited list of 'tags'. Messages will be returned if they match any of the tags passed in.
	Tags *string `json:"-" url:"tags,omitempty"`
	// Messages sent with the context of a Tenant
	TenantId *string `json:"-" url:"tenant_id,omitempty"`
	// The enqueued datetime of a message to filter out messages received before.
	EnqueuedAfter *string `json:"-" url:"enqueued_after,omitempty"`
	// The unique identifier used to trace the requests
	TraceId *string `json:"-" url:"traceId,omitempty"`
}

type ListMessagesResponse struct {
	// Paging information for the result set.
	Paging *Paging `json:"paging,omitempty" url:"paging,omitempty"`
	// An array of messages with their details.
	Results []*MessageDetails `json:"results,omitempty" url:"results,omitempty"`

	_rawJSON json.RawMessage
}

func (l *ListMessagesResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListMessagesResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListMessagesResponse(value)
	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListMessagesResponse) String() string {
	if len(l._rawJSON) > 0 {
		if value, err := core.StringifyJSON(l._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

type MessageDetails struct {
	// A unique identifier associated with the message you wish to retrieve (results from a send).
	Id string `json:"id" url:"id"`
	// The current status of the message.
	Status MessageStatus `json:"status,omitempty" url:"status,omitempty"`
	// A UTC timestamp at which Courier received the message request. Stored as a millisecond representation of the Unix epoch.
	Enqueued int64 `json:"enqueued" url:"enqueued"`
	// A UTC timestamp at which Courier passed the message to the Integration provider. Stored as a millisecond representation of the Unix epoch.
	Sent int64 `json:"sent" url:"sent"`
	// A UTC timestamp at which the Integration provider delivered the message. Stored as a millisecond representation of the Unix epoch.
	Delivered int64 `json:"delivered" url:"delivered"`
	// A UTC timestamp at which the recipient opened a message for the first time. Stored as a millisecond representation of the Unix epoch.
	Opened int64 `json:"opened" url:"opened"`
	// A UTC timestamp at which the recipient clicked on a tracked link for the first time. Stored as a millisecond representation of the Unix epoch.
	Clicked int64 `json:"clicked" url:"clicked"`
	// A unique identifier associated with the recipient of the delivered message.
	Recipient string `json:"recipient" url:"recipient"`
	// A unique identifier associated with the event of the delivered message.
	Event string `json:"event" url:"event"`
	// A unique identifier associated with the notification of the delivered message.
	Notification string `json:"notification" url:"notification"`
	// A message describing the error that occurred.
	Error *string `json:"error,omitempty" url:"error,omitempty"`
	// The reason for the current status of the message.
	Reason *Reason `json:"reason,omitempty" url:"reason,omitempty"`

	_rawJSON json.RawMessage
}

func (m *MessageDetails) UnmarshalJSON(data []byte) error {
	type unmarshaler MessageDetails
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = MessageDetails(value)
	m._rawJSON = json.RawMessage(data)
	return nil
}

func (m *MessageDetails) String() string {
	if len(m._rawJSON) > 0 {
		if value, err := core.StringifyJSON(m._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(m); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", m)
}

type MessageHistoryResponse struct {
	Results []map[string]interface{} `json:"results,omitempty" url:"results,omitempty"`

	_rawJSON json.RawMessage
}

func (m *MessageHistoryResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler MessageHistoryResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = MessageHistoryResponse(value)
	m._rawJSON = json.RawMessage(data)
	return nil
}

func (m *MessageHistoryResponse) String() string {
	if len(m._rawJSON) > 0 {
		if value, err := core.StringifyJSON(m._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(m); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", m)
}

type RenderOutputResponse struct {
	// An array of render output of a previously sent message.
	Results []*RenderOutput `json:"results,omitempty" url:"results,omitempty"`

	_rawJSON json.RawMessage
}

func (r *RenderOutputResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler RenderOutputResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RenderOutputResponse(value)
	r._rawJSON = json.RawMessage(data)
	return nil
}

func (r *RenderOutputResponse) String() string {
	if len(r._rawJSON) > 0 {
		if value, err := core.StringifyJSON(r._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}
