// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/trycourier/courier-go/v3/core"
)

type MergeProfileRequest struct {
	Profile map[string]interface{} `json:"profile,omitempty" url:"profile,omitempty"`
}

type GetListSubscriptionsRequest struct {
	// A unique identifier that allows for fetching the next set of message statuses.
	Cursor *string `json:"-" url:"cursor,omitempty"`
}

type ReplaceProfileRequest struct {
	Profile map[string]interface{} `json:"profile,omitempty" url:"profile,omitempty"`
}

type DeleteListSubscriptionResponse struct {
	status string

	_rawJSON json.RawMessage
}

func (d *DeleteListSubscriptionResponse) Status() string {
	return d.status
}

func (d *DeleteListSubscriptionResponse) UnmarshalJSON(data []byte) error {
	type embed DeleteListSubscriptionResponse
	var unmarshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*d = DeleteListSubscriptionResponse(unmarshaler.embed)
	d.status = "SUCCESS"
	d._rawJSON = json.RawMessage(data)
	return nil
}

func (d *DeleteListSubscriptionResponse) MarshalJSON() ([]byte, error) {
	type embed DeleteListSubscriptionResponse
	var marshaler = struct {
		embed
		Status string `json:"status"`
	}{
		embed:  embed(*d),
		Status: "SUCCESS",
	}
	return json.Marshal(marshaler)
}

func (d *DeleteListSubscriptionResponse) String() string {
	if len(d._rawJSON) > 0 {
		if value, err := core.StringifyJSON(d._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

type GetListSubscriptionsResponse struct {
	Paging *Paging `json:"paging,omitempty" url:"paging,omitempty"`
	// An array of lists
	Results []*GetListSubscriptionsList `json:"results,omitempty" url:"results,omitempty"`

	_rawJSON json.RawMessage
}

func (g *GetListSubscriptionsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler GetListSubscriptionsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetListSubscriptionsResponse(value)
	g._rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetListSubscriptionsResponse) String() string {
	if len(g._rawJSON) > 0 {
		if value, err := core.StringifyJSON(g._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}

type MergeProfileResponse struct {
	status string

	_rawJSON json.RawMessage
}

func (m *MergeProfileResponse) Status() string {
	return m.status
}

func (m *MergeProfileResponse) UnmarshalJSON(data []byte) error {
	type embed MergeProfileResponse
	var unmarshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*m = MergeProfileResponse(unmarshaler.embed)
	m.status = "SUCCESS"
	m._rawJSON = json.RawMessage(data)
	return nil
}

func (m *MergeProfileResponse) MarshalJSON() ([]byte, error) {
	type embed MergeProfileResponse
	var marshaler = struct {
		embed
		Status string `json:"status"`
	}{
		embed:  embed(*m),
		Status: "SUCCESS",
	}
	return json.Marshal(marshaler)
}

func (m *MergeProfileResponse) String() string {
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

type ProfileGetResponse struct {
	Profile     map[string]interface{} `json:"profile,omitempty" url:"profile,omitempty"`
	Preferences *RecipientPreferences  `json:"preferences,omitempty" url:"preferences,omitempty"`

	_rawJSON json.RawMessage
}

func (p *ProfileGetResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ProfileGetResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = ProfileGetResponse(value)
	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *ProfileGetResponse) String() string {
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

type ReplaceProfileResponse struct {
	status string

	_rawJSON json.RawMessage
}

func (r *ReplaceProfileResponse) Status() string {
	return r.status
}

func (r *ReplaceProfileResponse) UnmarshalJSON(data []byte) error {
	type embed ReplaceProfileResponse
	var unmarshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*r = ReplaceProfileResponse(unmarshaler.embed)
	r.status = "SUCCESS"
	r._rawJSON = json.RawMessage(data)
	return nil
}

func (r *ReplaceProfileResponse) MarshalJSON() ([]byte, error) {
	type embed ReplaceProfileResponse
	var marshaler = struct {
		embed
		Status string `json:"status"`
	}{
		embed:  embed(*r),
		Status: "SUCCESS",
	}
	return json.Marshal(marshaler)
}

func (r *ReplaceProfileResponse) String() string {
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

type SubscribeToListsRequest struct {
	Lists []*SubscribeToListsRequestListObject `json:"lists,omitempty" url:"lists,omitempty"`

	_rawJSON json.RawMessage
}

func (s *SubscribeToListsRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler SubscribeToListsRequest
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SubscribeToListsRequest(value)
	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SubscribeToListsRequest) String() string {
	if len(s._rawJSON) > 0 {
		if value, err := core.StringifyJSON(s._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}

type SubscribeToListsResponse struct {
	status string

	_rawJSON json.RawMessage
}

func (s *SubscribeToListsResponse) Status() string {
	return s.status
}

func (s *SubscribeToListsResponse) UnmarshalJSON(data []byte) error {
	type embed SubscribeToListsResponse
	var unmarshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*s = SubscribeToListsResponse(unmarshaler.embed)
	s.status = "SUCCESS"
	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SubscribeToListsResponse) MarshalJSON() ([]byte, error) {
	type embed SubscribeToListsResponse
	var marshaler = struct {
		embed
		Status string `json:"status"`
	}{
		embed:  embed(*s),
		Status: "SUCCESS",
	}
	return json.Marshal(marshaler)
}

func (s *SubscribeToListsResponse) String() string {
	if len(s._rawJSON) > 0 {
		if value, err := core.StringifyJSON(s._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}
