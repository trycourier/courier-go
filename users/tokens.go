// This file was auto-generated by Fern from our API Definition.

package users

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/trycourier/courier-go/v3/core"
)

// A list of tokens registered with the user.
type GetAllTokensResponse = []*UserToken

type GetUserTokenResponse struct {
	// Full body of the token. Must match token in URL.
	Token       *string     `json:"token,omitempty" url:"token,omitempty"`
	ProviderKey ProviderKey `json:"provider_key,omitempty" url:"provider_key,omitempty"`
	// ISO 8601 formatted date the token expires. Defaults to 2 months. Set to false to disable expiration.
	ExpiryDate *ExpiryDate `json:"expiry_date,omitempty" url:"expiry_date,omitempty"`
	// Properties sent to the provider along with the token
	Properties interface{} `json:"properties,omitempty" url:"properties,omitempty"`
	// Information about the device the token is associated with.
	Device *Device `json:"device,omitempty" url:"device,omitempty"`
	// Information about the device the token is associated with.
	Tracking *Tracking    `json:"tracking,omitempty" url:"tracking,omitempty"`
	Status   *TokenStatus `json:"status,omitempty" url:"status,omitempty"`
	// The reason for the token status.
	StatusReason *string `json:"status_reason,omitempty" url:"status_reason,omitempty"`

	_rawJSON json.RawMessage
}

func (g *GetUserTokenResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler GetUserTokenResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetUserTokenResponse(value)
	g._rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetUserTokenResponse) String() string {
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

type PatchUserTokenOpts struct {
	Patch []*PatchOperation `json:"patch,omitempty" url:"patch,omitempty"`

	_rawJSON json.RawMessage
}

func (p *PatchUserTokenOpts) UnmarshalJSON(data []byte) error {
	type unmarshaler PatchUserTokenOpts
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PatchUserTokenOpts(value)
	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PatchUserTokenOpts) String() string {
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

type UserToken struct {
	// Full body of the token. Must match token in URL.
	Token       *string     `json:"token,omitempty" url:"token,omitempty"`
	ProviderKey ProviderKey `json:"provider_key,omitempty" url:"provider_key,omitempty"`
	// ISO 8601 formatted date the token expires. Defaults to 2 months. Set to false to disable expiration.
	ExpiryDate *ExpiryDate `json:"expiry_date,omitempty" url:"expiry_date,omitempty"`
	// Properties sent to the provider along with the token
	Properties interface{} `json:"properties,omitempty" url:"properties,omitempty"`
	// Information about the device the token is associated with.
	Device *Device `json:"device,omitempty" url:"device,omitempty"`
	// Information about the device the token is associated with.
	Tracking *Tracking `json:"tracking,omitempty" url:"tracking,omitempty"`

	_rawJSON json.RawMessage
}

func (u *UserToken) UnmarshalJSON(data []byte) error {
	type unmarshaler UserToken
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UserToken(value)
	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UserToken) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}
