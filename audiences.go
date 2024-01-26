// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/trycourier/courier-go/v3/core"
)

type AudiencesListParams struct {
	// A unique identifier that allows for fetching the next set of audiences
	Cursor *string `json:"-"`
}

type AudienceMembersListParams struct {
	// A unique identifier that allows for fetching the next set of members
	Cursor *string `json:"-"`
}

type Audience struct {
	// A unique identifier representing the audience_id
	Id string `json:"id"`
	// The name of the audience
	Name string `json:"name"`
	// A description of the audience
	Description string  `json:"description"`
	Filter      *Filter `json:"filter,omitempty"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`

	_rawJSON json.RawMessage
}

func (a *Audience) UnmarshalJSON(data []byte) error {
	type unmarshaler Audience
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = Audience(value)
	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *Audience) String() string {
	if len(a._rawJSON) > 0 {
		if value, err := core.StringifyJSON(a._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

type AudienceListResponse struct {
	Items  []*Audience `json:"items,omitempty"`
	Paging *Paging     `json:"paging,omitempty"`

	_rawJSON json.RawMessage
}

func (a *AudienceListResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler AudienceListResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AudienceListResponse(value)
	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *AudienceListResponse) String() string {
	if len(a._rawJSON) > 0 {
		if value, err := core.StringifyJSON(a._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

type AudienceMemberListResponse struct {
	Items  []*AudienceMember `json:"items,omitempty"`
	Paging *Paging           `json:"paging,omitempty"`

	_rawJSON json.RawMessage
}

func (a *AudienceMemberListResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler AudienceMemberListResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AudienceMemberListResponse(value)
	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *AudienceMemberListResponse) String() string {
	if len(a._rawJSON) > 0 {
		if value, err := core.StringifyJSON(a._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

type AudienceUpdateResponse struct {
	Audience *Audience `json:"audience,omitempty"`

	_rawJSON json.RawMessage
}

func (a *AudienceUpdateResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler AudienceUpdateResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AudienceUpdateResponse(value)
	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *AudienceUpdateResponse) String() string {
	if len(a._rawJSON) > 0 {
		if value, err := core.StringifyJSON(a._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

type Filter struct {
	typeName           string
	SingleFilterConfig *SingleFilterConfig
	NestedFilterConfig *NestedFilterConfig
}

func NewFilterFromSingleFilterConfig(value *SingleFilterConfig) *Filter {
	return &Filter{typeName: "singleFilterConfig", SingleFilterConfig: value}
}

func NewFilterFromNestedFilterConfig(value *NestedFilterConfig) *Filter {
	return &Filter{typeName: "nestedFilterConfig", NestedFilterConfig: value}
}

func (f *Filter) UnmarshalJSON(data []byte) error {
	valueSingleFilterConfig := new(SingleFilterConfig)
	if err := json.Unmarshal(data, &valueSingleFilterConfig); err == nil {
		f.typeName = "singleFilterConfig"
		f.SingleFilterConfig = valueSingleFilterConfig
		return nil
	}
	valueNestedFilterConfig := new(NestedFilterConfig)
	if err := json.Unmarshal(data, &valueNestedFilterConfig); err == nil {
		f.typeName = "nestedFilterConfig"
		f.NestedFilterConfig = valueNestedFilterConfig
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, f)
}

func (f Filter) MarshalJSON() ([]byte, error) {
	switch f.typeName {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", f.typeName, f)
	case "singleFilterConfig":
		return json.Marshal(f.SingleFilterConfig)
	case "nestedFilterConfig":
		return json.Marshal(f.NestedFilterConfig)
	}
}

type FilterVisitor interface {
	VisitSingleFilterConfig(*SingleFilterConfig) error
	VisitNestedFilterConfig(*NestedFilterConfig) error
}

func (f *Filter) Accept(visitor FilterVisitor) error {
	switch f.typeName {
	default:
		return fmt.Errorf("invalid type %s in %T", f.typeName, f)
	case "singleFilterConfig":
		return visitor.VisitSingleFilterConfig(f.SingleFilterConfig)
	case "nestedFilterConfig":
		return visitor.VisitNestedFilterConfig(f.NestedFilterConfig)
	}
}

type AudienceUpdateParams struct {
	// The name of the audience
	Name *string `json:"name,omitempty"`
	// A description of the audience
	Description *string `json:"description,omitempty"`
	Filter      *Filter `json:"filter,omitempty"`
}