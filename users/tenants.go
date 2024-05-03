// This file was auto-generated by Fern from our API Definition.

package users

import (
	json "encoding/json"
	fmt "fmt"
	v3 "github.com/trycourier/courier-go/v3"
	core "github.com/trycourier/courier-go/v3/core"
)

type AddUserToSingleTenantsParams struct {
	Profile map[string]interface{} `json:"profile,omitempty" url:"profile,omitempty"`
}

type AddUserToMultipleTenantsParams struct {
	Tenants []*v3.UserTenantAssociation `json:"tenants,omitempty" url:"tenants,omitempty"`
}

type ListTenantsForUserParams struct {
	// The number of accounts to return
	// (defaults to 20, maximum value of 100)
	Limit *int `json:"-" url:"limit,omitempty"`
	// Continue the pagination with the next cursor
	Cursor *string `json:"-" url:"cursor,omitempty"`
}

type ListTenantsForUserResponse struct {
	Items []*v3.UserTenantAssociation `json:"items,omitempty" url:"items,omitempty"`
	// Set to true when there are more pages that can be retrieved.
	HasMore bool `json:"has_more" url:"has_more"`
	// A url that may be used to generate these results.
	Url string `json:"url" url:"url"`
	// A url that may be used to generate fetch the next set of results.
	// Defined only when `has_more` is set to true
	NextUrl *string `json:"next_url,omitempty" url:"next_url,omitempty"`
	// A pointer to the next page of results. Defined
	// only when `has_more` is set to true
	Cursor *string `json:"cursor,omitempty" url:"cursor,omitempty"`
	// Always set to `list`. Represents the type of this object.
	type_ string

	_rawJSON json.RawMessage
}

func (l *ListTenantsForUserResponse) Type() string {
	return l.type_
}

func (l *ListTenantsForUserResponse) UnmarshalJSON(data []byte) error {
	type embed ListTenantsForUserResponse
	var unmarshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*l = ListTenantsForUserResponse(unmarshaler.embed)
	l.type_ = "list"
	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListTenantsForUserResponse) MarshalJSON() ([]byte, error) {
	type embed ListTenantsForUserResponse
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*l),
		Type:  "list",
	}
	return json.Marshal(marshaler)
}

func (l *ListTenantsForUserResponse) String() string {
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
