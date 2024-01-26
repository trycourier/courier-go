// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/trycourier/courier-go/v3/core"
)

type AutomationAdHocInvokeParams struct {
	Brand      *string                `json:"brand,omitempty"`
	Data       map[string]interface{} `json:"data,omitempty"`
	Profile    *Profile               `json:"profile,omitempty"`
	Recipient  *string                `json:"recipient,omitempty"`
	Template   *string                `json:"template,omitempty"`
	Automation *Automation            `json:"automation,omitempty"`

	_rawJSON json.RawMessage
}

func (a *AutomationAdHocInvokeParams) UnmarshalJSON(data []byte) error {
	type unmarshaler AutomationAdHocInvokeParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AutomationAdHocInvokeParams(value)
	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *AutomationAdHocInvokeParams) String() string {
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

type AutomationInvokeParams struct {
	Brand     *string                `json:"brand,omitempty"`
	Data      map[string]interface{} `json:"data,omitempty"`
	Profile   *Profile               `json:"profile,omitempty"`
	Recipient *string                `json:"recipient,omitempty"`
	Template  *string                `json:"template,omitempty"`

	_rawJSON json.RawMessage
}

func (a *AutomationInvokeParams) UnmarshalJSON(data []byte) error {
	type unmarshaler AutomationInvokeParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AutomationInvokeParams(value)
	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *AutomationInvokeParams) String() string {
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

type AutomationInvokeResponse struct {
	RunId string `json:"runId"`

	_rawJSON json.RawMessage
}

func (a *AutomationInvokeResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler AutomationInvokeResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AutomationInvokeResponse(value)
	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *AutomationInvokeResponse) String() string {
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
