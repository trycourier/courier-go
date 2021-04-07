package courier

import (
	"context"
	"encoding/json"
)

type AutomationStep struct {
	Action           string      `json:"action"`
	Brand            string      `json:"brand,omitempty"`
	CancelationToken string      `json:"cancelation_token,omitempty"`
	Data             interface{} `json:"data,omitempty"`
	Duration         string      `json:"duration,omitempty"`
	If               string      `json:"if,omitempty"`
	List             string      `json:"list,omitempty"`
	Override         interface{} `json:"override,omitempty"`
	Profile          interface{} `json:"profile,omitempty"`
	Recipient        string      `json:"recipient,omitempty"`
	Ref              string      `json:"ref,omitempty"`
	Template         string      `json:"template,omitempty"`
	Until            string      `json:"until,omitempty"`
}

type Automation struct {
	CancelationToken string           `json:"cancelation_token,omitempty"`
	Steps            []AutomationStep `json:"steps"`
}

type AutomationInvokeBody struct {
	Automation Automation  `json:"automation"`
	Brand      string      `json:"brand,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Profile    interface{} `json:"profile,omitempty"`
	Recipient  string      `json:"recipient,omitempty"`
	Template   string      `json:"template,omitempty"`
}

type AutomationTemplateInvokeBody struct {
	Brand     string      `json:"brand,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Profile   interface{} `json:"profile,omitempty"`
	Recipient string      `json:"recipient,omitempty"`
	Template  string      `json:"template,omitempty"`
}

// Calls the POST /automations/invoke endpoint of the Courier API
func (c *Client) InvokeAutomation(ctx context.Context, body interface{}) (string, error) {
	bodyMap, err := toJSONMap(body)
	if err != nil {
		return "", err
	}

	response, err := c.API.SendRequestWithMaps(ctx, "POST", "/automations/invoke", bodyMap)
	if err != nil {
		return "", err
	}

	var runID string
	err = json.Unmarshal(response["runId"], &runID)
	if err != nil {
		return "", err
	}
	return runID, nil
}

// Calls the POST /automations/:templateId/invoke endpoint of the Courier API
func (c *Client) InvokeAutomationTemplate(ctx context.Context, templateId string, body interface{}) (string, error) {
	bodyMap, err := toJSONMap(body)
	if err != nil {
		return "", err
	}

	response, err := c.API.SendRequestWithMaps(ctx, "POST", "/automations/"+templateId+"/invoke", bodyMap)
	if err != nil {
		return "", err
	}

	var runID string
	err = json.Unmarshal(response["runId"], &runID)
	if err != nil {
		return "", err
	}
	return runID, nil
}
