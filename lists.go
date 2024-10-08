// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/trycourier/courier-go/v3/core"
)

type AddSubscribersToList struct {
	Recipients []*PutSubscriptionsRecipient `json:"recipients,omitempty" url:"recipients,omitempty"`
}

type GetSubscriptionForListRequest struct {
	// A unique identifier that allows for fetching the next set of list subscriptions
	Cursor *string `json:"-" url:"cursor,omitempty"`
}

type GetAllListsRequest struct {
	// A unique identifier that allows for fetching the next page of lists.
	Cursor *string `json:"-" url:"cursor,omitempty"`
	// "A pattern used to filter the list items returned. Pattern types supported: exact match on `list_id` or a pattern of one or more pattern parts. you may replace a part with either: `*` to match all parts in that position, or `**` to signify a wildcard `endsWith` pattern match."
	Pattern *string `json:"-" url:"pattern,omitempty"`
}

type SubscribeUserToListRequest struct {
	Preferences *RecipientPreferences `json:"preferences,omitempty" url:"preferences,omitempty"`
}

type RecipientPreferences struct {
	Categories    *NotificationPreferences `json:"categories,omitempty" url:"categories,omitempty"`
	Notifications *NotificationPreferences `json:"notifications,omitempty" url:"notifications,omitempty"`

	_rawJSON json.RawMessage
}

func (r *RecipientPreferences) UnmarshalJSON(data []byte) error {
	type unmarshaler RecipientPreferences
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RecipientPreferences(value)
	r._rawJSON = json.RawMessage(data)
	return nil
}

func (r *RecipientPreferences) String() string {
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

type List struct {
	Id      string `json:"id" url:"id"`
	Name    string `json:"name" url:"name"`
	Created *int   `json:"created,omitempty" url:"created,omitempty"`
	Updated *int   `json:"updated,omitempty" url:"updated,omitempty"`

	_rawJSON json.RawMessage
}

func (l *List) UnmarshalJSON(data []byte) error {
	type unmarshaler List
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = List(value)
	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *List) String() string {
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

type ListGetAllResponse struct {
	Paging *Paging `json:"paging,omitempty" url:"paging,omitempty"`
	Items  []*List `json:"items,omitempty" url:"items,omitempty"`

	_rawJSON json.RawMessage
}

func (l *ListGetAllResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListGetAllResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListGetAllResponse(value)
	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListGetAllResponse) String() string {
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

type ListGetSubscriptionsResponse struct {
	Paging *Paging                      `json:"paging,omitempty" url:"paging,omitempty"`
	Items  []*ListSubscriptionRecipient `json:"items,omitempty" url:"items,omitempty"`

	_rawJSON json.RawMessage
}

func (l *ListGetSubscriptionsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListGetSubscriptionsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListGetSubscriptionsResponse(value)
	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListGetSubscriptionsResponse) String() string {
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

type ListPutParams struct {
	Name        string                `json:"name" url:"name"`
	Preferences *RecipientPreferences `json:"preferences,omitempty" url:"preferences,omitempty"`

	_rawJSON json.RawMessage
}

func (l *ListPutParams) UnmarshalJSON(data []byte) error {
	type unmarshaler ListPutParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListPutParams(value)
	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListPutParams) String() string {
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

type PutSubscriptionsRecipient struct {
	RecipientId string                `json:"recipientId" url:"recipientId"`
	Preferences *RecipientPreferences `json:"preferences,omitempty" url:"preferences,omitempty"`

	_rawJSON json.RawMessage
}

func (p *PutSubscriptionsRecipient) UnmarshalJSON(data []byte) error {
	type unmarshaler PutSubscriptionsRecipient
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PutSubscriptionsRecipient(value)
	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PutSubscriptionsRecipient) String() string {
	if len(p._rawJSON) > 0 {
		if value, err := core.StringifyJSON(p._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type SubscribeUsersToListRequest struct {
	Recipients []*PutSubscriptionsRecipient `json:"recipients,omitempty" url:"recipients,omitempty"`
}
